FROM golang:1.15 as builder

WORKDIR /cli

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

RUN ls
RUN GOOS=linux GOARCH=amd64 CGO_ENABLED=0 make build-simple

FROM alpine:3.13

WORKDIR /karavel

COPY --from=builder /cli/bin/karavel /usr/local/bin/karavel

ENTRYPOINT [ "/usr/local/bin/karavel" ]
CMD [ "help" ]
