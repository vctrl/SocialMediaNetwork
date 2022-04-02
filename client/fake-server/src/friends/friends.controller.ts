import { Controller, Get, Param } from '@nestjs/common';
import { FriendsService } from './friends.service';

@Controller('friends')
export class FriendsController {
  constructor(private readonly friendsService: FriendsService) {}

  @Get(':id')
  get(@Param('id') id: string) {
    return this.friendsService.get(+id);
  }
}
