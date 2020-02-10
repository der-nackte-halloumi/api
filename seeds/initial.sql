INSERT INTO shops (id, name, address, lat, long)
VALUES
  (
    '7c3153b9-edb0-4ae4-a068-8c9d06532ea5',
    'Original Unverpackt',
    'Wiener Straße 16, 10999 Berlin',
    52.4978374,
    13.428037
  ),
  (
    'd7106377-c6e3-4d4e-bf49-b077e719fbe7',
    'YES FUTURE Positive Supermarket',
    'Carrer de Viladomat, 66, 08015 Barcelona, Spain',
    41.3781588,
    2.1572628
  );
INSERT INTO locales (id, locale)
VALUES
  ('469cfec0-b8f4-4694-b786-055fc4c449ec', 'en'),
  ('58728ba5-bc77-473f-9475-4a7d658897c2', 'de'),
  ('dd25b32b-50c5-4921-893a-d84aae90c38b', 'es');
INSERT INTO categories (id)
VALUES
  -- nuts
  ('bba08683-d337-4033-a087-bdb5366ca64c'),
  -- cheese
  ('4f5061bc-a388-4f1a-918f-cc894f1cd8d2');
INSERT INTO categories_translations (id, locale_id, category_id, value)
VALUES
  (
    'fc713602-b637-4b65-8eab-a62bf6dcec8c',
    '469cfec0-b8f4-4694-b786-055fc4c449ec',
    'bba08683-d337-4033-a087-bdb5366ca64c',
    'nuts'
  ),
  (
    '92d58005-bd79-4d24-b45c-4e80ca1d7133',
    '58728ba5-bc77-473f-9475-4a7d658897c2',
    'bba08683-d337-4033-a087-bdb5366ca64c',
    'Nüsse'
  ),
  (
    '5c4df4b8-bd02-43ea-922b-7bcf16f74940',
    'dd25b32b-50c5-4921-893a-d84aae90c38b',
    'bba08683-d337-4033-a087-bdb5366ca64c',
    'nueces'
  ),
  (
    '164bcccd-8266-4795-abfc-deb57c7d5712',
    '469cfec0-b8f4-4694-b786-055fc4c449ec',
    '4f5061bc-a388-4f1a-918f-cc894f1cd8d2',
    'cheese'
  ),
  (
    '4a33f3b5-3543-4896-b2b3-82993b8edced',
    '58728ba5-bc77-473f-9475-4a7d658897c2',
    '4f5061bc-a388-4f1a-918f-cc894f1cd8d2',
    'Käse'
  ),
  (
    '67a1d3ad-ef44-45a3-9df2-9cbf0648a061',
    'dd25b32b-50c5-4921-893a-d84aae90c38b',
    '4f5061bc-a388-4f1a-918f-cc894f1cd8d2',
    'queso'
  );
INSERT INTO products (id, category_id)
VALUES
  -- hazelnut
  (
    'babea3b0-8a91-4ba7-a565-2370baf1ab32',
    'bba08683-d337-4033-a087-bdb5366ca64c'
  ),
  -- almond
  (
    'dbf6b240-2dd7-48cf-b702-18f0d9c85c62',
    'bba08683-d337-4033-a087-bdb5366ca64c'
  ),
  -- halloumi
  (
    'f0b26402-979f-4cb0-a25a-7147c1dd62d6',
    '4f5061bc-a388-4f1a-918f-cc894f1cd8d2'
  );
INSERT INTO products_translations (id, locale_id, product_id, value)
VALUES
  (
    'c914f8c9-463c-4148-9d5b-68df2b432797',
    '469cfec0-b8f4-4694-b786-055fc4c449ec',
    'babea3b0-8a91-4ba7-a565-2370baf1ab32',
    'hazelnut'
  ),
  (
    '61738394-c99d-4104-9039-45dfff5c0740',
    '58728ba5-bc77-473f-9475-4a7d658897c2',
    'babea3b0-8a91-4ba7-a565-2370baf1ab32',
    'Haselnuss'
  ),
  (
    '5b1a50e4-0385-4b47-8f66-932fa45e49b0',
    'dd25b32b-50c5-4921-893a-d84aae90c38b',
    'babea3b0-8a91-4ba7-a565-2370baf1ab32',
    'avellana'
  ),
  (
    'cbd28565-710f-42ac-b8e1-97e104be2fba',
    '469cfec0-b8f4-4694-b786-055fc4c449ec',
    'dbf6b240-2dd7-48cf-b702-18f0d9c85c62',
    'almond'
  ),
  (
    '4ca98a0b-d445-439a-85dd-3052143b00b1',
    '58728ba5-bc77-473f-9475-4a7d658897c2',
    'dbf6b240-2dd7-48cf-b702-18f0d9c85c62',
    'Mandel'
  ),
  (
    '9ad50d85-7c81-4d8b-a4af-89725eb27f47',
    'dd25b32b-50c5-4921-893a-d84aae90c38b',
    'dbf6b240-2dd7-48cf-b702-18f0d9c85c62',
    'almendra'
  ),
  (
    '188988fd-8395-45aa-b1f5-311c92a0b9ba',
    '469cfec0-b8f4-4694-b786-055fc4c449ec',
    'f0b26402-979f-4cb0-a25a-7147c1dd62d6',
    'halloumi'
  ),
  (
    '8134d4ca-3a41-462f-8ca5-1ab7a96edca9',
    '58728ba5-bc77-473f-9475-4a7d658897c2',
    'f0b26402-979f-4cb0-a25a-7147c1dd62d6',
    'Halloumi'
  ),
  (
    'fd0000fc-fa52-4b8e-a8a5-749e7741c807',
    'dd25b32b-50c5-4921-893a-d84aae90c38b',
    'f0b26402-979f-4cb0-a25a-7147c1dd62d6',
    'halloumi'
  );
INSERT INTO shops_products (id, shop_id, product_id)
VALUES
  -- Original Unverpackt
  (
    '28fd3173-5279-4cd9-841d-f284e76eecf0',
    '7c3153b9-edb0-4ae4-a068-8c9d06532ea5',
    'babea3b0-8a91-4ba7-a565-2370baf1ab32'
  ),
  (
    'fc3901fe-07a8-4e9b-9095-766c8b1806f3',
    '7c3153b9-edb0-4ae4-a068-8c9d06532ea5',
    'dbf6b240-2dd7-48cf-b702-18f0d9c85c62'
  ),
  -- YES FUTURE
  (
    '5e941ec7-9562-4a64-897a-82be02e69fcf',
    'd7106377-c6e3-4d4e-bf49-b077e719fbe7',
    'dbf6b240-2dd7-48cf-b702-18f0d9c85c62'
  );
