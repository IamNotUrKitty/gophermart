BEGIN TRANSACTION;
  CREATE TABLE IF NOT EXISTS users (
		"id" uuid PRIMARY KEY DEFAULT gen_random_uuid(),
		"username" VARCHAR(250) NOT NULL UNIQUE,
		"password" VARCHAR(250)
		);
COMMIT;
