package seeder

import (
	"bytes"
	"context"
	"fmt"
	"log"
	"mime/multipart"
	"os"
	"path/filepath"
	"time"
	"web_doscom/internal/config"
	"web_doscom/internal/database"
	"web_doscom/internal/database/model/dto"
	"web_doscom/internal/database/model/entity"
	"web_doscom/internal/service"
	"web_doscom/internal/utils"

	"gorm.io/gorm"
)

func openSeedImage(path string) (
	*multipart.FileHeader,
	multipart.File,
	func(),
	error,
) {
	content, err := os.ReadFile(path)
	if err != nil {
		return nil, nil, nil, err
	}

	var body bytes.Buffer
	writer := multipart.NewWriter(&body)

	part, err := writer.CreateFormFile("file", filepath.Base(path))
	if err != nil {
		return nil, nil, nil, err
	}

	if _, err := part.Write(content); err != nil {
		return nil, nil, nil, err
	}

	if err := writer.Close(); err != nil {
		return nil, nil, nil, err
	}

	reader := multipart.NewReader(&body, writer.Boundary())

	form, err := reader.ReadForm(10 << 20)
	if err != nil {
		return nil, nil, nil, err
	}

	headers := form.File["file"]
	if len(headers) == 0 {
		form.RemoveAll()
		return nil, nil, nil, fmt.Errorf("file header not found")
	}

	fileHeader := headers[0]
	file, err := fileHeader.Open()
	if err != nil {
		form.RemoveAll()
		return nil, nil, nil, err
	}

	cleanup := func() {
		_ = file.Close()
		_ = form.RemoveAll()
	}

	return fileHeader, file, cleanup, nil

}

