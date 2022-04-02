export type Gender = 'MALE' | 'FEMALE';

export interface User {
  id: number;
  login: string;
  name: string;
  surname: string;
  age: number;
  sex: Gender;
  interests: string;
  city: string;
}
