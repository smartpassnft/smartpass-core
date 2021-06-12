FROM golang:latest

RUN mkdir /app

COPY ./src /app/

WORKDIR /app

EXPOSE 8000/tcp

RUN go build -o server .

CMD ["/app/server"]
