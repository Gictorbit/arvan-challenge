package discountdb

import (
	"context"
	dispb "github.com/gictorbit/arvan-challenge/protos/gen/discount/v1"
)

const eventsUsersQuery = `
	SELECT 
		u.id AS user_id,
		u.phone,
		EXTRACT(EPOCH FROM ue.timestamp)::BIGINT AS time_epoch
	FROM 
		user_events ue
	JOIN 
		users u ON ue.user_id = u.id
	WHERE 
		ue.event_code = $1
`

// EventUsers returns a report of users that participate in given event
func (ddb *DiscountDB) EventUsers(ctx context.Context, eventCode string) ([]*dispb.UserCodeUsage, error) {
	rows, err := ddb.GetPgConn().Query(ctx, eventsUsersQuery, eventCode)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var userUsages []*dispb.UserCodeUsage
	for rows.Next() {
		userUsage := &dispb.UserCodeUsage{}
		if e := rows.Scan(
			&userUsage.UserId,
			&userUsage.Phone,
			&userUsage.Timestamp); e != nil {
			return nil, e
		}
		userUsages = append(userUsages, userUsage)
	}
	if e := rows.Err(); e != nil {
		return nil, e
	}
	return userUsages, nil
}
