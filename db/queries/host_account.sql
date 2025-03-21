-- name: AddNewHostOtpQuery :one
INSERT INTO
    host_onboarding(
        username,
        company_name,
        company_email,
        registered,
        hosted_status,
        otp,
        created_at,
        expiry_at
    )
VALUES
    ($1, $2, $3, $4, $5, $6, NOW(), NOW() + INTERVAL '5 minutes')
RETURNING
    username,
    company_name,
    company_email,
    otp,
    expiry_at;

-- name: VerifyHostOTPAndCreateHostQuery :one
WITH verified_host AS (
    SELECT
        username,
        company_name,
        company_email,
        registered,
        hosted_status
    FROM host_onboarding
    WHERE
        host_onboarding.company_email = $1
        AND otp = $2
        AND expiry_at >= NOW()
    )
INSERT INTO
    host(
        username,
        company_name,
        company_email,
        registered,
        hosted_status,
        account_status,
        created_at,
        updated_at
    )
SELECT(
    vh.username,
    vh.company_name,
    vh.company_email,
    vh.registered,
    vh.hosted_status,
    'active',
    NOW(),
    NOW()
) FROM verified_host AS vh
WHERE
    vh.username IS NOT NULL
RETURNING
    username;

-- name: InvalidateHostOtpQuery :one
UPDATE host_onboarding
SET
    expiry_at = NOW()
WHERE
    username = $1
    AND expiry_at > NOW()
RETURNING
    username;
