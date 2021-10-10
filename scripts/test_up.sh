#!/bin/bash
set -e

cd -- $( dirname -- "${BASH_SOURCE[0]}" )

if docker start talon_mariadb talon_postgres &> /dev/null; then
    echo "containers already up"
else
    docker run --rm -d -p 127.0.0.1:3306:3306 --env-file dbtest.env --name talon_mariadb mariadb:latest
    docker run --rm -d -p 127.0.0.1:5432:5432 --env-file dbtest.env --name talon_postgres postgres:latest
    echo "containers started"

    if [[ "$1" == "wait" ]]; then
        echo "waiting 15s for startup"
        sleep 15
    fi
fi
