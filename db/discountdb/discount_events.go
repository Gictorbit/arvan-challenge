package discountdb

import (
	"context"
	dispb "github.com/gictorbit/arvan-challenge/protos/gen/discount/v1"
)

const eventsQuery = `
SELECT 
    code,
    title,
    description,
    gift_amount,
    max_users,
    user_count,
    start_time,
    end_time,
    created_at
FROM 
    events
WHERE 
    published = TRUE;
`

func (ddb *DiscountDB) GetPublishedEvents(ctx context.Context) ([]*dispb.Event, error) {
	rows, err := ddb.postgresConn.Query(ctx, `
        SELECT 
            code,
            title,
            description,
            gift_amount,
            max_users,
            start_time,
            end_time
        FROM 
            events
        WHERE 
            published = TRUE;
    `)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var events []*dispb.Event
	for rows.Next() {
		event := &dispb.Event{}
		if err := rows.Scan(&event.Code, &event.Title, &event.Description, &event.GiftAmount, &event.MaxUsers, &event.StartTime, &event.EndTime); err != nil {
			return nil, err
		}
		events = append(events, event)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return events, nil
}
