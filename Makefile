server-dev:
	gin --port 3000 --appPort 8181 go run main.go

client-dev:
	npm run client-dev-ssl

test:
	go test -coverprofile=core.cover.out -coverpkg=./server/ ./server/...

test-cover:
	go test -coverprofile=core.cover.out -coverpkg=./server/ ./server/...
	cat core.cover.out > coverage.out
	rm *.cover.out

view-cover:
		go tool cover -html=coverage.out
