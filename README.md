# graphql-preact-starter

A starter configuration for a self-contained GraphQL server/client Docker
image.

Uses Golang, Sqlite3 and GraphQL on the backend and Preact and GraphQL on
the frontend.

## Routes

In development, docker-compose will expose the server at <localhost:4510>. The
Production image will expose port `:8080` which may be bound as desired.

- `/` and any other route not stated below will serve the Preact UI.
- `/graphql` is the GraphQL query endpoint.
- `/graphiql` is the GraphiQL interactive UI (only in development).

## Server

The server uses [graphql-go] with inspiration from [this example][gql-example].
It mounts [`client`](#client)'s `.build` directory to `/public` and serves
those files at `/static/`, as well as serving `index.html` at all other
unspecified routes. It then serves the GraphQL query endpoint at `/graphql`.
In dev, it will serve the GraphiQL UI at `/graphiql`.

## Client

The client is a simple HTML web page that serves a Preact application to
consume [`server`](#server)'s GraphQL query endpoint.

## Scripts

This repo adheres loosely to the [Scripts to Rule Them All] convention:

- `script/setup` runs `docker-compose build` for development. There are two
  services in `docker-compose`, one for each Client and Server.
- `script/server` will rebuild, then serve with `docker-compose up`. Both
  services will rebuild and reload automatically. At present there is no HMR
  on the front end so you'll have to reload a page to see changes.
- `script/build` builds a single Docker image for Production using the
  `Dockerfile` at the root of the repo and tag it `:latest`.
- `script/test` runs tests and captures coverage. Run `script/test -h` for
  more information on usgae.

Since this repo also contains, in a way, two distinct applications, there
are also some helper scripts to make life a little easier:

- `script/npm` passes arguments to `npm` in the `client` directory, e.g.
  `script/npm run codegen`
- `script/go` passes arguments to `go` in the `server` directory, e.g.
  `script/go vet ./...`

[graphql-go]: https://github.com/graph-gophers/graphql-go
[gql-example]: https://github.com/tonyghita/graphql-go-example
[scripts to rule them all]: https://github.com/github/scripts-to-rule-them-all
