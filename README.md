# URL Shortener (*the task*)

## Routes

### /POST

- **localhost:8080/set**

#### *Request in body*

```json
{
    url: "https://google.com"
}
```

#### *Respose*

```json
{
    long:  "https://google.com",
    short: "http://localhost:8080/nYpDEvR"
}
```

***

### /GET

- **localhost:8080/info/:hash**

#### *Respose*

```json
{
    "ID": 1,
    "CreatedAt": "2021-01-24T17:35:49.2282227+03:00",
    "UpdatedAt": "2021-01-25T11:44:51.3115739+03:00",
    "DeletedAt": null,
    "hash": "nYpDEvR",
    "url": "https://google.com",
    "count": 2,
    "access": true,
    "code": 200
}
```

 ***

 ### /GET

- **localhost:8080/:hash**

#### *Redirect*

http://localhost:8080/nYpDEvR ---> https://google.com

 ***

## Build

### MS SQL Server

```sh
$ cd ./url-short-test
$ docker-compose up

```

### API Server

```sh
$ cd ./url-short-test/cmd/url-short
$ ./main.exe
```
