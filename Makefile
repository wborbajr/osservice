# ---------------------------------------------------------------------
# Copyright (c) 2020 WALDIR BORBA JUNIOR, B+ Tech. All Rights Reserved.
# Author(s): Anthony Potappel
# 			 Waldir Borba Junior
#
# This software may be modified and distributed under the terms of the
# MIT license. See the LICENSE file for details.
# ---------------------------------------------------------------------
#
#
# https://itnext.io/docker-makefile-x-ops-sharing-infra-as-code-parts-ea6fa0d22946
#

include docker.env

# Makefile
.EXPORT_ALL_VARIABLES:

#GO111MODULE=on
#GOPROXY=https://proxy.golang.org
#GONOSUMDB=off
#GOPRIVATE=*. internal.mycompany.com

export PROJECT_NAME
export DOCKER_TAG
export DOCKER_USER
export CONTAINER
export GOBIN ?= $(shell pwd)/bin

GOLINT = $(GOBIN)/golint

GO_FILES := $(shell \
	find . '(' -path '*/.*' -o -path './vendor' ')' -prune \
	-o -name '*.go' -print | cut -b3-)

#
GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean -cache -modcache
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get
GORUN=$(GOCMD) run
GOMOD=$(GOCMD) mod
#
DOKCMD=docker
DOKCOMPCMD=docker-compose
#
BINARY_NAME=osservice

.PHONY: default

default: help

# production: stop dang rebuild deploy

#
# GOLang CMD
# ------------------------------------------------------------------------------

gobuild:
	rm -f go.*
	rm -f vendor/
	# $(GOCLEAN)
	$(GOMOD) init github.com/wborbajr/osservice
	$(GOMOD) tidy

	@echo "########## GO Building starting ... "
	CGO_ENABLED=0 GOOS=linux $(GOBUILD) --trimpath -ldflags="-s -w" -o $(BINARY_NAME) main.go
	@echo "GO Build done..."

lint: $(GOLINT)
	@rm -rf lint.log
	@echo "Checking formatting..."
	@gofmt -d -s $(GO_FILES) 2>&1 | tee lint.log
	@echo "Checking vet..."
	@go vet ./... 2>&1 | tee -a lint.log
	@echo "Checking lint..."
	@$(GOLINT) ./... | tee -a lint.log
	@echo "Checking for unresolved FIXMEs..."
	@git grep -i fixme | grep -v -e vendor -e Makefile -e .md | tee -a lint.log
	@[ ! -s lint.log ]

#
# Production
# ------------------------------------------------------------------------------

prodbuild:
	@echo "\nStarting build...\n"
	$(DOKCOMPCMD) -f $(DKRFILEPROD) -f $(DKRFILEPRODOVR) build

prodrebuild:
	@echo "\nForcing Rebuild...\n"
	$(DOKCOMPCMD) -f $(DKRFILEPROD) -f $(DKRFILEPRODOVR) build --no-cache --force-rm --pull

prodstart:
	@echo "\nStarting production mode...\n"
	$(DOKCOMPCMD) -f $(DKRFILEPROD) -f $(DKRFILEPRODOVR) up -d

prodstop:
	@echo "\nstoping production mode...\n"
	$(DOKCOMPCMD) -f $(DKRFILEPROD) -f $(DKRFILEPRODOVR) stop

proddown:
	@echo "\nstoping production mode...\n"
	$(DOKCOMPCMD) -f $(DKRFILEPROD) -f $(DKRFILEPRODOVR) down -v --remove-orphans --rmi all

prodtop:
	$(DOKCOMPCMD) -f $(DKRFILEPROD) -f $(DKRFILEPRODOVR) top

prodps:
	$(DOKCOMPCMD) -f $(DKRFILEPROD) -f $(DKRFILEPRODOVR) ps

prodlogs:
	$(DOKCOMPCMD) -f $(DKRFILEPROD) -f $(DKRFILEPRODOVR) logs

prodevents:
	$(DOKCOMPCMD) -f $(DKRFILEPROD) -f $(DKRFILEPRODOVR) events

prodpause:
	$(DOKCOMPCMD) -f $(DKRFILEPROD) -f $(DKRFILEPRODOVR) pause

produnpause:
	$(DOKCOMPCMD) -f $(DKRFILEPROD) -f $(DKRFILEPRODOVR) unpause

