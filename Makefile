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

#
# Production
# ------------------------------------------------------------------------------

prodbuild:
	@echo "\nStarting build...\n"
	$(DOKCOMPCMD) -f $(DOCKERFILE) -f $(DOCKERFILEPROD) build

prodrebuild:
	@echo "\nForcing Rebuild...\n"
	$(DOKCOMPCMD) -f $(DOCKERFILE) -f $(DOCKERFILEPROD) build --no-cache --force-rm --pull

produp:
	@echo "\nStarting production mode...\n"
	$(DOKCOMPCMD) -f $(DOCKERFILE) -f $(DOCKERFILEPROD) up -d

prodstop:
	@echo "\nstoping production mode...\n"
	$(DOKCOMPCMD) -f $(DOCKERFILE) -f $(DOCKERFILEPROD) stop

proddown:
	@echo "\nstoping production mode...\n"
	$(DOKCOMPCMD) -f $(DOCKERFILE) -f $(DOCKERFILEPROD) down -v --remove-orphans --rmi all

prodtop:
	$(DOKCOMPCMD) -f $(DOCKERFILE) -f $(DOCKERFILEPROD) top

prodps:
	$(DOKCOMPCMD) -f $(DOCKERFILE) -f $(DOCKERFILEPROD) ps

prodlogs:
	$(DOKCOMPCMD) -f $(DOCKERFILE) -f $(DOCKERFILEPROD) logs

prodevents:
	$(DOKCOMPCMD) -f $(DOCKERFILE) -f $(DOCKERFILEPROD) events

prodpause:
	$(DOKCOMPCMD) -f $(DOCKERFILE) -f $(DOCKERFILEPROD) pause

produnpause:
	$(DOKCOMPCMD) -f $(DOCKERFILE) -f $(DOCKERFILEPROD) unpause

#
# Development
# ------------------------------------------------------------------------------

devbuild:
	@echo "\nStarting build...\n"
	$(DOKCOMPCMD) build

devrebuild:
	@echo "\nForcing Rebuild...\n"
	$(DOKCOMPCMD) build --no-cache --force-rm --pull

devup:
	@echo "\nStarting container...\n"
	$(DOKCOMPCMD) up -d

devstop:
	@echo "\nStoping container...\n"
	$(DOKCOMPCMD) stop

devdown:
	@echo "\nStoping container...\n"
	$(DOKCOMPCMD) down -v --remove-orphans --rmi all

devtop:
	$(DOKCOMPCMD) top

devps:
	$(DOKCOMPCMD) ps

devlogs:
	$(DOKCOMPCMD) logs

devevents:
	$(DOKCOMPCMD) events

devpause:
	$(DOKCOMPCMD) pause

devunpause:
	$(DOKCOMPCMD) unpause

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