import { Module } from '@nestjs/common'
import { ConfigModule } from '@nestjs/config'
import { TypeOrmModule } from '@nestjs/typeorm'
import { UserModule } from './user/user.module'
import { FirebaseModule } from 'nestjs-firebase'

@Module({
  imports: [
    ConfigModule,
    TypeOrmModule.forRoot({
      type: 'mysql',
      host: 'localhost',
      port: 3306,
      username: 'root',
      password: 'password',
      database: 'defaultdb',
      entities: ['./**/*.entity{.ts,.js}'],
      synchronize: true
    }),
    FirebaseModule.forRoot({
      googleApplicationCredential: './serviceAccount.json'
    }),
    UserModule
  ]
})
export class AppModule {}
