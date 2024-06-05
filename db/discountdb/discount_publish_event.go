package discountdb

import "context"

const publishEventQuery = `
	UPDATE events
		SET published = TRUE
	WHERE code = $1;
`

// PublishEvent publishes an event
func (ddb *DiscountDB) PublishEvent(ctx context.Context, eventCode string) error {
	_, err := ddb.GetPgConn().Exec(ctx, publishEventQuery, eventCode)
	return err
}
