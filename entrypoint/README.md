# entrypoint godoc


<!-- PROJECT LOGO -->
<br />



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

### 1- Command


Exec()
- We have our command function to make use of os.Exec() and here we create the template of the shell command.

### 2- Flag

Parse()
- We have Flag service and it helps us to parse command-line flags from os.Args[1:]

Bool()
- This function help us to make use of a boolean logic to decide whether we are going to use some sort of functions or not. It defines a bool flag with specified name.
  
String()
- Defines a string flag with specified name, default value, and usage string. The return value is the address of a string variable that stores the value of the flag

### 3- Flow


ExecuteCommand()
- It uses the shell type to execute commands.
  
Run()
- It basically take some sort of args like (provision, validate, healthCheck, logLevel) and depending to its value it execute relative yaml files.

### 4- Helper

ReadFile()
- Reads content of the yaml file and returns it

WriteFile()
- Writes content to a yaml file

Intersection()
- Basically we have two map and we compare them if there are some sort of values that should be changed according to its logic.

IntersectionHelper()
- Code blocks of the intersection logic. It compares two different map and runs the provided conditions.

contains()
- To check whether its content have the searched struct or not

UpdateValuesFile()
- Basically takes a interface and varsPath(which is path of the variable files) then it starts to override or leaves unchanged depending to intersection logic.



### 5- Klopac

Run()
- It might be considered as main function. It will execute some sort of code blocks depending to whether we are going to access klopac via a websocket or from command-line

### 6- Options

setFlags()
- It defines args' values and their specified names

### 7- Shell

Run()
- It runs the command string and return its result as output


### 8- Websocket

Enable()
- It creates a bidirectional channel with 1 buffer. And it also start to listen the websocket and basically what this function does is reads the messages that sent to channel. And it closes channel. After the close operation done. It uses select to define conditions.



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
