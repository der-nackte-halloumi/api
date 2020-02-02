ALTER TABLE "shops" DROP CONSTRAINT IF EXISTS "shops_city_id_fkey";

ALTER TABLE "shops_products" DROP CONSTRAINT IF EXISTS "shops_products_product_id_fkey";

ALTER TABLE "shops_products" DROP CONSTRAINT IF EXISTS "shops_products_shop_id_fkey";

ALTER TABLE "cities" DROP CONSTRAINT IF EXISTS "cities_country_id_fkey";

ALTER TABLE "products" DROP CONSTRAINT IF EXISTS "products_category_id_fkey";

ALTER TABLE "product_translations" DROP CONSTRAINT IF EXISTS ("product_translations_locale_id_fkey";

ALTER TABLE "product_translations" DROP CONSTRAINT IF EXISTS "product_translations_product_id_fkey";

ALTER TABLE "categories_translations" DROP CONSTRAINT IF EXISTS "categories_translations_locale_id_fkey";

ALTER TABLE "categories_translations" DROP CONSTRAINT IF EXISTS "categories_translations_category_id_fkey";

DROP TABLE IF EXISTS "shops";

DROP TABLE IF EXISTS "shops_products";

DROP TABLE IF EXISTS "cities";

DROP TABLE IF EXISTS "countries";

DROP TABLE IF EXISTS "products";

DROP TABLE IF EXISTS "product_translations";

DROP TABLE IF EXISTS "categories";

DROP TABLE IF EXISTS "categories_translations";

DROP TABLE IF EXISTS "locales";

