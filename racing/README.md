## Racing Service

The racing service has a `ListRaces` and a `GetRace` endpoint which lets you retrieve a list of races with the following properties:
- **id** - represents a unique identifier for the race.
- **meeting_id** - represents a unique identifier for the races meeting.
- **name** - is the official name given to the race.
- **number** - represents the number of the race.
- **vsible** - represents whether or not the race is visible.
- **advertised_start_time** - is the time the race is advertised to run.
- **status** - represents the status of a race, which is either open or closed.

`POST: /v1/list-races`

The `ListRaces` request takes an optional filter and orderBy parameters as Json. 
For the **filter paramater** you can specify the meetingIds within an array of integers and/or a showVisibleOnly boolean flag, or leave in blank. 
For the **order paramater** you can specify one or more columns incl. their sorting order, or leave in blank. 

`GET: /v1/races/{id}`

The `GetRace` request takes an id which is the a unique identifier for the race and returns a single race with the above properties, or a not found error if no race can be found with the id provided.

**See examples below** with API requests being made via the api gateway (port 8000), which forwards the requests to the racing service (port 9000)...
Also keep an eye on the **racing_test.go** file which includes unit tests illustrating how to consume the racing service methods for various use cases.

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

Curl request to get a race by id

```bash
curl -X "GET" "http://localhost:8000/v1/races/10" \
     -H 'Content-Type: application/json'
```