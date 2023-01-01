DROP TABLE IF EXISTS transactions CASCADE;

CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE transactions
(
    id         UUID PRIMARY KEY                  DEFAULT uuid_generate_v4(),
    amount     DOUBLE PRECISION         NOT NULL,
    datetime   TIMESTAMP WITH TIME ZONE NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP WITH TIME ZONE          DEFAULT CURRENT_TIMESTAMP
);

-- mock the first 1000 BTC
INSERT INTO transactions (amount, datetime)
VALUES (1000, '2019-10-05T13:00:00+00:00')