
OUTPUT_DIR="bin"

GOOS=linux GOARCH=amd64 go build -o "${OUTPUT_DIR}/main" main.go

zip "${OUTPUT_DIR}/main" "${OUTPUT_DIR}/main"