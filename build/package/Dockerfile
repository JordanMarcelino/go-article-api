FROM golang:1.21-alpine as builder

WORKDIR /app

COPY . .

#RUN go install -tags'mysql' github.com/golang-migrate/migrate/v4/cmd/migrate@latest
RUN go mod download && go mod verify
RUN go build -v -o /app/binary ./cmd/main.go

FROM alpine:3.14

COPY --from=builder /app/config.yaml /config.yaml
COPY --from=builder /app/binary /app/binary

EXPOSE 5050
ENTRYPOINT /app/binary