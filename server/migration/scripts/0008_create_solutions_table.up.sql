CREATE TABLE IF NOT EXISTS "solutions" (
  "id" BIGSERIAL PRIMARY KEY,
  "problem_id" INTEGER NOT NULL,
  "content" TEXT NOT NULL,
  "created_at" TIMESTAMP,
  "updated_at" TIMESTAMP,
  "deleted_at" TIMESTAMP,
  CONSTRAINT fk_problem
    FOREIGN KEY(problem_id)
        REFERENCES problems(id)
);