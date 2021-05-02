
OUTPUT_DIR="bin"

rm "${OUTPUT_DIR}/*"

GOOS=linux GOARCH=amd64 go build -o "${OUTPUT_DIR}/main" main.go

cd "${OUTPUT_DIR}"

zip main main

cd ..