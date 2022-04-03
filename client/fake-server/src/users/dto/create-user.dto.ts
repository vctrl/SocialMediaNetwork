import { IsIn, IsInt, IsString } from 'class-validator';
import { Gender, User } from '../../../../src/types/entities';
export class CreateUserDto implements Omit<User & { password: string }, 'id'> {
  @IsString()
  login: string;

  @IsString()
  password: string;

  @IsString()
  name: string;

  @IsString()
  surname: string;

  @IsInt()
  age: number;

  @IsIn(['MALE', 'FEMALE'])
  sex: Gender;

  @IsString()
  interests: string;

  @IsString()
  city: string;
}
