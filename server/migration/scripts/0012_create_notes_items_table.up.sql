CREATE TABLE IF NOT EXISTS "notes_items" (
    "id" UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    "note_id" UUID NOT NULL,
    "content" TEXT NOT NULL,
    "position" INTEGER NOT NULL,
    "created_at" TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    "updated_at" TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    "deleted_at" TIMESTAMP DEFAULT NULL,
    CONSTRAINT fk_note FOREIGN KEY(note_id) REFERENCES notes(id)
);