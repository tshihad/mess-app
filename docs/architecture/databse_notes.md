# Table Descritpion
## users
#### Fields
- user_id INTEGER PRIMARY KEY (index)
- username STRING
- email STRING
- session_key STRING
- password_digest STRING
- user_type INTEGER

## attendence
#### Fields
- date DATE (index)
- user_id FORIEGN KEY
- morining_att BOOLEAN
- noon_att BOOLEAN
- dinner_att BOOLEAN

## goods
#### Fields
- goods_id INTEGER PRIMARY KEY (index)
- name STRING
- created_at DATE
- updated_at DATE
- deleted_at DATE

## goods_purchase
#### Fields
- goods_purchase_id INTEGER PRIMARY KEY
- goods_id INTEGER FORIEGN KEY
- bill_id INTEGER FORIEGN KEY (index)
- date DATE (index)
- price FLOAT
- quantity FLOAT
- is_purchased BOOL

## bill
#### Fields
- bill_id INTEGER PRIMARY KEY (index)
- amount float
- title STRING
- description STRING
- date DATE (index)
- author_id INTEGER FORIEGN KEY(users)
- created_at DATE
- updated_at DATE
- deleted_at DATE

## payer
#### Fields
- payer_id INTEGER PRIMARY KEY
- bill_id INTEGER FORIEGN KEY
- amount FLOAT

## reciepient
#### Fields
- recipient_id INTEGER PRIMARY KEY
- bill_id INTEGER FORIEGN KEY
- amount FLOAT

## menus
#### Fields
- menu_id INTEGER PRIMARY KEY
- name STRING
- week INTEGER