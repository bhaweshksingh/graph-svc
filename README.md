# graph service service


## Pre requisites

- Docker
- Golang v1.14+
 
 
## Running App 

1. Build and run the app container.

`make app`

2. Inspect logs using docker 

`docker logs graph-svc-go -f`

## Verifying the Functionality

Getting chart data key Request
```shell script

curl -X POST 'http://localhost:8080/chart-data' \
--header 'Content-Type: application/json'
```

