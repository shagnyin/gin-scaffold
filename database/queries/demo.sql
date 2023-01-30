-- name: CreateDemo :exec
INSERT into `demo` (id, email, name, avatar)
values (?, ?, ?, ?);

-- name: DeleteDemoByID :exec
delete
from demo
where id = ?
limit 1;

-- name: SoftDeleteDemoById :exec
update demo
set deleted_at=?
where id = ?;

-- name: UpdateDemoByID :exec
UPDATE demo
set email=?,
    avatar=?
where id = ?;

-- name: CheckDemoById :one
select count(1)
from demo
where 1
  and id != ?
  and email = ?;

-- name: GetDemoByID :one
SELECT *
FROM `demo`
WHERE 1
  and id = ?
  and deleted_at = ''
limit 1;

-- name: GetDemoByEmail :one
SELECT *
FROM `demo`
WHERE 1
  and email = ?
  and deleted_at = ''
limit 1;

-- name: CountDemoByPage :one
SELECT count(1)
from demo
where 1
  and deleted_at = ''
limit 1;

-- name: ListDemoByPage :many
SELECT *
from demo
where 1
  and deleted_at = ''
order by created_at desc
limit ? offset ?;