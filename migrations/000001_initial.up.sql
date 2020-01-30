CREATE TABLE "shops" (
  "id" uuid PRIMARY KEY,
  "city_id" uuid
);

CREATE TABLE "shops_products" (
  "id" uuid PRIMARY KEY,
  "shop_id" uuid UNIQUE,
  "product_id" uuid UNIQUE
);

CREATE TABLE "cities" (
  "id" uuid PRIMARY KEY,
  "country_id" uuid,
  "lat" float,
  "long" float
);

CREATE TABLE "countries" (
  "id" uuid PRIMARY KEY
);

CREATE TABLE "products" (
  "id" uuid PRIMARY KEY,
  "category_id" uuid
);

CREATE TABLE "product_translations" (
  "id" uuid PRIMARY KEY,
  "locale_id" uuid UNIQUE,
  "product_id" uuid UNIQUE
);

CREATE TABLE "categories" (
  "id" uuid PRIMARY KEY
);

CREATE TABLE "categories_translations" (
  "id" uuid PRIMARY KEY,
  "locale_id" uuid UNIQUE,
  "category_id" uuid UNIQUE
);

CREATE TABLE "locales" (
  "id" uuid PRIMARY KEY,
  "locale" text
);

ALTER TABLE "shops" ADD FOREIGN KEY ("city_id") REFERENCES "cities" ("id");

ALTER TABLE "shops_products" ADD FOREIGN KEY ("shop_id") REFERENCES "shops" ("id");

ALTER TABLE "shops_products" ADD FOREIGN KEY ("product_id") REFERENCES "products" ("id");

ALTER TABLE "cities" ADD FOREIGN KEY ("country_id") REFERENCES "countries" ("id");

ALTER TABLE "products" ADD FOREIGN KEY ("category_id") REFERENCES "categories" ("id");

ALTER TABLE "product_translations" ADD FOREIGN KEY ("locale_id") REFERENCES "locales" ("id");

ALTER TABLE "product_translations" ADD FOREIGN KEY ("product_id") REFERENCES "products" ("id");

ALTER TABLE "categories_translations" ADD FOREIGN KEY ("locale_id") REFERENCES "locales" ("id");

ALTER TABLE "categories_translations" ADD FOREIGN KEY ("category_id") REFERENCES "categories" ("id");
