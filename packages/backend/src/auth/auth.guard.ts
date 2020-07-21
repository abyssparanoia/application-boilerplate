import { Injectable, CanActivate, ExecutionContext, ForbiddenException } from '@nestjs/common'
import { InjectFirebaseAdmin, FirebaseAdmin } from 'nestjs-firebase'

@Injectable()
export class AuthGuard implements CanActivate {
  constructor(@InjectFirebaseAdmin() private readonly firebase: FirebaseAdmin) {}

  async canActivate(context: ExecutionContext): Promise<boolean> {
    const req = context.switchToHttp().getRequest()
    const { authorization } = req.headers
    if (!authorization || !authorization.startsWith('Bearer ')) {
      return false
    }
    const idToken = authorization.slice(7, authorization.length)

    const decodedIdToken = await this.firebase.auth.verifyIdToken(idToken).catch(err => {
      console.error(err)
      throw new ForbiddenException(`${err.message}`)
    })

    req.decodedIdToken = { ...decodedIdToken }

    return true
  }
}
