# Table Descritpion
## users
#### Fields
- users_id INTEGER PRIMARY KEY
- users_name VARCHAR(50) NOT NULL
- email VARCHAR(50)
- session_key VARCHAR(64)
- password_digest VARCHAR(250)
- users_type INTEGER NOT NULL

## attendence
#### Fields
- att_date DATE
- users_id INTEGER REFERENCES users(users_id)
- morning_att BOOLEAN
- noon_att BOOLEAN
- dinner_att BOOLEAN
- PRIMARY KEY(att_date,users_id)

## goods
#### Fields
- goods_id INTEGER PRIMARY KEY
- goods_name VARCHAR(50) NOT NULL
- created_at DATE
- updated_at DATE
- deleted_at DATE

## goods_purchase
#### Fields
- goods_purchase_id INTEGER PRIMARY KEY
- goods_id INTEGER REFERENCES goods(goods_id)
- bill_id INTEGER REFERENCES bill(bill_id)
- purchase_date DATE
- price FLOAT
- quantity FLOAT
- is_purchased BOOLEAN

## bill
#### Fields
- bill_id INTEGER PRIMARY KEY
- amount float
- title VARCHAR(100)
- description VARCHAR(500)
- bill_date DATE
- author_id INTEGER REFERENCES users(users_id)
- created_at DATE
- updated_at DATE
- deleted_at DATE

## participent 
#### Fields
- users_id INTEGER REFERENCES users(users_id),
- bill_id INTEGER REFERENCES bill(bill_id),
- amount FLOAT,
- is_payer BOOLEAN,
- PRIMARY KEY(users_id,bill_id)

## menus
#### Fields
- menu_id INTEGER PRIMARY KEY,
- name VARCHAR(50),
- week INTEGER
