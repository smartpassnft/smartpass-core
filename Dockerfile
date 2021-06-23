FROM golang:latest as builder

RUN mkdir /app
COPY . /app/
WORKDIR /app
EXPOSE 8000/tcp
RUN go build -o server .

FROM debian:bullseye-slim

WORKDIR /tmp
COPY --from=builder /app/server /app/config /app/log .
CMD ["/tmp/server"]
