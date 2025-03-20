-- name: CreatePriceTierQuery :copyfrom
INSERT INTO
    price_tier(
        event_id,
        name,
        validity_start,
        validity_end,
        price,
        seat_available,
        total_seat
    )
VALUES ($1, $2, $3, $4, $5, $6, $6);
