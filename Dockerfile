FROM golang

WORKDIR /server
COPY . /server

CMD ["go", "run", "main.go"]
