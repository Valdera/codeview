CREATE TABLE IF NOT EXISTS "problems_sources" (
  "id" BIGSERIAL PRIMARY KEY,
  "problem_id" INTEGER NOT NULL,
  "source_id" INTEGER NOT NULL,
  "created_at" TIMESTAMP,
  "updated_at" TIMESTAMP,
  "deleted_at" TIMESTAMP,
  CONSTRAINT fk_problem
    FOREIGN KEY(problem_id)
        REFERENCES problems(id),
  CONSTRAINT fk_source
    FOREIGN KEY(source_id)
        REFERENCES sources(id)
);