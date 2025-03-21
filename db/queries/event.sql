-- name: CreateEventQuery :one
INSERT INTO
    event (
    title,
    type,
    host_id,
    description,
    venue,
    tags,
    age_limit,
    start_time
    )
VALUES
   ($1, $2, $3, $4, $5, $6, $7, $8)
RETURNING
    title,
    type,
    host_id,
    description,
    venue,
    tags,
    age_limit,
    start_time;

-- name: EditEventQuery :exec
-- Edit an existing Event based on the EventId
-- UPDATE;

-- name: PublishEventQuery :exec
-- Publish Event based on EventId and hostId
UPDATE event
SET
    visibility = 'published',
    updated_at = NOW()
WHERE
    id = $1 AND
    host_id = $2;

-- name: DeleteUnpublishedEventQuery :one
-- Delete an unpublished Event based on EventId and HostId
DELETE FROM event
WHERE
    host_id = $1 AND
    id = $2 AND
    visibility = 'draft'
RETURNING
    id;

-- name: FetchEventsPaginatedQuery :many
-- Fetch all Events in a paginated format
SELECT
    title,
    cover_picture_url,
    tags,
    venue,
    start_time,
    age_limit
FROM event
WHERE
    visibility = 'published'
ORDER BY
    start_time
LIMIT 25
OFFSET $1;

-- name: FetchEventByIdQuery :one
-- Fetch all Event details by EventId
SELECT
    e.title as event_name,
    e.type as event_type,
    e.description AS description,
    e.thumbnail_url AS event_poster_url,
    e.tags AS tags,
    e.venue AS venue,
    e.start_time AS start_time,
    e.end_time AS end_time,
    e.age_limit AS age_limit,
    array_agg(json_build_object(
        'first_name', u.first_name,
        'last_name', u.last_name,
        'avatar', u.avatar,
        'username', u.username
    )) AS artists,
    array_agg(json_build_object(
        'id', pt.id,
        'name', pt.name,
        'price', pt.price,
        'seat_available', pt.seat_available
    )) AS price_tiers
FROM
    event AS e
INNER JOIN event_artist AS ea ON
    e.id = ea.event_id
INNER JOIN user_account AS u ON
    u.username = ea.username
INNER JOIN price_tier AS pt ON
    e.id = pt.event_id
WHERE
    e.id = $1 AND
    e.visibility = 'published' AND
    pt.booking_status = 'open';

-- name: FetchEventByHostQuery :many
-- Fetch all Event by HostId
SELECT
    e.title,
    e.type,
    e.cover_picture_url,
    e.visibility,
    e.start_time
FROM event AS e
INNER JOIN host as h ON
    h.id = e.host_id
WHERE
    h.username = $1;

-- name: DeleteEventPosterQuery :one
-- Delete an Event poster based on eventId
UPDATE event
SET
    cover_picture_url = null,
    updated_at = NOW()
WHERE
    visibility = 'draft' AND
    id = $1 AND
    host_id = $2
RETURNING
    id;
