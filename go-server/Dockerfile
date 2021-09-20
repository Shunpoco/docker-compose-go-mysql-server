FROM golang:1.17

WORKDIR /app

COPY . .

RUN ls

RUN go get \
  && go build -o app

CMD ["./app"]

EXPOSE 5000
