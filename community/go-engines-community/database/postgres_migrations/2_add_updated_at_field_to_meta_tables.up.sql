ALTER TABLE entities ADD COLUMN updated_at TIMESTAMP NOT NULL DEFAULT now();
ALTER TABLE users ADD COLUMN updated_at TIMESTAMP NOT NULL DEFAULT now();
