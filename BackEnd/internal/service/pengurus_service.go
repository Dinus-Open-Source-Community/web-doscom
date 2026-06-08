package service

import (
	"context"
	"fmt"
	"log"
	"time"
	"web_doscom/internal/authorization"
	pengurusAuthorization "web_doscom/internal/authorization/pengurus"
	"web_doscom/internal/constants"
	"web_doscom/internal/database/model/dto"
	"web_doscom/internal/database/model/entity"
	"web_doscom/internal/utils"

	"github.com/mitchellh/mapstructure"
	"gorm.io/gorm"
)

type PengurusService struct {
	DB                  *gorm.DB
	PengurusModel       *entity.PengurusModel
	PengurusSosmedModel *entity.PengurusSosmedModel
	GalleryService      *GalleryService
}

func NewPengurusService(m *entity.PengurusModel, p *entity.PengurusSosmedModel, g *GalleryService, d *gorm.DB) *PengurusService {
	return &PengurusService{
		DB:                  d,
		PengurusModel:       m,
		PengurusSosmedModel: p,
		GalleryService:      g,
	}
}

func (p *PengurusService) UpdatePengurusSosmed(ctx context.Context, pengurusID int, sosmedUrl []string) ([]dto.PengurusSosmedResponse, error) {

	if len(sosmedUrl) == 0 {
		// delete all sosmed
		if err := p.PengurusSosmedModel.DeleteByPengurusID(ctx, pengurusID); err != nil {
			return nil, err
		}
	}

	// extract url info
	socialMediaInfo, err := utils.ExtractSocialMediaBatch(sosmedUrl)
	if err != nil {
		return nil, fmt.Errorf("failed to extract social media info: %w", err)
	}

	sosmedPayload := make([]dto.CreatePengurusSosmedPayload, len(sosmedUrl))
	for i, url := range socialMediaInfo {
		sosmedPayload[i] = dto.CreatePengurusSosmedPayload{
			PengurusID: pengurusID,
			Platform:   url.Platform,
			Username:   url.Username,
			Url:        url.URL,
			IsPrimary:  i == 0, // true hanya untuk index 0
		}
	}

	//delete all sosmed
	if err := p.PengurusSosmedModel.DeleteByPengurusID(ctx, pengurusID); err != nil {
		return nil, fmt.Errorf("failed to delete data pengurus sosmed: %w", err)
	}

	// re insert data sosmed
	sosmedResponse, err := p.PengurusSosmedModel.InsertPengurusSosmed(ctx, sosmedPayload)
	if err != nil {
		return nil, fmt.Errorf("failed to insert data pengurus sosmed: %w", err)
	}

	return sosmedResponse, nil
}

func (p *PengurusService) uploadProfilePhoto(ctx context.Context, userID int, userRole string, fileUpload *dto.UploadFileRequest) (string, error) {
	allow, err := pengurusAuthorization.CanEditPengurusField(userRole, "photo_url")
	if !allow || err != nil {
		return "", fmt.Errorf("sorry bro an error happened: %w", err)
	}

	now := time.Now()
	gallery := &dto.GalleryInsert{
		IDUsers:     userID,
		GalleryName: "foto profil pengurus",
		GalleryType: "pengurus",
		Description: "foto identitas diri yang mewakili pengurus doscom",
		EventDate: time.Date(
			now.Year(),
			now.Month(),
			now.Day(), 0, 0, 0, 0, time.UTC,
		),
	}

	_, fileURL, err := p.GalleryService.InsertGalleryAndFileUpload(ctx, gallery, fileUpload)
	if err != nil {
		return "", err
	}

	return fileURL, nil
}

func buildPenguruspayload(
	dataPengurus *dto.RegisterPengurusRequest,
	userRole string,
	divisi string,
	validPosition string,
) (*entity.Pengurus, error) {
	pengurusPayload := dto.PengurusPayload{
		Email:            dataPengurus.Email,
		Divisi:           divisi,
		Name:             dataPengurus.Name,
		Position:         validPosition,
		StartPeriodeYear: dataPengurus.StartPeriodeYear,
		EndPeriodeYear:   dataPengurus.EndPeriodeYear,
		PhotoURL:         dataPengurus.PhotoURL,
	}
	fillableFields, err := authorization.FilterRoleFieldPermission(userRole, &pengurusPayload)
	if err != nil {
		return nil, err
	}

	finalData := &entity.Pengurus{}
	decoder, err := mapstructure.NewDecoder(&mapstructure.DecoderConfig{
		TagName: "json",
		Result:  finalData,
	})
	if err != nil {
		return nil, err
	}
	if err := decoder.Decode(fillableFields); err != nil {
		return nil, fmt.Errorf("failed to decode data %w", err)
	}

	return finalData, nil
}

