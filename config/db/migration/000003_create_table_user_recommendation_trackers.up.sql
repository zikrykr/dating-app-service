BEGIN;

CREATE TABLE IF NOT EXISTS user_recommendation_tracker(
  id SERIAL PRIMARY KEY NOT NULL,
  user_id VARCHAR(100) NOT NULL,
  seen_user_id VARCHAR(100) NOT NULL,
  tracker_date DATE NOT NULL,
  CONSTRAINT unique_user_tracker_date UNIQUE (user_id, seen_user_id, tracker_date)
);

ALTER TABLE user_recommendation_tracker ADD CONSTRAINT user_recommendation_tracker_fk_users FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE NO ACTION ON UPDATE NO ACTION DEFERRABLE INITIALLY DEFERRED;
ALTER TABLE user_recommendation_tracker ADD CONSTRAINT user_recommendation_tracker_fk_seen_user_id FOREIGN KEY (seen_user_id) REFERENCES users(id) ON DELETE NO ACTION ON UPDATE NO ACTION DEFERRABLE INITIALLY DEFERRED;

COMMIT;