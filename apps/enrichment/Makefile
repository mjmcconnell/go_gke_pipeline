ROOT_DIR        := $(abspath $(dir $(lastword ${MAKEFILE_LIST})))
BUILD_DIR       := ${ROOT_DIR}/build
BINARY_DIR      := ${BUILD_DIR}/bin
MODULE          := ${GIT_HOST}/${GIT_USER_ID}/${GIT_PROJECT_ID}
GO_TESTFLAGS 	:= -coverpkg ./... -race -cover -coverprofile=coverage.out
GO_COVER_EXCLUDE := "_mock.go"

.PHONY: help
help: ## Lists all commands.
	${Q} fgrep -h "##" $(MAKEFILE_LIST) | fgrep -v fgrep | sed -e 's/\\$$//' | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

.PHONY: watcher.install
run: watcher.install
	cd cmd/enrichment/ && watcher

.PHONY: watcher.install
watcher.install: go.mod
	go install github.com/canthefason/go-watcher/cmd/watcher

.PHONY: go.build
go.build:
	rm -rf ${BUILD_DIR}
	go build -o ${BINARY_DIR}/enrichment github.com/mjmcconnell/go_gke_pipeline/apps/enrichment/cmd/enrichment
