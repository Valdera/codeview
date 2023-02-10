CREATE TABLE IF NOT EXISTS "problems" (
  "id" BIGSERIAL PRIMARY KEY,
  "title" VARCHAR(50) NOT NULL,
  "difficulty_id" INTEGER NOT NULL,
  "rating" INTEGER DEFAULT 0 NOT NULL, 
  CHECK ("rating" >= 0),
  CHECK ("rating" <= 5),
  "created_at" TIMESTAMP,
  "updated_at" TIMESTAMP,
  "deleted_at" TIMESTAMP,
  CONSTRAINT fk_difficulty
    FOREIGN KEY(difficulty_id)
        REFERENCES difficulties(id)
);