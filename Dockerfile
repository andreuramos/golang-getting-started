FROM golang:latest AS builder

WORKDIR /go/src/app
COPY . .

RUN go build -o /go/bin/app cmd/main.go

FROM scratch
COPY --from=builder /app /app
ENTRYPOINT ["app"]

