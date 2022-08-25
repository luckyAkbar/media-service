FROM golang:1.18-alpine

WORKDIR /app

RUN mkdir src
RUN mkdir bin
RUN mkdir image_storage
RUN mkdir video_storage

COPY go.mod .
COPY go.sum .

RUN go mod download
RUN go mod tidy

WORKDIR /app/src

COPY . .

RUN go build -o /app/bin/main main.go

WORKDIR /app

RUN rm -r /app/src

EXPOSE ${SERVER_PORT}

CMD ["/app/bin/main", "server"]