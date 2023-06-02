CREATE TABLE IF NOT EXISTS "problems_tags" (
  "id" UUID PRIMARY KEY DEFAULT gen_random_uuid(),
  "problem_id" UUID NOT NULL,
  "tag_id" UUID NOT NULL,
  "created_at" TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  "updated_at" TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  "deleted_at" TIMESTAMP DEFAULT NULL,
  CONSTRAINT fk_problem FOREIGN KEY(problem_id) REFERENCES problems(id),
  CONSTRAINT fk_tag FOREIGN KEY(tag_id) REFERENCES tags(id)
);