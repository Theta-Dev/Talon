#!/bin/bash
set -e

dialects=("sqlite" "mysql" "postgres")

cd -- "$( dirname -- "${BASH_SOURCE[0]}" )/../src"

../scripts/test_up.sh wait

go get ./...

printBlue() {
    echo -e "\e[30;46m ${1} \e[0m ${@:2}"
}

printGreen() {
    echo -e "\e[30;42m ${1} \e[0m ${@:2}"
}

printRed() {
    echo -e "\e[30;41m ${1} \e[0m ${@:2}"
}

failures=""

for dialect in "${dialects[@]}" ; do
    echo "------------------------"
    printBlue "TESTING" $dialect
    echo "------------------------"

    if DIALECT=${dialect} GOFLAGS="-count=1" go test -v -p 1 -timeout 1m ./...; then
        printGreen "PASS" $dialect
    else
        printRed "FAIL" $dialect
        failures+=${dialect}" "
    fi
done

echo "------------------------"
if [ -z "$failures" ]; then
    printGreen "ALL GOOD"
else
    printRed "FAILED TESTS" $failures
    exit 1
fi
