FROM golang:1.19.3-alpine
RUN apk update
RUN apk add alpine-sdk
COPY ./go-app /app/go
WORKDIR /app/go
RUN go build -o main .
CMD [ "/app/go/main" ]
