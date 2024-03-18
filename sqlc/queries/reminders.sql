
-- name: GetReminder :one
SELECT * FROM reminders
WHERE id = ? LIMIT 1;

-- name: ListReminders :many
SELECT * FROM reminders
ORDER BY id;

-- name: CreateReminder :exec
INSERT INTO reminders (title, `description`, remind_at)
VALUES (?, ?, ?);

-- name: UpdateReminder :exec
UPDATE reminders
SET title = ?, `description` = ?, remind_at = ?
WHERE id = ?;

-- name: DeleteReminder :exec
DELETE FROM reminders
WHERE id = ?;
