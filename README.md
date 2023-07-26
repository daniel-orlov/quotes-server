# quotes-server

[![Go Report Card](https://goreportcard.com/badge/github.com/daniel-orlov/quotes-server)](https://goreportcard.com/report/github.com/daniel-orlov/quotes-server)
![Go Pipeline](https://github.com/daniel-orlov/quotes-server/actions/workflows/go.yaml/badge.svg)

Quotes server is an example of how I normally build an HTTP server.

It started as a solution for a technical assignment and later turned into a showcase.
Currently I enhance it ad-hoc, see the list of the project issues for the planned improvements and add your own to challenge me.

For more details on the initial technical requirements, see [docs/000_Requirements.md](docs/000_Requirements.md)

Please also consult [docs/003_Potential_questions_and_answers_to_them.md](docs/003_Potential_questions_and_answers_to_them.md) for
additional information on the project.

## Requirements

+ [Go 1.19+](https://go.dev/dl/) installed (to run tests, start server or client without Docker)
+ [Docker](https://docs.docker.com/engine/install/) installed (to run docker)
+ [Docker Compose](https://docs.docker.com/compose/install/) installed (to run docker-compose), v2.0+ is required
+ [GolangCI-Lint](https://golangci-lint.run/usage/install/) installed (to run linter)
+ [GNU Make](https://www.gnu.org/software/make/) installed (to run Makefile)

NB: make commands tested to work both on GNU Make 3.81 (comes with the latest macOS) and 4.4.1 (the latest version at
the moment of writing).

## Getting started

### Environment variables

To avoid overcomplicating testing the solution of the project, I used envconfig with some reasonable defaults.
You can find them in the [config/config.go](config/config.go) file for server and in
the [pkg/client/config.go](pkg/client/config.go) file for client.

If you would like to set them manually, you can do it via environment variables:

### Server

| Name                 | Description                               | Default Value | Possible Values                 |
|----------------------|-------------------------------------------|---------------|---------------------------------|
| LOG_LEVEL            | Log level to use                          | debug         | debug, info, warn, error, fatal |
| LOG_FORMAT           | Log format to use                         | console       | console, json                   |
| GIN_MODE             | Gin mode to use                           | release       | release, debug                  |
| SERVER_PORT          | Port to listen on                         | 8080          | any port you find reasonable    |
| RATELIMITER_RATE     | Rate at which requests are allowed        | second        | second, minute                  |
| RATELIMITER_LIMIT    | Maximum number of requests allowed        | 5             |                                 |
| RATELIMITER_KEY      | Key to use for the ratelimiter            | client_ip     | client_ip                       |
| CHALLENGE_DIFFICULTY | Difficulty of the proof of work challenge | 20            | 1 to 30 (recommended)           |
| SALT_LENGTH          | Length of the salt                        | 8             |                                 |

### Client

| Name                    | Description                               | Default Value     | Possible Values                                                  |
|-------------------------|-------------------------------------------|-------------------|------------------------------------------------------------------|
| LOG_LEVEL               | Log level to use                          | debug             | debug, info, warn, error, fatal                                  |
| LOG_FORMAT              | Log format to use                         | console           | console, json                                                    |
| SERVER_HOST             | Host of the server to connect to          | localhost         | wherever server is hosted                                        |
| SERVER_PORT             | Port of the server to connect to          | 8080              | whichever port server is listebing on                            |
| REQUEST_PATH            | Path of the request to send to the server | /v1/quotes/random | whichever endpoint you want to hit on server                     |
| REQUEST_RATE_PER_SECOND | Number of requests per second to send     | 100               |                                                                  |
| REQUEST_COUNT           | Number of requests to send to the server  | 0                 | 0 means "run indefinetily", any positive number would limit that |

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

The structure of the project is inspired by Standard Go Project Layout.
For more information on the architectural thinking behind this structure,
see [docs/002_Project_architecture.md](docs/002_Project_architecture.md).

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
> > / **test/integration** - integration tests

## Contributing

I used [Conventional Commits](https://www.conventionalcommits.org/en/v1.0.0/) for commit messages.
I also recommend adhering
to [TDD (Red-Green-Refactor) principle](https://www.codecademy.com/article/tdd-red-green-refactor) when contributing.

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

## Infrastructure

The project is hosted on GitHub and uses GitHub Actions for CI/CD pipelines.
For actual infrastructure, see infrastructure repository - [link, WIP](https://github.com/daniel-orlov/quotes-infra)
