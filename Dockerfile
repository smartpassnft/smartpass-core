FROM golang:latest

RUN mkdir /app

COPY . /app/

WORKDIR /app

EXPOSE 8000/tcp

# RUN mv smartpass /usr/local/go/src

RUN go build -o server .

CMD ["/app/server"]
