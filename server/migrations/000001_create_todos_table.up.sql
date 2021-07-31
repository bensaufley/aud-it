CREATE TABLE todo_statuses (
  id INTEGER PRIMARY KEY,
  val TEXT NOT NULL UNIQUE
);

INSERT INTO todo_statuses (val)
VALUES
  ('Unstarted'),
  ('In Progress'),
  ('Abandoned'),
  ('Complete'),
  ('Deleted')
;

CREATE TABLE todos (
  ulid TEXT PRIMARY KEY,
  contents TEXT NOT NULL,
  todo_status_id INTEGER NOT NULL,
  FOREIGN KEY (todo_status_id) REFERENCES todo_statuses(id)
    ON DELETE RESTRICT
);
