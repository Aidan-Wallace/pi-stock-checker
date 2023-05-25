FROM golang:1.19 AS build-stage

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY main.go ./main.go
COPY pkg ./pkg

RUN go build -v -o /stock-checker ./main.go

FROM build-stage AS run-test-stage
RUN go test -v ./...
WORKDIR /
COPY config ./config
COPY --from=build-stage /stock-checker /stock-checker
ENTRYPOINT ["/stock-checker"]
