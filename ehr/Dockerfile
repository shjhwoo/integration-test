FROM golang:1.20

WORKDIR /ehr

COPY . .

RUN go mod download

RUN go mod tidy

CMD ["go", "run", "./cmd"]

EXPOSE 8080
