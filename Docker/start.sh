#!/bin/bash

# Exit on first error, print all commands.
# set -ev

# Grab the current directory
CURR_DIR=$(pwd)
DIR=$(dirname $0)
cd ${DIR}
DIR=$(pwd)
cd ${CURR_DIR}

display_help()
{
	echo "Usage $0 [,OPTIONS]
OPTIONS:
    -h|--help           Display this help message.
    -n|--node           Deploy node container.
    -p|--peer           Deploy extended peers containers.
    -f|--full           Deploy node and extend peers containers (like -n -p).
    -s|--start          Start containers containers.
    -S|--stop           Stop containers.
    -t|--tag            Tag hyperledger image as latest.
    -r|--remove         Remove containers.
    -a|--all            Make special actions on all docker containers on your machine.
"
}

# baseimage_version="x86_64-0.3.0"
baseimage_version="x86_64-0.1.0"

compose_init_file="docker-compose.yml"
compose_peers_file="docker-compose-extend-peers.yml"
compose_node_file="docker-compose-extend-node.yml"
compose_files="${compose_init_file}"
DEFAULT_CONFIG_FILE="default_config.json"
CONFIG_FILE=${DEFAULT_CONFIG_FILE}
CONFIG_TARGET_FILE="default.json"

NPM_CONFIG_DIR="${DIR}/../universal-payment/config"

ALL=false
STOP=false
START=false
REMOVE=false
TAG=false

while [ ! -z "$1" ] ; do
	case $1 in
		-h|--help)
			display_help
			exit 0
			;;
		-a|--all)
			ALL=true
			;;
		-f|--full)
			compose_files="${compose_files} ${compose_node_file} ${compose_peers_file}"
			;;
		-n|--node)
			compose_files="${compose_files} ${compose_node_file}"
			;;
		-p|--peers)
			compose_files="${compose_files} ${compose_peers_file}"
			CONFIG_FILE="production.json"
			;;
		-r|--remove)
			compose_files="${compose_files} ${compose_node_file} ${compose_peers_file}"
			REMOVE=true
			;;
		-s|--start)
			START=true
			;;
		-S|--stop)
			STOP=true
			;;
		-t|--tag)
			TAG=true
			;;
		*)
			echo "Unknown option '$1'"
			;;
	esac
	shift
done

COMPOSE_OPTS=""
for compose_file in ${compose_files} ; do
	COMPOSE_OPTS="${COMPOSE_OPTS} -f ${DIR}/${compose_file}"
done

if ${TAG} ; then
	# Tag the latest version of fabric-baseimage
	echo "Tag the latest version of fabric-baseimage"
	docker pull hyperledger/fabric-baseimage:${baseimage_version}
	docker tag hyperledger/fabric-baseimage:${baseimage_version} hyperledger/fabric-baseimage:latest
	exit 0
fi

cp ${NPM_CONFIG_DIR}/${CONFIG_FILE} ${NPM_CONFIG_DIR}/${CONFIG_TARGET_FILE}

# Clean up old docker containers

if ${ALL} ; then
	if ${REMOVE} ; then
		echo "Remove all containers"
		docker rm -vf `docker ps -aq` 2>/dev/null
	fi
	if ${START} ; then
		echo "Start all containers"
		docker start `docker ps -aq` 2>/dev/null
	elif ${STOP} ; then
		echo "Stop all containers"
		docker stop `docker ps -aq` 2>/dev/null
	fi
else 
	if ${STOP} ; then
		echo "Stop composed containers"
		docker-compose ${COMPOSE_OPTS} stop;
	else
		if ${REMOVE} ; then
			echo "Remove composed containers"
			docker-compose ${COMPOSE_OPTS} kill;
			docker-compose ${COMPOSE_OPTS} down;
		fi
		if ! ${REMOVE} && ! ${STOP} ; then
			echo "Rebuild composed containers"
			docker-compose ${COMPOSE_OPTS} build;
			docker-compose ${COMPOSE_OPTS} up -d ;
		fi
	fi
fi

echo "
Current containers:"
docker ps -a

exit 0
