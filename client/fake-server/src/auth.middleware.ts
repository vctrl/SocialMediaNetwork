import { Injectable, NestMiddleware } from '@nestjs/common';
import { Request, Response } from 'express';
import { User } from '../../src/types/entities';
import { app } from './main';
import { UsersService } from './users/users.service';

declare global {
  // eslint-disable-next-line @typescript-eslint/no-namespace
  namespace Express {
    interface Request {
      user?: User;
    }
  }
}

@Injectable()
export class AuthMiddleware implements NestMiddleware {
  async use(req: Request, res: Response, next: (error?: any) => void) {
    const userId = Number(req.cookies.userId);
    if (userId && !isNaN(userId)) req.user = await app.get(UsersService).getOne(userId);

    next();
  }
}
