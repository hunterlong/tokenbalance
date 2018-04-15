FROM golang

WORKDIR /go/src/github.com/hunterlong/tokenbalance
COPY . .

ENV GETH_SERVER=https://mainnet.infura.io/uy52ECefn575YC1ZaVNO
ENV PORT=8080
ENV GOPATH=/go
ENV PATH=/go/bin:/usr/local/go/bin:/usr/local/sbin:/usr/local/bin:/usr/sbin:/usr/bin:/sbin:/bin

RUN go get -d -v ./...
RUN go build .

EXPOSE $PORT

ENTRYPOINT ./tokenbalance start --geth=$GETH_SERVER --port $PORT --ip 0.0.0.0