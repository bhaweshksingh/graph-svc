# Graph service

This solution consumes and analyses the CSV to return graph data in API.
The linear trendline based on the daily count is calculated.

The response to the chart data endpoint includes:
- Chart Data Points
- Trendline start and end
- Trendline points

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


Response with limited Data

```shell
{
  "data": {
    "DataPoints": [
      {
        "Day": "01 Jan 2021",
        "VisitCount": 4
      },
      {
        "Day": "02 Jan 2021",
        "VisitCount": 4
      },
      {
        "Day": "03 Jan 2021",
        "VisitCount": 3
      },
      {
        "Day": "25 Jan 2022",
        "VisitCount": 1
      }
    ],
    "Trendline": {
      "Start": {
        "Day": "01 Jan 2021",
        "VisitCount": 3.0402978429640433
      },
      "End": {
        "Day": "25 Jan 2022",
        "VisitCount": 2.5790718493133755
      }
    },
    "TrendlinePoints": [
      {
        "Day": "01 Jan 2021",
        "VisitCount": 3.0402978429640433
      },
      {
        "Day": "02 Jan 2021",
        "VisitCount": 3.0391121720292085
      },
      {
        "Day": "03 Jan 2021",
        "VisitCount": 3.037926501094374
      },
      {
        "Day": "25 Jan 2022",
        "VisitCount": 2.5790718493133755
      }
    ]
  },
  "success": true
}
```
