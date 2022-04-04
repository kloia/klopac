
* go
  ```
  TEST
  go test ./... -v

  MOCK
  mockgen -source=pkg/shell/Shell.go -destination=mock/mock_shell.go -package=mock

  MANUAL TEST
  cd entrypoint
  go build cmd/main.go
  mv main ..
  ./main --valuesFile ../values.yaml --vars ../vars -> 

    ```