USE `defaultdb`;

CREATE TABLE `users` (
  `id` varchar(50) NOT NULL COMMENT 'firebase auth UID',
  `display_name` varchar(36) NOT NULL COMMENT 'display name',
  `icon_image_path` varchar(3000) NOT NULL COMMENT 'profile icon image path',
  `background_image_path` varchar(3000) NOT NULL COMMENT 'background image path',
  `profile` varchar(1024) NULL COMMENT 'profile text',
  `email` varchar(1024) NULL COMMENT 'email address',
  `created_at` TIMESTAMP DEFAULT CURRENT_TIMESTAMP COMMENT 'creation date',
  `updated_at` TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT 'updation date',
  `deleted_at` TIMESTAMP NULL COMMENT 'delete date',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4
COMMENT "users table";