func SeedPengurus(db *gorm.DB, galleryService *service.GalleryService, sosmedURL []string) error {

	ctx := context.Background()
	pengurusModel := entity.PengurusModel{DB: db}

	now := time.Now()

	pengurusList := []entity.Pengurus{
		{
			IDUser:           1,
			PhotoURL:         "",
			Email:            "111202315460@mhs.dinus.ac.id",
			Divisi:           "bph",
			Name:             "Aldio Sebastiansyah",
			Position:         "ketuaUmum",
			StartPeriodeYear: 2024,
			EndPeriodeYear:   2026,
			CreatedAt:        now,
			UpdatedAt:        now,
		},
		{
			IDUser:           2,
			PhotoURL:         "",
			Email:            "111202314916@mhs.dinus.ac.id",
			Divisi:           "bph",
			Name:             "Muhammad Imam Rafi'",
			Position:         "kepalaBidangHubunganMasyarakat",
			StartPeriodeYear: 2024,
			EndPeriodeYear:   2026,
			CreatedAt:        now,
			UpdatedAt:        now,
		},
		{
			IDUser:           3,
			PhotoURL:         "",
			Email:            "111202314933@mhs.dinus.ac.id",
			Divisi:           "bph",
			Name:             "Muhammad Ivan Rafsanjani",
			Position:         "kepalaBidangSumberDayaUmum",
			StartPeriodeYear: 2024,
			EndPeriodeYear:   2026,
			CreatedAt:        now,
			UpdatedAt:        now,
		},
		{
			IDUser:           4,
			PhotoURL:         "",
			Email:            "111202315298@mhs.dinus.ac.id",
			Divisi:           "bph",
			Name:             "Desvita Maharani",
			Position:         "sekretarisUmumI",
			StartPeriodeYear: 2024,
			EndPeriodeYear:   2026,
			CreatedAt:        now,
			UpdatedAt:        now,
		},
		{
			IDUser:           5,
			PhotoURL:         "",
			Email:            "112202407184@mhs.dinus.ac.id",
			Divisi:           "bph",
			Name:             "Inas Salwa Nuraini",
			Position:         "sekretarisUmumII",
			StartPeriodeYear: 2025,
			EndPeriodeYear:   2028,
			CreatedAt:        now,
			UpdatedAt:        now,
		},
		{
			IDUser:           6,
			PhotoURL:         "",
			Email:            "111202315443@mhs.dinus.ac.id",
			Divisi:           "bph",
			Name:             "Maxentia Kathleen",
			Position:         "bendaharaUmumI",
			StartPeriodeYear: 2024,
			EndPeriodeYear:   2026,
			CreatedAt:        now,
			UpdatedAt:        now,
		},
		{
			IDUser:           7,
			PhotoURL:         "",
			Email:            "111202415648@mhs.dinus.ac.id",
			Divisi:           "bph",
			Name:             "Widya Allifya Az Zahra",
			Position:         "bendaharaUmumII",
			StartPeriodeYear: 2025,
			EndPeriodeYear:   2028,
			CreatedAt:        now,
			UpdatedAt:        now,
		},
		{
			IDUser:           8,
			PhotoURL:         "",
			Email:            "111202315332@mhs.dinus.ac.id",
			Divisi:           "bph",
			Name:             "Harry LBI",
			Position:         "projectManagerI",
			StartPeriodeYear: 2024,
			EndPeriodeYear:   2026,
			CreatedAt:        now,
			UpdatedAt:        now,
		},
		{
			IDUser:           9,
			PhotoURL:         "",
			Email:            "111202415691@mhs.dinus.ac.id",
			Divisi:           "bph",
			Name:             "Fadhil Riyanto",
			Position:         "projectManagerII",
			StartPeriodeYear: 2024,
			EndPeriodeYear:   2028,
			CreatedAt:        now,
			UpdatedAt:        now,
		},
		{
			IDUser:           10,
			PhotoURL:         "",
			Email:            "111202314978@mhs.dinus.ac.id",
			Divisi:           "bph",
			Name:             "Wildanu Rafif Albaihaqi",
			Position:         "koordinatorHubunganMasyarakatExternal",
			StartPeriodeYear: 2024,
			EndPeriodeYear:   2026,
			CreatedAt:        now,
			UpdatedAt:        now,
		},
		{
			IDUser:           11,
			PhotoURL:         "",
			Email:            "111202315253@mhs.dinus.ac.id",
			Divisi:           "bph",
			Name:             "Sulthan Yustr Suwardhi",
			Position:         "hubunganMasyarakatExternal",
			StartPeriodeYear: 2024,
			EndPeriodeYear:   2026,
			CreatedAt:        now,
			UpdatedAt:        now,
		},
		{
			IDUser:           12,
			PhotoURL:         "",
			Email:            "111202415841@mhs.dinus.ac.id",
			Divisi:           "bph",
			Name:             "Wilma Auraruna Khalif",
			Position:         "hubunganMasyarakatExternal",
			StartPeriodeYear: 2025,
			EndPeriodeYear:   2028,
			CreatedAt:        now,
			UpdatedAt:        now,
		},
		{
			IDUser:           13,
			PhotoURL:         "",
			Email:            "111202415566@mhs.dinus.ac.id",
			Divisi:           "bph",
			Name:             "Danish Putra Utama",
			Position:         "hubunganMasyarakatExternal",
			StartPeriodeYear: 2025,
			EndPeriodeYear:   2028,
			CreatedAt:        now,
			UpdatedAt:        now,
		},
		{
			IDUser:           14,
			PhotoURL:         "",
			Email:            "111202315448@mhs.dinus.ac.id",
			Divisi:           "medcrev",
			Name:             "Dhion Nur Damanhuri",
			Position:         "koordinatorMediaCreative",
			StartPeriodeYear: 2024,
			EndPeriodeYear:   2026,
			CreatedAt:        now,
			UpdatedAt:        now,
		},
		{
			IDUser:           15,
			PhotoURL:         "",
			Email:            "111202315034@mhs.dinus.ac.id",
			Divisi:           "medcrev",
			Name:             "Mohammad Fikri Haikal",
			Position:         "mediaCreativeAnggota",
			StartPeriodeYear: 2024,
			EndPeriodeYear:   2026,
			CreatedAt:        now,
			UpdatedAt:        now,
		},
		{
			IDUser:           16,
			PhotoURL:         "",
			Email:            "111202315039@mhs.dinus.ac.id",
			Divisi:           "medcrev",
			Name:             "Megan Febriana Putri Johana",
			Position:         "mediaCreativeAnggota",
			StartPeriodeYear: 2024,
			EndPeriodeYear:   2026,
			CreatedAt:        now,
			UpdatedAt:        now,
		},
		{
			IDUser:           17,
			PhotoURL:         "",
			Email:            "114202304209@mhs.dinus.ac.id",
			Divisi:           "medcrev",
			Name:             "Garneta Ary Yuwaningtyas",
			Position:         "mediaCreativeAnggota",
			StartPeriodeYear: 2024,
			EndPeriodeYear:   2026,
			CreatedAt:        now,
			UpdatedAt:        now,
		},
		{
			IDUser:           18,
			PhotoURL:         "",
			Email:            "111202315017@mhs.dinus.ac.id",
			Divisi:           "medcrev",
			Name:             "Reno Dwi Aderelyan",
			Position:         "mediaCreativeAnggota",
			StartPeriodeYear: 2024,
			EndPeriodeYear:   2026,
			CreatedAt:        now,
			UpdatedAt:        now,
		},
		{
			IDUser:           19,
			PhotoURL:         "",
			Email:            "112202407174@mhs.dinus.ac.id",
			Divisi:           "medcrev",
			Name:             "Nayzala Aura Riyana",
			Position:         "mediaCreativeAnggota",
			StartPeriodeYear: 2025,
			EndPeriodeYear:   2028,
			CreatedAt:        now,
			UpdatedAt:        now,
		},
		{
			IDUser:           20,
			PhotoURL:         "",
			Email:            "111202415895@mhs.dinus.ac.id",
			Divisi:           "medcrev",
			Name:             "Elsandro Rivalito",
			Position:         "mediaCreativeAnggota",
			StartPeriodeYear: 2025,
			EndPeriodeYear:   2028,
			CreatedAt:        now,
			UpdatedAt:        now,
		},
		{
			IDUser:           21,
			PhotoURL:         "",
			Email:            "114202404506@mhs.dinus.ac.id",
			Divisi:           "medcrev",
			Name:             "Fatah Al Farid",
			Position:         "mediaCreativeAnggota",
			StartPeriodeYear: 2025,
			EndPeriodeYear:   2028,
			CreatedAt:        now,
			UpdatedAt:        now,
		},
		{
			IDUser:           22,
			PhotoURL:         "",
			Email:            "111202415585@mhs.dinus.ac.id",
			Divisi:           "medcrev",
			Name:             "Zovan Rizza Fannevi",
			Position:         "mediaCreativeAnggota",
			StartPeriodeYear: 2025,
			EndPeriodeYear:   2028,
			CreatedAt:        now,
			UpdatedAt:        now,
		},
		{
			IDUser:           23,
			PhotoURL:         "",
			Email:            "111202415989@mhs.dinus.ac.id",
			Divisi:           "medcrev",
			Name:             "Jovandya Ardhika",
			Position:         "mediaCreativeAnggota",
			StartPeriodeYear: 2025,
			EndPeriodeYear:   2028,
			CreatedAt:        now,
			UpdatedAt:        now,
		},
		{
			IDUser:           24,
			PhotoURL:         "",
			Email:            "111202415661@mhs.dinus.ac.id",
			Divisi:           "medcrev",
			Name:             "Amelia Ramadhani",
			Position:         "mediaCreativeAnggota",
			StartPeriodeYear: 2025,
			EndPeriodeYear:   2028,
			CreatedAt:        now,
			UpdatedAt:        now,
		},
		{
			IDUser:           25,
			PhotoURL:         "",
			Email:            "111202315018@mhs.dinus.ac.id",
			Divisi:           "medcrev",
			Name:             "Suryani Ayu Dewanti",
			Position:         "mediaCreativeAnggota",
			StartPeriodeYear: 2024,
			EndPeriodeYear:   2026,
			CreatedAt:        now,
			UpdatedAt:        now,
		},
		{
			IDUser:           26,
			PhotoURL:         "",
			Email:            "111202315244@mhs.dinus.ac.id",
			Divisi:           "medcrev",
			Name:             "Jonathan Naufal Farrel",
			Position:         "mediaCreativeAnggota",
			StartPeriodeYear: 2024,
			EndPeriodeYear:   2026,
			CreatedAt:        now,
			UpdatedAt:        now,
		},
		{
			IDUser:           27,
			PhotoURL:         "",
			Email:            "111202315203@mhs.dinus.ac.id",
			Divisi:           "medcrev",
			Name:             "Dian Hana Kartiko Sari",
			Position:         "mediaCreativeAnggota",
			StartPeriodeYear: 2024,
			EndPeriodeYear:   2026,
			CreatedAt:        now,
			UpdatedAt:        now,
		},
		{
			IDUser:           28,
			PhotoURL:         "",
			Email:            "111202415836@mhs.dinus.ac.id",
			Divisi:           "medcrev",
			Name:             "Adhifa fairuz zulfi",
			Position:         "mediaCreativeAnggota",
			StartPeriodeYear: 2025,
			EndPeriodeYear:   2028,
			CreatedAt:        now,
			UpdatedAt:        now,
		},
		{
			IDUser:           29,
			PhotoURL:         "",
			Email:            "111202315393@mhs.dinus.ac.id",
			Divisi:           "pemro",
			Name:             "Rico Andre Pratama",
			Position:         "koordinatorPemrograman",
			StartPeriodeYear: 2024,
			EndPeriodeYear:   2026,
			CreatedAt:        now,
			UpdatedAt:        now,
		},
		{
			IDUser:           30,
			PhotoURL:         "",
			Email:            "111202315458@mhs.dinus.ac.id",
			Divisi:           "pemro",
			Name:             "Felixs Togar Nugroho Siahaan",
			Position:         "pemrogramanAnggota",
			StartPeriodeYear: 2024,
			EndPeriodeYear:   2026,
			CreatedAt:        now,
			UpdatedAt:        now,
		},
		{
			IDUser:           31,
			PhotoURL:         "",
			Email:            "111202415776@mhs.dinus.ac.id",
			Divisi:           "pemro",
			Name:             "Husnul Fikri Averus",
			Position:         "pemrogramanAnggota",
			StartPeriodeYear: 2025,
			EndPeriodeYear:   2028,
			CreatedAt:        now,
			UpdatedAt:        now,
		},
		{
			IDUser:           32,
			PhotoURL:         "",
			Email:            "111202415651@mhs.dinus.ac.id",
			Divisi:           "pemro",
			Name:             "Naf'an Nur'alim",
			Position:         "pemrogramanAnggota",
			StartPeriodeYear: 2025,
			EndPeriodeYear:   2028,
			CreatedAt:        now,
			UpdatedAt:        now,
		},
		{
			IDUser:           33,
			PhotoURL:         "",
			Email:            "111202415546@mhs.dinus.ac.id",
			Divisi:           "pemro",
			Name:             "Rizki Ardiansyah Novianto",
			Position:         "pemrogramanAnggota",
			StartPeriodeYear: 2025,
			EndPeriodeYear:   2028,
			CreatedAt:        now,
			UpdatedAt:        now,
		},
		{
			IDUser:           34,
			PhotoURL:         "",
			Email:            "111202416082@mhs.dinus.ac.id",
			Divisi:           "pemro",
			Name:             "Muhammad Bagus Aditya",
			Position:         "pemrogramanAnggota",
			StartPeriodeYear: 2025,
			EndPeriodeYear:   2028,
			CreatedAt:        now,
			UpdatedAt:        now,
		},
		{
			IDUser:           35,
			PhotoURL:         "",
			Email:            "111202415833@mhs.dinus.ac.id",
			Divisi:           "pemro",
			Name:             "Jessica Leilani Handoko",
			Position:         "pemrogramanAnggota",
			StartPeriodeYear: 2025,
			EndPeriodeYear:   2028,
			CreatedAt:        now,
			UpdatedAt:        now,
		},
		{
			IDUser:           36,
			PhotoURL:         "",
			Email:            "111202415656@mhs.dinus.ac.id",
			Divisi:           "pemro",
			Name:             "Muhammad Ilham Maulana",
			Position:         "pemrogramanAnggota",
			StartPeriodeYear: 2025,
			EndPeriodeYear:   2028,
			CreatedAt:        now,
			UpdatedAt:        now,
		},
		{
			IDUser:           37,
			PhotoURL:         "",
			Email:            "111202415775@mhs.dinus.ac.id",
			Divisi:           "pemro",
			Name:             "Ghifari Wira Andanito",
			Position:         "pemrogramanAnggota",
			StartPeriodeYear: 2025,
			EndPeriodeYear:   2028,
			CreatedAt:        now,
			UpdatedAt:        now,
		},
		{
			IDUser:           38,
			PhotoURL:         "",
			Email:            "111202415726@mhs.dinus.ac.id",
			Divisi:           "pemro",
			Name:             "Fa'iz Maulana Habibi",
			Position:         "pemrogramanAnggota",
			StartPeriodeYear: 2025,
			EndPeriodeYear:   2028,
			CreatedAt:        now,
			UpdatedAt:        now,
		},
		{
			IDUser:           39,
			PhotoURL:         "",
			Email:            "111202415685@mhs.dinus.ac.id",
			Divisi:           "pemro",
			Name:             "Fajar Aziz Kurniawan",
			Position:         "pemrogramanAnggota",
			StartPeriodeYear: 2025,
			EndPeriodeYear:   2028,
			CreatedAt:        now,
			UpdatedAt:        now,
		},
		{
			IDUser:           40,
			PhotoURL:         "",
			Email:            "111202315467@mhs.dinus.ac.id",
			Divisi:           "data",
			Name:             "Fariz Hasim Arvianto",
			Position:         "koordinatorData",
			StartPeriodeYear: 2024,
			EndPeriodeYear:   2026,
			CreatedAt:        now,
			UpdatedAt:        now,
		},
		{
			IDUser:           41,
			PhotoURL:         "",
			Email:            "111202315020@mhs.dinus.ac.id",
			Divisi:           "data",
			Name:             "Brenendra Putra Oktaviansyah",
			Position:         "dataAnggota",
			StartPeriodeYear: 2024,
			EndPeriodeYear:   2026,
			CreatedAt:        now,
			UpdatedAt:        now,
		},
		{
			IDUser:           42,
			PhotoURL:         "",
			Email:            "111202315024@mhs.dinus.ac.id",
			Divisi:           "data",
			Name:             "Johana Oktavia Ramadhani",
			Position:         "dataAnggota",
			StartPeriodeYear: 2024,
			EndPeriodeYear:   2026,
			CreatedAt:        now,
			UpdatedAt:        now,
		},
		{
			IDUser:           43,
			PhotoURL:         "",
			Email:            "111202315031@mhs.dinus.ac.id",
			Divisi:           "data",
			Name:             "Chalida Abdat",
			Position:         "dataAnggota",
			StartPeriodeYear: 2024,
			EndPeriodeYear:   2026,
			CreatedAt:        now,
			UpdatedAt:        now,
		},
		{
			IDUser:           44,
			PhotoURL:         "",
			Email:            "111202315212@mhs.dinus.ac.id",
			Divisi:           "data",
			Name:             "Naya Nasywa Puspita Haryanto",
			Position:         "dataAnggota",
			StartPeriodeYear: 2024,
			EndPeriodeYear:   2026,
			CreatedAt:        now,
			UpdatedAt:        now,
		},
		{
			IDUser:           45,
			PhotoURL:         "",
			Email:            "111202315247@mhs.dinus.ac.id",
			Divisi:           "data",
			Name:             "Berlian Kusumayuda",
			Position:         "dataAnggota",
			StartPeriodeYear: 2025,
			EndPeriodeYear:   2028,
			CreatedAt:        now,
			UpdatedAt:        now,
		},
		{
			IDUser:           46,
			PhotoURL:         "",
			Email:            "111202315367@mhs.dinus.ac.id",
			Divisi:           "jaringan",
			Name:             "Muhammad Daniyal Haq",
			Position:         "koordinatorJaringan",
			StartPeriodeYear: 2024,
			EndPeriodeYear:   2026,
			CreatedAt:        now,
			UpdatedAt:        now,
		},
		{
			IDUser:           47,
			PhotoURL:         "",
			Email:            "111202314944@mhs.dinus.ac.id",
			Divisi:           "jaringan",
			Name:             "Yuanda Kusuma Aji",
			Position:         "jaringanAnggota",
			StartPeriodeYear: 2024,
			EndPeriodeYear:   2026,
			CreatedAt:        now,
			UpdatedAt:        now,
		},
		{
			IDUser:           48,
			PhotoURL:         "",
			Email:            "111202315462@mhs.dinus.ac.id",
			Divisi:           "jaringan",
			Name:             "Muhammad Ni'am Mawahib",
			Position:         "jaringanAnggota",
			StartPeriodeYear: 2024,
			EndPeriodeYear:   2026,
			CreatedAt:        now,
			UpdatedAt:        now,
		},
		{
			IDUser:           49,
			PhotoURL:         "",
			Email:            "112202307051@mhs.dinus.ac.id",
			Divisi:           "jaringan",
			Name:             "Arbinand Roffi Ilmi",
			Position:         "jaringanAnggota",
			StartPeriodeYear: 2024,
			EndPeriodeYear:   2026,
			CreatedAt:        now,
			UpdatedAt:        now,
		},
		{
			IDUser:           50,
			PhotoURL:         "",
			Email:            "111202315400@mhs.dinus.ac.id",
			Divisi:           "jaringan",
			Name:             "Faiz Aditya Dhananjaya",
			Position:         "jaringanAnggota",
			StartPeriodeYear: 2024,
			EndPeriodeYear:   2026,
			CreatedAt:        now,
			UpdatedAt:        now,
		},
		{
			IDUser:           51,
			PhotoURL:         "",
			Email:            "111202315457@mhs.dinus.ac.id",
			Divisi:           "jaringan",
			Name:             "Arya Pradipta",
			Position:         "jaringanAnggota",
			StartPeriodeYear: 2024,
			EndPeriodeYear:   2026,
			CreatedAt:        now,
			UpdatedAt:        now,
		},
		{
			IDUser:           52,
			PhotoURL:         "",
			Email:            "111202416045@mhs.dinus.ac.id",
			Divisi:           "jaringan",
			Name:             "Cherishta Joane Nungki",
			Position:         "jaringanAnggota",
			StartPeriodeYear: 2025,
			EndPeriodeYear:   2028,
			CreatedAt:        now,
			UpdatedAt:        now,
		},
		{
			IDUser:           53,
			PhotoURL:         "",
			Email:            "111202415662@mhs.dinus.ac.id",
			Divisi:           "jaringan",
			Name:             "Fujikawa Shinichi",
			Position:         "jaringanAnggota",
			StartPeriodeYear: 2025,
			EndPeriodeYear:   2028,
			CreatedAt:        now,
			UpdatedAt:        now,
		},
		{
			IDUser:           54,
			PhotoURL:         "",
			Email:            "111202415935@mhs.dinus.ac.id",
			Divisi:           "jaringan",
			Name:             "Aldi Firmansyah",
			Position:         "jaringanAnggota",
			StartPeriodeYear: 2025,
			EndPeriodeYear:   2028,
			CreatedAt:        now,
			UpdatedAt:        now,
		},
		{
			IDUser:           55,
			PhotoURL:         "",
			Email:            "111202415690@mhs.dinus.ac.id",
			Divisi:           "jaringan",
			Name:             "Nicko Galang Anarki",
			Position:         "jaringanAnggota",
			StartPeriodeYear: 2025,
			EndPeriodeYear:   2028,
			CreatedAt:        now,
			UpdatedAt:        now,
		},
		{
			IDUser:           56,
			PhotoURL:         "",
			Email:            "111202415805@mhs.dinus.ac.id",
			Divisi:           "jaringan",
			Name:             "Haikal Ega Pratama",
			Position:         "jaringanAnggota",
			StartPeriodeYear: 2025,
			EndPeriodeYear:   2028,
			CreatedAt:        now,
			UpdatedAt:        now,
		},
		{
			IDUser:           57,
			PhotoURL:         "",
			Email:            "111202415561@mhs.dinus.ac.id",
			Divisi:           "jaringan",
			Name:             "Fajar Maulana P",
			Position:         "jaringanAnggota",
			StartPeriodeYear: 2025,
			EndPeriodeYear:   2028,
			CreatedAt:        now,
			UpdatedAt:        now,
		},
		{
			IDUser:           58,
			PhotoURL:         "",
			Email:            "111202415571@mhs.dinus.ac.id",
			Divisi:           "jaringan",
			Name:             "Muhammad Naufal Tsaqif",
			Position:         "jaringanAnggota",
			StartPeriodeYear: 2025,
			EndPeriodeYear:   2028,
			CreatedAt:        now,
			UpdatedAt:        now,
		},
		{
			IDUser:           59,
			PhotoURL:         "",
			Email:            "111202415718@mhs.dinus.ac.id",
			Divisi:           "jaringan",
			Name:             "Raditya Perwira Putra",
			Position:         "jaringanAnggota",
			StartPeriodeYear: 2025,
			EndPeriodeYear:   2028,
			CreatedAt:        now,
			UpdatedAt:        now,
		},
		{
			IDUser:           60,
			PhotoURL:         "",
			Email:            "111202315025@mhs.dinus.ac.id",
			Divisi:           "jaringan",
			Name:             "Danendra Farrel Haryo Wibowo",
			Position:         "jaringanAnggota",
			StartPeriodeYear: 2024,
			EndPeriodeYear:   2026,
			CreatedAt:        now,
			UpdatedAt:        now,
		},
		{
			IDUser:           61,
			PhotoURL:         "",
			Email:            "111202315389@mhs.dinus.ac.id",
			Divisi:           "jaringan",
			Name:             "Ferdynand Ergy Pramudani",
			Position:         "jaringanAnggota",
			StartPeriodeYear: 2024,
			EndPeriodeYear:   2026,
			CreatedAt:        now,
			UpdatedAt:        now,
		},
	}

	for i := range pengurusList {
		pengurus := &pengurusList[i]
		photoPath := filepath.Join("storage", "uploads", "foto", fmt.Sprintf("%d.jpg", pengurus.IDUser))
		fileHeader, file, cleanup, err := openSeedImage(photoPath)
		if err != nil {
			return fmt.Errorf("failed to open photo user_id=%d: %w", pengurus.IDUser, err)
		}

		galleryData := &dto.GalleryInsert{
			IDUsers:     pengurus.IDUser,
			GalleryName: "foto profil pengurus " + pengurus.Name,
			GalleryType: "pengurus",
			Description: "foto identitas diri yang mewakili pengurus doscom",
			EventDate:   time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, time.UTC),
		}

		fileUpload := &dto.UploadFileRequest{
			FileHeader: fileHeader,
			File:       file,
			Folder:     "pengurus",
			UserID:     uint(pengurus.IDUser),
		}

		_, PhotoURL, err := galleryService.InsertGalleryAndFileUpload(ctx, galleryData, fileUpload)
		cleanup()
		if err != nil {
			return fmt.Errorf("failed to insert gallery user_id=%d: %w", pengurus.IDUser, err)
		}

		pengurus.PhotoURL = PhotoURL
		if err := pengurusModel.InsertPengurus(pengurus); err != nil {
			log.Printf("Failed to insert data pengurus %v", err)
			return fmt.Errorf("failed to insert data pengurus user_id=%d: %w", pengurus.IDUser, err)
		}

		SeedPengurusSosmed(db, pengurus.ID, sosmedURL)
	}

	return nil
}

