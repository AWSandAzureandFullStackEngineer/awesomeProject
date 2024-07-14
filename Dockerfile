# Build stage
FROM golang:1.9rc2-alpine AS build
WORKDIR /app
COPY go.mod ./

COPY . .
RUN CGO_ENABLED=0 go build -o app

# Final stage
FROM alpine:latest
WORKDIR /app
COPY --from=build /app/app .
EXPOSE 8080
CMD ["./app"]