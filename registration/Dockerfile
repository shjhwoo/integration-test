FROM golang:1.20

WORKDIR /registration

COPY . .

RUN go mod download

RUN go mod tidy

CMD ["go", "run", "./cmd"]

EXPOSE 8081