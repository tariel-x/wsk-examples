# Multiple files in the main package

1. Build `wsk -i action create assign assign.go --web true`.
2. Check action `wsk -i action get /guest/assign --url`.
3. Create api endpoint `wsk -i api create /assign get assign --response-type json` and get endpoint url.
4. Call url that looks like `http://HOST:9090/api/23bc46b1-71f6-4ed5-8c54-816aa4f8c502/assign?language=go`.
