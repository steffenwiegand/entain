package db

import (
	"database/sql"
	"testing"
	"time"

	"git.neds.sh/matty/entain/sports/proto/sports"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
)

// testing the sports.go/List filter paramater to ensure the right number of events is being returned
func TestEventsList(t *testing.T) {

	// setup event test data
	testEvents := []struct {
		id                  int64
		name                string
		categoryId          int64
		division            string
		country             string
		location            string
		advertisedStartTime *timestamppb.Timestamp
		visible             bool
	}{
		{id: 1, categoryId: 1, name: "Chelsea vs Arsenal", division: "Premier League", country: "England", location: "London", advertisedStartTime: timestamppb.New(time.Date(2009, 11, 21, 0, 0, 0, 0, time.UTC)), visible: true},
		{id: 2, categoryId: 1, name: "Tottenham vs Manchester", division: "Premier League", country: "England", location: "London", advertisedStartTime: timestamppb.New(time.Date(2009, 11, 22, 0, 0, 0, 0, time.UTC)), visible: true},
		{id: 3, categoryId: 2, name: "Butterly Final 100m", division: "Paris Olympics", country: "France", location: "Paris", advertisedStartTime: timestamppb.New(time.Date(2009, 11, 23, 0, 0, 0, 0, time.UTC)), visible: true},
		{id: 4, categoryId: 2, name: "Olympics Freestyle Final 100m", division: "Paris Olympics", country: "France", location: "Paris", advertisedStartTime: timestamppb.New(time.Date(2009, 11, 24, 0, 0, 0, 0, time.UTC)), visible: false},
		{id: 5, categoryId: 9, name: "Melbourne Grand Prix", division: "Formula 1", country: "Australia", location: "Melbourne", advertisedStartTime: timestamppb.New(time.Date(2009, 11, 25, 0, 0, 0, 0, time.UTC)), visible: false},
	}

	// setup test db
	db := setupTestDB(t)
	defer db.Close()

	// create table
	createTestTable(t, db)

	// insert test records
	insertTestEvents(t, db, testEvents)

	testSportsRepo := NewSportsRepo(db)

	// test cases
	testCases := []struct {
		name           string
		filter         *sports.ListEventsRequestFilter
		expectedLength int
	}{
		{name: "no filter", filter: &sports.ListEventsRequestFilter{}, expectedLength: 5},
		{name: "categoryIds filter", filter: &sports.ListEventsRequestFilter{CategoryIds: []int64{1, 2}}, expectedLength: 4},
		{name: "visible only filter", filter: &sports.ListEventsRequestFilter{ShowVisibleOnly: true}, expectedLength: 3},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			events, err := testSportsRepo.List(tc.filter, "")
			require.NoError(t, err)

			assert.Len(t, events, tc.expectedLength, "Unexpected number of events.")
		})
	}
}

// testing the sports.go/List function orderBy parameter, which determines the sorting order of the returned events
func TestEventsListOrderBy(t *testing.T) {

	// setup event test data
	testEvents := []struct {
		id                  int64
		name                string
		categoryId          int64
		division            string
		country             string
		location            string
		advertisedStartTime *timestamppb.Timestamp
		visible             bool
	}{
		{id: 1, categoryId: 5, name: "Chelsea vs Arsenal", division: "Premier League", country: "England", location: "London", advertisedStartTime: timestamppb.New(time.Date(2009, 11, 21, 0, 0, 0, 0, time.UTC)), visible: true},
		{id: 2, categoryId: 1, name: "Tottenham vs Manchester", division: "Premier League", country: "England", location: "London", advertisedStartTime: timestamppb.New(time.Date(2009, 11, 22, 0, 0, 0, 0, time.UTC)), visible: true},
		{id: 3, categoryId: 3, name: "Butterly Final 100m", division: "Paris Olympics", country: "France", location: "Paris", advertisedStartTime: timestamppb.New(time.Date(2009, 11, 23, 0, 0, 0, 0, time.UTC)), visible: true},
		{id: 4, categoryId: 4, name: "Butterly Final 100m", division: "Paris Olympics", country: "France", location: "Paris", advertisedStartTime: timestamppb.New(time.Date(2009, 11, 24, 0, 0, 0, 0, time.UTC)), visible: false},
		{id: 5, categoryId: 2, name: "Melbourne Grand Prix", division: "Formula 1", country: "Australia", location: "Melbourne", advertisedStartTime: timestamppb.New(time.Date(2009, 11, 25, 0, 0, 0, 0, time.UTC)), visible: false},
	}

	// setup test db
	db := setupTestDB(t)
	defer db.Close()

	// create table
	createTestTable(t, db)

	// insert test records
	insertTestEvents(t, db, testEvents)

	testEventsRepo := NewSportsRepo(db)

	// test cases
	testCases := []struct {
		name            string
		orderBy         string
		expectedIdOrder []int64
	}{
		{name: "no orderBy", orderBy: "", expectedIdOrder: []int64{1, 2, 3, 4, 5}},
		{name: "advertised_start_time DESC", orderBy: "advertised_start_time DESC", expectedIdOrder: []int64{5, 4, 3, 2, 1}},
		{name: "category_id", orderBy: "category_id", expectedIdOrder: []int64{2, 5, 3, 4, 1}},
		{name: "name, advertised_start_time DESC", orderBy: "name, advertised_start_time DESC", expectedIdOrder: []int64{4, 3, 1, 5, 2}},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			events, err := testEventsRepo.List(&sports.ListEventsRequestFilter{}, tc.orderBy)
			require.NoError(t, err)

			for i, event := range events {
				assert.Equal(t, int(tc.expectedIdOrder[i]), int(event.Id), "Unexpected event id for test case %s. Expected event id is %d. Returned was %d.", tc.name, tc.expectedIdOrder[i], event.Id)
			}
		})
	}
}

