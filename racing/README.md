## Racing Service

### ListRaces

The ListRaces API lets you retrieve a list of races with the following properties:
- **id** - represents a unique identifier for the race.
- **meeting_id** - represents a unique identifier for the races meeting.
- **name** - is the official name given to the race.
- **number** - represents the number of the race.
- **vsible** - represents whether or not the race is visible.
- **advertised_start_time** - is the time the race is advertised to run.
- **status** - represents the status of a race, which is either open or closed.

The ListRaces request takes an optional filter and orderBy parameter. 
For the **filter paramater** you can specify the meetingIds within an array of integers and/or a showVisibleOnly boolean flag, or leave in blank. 
For the **order paramater** you can specify one or more columns incl. their sorting order, or leave in blank. 
**See examples below...**

Curl request with no filter
```bash
curl -X "POST" "http://localhost:8000/v1/list-races" \
     -H 'Content-Type: application/json' \
     -d $'{
  "filter": {}
}'
```

Curl request with meetingIds filter (targeting meetingId property of races)

```bash
curl -X "POST" "http://localhost:8000/v1/list-races" \
     -H 'Content-Type: application/json' \
     -d $'{
  "filter": {"meetingIds":[1,2,3]}
}'
```

Curl request with showVisibleOnly filter (targeting vivible property of races)

```bash
curl -X "POST" "http://localhost:8000/v1/list-races" \
     -H 'Content-Type: application/json' \
     -d $'{
  "filter": {"showVisibleOnly":true}
}'
```

Curl request with orderBy paramater (targeting any field within races: id, meeting_id, name, number, visible, advertised_start_time)

```bash
curl -X "POST" "http://localhost:8000/v1/list-races" \
     -H 'Content-Type: application/json' \
     -d $'{
  "filter": {},
  "orderBy": "meeting_id"
}'
```

Curl request with multiple orderBy paramaters (targeting any field within races: id, meeting_id, name, number, visible, advertised_start_time)

```bash
curl -X "POST" "http://localhost:8000/v1/list-races" \
     -H 'Content-Type: application/json' \
     -d $'{
  "filter": {},
  "orderBy": "meeting_id, name DESC"
}'
```

Also keep an eye on the **racing_test.go** file which includes a unit test illustrating how to consume the racing service method for various use cases.