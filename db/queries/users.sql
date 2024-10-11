-- name: GetUser :one
select * from users where id = $1;

-- name: GetUserByEmail :one
select * from users where email = $1;

-- name: GetUsers :many
select * from users where deleted_at is null;

-- name: CreateUser :one
insert into users (email, name, image_url) values ($1, $2, $3) returning *;

-- name: UpdateUser :one
update users set email = $1, name = $2, image_url = $3, updated_at = current_timestamp where id = $4 returning *;

-- name: DeleteUser :exec
update users set deleted_at = current_timestamp where id = $1;