#!/bin/bash

set -e
set -x

function updaterepo {
    REPO=${1}
    LOCAL_REPO_PATH="/data/${REPO}"
    if ! [[ -d "${LOCAL_REPO_PATH}" ]]; then
        echo "No repo ${LOCAL_REPO_PATH} found"
        exit 1
    fi
    echo "Updating ${REPO}"
    find ${LOCAL_REPO_PATH} -name repodata | xargs -n 1 rm -rf
    time createrepo --update -s sha "${LOCAL_REPO_PATH}"

}


if [ "$1" = 'updaterepo' ]; then
    echo "[Run] Update repo"
    updaterepo $2
    exit 0
fi

echo "[RUN]: Builtin command not provided [updaterepo]"
echo "[RUN]: $@"

exec "$@"