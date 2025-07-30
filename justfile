#!/usr/bin/env just --justfile

clean:
    cd go-api && go clean
    cd dotnet-api && dotnet clean

go_build:
    cd go-api && docker build -t golang-api:local .

dotnet_build:
    cd dotnet-api && docker build -t dotnet-api:local .

go_deploy:
    kubectl apply -f k8s/golang-deployment.yaml
    minikube service golang-api

dotnet_deploy:
    kubectl apply -f k8s/dotnet-deployment.yaml
    minikube service dotnet-api

go_restart:
    kubectl delete pod $(kubectl get pods | grep golang-api | awk '{print $1}')

dotnet_restart:
    kubectl delete pod $(kubectl get pods | grep dotnet-api | awk '{print $1}')

dashboard:
    minikube dashboard