DROP TYPE IF EXISTS "collection_type";

CREATE TYPE "collection_type" AS ENUM ('NOTE', 'PROBLEM');

CREATE TABLE IF NOT EXISTS "collections" (
    "id" UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    "title" VARCHAR(50) NOT NULL,
    "description" TEXT NOT NULL,
    "emoji" VARCHAR(10) NOT NULL,
    "type" collection_type NOT NULL,
    "created_at" TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    "updated_at" TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    "deleted_at" TIMESTAMP DEFAULT NULL
);