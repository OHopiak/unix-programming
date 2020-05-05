#!/usr/bin/env bash

set -e
PROGRAM=$0

function main() {
  filename=$1
  if ! [[ -f "${filename}" ]]; then
    help
    exit 2
  fi

  sed 's/[^a-zA-Z \n]//g' "${filename}" \
  | awk '{print tolower($0)}' \
  | sed -e 's/\s/\n/g' \
  | sed -e '/^$/d' \
  | sort \
  | uniq -c \
  | sort -k1nr\
  | awk 'FNR <= 10 {print $0}'

}

function usage() {
cat <<EOF
Usage: ${PROGRAM} SRC
       ${PROGRAM} [OPTION]
EOF
}

function help(){
cat <<EOF
$(usage)
  -h, --help                 give this help list

Parses file SRC, changes all letters to lowercase, changes every blank space to the newline,
sorts the result, counts repeated words, sorts the count results descending by count,
show top 10 results.
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

main "$1"