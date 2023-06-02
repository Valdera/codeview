CREATE TABLE IF NOT EXISTS "questions" (
  "id" UUID PRIMARY KEY DEFAULT gen_random_uuid(),
  "problem_id" UUID NOT NULL,
  "content" TEXT NOT NULL,
  "created_at" TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  "updated_at" TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  "deleted_at" TIMESTAMP DEFAULT NULL,
  CONSTRAINT fk_problem FOREIGN KEY(problem_id) REFERENCES problems(id)
);