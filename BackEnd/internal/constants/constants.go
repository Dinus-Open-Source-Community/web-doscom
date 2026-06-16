package constants

import "regexp"

// ================= ROLE =================
const (
	RoleAdmin       = "admin"
	RoleKoordinator = "koor"
	RolePengurus    = "pengurus"
)

// role key
const (
	RoleKeyPemroAnggota    = "pemroAnggota"
	RoleKeyJaringanAnggota = "jaringanAnggota"
	RoleKeyMedcrevAnggota  = "medcrevAnggota"
	RoleKeyDataAnggota     = "dataAnggota"
	RoleKeyBphAnggota      = "BPHAnggota"
	RoleKeyKoorPemro       = "KoorPemro"
	RoleKeyKoorJaringan    = "KoorJaringan"
	RoleKeyKoorData        = "KoorData"
	RoleKeyKoorMedcrev     = "KoorMedcrev"
	RoleKeyBPH             = "BPH"
	RoleKeySuperAdmin      = "SuperAdmin"
)

var ValidPosition = map[string]string{
	"ketum":        RoleAdmin,
	"sdm":          RoleKoordinator,
	"pm":           RoleKoordinator,
	"pmAng":        RoleKoordinator,
	"KoorPemro":    RoleKoordinator,
	"KoorJaringan": RoleKoordinator,
	"KoorMedcrev":  RoleKoordinator,
	"KoorData":     RoleKoordinator,
	"sekum":        RoleKoordinator,
	"sekAng":       RoleKoordinator,
	"pr":           RoleKoordinator,
	"bendum":       RolePengurus,
	"humas":        RolePengurus,
	"bendAng":      RolePengurus,
	"PemroAng":     RolePengurus,
	"JaringanAng":  RolePengurus,
	"MedcrevAng":   RolePengurus,
	"DataAng":      RolePengurus,
}

var PositionGroup = map[string][]string{
	"bph":      {"ketum", "sdm", "pr", "pm", "pmang", "sekum", "sekang", "bendum", "bendang"},
	"pemro":    {"KoorPemro", "PemroAng"},
	"jaringan": {"KoorJaringan", "JaringanAng"},
	"medcrev":  {"KoorMedcrev", "MedcrevAng"},
	"data":     {"KoorData", "DataAng"},
}

type Divisioninfo struct {
	Role   string
	Divisi string
}

var RoleGroup = map[string]Divisioninfo{
	"pemroAnggota":    {RolePengurus, "pemro"},
	"jaringanAnggota": {RolePengurus, "jaringan"},
	"medcrevAnggota":  {RolePengurus, "medcrev"},
	"dataAnggota":     {RolePengurus, "data"},
	"BPHAnggota":      {RolePengurus, "bph"},
	"KoorPemro":       {RoleKoordinator, "pemro"},
	"KoorJaringan":    {RoleKoordinator, "jaringan"},
	"KoorData":        {RoleKoordinator, "data"},
	"KoorMedcrev":     {RoleKoordinator, "medcrev"},
	"BPH":             {RoleKoordinator, "bph"},
	"SuperAdmin":      {RoleAdmin, "bph"},
}

var RoleLevel = map[string]int{
	RoleAdmin:       1,
	RoleKoordinator: 2,
	RolePengurus:    3,
}

var RoleFieldPermission = map[string][]string{
	RoleAdmin:       {"name", "email", "divisi", "position", "start_periode_year", "end_periode_year", "photo_url"},
	RoleKoordinator: {"name", "email", "divisi", "position", "start_periode_year", "end_periode_year", "photo_url"},
	RolePengurus:    {"name", "email", "divisi", "start_periode_year", "end_periode_year", "photo_url"},
}

// map auto assign role koor to anggota role
var AutoAsignRole = map[string]string{
	"KoorPemro":    "pemroAnggota",
	"KoorJaringan": "jaringanAnggota",
	"KoorData":     "medcrevAnggota",
	"KoorMedcrev":  "dataAnggota",
	"BPH":          "BPHAnggota",
}

// ================= BLOG STATUS =================

const (
	StatusDraft     = "draft"
	StatusPublished = "published"
	StatusUnpublish = "unpublished"
	StatusRejected  = "rejected"
	StatusPending   = "pending_review"
)

// regex password
var (
	AtLeastOneUppercase   = regexp.MustCompile(`[A-Z]`)
	AtLeastOneLowercase   = regexp.MustCompile(`[a-z]`)
	AtLeastOneNumeric     = regexp.MustCompile(`[0-9]`)
	AtLeastOneSpecialChar = regexp.MustCompile(`[^A-Za-z0-9]`)
	EightCharsOrMore      = regexp.MustCompile(`.{8,}`)
)
