FROM golang

WORKDIR /go/src
ADD api /go/src/minesweeper-API/api
ADD main /go/src/minesweeper-API/main
ADD minesweeper /go/src/minesweeper-API/minesweeper
ADD persistence /go/src/minesweeper-API/persistence
ADD types /go/src/minesweeper-API/types
ADD Gopkg.toml /go/src/minesweeper-API

RUN mkdir /go/src/minesweeper-API/builds

RUN curl https://raw.githubusercontent.com/golang/dep/master/install.sh | sh

WORKDIR /go/src/minesweeper-API

RUN dep ensure

RUN go build -o builds/minesweeper_api main/main.go

EXPOSE 8080

CMD ./builds/minesweeper_api
