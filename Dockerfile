FROM  golang:1.21.3 as builder

WORKDIR /app/cmd

COPY go.mod ./

RUN go mod tidy

COPY . .

RUN go build -o main .

FROM alpine:latest

RUN apk --no-cache add ca-certificates

WORKDIR /root/

COPY --from=builder /app/cmd/main .

EXPOSE 8080

CMD [ "./main" ]