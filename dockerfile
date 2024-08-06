FROM golang:1.22.5

ENV config=docker

WORKDIR /app

COPY ./ /app

RUN go mod download

COPY ./wait-for-it.sh /usr/local/bin/wait-for-it.sh
RUN chmod +x /usr/local/bin/wait-for-it.sh

RUN go build -o main ./cmd/api

EXPOSE 3000

CMD [ "wait-for-it.sh","db:5432","--","./main" ]