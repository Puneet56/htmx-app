CREATE TABLE reminders (
  id SERIAL PRIMARY KEY,
  title TEXT NOT NULL,
  `description` TEXT NULL DEFAULT '',
  remind_at TIMESTAMP NOT NULL
);