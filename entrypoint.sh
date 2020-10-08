#!/bin/sh

postgres=postgres://${DB_USER}:${DB_PASSWORD}@db/${DB_NAME}?sslmode=disable

go get -tags 'postgres' -u github.com/golang-migrate/migrate/v4/cmd/migrate/

migrate -source file://migrations -database $postgres up

go get github.com/cespare/reflex

exec "$@"
