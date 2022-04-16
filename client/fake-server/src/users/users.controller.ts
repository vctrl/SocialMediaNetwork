import { Controller, Get, Post, Body, Patch, Param, Delete, HttpException, Response, UseGuards } from '@nestjs/common';
import { UsersService } from './users.service';
import { CreateUserDto } from './dto/create-user.dto';
import { UpdateUserDto } from './dto/update-user.dto';
import { LoginDto } from './dto/login.dto';
import { Response as EResponse } from 'express';
import { GetUser } from './user.decorator';
import { AuthGuard } from './auth.guard';
import { User } from '../../../src/types/entities';

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
    res.send(this.usersService.publicData(user));
  }

  @Post('logout')
  @UseGuards(AuthGuard)
  async logout(@Response() res: EResponse) {
    res.cookie('userId', -1, { maxAge: -1 });
    res.send();
  }

  @Get('me')
  @UseGuards(AuthGuard)
  me(@GetUser() user: User) {
    return this.usersService.publicData(user);
  }

  @Get(':ids')
  async findMany(@Param('ids') ids: string) {
    const idsParsed = ids.split(',').map((id) => +id);
    const users = await this.usersService.getMany(idsParsed);
    return users.map((user) => this.usersService.publicData(user));
  }

  @Patch(':id')
  async update(@Param('id') id: string, @Body() updateUserDto: UpdateUserDto) {
    const user = await this.usersService.update(+id, updateUserDto);
    return this.usersService.publicData(user);
  }

  @Delete(':id')
  async remove(@Param('id') id: string) {
    const user = await this.usersService.remove(+id);
    return this.usersService.publicData(user);
  }
}
