export interface User {
  id: number;
  username: string;
  email: string;
  password: string;

  displayName?: string;
  photoURL?: string;
}
