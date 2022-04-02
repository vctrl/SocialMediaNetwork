import { User } from '../../../../src/types/entities';

export type CreateUserDto = Omit<User & { password: string }, 'id'>;
