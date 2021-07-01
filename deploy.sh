#!/bin/bash

if [[ -z "${LA_SSH}" ]]; then
	echo "Missing LA_SSH variable"
	exit 1
fi

services="populus amnis verum nuntius"

for service in $services; do
	echo ----------BUILDING $service----------
	docker buildx build -t satvikr/liveassist_$service:latest -f ./docker/Dockerfile.$service .
	echo ----------SENDING $service----------
	docker save satvikr/liveassist_$service:latest | bzip2 | ssh $LA_SSH "bunzip2 | sudo docker load"
done

files="./deploy/docker-compose.yml ./deploy/config/.env"

for file in $files; do
	bname="$(basename $file)"
	echo ----------SENDING $file----------
	cat $file | bzip2 | ssh $LA_SSH "bunzip2 > $bname"
done

echo ----------RESTARTING SERVICES----------
ssh $LA_SSH "sudo docker-compose stop"
for service in $services; do
	ssh $LA_SSH "sudo docker-compose rm -f -- $service"
done
ssh $LA_SSH "sudo docker-compose --env-file .env up -d"
echo ----------PRUNING OLD IMAGES----------
ssh $LA_SSH "sudo docker image prune -f"