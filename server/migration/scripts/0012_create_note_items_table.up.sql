CREATE TABLE IF NOT EXISTS "note_items" (
    "id" UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    "header" VARCHAR(50) NOT NULL,
    "note_id" UUID NOT NULL,
    "content" TEXT NOT NULL,
    "created_at" TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    "updated_at" TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    "deleted_at" TIMESTAMP DEFAULT NULL,
    CONSTRAINT fk_note FOREIGN KEY(note_id) REFERENCES notes(id)
);