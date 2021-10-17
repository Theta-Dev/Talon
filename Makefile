SRC_DIR=./src
SCRIPT_DIR=./scripts
UI_MENU_DIR=./ui/menu

setup:
	go get -t ./...
	cd ${UI_MENU_DIR} && npm install

go-lint:
	golangci-lint run

js-lint:
	cd ${UI_MENU_DIR} && npm run ci

lint: go-lint js-lint

test:
	${SCRIPT_DIR}/test_run.sh

clean:
	${SCRIPT_DIR}/test_down.sh
	rm -rf ${UI_MENU_DIR}/dist
