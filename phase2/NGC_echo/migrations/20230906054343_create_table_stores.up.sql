CREATE TABLE stores (
   ID SERIAL PRIMARY KEY,
   Name VARCHAR(255),
   Address VARCHAR(255)
);

CREATE TABLE store_details (
   ID SERIAL PRIMARY KEY,
   store_id INT,
   weather VARCHAR(255),
   latitude FLOAT,
   longitude FLOAT,
   total_sales FLOAT,
   rating FLOAT,
   FOREIGN KEY (store_id) REFERENCES stores(ID)
);
