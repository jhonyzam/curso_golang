# Products API

Products API is an example project, created with the objective of setting some quality standards for the Solution Team's Go projects.
It already follows Go project structure best practices, examples of CI configuration, API documentation, tests and environment setup for local tests.


## REST API Documentation

All available REST API documentation exposed by the project was documented using the [OpenAPI](https://www.openapis.org/) standard.

To view this documentation locally use the following command:
```bash
make run-swagger
```
The same documentation is available online and can be accessed by clicking [here](https://sd-projects.gitlab.neoway.io/products/api/index.html).


_API routes documentation generate by [Swagger-UI](https://github.com/swagger-api/swagger-ui)._


## Getting Started 

### Prerequisites

- [Golang](http://golang.org/)(>11.0)
- [GNU Make](https://www.gnu.org/software/make/)
- [Docker](http://docker.com)
- [Docker Compose](https://docs.docker.com/compose/install/)


### Environment variables

```bash
LOGGER_LEVEL=error #values: error; warn; info; debug;
PRODUCTS_PORT=8080
POSTGRES_URL=postgres://postgres:admin@$localhost:5432/products?sslmode=disable
API_CLIENT_ID=app-id
API_CLIENT_SECRET=app-secret
NEOWAY_URL=https://api.neoway.com.br
```

### Installing and running locally
```shell script
# Install dependencies
make install

# Run postgres locally as a container
make env

# Run server locally
make run

# Run server locally with custom environment variables
POSTGRES_URL=postgres://postgres:admin@$localhost:5432/products?sslmode=disable \
PRODUCTS_PORT=5001 \
LOGGER_LEVEL=debug \
API_CLIENT_ID=app-id
API_CLIENT_SECRET=app-secret
NEOWAY_URL=https://api.neoway.com.br
make run
```

### Setting up git hooks

After cloning the repository, change the git hooks path so it's only possible to commit code with the required quality.

```shell script
make git-config
```

## Running the tests and coverage report

To view report of tests locally use the following command:

```bash
make env # prepares environment for testing
make test
```PostgreSQL - products@localhost

## Running the lint verification

```bash
make lint
```
_Lint report generate by [GolangCI-lint](https://github.com/golangci/golangci-lint)._

## Running vulnerability check in Go dependencies
```bash
make audit
```
_Audit report generate by [Nancy](https://github.com/sonatype-nexus-community/nancy)._


## Deployment

### Build

```bash
make build
```

### Create release image, add tag and push

```bash
make image tag push
```

### Run registry image locally

```bash
make run-docker

make remove-docker
```
  
## Inspiration

### Package organization

The package structure used in this project was inspired by the [golang-standards](https://github.com/golang-standards/project-layout) project.

### Project layers organization

The project layers structure used in this project was inspired by the **Hexagonal Architecture** (Ports & Adapters).  

  
## Contributing
See [CONTRIBUTING](CONTRIBUTING.md) documentation for more details.

  
## Changelog
See [CHANGELOG](CHANGELOG.md) documentation for more details.