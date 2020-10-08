#!/bin/bash

PROJECT_NAME=$1

fileList="Dockerfile go.mod .env.dist docker-compose.yml main.go router/router.go router/healthy_check.go .gitignore"

for f in ${fileList};
do
  if [[ "$OSTYPE" == "darwin"* ]]; then
    sed -i '' "s/APP_NAME/$PROJECT_NAME/g" ${f};
  else
    sed -i "s/APP_NAME/$PROJECT_NAME/g" ${f};
  fi
done

swag init

rm config.sh
