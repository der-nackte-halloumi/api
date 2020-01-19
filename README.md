# api

The API.

[![Build Status](https://travis-ci.com/der-nackte-halloumi/api.svg?branch=master)](https://travis-ci.com/der-nackte-halloumi/api)

## File structure

- `cmd`: entrypoints to the application (e.g. REST / RPC / CLI / ...)
- `pkg`: reusable stuff which is portable between projects, "utils" and these things
- `scripts`: startup scripts / migrations and this shizzle
- `datastore`: everything related to the database access, the concrete implementation of the repository interfaces
- `usecase`: usescase the user wants to be accomplish (these are basically services). A single usecase holds the interface for a service and the repository and the implementation of the service as well
- `interface`: the actual interface to the application, here it will be a REST APInterface
- `domain`: bridge between datastore and interface (here we perform validation and these things)
