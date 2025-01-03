#!/bin/bash

# Nazwy i tagi obrazów
LOCAL_IMAGE_NAME="ise-mock"
LOCAL_TAG_LATEST="latest"
VERSION_TAG="0.04"
REGISTRY_URL="registry.local:32000"
REMOTE_IMAGE_NAME="$REGISTRY_URL/$LOCAL_IMAGE_NAME:$VERSION_TAG"

echo "Budowanie obrazu Docker..."
docker build . \
  -t "$LOCAL_IMAGE_NAME:$LOCAL_TAG_LATEST" \
  -t "$LOCAL_IMAGE_NAME:$VERSION_TAG" \
  -t "$REMOTE_IMAGE_NAME"

if [ $? -eq 0 ]; then
  echo "Obraz zbudowany pomyślnie."
else
  echo "Błąd podczas budowania obrazu."
  exit 1
fi

echo "Pushing obraz do rejestru Docker..."
docker push "$REMOTE_IMAGE_NAME"

if [ $? -eq 0 ]; then
  echo "Obraz wypchnięty pomyślnie do rejestru: $REGISTRY_URL"
else
  echo "Błąd podczas pushu obrazu do rejestru."
  exit 1
fi