package constants

const (
	RoleAdmin       = "admin"
	RoleKoordinator = "koor"
	RolePengurus    = "pengurus"
)

var ValidPosition = map[string]string{
	"ketum":        RoleAdmin,
	"sdm":          RoleKoordinator,
	"pr":           RolePengurus,
	"pm":           RoleKoordinator,
	"pmAng":        RoleKoordinator,
	"KoorPemro":    RoleKoordinator,
	"KoorJaringan": RoleKoordinator,
	"KoorMedcrev":  RoleKoordinator,
	"KoorData":     RoleKoordinator,
	"sekum":        RoleKoordinator,
	"bendum":       RolePengurus,
	"sekAng":       RolePengurus,
	"bendAng":      RolePengurus,
	"PemroAng":     RolePengurus,
	"JaringanAng":  RolePengurus,
	"MedcrevAng":   RolePengurus,
	"DataAng":      RolePengurus,
}

// position group
var PositionGroup = map[string][]string{
	"bph":      {"ketum", "sdm", "pr", "pm", "pmang", "sekum", "sekang", "bendum", "bendang"},
	"pemro":    {"KoorPemro", "PemroAng"},
	"jaringan": {"KoorJaringan", "JaringanAng"},
	"medcrev":  {"KoorMedcrev", "MedcrevAng"},
	"data":     {"KoorData", "DataAng"},
}

type divitioninfo struct {
	Role   string
	Divisi string
}

var RoleGroup = map[string]divitioninfo{
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
	RoleAdmin:       {"name", "email", "divisi", "position", "sosmed", "period", "photo_url"},
	RoleKoordinator: {"name", "email", "divisi", "position", "sosmed", "period"},
	RolePengurus:    {"name", "email", "divisi", "sosmed", "period", "photo_url"},
}

// map auto assign role koor to anggota role
var AutoAsignRole = map[string]string{
	"KoorPemro":    "pemroAnggota",
	"KoorJaringan": "jaringanAnggota",
	"KoorData":     "medcrevAnggota",
	"KoorMedcrev":  "dataAnggota",
	"BPH":          "BPHAnggota",
}
