#!/bin/sh

postgres=postgres://${DB_USER}:${DB_PASSWORD}@db/${DB_NAME}?sslmode=disable

go get -tags 'postgres' -u github.com/golang-migrate/migrate/v4/cmd/migrate/

migrate -source file://migrations -database $postgres up

ORG_NAME=go-testfixtures
REPO_NAME=testfixtures
LATEST_VERSION=$(curl -s https://api.github.com/repos/${ORG_NAME}/${REPO_NAME}/releases/latest | \
  grep "tag_name" | cut -d'v' -f2 | cut -d'"' -f1)

mkdir -p cmd && \
cd cmd && \
wget https://github.com/${ORG_NAME}/${REPO_NAME}/releases/download/v${LATEST_VERSION}//testfixtures_linux_amd64.tar.gz \
-O testfixtures.tar.gz && \
tar -xvf testfixtures.tar.gz && \
mv testfixtures /bin/testfixtures && \
cd ../ && \
rm -rf cmd

testfixtures -d postgres --dangerous-no-test-database-check -c "$postgres" -D fixtures

go get github.com/cespare/reflex

exec "$@"
