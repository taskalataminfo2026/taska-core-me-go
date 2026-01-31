# ====== Etapa de compilación ======
FROM golang:1.22-alpine AS builder

WORKDIR /app
RUN apk add --no-cache git ca-certificates

COPY go.mod go.sum ./

ARG GITHUB_TOKEN
RUN git config --global url."https://${GITHUB_TOKEN}@github.com/".insteadOf "git@github.com:"

ENV GOPRIVATE=github.com/taskalataminfo2026

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o taska-auth ./cmd/api/main.go


# ====== Etapa de producción ======
FROM alpine:3.19

WORKDIR /root
RUN apk --no-cache add ca-certificates

COPY --from=builder /app/taska-auth .

EXPOSE 8080
CMD ["./taska-auth"]
