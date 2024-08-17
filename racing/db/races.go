package db

import (
	"database/sql"
	"fmt"
	"regexp"
	"strings"
	"sync"
	"time"

	"github.com/golang/protobuf/ptypes"
	_ "github.com/mattn/go-sqlite3"

	"git.neds.sh/matty/entain/racing/proto/racing"
)

// RacesRepo provides repository access to races.
type RacesRepo interface {
	// Init will initialise our races repository.
	Init() error

	// List will return a list of races.
	List(filter *racing.ListRacesRequestFilter, orderBy string) ([]*racing.Race, error)

	// Get returns a race for a given race id
	Get(id int64) (*racing.Race, error)
}

type racesRepo struct {
	db   *sql.DB
	init sync.Once
}

// NewRacesRepo creates a new races repository.
func NewRacesRepo(db *sql.DB) RacesRepo {
	return &racesRepo{db: db}
}

// Init prepares the race repository dummy data.
func (r *racesRepo) Init() error {
	var err error

	r.init.Do(func() {
		// For test/example purposes, we seed the DB with some dummy races.
		err = r.seed()
	})

	return err
}

// Lists all races. The paramter filter, orderBy can be left blank. By default all races are returned, sorted by advertised_start_time.
func (r *racesRepo) List(filter *racing.ListRacesRequestFilter, orderBy string) ([]*racing.Race, error) {
	var (
		err   error
		query string
		args  []interface{}
	)

	query = getRaceQueries()[racesList]

	query, args = r.applyFilter(query, filter, orderBy)

	query = r.applyOrderBy(query, orderBy)

	rows, err := r.db.Query(query, args...)
	if err != nil {
		return nil, err
	}

	return r.scanRaces(rows)
}

func (r *racesRepo) applyFilter(query string, filter *racing.ListRacesRequestFilter, orderBy string) (string, []interface{}) {
	var (
		clauses []string
		args    []interface{}
	)

	if filter == nil {
		return query, args
	}

	if len(filter.MeetingIds) > 0 {
		clauses = append(clauses, "meeting_id IN ("+strings.Repeat("?,", len(filter.MeetingIds)-1)+"?)")

		for _, meetingID := range filter.MeetingIds {
			args = append(args, meetingID)
		}
	}

	if filter.ShowVisibleOnly {
		clauses = append(clauses, "visible=1")
	}

	if len(clauses) != 0 {
		query += " WHERE " + strings.Join(clauses, " AND ")
	}

	return query, args
}

// applyOrderBy appends an ORDER BY clause to the query by validating the order by paramter
func (r *racesRepo) applyOrderBy(query string, orderBy string) string {
	// default case
	if strings.TrimSpace(orderBy) == "" {
		query += " ORDER BY advertised_start_time"
		return query
	}

	orderFields := strings.Split(orderBy, ",")

	var orderClauses []string

	// loop through each field provided and validate it
	for _, field := range orderFields {
		field = strings.TrimSpace(field)

		if validateOrderByField(field) && field != "" {
			orderClauses = append(orderClauses, field)
		}
	}

	// append the order clauses to the query
	if len(orderClauses) > 0 {
		query += " ORDER BY " + strings.Join(orderClauses, ", ")
	} else {
		// If no valid order clauses, default to ordering by advertised_start_time
		query += " ORDER BY advertised_start_time"
	}

	return query
}

// validateOrderByField checks if the field is a valid column name
func validateOrderByField(field string) bool {

	// regular expression to allow alphanumeric characters, underscores, and sorting keywords
	// it does not allow SQL keywords (like INSERT, UPDATE) or special characters
	re := regexp.MustCompile(`^[a-zA-Z0-9_]+(\s+(ASC|DESC))?$`)
	return re.MatchString(field)
}

func (m *racesRepo) scanRaces(
	rows *sql.Rows,
) ([]*racing.Race, error) {
	var races []*racing.Race

	for rows.Next() {
		var race racing.Race
		var advertisedStart time.Time

		if err := rows.Scan(&race.Id, &race.MeetingId, &race.Name, &race.Number, &race.Visible, &advertisedStart, &race.Status); err != nil {
			if err == sql.ErrNoRows {
				return nil, nil
			}

			return nil, err
		}

		ts, err := ptypes.TimestampProto(advertisedStart)
		if err != nil {
			return nil, err
		}

		race.AdvertisedStartTime = ts

		races = append(races, &race)
	}

	return races, nil
}

// Gets a single race by id.
func (r *racesRepo) Get(id int64) (*racing.Race, error) {
	var (
		query           string
		advertisedStart time.Time
	)

	query = getRaceQuery(id)[race]

	var race racing.Race

	row := r.db.QueryRow(query, id)
	err := row.Scan(&race.Id, &race.MeetingId, &race.Name, &race.Number, &race.Visible, &advertisedStart, &race.Status)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, ErrRaceNotFound(id)
		}
		return nil, err
	}

	return &race, nil
}

// ErrRaceNotFound is returned when a race is not found.
func ErrRaceNotFound(id int64) error {
	return fmt.Errorf("race with id %d not found", id)
}
