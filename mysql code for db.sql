CREATE TABLE `customer_order`(
    `customer_order_id` INT NOT NULL AUTO_INCREMENT PRIMARY KEY,
    `order_id` INT UNSIGNED NOT NULL AUTO_INCREMENT,
    `order_date` TIMESTAMP NOT NULL,
    `customer_id` INT NOT NULL,
    `shipping_method_id` BIGINT NOT NULL
);
CREATE TABLE `book`(
    `book_id` INT UNSIGNED NOT NULL AUTO_INCREMENT PRIMARY KEY,
    `name` VARCHAR(255) NOT NULL,
    `publish_date` TIMESTAMP NOT NULL,
    `price` INT NOT NULL,
    `sku` VARCHAR(255) NOT NULL,
    `description` VARCHAR(255) NOT NULL,
    `created_on` TIMESTAMP NOT NULL,
    `updated_on` TIMESTAMP NOT NULL
);
CREATE TABLE `shipping_method`(
    `shipping_method_id` INT UNSIGNED NOT NULL AUTO_INCREMENT PRIMARY KEY,
    `name` VARCHAR(255) NOT NULL
);
CREATE TABLE `order_line`(
    `line_id` INT UNSIGNED NOT NULL AUTO_INCREMENT PRIMARY KEY,
    `order_id` INT NOT NULL,
    `book_id` INT NOT NULL,
    `price` INT NOT NULL
);
CREATE TABLE `transactionPayment`(
    `transaction_id` INT UNSIGNED NOT NULL AUTO_INCREMENT PRIMARY KEY,
    `amount` INT NOT NULL,
    `line_id` INT NOT NULL,
    `shipping_method_id` INT NOT NULL,
    `card_number` BIGINT NOT NULL,
    `expiration_date` DATE NOT NULL,
    `cvv` VARCHAR(4) NOT NULL,
    `status` VARCHAR(255) NOT NULL DEFAULT 'pending',
    `customer_id` INT NOT NULL
);
CREATE TABLE `customer`(
    `customer_id` INT UNSIGNED NOT NULL AUTO_INCREMENT PRIMARY KEY,
    `first_name` VARCHAR(255) NOT NULL,
    `last_name` VARCHAR(255) NOT NULL,
    `email` VARCHAR(255) NOT NULL,
    `password` VARCHAR(255) NOT NULL,
    `city` VARCHAR(255) NOT NULL
);
ALTER TABLE
    `order_line` ADD CONSTRAINT `order_line_book_id_foreign` FOREIGN KEY(`book_id`) REFERENCES `book`(`book_id`);
ALTER TABLE
    `transactionPayment` ADD CONSTRAINT `transactionpayment_line_id_foreign` FOREIGN KEY(`line_id`) REFERENCES `order_line`(`line_id`);
ALTER TABLE
    `transactionPayment` ADD CONSTRAINT `transactionpayment_shipping_method_id_foreign` FOREIGN KEY(`shipping_method_id`) REFERENCES `shipping_method`(`shipping_method_id`);
ALTER TABLE
    `transactionPayment` ADD CONSTRAINT `transactionpayment_customer_id_foreign` FOREIGN KEY(`customer_id`) REFERENCES `customer`(`customer_id`);
ALTER TABLE
    `customer_order` ADD CONSTRAINT `customer_order_customer_id_foreign` FOREIGN KEY(`customer_id`) REFERENCES `customer`(`customer_id`);
ALTER TABLE
    `customer_order` ADD CONSTRAINT `customer_order_order_id_foreign` FOREIGN KEY(`order_id`) REFERENCES `order_line`(`order_id`);
ALTER TABLE
    `customer_order` ADD CONSTRAINT `customer_order_shipping_method_id_foreign` FOREIGN KEY(`shipping_method_id`) REFERENCES `shipping_method`(`shipping_method_id`);