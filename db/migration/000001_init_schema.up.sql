CREATE TABLE "customer"(
    "customer_id" SERIAL PRIMARY KEY,
    "first_name" VARCHAR(255) NOT NULL,
    "last_name" VARCHAR(255) NOT NULL,
    "email" VARCHAR(255) NOT NULL,
    "password" VARCHAR(255) NOT NULL,
    "city" VARCHAR(255) NOT NULL,
    "phone_number" VARCHAR(255) NOT NULL
);
CREATE TABLE "book"(
    "book_id" SERIAL PRIMARY KEY,
    "name" VARCHAR(255) NOT NULL,
    "publish_date" VARCHAR(255) NOT NULL,
    "price" INT NOT NULL,
    "sku" VARCHAR(255) NOT NULL,
    "description" VARCHAR(255) NOT NULL,
    "created_at" timestamptz NOT NULL DEFAULT (now()),
    "updated_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "shipping_method"(
    "shipping_method_id" SERIAL PRIMARY KEY,
    "name" VARCHAR(255) NOT NULL,
    "cost" INT NOT NULL
);

CREATE TABLE "customer_order"(
    "customer_order_id" SERIAL PRIMARY KEY,
    "order_date" VARCHAR(255) NOT NULL,
    "customer_id" INT NOT NULL ,
    "shipping_method_id" INT NOT NULL 
);

CREATE TABLE "order_line"(
    "line_id" SERIAL PRIMARY KEY,
    "order_id" INT NOT NULL ,
    "book_id" INT NOT NULL ,
    "price" INT NOT NULL
);

ALTER TABLE "order_line" ADD FOREIGN KEY ("order_id") REFERENCES "customer_order" ("customer_order_id");
ALTER TABLE "order_line" ADD FOREIGN KEY ("book_id") REFERENCES "book" ("book_id");
ALTER TABLE "customer_order" ADD FOREIGN KEY ("customer_id") REFERENCES "customer" ("customer_id");
