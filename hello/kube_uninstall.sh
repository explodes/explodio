#!/usr/bin/env bash
SERVICENAME='hello'

current_context() {
  kubectl config view -o=jsonpath='{.current-context}'
}

switch_context() {
  kubectl config use-context "${1}"
}

cur="$(current_context)"
__change_context="Y"
__continue="n"
if [ "$cur" != docker-desktop ] && [ "$cur" != minikube ]; then
    read  -p "Switch to docker-desktop? [$__change_context] " __change_context
    if [[ "${__change_context}" != "n" ]]; then
        switch_context docker-desktop
    else
        read  -p "You are using the context ($cur). Continue? [$__continue] " __continue
        if [[ "${__continue}" != "Y" ]]; then
            echo "Exiting"
            exit 0
        fi
    fi
fi

kustomize build deploy/local | sed -e 's/((COMPONENT_NAME))/'"${SERVICENAME}"'/g' | kubectl delete -f -
