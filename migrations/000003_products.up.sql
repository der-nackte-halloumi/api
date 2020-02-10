ALTER TABLE IF EXISTS "product_translations"
  RENAME TO "products_translations";
ALTER TABLE IF EXISTS "products_translations"
ADD
  COLUMN "value" text NOT NULL;
ALTER TABLE IF EXISTS "categories_translations"
ADD
  COLUMN "value" text NOT NULL;
-- unique combination of product id and locale id
ALTER TABLE "products_translations" DROP CONSTRAINT IF EXISTS "product_translations_locale_id_key";
ALTER TABLE "products_translations" DROP CONSTRAINT IF EXISTS "product_translations_product_id_key";
ALTER TABLE "products_translations"
ADD
  CONSTRAINT "productss_translations_product_id_locale_id_key" UNIQUE ("product_id", "locale_id");
-- unique combination of category id and locale id
ALTER TABLE "categories_translations" DROP CONSTRAINT IF EXISTS "categories_translations_category_id_key";
ALTER TABLE "categories_translations" DROP CONSTRAINT IF EXISTS "categories_translations_locale_id_key";
ALTER TABLE "categories_translations"
ADD
  CONSTRAINT "categories_translations_category_id_locale_id_key" UNIQUE ("category_id", "locale_id");
-- unique combination of shop id and product id
ALTER TABLE "shops_products" DROP CONSTRAINT IF EXISTS "shops_products_product_id_key";
ALTER TABLE "shops_products" DROP CONSTRAINT IF EXISTS "shops_products_shop_id_key";
ALTER TABLE "shops_products"
ADD
  CONSTRAINT "shops_products_shop_id_product_id_key" UNIQUE ("shop_id", "product_id");
