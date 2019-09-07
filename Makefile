CONTAINERNAME_001=mongo001
CONTAINERNAME_002=mongo002
CONTAINERNAME_003=mongo003
DOCKER_RUN_OPTIONS=

all_container=`docker ps -a -q`
active_container=`docker ps -q`
images=`docker images | awk '/^<none>/ { print $$3 }'`

default: build
build:
	docker-compose build
restart: stop start
start: up
up:
	docker-compose up -d && docker-compose logs
stop: down
down:
	docker-compose down

logs:
	docker-compose logs
logsf:
	docker-compose logs -f

console: attach
attach:
	docker exec -it $(CONTAINERNAME_001) /bin/bash --login
console2: attach2
attach2:
	docker exec -it $(CONTAINERNAME_002) /bin/bash --login
console3: attach3
attach3:
	docker exec -it $(CONTAINERNAME_003) /bin/bash --login

setuprs:
	docker exec -it $(CONTAINERNAME_001) /scripts/setuprs


clean: clean_container clean_images
clean_images:
	@if [ "$(images)" != "" ] ; then \
		docker rmi $(images); \
	fi
clean_container:
	@for a in $(all_container) ; do \
		for b in $(active_container) ; do \
			if [ "$${a}" = "$${b}" ] ; then \
				continue 2; \
			fi; \
		done; \
		docker rm $${a}; \
	done
