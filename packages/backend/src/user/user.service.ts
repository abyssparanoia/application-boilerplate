import { Injectable } from '@nestjs/common'
import { InjectRepository } from '@nestjs/typeorm'
import { User } from './user.entity'
import { Repository } from 'typeorm'

@Injectable()
export class UserService {
  constructor(
    @InjectRepository(User)
    private userRepository: Repository<User>
  ) {}
  getHello(): string {
    return 'Hello World!'
  }

  getById(id: string) {
    return this.userRepository.findOneOrFail(id)
  }
}
