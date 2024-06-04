-- +goose Up
-- +goose StatementBegin
CREATE OR REPLACE FUNCTION apply_gift_code(p_phone VARCHAR, p_code UUID)
    RETURNS TABLE
            (
                success     BOOLEAN,
                message     TEXT,
                new_balance NUMERIC
            )
AS
$$
DECLARE
    v_user_id INTEGER;
    v_event   RECORD;
BEGIN
    -- Find user ID by phone number
    SELECT id
    INTO v_user_id
    FROM users
    WHERE phone = p_phone;

    IF NOT FOUND THEN
        RETURN QUERY SELECT FALSE, 'User not found', 0;
        RETURN;
    END IF;

    -- Lock the event row for update
    SELECT *
    INTO v_event
    FROM events
    WHERE code = p_code
        FOR UPDATE;

    IF NOT FOUND THEN
        RETURN QUERY SELECT FALSE, 'Gift code not found', 0;
        RETURN;
    END IF;

    -- Check if event is currently active
    IF NOW() < v_event.start_time OR NOW() > v_event.end_time THEN
        RETURN QUERY SELECT FALSE, 'Gift code is not active', 0;
        RETURN;
    END IF;

    -- Check if event is published and if max user limit is reached
    IF v_event.published = FALSE OR v_event.user_count >= v_event.max_users THEN
        RETURN QUERY SELECT FALSE, 'Gift code not valid or already fully used', 0;
        RETURN;
    END IF;

    -- Check if user has already used this gift code
    IF EXISTS (SELECT 1 FROM user_events WHERE user_id = v_user_id AND event_code = p_code) THEN
        RETURN QUERY SELECT FALSE, 'User has already applied this gift code', 0;
        RETURN;
    END IF;

    -- Apply gift code to user's balance
    UPDATE users
    SET balance = balance + v_event.gift_amount
    WHERE id = v_user_id
    RETURNING balance INTO new_balance;

    -- Record the usage of the gift code
    INSERT INTO user_events (user_id, event_code)
    VALUES (v_user_id, p_code);

    -- Update the event user count
    UPDATE events
    SET user_count = user_count + 1
    WHERE code = p_code;

    RETURN QUERY SELECT TRUE, 'Gift code applied successfully', new_balance;
EXCEPTION
    WHEN OTHERS THEN
        RETURN QUERY SELECT FALSE, 'Error applying gift code', 0;
END;
$$ LANGUAGE plpgsql;

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP FUNCTION IF EXISTS apply_gift_code;
-- +goose StatementEnd