package networks

import (
	"github.com/filecoin-project/go-state-types/abi"

	"github.com/filecoin-project/go-filecoin/internal/pkg/config"
)

type NetworkConf struct {
	Bootstrap config.BootstrapConfig
	Network   config.NetworkParamsConfig
}

func Testnet() *NetworkConf {

	return &NetworkConf{
		Bootstrap: config.BootstrapConfig{
			Addresses: []string{
				"/dns4/bootstrap-0.testnet.fildev.network/tcp/1347/p2p/12D3KooWJTUBUjtzWJGWU1XSiY21CwmHaCNLNYn2E7jqHEHyZaP7",
				"/dns4/bootstrap-1.testnet.fildev.network/tcp/1347/p2p/12D3KooW9yeKXha4hdrJKq74zEo99T8DhriQdWNoojWnnQbsgB3v",
				"/dns4/bootstrap-2.testnet.fildev.network/tcp/1347/p2p/12D3KooWCrx8yVG9U9Kf7w8KLN3Edkj5ZKDhgCaeMqQbcQUoB6CT",
				"/dns4/bootstrap-4.testnet.fildev.network/tcp/1347/p2p/12D3KooWPkL9LrKRQgHtq7kn9ecNhGU9QaziG8R5tX8v9v7t3h34",
				"/dns4/bootstrap-3.testnet.fildev.network/tcp/1347/p2p/12D3KooWKYSsbpgZ3HAjax5M1BXCwXLa6gVkUARciz7uN3FNtr7T",
				"/dns4/bootstrap-5.testnet.fildev.network/tcp/1347/p2p/12D3KooWQYzqnLASJAabyMpPb1GcWZvNSe7JDcRuhdRqonFoiK9W",
				"/dns4/lotus-bootstrap.forceup.cn/tcp/41778/p2p/12D3KooWFQsv3nRMUevZNWWsY1Wu6NUzUbawnWU5NcRhgKuJA37C",
				"/dns4/bootstrap-0.starpool.in/tcp/12757/p2p/12D3KooWGHpBMeZbestVEWkfdnC9u7p6uFHXL1n7m1ZBqsEmiUzz",
				"/dns4/bootstrap-1.starpool.in/tcp/12757/p2p/12D3KooWQZrGH1PxSNZPum99M1zNvjNFM33d1AAu5DcvdHptuU7u",
				"/dns4/node.glif.io/tcp/1235/p2p/12D3KooWBF8cpp65hp2u9LK5mh19x67ftAam84z9LsfaquTDSBpt",
				"/dns4/bootstrap-0.ipfsmain.cn/tcp/34721/p2p/12D3KooWQnwEGNqcM2nAcPtRR9rAX8Hrg4k9kJLCHoTR5chJfz6d",
				"/dns4/bootstrap-1.ipfsmain.cn/tcp/34723/p2p/12D3KooWMKxMkD5DMpSWsW7dBddKxKT7L2GgbNuckz9otxvkvByP",
			},
			MinPeerThreshold: 1,
			Period:           "30s",
		},
		Network: config.NetworkParamsConfig{
			ConsensusMinerMinPower: 10 << 40,
			ReplaceProofTypes: []int64{
				int64(abi.RegisteredSealProof_StackedDrg32GiBV1),
				int64(abi.RegisteredSealProof_StackedDrg64GiBV1),
			},
		},
	}
}
