export interface User {
  id: number;
  username: string;
  email: string;
  role: string;
  full_name: string;
}

export interface UserPatch {
  username?: string;
  email?: string;
  fullname?: string;
}

export interface ChangePasswordPayload {
  old_password: string;
  new_password: string;
}

export interface AdminChangePasswordPayload {
  new_password: string;
}
