FROM golang:1.21-alpine as builder

WORKDIR /app

COPY . .

RUN go mod download && go mod tidy && go mod verify

RUN go build -v -o /app/binary ./cmd/main.go

FROM alpine:3.14

COPY --from=builder /app/config.yaml /config.yaml
COPY --from=builder /app/db/migrations /db/migrations
COPY --from=builder /app/binary /app/binary

EXPOSE 5050
ENV JWT_SECRET="@secret$" POSTGRES_USER=postgres POSTGRES_PASSWORD=postgres DB_HOST=database
ENTRYPOINT /app/binary