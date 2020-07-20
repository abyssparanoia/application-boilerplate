import { Controller, Get, Post, Body } from '@nestjs/common'
import { SignInRequest } from './user.request'
import { UserService } from './user.service'

@Controller('users')
export class UserController {
  constructor(private readonly userService: UserService) {}

  @Get()
  getHello() {
    return this.userService.getById('DUMMY_USER_ID')
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
