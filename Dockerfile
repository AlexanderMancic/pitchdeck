# Build stage
FROM golang:1.24.4-alpine AS builder
WORKDIR /app
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o /pitchdeck

# Runtime stage
FROM alpine:3.18
WORKDIR /app
COPY --from=builder /pitchdeck /app/pitchdeck
COPY --from=builder /app/templates /app/templates
COPY --from=builder /app/static /app/static
EXPOSE 8080
CMD ["/app/pitchdeck"]
