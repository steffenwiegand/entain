package db

import (
	"database/sql"
	"testing"
	"time"

	"git.neds.sh/matty/entain/racing/proto/racing"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
)

// testing the races.go/List filter paramater to ensure the right number of races is being returned
func TestRacesList(t *testing.T) {

	// setup racing test data
	testRaces := []struct {
		id                  int64
		meetingId           int64
		name                string
		number              int64
		visible             bool
		advertisedStartTime *timestamppb.Timestamp
	}{
		{id: 1, meetingId: 1, name: "First Race", number: 91, visible: true, advertisedStartTime: timestamppb.New(time.Date(2009, 11, 21, 0, 0, 0, 0, time.UTC))},
		{id: 2, meetingId: 1, name: "Second Race", number: 92, visible: true, advertisedStartTime: timestamppb.New(time.Date(2009, 11, 22, 0, 0, 0, 0, time.UTC))},
		{id: 3, meetingId: 2, name: "Third Race", number: 93, visible: true, advertisedStartTime: timestamppb.New(time.Date(2009, 11, 23, 0, 0, 0, 0, time.UTC))},
		{id: 4, meetingId: 2, name: "Fourth Race", number: 94, visible: false, advertisedStartTime: timestamppb.New(time.Date(2009, 11, 24, 0, 0, 0, 0, time.UTC))},
		{id: 5, meetingId: 9, name: "Fifth Race", number: 95, visible: false, advertisedStartTime: timestamppb.New(time.Date(2009, 11, 25, 0, 0, 0, 0, time.UTC))},
	}

	// setup test db
	db := setupTestDB(t)
	defer db.Close()

	// create table
	createTestTable(t, db)

	// insert test records
	insertTestRaces(t, db, testRaces)

	testRacesRepo := NewRacesRepo(db)

	// test cases
	testCases := []struct {
		name           string
		filter         *racing.ListRacesRequestFilter
		expectedLength int
	}{
		{name: "no filter", filter: &racing.ListRacesRequestFilter{}, expectedLength: 5},
		{name: "meetingIds filter", filter: &racing.ListRacesRequestFilter{MeetingIds: []int64{1, 2}}, expectedLength: 4},
		{name: "visible only filter", filter: &racing.ListRacesRequestFilter{ShowVisibleOnly: true}, expectedLength: 3},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			races, err := testRacesRepo.List(tc.filter, "")
			require.NoError(t, err)

			assert.Len(t, races, tc.expectedLength, "Unexpected number of races.")
		})
	}
}

// testing the races.go/List function orderBy parameter, which determines the sorting order of the returned races
func TestRacesListOrderBy(t *testing.T) {

	// setup racing test data
	testRaces := []struct {
		id                  int64
		meetingId           int64
		name                string
		number              int64
		visible             bool
		advertisedStartTime *timestamppb.Timestamp
	}{
		{id: 1, meetingId: 5, name: "Horse Race", number: 91, visible: true, advertisedStartTime: timestamppb.New(time.Date(2009, 11, 21, 0, 0, 0, 0, time.UTC))},
		{id: 2, meetingId: 1, name: "Dog Race", number: 92, visible: true, advertisedStartTime: timestamppb.New(time.Date(2009, 11, 22, 0, 0, 0, 0, time.UTC))},
		{id: 3, meetingId: 3, name: "Rabbit Race", number: 93, visible: true, advertisedStartTime: timestamppb.New(time.Date(2009, 11, 23, 0, 0, 0, 0, time.UTC))},
		{id: 4, meetingId: 4, name: "Rabbit Race", number: 94, visible: false, advertisedStartTime: timestamppb.New(time.Date(2009, 11, 24, 0, 0, 0, 0, time.UTC))},
		{id: 5, meetingId: 2, name: "Snake Race", number: 95, visible: false, advertisedStartTime: timestamppb.New(time.Date(2009, 11, 25, 0, 0, 0, 0, time.UTC))},
	}

	// setup test db
	db := setupTestDB(t)
	defer db.Close()

	// create table
	createTestTable(t, db)

	// insert test records
	insertTestRaces(t, db, testRaces)

	testRacesRepo := NewRacesRepo(db)

	// test cases
	testCases := []struct {
		name            string
		orderBy         string
		expectedIdOrder []int64
	}{
		{name: "no orderBy", orderBy: "", expectedIdOrder: []int64{1, 2, 3, 4, 5}},
		{name: "advertised_start_time DESC", orderBy: "advertised_start_time DESC", expectedIdOrder: []int64{5, 4, 3, 2, 1}},
		{name: "meeting_id", orderBy: "meeting_id", expectedIdOrder: []int64{2, 5, 3, 4, 1}},
		{name: "name, advertised_start_time DESC", orderBy: "name, advertised_start_time DESC", expectedIdOrder: []int64{2, 1, 4, 3, 5}},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			races, err := testRacesRepo.List(&racing.ListRacesRequestFilter{}, tc.orderBy)
			require.NoError(t, err)

			for i, race := range races {
				assert.Equal(t, int(tc.expectedIdOrder[i]), int(race.Id), "Unexpected race id for test case %s. Expected race id is %d. Returned was %d.", tc.name, tc.expectedIdOrder[i], race.Id)
			}
		})
	}
}

