import { Entity, Column, PrimaryColumn, CreateDateColumn, UpdateDateColumn, DeleteDateColumn } from 'typeorm'

interface IUser {
  id: string
  displayName: string
  iconImagePath: string
  backgroundImagePath: string
  profile?: string
  email?: string
  createdAt: Date
  updatedAt: Date
  deletedAt?: Date
}

@Entity({ name: 'users' })
export class User implements IUser {
  @PrimaryColumn()
  id: string

  @Column({ name: 'display_name', length: 36 })
  displayName: string

  @Column({ name: 'icon_image_path', length: 3000 })
  iconImagePath: string

  @Column({ name: 'background_image_path', length: 3000 })
  backgroundImagePath: string

  @Column({ name: 'profile', length: '1024', nullable: true })
  profile?: string

  @Column({ name: 'email', length: '1024', nullable: true })
  email?: string

  @CreateDateColumn({ name: 'created_at', type: 'timestamp' })
  readonly createdAt: Date

  @UpdateDateColumn({ name: 'updated_at', type: 'timestamp' })
  readonly updatedAt: Date

  @DeleteDateColumn({ name: 'deleted_at', type: 'timestamp' })
  readonly deletedAt?: Date

  new(param: IUser) {
    this.id = param.id
    this.displayName = param.displayName
    this.iconImagePath = param.iconImagePath
    this.backgroundImagePath = param.backgroundImagePath
    this.profile = param.profile
    this.email = param.email
  }
}
