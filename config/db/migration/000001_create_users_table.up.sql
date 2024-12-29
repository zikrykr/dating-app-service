BEGIN;

CREATE TABLE IF NOT EXISTS users(
  id VARCHAR(100) PRIMARY KEY NOT NULL,
  name VARCHAR(100) NOT NULL,
  email VARCHAR(100) NOT NULL,
  date_of_birth TIMESTAMP,
  gender VARCHAR(50),
  profile_image_url VARCHAR(200),
  description TEXT,
  password TEXT NOT NULL,
  created_at TIMESTAMP,
  updated_at TIMESTAMP,
  deleted_at TIMESTAMP
);

COMMIT;