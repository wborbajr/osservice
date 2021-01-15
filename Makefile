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

export PROJECT_NAME
export DOCKER_TAG
export DOCKER_USER
export CONTAINER

.PHONY: default

default: help

production: stop dang rebuild deploy

gobuild:
	rm -f go.*
	rm -f vendor/
	go mod init github.com/wborbajr/osservice
	go mod tidy

	@echo "########## GO Building starting ... "
	CGO_ENABLED=0 GOOS=linux go build --trimpath -ldflags="-s -w" -o osservice main.go
	@echo "GO Build done..."

produp:
	@echo "\nStarting production mode...\n"
	docker-compose -f docker-compose.yml -f docker-compose.prod.yml up -d

proddown:
	@echo "\nstoping production mode...\n"
	docker-compose -f docker-compose.yml -f docker-compose.prod.yml down --remove-orphans

build:
	@echo "\nStarting build...\n"
	docker-compose -f ${DOCKERFILE} build

rebuild:
	@echo "\nForcing Rebuild...\n"
	docker-compose -f ${DOCKERFILE} build --no-cache --force-rm --pull

start:
	@echo "\nStarting container...\n"
	docker-compose -f ${DOCKERFILE} up -d

stop:
	@echo "\nStoping container...\n"
	docker-compose -f ${DOCKERFILE} down --remove-orphans

stopall:
	@echo "\nStoping container...\n"
	docker-compose -f ${DOCKERFILE} down --remove-orphans --rmi all

exec:
	@echo "\nEntering container...\n"
	docker exec -ti ${CONTAINER} /bin/bash

deploy:
	@echo "\nStarting delpoy...\n"
	docker push ${DOCKER_USER}/${PROJECT_NAME}:${DOCKER_TAG}

dang:
	@echo "\nStarting dangling removal\n"
	docker rmi $$(docker images -q -f dangling=true)

develop:
	docker-compose up --force-recreate --build && docker-compose down --remove-orphans

restart:
	docker-compose -f ${DOCKERFILE} restart

pause:
	docker-compose -f ${DOCKERFILE} pause

unpause:
	docker-compose -f ${DOCKERFILE} unpause

top:
	docker-compose -f ${DOCKERFILE} top

ps:
	docker-compose -f ${DOCKERFILE} ps

logs:
	docker-compose -f ${DOCKERFILE} logs

events:
	docker-compose -f ${DOCKERFILE} events

prune:
	docker system prune -a -f --volumes

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