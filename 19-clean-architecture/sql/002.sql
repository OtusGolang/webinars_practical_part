CREATE INDEX owner_start_time_idx ON events USING btree (owner, start_time);
