FROM golang:1.17-alpine
WORKDIR /

ENV GOOS linux

COPY go.mod ./
#COPY go.sum ./

RUN go mod download

COPY . .
RUN go build -o zedis

EXPOSE ${ZEDIS_PORT}
CMD ./zedis