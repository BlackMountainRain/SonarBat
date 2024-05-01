#!/bin/bash
set -eo pipefail

# for the forked repo
org_name=$1
image_path=$2
push=0
if [[ $3 == "--push" ]]; then
  push=1
fi

echo -e "Organization: $org_name"
echo -e "Image relative path: $image_path"

dockerfile_path="$image_path/Dockerfile"
info_path="$image_path/info"
if [ ! -f "$dockerfile_path" ]; then
  echo "No Dockerfile found"
  exit 1
fi
if [ ! -f "$info_path" ]; then
  echo "No build info found for Dockerfile"
  exit 1
fi

# load env
source "$info_path"

# override org name
if [[ -n "$org_name" ]]; then
  DOCKER_ORG=$(echo "$org_name" | tr '[:upper:]' '[:lower:]')
fi

tags=(latest)
# add version tags
if [[ -n $DOCKER_TAG ]]; then
  # check if ref name is a version number
  if [[ $DOCKER_TAG =~ ^v[0-9]+\.[0-9]+\.[0-9]+$ ]]; then
    major_version=$(echo "$DOCKER_TAG" | cut -d. -f1)
    minor_version=$(echo "$DOCKER_TAG" | cut -d. -f1,2)
    tags+=("$major_version" "$minor_version")
  fi
  tags+=("$DOCKER_TAG")
fi
echo "Tags: ${tags[*]}"

DOCKER_REPOSITORY=$DOCKER_REGISTRY/$DOCKER_ORG/$DOCKER_IMAGE
echo "Repo: $DOCKER_REPOSITORY"
echo "Base dir: $DOCKER_BASE_DIR"

args=""
for tag in "${tags[@]}"; do
  args+=" -t $DOCKER_REPOSITORY:$tag"
done
if [[ $push -eq 1 ]]; then
  args+=" --push"
fi

docker buildx build \
  $args \
  --platform linux/386,linux/amd64,linux/arm64/v8,linux/arm/v7 \
  --provenance=false \
  -f "$dockerfile_path" "$DOCKER_BASE_DIR"
