FROM golang

WORKDIR /go/src/github.com/hunterlong/tokenbalance
COPY . .

ENV GETH_SERVER=https://mainnet.infura.io/uy52ECefn575YC1ZaVNO
ENV PORT=8080

RUN go get -d -v ./...
RUN go install

EXPOSE $PORT

ENTRYPOINT tokenbalance start --geth=$GETH_SERVER --port $PORT --ip 0.0.0.0