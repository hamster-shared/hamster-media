#!/opt/homebrew/bin/zsh

set -ex;

export hamster_media_version=$(date "+%Y%m%d%H%M%S")

#build
docker buildx build -t hamstershare/hamster-media:${hamster_media_version} --platform=linux/amd64 --push .

envsubst < deploy.yml | kubectl -n hamster apply -f -