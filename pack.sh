#!/bin/bash

# exit when any command fails
set -e

OUTPUT_DIR="bin"

# Remove Output Directory if it exists
[[ -d "${OUTPUT_DIR}" ]] && rm -r "${OUTPUT_DIR}"

# Recreate Output Directory
mkdir "${OUTPUT_DIR}"

echo "===== CURRENT DIR ====="
ls -la

echo "===== BUILDING GOLANG ====="

GOOS=linux GOARCH=amd64 go build -o "${OUTPUT_DIR}/main" main.go

# zip up deployment archive
(
  echo "===== ZIPPING BINARIES ====="
  cd "${OUTPUT_DIR}" || exit
  zip deployment main
)

echo "===== BIN CONTENTS  ====="
ls -l bin
echo "===== ZIP CONTENTS  ====="
unzip -l bin/deployment.zip