package service

import (
	"context"
	"fmt"
	"time"
	"web_doscom/internal/constants"
	"web_doscom/internal/database/model"
	"web_doscom/internal/utils"

	"github.com/mitchellh/mapstructure"
)

type PengurusService struct {
	PengurusModel  *model.PengurusModel
	GalleryService *GalleryService
}

func NewPengurusService(m *model.PengurusModel, g *GalleryService) *PengurusService {
	return &PengurusService{
		PengurusModel:  m,
		GalleryService: g,
	}
}

func (p *PengurusService) RolePositionAuthorization(idParams, currentUserID int, userRole string) (string, error) {
	// white list role
	actor, ok := constants.RoleGroup[userRole]
	if !ok {
		return "", fmt.Errorf("role not valid")
	}

	actorRole := actor.Role
	// cek role pengurus
	if actorRole == constants.RolePengurus && currentUserID != idParams {
		return "", fmt.Errorf("You are not allowed to update this data")
	}

	userValid, err := p.PengurusModel.GetPengurusById(idParams)
	if err != nil {
		return "", err
	}
	// whitelist position
	targetPositionRole := constants.ValidPosition[userValid.Position]

	// ckeck user level role
	actorLevel := constants.RoleLevel[actorRole]
	targetLevel := constants.RoleLevel[targetPositionRole]

	if actorLevel <= targetLevel && currentUserID != idParams {
		return "", fmt.Errorf("You are not allowed to update this data")
	}

	return actorRole, nil
}

func (p *PengurusService) CreatePengurus(
	ctx context.Context,
	currentUserID int,
	userRole string,
	dataPengurus *model.RegisterPengurusRequest,
	fileUpload *model.UploadFileRequest,
) (*model.PengurusResponse, error) {

	// auto assign position and divisi
	divisi, validPosition, err := utils.SetDivitionAndPositionByRole(dataPengurus.Position, userRole)
	if err != nil {
		return nil, err
	}

	// auto assign user_id
	roleUser, ok := constants.RoleGroup[userRole]
	if !ok {
		return nil, fmt.Errorf("role not valid")
	}
	switch roleUser.Role {
	case constants.RolePengurus:
		dataPengurus.UserID = currentUserID

	case constants.RoleKoordinator:
		if dataPengurus.UserID == 0 {
			dataPengurus.UserID = currentUserID
		}
	case constants.RoleAdmin:
		if dataPengurus.UserID == 0 && validPosition != "ketum" {
			return nil, fmt.Errorf("userId must be given")
		}
	}

	// filter field insert by role
	data := model.PengurusPatch{
		Email:    dataPengurus.Email,
		Divisi:   divisi,
		Name:     dataPengurus.Name,
		Position: validPosition,
		Sosmed:   dataPengurus.Sosmed,
		Period:   dataPengurus.Period,
		PhotoURL: dataPengurus.PhotoURL,
	}
	fillableFields, err := utils.FilterRoleFieldPermission(userRole, &data)
	if err != nil {
		return nil, err
	}

	var finalData *model.Pengurus
	if err := mapstructure.Decode(fillableFields, &finalData); err != nil {
		return nil, fmt.Errorf("failed to decode data")
	}

	finalData.UserID = dataPengurus.UserID
	if finalData.Sosmed == "" {
		finalData.Sosmed = "instagram"
	}
	// upload photo
	if _, canUploadPhoto := fillableFields["photo_url"]; canUploadPhoto {
		now := time.Now()
		gallery := &model.GalleryInsert{
			IDUsers:     finalData.UserID,
			GalleryName: "foto profil pengurus",
			GalleryType: "pengurus",
			Description: "foto identitas diri yang mewakili pengurus doscom",
			EventDate: time.Date(
				now.Year(),
				now.Month(),
				now.Day(), 0, 0, 0, 0, time.UTC,
			),
		}

		_, fileURL, err := p.GalleryService.InsertGalleryAndFileUpload(
			ctx,
			gallery,
			fileUpload,
		)

		if err != nil {
			return nil, err
		}

		finalData.PhotoURL = fileURL
	}

	// insert data
	if err := p.PengurusModel.InsertPengurus(finalData); err != nil {
		return nil, err
	}

	return &model.PengurusResponse{
		ID:       finalData.ID,
		PhotoURL: finalData.PhotoURL,
		Email:    finalData.Email,
		Divisi:   finalData.Divisi,
		Name:     finalData.Name,
		Position: finalData.Position,
		Sosmed:   finalData.Sosmed,
		Period:   finalData.Period,
	}, nil

}

