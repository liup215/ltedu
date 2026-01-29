export interface InitSuperUserRequest {
  username: string;
  password: string;
  email: string;
  mobile: string;
}

export interface InitSuperUserResponseData {
  id: number;
  username: string;
  email: string;
  mobile: string;
}
