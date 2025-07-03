FROM golang:latest

RUN mkdir /app

ADD . /app/

WORKDIR /app

RUN go get gopkg.in/telebot.v4
RUN go get github.com/redis/go-redis/v9
RUN go get github.com/joho/godotenv

RUN go run bot.go