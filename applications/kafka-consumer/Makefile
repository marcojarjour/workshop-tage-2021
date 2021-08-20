
# VARIABLES
export GO111MODULE = on


# CONFIG
.PHONY: help print-variables
.DEFAULT_GOAL := help


# ACTIONS

## application

build :		## Build application
	go build ./...

build-linux :		## Build application for linux
	CGO_ENABLED=1 GOOS=linux GOARCH=amd64 go build -tags netgo -o /bin/app -ldflags "-w -s"

run :		## Run application
	godotenv -f local.env go run main.go


## container

__check-container-tag :
	@[ "$(CONTAINER_TAG)" ] || ( echo "Missing container tag (CONTAINER_TAG), please define it and retry"; exit 1 )

container-build : __check-container-tag		## Build container
	DOCKER_BUILDKIT=0 docker build . -t bygui86/kafka-consumer:$(CONTAINER_TAG) --no-cache

container-build-with-buildkit : __check-container-tag		## Build container
	DOCKER_BUILDKIT=1 docker build . -t bygui86/kafka-producer:$(CONTAINER_TAG) --no-cache

container-push : __check-container-tag		## Push container to Docker hub
	docker push bygui86/kafka-consumer:$(CONTAINER_TAG)


## helpers

help :		## Help
	@echo ""
	@echo "*** \033[33mMakefile help\033[0m ***"
	@echo ""
	@echo "Targets list:"
	@grep -E '^[a-zA-Z_-]+ :.*?## .*$$' $(MAKEFILE_LIST) | sort -k 1,1 | awk 'BEGIN {FS = ":.*?## "}; {printf "\t\033[36m%-30s\033[0m %s\n", $$1, $$2}'
	@echo ""

print-variables :		## Print variables values
	@echo ""
	@echo "*** \033[33mMakefile variables\033[0m ***"
	@echo ""
	@echo "- - - makefile - - -"
	@echo "MAKE: $(MAKE)"
	@echo "MAKEFILES: $(MAKEFILES)"
	@echo "MAKEFILE_LIST: $(MAKEFILE_LIST)"
	@echo "- - -"
	@echo ""
