FROM golang:1.22-alpine

WORKDIR /app
COPY go.mod go.sum ./

RUN go mod download

COPY . .

ENV GOOS=linux CGO_ENABLED=0

RUN go build -o /loslogger ./cmd/main.go

EXPOSE ${LOSLOGGER_API_PORT}

CMD ["/loslogger"]
