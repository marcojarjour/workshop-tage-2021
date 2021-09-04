
# VARIABLES
# -


# CONFIG
.PHONY: help print-variables
.DEFAULT_GOAL := help


# ACTIONS

## minikube

start-docker-minikube :		## Start Minikube with Docker driver
	minikube start --driver=docker --cpus=8 --memory=12288 --disk-size=50g

start-hyperkit-minikube :		## Start Minikube with Hyperkit driver
	minikube start --driver=hyperkit --cpus=8 --memory=12288 --disk-size=50g

start-virtualbox-minikube :		## Start Minikube with VirtualBox driver
	minikube start --driver=virtualbox --cpus=8 --memory=12288 --disk-size=50g

stop-minikube :		## Stop Minikube
	minikube stop

delete-minikube :		## Delete Minikube
	minikube delete

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
