swagger:
	@swagger generate spec -o ./swagger.yml --scan-models
	@swagger serve ./swagger.yml

run:
	CGO_ENABLED=0 GOARCH=amd64 GOOS=linux go build -o ./bin -ldflags="-w -s" ./...
	./bin/backend

debug:
	CGO_ENABLED=0 GOARCH=amd64 GOOS=linux go build -o ./bin -gcflags="all=-N -l" ./...
	@dlv --listen=:2345 --headless=true --api-version=2 exec ./bin/backend