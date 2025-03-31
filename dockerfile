FROM golang:1.24-alpine

ENV GIN_MODE=release
WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY . . 
RUN go build -o /task-manager ./cmd

EXPOSE 8080

CMD [ "/task-manager" ]

