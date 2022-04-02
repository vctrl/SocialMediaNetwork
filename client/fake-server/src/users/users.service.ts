import { Injectable } from '@nestjs/common';
import { CreateUserDto } from './dto/create-user.dto';
import { UpdateUserDto } from './dto/update-user.dto';
import { User } from '../../../src/types/entities';

let nextId = 1;
const users = new Map<User['id'], User & { password: string }>();

@Injectable()
export class UsersService {
  create(dto: CreateUserDto) {
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

  findAll() {
    return Array.from(users.values());
  }

  getOne(id: User['id']) {
    return users.get(id);
  }

  getMany(ids: User['id'][]) {
    return ids.map((id) => this.getOne(id));
  }

  findByLogin(login: User['login']) {
    return Array.from(users.values()).find((user) => user.login === login);
  }

  update(id: User['id'], dto: UpdateUserDto) {
    const user = users.get(id);
    for (const key in dto) {
      user[key] = dto[key];
    }

    return user;
  }

  remove(id: User['id']) {
    users.delete(id);
  }
}
