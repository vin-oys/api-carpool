CREATE TYPE user_role AS ENUM (
  'super_administrator',
  'administrator',
  'driver',
  'passenger'
);

CREATE TYPE gender AS ENUM (
  'male',
  'female'
);

CREATE TYPE category AS ENUM (
  'adult',
  'child'
);

CREATE TABLE "users" (
											 "id" INT PRIMARY KEY,
											 "username" INT UNIQUE,
											 "password" VARCHAR NOT NULL,
											 "contact_number" VARCHAR UNIQUE,
											 "created_at" TIMESTAMPTZ NOT NULL DEFAULT (now()),
											 "updated_at" TIMESTAMPTZ NOT NULL,
											 "role_id" user_role
);

CREATE TABLE "models" (
												"id" INT PRIMARY KEY,
												"name" VARCHAR UNIQUE NOT NULL,
												"created_at" TIMESTAMPTZ NOT NULL DEFAULT (now()),
												"updatedAt" TIMESTAMPTZ NOT NULL
);

CREATE TABLE "cars" (
											"id" INT PRIMARY KEY,
											"plate" VARCHAR UNIQUE NOT NULL,
											"model_id" INT NOT NULL,
											"created_at" TIMESTAMPTZ NOT NULL DEFAULT (now()),
											"updated_at" TIMESTAMPTZ NOT NULL
);

CREATE TABLE "schedules" (
													 "id" INT PRIMARY KEY,
													 "departure_date" DATE NOT NULL,
													 "departure_time" TIME NOT NULL,
													 "pickup" JSONB NOT NULL,
													 "dropoff" JSONB NOT NULL,
													 "driver_id" INT,
													 "car_id" INT,
													 "created_at" TIMESTAMPTZ NOT NULL DEFAULT (now()),
													 "updated_at" TIMESTAMPTZ NOT NULL
);

CREATE TABLE "schedule_passengers" (
																		 "id" INT PRIMARY KEY,
																		 "schedule_id" INT,
																		 "passenger_id" INT,
																		 "gender" gender,
																		 "category" category,
																		 "seat" INT,
																		 "created_at" TIMESTAMPTZ NOT NULL DEFAULT (now()),
																		 "updated_at" TIMESTAMPTZ NOT NULL
);

COMMENT ON COLUMN "schedules"."driver_id" IS 'When carpool confirmed';

COMMENT ON COLUMN "schedules"."car_id" IS 'When carpool confirmed';

ALTER TABLE "cars" ADD FOREIGN KEY ("model_id") REFERENCES "models" ("id");

ALTER TABLE "schedules" ADD FOREIGN KEY ("driver_id") REFERENCES "users" ("id");

ALTER TABLE "schedules" ADD FOREIGN KEY ("car_id") REFERENCES "cars" ("id");

ALTER TABLE "schedule_passengers" ADD FOREIGN KEY ("schedule_id") REFERENCES "schedules" ("id");

ALTER TABLE "schedule_passengers" ADD FOREIGN KEY ("passenger_id") REFERENCES "users" ("id");
