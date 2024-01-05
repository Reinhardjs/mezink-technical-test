#### Built With

* Go (Echo)
* Docker

<br>

## Running in local
---
If you want to run this project on your local machine, do the followings

<br>

### Without using docker
---

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

<br>

### Using docker
---

```
docker build -t mezink .
```

<br>

### Available endpoints :
- `POST` localhost:9090/records

<br>

---

### Example Request Body Payload for `POST` /records :
```
{
    "startDate": "2024-01-01",
    "endDate": "2024-12-31",
    "minCount": 0,
    "maxCount": 1000
}
```

### Example Response Body Payload for `POST` /records :
```
{
    "code": 0,
    "msg": "Success",
    "records": [
        {
            "id": 1,
            "name": "name 1",
            "marks": [
                100,
                50,
                50
            ],
            "created_at": "2024-01-05T13:12:46+07:00"
        }
    ]
}
```

<br>

### Table Structure
---

![image](https://github.com/Reinhardjs/mezink-technical-test/assets/7758970/5c5d9279-f0c6-4f1a-a70b-40cdf8ac90c4)
