FROM golang:1.21.4-alpine3.18 AS builder

WORKDIR /go/src/

COPY . .

RUN go mod download

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o app .

FROM gcr.io/distroless/static-debian12

COPY --from=builder /go/src/app /app

CMD ["/app"]

EXPOSE 8000