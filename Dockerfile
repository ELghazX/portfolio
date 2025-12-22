# Frontend
FROM oven/bun:alpine AS frontend-builder
WORKDIR /app
COPY package.json bun.lock ./
RUN bun install
COPY . .
RUN bun x tailwindcss -i static/assets/css/app.css -o static/assets/dist/tailwind.css --minify

# Go Builder
FROM golang:1.25-alpine AS go-builder
WORKDIR /app

RUN go install github.com/a-h/templ/cmd/templ@latest
RUN wget -qO- https://github.com/sqlc-dev/sqlc/releases/download/v1.27.0/sqlc_1.27.0_linux_amd64.tar.gz | tar -xz -C /usr/local/bin

COPY go.mod go.sum ./
RUN go mod download
COPY . .

RUN sqlc generate
RUN templ generate

COPY --from=frontend-builder /app/static/assets/dist/tailwind.css ./static/assets/dist/tailwind.css

RUN CGO_ENABLED=0 GOOS=linux go build -o /app/main ./cmd/api/main.go

# Runner 
FROM alpine:latest
RUN apk add --no-cache ca-certificates tzdata
WORKDIR /app

COPY --from=go-builder /app/main .
COPY --from=go-builder /app/static ./static

EXPOSE 8080
CMD ["./main"]
