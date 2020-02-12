ALTER TABLE IF EXISTS "categories_translations" DROP COLUMN "value";
ALTER TABLE IF EXISTS "products_translations" DROP COLUMN "value";
ALTER TABLE IF EXISTS "products_translations"
  RENAME TO "product_translations";
-- TODO: add constraints
