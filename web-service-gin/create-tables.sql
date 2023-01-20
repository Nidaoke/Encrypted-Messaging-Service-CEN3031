DROP TABLE IF EXISTS accounts;
CREATE TABLE accounts (
  username      VARCHAR(32) NOT NULL,
  password     VARCHAR(32) NOT NULL,
  PRIMARY KEY (`username`)
);

INSERT INTO accounts
  (username, password)
VALUES
  ('user1', 'password123'),
  ('user2', 'password321');