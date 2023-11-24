CREATE DATABASE Database1;
USE Database1;

CREATE TABLE user(
	user_id int AUTO_INCREMENT PRIMARY KEY,
    name_user  VARCHAR(255) NOT NULL,
    rol varchar(50) NOT NULL,
    mail varchar(150) NOT NULL, 
    pwd varchar(250) NOT NULL,
    constraint chk_rol CHECK (rol="customer" or rol="store administrator")
);

CREATE TABLE store (
    store_id INT AUTO_INCREMENT PRIMARY KEY,
    name_store VARCHAR(255) NOT NULL,
    address VARCHAR(255),
    user_id int, 
    foreign key (user_id) references user(user_id)
);

CREATE TABLE reviews(
	rating_id int AUTO_INCREMENT PRIMARY KEY,
    store_id int NOT NULL,
    customer_id int NOT NULL,
    rating int NOT NULL CHECK (rating >= 1 AND rating <= 5),
    comment TEXT
);

CREATE TABLE store_article_price (
    id_price INT AUTO_INCREMENT PRIMARY KEY,
    store_id INT,
    article_id INT,
    price DECIMAL(10, 2),
    content int,
    unitMeasurement varchar(50),
    FOREIGN KEY (store_id) REFERENCES store(store_id)
);
CREATE DATABASE Database2;
USE Database2;
CREATE TABLE article (
    article_id INT AUTO_INCREMENT PRIMARY KEY,
    name_article VARCHAR(255) NOT NULL,
    brand varchar (250),
    category varchar(250)
);

 