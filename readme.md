# gRPC API Testing for xray-core using Go

This repository contains test examples for interacting with the gRPC API provided by `xray-core` using Go. It demonstrates how to set up and execute tests to validate various gRPC commands available in `xray-core`.

## Prerequisites

- Go installed on your local environment.
- `xray-core` executable, configured and running locally.
- Protobuf/gRPC dependencies properly installed.

Ensure `xray-core` is running before starting the tests:

```sh
xray-core run -config ./config.json
```

## Configuration

To successfully run the tests, you need to have following configuration:

```json
"policy": {
    "system": {
        "statsInboundDownlink": true,
        "statsInboundUplink": true
    }
}
```

With debug logging enabled, you should see entries like:

```plaintext
[Debug] app/stats: create new counter inbound>>>api>>>traffic>>>uplink
[Debug] app/stats: create new counter inbound>>>api>>>traffic>>>downlink
```

## Tests Overview

The following test cases are covered in this project:

### 1. **GetSystemStats**
This test retrieves system stats using `GetSysStats()` and checks for:

- Number of Goroutines (`NumGoroutine`)
- Number of Garbage Collections (`NumGC`)
- Memory allocation statistics (`Alloc`, `TotalAlloc`, `Sys`, `Mallocs`, `Frees`, `LiveObjects`, `PauseTotalNs`, `Uptime`)

These stats must have non-zero values to ensure `xray-core` is actively tracking metrics.

### 2. **GetStats - Existing Counter**
Tests fetching statistics for an existing counter, such as:

- `"inbound>>>api>>>traffic>>>uplink"`

The test checks that the value of the counter is greater than zero, validating the presence of traffic data.

### 3. **GetStats - Non-Existing Counter**
Attempts to retrieve stats for a non-existing counter:

- `"non-existing-counter"`

An error is expected here, ensuring that the API correctly handles requests for invalid stats names.

### 4. **QueryStats - All Existing Counters**
Queries all existing counters using `QueryStats()`. This test ensures that:

- A non-empty list of stats is returned.
  
This is useful to confirm the system tracks various stats and they are accessible through the API.

### 5. **QueryStats - Non-Existing Counters**
Attempts to query counters with a specific pattern that doesn’t match any existing counter:

- `"non-existing-counter"`

This ensures the API correctly returns an empty list when no matching counters exist.

## Running Tests

To run the tests, you can use the standard Go testing command:

```sh
go test ./...
```

The tests are structured in a way that each one can be run independently or as part of the complete suite.

## Helper Functions

The tests rely on helper functions to set up and tear down the gRPC connection:

- `prepareGrpcConn(t, ctx, address)`: Prepares the gRPC connection to the specified address.
- `prepareStatsClient(t, conn)`: Sets up a stats service client to interact with the `xray-core` stats API.

These helper functions are designed to keep the test code clean and reusable.

## Dependencies

- [Testify](https://github.com/stretchr/testify) - A powerful Go toolkit that provides many useful assertions for testing.
- `xray-core` gRPC stubs, generated from the `.proto` files included in the project.

## License

This project is open source and licensed under the [MIT License](LICENSE).

## Contributing

Feel free to open issues or submit PRs if you find bugs or have suggestions for additional features or test cases.

## Author

Timofei Belanenko
```