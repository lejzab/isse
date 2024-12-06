#!/bin/bash

NAMESPACE="ise-mock"
MANIFEST_DIR="k8s"

echo "Usuwanie namespace: $NAMESPACE"
kubectl delete namespace $NAMESPACE --wait

echo "Tworzenie namespace: $NAMESPACE"
kubectl create namespace $NAMESPACE

echo "Stosowanie konfiguracji z katalogu: $MANIFEST_DIR"
kubectl -n $NAMESPACE apply -f $MANIFEST_DIR

echo "Proces zakończony."