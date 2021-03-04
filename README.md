# messari-metrics

## Key takeaways

1. Used Goa (https://github.com/goadesign/goa) to generate design-based service, see: `design/design.go`
    - Took some extra time on this as I've been wanting to learn Goa
2. Implemented Messari API client with two endpoints, see: `messari/client.go`
3. Implemented Metrics Service API, see: `endpoints.go`
4. Implemented an AssetCache for reduced aggregate request times, see: `asset_cache.go`
    - I would move this to a Redis-backed cache in a real world scenario
5. Wrote some tests, but would write a lot more in the real world

## Setup

1. Check out project: `git clone git@github.com:jordanfowler/messari-metrics.git`
2. From project root: `go mod download`
3. Export your Messari API Key: `export MESSARI_API_KEY=...` 
4. From project root: `go run .`

## API Requests

1. Single asset metrics (replace "{slug}" with any asset on Messari): `curl localhost:8000/asset/{slug}`
2. Aggregate asset metrics (query params optional): `curl localhost:8000/aggregate[?tags=...|sector=...]`

## Caveats

- Aggregate endpoint will not show full results until all asset tags & sectors have been cached, time required: ~ 15 - 25s 
- Until removing Messari API rate limiting, aggregate endpoint may become unresponsive