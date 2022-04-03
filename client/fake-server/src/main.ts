import { INestApplication, ValidationPipe } from '@nestjs/common';
import { NestFactory } from '@nestjs/core';
import { AppModule } from './app.module';
import * as cookies from 'cookie-parser';

export let app: INestApplication;

async function bootstrap() {
  app = await NestFactory.create(AppModule);
  app.use(cookies());
  app.useGlobalPipes(new ValidationPipe());
  await app.listen(3000);
}
bootstrap();
