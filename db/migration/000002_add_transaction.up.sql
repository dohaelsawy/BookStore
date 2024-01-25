
CREATE TABLE "transaction_payment"(
    "transaction_id" SERIAL PRIMARY KEY,
    "amount" INT NOT NULL,
    "line_id" INT NOT NULL,
    "shipping_method_id" INT NOT NULL,
    "card_number" INT NOT NULL,
    "expiration_date" DATE NOT NULL,
    "cvv" VARCHAR(4) NOT NULL,
    "status" VARCHAR(255) NOT NULL DEFAULT 'pending',
    "customer_id" INT NOT NULL
);

ALTER TABLE "transaction_payment" ADD FOREIGN KEY ("line_id") REFERENCES "order_line"("line_id");
ALTER TABLE "transaction_payment" ADD FOREIGN KEY("shipping_method_id") REFERENCES "shipping_method"("shipping_method_id");
ALTER TABLE "transaction_payment" ADD FOREIGN KEY("customer_id") REFERENCES "customer" ( "customer_id");
