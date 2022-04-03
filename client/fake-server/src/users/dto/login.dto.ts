import { IsString } from 'class-validator';
import { User } from '../../../../src/types/entities';

export class LoginDto {
  @IsString()
  login: User['login'];

  @IsString()
  password: string;
}
