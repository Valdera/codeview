CREATE TABLE IF NOT EXISTS "problems_tags" (
  "id" BIGSERIAL PRIMARY KEY,
  "problem_id" INTEGER NOT NULL,
  "tag_id" INTEGER NOT NULL,
  "created_at" TIMESTAMP,
  "updated_at" TIMESTAMP,
  "deleted_at" TIMESTAMP,
  CONSTRAINT fk_problem
    FOREIGN KEY(problem_id)
        REFERENCES problems(id),
  CONSTRAINT fk_tag
    FOREIGN KEY(tag_id)
        REFERENCES tags(id)
);