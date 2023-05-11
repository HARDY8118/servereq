# servereq
Simple Go server to print http request

## Build

```sh
go build -o servereq main.go
```

## Start

```sh
./servereq [PORT]
```

Make sure the port provided is free for use.

### Example

```sh
./servereq 8000
```

which will output

```sh
Server starting on PORT: 8000
```

## Output format

On request to running server similar to
```sh
curl -X POST -H 'Content-Type: application/json' --data '{"id": 1}' http://localhost:8000/users
```
similar output is generated
```sh
HTTP/1.1 POST localhost:8000/users
Content-Type: [application/json]
Content-Length: [9]
User-Agent: [curl/7.81.0]
Accept: [*/*]
{"id": 1}
```

## Reserved path

- `/version` will respond with the version number to the request.
`curl -X POST -H 'Content-Type: application/json' --data '{"id": 1}' http://localhost:8000/users`
will respond with something similar to `Version: 0.0.1`