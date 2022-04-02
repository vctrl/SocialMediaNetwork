import { Controller, Delete, Get, Param, Post, Request, UseGuards } from '@nestjs/common';
import { Request as ERequest } from 'express';
import { AuthGuard } from 'src/auth.guard';
import { FriendsService } from './friends.service';

@Controller('friends')
export class FriendsController {
  constructor(private readonly friendsService: FriendsService) {}

  @Get()
  @UseGuards(AuthGuard)
  get(@Request() req: ERequest) {
    return this.friendsService.get(req.user.id);
  }

  @Get('requests')
  @UseGuards(AuthGuard)
  getRequests(@Request() req: ERequest) {
    return this.friendsService.getRequestsIncoming(req.user.id);
  }

  @Get('sent_requests')
  @UseGuards(AuthGuard)
  getSentRequests(@Request() req: ERequest) {
    return this.friendsService.getRequestsSent(req.user.id);
  }

  @Post(':userId')
  @UseGuards(AuthGuard)
  sendRequest(@Request() req: ERequest, @Param('userId') userId: string) {
    return this.friendsService.createRequest(req.user.id, +userId);
  }

  @Delete(':userId')
  @UseGuards(AuthGuard)
  async remove(@Request() req: ERequest, @Param('userId') userId: string) {
    return this.friendsService.remove(req.user.id, +userId);
  }

  @Post(':userId/accept')
  @UseGuards(AuthGuard)
  acceptRequest(@Request() req: ERequest, @Param('userId') userId: string) {
    return this.friendsService.acceptRequest(+userId, req.user.id);
  }

  @Delete('requests/:userId')
  @UseGuards(AuthGuard)
  async removeRequest(@Request() req: ERequest, @Param('userId') userId: string) {
    await Promise.all([
      this.friendsService.removeRequest(+userId, req.user.id),
      this.friendsService.removeRequest(req.user.id, +userId),
    ]);
  }
}
