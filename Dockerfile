FROM golang:alpine

WORKDIR /var/www/html/apps/reporting
COPY . /var/www/html/apps/reporting

RUN go build -o main .

CMD ["/var/www/html/apps/reporting/main"]