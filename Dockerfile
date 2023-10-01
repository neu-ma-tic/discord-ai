# syntax=docker/dockerfile:1

# Build the application from source
FROM golang AS build-stage

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY *.go ./

RUN CGO_ENABLED=0 GOOS=linux go build -o /dgo-api

# Deploy the application binary into a lean image
FROM alpine AS build-release-stage

WORKDIR /

COPY --from=build-stage /dgo-api /dgo-api

ENTRYPOINT ["/dgo-api"]