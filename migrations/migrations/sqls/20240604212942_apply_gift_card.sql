-- +goose Up
-- +goose StatementBegin
CREATE OR REPLACE FUNCTION apply_gift_code(p_phone VARCHAR, p_code UUID)
    RETURNS TABLE
            (
                message       TEXT,
                new_balance   NUMERIC,
                user_id       INTEGER,
                event_code    UUID,
                event_title   VARCHAR,
                event_desc    VARCHAR,
                gift_amount   NUMERIC
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
        RETURN QUERY SELECT 'User not found', 0, NULL, NULL, NULL, NULL, 0;
        RETURN;
    END IF;

    -- Lock the event row for update
    SELECT *
    INTO v_event
    FROM events
    WHERE code = p_code
        FOR UPDATE;

    IF NOT FOUND THEN
        RETURN QUERY SELECT 'Gift code not found', 0, v_user_id, p_code, NULL, NULL, 0;
        RETURN;
    END IF;

    -- Check if event is currently active
    IF NOW() < v_event.start_time OR NOW() > v_event.end_time THEN
        RETURN QUERY SELECT 'Gift code is not active', 0, v_user_id, p_code, v_event.title, v_event.description, 0;
        RETURN;
    END IF;

    -- Check if event is published and if max user limit is reached
    IF v_event.published = FALSE OR v_event.user_count >= v_event.max_users THEN
        RETURN QUERY SELECT 'Gift code not valid or already fully used', 0, v_user_id, p_code, v_event.title, v_event.description, 0;
        RETURN;
    END IF;

    -- Check if user has already used this gift code
    IF EXISTS (SELECT 1 FROM user_events WHERE user_id = v_user_id AND event_code = p_code) THEN
        RETURN QUERY SELECT 'User has already applied this gift code', 0, v_user_id, p_code, v_event.title, v_event.description, 0;
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

    RETURN QUERY SELECT 'Gift code applied successfully', new_balance, v_user_id, p_code, v_event.title, v_event.description, v_event.gift_amount;
EXCEPTION
    WHEN OTHERS THEN
        RETURN QUERY SELECT 'Error applying gift code', 0, v_user_id, p_code, v_event.title, v_event.description, 0;
END;
$$ LANGUAGE plpgsql;



-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP FUNCTION IF EXISTS apply_gift_code;
-- +goose StatementEnd
