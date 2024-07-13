CREATE TABLE scheduler (
  id INTEGER PRIMARY KEY AUTOINCREMENT,
  date DATE NOT NULL,
  title VARCHAR(256) NOT NULL,
  comment TEXT,
  repeat VARCHAR(128)
);

CREATE INDEX date_index ON scheduler (date);