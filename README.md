# Golang Web API Template

## API Documentation

Swagger API documentation can be found at `/v1/swagger`

## Getting started

You need to execute `config.sh` before any action on this template. This script take one arg that is the project name.

```sh
$ ./config.sh my-super-api
```

After that, you can run `make init` to copy all configuration files, then you can run `docker-compose up`.

## Tests

Run tests using the Make recipe `test`

```sh
$ make test
```

To add verbosity to `go test` command

```sh
$ make test VERBOSE=-v
```

## Code Quality

Your commits names should follow the [@commitlint/config-conventional](https://github.com/conventional-changelog/commitlint/tree/master/@commitlint/config-conventional) rules.

Check and fix your code using:

```sh
$ make coding-style
```

## Pre-commit

Install pre-commit following the [official documentation](https://pre-commit.com/#installation)

Setup your pre-commit hooks using:

```sh
# pre-commit hooks
$ pre-commit install
# message commit hooks
$ pre-commit install --hook-type commit-msg
```

For additional information check [pre-commit docs](https://pre-commit.com)
