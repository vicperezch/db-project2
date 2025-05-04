FROM golang:1.24

WORKDIR /app/src

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN go build -v -o /app/main .

EXPOSE 8080

CMD ["/app/main"]
