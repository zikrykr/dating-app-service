BEGIN;

CREATE TYPE swipe_type_enum AS ENUM('like', 'pass');

CREATE TABLE IF NOT EXISTS user_swipes(
  id SERIAL PRIMARY KEY NOT NULL,
  user_id VARCHAR(100) NOT NULL,
  swiped_user_id VARCHAR(100) NOT NULL,
  swipe_type swipe_type_enum,
  created_at TIMESTAMP,
  updated_at TIMESTAMP,
  deleted_at TIMESTAMP,
  CONSTRAINT unique_user_swiped_id UNIQUE (user_id, swiped_user_id)
);

ALTER TABLE user_swipes ADD CONSTRAINT user_swipes_fk_users FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE NO ACTION ON UPDATE NO ACTION DEFERRABLE INITIALLY DEFERRED;
ALTER TABLE user_swipes ADD CONSTRAINT user_swipes_fk_swiped_user_id FOREIGN KEY (swiped_user_id) REFERENCES users(id) ON DELETE NO ACTION ON UPDATE NO ACTION DEFERRABLE INITIALLY DEFERRED;

COMMIT;