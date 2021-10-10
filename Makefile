SRC_DIR=./src
SCRIPT_DIR=./scripts
UI_MENU_DIR=./ui/menu

prep:
	go get ./...
	cd ${UI_MENU_DIR} && npm install

lint:
	golangci-lint run
	cd ${UI_MENU_DIR} && npm run ci

test:
	${SCRIPT_DIR}/test_run.sh
