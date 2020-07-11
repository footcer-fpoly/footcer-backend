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

CREATE TABLE "stadium"(
"stadium_id" text NOT NULL UNIQUE,
"name_stadium" text NOT NULL,
"address" text NOT NULL,
"description" text NOT NULL,
"image" text NOT NULL,
"start_time" text NOT NULL,
"end_time" text NOT NULL,
"category" text NOT NULL,
"latitude" numeric NOT NULL,
"longitude" numeric NOT NULL,
"ward" text NOT NULL,
"district" text NOT NULL,
"city" text NOT NULL,
"time_peak" text NOT NULL,
"time_order" text NOT NULL,
"user_id" text NOT NULL,
"created_at" TIMESTAMPTZ NOT NULL,
"updated_at" TIMESTAMPTZ NOT NULL,

FOREIGN KEY ("user_id") REFERENCES "users" ("user_id"),
CONSTRAINT stadium_pkey PRIMARY KEY (stadium_id)

);
CREATE TABLE "stadium_collage"(
"stadium_collage_id" text NOT NULL UNIQUE,
"name_stadium_collage" text NOT NULL,
"amount_people" text NOT NULL,
"price_normal" numeric NOT NULL,
"price_peak" numeric NOT NULL,
"stadium_id" text NOT NULL,
"created_at" TIMESTAMPTZ NOT NULL,
"updated_at" TIMESTAMPTZ NOT NULL,

FOREIGN KEY (stadium_id) REFERENCES stadium (stadium_id),
CONSTRAINT stadium_collage_pkey PRIMARY KEY (stadium_collage_id)
);
CREATE TABLE "service"(
"service_id" text NOT NULL UNIQUE,
"name_service" text NOT NULL,
"price_service" text NOT NULL,
"image" text NOT NULL,
"stadium_id" text NOT NULL,
FOREIGN KEY (stadium_id) REFERENCES stadium (stadium_id),
CONSTRAINT service_pkey PRIMARY KEY (service_id)
);
CREATE TABLE "review"
(
"review_id" text NOT NULL UNIQUE,
"content" text NOT NULL,
"rate" float NOT NULL,
"stadium_id" text NOT NULL,
"user_id" text NOT NULL,
"created_at" DATE NOT NULL,
"updated_at" DATE NOT NULL,

FOREIGN KEY (stadium_id) REFERENCES stadium (stadium_id),
FOREIGN KEY (user_id) REFERENCES users (user_id),
CONSTRAINT review_pkey PRIMARY KEY (review_id)

);
CREATE TABLE "team"
(
"team_id" text NOT NULL UNIQUE,
"name" text NOT NULL,
"level" text NOT NULL,
"place" text NOT NULL,
"description" text NOT NULL,
"avatar" text NOT NULL,
"background" text NOT NULL,
"leader_id" text NOT NULL,
"created_at" DATE NOT NULL,
"updated_at" DATE NOT NULL,

FOREIGN KEY (leader_id) REFERENCES users (user_id),
CONSTRAINT team_pkey PRIMARY KEY (team_id)

);
CREATE TABLE "team_details"
(
"team_details_id" text NOT NULL UNIQUE,
"team_id" text NOT NULL,
"user_id" text NOT NULL,
"role" text NOT NULL,
"accept" text NOT NULL,
"created_at" DATE NOT NULL,
"updated_at" DATE NOT NULL,

FOREIGN KEY (team_id) REFERENCES team (team_id),
FOREIGN KEY (user_id) REFERENCES users (user_id),
CONSTRAINT team_details_pkey PRIMARY KEY (team_details_id)
);

-- +migrate Down
DROP TABLE "review";
DROP TABLE "service";
DROP TABLE "stadium_collage";
DROP TABLE "stadium";
DROP TABLE "team_details";

DROP TABLE "team";
DROP TABLE "users";

-- DROP TABLE "game";


