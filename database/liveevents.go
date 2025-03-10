package database

import (
	"context"
	"database/sql"
	"time"
)

// AllActiveEvents retrieves all active events (where end_time is null)
func (db *DB) AllActiveEvents(ctx context.Context) ([]LiveEvent, error) {
	sqlStmt := `SELECT id, title, description, start_time, end_time, rewards 
				FROM live_events 
				WHERE end_time IS NULL 
				ORDER BY id`

	rows, err := db.conn.QueryContext(ctx, sqlStmt)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	events := []LiveEvent{}

	for rows.Next() {
		var (
			event     LiveEvent
			startTime int64
			endTime   sql.NullInt64
		)

		if err := rows.Scan(
			&event.ID,
			&event.Title,
			&event.Description,
			&startTime,
			&endTime,
			&event.Rewards,
		); err != nil {
			return nil, err
		}

		event.StartTime = time.Unix(startTime, 0)
		if endTime.Valid {
			t := time.Unix(endTime.Int64, 0)
			event.EndTime = &t
		}

		events = append(events, event)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return events, nil
}

// AllEvent retrieves all events
func (db *DB) AllEvent(ctx context.Context) ([]LiveEvent, error) {
	sqlStmt := `SELECT id, title, description, start_time, end_time, rewards 
				FROM live_events 
				ORDER BY id`

	rows, err := db.conn.QueryContext(ctx, sqlStmt)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	events := []LiveEvent{}

	for rows.Next() {
		var (
			event     LiveEvent
			startTime int64
			endTime   sql.NullInt64
		)

		if err := rows.Scan(
			&event.ID,
			&event.Title,
			&event.Description,
			&startTime,
			&endTime,
			&event.Rewards,
		); err != nil {
			return nil, err
		}

		event.StartTime = time.Unix(startTime, 0)

		if endTime.Valid {
			t := time.Unix(endTime.Int64, 0)
			event.EndTime = &t
		}

		events = append(events, event)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return events, nil
}

// GetEvent retrieves one event from db with is ID
func (db *DB) GetEvent(ctx context.Context, id int) (LiveEvent, error) {
	sqlStmt := `SELECT * FROM live_events where id = $1`

	row := db.conn.QueryRowContext(ctx, sqlStmt, id)

	var (
		event     LiveEvent
		startTime int64
		endTime   sql.NullInt64
	)

	if err := row.Scan(
		&event.ID,
		&event.Title,
		&event.Description,
		&startTime,
		&endTime,
		&event.Rewards,
	); err != nil {
		return LiveEvent{}, err
	}

	event.StartTime = time.Unix(startTime, 0)
	if endTime.Valid {
		t := time.Unix(endTime.Int64, 0)
		event.EndTime = &t
	}

	return event, nil
}

// CreateEvent creates one event in db
func (db *DB) CreateEvent(ctx context.Context, event LiveEvent) (int64, error) {
	sqlStmt := `INSERT INTO live_events (title, description, start_time, rewards)
				VALUES ($1,$2,$3,$4)`

	result, err := db.conn.ExecContext(ctx, sqlStmt, event.Title, event.Description, event.StartTime.Unix(), event.Rewards)
	if err != nil {
		return 0, err
	}

	return result.LastInsertId()
}

// DeleteEvent deletes one event in db
func (db *DB) DeleteEvent(ctx context.Context, id int) error {
	sqlStmt := `DELETE FROM live_events 
				WHERE id = $1`

	_, err := db.conn.ExecContext(ctx, sqlStmt, id)

	return err
}

// UpdateEvent updates one event in db
func (db *DB) UpdateEvent(ctx context.Context, event LiveEvent) error {
	sqlStmt := `UPDATE live_events 
				SET title = $1,
					description = $2,
					end_time = $3,
					rewards = $4
				WHERE 
					id = $5
				`

	_, err := db.conn.ExecContext(ctx, sqlStmt, event.Title, event.Description, event.EndTime.Unix(), event.Rewards, event.ID)

	return err
}
