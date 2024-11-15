BEGIN;

CREATE TABLE balance (
   id BIGSERIAL PRIMARY KEY,
   user_id uuid REFERENCES users (id),
   current DOUBLE PRECISION,
   withdrawn DOUBLE PRECISION
);


CREATE UNIQUE INDEX user_id_balance_idx ON balance(user_id);

COMMIT;
