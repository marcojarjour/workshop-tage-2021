
# VARIABLES
# -


# CONFIG
.PHONY: help print-variables
.DEFAULT_GOAL := help


# ACTIONS

build-all :		## Build all applications
	cd standalone/ && make build
	cd http-client/ && make build
	cd http-server/ && make build
	cd http-server-db/ && make build
	cd grpc-client/ && make build
	cd grpc-server/ && make build
	cd kafka-consumer/ && make build
	cd kafka-producer/ && make build

__check-container-tag-all :
	@[ "$(CONTAINER_TAG)" ] || ( echo "Missing container tag (CONTAINER_TAG), please define it and retry"; exit 1 )

container-build-push-all : __check-container-tag-all		## Build and push ontainers of all applications
	cd standalone/ && make container-build container-push CONTAINER_TAG=$(CONTAINER_TAG)
	cd http-client/ && make container-build container-push CONTAINER_TAG=$(CONTAINER_TAG)
	cd http-server/ && make container-build container-push CONTAINER_TAG=$(CONTAINER_TAG)
	cd http-server-db/ && make container-build container-push CONTAINER_TAG=$(CONTAINER_TAG)
	cd grpc-client/ && make container-build container-push CONTAINER_TAG=$(CONTAINER_TAG)
	cd grpc-server/ && make container-build container-push CONTAINER_TAG=$(CONTAINER_TAG)
	cd kafka-consumer/ && make container-build container-push CONTAINER_TAG=$(CONTAINER_TAG)
	cd kafka-producer/ && make container-build container-push CONTAINER_TAG=$(CONTAINER_TAG)


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
