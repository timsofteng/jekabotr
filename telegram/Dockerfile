FROM golang:1.21.0-alpine

ARG GITHUB_ACCESS_TOKEN
ARG GOPRIVATE

RUN apk add --no-cache git
CMD ["git","--version"]

RUN git config --global \
url."https://$GITHUB_ACCESS_TOKEN:@github.com/".insteadOf \
"https://github.com/"

WORKDIR ./telegram 

COPY go.mod go.sum .
RUN go mod download

COPY . .
RUN go build

CMD ["./telegram"]
