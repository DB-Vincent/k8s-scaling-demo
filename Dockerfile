# Build stage
FROM golang:1.23-alpine AS build
WORKDIR /app
COPY . .
RUN go build -o demo .

# Runtime stage
FROM alpine:3.20
WORKDIR /app
COPY --from=build /app/demo .
COPY --from=build /app/frontend frontend/
CMD ["/app/demo"]
