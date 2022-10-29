#!/bin/bash
# Need to login to registry.cloud.okteto.net
# Need cluster config for cloud okteto

DIRECTORY=$(pwd)
BACKEND="${DIRECTORY}/backend"
BOT="${DIRECTORY}/bot"

DEPLOY_MANIFEST="${DIRECTORY}/publish/prod/deployment.yml"
buildAndPushBackend (){
   local REGISTRY=registry.cloud.okteto.net/kevin-vargas
   local IMAGE_TAG="${REGISTRY}/$1"
   echo $IMAGE_TAG
   echo $2
   docker build -t $IMAGE_TAG $2
   docker push $IMAGE_TAG
}

apply() {
    kubectl apply -f $1
}

log() {
    echo -e "${1}"
}

logt() {
    log "\t ♥♥♥ ${1} \t ♥♥♥"
}

{

    logt "building backend and pusing image"
    buildAndPushBackend be-futbol-matches $BACKEND
    buildAndPushBackend bot-futbol-matches $BOT

    logt "deploying kubernetes objects"
    apply $DEPLOY_MANIFEST

}