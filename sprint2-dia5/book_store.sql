DROP DATABASE IF EXISTS book_store;

CREATE DATABASE book_store;

USE book_store;

CREATE TABLE authors (
	id INT NOT NULL AUTO_INCREMENT PRIMARY KEY,
    author_name VARCHAR(255) NOT NULL
);

INSERT INTO authors (author_name)
VALUES ("J. R. R. Tolkien"),
("George R. R. Martin");

CREATE TABLE books (
	id INT NOT NULL AUTO_INCREMENT PRIMARY KEY,
	book_name VARCHAR(255) NOT NULL,
    author_id INT NOT NULL,
    FOREIGN KEY (author_id) REFERENCES authors(id)
);

INSERT INTO books (book_name, author_id)
VALUES ("O Senhor dos Anéis - A Sociedade do Anel", 1),
("As Crônicas de Gelo e Fogo", 2);

CREATE TABLE customers (
	id INT NOT NULL AUTO_INCREMENT PRIMARY KEY,
    customer_name VARCHAR(255) NOT NULL
);

INSERT INTO customers (customer_name)
VALUES ("Maurício");

CREATE TABLE rental (
	book_id INT NOT NULL,
    customer_id INT NOT NULL,
    rental_date DATETIME DEFAULT NOW(),
    return_date DATETIME DEFAULT null,
    FOREIGN KEY (book_id) REFERENCES books(id),
    FOREIGN KEY (customer_id) REFERENCES customers(id),
    PRIMARY KEY (book_id, customer_id)
);

INSERT INTO rental (book_id, customer_id)
VALUES (2, 1);

SELECT b.book_name AS book, a.author_name AS author,
	c.customer_name AS customer, r.rental_date, r.return_date FROM books AS b
INNER JOIN authors AS a ON a.id = b.author_id
INNER JOIN rental AS r ON r.book_id = b.id
INNER JOIN customers AS c on r.customer_id = c.id;
