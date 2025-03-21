-- name: VerifyArtistQuery :exec
SELECT
    username
FROM
    user_account
WHERE
    username = ANY ($1::TEXT[]);

-- name: AddArtistsQuery :copyfrom
INSERT INTO
    event_artist(
        event_id,
        username
    )
VALUES ($1, $2);
