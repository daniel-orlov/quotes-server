# quotes-server

## Requirements

+ [Go 1.19+](https://go.dev/dl/) installed (to run tests, start server or client without Docker)
+ [Docker](https://docs.docker.com/engine/install/) installed (to run docker)
+ [Docker Compose](https://docs.docker.com/compose/install/) installed (to run docker-compose)
+ [GNU Make](https://www.gnu.org/software/make/) installed (to run Makefile)
+ [GolangCI-Lint](https://golangci-lint.run/usage/install/) installed (to run linter)

## Getting started

### Start server and client via docker-compose:

```
make start
```

### Run only server:

```
make run-server
```

### Run only client:

```
make run-client
```

### Launch tests:

```
make test
```

## Project structure

> / **quotes-server**
>
> > / **.github**
> > > / **workflows** - GitHub Actions workflows for running CI/CD pipelines
>
> > / **build** - Dockerfiles for building images
>
> > / **cmd** - entrypoints for server and client
>
> > / **configs** - configuration file for server (the client one is in the `/pkg/client` directory)
>
> > / **deploy** - docker-compose file for running server and client
>
> > / **docs** - documentation files
>
> > / **internal** - internal packages
>
> > > / **domain** - domain entities, such as models and services, containing business logic
>
> > > / **storage** - storage layer, containing repositories and database models
>
> > > / **transport** - transport layer, containing handlers and middlewares
>
> > / **pkg** - public packages
>

## Contributing

I used [Conventional Commits](https://www.conventionalcommits.org/en/v1.0.0/) for commit messages.
I also recommend adhering to TDD (Red-Green-Refactor) principle when contributing.

### Linting

To run linter, use the following command:

```
make lint-host
```

### Testing

To run tests, use the following command:

```
make test
```

## More information

For more information, please, refer to /docs directory.