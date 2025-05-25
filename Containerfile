FROM golang:1.24-alpine

LABEL org.opencontainers.image.source=https://github.com/victorposada/gerrit-wui


WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download
RUN go mod verify
RUN apk add curl

COPY . .

CMD ["go", "run", "cmd/gerrit-wui/main.go"]
