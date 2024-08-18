package db

import "fmt"

const (
	racesList = "list"
	race      = "race"
)

func getRaceQueries() map[string]string {
	return map[string]string{
		racesList: `
			SELECT 
				id, 
				meeting_id, 
				name, 
				number, 
				visible, 
				advertised_start_time,
				CASE 
					WHEN advertised_start_time < CURRENT_TIMESTAMP THEN 'CLOSED'
					ELSE 'OPEN'
				END AS status
			FROM races
		`,
	}
}

func getRaceQuery(id int64) map[string]string {
	return map[string]string{
		race: fmt.Sprintf(`
			SELECT 
				id, 
				meeting_id, 
				name, 
				number, 
				visible, 
				advertised_start_time,
				CASE 
					WHEN advertised_start_time < CURRENT_TIMESTAMP THEN 'CLOSED'
					ELSE 'OPEN'
				END AS status
			FROM races
			WHERE id=%d
		`, id),
	}
}
