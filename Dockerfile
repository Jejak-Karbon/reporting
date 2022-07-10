FROM golang:alpine

WORKDIR /apps/reporting
COPY . /apps/reporting

RUN go build -o main .

CMD ["/apps/reporting/main"]