import { User } from '../../../../src/types/entities';

export type LoginDto = {
  login: User['login'];
  password: string;
};
