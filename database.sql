CREATE TABLE category (
    id SERIAL PRIMARY KEY,
    category_name VARCHAR(100)
);
CREATE TABLE location (
    id SERIAL PRIMARY KEY,
    address VARCHAR(100),
    city VARCHAR(100),
    province VARCHAR(100),
    item_location VARCHAR(100)
);


CREATE TABLE item (
    id SERIAL PRIMARY KEY,
    item_code VARCHAR(50) NOT NULL UNIQUE,
    name VARCHAR(100) NOT NULL,
    category_id INT,
    location_id INT,
    price INT NOT NULL,
    stock INT NOT NULL DEFAULT 0,
    FOREIGN KEY (category_id) REFERENCES category(id),
    FOREIGN KEY (location_id) REFERENCES lcoation(id)
);

INSERT INTO category(category_name) VALUES 
('Elektronik'),
('Pakaian');

INSERT INTO location(address, city, province, item_position) VALUES
('jln. Ryacudu 340', 'Palembang', 'South Sumatra', 'Sriwijaya Store'),
('jln. Ryacudu 026', 'Palembang', 'South Sumatra', 'Sriwijaya Werehouse, Shelf A number 1');

INSERT INTO item(item_code, name, category_id, location_id, price, stock) VALUES
('P001', 'Redmi 3A', 1, 1, 1000000, 20),
('P002', 'Jaket Tactical', 2, 2, 200000, 20);