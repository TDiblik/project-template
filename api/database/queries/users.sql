-- name: GetUserById :one
SELECT * FROM users
WHERE id = $1 LIMIT 1;

-- name: GetUserByEmailOrOauth :one
-- Corresponds to utils.SelectUserByEmailQuery
SELECT * FROM users
WHERE email = $1
   OR github_email = $1
   OR google_email = $1
   OR facebook_email = $1
   OR spotify_email = $1
LIMIT 1;

-- name: CheckEmailExists :one
-- Corresponds to utils.UserEmailExistsQuery
SELECT EXISTS(
    SELECT 1 FROM users
    WHERE email = $1
       OR github_email = $1
       OR google_email = $1
       OR facebook_email = $1
       OR spotify_email = $1
);

-- name: GetUserAuthByEmail :one
-- Corresponds to utils.SelectIdAndPasswordHashByEmailQuery
SELECT id, password_hash
FROM users
WHERE email = $1
   OR github_email = $1
   OR google_email = $1
   OR facebook_email = $1
   OR spotify_email = $1
LIMIT 1;

-- name: CheckHandleExists :one
SELECT EXISTS(SELECT 1 FROM users WHERE handle = $1);

-- name: CreateUser :one
INSERT INTO users (
    email, email_verified, password_hash, handle, first_name, last_name, avatar_url,
    github_id, github_email, github_handle, github_url,
    google_id, google_email,
    facebook_id, facebook_email, facebook_url,
    spotify_id, spotify_email, spotify_url
) VALUES (
    $1, $2, $3, $4, $5, $6, $7,
    $8, $9, $10, $11,
    $12, $13,
    $14, $15, $16,
    $17, $18, $19
)
RETURNING *;

-- name: UpdateUserFull :one
-- Used in CreateOrUpdateUser to merge accounts
UPDATE users SET
    password_hash = COALESCE(NULLIF($2, ''), password_hash), -- Update only if new value is provided
    handle = COALESCE($3, handle),
    first_name = COALESCE($4, first_name),
    last_name = COALESCE($5, last_name),
    avatar_url = COALESCE($6, avatar_url),
    email_verified = $7,

    github_id = COALESCE($8, github_id),
    github_email = COALESCE($9, github_email),
    github_handle = COALESCE($10, github_handle),
    github_url = COALESCE($11, github_url),

    google_id = COALESCE($12, google_id),
    google_email = COALESCE($13, google_email),

    facebook_id = COALESCE($14, facebook_id),
    facebook_email = COALESCE($15, facebook_email),
    facebook_url = COALESCE($16, facebook_url),

    spotify_id = COALESCE($17, spotify_id),
    spotify_email = COALESCE($18, spotify_email),
    spotify_url = COALESCE($19, spotify_url),

    updated_at = NOW()
WHERE id = $1
RETURNING *;

-- name: UpdateUserPreferences :one
-- Used in PatchUserMeHandler
UPDATE users SET
    first_name = COALESCE(NULLIF(@first_name::text, ''), first_name),
    last_name = COALESCE(NULLIF(@last_name::text, ''), last_name),
    prefered_theme = COALESCE(NULLIF(@prefered_theme::text, ''), prefered_theme),
    prefered_language = COALESCE(NULLIF(@prefered_language::text, ''), prefered_language),
    updated_at = NOW()
WHERE id = @id
RETURNING *;

-- name: UpdateUserAvatar :one
-- Used in PostUserMeAvatarHandler
UPDATE users SET
    avatar_url = $2,
    updated_at = NOW()
WHERE id = $1
RETURNING *;

-- name: SetLastLoginNow :one
-- Used in GetJwtPostLogin
UPDATE users
SET last_login_at = NOW()
WHERE id = $1
RETURNING *;
