import { Controller, Get, Post, Body, Param } from '@nestjs/common'
import { SignInRequest } from './user.request'
import { UserService } from './user.service'

@Controller('users')
export class UserController {
  constructor(private readonly userService: UserService) {}

  @Get(':userId')
  get(@Param('userId') userId: string) {
    return this.userService.getById(userId)
  }

  @Post()
  async signIn(@Body() signInRequestDto: SignInRequest) {
    // eslint-disable-next-line no-console
    console.log(signInRequestDto)
    return {
      accessToken: 'accessToken',
      refreshToken: 'refreshToken'
    }
  }
}
