FROM golang:1.16

WORKDIR /go/src/

ENV PATH="/go/bin:${PATH}"

COPY . .


EXPOSE 8000

RUN GOOS=linux go build main.go

ENTRYPOINT [ "./ifc3-desafio-go" ]