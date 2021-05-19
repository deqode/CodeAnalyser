
get-linter:
	go get github.com/mgechev/revive

run-linter:
	revive -formatter friendly ./...

get-fmt:
	go get fmt

run-fmt:
	go fmt ./...


tidy:
	go mod tidy
	#go mod vendor

proto-buf:
	buf lint
	buf generate

# Run all the linters
lint: get-linter get-fmt run-fmt run-linter

#Get the test coverage
test:
	go test -cover ./...
	mkdir -p cover
	go test -coverprofile=./cover/coverage.out ./...
	go tool cover -html=./cover/coverage.out

