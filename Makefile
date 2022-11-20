##@ Help
.PHONY: help
help: ## Display this help screen
	@awk 'BEGIN {FS = ":.*##"; printf "\nUsage:\n  make \033[36m<target>\033[0m\n"} /^[a-zA-Z_-]+:.*?##/ { printf "  \033[36m%-15s\033[0m %s\n", $$1, $$2 } /^##@/ { printf "\n\033[1m%s\033[0m\n", substr($$0, 5) } ' $(MAKEFILE_LIST)


##@ Build
TARGETS := $(shell ls -d cmd/* | cut -f2 -d'/')


##@ Go
.PHONY: vendor
vendor: ## Install dependencies
	go mod vendor

.PHONY: test
test: ## Run tests
	go test -v ./...

TEST_FN ?= .
.PHONY: test-fn
test-fn: ## Run tests for a specific function
	go test -v ./... -run $(TEST_FN)


##@ Vagrant
VAGRANT_BOX ?= "generic/ubuntu2204"

.PHONY: vagrant-up
vagrant-up: ## Start the vagrant box
	@sed 's|config.vm.box = ".*"|config.vm.box = $(VAGRANT_BOX)|' Vagrantfile > Vagrantfile.tmp && mv Vagrantfile.tmp Vagrantfile
	vagrant up

.PHONY: vagrant-reload
vagrant-reload: ## Reload the vagrant box
	vagrant reload

.PHONY: vagrant-ssh
vagrant-ssh: ## SSH into the vagrant box, password: vagrant
	vagrant ssh

.PHONY: vagrant-stop
vagrant-stop: ## Stop the vagrant box
	vagrant halt

.PHONY: vagrant-destroy
vagrant-destroy: ## Destroy the vagrant box
	vagrant destroy

.PHONY: vagrant-rm-box
vagrant-rm-box: ## Remove the vagrant box
	vagrant box remove $(VAGRANT_BOX)
