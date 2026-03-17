CREATE TABLE IF NOT EXISTS events (
  id          INTEGER PRIMARY KEY,
  title       VARCHAR(50)  NOT NULL,
  description VARCHAR(500),
  started_at  DATETIME     NOT NULL DEFAULT (unixepoch()),
  ended_at    DATETIME,
  category    VARCHAR(50),
  project     VARCHAR(50),
  origin      VARCHAR(50)  NOT NULL DEFAULT 'manual'
);
