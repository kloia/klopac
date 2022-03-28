# entrypoint godoc


<!-- PROJECT LOGO -->
<br />



<!-- TABLE OF CONTENTS -->
<details>
  <summary>Packages(functions) used in entrypoint</summary>
  <ol>
    <li>
      <a href="/pkg/command/Command.go">Command</a>
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


## Getting Started

This is an example of how you may give instructions on setting up entrypoint locally.
To get a local copy up and running follow these simple example steps.

### Prerequisites

This is an example of how to list things you need to use the software and how to install them.
* go
  ```sh
  git clone git@github.com:kloia/klopac.git && cd klopac

  BUILD-KLOPAC RUNNER
  sh builder.sh

  OPENSOURCE
  docker run --rm -v /var/run/docker.sock:/var/run/docker.sock runner --provision --validate

  WEBSOCKET - HTTP
  docker run --rm -v /var/run/docker.sock:/var/run/docker.sock runner --websocket --uri localhost:80

  TEST
  go test ./... -v

  MOCK
  mockgen -source=pkg/shell/Shell.go -destination=mock/mock_shell.go -package=mock

  MANUAL TEST
  go build cmd/main.go
  mv main ..
  ./main

    ```
<!-- CONTRIBUTING -->
## Contributing

Contributions are what make the open source community such an amazing place to learn, inspire, and create. Any contributions you make are **greatly appreciated**.

If you have a suggestion that would make this better, please fork the repo and create a pull request. You can also simply open an issue with the tag "enhancement".
Don't forget to give the project a star! Thanks again!

1. Fork the Project
2. Create your Feature Branch (`git checkout -b feature/AmazingFeature`)
3. Commit your Changes (`git commit -m 'Add some AmazingFeature'`)
4. Push to the Branch (`git push origin feature/AmazingFeature`)
5. Open a Pull Request

<p align="right">(<a href="#top">back to top</a>)</p>
