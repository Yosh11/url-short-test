# URL Shortener v2 (*the task*)

## Routes

 ***

### /POST

- **localhost:8080/urls/set**

#### *Request in body*

```json
{
    "url": "https://yandex.ru/"
}
```

#### *Response*

```json
{
    "long":  "https://yandex.ru/",
    "short": "http://localhost:8080/urls/AQwwD2z"
}
```

***

### /GET

- **localhost:8080/urls/info/:hash**

#### *Response*

```json
{
    "Id": "619b7d1b525c86dea2ecb00f",
    "created-at": "2021-11-22T11:20:59.505Z",
    "updated_at": "2021-11-22T11:21:07.612Z",
    "deleted_at": null,
    "hash": "AQwwD2z",
    "url": "https://yandex.ru/",
    "count": 1,
    "access": true,
    "code": 0
}
```

 ***

### /GET

- **localhost:8080/urls/:hash**

#### *Redirect*

http://localhost:8080/urls/AQwwD2z ---> https://yandex.ru/

 ***

### /DELETE

- **localhost:8080/urls/:hash**

#### *Response*

```json
{
    "message": "url removed"
}
```

 ***

## Startup

```sh
$ cd ./
$ docker-compose up --build
```
 ***