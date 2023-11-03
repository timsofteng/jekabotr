FROM golang:1.21.0-alpine
WORKDIR ./images 

COPY . .
RUN go build

CMD ["./images"]
