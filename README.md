# pusher

A Prometheus push example with TelmetryTower.

TelemetryTower tsdb implement Prometheus PushGateway protocol.

## How to run


Before to run code, update the real basic auth account [main.go#L54](https://github.com/telemetrytower/pusher/blob/master/main.go#L54) please.

```
git clone https://github.com/telemetrytower/pusher.git
cd pusher
go mod tidy && go mod vendor 
go run main.go
```
