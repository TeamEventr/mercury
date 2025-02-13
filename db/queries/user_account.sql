-- name: DBCheckUserExist :one
-- Check whether a user_account exists or not.
SELECT EXISTS (
  SELECT 1
  FROM 
    user_account
  WHERE 
    username = $1
); 

-- name: DBCheckUserBanned :one
-- Check whether a user_account is banned or not.
SELECT EXISTS (
  SELECT 1 
  FROM 
    user_account
  WHERE
    username = $1
    AND status = 'banned'
);

-- name: DBCheckUserDisabled :one
-- Check whether a user_account is disabled or not.
SELECT EXISTS (
  SELECT 1 
  FROM
    user_account
  WHERE
    username = $1
    AND status = 'disabled'
);

-- name: DBLoginUser :one
-- Check login credentials post hashing the password against an active account.
SELECT 
  username,
  first_name,
  last_name,
  gender,
  email,
  avatar,
  city
FROM 
  user_account
WHERE
  username = $1
  AND password_login = true
  AND password = $2
  AND status = 'active';

-- name: DBOnboardUser :one
-- Insert a onboarding record with an otp and an expiry time.
INSERT INTO 
  user_onboarding (
    username, 
    password, 
    email, 
    otp, 
    expiry_at
  )
VALUES
  ($1, $2, $3, $4, $5)
RETURNING 
  username, 
  email, 
  otp, 
  expiry_at;

-- name: DBVerifyUserOTP :one
SELECT
  username,
  password,
  email
FROM
  user_onboarding
WHERE
  username = $1
  AND otp = $2
  AND expiry_at >= NOW() - INTERVAL '5 minutes';

-- name: DBGetUserOnboardingOTP :one
SELECT *
FROM
  user_onboarding
WHERE
  username = $1
  AND expiry_at >= NOW() - INTERVAL '5 minutes';

-- name: DBCreateUserAccount :one
INSERT INTO
  user_account (
    username, 
    password_login, 
    password, 
    email, 
    loggedin_at, 
    refresh_token
  )
VALUES
  ($1, true, $2, $3, $4, $5)
RETURNING 
  username, 
  first_name, 
  middle_name,
  last_name,
  gender,
  email, 
  avatar,
  city,
  refresh_token;

-- name: DBEditUserAccount :one
UPDATE 
  user_account
SET
  first_name = $2,
  middle_name = $3,
  last_name = $4,
  gender = $5,
  city = $5
WHERE
  username = $1
  AND status = 'active'
RETURNING
  username,
  first_name,
  middle_name,
  last_name,
  gender,
  city;

-- name: DBEditUserAvatar :one
UPDATE
  user_account
SET
  avatar = $2
WHERE
  username = $1
  AND status = 'active'
RETURNING
  avatar;

-- name: DBDeleteUserAvatar :one
UPDATE
  user_account
SET
  avatar = null
WHERE
  username = $1
  AND status = 'active'
RETURNING
  avatar;

-- name: DBDisableUserAccount :exec
UPDATE
  user_account
SET
  status = 'disabled'
WHERE
  username = $1
  AND status = 'active';

-- name: DBEnableUserAccount :exec
UPDATE
  user_account
SET
  status = 'active'
WHERE
  username = $1
  AND status = 'disabled';
