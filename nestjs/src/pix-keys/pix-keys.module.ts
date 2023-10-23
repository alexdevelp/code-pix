import { Module } from '@nestjs/common';
import { TypeOrmModule } from '@nestjs/typeorm';
import { PixKey } from './entities/pix-key.entity';
import { PixKeysService } from './pix-keys.service';
import { PixKeysController } from './pix-keys.controller';
import { BankAccount } from '../bank-accounts/entities/bank-account.entity';

@Module({
  imports: [TypeOrmModule.forFeature([PixKey, BankAccount])],
  controllers: [PixKeysController],
  providers: [PixKeysService],
})
export class PixKeysModule {}
