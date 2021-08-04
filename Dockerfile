FROM golang

RUN apt-get update && apt-get -q install -y \
    curl \
    dnsutils \
    gettext

WORKDIR /go/delivery

COPY . . 

RUN go build .

CMD ["./webserver"]

