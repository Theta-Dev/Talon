#!/bin/bash
set -e

docker kill talon_mariadb talon_postgres
echo "containers stopped"