#
# Development
# ------------------------------------------------------------------------------

devbuild:
	@echo "\nStarting build...\n"
	$(DOKCOMPCMD) -f $(DKRFILEDEV) -f $(DKRFILEDEVOVR) build

devrebuild:
	@echo "\nForcing Rebuild...\n"
	$(DOKCOMPCMD) -f $(DKRFILEDEV) -f $(DKRFILEDEVOVR) build --no-cache --force-rm --pull

devstart:
	@echo "\nStarting container...\n"
	$(DOKCOMPCMD) -f $(DKRFILEDEV) -f $(DKRFILEDEVOVR) up -d

devstop:
	@echo "\nStoping container...\n"
	$(DOKCOMPCMD) -f $(DKRFILEDEV) -f $(DKRFILEDEVOVR) stop

devdown:
	@echo "\nStoping container...\n"
	$(DOKCOMPCMD) -f $(DKRFILEDEV) -f $(DKRFILEDEVOVR) down -v --remove-orphans --rmi all

devtop:
	$(DOKCOMPCMD) -f $(DKRFILEDEV) -f $(DKRFILEDEVOVR) top

devps:
	$(DOKCOMPCMD) -f $(DKRFILEDEV) -f $(DKRFILEDEVOVR) ps

devlogs:
	$(DOKCOMPCMD) -f $(DKRFILEDEV) -f $(DKRFILEDEVOVR) logs

devevents:
	$(DOKCOMPCMD) -f $(DKRFILEDEV) -f $(DKRFILEDEVOVR) events

devpause:
	$(DOKCOMPCMD) -f $(DKRFILEDEV) -f $(DKRFILEDEVOVR) pause

devunpause:
	$(DOKCOMPCMD) -f $(DKRFILEDEV) -f $(DKRFILEDEVOVR) unpause

#
# Docker CMD
# ------------------------------------------------------------------------------

exec:
	@echo "\nEntering container...\n"
	$(DOKCMD) exec -ti ${CONTAINER} /bin/bash

deploy:
	@echo "\nStarting delpoy...\n"
	$(DOKCMD) push ${DOCKER_USER}/${PROJECT_NAME}:${DOCKER_TAG}

dang:
	@echo "\nStarting dangling removal\n"
	$(DOKCMD) rmi $$($(DOKCMD) images -q -f dangling=true)

prune:
	$(DOKCMD) system prune -a -f --volumes

remove:
	$(DOKCMD) rm $$($(DOKCMD) ps -a -q) -f

help:
	@echo ''
	@echo 'Usage: make [TARGET] [EXTRA_ARGUMENTS]'
	@echo 'Targets:'
	@echo '  build    	build docker --image-- for current user: $(HOST_USER)(uid=$(HOST_UID))'
	@echo '  rebuild  	rebuild docker --image-- for current user: $(HOST_USER)(uid=$(HOST_UID))'
	@echo '  test     	test docker --container-- for current user: $(HOST_USER)(uid=$(HOST_UID))'
	@echo '  service   	run as service --container-- for current user: $(HOST_USER)(uid=$(HOST_UID))'
	@echo '  login   	run as service and login --container-- for current user: $(HOST_USER)(uid=$(HOST_UID))'
	@echo '  clean    	remove docker --image-- for current user: $(HOST_USER)(uid=$(HOST_UID))'
	@echo '  prune    	shortcut for docker system prune -af. Cleanup inactive containers and cache.'
	@echo '  shell      run docker --container-- for current user: $(HOST_USER)(uid=$(HOST_UID))'
	@echo ''
	@echo 'Extra arguments:'
	@echo 'cmd=:	make cmd="whoami"'
	@echo '# user= and uid= allows to override current user. Might require additional privileges.'
	@echo 'user=:	make shell user=root (no need to set uid=0)'
	@echo 'uid=:	make shell user=dummy uid=4000 (defaults to 0 if user= set)'

# 	docker container rm $$(docker ps -aq) -f
# 	docker image rm $$(docker images --format "{{.ID}}" --filter "dangling=true")
# 	docker volume ls -f dangling=true

# $ docker run -dt <image>
# $ docker exec -it <container> <command>

# docker run -dt --name nfeimport wborbajr/nfeimport:latest
# docker exec -it nfeimport /bin/bash