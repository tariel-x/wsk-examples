# Standalone hello with action loop and subpackage

1. Create `exec` from `main.go`: `cp main.go exec`.
2. Make archive with action: `zip action.zip exec handler/*`.
3. Build `wsk -i action create standalone-multiple action.zip --web true --kind go:default`.
4. Check action `wsk -i action get /guest/standalone-multiple --url`.
5. Create api endpoint `wsk -i api create /standalone-multiple get standalone-multiple --response-type json` and get endpoint url.
6. Call url that looks like `http://HOST:9090/api/23bc46b1-71f6-4ed5-8c54-816aa4f8c502/multiple?name=vasya`.
