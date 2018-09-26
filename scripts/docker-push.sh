#!/bin/bash

set -e

CWD="$( cd "$( dirname "${BASH_SOURCE[0]}" )" >/dev/null && pwd )"
TAG=$(${CWD}/git-tag.sh)

docker push "gostudygroup/agenda-api:${TAG}"
