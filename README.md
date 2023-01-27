## Examples of Golang Kernel TLS

This repo is to provide examples of using Golang Kernel TLS (https://github.com/secure-for-ai/goktls)
 
#### Echo
- Enable the Kernel TLS on Linux.
    ```bash
    sudo modprobe tls
    ```

- Disable KTLS
    ```bash
    # Disable the tls kernel module
    sudo modprobe -r tls
    ```
    Disable KTLS in golang.

    Set env `GOKTLS=0`, or omit it in the command.

- Run the server
    ```bash
    GOKTLS=1 go run server.go
    ```
    To print out the debug information, run with `-tags=debug`.
    ```bash
    GOKTLS=1 go run -tags=debug server.go
    ```

- Run the client
    ```bash
    GOKTLS=1 go run client.go
    ```
    To print out the debug information, run with `-tags=debug`.
    ```bash
    GOKTLS=1 go run -tags=debug client.go
    ```

- Run clients to test different ciphersuites and TLS versions.
    ```bash
    GOKTLS=1 go run client_test_tls_config.go
    ```
    To print out the debug information, run with `-tags=debug`.
    ```bash
    GOKTLS=1 go run -tags=debug client_test_tls_config.go
    ```
