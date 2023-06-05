CREATE TABLE IF NOT EXISTS "collections_items" (
    "id" UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    "collection_id" UUID NOT NULL,
    "item_id" TEXT NOT NULL,
    "position" INTEGER NOT NULL,
    "created_at" TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    "updated_at" TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    "deleted_at" TIMESTAMP DEFAULT NULL,
    CONSTRAINT fk_collection FOREIGN KEY(collection_id) REFERENCES collections(id)
);