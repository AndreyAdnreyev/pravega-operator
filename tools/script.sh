#! /bin/bash

if [ "$#" -ne 2 ]; then
	echo "Error : Invalid number of arguments"
	echo "Usage: ./script.sh namespace name"
	exit 1
fi


function UpgradingToPoperator(){
echo $1 $2
local namespace=$1

local name=$2
echo namespace is ${namespace}

sed -i "s/namespace.*/namespace: $namespace "/ ./webhook.yaml

kubectl create -f ./webhook.yaml

sed -i "s/name.*/name: ${name}-supported-upgrade-paths"/ ./version_map.yaml

kubectl create -f  ./version_map.yaml

kubectl create -f  ./secret.yaml

kubectl apply -f  ./operator.yaml

kubectl apply -f  ./crd.yaml

echo "helm install charts/bookkeeper --name bkop --namespace $namespace"

helm install charts/bookkeeper --name bkop --namespace $namespace

kubectl apply -f role.yaml


}

UpgradingToPoperator $1 $2