NAME=ekolo
SWAG_GFILE=app/app.go
SWAG_OUT=app/docs

test:
	go test -v -failfast -count=1 -cover -covermode=count -coverprofile=coverage.out ./...
	go tool cover -func coverage.out

deps:
	go install github.com/swaggo/swag/cmd/swag@latest
	go get -u github.com/swaggo/files

swag.gen:
	swag init --parseInternal -g ${SWAG_GFILE} -output ${SWAG_OUT}

build: swag.gen
	go build -o bin/${NAME} cmd/main.go

run: build
	./bin/${NAME}
