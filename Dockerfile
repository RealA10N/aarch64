FROM golang:1.23

WORKDIR /aarch64codegen

COPY . .

RUN go mod download

CMD ["go", "test", "-v", "./..."]