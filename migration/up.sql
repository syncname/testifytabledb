CREATE TABLE "users" (
    "name" varchar PRIMARY KEY,
    "email" varchar UNIQUE NOT NULL,
    "created_at" timestamptz NOT NULL DEFAULT (now())
);