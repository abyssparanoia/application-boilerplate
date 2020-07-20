import { IsString } from 'class-validator'

export class SignInRequest {
  @IsString()
  userID!: string

  @IsString()
  password!: string
}
