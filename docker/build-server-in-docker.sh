#!/bin/bash

set -e

PN="${BASH_SOURCE[0]##*/}"
PD="$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )"

renice 15 $$
cd "${PD}/.."

export GOBUILDER_BUILD_VERSION="$(cat web/package.json | grep '\"version\"' | grep -o '[0-9.]\+')"

gobuilder version -c ">=0.1.7" &>/dev/null || go get -u -v "github.com/tsaikd/gobuilder"
gobuilder --all --check --test --debug
