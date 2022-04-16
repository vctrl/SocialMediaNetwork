import { HttpException, Injectable } from '@nestjs/common';
import { CreateUserDto } from './dto/create-user.dto';
import { UpdateUserDto } from './dto/update-user.dto';
import { User } from '../../../src/types/entities';

let nextId = 2;
const users = new Map<User['id'], User & { password: string }>([
  [
    1,
    {
      id: 1,
      login: 'user',
      password: 'password',
      name: 'John',
      surname: 'Doe',
      sex: 'MALE',
      age: 22,
      city: 'Moscow',
      interests: 'Web development',
    },
  ],
]);

@Injectable()
export class UsersService {
  async create(dto: CreateUserDto) {
    if (await this.findByLogin(dto.login)) throw new HttpException('User with this login already exists', 400);

    const id = nextId++;
    users.set(id, {
      age: dto.age,
      city: dto.city,
      id: id,
      interests: dto.interests,
      login: dto.login,
      name: dto.name,
      sex: dto.sex,
      surname: dto.surname,
      password: dto.password,
    });

    return this.getOne(id);
  }

  async findAll() {
    return Array.from(users.values());
  }

  async getOne(id: User['id']) {
    return users.get(id);
  }

  async getMany(ids: User['id'][]) {
    return Promise.all(ids.map((id) => this.getOne(id)));
  }

  async findByLogin(login: User['login']) {
    return Array.from(users.values()).find((user) => user.login === login);
  }

  async update(id: User['id'], dto: UpdateUserDto) {
    const user = users.get(id);
    for (const key in dto) {
      user[key] = dto[key];
    }

    return user;
  }

  async remove(id: User['id']) {
    const user = await this.getOne(id);
    if (!user) return null;

    users.delete(id);
    return user;
  }

  publicData(user: User & { password?: string }) {
    const { password, ...publicUser } = user;
    return publicUser;
  }
}
