# messari-metrics

## Setup

1. Check out project: `git clone git@github.com:jordanfowler/messari-metrics.git`
2. From project root: `go mod download`
3. Export your Messari API Key: `export MESSARI_API_KEY=...` 
4. From project root: `go run .`

## API Requests

1. Single asset metrics (replace "{slug}" with any asset on Messari): `curl localhost:8000/asset/{slug}`
2. Aggregate asset metrics (query params optional): `curl localhost:8000/aggregate[?tags=...|sector=...]`