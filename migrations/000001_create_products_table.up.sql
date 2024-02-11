CREATE TABLE IF NOT EXISTS public.products (
  id uuid,
  name text DEFAULT '',
  price decimal(10,2) DEFAULT 0,
  imageurl text DEFAULT '',
  CONSTRAINT products_pk PRIMARY KEY (id)
);