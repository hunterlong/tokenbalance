FROM alpine:latest
MAINTAINER "Hunter Long (https://github.com/hunterlong)"

ARG VERSION
RUN apk --no-cache add libstdc++ ca-certificates curl jq

# make build-alpine
ADD ./build/alpine-linux-amd64 /usr/local/bin/tokenbalance

WORKDIR /app
ENV GETH_SERVER=https://eth.coinapp.io
ENV PORT=8080
EXPOSE $PORT
HEALTHCHECK --interval=30s --timeout=10s --retries=5 CMD curl -s "http://localhost:$PORT/health" | jq -r -e ".online==true"
ENTRYPOINT tokenbalance start --geth $GETH_SERVER --port $PORT --ip 0.0.0.0