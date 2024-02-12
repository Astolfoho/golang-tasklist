FROM golang:1.21.6-alpine

WORKDIR /app

COPY go.mod ./
RUN go mod download

COPY *.go ./

RUN go get
RUN go build -o /dist

EXPOSE 8080

CMD [ "/dist" ]