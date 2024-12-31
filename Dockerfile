FROM golang:1.23

WORKDIR /app

COPY . .

RUN go build -o api main.go

CMD [ "./api" ]

EXPOSE 8080