// testing the sports.go/List status field, which is determined by the advertised start time of an event
func TestEventsListStatus(t *testing.T) {

	// setup event test data
	testEvents := []struct {
		id                  int64
		name                string
		categoryId          int64
		division            string
		country             string
		location            string
		advertisedStartTime *timestamppb.Timestamp
		visible             bool
	}{
		{id: 1, categoryId: 1, name: "Chelsea vs Arsenal", division: "Premier League", country: "England", location: "London", advertisedStartTime: timestamppb.New(time.Date(2009, 11, 21, 0, 0, 0, 0, time.UTC)), visible: true},
		{id: 2, categoryId: 1, name: "Tottenham vs Manchester", division: "Premier League", country: "England", location: "London", advertisedStartTime: timestamppb.New(time.Date(2009, 11, 22, 0, 0, 0, 0, time.UTC)), visible: true},
		{id: 3, categoryId: 2, name: "Butterly Final 100m", division: "Paris Olympics", country: "France", location: "Paris", advertisedStartTime: timestamppb.New(time.Now().Add(time.Hour)), visible: true},
		{id: 4, categoryId: 2, name: "Olympics Freestyle Final 100m", division: "Paris Olympics", country: "France", location: "Paris", advertisedStartTime: timestamppb.New(time.Now().Add(2 * time.Hour)), visible: false},
		{id: 5, categoryId: 9, name: "Melbourne Grand Prix", division: "Formula 1", country: "Australia", location: "Melbourne", advertisedStartTime: timestamppb.New(time.Now().Add(3 * time.Hour)), visible: false},
	}

	// setup test db
	db := setupTestDB(t)
	defer db.Close()

	// create table
	createTestTable(t, db)

	// insert test records
	insertTestEvents(t, db, testEvents)

	testEventsRepo := NewSportsRepo(db)

	// test cases
	testCases := []struct {
		name           string
		id             int64
		exptecedStatus string
	}{
		{name: "closed status", id: 1, exptecedStatus: "CLOSED"},
		{name: "closed status", id: 2, exptecedStatus: "CLOSED"},
		{name: "open status", id: 3, exptecedStatus: "OPEN"},
		{name: "open status", id: 4, exptecedStatus: "OPEN"},
		{name: "open status", id: 5, exptecedStatus: "OPEN"},
	}

	events, err := testEventsRepo.List(&sports.ListEventsRequestFilter{}, "")
	require.NoError(t, err)

	for i, event := range events {
		assert.Equal(t, testCases[i].exptecedStatus, event.Status, "Unexpected event status for event id %d", event.Id)
	}
}

