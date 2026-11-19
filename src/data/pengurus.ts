// dummy: replace before launch
// Replace `dummyPeriods`/`dummyRoles` and the `pengurus` derivation below with real data.

export type PengurusMember = {
  name: string;
  role: string;
  period: string;
  division: string;
  description: string;
  socialMedia?: { platform: string; url: string }[];
  history?: { role: string; period: string }[];
  /** Omit to render the initials fallback. */
  image?: string;
};

const dummyPeriods = [
  "2023/2024",
  "2024/2025",
  "2025/2026",
  "2026/2027",
  "2027/2028",
  "2028/2029",
  "2029/2030",
];

const dummyRoles = [
  "Ketua Umum",
  "Sekertaris Umum 1",
  "Sekertaris Umum 2",
  "Bendahara Umum 1",
  "Bendahara Umum 2",
  "Project Manager 1",
  "Project Manager 2",
  "Tech Lead 1",
  "Tech Lead 2",
  "Kabid SDM",
  "Koordinator Humas",
  "Koordinator Pemrograman",
  "Koordinator Jaringan",
  "Koordinator Creative Media",
  "Koordinator Data",
  // 16 - 30: Pengurus divisi di bawah koordinator
  "Pengurus Humas",
  "Pengurus Pemrograman",
  "Pengurus Jaringan",
  "Pengurus Creative Media",
  "Pengurus Data",
  "Pengurus Humas",
  "Pengurus Pemrograman",
  "Pengurus Jaringan",
  "Pengurus Creative Media",
  "Pengurus Data",
  "Pengurus Humas",
  "Pengurus Pemrograman",
  "Pengurus Jaringan",
  "Pengurus Creative Media",
  "Pengurus Data",
];

const roleDetails: Record<string, { division: string; description: string }> = {
  "Ketua Umum": {
    division: "Pengurus Inti",
    description:
      "Mengkoordinasikan jalannya organisasi dan menyelaraskan kerja pengurus.",
  },
  "Sekertaris Umum 1": {
    division: "Pengurus Inti",
    description:
      "Mendukung administrasi, dokumentasi, dan komunikasi internal organisasi.",
  },
  "Sekertaris Umum 2": {
    division: "Pengurus Inti",
    description:
      "Mendukung administrasi, dokumentasi, dan komunikasi internal organisasi.",
  },
  "Bendahara Umum 1": {
    division: "Pengurus Inti",
    description: "Mendukung pencatatan serta pengelolaan keuangan organisasi.",
  },
  "Bendahara Umum 2": {
    division: "Pengurus Inti",
    description: "Mendukung pencatatan serta pengelolaan keuangan organisasi.",
  },
  "Project Manager 1": {
    division: "Project",
    description:
      "Mengatur perencanaan, koordinasi, dan perkembangan proyek komunitas.",
  },
  "Project Manager 2": {
    division: "Project",
    description:
      "Mengatur perencanaan, koordinasi, dan perkembangan proyek komunitas.",
  },
  "Tech Lead 1": {
    division: "Project",
    description:
      "Mengarahkan koordinasi teknis dan proses pengembangan proyek.",
  },
  "Tech Lead 2": {
    division: "Project",
    description:
      "Mengarahkan koordinasi teknis dan proses pengembangan proyek.",
  },
  "Kabid SDM": {
    division: "Sumber Daya Manusia",
    description:
      "Mendukung pengembangan anggota dan koordinasi sumber daya manusia.",
  },
  "Koordinator Humas": {
    division: "Humas",
    description: "Mengkoordinasikan kegiatan dan anggota divisi Humas.",
  },
  "Koordinator Pemrograman": {
    division: "Pemrograman",
    description: "Mengkoordinasikan kegiatan dan anggota divisi Pemrograman.",
  },
  "Koordinator Jaringan": {
    division: "Jaringan",
    description: "Mengkoordinasikan kegiatan dan anggota divisi Jaringan.",
  },
  "Koordinator Creative Media": {
    division: "Creative Media",
    description:
      "Mengkoordinasikan kegiatan dan anggota divisi Creative Media.",
  },
  "Koordinator Data": {
    division: "Data",
    description: "Mengkoordinasikan kegiatan dan anggota divisi Data.",
  },
  "Pengurus Humas": {
    division: "Humas",
    description: "Mendukung kegiatan dan program kerja divisi Humas.",
  },
  "Pengurus Pemrograman": {
    division: "Pemrograman",
    description: "Mendukung kegiatan dan program kerja divisi Pemrograman.",
  },
  "Pengurus Jaringan": {
    division: "Jaringan",
    description: "Mendukung kegiatan dan program kerja divisi Jaringan.",
  },
  "Pengurus Creative Media": {
    division: "Creative Media",
    description: "Mendukung kegiatan dan program kerja divisi Creative Media.",
  },
  "Pengurus Data": {
    division: "Data",
    description: "Mendukung kegiatan dan program kerja divisi Data.",
  },
};

export const pengurus: PengurusMember[] = dummyPeriods.flatMap((period) =>
  dummyRoles.map((role, index) => ({
    name: `Nama Pengurus ${index + 1}`,
    role,
    period,
    ...roleDetails[role],
    socialMedia: [
      { platform: "Instagram", url: "https://instagram.com/" },
      { platform: "LinkedIn", url: "https://linkedin.com/in/" },
    ],
    history: [
      { role: "Anggota Pemrograman", period: "2022/2023" },
      { role: "Pengurus Pemrograman", period: "2023/2024" },
      { role, period },
    ],
    image:
      (index + 1) % 9 === 0
        ? undefined
        : `/images/avatar-${(index % 5) + 1}.jpg`,
  })),
);

/** Every period present in `pengurus`, newest first. */
export const pengurusPeriods: string[] = [
  ...new Set(pengurus.map((person) => person.period)),
].sort((a, b) => Number.parseInt(b, 10) - Number.parseInt(a, 10));
