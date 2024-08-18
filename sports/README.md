## Sports Service

The racing service has a `ListEvents` and a `GetEvent` endpoint which lets you retrieve a list of events with the following properties:
- **id** - represents a unique identifier for the event.
  **name** - represents the official name given to the event.
- **category_id** - represents a unique identifier for the event category which identiefies as soccer, baseball, basketball, etc.
- **division** - represents where within the division the event sits in (premier league, championship, league one, league two, etc).
- **country** - represents which country the event takes place
- **location** - represents which country the event takes place
- **advertised_start_time** - is the time the event is advertised to run.
- **status** - reflects the current status of an event. Values are `OPEN` or `CLOSED` or `CANCELED`. 
- **visible** - represents whether or not the event is visible.

`POST: /v1/list-events`

The `ListEvents` request takes an optional filter and orderBy parameters as Json. 
For the **filter paramater** you can specify the meetingIds within an array of integers and/or a showVisibleOnly boolean flag, or leave in blank. 
For the **order paramater** you can specify one or more columns incl. their sorting order, or leave in blank. 

`GET: /v1/events/{id}`

The `GetEvent` request takes an id which is the a unique identifier for an event and returns a single event with the above properties, or a not found error if no event can be found with the id provided.

**See examples below** with API requests being made via the api gateway (port 8000), which forwards the requests to the sporting service (port 10000)...
Also keep an eye on the **sports_test.go** file which includes unit tests illustrating how to consume the sports service methods for various use cases.

In a terminal window, start the sports service...

```bash
cd ./sports

go build && ./sports
➜ INFO[0000] gRPC server listening on: localhost:10000
```

Curl request with no filter
```bash
curl -X "POST" "http://localhost:8000/v1/list-events" \
     -H 'Content-Type: application/json' \
     -d $'{
  "filter": {}
}'
```

Curl request with meetingIds filter (targeting meetingId property of events)

```bash
curl -X "POST" "http://localhost:8000/v1/list-events" \
     -H 'Content-Type: application/json' \
     -d $'{
  "filter": {"categoryIds":[1,2,3]}
}'
```

Curl request with showVisibleOnly filter (targeting vivible property of events)

```bash
curl -X "POST" "http://localhost:8000/v1/list-events" \
     -H 'Content-Type: application/json' \
     -d $'{
  "filter": {"showVisibleOnly":true}
}'
```

Curl request with orderBy paramater (targeting any field within events table)

```bash
curl -X "POST" "http://localhost:8000/v1/list-events" \
     -H 'Content-Type: application/json' \
     -d $'{
  "filter": {},
  "orderBy": "name DESC, category_id"
}'
```

Curl request to get a event by id

```bash
curl -X "GET" "http://localhost:8000/v1/events/10" \
     -H 'Content-Type: application/json'
```