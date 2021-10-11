#!/bin/bash
set -e

if [[ -n $(docker ps -q -f name=talon_) ]]; then
	docker kill talon_mariadb talon_postgres
	echo "containers stopped"
fi
