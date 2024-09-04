## Simple HTTP server on Golang.

The server supports only POST request, with obligated "id" (int) and "cmd" (string) keys. The response will be the same but with one additional field - "result". The result is random color. It will be the same as in previous response if the previous request had the same id as current. The server supports simple local persistent state.

Example:
```
>>> curl localhost:8080/process -d '{"id": 123, "cmd": "test"}'
{"cmd":"test","id":123,"result":"green"}

>>> curl localhost:8080/process -d '{"id": 124, "cmd": "test"}'
{"cmd":"test","id":123,"result":"red"}

>>> curl localhost:8080/process -d '{"id": 124, "cmd": "TEST TEST"}'
{"cmd":"TEST TEST","id":123,"result":"red"}
```
