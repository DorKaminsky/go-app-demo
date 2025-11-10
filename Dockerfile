FROM golang:1.21-alpine AS builder

COPY . .

RUN go mod download

RUN go build -o go-app-demo .

FROM golang:1.21-alpine


COPY --from=builder /app/go-app-demo .
COPY --from=builder /app/VERSION .

EXPOSE 8080

CMD ["./go-app-demo"]
