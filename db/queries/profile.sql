-- name: GetMyProfileQuery :one
SELECT
    username,
    email,
    first_name,
    middle_name,
    last_name,
    dob,
    avatar,
    gender,
    city,
    phone_number
FROM user_account
WHERE
    username = $1
    AND status = 'active'
    AND loggedin_at IS NOT NULL
LIMIT 1;

-- name: EditMyProfileQuery :one
UPDATE user_account
SET
    first_name = COALESCE($2, first_name),
    middle_name = COALESCE($3, middle_name),
    last_name = COALESCE($4, last_name),
    gender = COALESCE($5, gender),
    dob = COALESCE($6, dob),
    city = COALESCE($7, city),
    phone_number = COALESCE($8, phone_number),
    updated_at = NOW()
WHERE
    username = $1
    AND status = 'active'
    AND loggedin_at IS NOT NULL
RETURNING
    username,
    email,
    avatar,
    first_name,
    middle_name,
    last_name,
    gender,
    dob,
    city,
    phone_number;

-- name: DeleteCustomerPfp :one
UPDATE user_account
SET
    avatar = NULL
WHERE
    username = $1
    AND status = 'active'
    AND loggedin_at IS NOT NULL
RETURNING
    username,
    avatar;
