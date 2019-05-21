
-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied
DROP TABLE IF EXISTS cluster CASCADE;
DROP TABLE IF EXISTS cluster_users CASCADE;

CREATE TABLE cluster(
    cluster_id SERIAL PRIMARY KEY,
    author_id INTEGER REFERENCES users(users_id),
    title VARCHAR(50),
    description VARCHAR(500),
    created_at DATE NOT NULL DEFAULT CURRENT_DATE,
    updated_at DATE,
    deleted_at DATE
);

CREATE TABLE cluster_users(
    users_id INTEGER REFERENCES users(users_id),
    cluster_id INTEGER REFERENCES cluster(cluster_id),
    created_at DATE NOT NULL DEFAULT CURRENT_DATE,
    updated_at DATE,
    deleted_at DATE,
    PRIMARY KEY(users_id,cluster_id)
);

ALTER TABLE attendence ADD COLUMN cluster_id INTEGER REFERENCES cluster(cluster_id);
ALTER TABLE attendence DROP CONSTRAINT attendence_pkey;
ALTER TABLE attendence ADD PRIMARY KEY(att_date,users_id,cluster_id);

ALTER TABLE goods DROP COLUMN created_at,DROP COLUMN updated_at,DROP COLUMN deleted_at;
ALTER TABLE goods ADD COLUMN cluster_id INTEGER REFERENCES cluster(cluster_id);

ALTER TABLE bill DROP CONSTRAINT bill_pkey CASCADE;
ALTER TABLE bill ADD COLUMN cluster_id INTEGER REFERENCES cluster(cluster_id);
ALTER TABLE bill ADD PRIMARY KEY(bill_id,cluster_id);
-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back

DROP TABLE IF EXISTS cluster CASCADE;
DROP TABLE IF EXISTS cluster_users CASCADE;