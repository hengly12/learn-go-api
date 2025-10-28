# Use Go 1.25 as the base image (matches your go.mod)
FROM golang:1.25.3 AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o main .

# ---- Runtime Stage ----
FROM gcr.io/distroless/base-debian12
WORKDIR /app
COPY --from=builder /app/main .

EXPOSE 8080
CMD ["./main"]
