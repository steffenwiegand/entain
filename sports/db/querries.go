package db

import "fmt"

const (
	eventsList = "list"
	event      = "event"
)

func getEventQueries() map[string]string {
	return map[string]string{
		eventsList: `
			SELECT 
				id,
				name,
				category_id,
				division,
				country,
				location,
				advertised_start_time,
				CASE 
					WHEN advertised_start_time < CURRENT_TIMESTAMP THEN 'CLOSED'
					ELSE 'OPEN'
				END AS status,
				visible
			FROM events
		`,
	}
}

func getEventQuery(id int64) map[string]string {
	return map[string]string{
		event: fmt.Sprintf(`
			SELECT 
				id,
				name,
				category_id,
				division,
				country,
				location,
				advertised_start_time,
				CASE 
					WHEN advertised_start_time < CURRENT_TIMESTAMP THEN 'CLOSED'
					ELSE 'OPEN'
				END AS status,
				visible
			FROM events
			WHERE id=%d
		`, id),
	}
}
