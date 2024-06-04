-- +goose Up
-- +goose StatementBegin

CREATE EXTENSION IF NOT EXISTS "uuid-ossp";
CREATE TABLE IF NOT EXISTS users
(
    id         SERIAL PRIMARY KEY,
    uuid       UUID         NOT NULL UNIQUE DEFAULT uuid_generate_v4(),
    phone      VARCHAR(255) NOT NULL UNIQUE,
    balance    NUMERIC                      DEFAULT 0,
    created_at TIMESTAMP    NOT NULL        DEFAULT NOW()
);

-- transactions table to store transaction history
CREATE TABLE IF NOT EXISTS transactions
(
    transaction_id SERIAL PRIMARY KEY,
    user_id        INTEGER REFERENCES users (id),
    amount         NUMERIC,
    description    TEXT,
    timestamp      TIMESTAMP DEFAULT NOW()
);

-- events table to store gift code information
CREATE TABLE IF NOT EXISTS events
(
    code        UUID         NOT NULL UNIQUE DEFAULT uuid_generate_v4(),
    title       VARCHAR(250) NOT NULL,
    description VARCHAR(500),
    gift_amount NUMERIC,
    max_users   INTEGER,
    user_count  INTEGER                      DEFAULT 0,
    start_time  TIMESTAMP    NOT NULL,
    end_time    TIMESTAMP    NOT NULL,
    published   BOOLEAN      NOT NULL        DEFAULT false,
    created_at  TIMESTAMP                    DEFAULT NOW()
);

-- user_events table to track gift code usage by users
CREATE TABLE IF NOT EXISTS user_events
(
    user_id    INTEGER REFERENCES users (id),
    event_code UUID REFERENCES events (code),
    timestamp  TIMESTAMP DEFAULT NOW(),
    PRIMARY KEY (user_id, event_code)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS user_events;
DROP TABLE IF EXISTS transactions;
DROP TABLE IF EXISTS events;
DROP TABLE IF EXISTS users;
-- +goose StatementEnd