// testing the races.go/List status field, which is determined by the advertised start time of a race
func TestRacesListStatus(t *testing.T) {

	// setup racing test data
	testRaces := []struct {
		id                  int64
		meetingId           int64
		name                string
		number              int64
		visible             bool
		advertisedStartTime *timestamppb.Timestamp
	}{
		{id: 1, meetingId: 5, name: "Horse Race", number: 91, visible: true, advertisedStartTime: timestamppb.New(time.Date(2009, 11, 21, 0, 0, 0, 0, time.UTC))},
		{id: 2, meetingId: 1, name: "Dog Race", number: 92, visible: true, advertisedStartTime: timestamppb.New(time.Date(2009, 11, 22, 0, 0, 0, 0, time.UTC))},
		{id: 3, meetingId: 3, name: "Rabbit Race", number: 93, visible: true, advertisedStartTime: timestamppb.New(time.Now().Add(time.Hour))},
		{id: 4, meetingId: 4, name: "Rabbit Race", number: 94, visible: true, advertisedStartTime: timestamppb.New(time.Now().Add(2 * time.Hour))},
		{id: 5, meetingId: 2, name: "Rabbit Race", number: 95, visible: true, advertisedStartTime: timestamppb.New(time.Now().Add(3 * time.Hour))},
	}

	// setup test db
	db := setupTestDB(t)
	defer db.Close()

	// create table
	createTestTable(t, db)

	// insert test records
	insertTestRaces(t, db, testRaces)

	testRacesRepo := NewRacesRepo(db)

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

	races, err := testRacesRepo.List(&racing.ListRacesRequestFilter{}, "")
	require.NoError(t, err)

	for i, race := range races {
		assert.Equal(t, testCases[i].exptecedStatus, race.Status, "Unexpected race status for race id %d", race.Id)
	}
}

// testing the races.go/GetRace service method
func TestGetRace(t *testing.T) {

	// setup racing test data
	testRaces := []struct {
		id                  int64
		meetingId           int64
		name                string
		number              int64
		visible             bool
		advertisedStartTime *timestamppb.Timestamp
	}{
		{id: 1, meetingId: 5, name: "Horse Race", number: 91, visible: true, advertisedStartTime: timestamppb.New(time.Date(2009, 11, 21, 0, 0, 0, 0, time.UTC))},
		{id: 2, meetingId: 1, name: "Dog Race", number: 92, visible: true, advertisedStartTime: timestamppb.New(time.Date(2009, 11, 22, 0, 0, 0, 0, time.UTC))},
		{id: 3, meetingId: 3, name: "Rabbit Race", number: 93, visible: true, advertisedStartTime: timestamppb.New(time.Now().Add(time.Hour))},
		{id: 4, meetingId: 4, name: "Rabbit Race", number: 94, visible: true, advertisedStartTime: timestamppb.New(time.Now().Add(2 * time.Hour))},
		{id: 5, meetingId: 2, name: "Rabbit Race", number: 95, visible: true, advertisedStartTime: timestamppb.New(time.Now().Add(3 * time.Hour))},
	}

	// setup test db
	db := setupTestDB(t)
	defer db.Close()

	// create table
	createTestTable(t, db)

	// insert test records
	insertTestRaces(t, db, testRaces)

	testRacesRepo := NewRacesRepo(db)

	// test cases
	testCases := []struct {
		name       string
		id         int64
		exptecedId int64
	}{
		{name: "GetRace id 1 test", id: 1, exptecedId: 1},
		{name: "GetRace id 2 test", id: 2, exptecedId: 2},
		{name: "GetRace id 3 test", id: 3, exptecedId: 3},
		{name: "GetRace id 4 test", id: 4, exptecedId: 4},
		{name: "GetRace id 5 test", id: 5, exptecedId: 5},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			race, err := testRacesRepo.Get(tc.id)
			require.NoError(t, err)

			assert.Equal(t, tc.exptecedId, race.Id, "Unexpected race returned for race id %d", race.Id)
		})
	}
}

// testing the races.go/GetRace service method for not found
func TestGetRaceNotFound(t *testing.T) {
	// setup test db
	db := setupTestDB(t)
	defer db.Close()

	// create table
	createTestTable(t, db)

	testRacesRepo := NewRacesRepo(db)

	race, err := testRacesRepo.Get(1)

	// expected error from races.go
	expectedErr := ErrRaceNotFound(1)

	assert.Error(t, err)
	assert.Equal(t, expectedErr, err, "Unexpected error for race id 1. Error did not match the expected not found error.")

	// make sure no race was returned
	assert.Nil(t, race, "Unexpected GetRace result for race id 1. An actual race was returned.")
}

func setupTestDB(t *testing.T) *sql.DB {
	db, err := sql.Open("sqlite3", ":memory:")
	require.NoError(t, err, "Unable to open db")
	return db
}

func createTestTable(t *testing.T, db *sql.DB) {
	statement, err := db.Prepare(`CREATE TABLE IF NOT EXISTS races (id INTEGER PRIMARY KEY, meeting_id INTEGER, name TEXT, number INTEGER, visible INTEGER, advertised_start_time DATETIME)`)
	require.NoError(t, err, "Failed to prepare create table statement")
	_, err = statement.Exec()
	require.NoError(t, err, "Failed to execute create table statement")
}

func insertTestRaces(t *testing.T, db *sql.DB, testRaces []struct {
	id                  int64
	meetingId           int64
	name                string
	number              int64
	visible             bool
	advertisedStartTime *timestamppb.Timestamp
}) {
	for _, race := range testRaces {
		statement, err := db.Prepare(`INSERT OR IGNORE INTO races(id, meeting_id, name, number, visible, advertised_start_time) VALUES (?, ?, ?, ?, ?, ?)`)
		require.NoError(t, err, "Failed to prepare insert statement")
		_, err = statement.Exec(
			race.id,
			race.meetingId,
			race.name,
			race.number,
			race.visible,
			race.advertisedStartTime.AsTime(),
		)
		require.NoError(t, err, "Failed to insert test race")
	}
}
