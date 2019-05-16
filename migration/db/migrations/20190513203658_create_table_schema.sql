
-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied

DROP TABLE IF EXISTS users CASCADE;
DROP TABLE IF EXISTS attendence CASCADE;
DROP TABLE IF EXISTS goods CASCADE;
DROP TABLE IF EXISTS goods_purchase CASCADE;
DROP TABLE IF EXISTS bill CASCADE;
DROP TABLE IF EXISTS participent CASCADE;
DROP TABLE IF EXISTS menus CASCADE;

CREATE TABLE users(
    users_id SERIAL PRIMARY KEY,
    users_name VARCHAR(50) NOT NULL,
    email VARCHAR(50),
    session_key VARCHAR(64),
    password_digest VARCHAR(250),
    users_type INTEGER NOT NULL,
    UNIQUE(users_name)
);

CREATE TABLE attendence(
    att_date DATE,
    users_id INTEGER REFERENCES users(users_id),
    morning_att BOOLEAN,
    noon_att BOOLEAN,
    dinner_att BOOLEAN,
    PRIMARY KEY(att_date,users_id)
);

CREATE TABLE goods
(
    goods_id SERIAL PRIMARY KEY, 
    goods_name VARCHAR(50) NOT NULL,
    created_at DATE,
    updated_at DATE,
    deleted_at DATE
);

CREATE TABLE bill(
    bill_id SERIAL PRIMARY KEY,
    amount float,
    title VARCHAR(100),
    description VARCHAR(500),
    bill_date DATE,
    author_id INTEGER REFERENCES users(users_id),
    created_at DATE,
    updated_at DATE,
    deleted_at DATE
);

CREATE TABLE goods_purchase(
    goods_id INTEGER REFERENCES goods(goods_id),
    bill_id INTEGER REFERENCES bill(bill_id),
    purchase_date DATE,
    price FLOAT,
    quantity FLOAT,
    is_purchased BOOLEAN,
    PRIMARY KEY(goods_id,bill_id)
);

CREATE TABLE participent(
    users_id INTEGER REFERENCES users(users_id),
    bill_id INTEGER REFERENCES bill(bill_id),
    amount FLOAT,
    is_payer BOOLEAN,
    PRIMARY KEY(users_id,bill_id)
);

CREATE TABLE menus(
    menu_id SERIAL PRIMARY KEY,
    name VARCHAR(50),
    week INTEGER
);

-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back

DROP TABLE users CASCADE;
DROP TABLE attendence CASCADE;
DROP TABLE goods CASCADE;
DROP TABLE goods_purchase CASCADE;
DROP TABLE bill CASCADE;
DROP TABLE payer CASCADE;
DROP TABLE reciepient CASCADE;
DROP TABLE menus CASCADE;