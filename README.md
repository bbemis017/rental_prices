# Apartment Notifier


## Local Development

### Configuration
example.envrc provides an example of the environment variables required to
run this application. Some variables such as SCRAPEIT_NET_KEY are secrets
and you may need to generate your own.

To setup your environment more easily you can use [direnv](https://direnv.net/). With direnv installed you should be able to
copy example.envrc to .envrc, input your own values for your environment, and tell direnv to export your variables:

```bash
cp example.envrc .envrc
direnv allow
```
With this in your bashrc you can export
the environment variables by calling

```bash
export_env example.env
```

### Running locally
Start lambda function locally
```bash
go run Main.go
```

call lambda function with client
```bash
go run client/client.go
```

## Manually Deploy to AWS Lambda

Build distributable zip file
```bash
./pack.sh
```

Upload zip file to [AWS Lambda](https://console.aws.amazon.com/)