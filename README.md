# Apartment Notifier


## Local Development

### Configuration
example.env provides an example of the environment variables required to
run this application. Some variables such as SCRAPEIT_NET_KEY are secrets
and you may need to generate your own.

To setup your environment more easily you can add this snippet to your .bashrc file

```bash
export_env() {
   temp_file="export_env.temp"
   cat "${1}" | grep -v "^$" | grep -v "^#" > ${temp_file}
   while IFS= read -r line; do
          export "${line}"
   done < ${temp_file}
   rm ${temp_file}
}
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