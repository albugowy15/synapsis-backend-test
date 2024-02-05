CREATE OR REPLACE FUNCTION trigger_set_timestamp()
RETURNS TRIGGER AS $$
BEGIN
  NEW.updated_at = NOW();
  RETURN NEW;
END;
$$ LANGUAGE plpgsql;

CREATE TABLE "users" (
  "id" bigint PRIMARY KEY,
  "username" varchar(30) UNIQUE,
  "email" varchar(50) UNIQUE,
  "fullname" varchar(100),
  "password" varchar(16),
  "address" varchar(200),
  "phone_number" varchar(14) UNIQUE,
  "created_at" TIMESTAMPTZ DEFAULT NOW() NOT NULL,
  "updated_at" TIMESTAMPTZ DEFAULT NOW() NOT NULL
);
CREATE TRIGGER set_timestamp
BEFORE UPDATE ON "users"
FOR EACH ROW
EXECUTE PROCEDURE trigger_set_timestamp();

CREATE TABLE "shopping_carts" (
  "id" bigint PRIMARY KEY,
  "user_id" bigint NOT NULL,
  "created_at" TIMESTAMPTZ DEFAULT NOW() NOT NULL,
  "updated_at" TIMESTAMPTZ DEFAULT NOW() NOT NULL
);
CREATE TRIGGER set_timestamp
BEFORE UPDATE ON "shopping_carts"
FOR EACH ROW
EXECUTE PROCEDURE trigger_set_timestamp();

CREATE TABLE "items" (
  "id" bigint PRIMARY KEY,
  "quantity" smallint,
  "shopping_cart_id" bigint,
  "product_id" bigint,
  "created_at" TIMESTAMPTZ DEFAULT NOW() NOT NULL,
  "updated_at" TIMESTAMPTZ DEFAULT NOW() NOT NULL
);
CREATE TRIGGER set_timestamp
BEFORE UPDATE ON "items"
FOR EACH ROW
EXECUTE PROCEDURE trigger_set_timestamp();

CREATE TABLE "products" (
  "id" bigint PRIMARY KEY,
  "name" varchar(100) NOT NULL,
  "description" text,
  "price" numeric(10, 2) NOT NULL,
  "stock" int NOT NULL,
  "created_at" TIMESTAMPTZ DEFAULT NOW() NOT NULL,
  "updated_at" TIMESTAMPTZ DEFAULT NOW() NOT NULL
);
CREATE TRIGGER set_timestamp
BEFORE UPDATE ON "products"
FOR EACH ROW
EXECUTE PROCEDURE trigger_set_timestamp();

CREATE TABLE "categories" (
  "id" bigint PRIMARY KEY,
  "name" varchar(20) NOT NULL,
  "description" varchar(150),
  "created_at" TIMESTAMPTZ DEFAULT NOW() NOT NULL,
  "updated_at" TIMESTAMPTZ DEFAULT NOW() NOT NULL
);
CREATE TRIGGER set_timestamp
BEFORE UPDATE ON "categories"
FOR EACH ROW
EXECUTE PROCEDURE trigger_set_timestamp();

CREATE TABLE "products_categories" (
  "id" bigint PRIMARY KEY,
  "product_id" bigint,
  "category_id" bigint,
  "created_at" TIMESTAMPTZ DEFAULT NOW() NOT NULL,
  "updated_at" TIMESTAMPTZ DEFAULT NOW() NOT NULL
);
CREATE TRIGGER set_timestamp
BEFORE UPDATE ON "products_categories"
FOR EACH ROW
EXECUTE PROCEDURE trigger_set_timestamp();

CREATE TABLE "payment_methods" (
  "id" int PRIMARY KEY,
  "type" char(20),
  "merchant" varchar(30),
  "fee" numeric(10, 2),
  "created_at" TIMESTAMPTZ DEFAULT NOW() NOT NULL,
  "updated_at" TIMESTAMPTZ DEFAULT NOW() NOT NULL
);
CREATE TRIGGER set_timestamp
BEFORE UPDATE ON "payment_methods"
FOR EACH ROW
EXECUTE PROCEDURE trigger_set_timestamp();

CREATE TABLE "payment_transactions" (
  "id" bigint PRIMARY KEY,
  "total_product_price" numeric(10, 2),
  "tax" numeric(10, 2),
  "total_price" numeric(10, 2),
  "payment_method_id" int,
  "status" char(10),
  "qr_code" varchar(100),
  "virtual_account" varchar(30),
  "account_number" varchar(30),
  "user_id" bigint,
  "created_at" TIMESTAMPTZ DEFAULT NOW() NOT NULL,
  "updated_at" TIMESTAMPTZ DEFAULT NOW() NOT NULL
);
CREATE TRIGGER set_timestamp
BEFORE UPDATE ON "payment_transactions"
FOR EACH ROW
EXECUTE PROCEDURE trigger_set_timestamp();

ALTER TABLE "shopping_carts" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id");

ALTER TABLE "items" ADD FOREIGN KEY ("shopping_cart_id") REFERENCES "shopping_carts" ("id");

ALTER TABLE "items" ADD FOREIGN KEY ("product_id") REFERENCES "products" ("id");

ALTER TABLE "products_categories" ADD FOREIGN KEY ("product_id") REFERENCES "products" ("id");

ALTER TABLE "products_categories" ADD FOREIGN KEY ("category_id") REFERENCES "categories" ("id");

ALTER TABLE "payment_transactions" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id");

ALTER TABLE "payment_transactions" ADD FOREIGN KEY ("payment_method_id") REFERENCES "payment_methods" ("id");
