# Apartment Notifier


### Local Development

Start lambda function locally
```bash
_LAMBDA_SERVER_PORT=8001 go run Main.go
```

call lambda function with client
```bash
go run client/client.go
```

### Manually Deploy to AWS Lambda

Build distributable zip file
```bash
./pack.sh
```

Upload zip file to [AWS Lambda](https://console.aws.amazon.com/)