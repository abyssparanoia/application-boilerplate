import { Injectable, NotFoundException } from '@nestjs/common'
import { InjectRepository } from '@nestjs/typeorm'
import { User } from './user.entity'
import { Repository } from 'typeorm'

@Injectable()
export class UserService {
  constructor(
    @InjectRepository(User)
    private userRepository: Repository<User>
  ) {}

  async getById(id: string) {
    const user = await this.userRepository.findOne(id)
    if (!user) {
      throw new NotFoundException(`user ${id} not found`)
    }
    return user
  }
}
