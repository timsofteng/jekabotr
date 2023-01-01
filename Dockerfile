FROM golang:1.19-alpine as builder
LABEL stage=gobuilder
WORKDIR /usr/src/jekabot
COPY ./app .
RUN go mod tidy
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o jekabot .

FROM alpine:latest  
RUN apk --no-cache add ca-certificates
WORKDIR /root/
COPY --from=builder /usr/src/jekabot/jekabot .
COPY --from=builder /usr/src/jekabot/config.yaml .
EXPOSE 8000
CMD [ "./jekabot" ]
