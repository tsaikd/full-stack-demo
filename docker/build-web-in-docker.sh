#!/bin/bash

set -e

PN="${BASH_SOURCE[0]##*/}"
PD="$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )"

renice 15 $$
cd "${PD}/../web"

type yarn >/dev/null 2>&1 || npm install -g yarn
yarn install

npm run build
