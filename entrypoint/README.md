
```
BUILD-KLOPAC RUNNER
sh builder.sh

OPENSOURCE
docker run --rm -v /var/run/docker.sock:/var/run/docker.sock runner --provision --validate

WEBSOCKET - HTTP
docker run --rm -v /var/run/docker.sock:/var/run/docker.sock runner --websocket

TEST
go test ./... -v

MOCK
mockgen -source=pkg/shell/Shell.go -destination=mock/mock_shell.go -package=mock


```
