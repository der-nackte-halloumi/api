# api

The API.

[![Build Status](https://travis-ci.com/der-nackte-halloumi/api.svg?branch=master)](https://travis-ci.com/der-nackte-halloumi/api)

## Installation

Ensure `go`, `Docker` and if needed `pgcli` are installed and up-to-date. Copy `.env.dist` to `.env` and update all empty values.

## Development

A series of `make` commands support the local development isolated from the other pieces:

```sh
# create a local database container
make db_create
# re-/start the local database container
make db_start
# stop the local database container
make db_stop
# remove the local database container including a possible volume
make db_clean
# calls pgcli to connect to the local database container, requires pgcli to be in path
make db_cli
# runs all migrations in the local database container
make migrations
# runs the acutal API
make run
```

## File structure

- `cmd`: entrypoints to the application (e.g. REST / RPC / CLI / ...)
- `pkg`: reusable stuff which is portable between projects, "utils" and these things
- `scripts`: startup scripts / migrations and this shizzle
- `datastore`: everything related to the database access, the concrete implementation of the repository interfaces
- `usecase`: usescase the user wants to be accomplish (these are basically services). A single usecase holds the interface for a service and the repository and the implementation of the service as well
- `interface`: the actual interface to the application, here it will be a REST APInterface
- `domain`: bridge between datastore and interface (here we perform validation and these things)