// testing the sports.go/GetEvent service method
func TestGetEvent(t *testing.T) {

	// setup event test data
	testEvents := []struct {
		id                  int64
		name                string
		categoryId          int64
		division            string
		country             string
		location            string
		advertisedStartTime *timestamppb.Timestamp
		visible             bool
	}{
		{id: 1, categoryId: 1, name: "Chelsea vs Arsenal", division: "Premier League", country: "England", location: "London", advertisedStartTime: timestamppb.New(time.Date(2009, 11, 21, 0, 0, 0, 0, time.UTC)), visible: true},
		{id: 2, categoryId: 1, name: "Tottenham vs Manchester", division: "Premier League", country: "England", location: "London", advertisedStartTime: timestamppb.New(time.Date(2009, 11, 22, 0, 0, 0, 0, time.UTC)), visible: true},
		{id: 3, categoryId: 8, name: "Butterly Final 100m", division: "Paris Olympics", country: "France", location: "Paris", advertisedStartTime: timestamppb.New(time.Date(2009, 11, 23, 0, 0, 0, 0, time.UTC)), visible: true},
		{id: 4, categoryId: 8, name: "Butterly Final 100m", division: "Paris Olympics", country: "France", location: "Paris", advertisedStartTime: timestamppb.New(time.Date(2009, 11, 24, 0, 0, 0, 0, time.UTC)), visible: false},
		{id: 5, categoryId: 2, name: "Melbourne Grand Prix", division: "Formula 1", country: "Australia", location: "Melbourne", advertisedStartTime: timestamppb.New(time.Date(2009, 11, 25, 0, 0, 0, 0, time.UTC)), visible: false},
	}

	// setup test db
	db := setupTestDB(t)
	defer db.Close()

	// create table
	createTestTable(t, db)

	// insert test records
	insertTestEvents(t, db, testEvents)

	testSportsRepo := NewSportsRepo(db)

	// test cases
	testCases := []struct {
		name       string
		id         int64
		exptecedId int64
	}{
		{name: "GetEvent id 1 test", id: 1, exptecedId: 1},
		{name: "GetEvent id 2 test", id: 2, exptecedId: 2},
		{name: "GetEvent id 3 test", id: 3, exptecedId: 3},
		{name: "GetEvent id 4 test", id: 4, exptecedId: 4},
		{name: "GetEvent id 5 test", id: 5, exptecedId: 5},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			event, err := testSportsRepo.Get(tc.id)
			require.NoError(t, err)

			assert.Equal(t, tc.exptecedId, event.Id, "Unexpected event returned for event id %d", event.Id)
		})
	}
}

// testing the sports.go/GetEvent service method for not found
func TestGetEventNotFound(t *testing.T) {
	// setup test db
	db := setupTestDB(t)
	defer db.Close()

	// create table
	createTestTable(t, db)

	testSportsRepo := NewSportsRepo(db)

	event, err := testSportsRepo.Get(1)

	// expected error from events.go
	expectedErr := ErrEventNotFound(1)

	assert.Error(t, err)
	assert.Equal(t, expectedErr, err, "Unexpected error for event id 1. Error did not match the expected not found error.")

	// make sure no event was returned
	assert.Nil(t, event, "Unexpected GetEvent result for event id 1. An actual event was returned.")
}

func setupTestDB(t *testing.T) *sql.DB {
	db, err := sql.Open("sqlite3", ":memory:")
	require.NoError(t, err, "Unable to open db")
	return db
}

func createTestTable(t *testing.T, db *sql.DB) {
	statement, err := db.Prepare(`CREATE TABLE IF NOT EXISTS events (id INTEGER PRIMARY KEY, name TEXT, category_id INTEGER, division TEXT, country TEXT, location TEXT, advertised_start_time DATETIME, visible INTEGER)`)
	require.NoError(t, err, "Failed to prepare create table statement")
	_, err = statement.Exec()
	require.NoError(t, err, "Failed to execute create table statement")
}

func insertTestEvents(t *testing.T, db *sql.DB, testEvents []struct {
	id                  int64
	name                string
	categoryId          int64
	division            string
	country             string
	location            string
	advertisedStartTime *timestamppb.Timestamp
	visible             bool
}) {
	for _, event := range testEvents {
		statement, err := db.Prepare(`INSERT OR IGNORE INTO events(id, name, category_id, division, country, location, advertised_start_time, visible) VALUES (?, ?, ?, ?, ?, ?, ?, ?)`)
		require.NoError(t, err, "Failed to prepare insert statement")
		defer statement.Close()

		_, err = statement.Exec(
			event.id,
			event.name,
			event.categoryId,
			event.division,
			event.country,
			event.location,
			event.advertisedStartTime.AsTime(),
			event.visible,
		)
		require.NoError(t, err, "Failed to insert test event")
	}
}
