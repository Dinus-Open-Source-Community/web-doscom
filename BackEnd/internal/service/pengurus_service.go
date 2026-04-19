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

func (p *PengurusService) RolePositionAuthorization(ctx context.Context, idParams, currentUserID int, userRole string) (string, error) {
	// check if data exist
	userValid, err := p.PengurusModel.GetPengurusById(ctx, idParams)
	if err != nil {
		return "", err
	}

	// white list role
	actor, ok := constants.RoleGroup[userRole]
	if !ok {
		return "", fmt.Errorf("role not valid")
	}

	actorRole := actor.Role
	// SuperAdmin bypass
	if actorRole == constants.RoleAdmin {
		return actorRole, nil
	}

	// cek role pengurus
	if actorRole == constants.RolePengurus && currentUserID != idParams {
		return "", fmt.Errorf("You are not allowed to update this data")
	}

	// whitelist position
	targetPositionRole := constants.ValidPosition[userValid.Position]

	// ckeck user level role
	actorLevel := constants.RoleLevel[actorRole]
	targetLevel := constants.RoleLevel[targetPositionRole]

	if actorLevel >= targetLevel && currentUserID != idParams {
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
	divisi, validPosition, err := utils.SetDivitionAndPositionByRole(dataPengurus.Position, dataPengurus.Divisi, userRole)
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
		if dataPengurus.UserID == 0 {
			dataPengurus.UserID = currentUserID
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
	roleUser, err := p.RolePositionAuthorization(ctx, idParams, currentUserID, userRole)
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
		// Pastikan kita ambil data pengurus dulu untuk tahu UserID-nya yang asli
		targetPengurus, err := p.PengurusModel.GetPengurusById(ctx, idParams)
		if err != nil {
			return nil, err
		}

		// Pastikan juga bungkusan fileUpload menggunakan UserID asli dari tabel users (angka 2)
		fileUpload.UserID = uint(targetPengurus.UserID)

		// update file upload and gallery
		now := time.Now()
		gallery := &model.GalleryInsert{
			IDUsers:     targetPengurus.UserID, // Correctly link to the UserID, not Pengurus ID
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

func (p *PengurusService) GetAllPengurusBaseOnDivision(
	ctx context.Context,
	userRole, divisi string,
) ([]model.PengurusResponse, error) {
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

func (p *PengurusService) GetAllPengurusByDivision(
	ctx context.Context,
	division string,
) ([]model.PengurusResponse, error) {
	if division == "" {
		return nil, fmt.Errorf("division must be provided")
	}

	pengurusResponse, err := p.PengurusModel.GetAllPengurusByDivisi(ctx, division)
	if err != nil {
		return nil, fmt.Errorf("terjadi error ketika ambil data %w", err)
	}

	dataPengurus := make([]model.PengurusResponse, 0, len(pengurusResponse))
	for _, data := range pengurusResponse {
		dataPengurus = append(dataPengurus, model.PengurusResponse{
			ID:       data.ID,
			PhotoURL: data.PhotoURL,
			Email:    data.Email,
			Divisi:   data.Divisi,
			Name:     data.Name,
			Position: data.Position,
			Sosmed:   data.Sosmed,
			Period:   data.Period,
		})
	}

	return dataPengurus, nil
}

func (p *PengurusService) GetPengurusByID(ctx context.Context, id int, userRole string, userID int) (model.PengurusResponse, error) {
	validRole, err := constants.GetRoleInfo(userRole)
	if err != nil {
		return model.PengurusResponse{}, fmt.Errorf("role not valid")
	}

	pengurusResponse, err := p.PengurusModel.GetPengurusById(ctx, id)
	if err != nil {
		return model.PengurusResponse{}, fmt.Errorf("error while getting the data %w", err)
	}

	switch validRole.Role {
	case constants.RoleKoordinator:
		if validRole.Divisi != pengurusResponse.Divisi {
			return model.PengurusResponse{}, fmt.Errorf("you can not see other division bro %w", err)
		}
	case constants.RolePengurus:
		if userID != pengurusResponse.ID {
			return model.PengurusResponse{}, fmt.Errorf("you can't see other data bro")
		}
	default:
		return model.PengurusResponse{}, fmt.Errorf("role not valid")
	}

	pengurusDataResponse := model.PengurusResponse{
		ID:       pengurusResponse.ID,
		PhotoURL: pengurusResponse.PhotoURL,
		Email:    pengurusResponse.Email,
		Divisi:   pengurusResponse.Divisi,
		Name:     pengurusResponse.Name,
		Position: pengurusResponse.Position,
		Sosmed:   pengurusResponse.Sosmed,
		Period:   pengurusResponse.Period,
	}
	return pengurusDataResponse, nil
}

func (p *PengurusService) DeletePengurusById(ctx context.Context, idPengurus int, userRole string) error {
	role, ok := constants.RoleGroup[userRole]
	if !ok {
		return fmt.Errorf("role not valid")
	}

	targetPengurus, err := p.PengurusModel.GetPengurusById(ctx, idPengurus)
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
