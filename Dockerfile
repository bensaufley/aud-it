FROM golang:1.16-alpine as go-builder

WORKDIR /go/src/github.com/bensaufley/aud-it
COPY ./server/go.mod ./server/go.sum ./
RUN go mod download

RUN apk add --no-cache \
        gcc \
        libc6-compat \
        musl-dev

COPY ./server/ /go/src/github.com/bensaufley/aud-it/

RUN CGO_ENABLED=1 GOARCH=amd64 GOOS=linux go build -a -installsuffix 'static' -o serve ./cmd/serve

FROM node:16.5.0 as node-builder

WORKDIR /tmp
COPY ./client/package*.json ./
RUN npm install

WORKDIR /usr/src/aud-it
RUN mv /tmp/package*.json /tmp/node_modules ./

COPY ./client/ /usr/src/aud-it/
COPY ./server/internal/schema/graphql/ /usr/src/aud-it/src/graphql/schema/

ENV NODE_ENV=production

RUN npm run build

RUN npm prune --production

FROM alpine:3.14

RUN apk add --no-cache \
        gcc \
        libc6-compat

WORKDIR /app
COPY --from=go-builder /go/src/github.com/bensaufley/aud-it/serve .
COPY --from=node-builder /usr/src/aud-it/.build /public
COPY ./server/migrations/ ./migrations/

VOLUME [ "/storage" ]

EXPOSE 8080

CMD [ "./serve" ]
