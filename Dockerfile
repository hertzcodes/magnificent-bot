FROM golang:1.23 as builder

WORKDIR /github.com/hertzCodes/magnificent-bot

RUN CGO_ENABLED=0 go build -o ./bot cmd/main.go

FROM alpine as deploy

RUN apk add --no-cache tzdata
ENV TZ=Asia/Tehran

WORKDIR /github.com/hertzCodes/magnificent-bot

COPY --from=builder /github.com/hertzCodes/magnificent-bot ./bot

CMD [ "./bot" , "--config" , "/etc/config.json"]