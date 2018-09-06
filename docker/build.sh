#!/bin/bash

set -e

PN="${BASH_SOURCE[0]##*/}"
PD="$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )"

if [ "${1}" ]; then
	branch="-${1}"
else
	branch=""
fi

pushd "${PD}/.." >/dev/null

orgname="tsaikd"
projname="full-stack-demo"
repo="github.com/${orgname}/${projname}"
githash="$(git rev-parse HEAD | cut -c1-6)"
servename="${orgname}/${projname}:serve${branch}-$(date "+%Y%m%d-%H%M%S")-${githash}"
cachedir="/var/tmp/${orgname}-${projname}-cache"
buildgoimg="golang:1.11"
buildwebimg="tsaikd/node-yarn:latest"

function isDarwin() {
	[ "$(uname -s)" == "Darwin" ]
}

function timestamp() {
	date "+%Y-%m-%dT%H:%M:%S%z"
}

function cachetime() {
	if isDarwin ; then
		date -v-7d +%s
	else
		date +%s -d -7day
	fi
}

function modified() {
	if isDarwin ; then
		stat -f %m "$@"
	else
		stat -c %Y "$@"
	fi
}

if isDarwin ; then
	cachedir="/tmp/${orgname}-${projname}-cache"
fi

if [ -d "${cachedir}" ] ; then
	if [ "$(modified "${cachedir}")" -lt "$(cachetime)" ] ; then
		rm -rf "${cachedir}" || true
		docker pull "${buildgoimg}"
		docker pull "${buildwebimg}"
	fi
else
	docker pull "${buildgoimg}"
	docker pull "${buildwebimg}"
fi

echo "[$(timestamp)] build ${projname} web dist (${githash})"
docker run --rm \
	-e "VUE_APP_GITCOMMIT=${githash}" \
	-w "/${projname}" \
	-v "${PWD}:/${projname}" \
	-v "${cachedir}/node_modules:/${projname}/web/node_modules" \
	-v "${cachedir}/yarn:/usr/local/share/.cache/yarn" \
	-v "${cachedir}/npm:/root/.npm" \
	"${buildwebimg}" \
	"./docker/build-web-in-docker.sh"

echo "[$(timestamp)] build ${projname} server binary (${githash})"
docker run --rm \
	-e CGO_ENABLED=0 \
	-e HN_DBCONNURL="${HN_DBCONNURL}" \
	-e HN_ELASTIC="${HN_ELASTIC}" \
	-w "/go/src/${repo}" \
	-v "${PWD}:/go/src/${repo}" \
	-v "${cachedir}/go/src:/go/src" \
	-v "${cachedir}/go/bin:/go/bin" \
	"${buildgoimg}" \
	"./docker/build-server-in-docker.sh"

echo "[$(timestamp)] build ${projname} runtime docker (${githash})"
docker build --force-rm --pull -f docker/Dockerfile-serve -t "${servename}" .

if [ "${DOCKER_SERVE_NAME}" ] ; then
	docker tag "${servename}" "${DOCKER_SERVE_NAME}"
fi

echo "[$(timestamp)] ${projname} finished (${githash})"

popd >/dev/null
