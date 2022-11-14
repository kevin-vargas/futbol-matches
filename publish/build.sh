#!/bin/bash
# Need to login to registry.cloud.okteto.net
# Need cluster config for cloud okteto

DIRECTORY=$(pwd)
BACKEND="${DIRECTORY}/backend"
BOT="${DIRECTORY}/bot"
FRONT="${DIRECTORY}/frontend"
REGISTRY=$1
DEPLOY_MANIFEST="${DIRECTORY}/${2}/deployment.yml"

buildAndPush (){
   local IMAGE_TAG="${REGISTRY}/$1"
   docker build -t $IMAGE_TAG $2
   docker push $IMAGE_TAG
}

apply() {
    kubectl apply -f $1
}

cleanUP() {
    local JOB="front-job"
    local NAMESPACE="tacs"
    kubectl delete job $JOB -n $NAMESPACE
}

log() {
    echo -e "${1}"
}

logt() {
    log "\t ♥♥♥ ${1} \t ♥♥♥"
}

{

    logt "building backend and pusing image"
    buildAndPush be-futbol-matches $BACKEND
    buildAndPush bot-futbol-matches $BOT
    buildAndPush front-futbol-matches $FRONT

    logt "deploying kubernetes objects"
    apply $DEPLOY_MANIFEST
    
    logt "cleanUP environment"
    cleanUP
}