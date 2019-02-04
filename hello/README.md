# Hello

1. Create action `wsk -i action create helloGo hello.go --web true`.
2. Check action `wsk -i action get /guest/helloGo --url`.
3. Create api endpoint `wsk -i api create /hello /world get helloGo --response-type json` and get endpoint url.
4. Call url that looks like `http://HOST:9090/api/23bc46b1-71f6-4ed5-8c54-816aa4f8c502/hello/world?name=UserName`.