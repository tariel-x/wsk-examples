# Deploy yaml

## Setup

Add

```
APIGW_ACCESS_TOKEN=KEY
```

to `~/.wskprops` where `KEY` is same to `AUTH` from the same file.

## Deploy 

1. `wskdeploy -m deploy.yml`
2. See apis list `wsk -i api list`
3. Call url like `http://10.129.0.29:9090/api/23bc46b1-71f6-4ed5-8c54-816aa4f8c502/assign/reviewer?language=go`
4. Or invoke sequence with ` wsk -i action invoke /guest/assignee_pkg/assign_mr_sequence -p language go -b`
