DROP TABLE IF EXISTS accounts;
CREATE TABLE accounts (
  username      VARCHAR(64) NOT NULL,
  password     VARCHAR(64) NOT NULL,
  PRIMARY KEY (`username`)
)
