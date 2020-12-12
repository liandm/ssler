# ssler

clone and run this command

```shell
./ssler --country="NL" --state="North Holland" --city="Amsterdam" --cn="Liandm CA" --o="Liandm, Ltd." --ou="IT Department" --d="localhost,127.0.0.1,::1,example.test,api.example.dev" create
```

```shell
curl -sSfL https://raw.githubusercontent.com/liandm/ssler/main/ssler | sh -s -- -b $(go env GOPATH)/bin v0.1
```