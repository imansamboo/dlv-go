# Go Project Debugging with Delve

This project demonstrates how to set up remote debugging for a Go application using Delve.

## Prerequisites

- **Go**: Ensure that Go is installed on your machine. You can download and install Go from [https://golang.org/dl/](https://golang.org/dl/).
- **Delve**: Install Delve for Go debugging by running:

  ```bash
  go install github.com/go-delve/delve/cmd/dlv@latest
  ```

## Running the Delve Debug Server

To start the Delve debug server on your remote (or local) machine, use the following command. This will run Delve in headless mode, allowing you to connect to it remotely:

```bash
dlv debug main.go --headless --listen=:2345 --api-version=2
```

- **--headless**: Runs the Delve server without a user interface.
- **--listen=:2345**: The Delve server listens on port 2345.
- **--api-version=2**: Specifies the API version for compatibility.

## Attaching the Delve Client

Once the Delve server is running, you can attach to it from your local machine using the following command:

```bash
dlv connect <remote-ip>:2345
```

Replace `<remote-ip>` with the IP address of the machine running the Delve server (use `127.0.0.1` for local debugging).

## Common Delve Commands

Once you have attached the client, you can use the following commands in the Delve CLI:

- **Continue execution**:
  
  ```bash
  continue
  ```

- **Set a breakpoint** at a specific line:

  ```bash
  break main.go:19
  ```

- **Print the value of a variable** (e.g., `task1`):

  ```bash
  print task1
  ```

  or

  ```bash
  p task1
  ```

- **List all breakpoints**:

  ```bash
  breakpoints
  ```

- **Step through code one line at a time**:

  ```bash
  step
  ```

- **Next (step over a function call)**:

  ```bash
  next
  ```

- **List the current stack trace**:

  ```bash
  stack
  ```

- **Restart the program**:

  ```bash
  restart
  ```

- **Quit Delve**:

  ```bash
  quit
  ```

## Sample Debugging Session

1. Start the Delve server on your remote machine:
   ```bash
   dlv debug main.go --headless --listen=:2345 --api-version=2
   ```

2. Connect to the Delve server from your local machine:
   ```bash
   dlv connect <remote-ip>:2345
   ```

3. Set a breakpoint at line 19:
   ```bash
   break main.go:19
   ```

4. Continue program execution:
   ```bash
   continue
   ```

5. Once the breakpoint is hit, print the value of `task1`:
   ```bash
   print task1
   ```

6. Step through the code:
   ```bash
   step
   ```

7. Quit the debugging session:
   ```bash
   quit
   ```

## License

This project is licensed under the MIT License.
