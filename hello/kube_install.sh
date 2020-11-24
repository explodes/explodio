#!/usr/bin/env bash
SERVICENAME='hello'

__abort() {
    echo "$@"
    exit 1
}

check_current_and_correct_context() {
    local -r CURRENT_CONTEXT="$(kubectl config view -o=jsonpath='{.current-context}')"
    local __change_context="Y"
    local __continue="n"
    if [ "$CURRENT_CONTEXT" != docker-desktop ] && [ "$CURRENT_CONTEXT" != minikube ]; then
        read  -p "Switch to docker-desktop? [$__change_context] " __change_context
        if [[ "${__change_context}" != "n" ]]; then
            kubectl config use-context docker-desktop
        else
            read  -p "You are using the context ($CURRENT_CONTEXT). Continue? [$__continue] " __continue
            if [[ "${__continue}" != "Y" ]]; then
                echo "Exiting"
                exit 0
            fi
        fi
    fi
}

build_docker_images() {
    local -r IMAGE_NAME=$1

    eval $(minikube docker-env)

    if [ -f Dockerfile."${IMAGE_NAME}" ]; then
        docker build -t explodio/"${IMAGE_NAME}" -f Dockerfile."${IMAGE_NAME}" . || __abort "failed to build image"
    fi

    if [ -f Dockerfile.flyway ]; then
        docker build -t explodio/"${IMAGE_NAME}"/flyway -f Dockerfile.flyway . || __abort "failed to build flyway image"
    fi
}

deploy_local() {
    local -r IMAGE_NAME=$1
    kustomize build deploy/local | kubectl apply -f - || __abort "failed to apply configuration"
}

check_current_and_correct_context || __abort "failed to apply context"
build_docker_images $SERVICENAME
deploy_local $SERVICENAME
