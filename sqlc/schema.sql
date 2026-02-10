CREATE TABLE IF NOT EXISTS users (
  id     INTEGER PRIMARY KEY AUTOINCREMENT,
  name   TEXT    NOT NULL,
  email  TEXT    NOT NULL UNIQUE
);

CREATE TABLE IF NOT EXISTS movies (
  id               INTEGER PRIMARY KEY AUTOINCREMENT,
  user_id          INTEGER NOT NULL,
  name             TEXT    NOT NULL DEFAULT '',
  watched_date     TEXT    NOT NULL,
  poster_url       TEXT    NOT NULL DEFAULT '',
  diary            TEXT    NOT NULL DEFAULT '',
  description      TEXT    NOT NULL DEFAULT '',
  location_watched TEXT    NOT NULL DEFAULT '',
  release_date     TEXT    NOT NULL DEFAULT '',
  FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE
);