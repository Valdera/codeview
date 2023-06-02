CREATE TABLE IF NOT EXISTS "problems_sources" (
  "id" UUID PRIMARY KEY DEFAULT gen_random_uuid(),
  "problem_id" UUID NOT NULL,
  "source_id" UUID NOT NULL,
  "created_at" TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  "updated_at" TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  "deleted_at" TIMESTAMP DEFAULT NULL,
  CONSTRAINT fk_problem FOREIGN KEY(problem_id) REFERENCES problems(id),
  CONSTRAINT fk_source FOREIGN KEY(source_id) REFERENCES sources(id)
);