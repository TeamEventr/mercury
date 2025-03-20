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
    start_time
;

-- name: EditEventQuery :one
-- Edit an existing Event based on the EventId
-- UPDATE;

-- name: PublishEventQuery :one
-- Publish Event based on EventId and hostId
-- UPDATE;

-- name: DeleteUnpublishedEventQuery :one
-- Delete an unpublished Event based on EventId and HostId
-- DELETE;

-- name: FetchEventsPaginatedQuery :many
-- Fetch all Events in a paginated format
-- SELECT;

-- name: FetchEventByIdQuery :one
-- Fetch all Event details by EventId
-- SELECT;

-- name: FetchEventByHostQuery :many
-- Fetch all Event by HostId
-- SELECT;

-- name: AddEventPosterQuery :one
-- Add an Event poster based on EventId
-- UPDATE;

-- name: DeleteEventPosterQuery :exec
-- Delete an Event poster based on eventId
-- DELETE;
