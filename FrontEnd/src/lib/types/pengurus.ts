export interface PengurusSosmed {
  platform: string;
  username: string;
  url: string;
  is_primary: boolean;
}

export interface PengurusPublic {
  id: number;
  photo_url: string;
  divisi: string;
  name: string;
  position: string;
  sosmed: PengurusSosmed[];
  start_periode_year: number;
  end_periode_year: number;
}

export interface Pengurus extends PengurusPublic {
  id_user: number;
  email: string;
  created_at: string;
  updated_at: string;
}
