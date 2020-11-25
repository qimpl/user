#!/bin/bash

set -e

docker-compose exec -T api bash <<EOF
  testfixtures -d postgres --dangerous-no-test-database-check -c "postgres://\${DB_USER}:\${DB_PASSWORD}@db/\${DB_TEST_NAME}?sslmode=disable" -D fixtures
  go test $VERBOSE ./...
EOF
