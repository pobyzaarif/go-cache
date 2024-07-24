# go-cache
Golang cache library/adapter with TTL implementation, key:value store using memory (Memcached), and Redis.
- You can use memory cache when you run the app in a single machine/local env, temporary mode
- Use Redis when you need persistent cache data

## Implementation
Just take a look at the test file ( *_test.go )

## Run test
```
go test -count=1 -coverprofile=coverage.out -coverpkg=./... && go tool cover -html=coverage.out -o cover.html
```
