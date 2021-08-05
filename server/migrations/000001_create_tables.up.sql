CREATE TABLE import_statuses (
  ulid INTEGER PRIMARY KEY,
  val TEXT NOT NULL UNIQUE
);

INSERT INTO import_statuses (val)
VALUES
  ('Pending'),
  ('Complete')
;

CREATE TABLE books (
  ulid TEXT PRIMARY KEY,
  title TEXT NOT NULL,
  duration INTEGER NOT NULL,
  import_status_ulid INTEGER NOT NULL,
  FOREIGN KEY (import_status_ulid) REFERENCES import_statuses(ulid)
    ON DELETE RESTRICT
);

CREATE TABLE authors (
  ulid TEXT PRIMARY KEY,
  first_name TEXT,
  last_name TEXT NOT NULL
);

CREATE TABLE book_authors (
  ulid TEXT PRIMARY KEY,
  book_ulid TEXT NOT NULL,
  author_ulid TEXT NOT NULL,
  meta TEXT,
  FOREIGN KEY (book_ulid) REFERENCES books(ulid)
    ON DELETE CASCADE,
  FOREIGN KEY (author_ulid) REFERENCES authors(ulid)
    ON DELETE CASCADE
);

CREATE UNIQUE INDEX distinct_book_authors
ON book_authors (book_ulid, author_ulid);

CREATE TABLE genres (
  ulid TEXT PRIMARY KEY,
  name TEXT NOT NULL UNIQUE
);

CREATE TABLE book_genres (
  ulid TEXT PRIMARY KEY,
  book_ulid TEXT NOT NULL,
  genre_ulid TEXT NOT NULL,
  FOREIGN KEY (book_ulid) REFERENCES books(ulid)
    ON DELETE CASCADE,
  FOREIGN KEY (genre_ulid) REFERENCES genres(ulid)
    ON DELETE CASCADE
);

CREATE UNIQUE INDEX distinct_book_genres
ON book_genres (book_ulid, genre_ulid);

CREATE TABLE users (
  ulid TEXT PRIMARY KEY,
  username TEXT NOT NULL UNIQUE
);

CREATE TABLE lists (
  ulid TEXT PRIMARY KEY,
  name TEXT NOT NULL,
  user_ulid TEXT NOT NULL,
  FOREIGN KEY (user_ulid) REFERENCES users(ulid)
    ON DELETE CASCADE
);

CREATE TABLE list_books (
  ulid TEXT PRIMARY KEY,
  list_ulid TEXT NOT NULL,
  book_ulid TEXT NOT NULL,
  FOREIGN KEY (list_ulid) REFERENCES lists(ulid)
    ON DELETE CASCADE,
  FOREIGN KEY (book_ulid) REFERENCES books(ulid)
    ON DELETE CASCADE
);
