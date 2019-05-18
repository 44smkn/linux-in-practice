#!/usr/bin/env bash

usage_exit() {
        echo "Usage: $0 [-d dir]" 1>&2
        exit 1
}

while getopts d:h OPT
do
    case $OPT in
        d)  VALUE_D=$OPTARG
            ;;
        h)  usage_exit
            ;;
        \?) usage_exit
            ;;
    esac
done

shift $((OPTIND - 1))

BASE_DIR=$(dirname ${0})
TARGET_DIR=${BASE_DIR}/${VALUE_D}
if [[ ! -e ${TARGET_DIR} ]]; then
    echo "${TARGET_DIR} is not exists."
fi

cp -f ${BASE_DIR}/Dockerfile ${TARGET_DIR}
cp -f ${BASE_DIR}/go.mod ${TARGET_DIR}
TAG=$(echo ${VALUE_D} | awk -F'[/]' '{print $2}')
docker build -t linux-in-practice:${TAG} ${TARGET_DIR}/ > /dev/null
docker run --rm --cap-add=SYS_PTRACE --security-opt="seccomp=unconfined" linux-in-practice:${TAG}
rm ${TARGET_DIR}/Dockerfile
rm ${TARGET_DIR}/go.mod

echo "completed"