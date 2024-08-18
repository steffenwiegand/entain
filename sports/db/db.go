package db

import (
	"math/rand"
	"time"

	"syreclabs.com/go/faker"
)

func (r *sportsRepo) seed() error {
	statement, err := r.db.Prepare(`CREATE TABLE IF NOT EXISTS events (id INTEGER PRIMARY KEY, name TEXT, category_id INTEGER, division TEXT, country TEXT, location TEXT, advertised_start_time DATETIME, visible BOOLEAN)`)
	if err == nil {
		_, err = statement.Exec()
	}

	for i := 1; i <= 100; i++ {
		statement, err = r.db.Prepare(`INSERT OR IGNORE INTO events (id, name, category_id, division, country, location, advertised_start_time, visible) VALUES (?,?,?,?,?,?,?,?)`)
		if err == nil {
			_, err = statement.Exec(
				i,                             // id
				getRandomSportsEvent(),        // name
				faker.Number().Between(1, 10), // category_id
				faker.Team().Name(),           // division
				faker.Address().Country(),     // country
				faker.Address().City(),        // location
				faker.Time().Between(time.Now().AddDate(0, 0, -1), time.Now().AddDate(0, 0, 2)).Format(time.RFC3339), // advertised_start_time
				faker.Number().Between(0, 1), // visible
			)
		}
	}

	return err
}

func getRandomSportsEvent() string {
	events := []string{
		"Mike Tyson vs Muhammad Ali",
		"Bayern Munich vs Real Madrid - Champions League",
		"USA vs France - FIFA World Cup",
		"Warriors vs San Antonio Spurs - NBA Play Offs",
		"LA Lakers vs Dallas Mavericks - NBA Play Offs",
		"Day 1 - Tour de France",
		"Day 2 - Tour de France",
		"Day 3 - Tour de France",
		"Day 4 - Tour de France",
		"Day 5 - Tour de France",
	}
	rand.Seed(time.Now().UnixNano())
	return events[rand.Intn(len(events))]
}
