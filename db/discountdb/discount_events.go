package discountdb

import (
	"context"
	dispb "github.com/gictorbit/arvan-challenge/protos/gen/discount/v1"
)

const publishedEventsQuery = `
	SELECT 
		code,
		title,
		description,
		gift_amount,
		max_users,
		EXTRACT(EPOCH FROM start_time)::BIGINT AS start_time_epoch,
		EXTRACT(EPOCH FROM end_time)::BIGINT AS end_time_epoch
	FROM 
		events
	WHERE 
		published = TRUE AND NOW() < end_time;
`

// GetPublishedEvents returns all published events
func (ddb *DiscountDB) GetPublishedEvents(ctx context.Context) ([]*dispb.Event, error) {
	rows, err := ddb.postgresConn.Query(ctx, publishedEventsQuery)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var events []*dispb.Event
	for rows.Next() {
		event := &dispb.Event{}
		if err := rows.Scan(
			&event.Code,
			&event.Title,
			&event.Description,
			&event.GiftAmount,
			&event.MaxUsers,
			&event.StartTime,
			&event.EndTime); err != nil {
			return nil, err
		}
		events = append(events, event)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return events, nil
}
