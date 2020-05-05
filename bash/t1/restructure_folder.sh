#!/bin/bash
set -e

PROGRAM="$0"

function restructure(){
folder=$1

if ! [[ -d "${folder}" ]]; then
  help
  exit 2
fi

cd "$folder"
for file in *; do
    if [[ -f "${file}" ]] && [[ "${file}" == *"."* ]]; then
        subfolder=$(echo "${file}" | cut -d. -f1)
        mkdir -p "$subfolder"
        mv "${file}" "$subfolder"
    fi
done
}

function usage() {
cat <<EOF
Usage: ${PROGRAM} DIRECTORY
       ${PROGRAM} [OPTION]
EOF
}

function help(){
cat <<EOF
$(usage)
  -h, --help                 give this help list

Moves all files in the DIRECTORY to a folder inside DIRECTORY
named as a part ot the file before the first period
EOF
}

if [[ "$1" == "--help" ]] || [[ "$1" == "-h" ]]; then
  help
  exit 0
fi

if [[ "$#" -ne 1 ]]; then
  help
  exit 1
fi

restructure "$1"

