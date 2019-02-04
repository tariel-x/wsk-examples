# Standalone hello

Run at local: `go run main.go '{"name": "username"}'`

1. Build `GOOS=linux GOARCH=amd64 go build -o exec main.go`
2. Pack `zip action.zip exec`
3. Deploy `wsk -i action create helloGoStandalone action.zip --native --web true`
4. Create api endpoint `wsk -i api create /hello /world get helloGoStandalone --response-type json` and get endpoint url.
5. Call url that looks like `http://HOST:9090/api/23bc46b1-71f6-4ed5-8c54-816aa4f8c502/hello/world?name=UserName`.
