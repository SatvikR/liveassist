#!/bin/bash

if [[ -z "${LA_SSH}" ]]; then
	echo "Missing LA_SSH variable"
	exit 1
fi

docker build -t satvikr/liveassist_populus:latest -f ./docker/Dockerfile.populus .
docker save satvikr/liveassist_populus:latest | bzip2 | ssh $LA_SSH "bunzip2 | sudo docker load"

docker build -t satvikr/liveassist_amnis:latest -f ./docker/Dockerfile.amnis .
docker save satvikr/liveassist_amnis:latest | bzip2 | ssh $LA_SSH "bunzip2 | sudo docker load"

docker build -t satvikr/liveassist_verum:latest -f ./docker/Dockerfile.verum .
docker save satvikr/liveassist_verum:latest | bzip2 | ssh $LA_SSH "bunzip2 | sudo docker load"

docker build -t satvikr/liveassist_nuntius:latest -f ./docker/Dockerfile.nuntius .
docker save satvikr/liveassist_nuntius:latest | bzip2 | ssh $LA_SSH "bunzip2 | sudo docker load"

cat ./deploy/docker-compose.yml | bzip2 | ssh $LA_SSH "bunzip2 >> docker-compose.yml"
ssh $LA_SSH "sudo docker-compose up -d"