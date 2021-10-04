#!/bin/bash

cd -- "$( dirname -- "${BASH_SOURCE[0]}" )/.."

docker run --rm -d -p 127.0.0.1:3306:3306 --env-file ./scripts/dbtest.env --name mariadb mariadb
docker run --rm -d -p 127.0.0.1:5432:5432 --env-file ./scripts/dbtest.env --name postgres postgres
