#### Built With

* Go (Echo)
* Docker

## Running in local
If you want to run this project on your local machine, do the followings

### Without using docker

```
go get -u -t -d -v ./...
```

```
go mod download
```

```
go mod tidy
```

```
go run main.go
```

### Using docker

```
docker build -t mezink .
```

<br>

available endpoints :
- `POST` localhost:9090/records
- 

### Example Request Body Payload for `POST` /records :
```
{
    "startDate": "2024-01-01",
    "endDate": "2024-12-31",
    "minCount": 0,
    "maxCount": 1000
}
```
