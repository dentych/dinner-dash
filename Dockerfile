# Build backend
FROM golang:alpine

# Install prerequisites
RUN apk add git

COPY . /build

WORKDIR /build

RUN go build -o dinner-dash cmd/main.go

# Create dinner-dash backend image
FROM alpine

WORKDIR /app

COPY --from=0 /build/dinner-dash /app
COPY ./migrations /app/migrations

EXPOSE 8080

CMD ["./dinner-dash"]
