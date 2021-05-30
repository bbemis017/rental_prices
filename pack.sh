
OUTPUT_DIR="bin"

rm "${OUTPUT_DIR}/main"
rm "${OUTPUT_DIR}/main.zip"

GOOS=linux GOARCH=amd64 go build -o "${OUTPUT_DIR}/main" main.go

cd "${OUTPUT_DIR}"

zip main main

cd ..