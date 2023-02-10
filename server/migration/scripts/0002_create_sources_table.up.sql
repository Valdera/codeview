CREATE TABLE IF NOT EXISTS "sources" (
  "id" BIGSERIAL PRIMARY KEY,
  "label" VARCHAR(50) UNIQUE NOT NULL,
  "color" VARCHAR(7) NOT NULL,
  "created_at" TIMESTAMP,
  "updated_at" TIMESTAMP,
  "deleted_at" TIMESTAMP
);