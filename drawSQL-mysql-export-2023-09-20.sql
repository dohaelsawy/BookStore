CREATE TABLE `customer_order`(
    `customer_order_id` BIGINT UNSIGNED NOT NULL AUTO_INCREMENT PRIMARY KEY,
    `order_date` VARCHAR(255) NOT NULL,
    `customer_id` BIGINT NOT NULL,
    `shipping_method_id` BIGINT NOT NULL
);
CREATE TABLE `book`(
    `book_id` INT UNSIGNED NOT NULL AUTO_INCREMENT PRIMARY KEY,
    `name` VARCHAR(255) NOT NULL,
    `publish_dat` VARCHAR(255) NOT NULL,
    `price` INT NOT NULL,
    `sku` VARCHAR(255) NOT NULL,
    `description` VARCHAR(255) NOT NULL,
    `created_on` VARCHAR(255) NOT NULL,
    `updated_on` VARCHAR(255) NOT NULL
);
CREATE TABLE `shipping_method`(
    `shipping_method_id` INT UNSIGNED NOT NULL AUTO_INCREMENT PRIMARY KEY,
    `name` VARCHAR(255) NOT NULL,
    `cost` INT NOT NULL
);
CREATE TABLE `order_line`(
    `line_id` INT UNSIGNED NOT NULL AUTO_INCREMENT PRIMARY KEY,
    `order_id` BIGINT NOT NULL,
    `book_id` BIGINT NOT NULL,
    `price` BIGINT NOT NULL
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
    `customer_order` ADD CONSTRAINT `customer_order_customer_id_foreign` FOREIGN KEY(`customer_id`) REFERENCES `customer`(`customer_id`);
ALTER TABLE
    `customer_order` ADD CONSTRAINT `customer_order_shipping_method_id_foreign` FOREIGN KEY(`shipping_method_id`) REFERENCES `shipping_method`(`shipping_method_id`);
ALTER TABLE
    `order_line` ADD CONSTRAINT `order_line_order_id_foreign` FOREIGN KEY(`order_id`) REFERENCES `customer_order`(`customer_order_id`);