FROM golang:1.19-alpine
WORKDIR /usr/src/jekabot
COPY ./app .
RUN go mod tidy
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o jekabot .

FROM alpine:latest  
RUN apk --no-cache add ca-certificates
WORKDIR /root/
COPY --from=0 /usr/src/jekabot ./
EXPOSE 8000
CMD [ "jekabot" ]
