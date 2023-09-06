CREATE TABLE Transactions (
    id SERIAL PRIMARY KEY,
    user_id INT,
    product_id INT,
    store_id INT,
    quantity INT,
    total_amount FLOAT,
    FOREIGN KEY (user_id) REFERENCES users(id),
    FOREIGN KEY (product_id) REFERENCES products(id),
    FOREIGN KEY (store_id) REFERENCES stores(id)
);
