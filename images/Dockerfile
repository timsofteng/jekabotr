FROM golang:1.21.0-alpine
WORKDIR ./telegram 

COPY . .
RUN go build

CMD ["./telegram"]
