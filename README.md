# Hazelcast Go Client Logger Sample

This sample demonstrates how to use a [zap](https://github.com/uber-go/zap) custom logger for the [Hazelcast Go Client](https://github.com/hazelcast/hazelcast-go-client).

## Prerequisites

* Go 1.15 or better.
* The example uses [Hazelcast Cloud](https://cloud.hazelcast.com/), which is free to use. But it can be adapted to use an on-prem Hazelcast instance. 

## Usage

1. Clone this repo:
   ```
   git clone https://github.com/yuce/hazelcast-go-client-logger-sample
   ```
2. Download the dependencies:
   ```
   go mod tidy
   ```
3. Update `clusterName` and `cloudToken` constants in `main.go` with the correct values.
4. Run the sample
   ```
   go build .
   ./hazelcast-go-client-logger-sample
   ``` 

## License

[Apache 2 License](https://github.com/yuce/hazelcast-go-client-logger-sample/blob/master/LICENSE).

Copyright (c) 2008-2021, Hazelcast, Inc. All Rights Reserved.
