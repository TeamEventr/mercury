-- name: CheckUserExistQuery :one
-- Check whether a user_account exists or not.
SELECT EXISTS (
  SELECT 1
  FROM
    user_account
  WHERE
    username = $1
);

-- name: CheckUserBannedQuery :one
-- Check whether a user_account is banned or not.
SELECT EXISTS (
  SELECT 1
  FROM
    user_account
  WHERE
    username = $1
    AND status = 'banned'
);

-- name: CheckUserDisabledQuery :one
-- Check whether a user_account is disabled or not.
SELECT EXISTS (
  SELECT 1
  FROM
    user_account
  WHERE
    username = $1
    AND status = 'disabled'
);

-- name: AddNewCustomerOtpQuery :one
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
  ($1, $2, $3, $4, NOW() + INTERVAL '5 minutes')
RETURNING
  username,
  email,
  otp,
  expiry_at;

-- name: VerifyOTPAndCreateAccountQuery :one
-- OTP verification and account creation in a single query
WITH verified_user AS (
    SELECT
        username,
        password,
        email
    FROM user_onboarding
    WHERE
        user_onboarding.email = $1
        AND otp = $2
        AND expiry_at >= NOW()
    )
INSERT INTO
    user_account(
        username,
        password_login,
        password,
        email,
        status,
        created_at,
        updated_at
    )
SELECT (
    vu.username,
    true,
    vu.password,
    vu.email,
    'active',
    NOW(),
    NOW()
) FROM verified_user AS vu
WHERE
    vu.username IS NOT NULL
RETURNING
    username;

-- name: InvalidateOtpQuery :one
UPDATE user_onboarding
SET
    expiry_at = NOW()
WHERE
    username = $1
    AND expiry_at > NOW()
RETURNING
    username;
