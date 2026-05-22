package service

import (
	"context"
	"fmt"
	"time"
	"web_doscom/internal/authorization"
	pengurusAuthorization "web_doscom/internal/authorization/pengurus"
	"web_doscom/internal/constants"
	"web_doscom/internal/database/model/dto"
	"web_doscom/internal/database/model/entity"
	"web_doscom/internal/utils"

	"github.com/mitchellh/mapstructure"
)

type PengurusService struct {
	PengurusModel       *entity.PengurusModel
	PengurusSosmedModel *entity.PengurusSosmedModel
	GalleryService      *GalleryService
}

func NewPengurusService(m *entity.PengurusModel, p *entity.PengurusSosmedModel, g *GalleryService) *PengurusService {
	return &PengurusService{
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

func (p *PengurusService) CreatePengurus(
	ctx context.Context,
	currentUserID int,
	userRole string,
	dataPengurus *dto.RegisterPengurusRequest,
	fileUpload *dto.UploadFileRequest,
) (*dto.PengurusResponse, error) {

	// auto assign position and divisi
	divisi, validPosition, err := authorization.SetDivitionAndPositionByRole(dataPengurus.Position, dataPengurus.Divisi, userRole)
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
	data := dto.PengurusPayload{
		Email:            dataPengurus.Email,
		Divisi:           divisi,
		Name:             dataPengurus.Name,
		Position:         validPosition,
		StartPeriodeYear: dataPengurus.StartPeriodeYear,
		EndPeriodeYear:   dataPengurus.EndPeriodeYear,
		PhotoURL:         dataPengurus.PhotoURL,
	}

	fillableFields, err := authorization.FilterRoleFieldPermission(userRole, &data)
	if err != nil {
		return nil, err
	}

	var finalData *entity.Pengurus
	if err := mapstructure.Decode(fillableFields, &finalData); err != nil {
		return nil, fmt.Errorf("failed to decode data")
	}

	// upload photo
	if _, canUploadPhoto := fillableFields["photo_url"]; canUploadPhoto {
		now := time.Now()
		gallery := &dto.GalleryInsert{
			IDUsers:     finalData.IDUser,
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

	// insert data pengurus
	if err := p.PengurusModel.InsertPengurus(finalData); err != nil {
		return nil, err
	}

	// insert data sosmed pengurus
	var socialMediaResponse []dto.PengurusSosmedResponse
	if len(dataPengurus.Sosmed) != 0 {

		socialMediaInfo, err := utils.ExtractSocialMediaBatch(dataPengurus.Sosmed)
		if err != nil {
			return nil, fmt.Errorf("failed to extract social media info %w", err)
		}

		socialMediaInsert := make([]dto.CreatePengurusSosmedPayload, len(socialMediaInfo))
		for i, info := range socialMediaInfo {
			socialMediaInsert[i] = dto.CreatePengurusSosmedPayload{
				PengurusID: finalData.ID,
				Platform:   info.Platform,
				Username:   info.Username,
				Url:        info.URL,
				IsPrimary:  i == 0, // true hanya untuk index 0
			}
		}

		socialMediaResponse, err = p.PengurusSosmedModel.InsertPengurusSosmed(ctx, socialMediaInsert)
		if err != nil {
			return nil, fmt.Errorf("failed to insert social media data %w", err)
		}
	}

	return &dto.PengurusResponse{
		ID:               finalData.ID,
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
) (*dto.PengurusPublicResponse, error) {

	var (
		updatedPengurus *dto.PengurusResponse
		error           error
	)

	userData, err := p.PengurusModel.GetPengurusById(ctx, idParams)
	if err != nil {
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
		targetPengurus, err := p.PengurusModel.GetPengurusById(ctx, idParams)
		if err != nil {
			return nil, err
		}

		// Pastikan juga bungkusan fileUpload menggunakan UserID asli dari tabel users (angka 2)
		fileUpload.UserID = uint(targetPengurus.IDUser)

		// update file upload and gallery
		now := time.Now()
		gallery := &dto.GalleryInsert{
			IDUsers:     targetPengurus.IDUser, // Correctly link to the UserID, not Pengurus ID
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

	// update data sosmed
	var sosmedResponse []dto.PengurusSosmedResponse
	if len(dataPengurus.Sosmed) != 0 {
		sosmedResponse, err = p.UpdatePengurusSosmed(ctx, updatedPengurus.ID, dataPengurus.Sosmed)
		if err != nil {
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

	pengurusResponse, err := p.PengurusModel.GetAllPengurusByDivisi(ctx, division)
	if err != nil {
		return nil, fmt.Errorf("terjadi error ketika ambil data %w", err)
	}

	dataPengurus := make([]dto.PengurusResponse, 0, len(pengurusResponse))
	for _, data := range pengurusResponse {
		dataPengurus = append(dataPengurus, dto.PengurusResponse{
			ID:               data.ID,
			PhotoURL:         data.PhotoURL,
			Email:            data.Email,
			Divisi:           data.Divisi,
			Name:             data.Name,
			Position:         data.Position,
			StartPeriodeYear: data.StartPeriodeYear,
			EndPeriodeYear:   data.EndPeriodeYear,
		})
	}

	return dataPengurus, nil
}

func (p *PengurusService) GetPengurusByID(ctx context.Context, id int, userRole string, userID int) (dto.PengurusResponse, error) {
	validRole, err := authorization.GetRoleInfo(userRole)
	if err != nil {
		return dto.PengurusResponse{}, fmt.Errorf("role not valid")
	}

	pengurusResponse, err := p.PengurusModel.GetPengurusById(ctx, id)
	if err != nil {
		return dto.PengurusResponse{}, fmt.Errorf("error while getting the data %w", err)
	}

	switch validRole.Role {
	case constants.RoleKoordinator:
		if validRole.Divisi != pengurusResponse.Divisi {
			return dto.PengurusResponse{}, fmt.Errorf("you can not see other division bro %w", err)
		}
	case constants.RolePengurus:
		if userID != pengurusResponse.ID {
			return dto.PengurusResponse{}, fmt.Errorf("you can't see other data bro")
		}
	default:
		return dto.PengurusResponse{}, fmt.Errorf("role not valid")
	}

	return *pengurusResponse, nil
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
