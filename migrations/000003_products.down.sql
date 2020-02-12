ALTER TABLE IF EXISTS "categories_translations" DROP COLUMN "value";
ALTER TABLE IF EXISTS "products_translations" DROP COLUMN "value";
ALTER TABLE "products_translations" DROP CONSTRAINT IF EXISTS "productss_translations_product_id_locale_id_key";
ALTER TABLE "products_translations"
ADD
  CONSTRAINT "product_translations_locale_id_key" UNIQUE ("locale_id");
ALTER TABLE "products_translations"
ADD
  CONSTRAINT "product_translations_product_id_key" UNIQUE ("product_id");
ALTER TABLE "categories_translations" DROP CONSTRAINT IF EXISTS "categories_translations_category_id_locale_id_key";
ALTER TABLE "categories_translations"
ADD
  CONSTRAINT "categories_translations_category_id_key" UNIQUE ("category_id");
ALTER TABLE "categories_translations"
ADD
  CONSTRAINT "categories_translations_locale_id_key" UNIQUE ("locale_id");
ALTER TABLE "shops_products" DROP CONSTRAINT IF EXISTS "shops_products_shop_id_product_id_key";
ALTER TABLE "shops_products"
ADD
  CONSTRAINT "shops_products_product_id_key" UNIQUE ("product_id");
ALTER TABLE "shops_products"
ADD
  CONSTRAINT "shops_products_shop_id_key" UNIQUE ("shop_id");
ALTER TABLE IF EXISTS "products_translations"
  RENAME TO "product_translations";
