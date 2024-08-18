package db

import (
	"database/sql"
	"fmt"
	"regexp"
	"strings"
	"sync"
	"time"

	"git.neds.sh/matty/entain/sports/proto/sports"
	"github.com/golang/protobuf/ptypes"
	_ "github.com/mattn/go-sqlite3"
)

// SportsRepo provides repository access to events.
type SportsRepo interface {
	// Init will initialise our sports repository.
	Init() error

	// List will return a list of sports.
	List(filter *sports.ListEventsRequestFilter, orderBy string) ([]*sports.Event, error)

	// Get returns an event for a given event id
	Get(id int64) (*sports.Event, error)
}

type sportsRepo struct {
	db   *sql.DB
	init sync.Once
}

// NewSportsRepo creates a new sports repository.
func NewSportsRepo(db *sql.DB) SportsRepo {
	return &sportsRepo{db: db}
}

// Init prepares the sports repository dummy data.
func (r *sportsRepo) Init() error {
	var err error

	r.init.Do(func() {
		// For test/example purposes, we seed the DB with some dummy events.
		err = r.seed()
	})

	return err
}

// Lists all events. The paramter filter, orderBy can be left blank. By default all events are returned, sorted by advertised_start_time.
func (r *sportsRepo) List(filter *sports.ListEventsRequestFilter, orderBy string) ([]*sports.Event, error) {
	var (
		err   error
		query string
		args  []interface{}
	)

	query = getEventQueries()[eventsList]

	query, args = r.applyFilter(query, filter, orderBy)

	query = r.applyOrderBy(query, orderBy)

	rows, err := r.db.Query(query, args...)
	if err != nil {
		return nil, err
	}

	return r.scanEvents(rows)
}

func (r *sportsRepo) applyFilter(query string, filter *sports.ListEventsRequestFilter, orderBy string) (string, []interface{}) {
	var (
		clauses []string
		args    []interface{}
	)

	if filter == nil {
		return query, args
	}

	if len(filter.CategoryIds) > 0 {
		clauses = append(clauses, "category_id IN ("+strings.Repeat("?,", len(filter.CategoryIds)-1)+"?)")

		for _, categoryID := range filter.CategoryIds {
			args = append(args, categoryID)
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
func (r *sportsRepo) applyOrderBy(query string, orderBy string) string {
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

func (m *sportsRepo) scanEvents(
	rows *sql.Rows,
) ([]*sports.Event, error) {
	var events []*sports.Event

	for rows.Next() {
		var event sports.Event
		var advertisedStart time.Time

		if err := rows.Scan(
			&event.Id,
			&event.Name,
			&event.CategoryId,
			&event.Division,
			&event.Country,
			&event.Location,
			&advertisedStart,
			&event.Status,
			&event.Visible); err != nil {
			if err == sql.ErrNoRows {
				return nil, nil
			}

			return nil, err
		}

		ts, err := ptypes.TimestampProto(advertisedStart)
		if err != nil {
			return nil, err
		}

		event.AdvertisedStartTime = ts

		events = append(events, &event)
	}

	return events, nil
}

// Gets a single event by id.
func (r *sportsRepo) Get(id int64) (*sports.Event, error) {
	var (
		query           string
		advertisedStart time.Time
	)

	query = getEventQuery(id)[event]

	var event sports.Event

	row := r.db.QueryRow(query, id)
	err := row.Scan(
		&event.Id,
		&event.Name,
		&event.CategoryId,
		&event.Division,
		&event.Country,
		&event.Location,
		&advertisedStart,
		&event.Status,
		&event.Visible)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, ErrEventNotFound(id)
		}
		return nil, err
	}

	return &event, nil
}

// ErrEventNotFound is returned when a event is not found.
func ErrEventNotFound(id int64) error {
	return fmt.Errorf("event with id %d not found", id)
}
