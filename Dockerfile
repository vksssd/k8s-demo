FROM  golang:1.21.3 as builder

WORKDIR /app

COPY go.mod ./

RUN go mod tidy

COPY . .

RUN echo "checking file in cmd"

RUN ls 

RUN CGO_ENABLED=0 GOOS=linux go build -o ./cmd/main ./cmd

FROM alpine:latest

RUN apk --no-cache add ca-certificates

WORKDIR /root/

COPY --from=builder /app/cmd/main .

EXPOSE 8080

CMD [ "./main" ]