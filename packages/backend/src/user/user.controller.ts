import { Controller, Get, Post, Param, UseGuards } from '@nestjs/common'
import { UserService } from './user.service'
import { AuthGuard } from '../auth/auth.guard'
import { AuthUser, IAuthUser } from '../auth/auth.decorator'

@Controller('users')
export class UserController {
  constructor(private readonly userService: UserService) {}

  @Get(':userId')
  get(@Param('userId') userId: string) {
    return this.userService.getById(userId)
  }

  @Post('authorize')
  @UseGuards(AuthGuard)
  async signIn(@AuthUser() param: IAuthUser) {
    console.log(param)
    return {
      accessToken: 'accessToken',
      refreshToken: 'refreshToken'
    }
  }
}
