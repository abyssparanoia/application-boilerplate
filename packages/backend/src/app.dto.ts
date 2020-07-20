import { IsString } from 'class-validator'

export class SignInRequestDto {
  @IsString()
  userID!: string

  @IsString()
  password!: string
}
