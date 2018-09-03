# Build backend
FROM golang:alpine

# Install prerequisites
RUN apk add git

COPY . /build

WORKDIR /build

RUN go build -o dinner-dash main.go

# Create dinner-dash backend image
FROM alpine

WORKDIR /app

COPY --from=0 /build/dinner-dash /app

EXPOSE 8080

CMD ["./dinner-dash"]
