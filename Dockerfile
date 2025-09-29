FROM golang:1.25

WORKDIR /app

COPY . .

RUN go build -o /app/generator ./cmd/http

EXPOSE 4000

CMD ["/app/generator"]
