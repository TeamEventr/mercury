-- name: FetchMyTicketsQuery :many
-- This query will return paginated results with 25 tickets at a time as
--  number of tickets would keep growing and we do not want to return back
--  everything every single time.
WITH active_user_tickets AS (
    SELECT
        t.id AS ticket_id,
        t.seats_booked AS seats_booked,
        e.title AS event_name,
        pt.name AS price_tier_name,
        t.created_at AS created_at
    FROM
        ticket t
    JOIN user_account AS u ON
        t.username = u.username
    JOIN event AS e ON
        t.event_id = e.id
    JOIN price_tier AS pt ON
        t.tier_id = pt.id
    WHERE
        u.username = $1
        AND u.loggedin_at IS NOT NULL
        AND u.status = 'active'
)
SELECT
    ticket_id,
    seats_booked,
    event_name,
    price_tier_name,
    created_at
FROM
    active_user_tickets
ORDER BY
    created_at DESC
LIMIT 25
OFFSET $2;
