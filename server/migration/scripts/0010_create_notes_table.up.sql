DROP TYPE IF EXISTS "note_status";

CREATE TYPE "note_status" AS ENUM ('DRAFT', 'PUBLISHED');

CREATE TABLE IF NOT EXISTS "notes" (
    "id" UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    "title" VARCHAR(50) NOT NULL,
    "emoji" VARCHAR(10) NOT NULL,
    "status" note_status NOT NULL DEFAULT 'DRAFT',
    "created_at" TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    "updated_at" TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    "deleted_at" TIMESTAMP DEFAULT NULL
);