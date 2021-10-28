module github.com/onomyprotocol/tm-load-test

go 1.15

require (
	github.com/gorilla/websocket v1.4.2
	github.com/prometheus/client_golang v1.10.0
	github.com/puneetsingh166/onomy-sdk v0.43.1
	github.com/satori/go.uuid v1.2.0
	github.com/sirupsen/logrus v1.8.1
	github.com/spf13/cobra v1.1.3
	github.com/tendermint/tendermint v0.34.13
	golang.org/x/crypto v0.0.0-20201221181555-eec23a3978ad
)

replace github.com/gogo/protobuf => github.com/regen-network/protobuf v1.3.3-alpha.regen.1
