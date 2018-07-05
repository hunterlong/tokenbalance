FROM alpine:latest

ENV VERSION=1.52

RUN apk --no-cache add libstdc++ ca-certificates
RUN wget -q https://github.com/hunterlong/tokenbalance/releases/download/v$VERSION/tokenbalance-linux-alpine.tar.gz && \
      tar -xvzf tokenbalance-linux-alpine.tar.gz && \
      chmod +x tokenbalance && \
      mv tokenbalance /usr/local/bin/tokenbalance

WORKDIR /app
ENV GETH_SERVER=https://mainnet.infura.io/uy52ECefn575YC1ZaVNO
ENV PORT=8080

EXPOSE $PORT

ENTRYPOINT tokenbalance start --geth $GETH_SERVER --port $PORT --ip 0.0.0.0