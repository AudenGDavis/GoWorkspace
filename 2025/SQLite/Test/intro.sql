-- CREATE TABLE users (
--  id INTEGER PRIMARY KEY,
-- 	name TEXT NOT NULL,
-- 	username TEXT NOT NULL UNIQUE,
-- 	email TEXT,
-- 	age INTEGER,
-- 	created_at DATETIME DEFAULT CURRENT_TIMESTAMP
-- );


-- ALTER TABLE users ADD COLUMN status TEXT;
-- DROP TABLE users;

-- INSERT INTO users (name , username)
-- VALUES ('john smith','js'), ('sal smith', 'ss'), ("cole connor", 'cc');

-- SELECT * FROM users

-- UPDATE users SET email = 'newemail@gmail.com' WHERE id=1;
-- SELECT * FROM users

DELETE FROM users;
SELECT * FROM users;