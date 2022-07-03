FROM golang:go 1.18
ENV ROOT=/go/src/app
RUN mkdir ${ROOT}
WORKDIR ${ROOT}

COPY go.mod go.sum ./
COPY . .
RUN go get

RUN CGO_ENABLED=0 GOOS=linux go build -o $ROOT/binary
EXPOSE 3000
CMD ["/go/src/app/binary"]
