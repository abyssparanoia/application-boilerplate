import { createParamDecorator, ExecutionContext } from '@nestjs/common'
import { auth } from 'firebase-admin'

export type IAuthUser = auth.DecodedIdToken

export const AuthUser = createParamDecorator(
  (_: unknown, ctx: ExecutionContext): IAuthUser => {
    const request = ctx.switchToHttp().getRequest()
    return { ...request.decodedIdToken }
  }
)
