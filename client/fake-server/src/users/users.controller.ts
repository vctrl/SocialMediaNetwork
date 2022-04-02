import { Controller, Get, Post, Body, Patch, Param, Delete, HttpException, Response } from '@nestjs/common';
import { UsersService } from './users.service';
import { CreateUserDto } from './dto/create-user.dto';
import { UpdateUserDto } from './dto/update-user.dto';
import { LoginDto } from './dto/login.dto';
import { Response as EResponse } from 'express';

@Controller('users')
export class UsersController {
  constructor(private readonly usersService: UsersService) {}

  @Post('register')
  register(@Body() dto: CreateUserDto) {
    return this.usersService.create(dto);
  }

  @Post('login')
  async login(@Body() dto: LoginDto, @Response() res: EResponse) {
    const user = await this.usersService.findByLogin(dto.login);
    if (!user || user.password !== dto.password) throw new HttpException('Wrong credentials', 403);

    res.cookie('userId', user.id);
    res.send(user);
  }

  @Get(':ids')
  findMany(@Param('ids') ids: string) {
    const idsParsed = ids.split(',').map((id) => +id);
    return this.usersService.getMany(idsParsed);
  }

  @Patch(':id')
  update(@Param('id') id: string, @Body() updateUserDto: UpdateUserDto) {
    return this.usersService.update(+id, updateUserDto);
  }

  @Delete(':id')
  remove(@Param('id') id: string) {
    return this.usersService.remove(+id);
  }
}
