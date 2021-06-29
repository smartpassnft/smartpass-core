FROM golang:latest as builder

RUN mkdir /app
COPY . /app/
WORKDIR /app
EXPOSE 8000/tcp
RUN go build -o server .

CMD ["/app/server"]
