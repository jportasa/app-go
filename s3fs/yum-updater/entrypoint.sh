#!/bin/bash

set -e
set -x

function defaults {
    : ${SYNC_LOCAL_REPO_PREFIX="/data"}

    export SYNC_DEST SYNC_DELETE SYNC_REPOS_FILE SYNC_LOCAL_REPO_PREFIX SYNC_DRYRUN
}

function lock {
    REPO_PATH=$1
    LOCKFILE="${REPO_PATH}/lock"
    echo "Getting lock on ${LOCKFILE}"
    lockfile ${LOCKFILE}
    trap 'unlock ${REPO_PATH}' EXIT SIGINT SIGTERM SIGHUP
}

function unlock {
    REPO_PATH=$1
    LOCKFILE="${REPO_PATH}/lock"
    echo "Removing ${LOCKFILE}."
    rm -f ${LOCKFILE}
    trap - EXIT
}

function updaterepo {
    REPO=${1}
    LOCAL_REPO_PATH="${SYNC_LOCAL_REPO_PREFIX}/${REPO}"
    if ! [[ -d "${LOCAL_REPO_PATH}" ]]; then
        echo "No repo ${LOCAL_REPO_PATH} found"
        exit 1
    fi
    echo "Updating ${REPO}"
    lock ${LOCAL_REPO_PATH}

    echo "Lock acquired, updating repo"
    find ${LOCAL_REPO_PATH} -name repodata | xargs -n 1 rm -rf
    time createrepo --update -s sha "${LOCAL_REPO_PATH}"

    unlock ${LOCAL_REPO_PATH}
}


echo "HOME is ${HOME}"
echo "WHOAMI is `whoami`"

defaults


if [ "$1" = 'updaterepo' ]; then
    echo "[Run] Update repo"
    updaterepo $2
    exit 0
fi



echo "[RUN]: Builtin command not provided [updaterepo]"
echo "[RUN]: $@"

exec "$@"