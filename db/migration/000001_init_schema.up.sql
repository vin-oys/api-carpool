CREATE TYPE "user_role" AS ENUM (
	'super_administrator',
	'administrator',
	'driver',
	'passenger'
	);

CREATE TYPE "category" AS ENUM (
	'adult',
	'child'
	);

CREATE TYPE "country" AS ENUM (
	'malaysia',
	'singapore'
	);

CREATE TABLE "user"
(
	"id"             SERIAL PRIMARY KEY,
	"username"       VARCHAR UNIQUE NOT NULL,
	"password"       VARCHAR        NOT NULL,
	"firstname"      VARCHAR,
	"lastname"       VARCHAR,
	"contact_number" VARCHAR UNIQUE NOT NULL,
	"created_at"     TIMESTAMP      NOT NULL DEFAULT (now()),
	"updated_at"     TIMESTAMP,
	"role_id"        USER_ROLE      NOT NULL
);

CREATE TABLE "car"
(
	"plate_id"   VARCHAR UNIQUE PRIMARY KEY NOT NULL,
	"pax"        INT                        NOT NULL,
	"created_at" TIMESTAMP                  NOT NULL DEFAULT (now()),
	"updated_at" TIMESTAMP
);

CREATE TABLE "schedule"
(
	"id"               SERIAL PRIMARY KEY,
	"departure_date"   DATE      NOT NULL,
	"departure_time"   TIME      NOT NULL,
	"pickup"           JSONB     NOT NULL,
	"drop_off"         JSONB     NOT NULL,
	"pickup_country"   COUNTRY   NOT NULL,
	"drop_off_country" COUNTRY   NOT NULL,
	"driver_id"        INT,
	"plate_id"         VARCHAR,
	"created_at"       TIMESTAMP NOT NULL DEFAULT (now()),
	"updated_at"       TIMESTAMP,
	CONSTRAINT fk_schedule_driver_id
		FOREIGN KEY ("driver_id")
			REFERENCES "user" ("id"),
	CONSTRAINT fk_schedule_plate_id
		FOREIGN KEY ("plate_id")
			REFERENCES "car" ("plate_id")
);

CREATE TABLE "schedule_passenger"
(
	"id"           SERIAL PRIMARY KEY,
	"schedule_id"  INT,
	"passenger_id" INT       NOT NULL,
	"category"     CATEGORY  NOT NULL,
	"seat"         INT,
	"created_at"   TIMESTAMP NOT NULL DEFAULT (now()),
	"updated_at"   TIMESTAMP,
	CONSTRAINT fk_schedule_passenger_schedule_id
		FOREIGN KEY ("schedule_id") REFERENCES "schedule" ("id"),
	CONSTRAINT fk_schedule_passenger_passenger_id
		FOREIGN KEY ("passenger_id") REFERENCES "user" ("id")
);

COMMENT ON COLUMN "schedule"."driver_id" IS 'When carpool confirmed';

COMMENT ON COLUMN "schedule"."plate_id" IS 'When carpool confirmed';