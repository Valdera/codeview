CREATE TABLE IF NOT EXISTS "problems" (
  "id" UUID PRIMARY KEY DEFAULT gen_random_uuid(),
  "title" VARCHAR(50) NOT NULL,
  "emoji" VARCHAR(10) NOT NULL,
  "difficulty_id" UUID NOT NULL,
  "rating" INTEGER DEFAULT 0 NOT NULL,
  CHECK ("rating" >= 0),
  CHECK ("rating" <= 5),
  "created_at" TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  "updated_at" TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  "deleted_at" TIMESTAMP DEFAULT NULL,
  CONSTRAINT fk_difficulty FOREIGN KEY(difficulty_id) REFERENCES difficulties(id)
);