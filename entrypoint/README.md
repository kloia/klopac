# entrypoint godoc

## Getting Started

This is an example of how you may give instructions on setting up entrypoint locally.
To get a local copy up and running follow these simple example steps.

### Prerequisites

This is an example of how to list things you need to use the software and how to install them.

```
$ go version: go version go1.18
```

### Usage
  ```
  cd entrypoint
  go build cmd/main.go
  mv main ..
  cd -
  ./main --valuesFile ./values.yaml --varsPath ./vars
```

### Configuration Options

| Option      | Default              | Description                                                                     |
|-------------|----------------------|---------------------------------------------------------------------------------|
| provision   | false                | It executes provisioner                                                         |   
| validate    | false                | It executes both provisioner and validator                                      |   
| healthcheck | false                | It executes finalizer                                                           |   
| websocket   | false                | It helps to make use of websocket connection - required uri, username, password |   
| uri         | ""                   | websocket uri                                                                   |   
| username    | ""                   | username for websocket connection                                               |   
| password    | ""                   | password for websocket connection                                               |   
| logLevel    | INFO                 | It sets the level of the producing logs                                         |   
| logFile     | /data/entrypoint.log | It sets the filename for logs                                                   |  
| valuesFile  | /data/values.yaml    | Value File                                                                      |   
| dataPath    | /data/               | Data File Path                                                                  |   
| varsPath    | /data/vars           | Variable File                                                                   |   
| bundleFile  | /data/bundle.tar.gz        | Bundle File                                                                     |   


<!-- TABLE OF CONTENTS -->
<details>
  <summary>Packages(functions) used in entrypoint</summary>
  <ol>
    <li>
      <a href="/entrypoint/pkg/command/Command.go">Command</a>
    </li>
    <li>
      <a href="/entrypoint/pkg/flag/Flag.go">Flag</a>
    </li>
    <li><a href="/entrypoint/pkg/flow/Flow.go">Flow</a></li>
    <li><a href="/entrypoint/pkg/helper/Helper.go">Helper</a></li>
    <li><a href="/entrypoint/pkg/klopac/Klopac.go">Klopac</a></li>
    <li><a href="/entrypoint/pkg/option/Options.go">Options</a></li>
    <li><a href="/entrypoint/pkg/shell/Shell.go">Shell</a></li>
    <li><a href="/entrypoint/pkg/websocket/WebSocket.go">Websocket</a></li>
  </ol>
</details>