func (p *PengurusService) CreatePengurus(
	ctx context.Context,
	currentUserID int,
	userRole string,
	dataPengurus *dto.RegisterPengurusRequest,
	fileUpload *dto.UploadFileRequest,
) (*dto.PengurusResponse, error) {

	userValidDivisi, err := authorization.GetRoleInfo(userRole)
	if err != nil {
		return nil, err
	}
	// auto assign position and divisi
	divisi, validPosition, err := authorization.SetDivitionAndPositionByRole(dataPengurus.Position, dataPengurus.Divisi, userRole)
	if err != nil {
		return nil, err
	}
	log.Printf("[service] division: %s and position: %s", divisi, validPosition)

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
	finalData, err := buildPenguruspayload(dataPengurus, userRole, divisi, validPosition)
	if err != nil {
		return nil, err
	}

	// upload photo
	if fileUpload != nil {
		fileURL, err := p.uploadProfilePhoto(ctx, dataPengurus.UserID, userRole, fileUpload)
		if err != nil {
			return nil, fmt.Errorf("terjadi error ketika upload gambar: %w", err)
		}

		finalData.PhotoURL = fileURL
	}

	// insert data pengurus
	if userValidDivisi.Role == constants.RolePengurus {
		finalData.Position = validPosition
	}
	finalData.IDUser = dataPengurus.UserID
	// log.Printf("[service] userID: %d & position: %s", finalData.IDUser, finalData.Position)
	if err := p.PengurusModel.InsertPengurus(finalData); err != nil {
		return nil, err
	}

	// insert data sosmed pengurus
	var socialMediaResponse []dto.PengurusSosmedResponse
	if len(dataPengurus.Sosmed) != 0 {
		socialMediaResponse, err = p.BuildSosmedAndSave(ctx, finalData.ID, dataPengurus.Sosmed)
		if err != nil {
			return nil, err
		}
	}

	return &dto.PengurusResponse{
		ID:               finalData.ID,
		IDUser:           finalData.IDUser,
		PhotoURL:         finalData.PhotoURL,
		Email:            finalData.Email,
		Divisi:           finalData.Divisi,
		Name:             finalData.Name,
		Position:         finalData.Position,
		Sosmed:           socialMediaResponse,
		StartPeriodeYear: finalData.StartPeriodeYear,
		EndPeriodeYear:   finalData.EndPeriodeYear,
	}, nil

}

