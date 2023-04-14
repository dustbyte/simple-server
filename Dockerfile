FROM golang

WORKDIR /server
COPY . /server

RUN go build .

CMD /server/simple-server
