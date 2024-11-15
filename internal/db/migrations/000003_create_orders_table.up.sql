BEGIN;
CREATE TABLE IF NOT EXISTS orders (
		"number" VARCHAR(255) NOT NULL PRIMARY KEY,
		"status" status DEFAULT 'NEW',
		"user_id" uuid REFERENCES users (id),
		"accrual" DOUBLE PRECISION,
		"uploaded_at" TIMESTAMP WITH TIME ZONE DEFAULT NOW()
	);
COMMIT;
