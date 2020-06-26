-- +migrate Up
CREATE TABLE "users"
(
    "user_id" text NOT NULL UNIQUE,
    "phone" text NOT NULL UNIQUE ,
    "email" text NOT NULL UNIQUE ,
    "password" text NULL ,
    "avatar" text NOT NULL,
    "display_name" text NOT NULL,
    "role" text NOT NULL,
    "birthday" text NOT NULL,
    "position" text NOT NULL,
    "level" text  NOT NULL,
    "verify" text NOT NULL,
    "created_at" TIMESTAMPTZ NOT NULL,
    "updated_at" TIMESTAMPTZ NOT NULL,
    CONSTRAINT users_pkey PRIMARY KEY (user_id)
);
-- +migrate Down
DROP TABLE "users";
-- DROP TABLE "orders";
-- DROP TABLE "product";
-- DROP TABLE "cate";
-- DROP TABLE "card";

