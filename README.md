# AGROWS Usage Example

This repository provides a complete example application using AGROWS for RPC over WebSockets with a Golang backend and a WebAssembly (WASM) client.

## Overview

The example demonstrates how to use AGROWS to define RPC functions in Golang, generate client and server code, and build a full application with development features like auto-reloading.

## Installation and Setup

### Prerequisites

- Go 1.22.5 or higher
- Node.js (for bun)
- Rust (for rustywind)
- Tailwind CSS

### Installation

1. Clone the repository:
    ```sh
    git clone https://github.com/codeupdateandmodificationsystem/agrows-usage-example.git
    cd agrows-usage-example
    ```

2. Install dependencies:
    ```sh
    go mod download
    npm install -g bun
    cargo install rustywind
    ```

## Usage

### Generating Client and Server Code

1. Generate the client and server code using the `agrows` CLI:
    ```sh
    just generate
    ```

### Running the Application

To start the application with auto-reloading:
```sh
just watch
```

### Example Functions

The example includes two RPC functions:

- `SayHello(name string) string`: Sends a greeting.
- `CrazyMath(inp CalcInput) string`: Performs some math operations.

The `CalcInput` type is defined as:

```go
type CalcInput struct {     
    A int     
    B int
}
```

### Web Interface

The web interface is built using HTMX and TypeScript. It provides a simple UI for interacting with the RPC functions.

### Development

The repository includes a `justfile` with helpful commands for development:

- `just watch`: Builds and runs the application with auto-reloading.
- `just generate`: Generates client and server code using AGROWS.

## Contributing

Contributions are welcome! Please fork the repository and submit a pull request.

## License

This project is licensed under the GPL license. See `LICENSE` for details.