func (p *PengurusService) UpdateDataPengurus(
	ctx context.Context,
	idParams int,
	currentUserID int,
	userRole string,
	dataPengurus *dto.PengurusPatch,
	fileUpload *dto.UploadFileRequest,
	isSelfUpdate bool,
) (*dto.PengurusPublicResponse, error) {

	var (
		updatedPengurus *dto.PengurusResponse
		error           error
	)

	log.Printf("[service] service ter execute")
	userData, err := p.PengurusModel.GetPengurusByUserID(ctx, idParams)
	if err != nil {
		log.Printf("[service] userID: %d", idParams)
		return nil, err
	}

	dataUser := entity.Pengurus{
		ID:               userData.ID,
		IDUser:           userData.IDUser,
		PhotoURL:         userData.PhotoURL,
		Email:            userData.Email,
		Divisi:           userData.Divisi,
		Name:             userData.Name,
		Position:         userData.Position,
		StartPeriodeYear: userData.StartPeriodeYear,
		EndPeriodeYear:   userData.EndPeriodeYear,
	}
	// authorization check for update data
	roleUser, err := pengurusAuthorization.RolePositionAuthorization(
		ctx,
		idParams,
		currentUserID,
		userRole,
		dataUser,
	)
	if err != nil {
		return nil, err
	}

	dataPengurusPayload := dto.PengurusPayload{
		Email:            dataPengurus.Email,
		Divisi:           dataPengurus.Divisi,
		Name:             dataPengurus.Name,
		StartPeriodeYear: dataPengurus.StartPeriodeYear,
		EndPeriodeYear:   dataPengurus.EndPeriodeYear,
		Position:         dataPengurus.Position,
		PhotoURL:         dataPengurus.PhotoURL,
	}
	// filter fileld update by role and update profile
	editableFields, err := authorization.FilterRoleFieldPermission(userRole, &dataPengurusPayload)
	if err != nil {
		return nil, err
	}

	canUpdatePhoto := roleUser == constants.RoleAdmin || (roleUser == constants.RolePengurus && currentUserID == idParams)
	if fileUpload != nil {
		if !canUpdatePhoto {
			return nil, fmt.Errorf("koordinator tidak dapat memperbarui foto pengurus")
		}
		// Pastikan kita ambil data pengurus dulu untuk tahu UserID-nya yang asli
		targetPengurus, err := p.PengurusModel.GetPengurusByUserID(ctx, idParams)
		if err != nil {
			log.Printf("[service] error disini 3")
			return nil, err
		}

		// Pastikan juga bungkusan fileUpload menggunakan UserID asli dari tabel users
		fileUpload.UserID = uint(targetPengurus.IDUser)

		// update file upload and gallery
		now := time.Now()
		gallery := &dto.GalleryInsert{
			IDUsers:     targetPengurus.IDUser,
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

	updatedPengurus = new(dto.PengurusResponse)
	// update data pengurus
	if isSelfUpdate {
		updatedPengurus, error = p.PengurusModel.UpdatePengurusPartialByUserID(idParams, editableFields)
		if error != nil {
			log.Printf("[service] error disini 3")
			return nil, error
		}
	} else {
		updatedPengurus, error = p.PengurusModel.UpdatePenguruspartial(idParams, editableFields)
		if error != nil {
			log.Printf("[service] error disini 2")
			return nil, error
		}
	}

	// update data sosmed
	var sosmedResponse []dto.PengurusSosmedResponse
	if len(dataPengurus.Sosmed) != 0 {
		sosmedResponse, err = p.UpdatePengurusSosmed(ctx, updatedPengurus.ID, dataPengurus.Sosmed)
		if err != nil {
			log.Printf("[service] error disini 1")
			return nil, err
		}
	}

	return &dto.PengurusPublicResponse{
		ID:               updatedPengurus.ID,
		PhotoURL:         updatedPengurus.PhotoURL,
		Divisi:           updatedPengurus.Divisi,
		Name:             updatedPengurus.Name,
		Position:         updatedPengurus.Position,
		Sosmed:           sosmedResponse,
		StartPeriodeYear: updatedPengurus.StartPeriodeYear,
		EndPeriodeYear:   updatedPengurus.EndPeriodeYear,
	}, nil
}

func (p *PengurusService) GetPengurusBaseOnDivision(
	ctx context.Context,
	userRole, divisi string,
) ([]dto.PengurusResponse, error) {
	// cek role
	role, ok := constants.RoleGroup[userRole]
	if !ok {
		return nil, fmt.Errorf("role not valid")
	}

	var (
		pengurusResponse []dto.PengurusResponse
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
) ([]dto.PengurusResponse, error) {
	if division == "" {
		return nil, fmt.Errorf("division must be provided")
	}

	pengurusResponse, err := p.PengurusModel.GetPengurusByDivisi(ctx, division)
	if err != nil {
		return nil, fmt.Errorf("terjadi error ketika ambil data %w", err)
	}

	return pengurusResponse, nil
}

func (p *PengurusService) GetPengurusByUserID(ctx context.Context, userRole string, userID int) (dto.PengurusResponse, error) {
	validRole, err := authorization.GetRoleInfo(userRole)
	if err != nil {
		log.Printf("[service] error disini, userRole: %s, userDivisi: %s", validRole.Role, validRole.Divisi)
		return dto.PengurusResponse{}, fmt.Errorf("role not valid")
	}

	pengurusResponse, err := p.PengurusModel.GetPengurusByUserID(ctx, userID)
	if err != nil {
		return dto.PengurusResponse{}, fmt.Errorf("error while getting the data: %w", err)
	}

	switch validRole.Role {
	case constants.RoleKoordinator:
		if validRole.Divisi != pengurusResponse.Divisi {
			return dto.PengurusResponse{}, fmt.Errorf("you can not see other division bro %w", err)
		}
	case constants.RolePengurus:
		if userID != pengurusResponse.IDUser {
			return dto.PengurusResponse{}, fmt.Errorf("you can't see other data bro")
		}
	case constants.RoleAdmin:
		// do nothing
	default:
		log.Printf("[service] error disini wak")
		return dto.PengurusResponse{}, fmt.Errorf("role not valid")
	}

	return *pengurusResponse, nil
}

func (p *PengurusService) GetPengurusByID(ctx context.Context, userRole string, id int) (dto.PengurusResponse, error) {
	validRole, err := authorization.GetRoleInfo(userRole)
	if err != nil {
		return dto.PengurusResponse{}, fmt.Errorf("role not valid")
	}

	pengurusResponse, err := p.PengurusModel.GetPengurusByID(ctx, id)
	if err != nil {
		return dto.PengurusResponse{}, fmt.Errorf("error while getting the data: %w", err)
	}

	switch validRole.Role {
	case constants.RoleKoordinator:
		if validRole.Divisi != pengurusResponse.Divisi {
			return dto.PengurusResponse{}, fmt.Errorf("you can not see other division bro %w", err)
		}
	case constants.RolePengurus:
		if id != pengurusResponse.ID {
			return dto.PengurusResponse{}, fmt.Errorf("you can't see other data bro")
		}
	case constants.RoleAdmin:
	default:
		return dto.PengurusResponse{}, fmt.Errorf("role not valid")
	}

	return *pengurusResponse, nil
}

func (p *PengurusService) DeletePengurusById(ctx context.Context, idPengurus int, userRole string) error {
	validRole, ok := constants.RoleGroup[userRole]
	if !ok {
		return fmt.Errorf("role not valid")
	}

	targetPengurus, err := p.PengurusModel.GetPengurusByID(ctx, idPengurus)
	if err != nil {
		return err
	}

	switch validRole.Role {
	case constants.RoleAdmin:
		// lakukan trancation delete pengurus and pengurus sosmed
	case constants.RoleKoordinator:
		if validRole.Divisi != targetPengurus.Divisi {
			return fmt.Errorf("you can not delete data from other division, %w", err)
		}
	case constants.RolePengurus:
		return fmt.Errorf("you can't delete your own data %w", err)
	default:
		return fmt.Errorf("role not valid")
	}

	// begin transaction
	err = p.DB.Transaction(func(tx *gorm.DB) error {

		modelPengurus := p.PengurusModel.WithTx(tx)
		modelPengurusSosmed := p.PengurusSosmedModel.WithTx(tx)

		// delete sosmed first
		if err := modelPengurusSosmed.DeleteByPengurusID(ctx, idPengurus); err != nil {
			return err
		}

		// delete data pengurus
		if err := modelPengurus.DeletePengurus(ctx, idPengurus); err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		return fmt.Errorf("failed to delete data %w", err)
	}

	return nil
}

func (p *PengurusService) BuildSosmedAndSave(ctx context.Context, pengurusID int, sosmedURL []string) ([]dto.PengurusSosmedResponse, error) {
	socialMediaInfo, err := utils.ExtractSocialMediaBatch(sosmedURL)
	if err != nil {
		return nil, fmt.Errorf("failed to extract social media info: %w", err)
	}

	socialMediaInsert := make([]dto.CreatePengurusSosmedPayload, len(socialMediaInfo))
	for i, info := range socialMediaInfo {
		socialMediaInsert[i] = dto.CreatePengurusSosmedPayload{
			PengurusID: pengurusID,
			Platform:   info.Platform,
			Username:   info.Username,
			Url:        info.URL,
			IsPrimary:  i == 0, // true hanya untuk index 0
		}
	}

	// insert pengurus sosmed
	socialMediaResponse, err := p.PengurusSosmedModel.InsertPengurusSosmed(ctx, socialMediaInsert)
	if err != nil {
		return nil, fmt.Errorf("failed to insert social media data: %w", err)
	}

	return socialMediaResponse, nil
}