func SeedPengurusSosmed(db *gorm.DB, pengurusID int, sosmedUrl []string) {
	ctx := context.Background()
	pengurusSosmedModel := entity.PengurusSosmedModel{DB: db}
	// parse url info
	socialMediaInfo, err := utils.ExtractSocialMediaBatch(sosmedUrl)
	if err != nil {
		log.Printf("Failed to extract social media info %v", err)
		return
	}

	sosmedPayload := make([]dto.CreatePengurusSosmedPayload, len(socialMediaInfo))
	for i, info := range socialMediaInfo {
		sosmedPayload[i] = dto.CreatePengurusSosmedPayload{
			PengurusID: pengurusID,
			Platform:   info.Platform,
			Username:   info.Username,
			Url:        info.URL,
			IsPrimary:  i == 0, // true hanya untuk index 0
		}
	}

	_, err = pengurusSosmedModel.InsertPengurusSosmed(ctx, sosmedPayload)
	if err != nil {
		log.Printf("Failed to insert data pengurus sosmed %v", err)
		return
	}

	log.Println("Seed table 'pengurus' completed.")
}

func RunSeedPengurus(db *gorm.DB, minioClient *config.MinioClient) {
	models := database.NewModel(db)
	storageService := service.NewStorageService(
		minioClient,
		&models.FileUploads,
	)
	galleryService := service.NewGalleryService(
		&models.Gallery,
		storageService,
	)

	sosmedURL := []string{
		"https://www.instagram.com/dhioonnn/",
		"https://www.linkedin.com/in/dhion-nur-damanhuri-2bb863275/",
		"https://github.com/IKOPOO",
	}

	if err := SeedPengurus(db, galleryService, sosmedURL); err != nil {
		log.Fatalf("Failed to seed pengurus: %v", err)
	}
}
