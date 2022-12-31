DROP TABLE IF EXISTS transaction CASCADE;

CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE transaction
(
    id         UUID PRIMARY KEY                  DEFAULT uuid_generate_v4(),
    amount     DOUBLE PRECISION         NOT NULL,
    datetime   TIMESTAMP WITH TIME ZONE NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP WITH TIME ZONE          DEFAULT CURRENT_TIMESTAMP
);
