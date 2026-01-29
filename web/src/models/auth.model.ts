// --- Login Related Types ---
export interface LoginResponseData {
  token: string;
  expire: string; // Or Date, depending on serialization. Let's assume string for now.
}

// Define the structure for login credentials
export interface LoginCredentials {
  username?: string;
  password?: string;
}

export interface RegistrationPayload {
  username: string;
  email: string;
  mobile?: string;
  password?: string; // Should ideally be required for registration logic by the calling component
}

// --- Change Password Related Types ---
export interface ChangePasswordRequest {
  oldPassword: string;
  newPassword: string;
}

export interface ChangePasswordResponse {
  code: number;
  message: string;
  data?: null;
}
