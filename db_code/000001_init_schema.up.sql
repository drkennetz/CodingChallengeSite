CREATE TYPE "dev_level" AS ENUM (
  'beginner',
  'junior',
  'midlevel',
  'senior'
);

CREATE TABLE "accounts" (
  "id" bigserial PRIMARY KEY,
  "full_name" varchar NOT NULL,
  "email" varchar NOT NULL,
  "opted_in" boolean NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now()),
  "updated_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "users" (
  "id" bigserial PRIMARY KEY,
  "account_id" bigint NOT NULL,
  "admin_user" boolean NOT NULL DEFAULT false,
  "username" varchar NOT NULL,
  "password" varchar NOT NULL,
  "grade" dev_level NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now()),
  "updated_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "categories" (
  "id" bigserial PRIMARY KEY,
  "category" varchar NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "questions" (
  "id" bigserial PRIMARY KEY,
  "challenge_name" varchar NOT NULL,
  "description" text NOT NULL,
  "example" text NOT NULL,
  "difficulty" int4 NOT NULL,
  "complexity" text NOT NULL,
  "completion_time" int4 NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now()),
  "updated_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "question_categories" (
  "id" bigserial PRIMARY KEY,
  "question_id" bigint NOT NULL,
  "category_id" bigint NOT NULL
);

CREATE TABLE "question_test_cases" (
  "id" bigserial PRIMARY KEY,
  "question_id" bigint NOT NULL,
  "inputs" json NOT NULL,
  "outputs" json NOT NULL
);

CREATE TABLE "user_question_score" (
  "id" bigserial PRIMARY KEY,
  "user_id" bigint NOT NULL,
  "question_id" bigint NOT NULL,
  "score" int4 NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now()),
  "updated_at" timestamptz NOT NULL DEFAULT (now())
);

ALTER TABLE "users" ADD FOREIGN KEY ("account_id") REFERENCES "accounts" ("id");

ALTER TABLE "question_categories" ADD FOREIGN KEY ("question_id") REFERENCES "questions" ("id");

ALTER TABLE "question_categories" ADD FOREIGN KEY ("category_id") REFERENCES "categories" ("id");

ALTER TABLE "question_test_cases" ADD FOREIGN KEY ("question_id") REFERENCES "questions" ("id");

ALTER TABLE "user_question_score" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id");

ALTER TABLE "user_question_score" ADD FOREIGN KEY ("question_id") REFERENCES "questions" ("id");

CREATE UNIQUE INDEX ON "accounts" ("email");

CREATE UNIQUE INDEX ON "users" ("account_id", "username");

CREATE INDEX ON "users" ("grade");

CREATE INDEX ON "question_categories" ("category_id");

CREATE INDEX ON "question_test_cases" ("question_id");

CREATE INDEX ON "user_question_score" ("user_id");

CREATE INDEX ON "user_question_score" ("question_id");
