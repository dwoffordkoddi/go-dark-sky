FROM golang:1.14 as intermediate

ARG ACCESS_TOKEN_USR="nothing"
ARG ACCESS_TOKEN_PWD="nothing"
WORKDIR /tmp
RUN apt-get update
RUN apt-get install git gcc
RUN wget "https://www.apache.org/dyn/mirrors/mirrors.cgi?action=download&filename=pulsar/pulsar-2.5.0/DEB/apache-pulsar-client.deb"\
    -O pulsar.deb
RUN wget "https://www.apache.org/dyn/mirrors/mirrors.cgi?action=download&filename=pulsar/pulsar-2.5.0/DEB/apache-pulsar-client-dev.deb"\
    -O pulsar-dev.deb
RUN apt install ./pulsar.deb
RUN apt install ./pulsar-dev.deb

RUN printf "machine github.com\n\
    login ${ACCESS_TOKEN_USR}\n\
    password ${ACCESS_TOKEN_PWD}\n\
    \n\
    machine api.github.com\n\
    login ${ACCESS_TOKEN_USR}\n\
    password ${ACCESS_TOKEN_PWD}\n"\
    >> /root/.netrc
RUN chmod 600 /root/.netrc

WORKDIR /go/src/app
COPY . .

RUN go get -d -v ./...
RUN go build

FROM golang:1.14 AS final
COPY --from=intermediate /go/ /go
COPY --from=intermediate /usr/lib/ /usr/lib/
WORKDIR /go/src/app

EXPOSE 8080/tcp
ENV ENV Dev

CMD ["./koddi-framework-starter"]