func (p *PengurusService) UpdateDataPengurus(
	ctx context.Context,
	idParams int,
	currentUserID int,
	userRole string,
	dataPengurus *model.PengurusPatch,
	fileUpload *model.UploadFileRequest,
) (*model.PengurusResponse, error) {

	var (
		updatedPengurus *model.PengurusResponse
		error           error
	)

	// authorization check for update data
	roleUser, err := p.RolePositionAuthorization(idParams, currentUserID, userRole)
	if err != nil {
		return nil, err
	}

	// filter fileld update by role and update profile
	editableFields, err := utils.FilterRoleFieldPermission(userRole, dataPengurus)
	if err != nil {
		return nil, err
	}

	canUpdatePhoto := roleUser == constants.RoleAdmin || (roleUser == constants.RolePengurus && currentUserID == idParams)
	if fileUpload != nil {
		if !canUpdatePhoto {
			return nil, fmt.Errorf("koordinator tidak dapat memperbarui foto pengurus")
		}
		// update file upload and gallery
		now := time.Now()
		gallery := &model.GalleryInsert{
			IDUsers:     idParams,
			GalleryName: "foto profil pengurus",
			GalleryType: "pengurus",
			Description: "foto identitas diri yang mewakili pengurus doscom",
			EventDate: time.Date(
				now.Year(),
				now.Month(),
				now.Day(), 0, 0, 0, 0, time.UTC,
			),
		}
		_, fileURL, err := p.GalleryService.InsertGalleryAndFileUpload(
			ctx,
			gallery,
			fileUpload,
		)
		if err != nil {
			return nil, err
		}
		editableFields["photo_url"] = fileURL

	}

	// update data pengurus
	updatedPengurus, error = p.PengurusModel.UpdatePengurusPartial(idParams, editableFields)
	if error != nil {
		return nil, error
	}

	return updatedPengurus, nil
}

func (p *PengurusService) GetAllPengurusBaseOnDivision(ctx context.Context, userRole, divisi string) ([]model.PengurusResponse, error) {
	// cek role
	role, ok := constants.RoleGroup[userRole]
	if !ok {
		return nil, fmt.Errorf("role not valid")
	}

	var (
		pengurusResponse []model.PengurusResponse
		err              error
	)
	switch role.Role {
	case constants.RoleAdmin:
		// get all data from any division
		pengurusResponse, err = p.PengurusModel.GetPengurusByDivisi(ctx, divisi)
		if err != nil {
			return nil, err
		}
	case constants.RoleKoordinator:
		// can only get data from same division
		pengurusResponse, err = p.PengurusModel.GetPengurusByDivisi(ctx, role.Divisi)
		if err != nil {
			return nil, err
		}
	case constants.RolePengurus:
		// can't get another pengurus data
		return nil, fmt.Errorf("forbidden, you can't see other data")
	default:
		return nil, fmt.Errorf("role not valid")
	}

	return pengurusResponse, nil

}

func (p *PengurusService) DeletePengurusById(ctx context.Context, idPengurus int, userRole string) error {
	role, ok := constants.RoleGroup[userRole]
	if !ok {
		return fmt.Errorf("role not valid")
	}

	targetPengurus, err := p.PengurusModel.GetPengurusById(idPengurus)
	if err != nil {
		return err
	}

	switch role.Role {
	case constants.RoleAdmin:
		if err := p.PengurusModel.DeletePengurus(ctx, idPengurus); err != nil {
			return err
		}
	case constants.RoleKoordinator:
		if role.Divisi != targetPengurus.Divisi {
			return fmt.Errorf("you can not delete data from other division, %w", err)
		}
		if err := p.PengurusModel.DeletePengurus(ctx, idPengurus); err != nil {
			return err
		}
	case constants.RolePengurus:
		return fmt.Errorf("you can't delete your own data %w", err)
	default:
		return fmt.Errorf("role not valid")
	}

	return nil
}
