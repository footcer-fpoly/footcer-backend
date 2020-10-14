-- +migrate Up
CREATE TABLE "users"
(
    "user_id" text NOT NULL UNIQUE,
    "phone" text NOT NULL UNIQUE ,
    "email" text NOT NULL ,
    "password" text NULL ,
    "avatar" text NOT NULL,
    "display_name" text NOT NULL,
    "role" text NOT NULL,
    "birthday" text NOT NULL,
    "position" text NOT NULL,
    "level" text  NOT NULL,
    "verify" text NOT NULL,
    "token_notify" text NOT NULL,
    "created_at" TIMESTAMPTZ NOT NULL,
    "updated_at" TIMESTAMPTZ NOT NULL,
    CONSTRAINT users_pkey PRIMARY KEY (user_id)
);

CREATE TABLE "stadium"(
"stadium_id" text NOT NULL UNIQUE,
"user_id" text NOT NULL,
"name_stadium" text NOT NULL,
"address" text NOT NULL,
"description" text NOT NULL,
"image" text NOT NULL,
"category" text NOT NULL,
"latitude" numeric NOT NULL,
"longitude" numeric NOT NULL,
"ward" text NOT NULL,
"district" text NOT NULL,
"city" text NOT NULL,
"created_at" TIMESTAMPTZ NOT NULL,
"updated_at" TIMESTAMPTZ NOT NULL,

FOREIGN KEY ("user_id") REFERENCES "users" ("user_id"),
CONSTRAINT stadium_pkey PRIMARY KEY (stadium_id)

);
CREATE TABLE "stadium_collage"(
"stadium_collage_id" text NOT NULL UNIQUE,
"stadium_id" text NOT NULL,
"name_stadium_collage" text NOT NULL,
"amount_people" text NOT NULL,
"start_time" text NOT NULL,
"end_time" text NOT NULL,
"play_time" text NOT NULL,
"created_at" TIMESTAMPTZ NOT NULL,
"updated_at" TIMESTAMPTZ NOT NULL,

FOREIGN KEY (stadium_id) REFERENCES stadium (stadium_id),
CONSTRAINT stadium_collage_pkey PRIMARY KEY (stadium_collage_id)
);

CREATE TABLE "stadium_details"(
"stadium_detail_id" text NOT NULL UNIQUE,
"stadium_collage_id" text NOT NULL,
"start_time_detail" text NOT NULL,
"end_time_detail" text NOT NULL,
"price" numeric NOT NULL,
"description" text NOT NULL,
"has_order" BOOLEAN NOT NULL DEFAULT FALSE,
"created_at" TIMESTAMPTZ NOT NULL,
"updated_at" TIMESTAMPTZ NOT NULL,

FOREIGN KEY (stadium_collage_id) REFERENCES stadium_collage (stadium_collage_id),
CONSTRAINT stadium_details_pkey PRIMARY KEY (stadium_detail_id)
);

CREATE TABLE "service"(
"service_id" text NOT NULL UNIQUE,
"stadium_id" text NOT NULL,
"name_service" text NOT NULL,
"price_service" text NOT NULL,
"image" text NOT NULL,
FOREIGN KEY (stadium_id) REFERENCES stadium (stadium_id),
CONSTRAINT service_pkey PRIMARY KEY (service_id)
);
CREATE TABLE "review"
(
"review_id" text NOT NULL UNIQUE,
"user_id" text NOT NULL,
"stadium_id" text NOT NULL,
"content" text NOT NULL,
"rate" float NOT NULL,
"created_at_rv" DATE NOT NULL,
"updated_at_rv" DATE NOT NULL,

FOREIGN KEY (stadium_id) REFERENCES stadium (stadium_id),
FOREIGN KEY (user_id) REFERENCES users (user_id),
CONSTRAINT review_pkey PRIMARY KEY (review_id)

);
CREATE TABLE "team"
(
"team_id" text NOT NULL UNIQUE,
"leader_id" text NOT NULL,
"name" text NOT NULL,
"level" text NOT NULL,
"place" text NOT NULL,
"description" text NOT NULL,
"avatar" text NOT NULL,
"background" text NOT NULL,
"created_at" DATE NOT NULL,
"updated_at" DATE NOT NULL,

FOREIGN KEY (leader_id) REFERENCES users (user_id),
CONSTRAINT team_pkey PRIMARY KEY (team_id)

);
CREATE TABLE "team_details"
(
"team_details_id" text NOT NULL UNIQUE,
"teams_id" text NOT NULL,
"user_id" text NOT NULL,
"role_team" text NOT NULL,
"accept" text NOT NULL,
"created_at" DATE NOT NULL,
"updated_at" DATE NOT NULL,

FOREIGN KEY (teams_id) REFERENCES team (team_id),
FOREIGN KEY (user_id) REFERENCES users (user_id),
CONSTRAINT team_details_pkey PRIMARY KEY (team_details_id)
);
CREATE TABLE "orders"
(
"order_id" text NOT NULL UNIQUE,
"user_id" text NOT NULL,
"stadium_detail_id" text NOT NULL,
"time" text NOT NULL,
"description" text NOT NULL,
"price" numeric NOT NULL,
"finish" BOOLEAN NOT NULL,
"accept" BOOLEAN NOT NULL,
"order_created_at" TIMESTAMPTZ NOT NULL,
"order_updated_at" TIMESTAMPTZ NOT NULL,

FOREIGN KEY (stadium_detail_id) REFERENCES stadium_details (stadium_detail_id),
FOREIGN KEY (user_id) REFERENCES users (user_id),

CONSTRAINT order_id_pkey PRIMARY KEY (order_id)
);

CREATE TABLE "game"
(
"game_id" text NOT NULL UNIQUE,
"date" DATE NOT NULL,
"hour" TIME NOT NULL,
"type" text NOT NULL,
"score" text NOT NULL,
"description" text NOT NULL,
"finish" text NOT NULL,
"stadium_id" text NOT NULL,
"team_id_host" text NOT NULL,
"team_id_guest" text  NULL,
"game_created_at" TIMESTAMPTZ NOT NULL,
"game_updated_at" TIMESTAMPTZ NOT NULL,

FOREIGN KEY (stadium_id) REFERENCES stadium (stadium_id),
FOREIGN KEY (team_id_host) REFERENCES team (team_id),
FOREIGN KEY (team_id_guest) REFERENCES team (team_id),

CONSTRAINT game_id_pkey PRIMARY KEY (game_id)
);


CREATE TABLE "game_temp"
(
"game_temp_id" text NOT NULL UNIQUE,
"game_id" text NOT NULL,
"team_id" text NOT NULL,

FOREIGN KEY (game_id) REFERENCES game (game_id),
FOREIGN KEY (team_id) REFERENCES team (team_id),

CONSTRAINT game_temp_id_pkey PRIMARY KEY (game_temp_id)
);

-- +migrate Down
DROP TABLE "review";
DROP TABLE "service";
DROP TABLE "orders";
DROP TABLE "stadium_collage";
DROP TABLE "team_details";
DROP TABLE "game_temp";
DROP TABLE "game";
DROP TABLE "team";
DROP TABLE "stadium";
DROP TABLE "users";
