-- +goose Up
CREATE TABLE `users` (
  `id`                       VARCHAR(50)  NOT NULL COMMENT "UID",
  `display_name`             VARCHAR(100) NOT NULL COMMENT "name",
  `email`                    VARCHAR(191) NOT NULL COMMENT "Email",
  `icon_image_path`          VARCHAR(20)  NOT NULL COMMENT "icon image path",
  `background_image_path`    VARCHAR(20)  NOT NULL COMMENT "background image path",
  `profile`                  VARCHAR(20)  NOT NULL COMMENT "profile",
  `created_at`               TIMESTAMP   NOT NULL COMMENT "creation date",
  `updated_at`               TIMESTAMP   NOT NULL COMMENT "update date",
  CONSTRAINT `users_pkey` PRIMARY KEY (`id`),
  UNIQUE `users_unique_email` (`email`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4
COMMENT "ユーザー";

-- +goose Down
DROP TABLE users;
