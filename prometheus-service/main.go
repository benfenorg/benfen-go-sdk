package main

import (
	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"net/http"
)

const ChainName = "test"

var bfcSystemMonitor = GlobalMonitor{
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
}

type GlobalMonitor struct {
	DaoVotingDelay       prometheus.Gauge
	DaoVotingPeriod      prometheus.Gauge
	DaoVotingQuorum      prometheus.Gauge
	DaoVotingActionDelay prometheus.Gauge
	VaultInfo            *prometheus.GaugeVec
	DexInfo              *prometheus.GaugeVec
	NftInfo              *prometheus.GaugeVec
	StakingInfo          *prometheus.GaugeVec
	DaoProposalCount     prometheus.Gauge
	DaoVotesCount        prometheus.Gauge
	DaoVotingBFCCount    prometheus.Gauge
}

func main() {
	// 创建一个自定义的注册表

	registry := prometheus.NewRegistry()

	// 创建一个简单呃 gauge 指标。
	DaoVotingDelay := prometheus.NewGauge(
		prometheus.GaugeOpts{
			Name: "bfc_dao_voting_delay",
			Help: "DaoVotingDelay-u64",
		},
	)

	DaoVotingPeriod := prometheus.NewGauge(
		prometheus.GaugeOpts{
			Name: "bfc_dao_voting_period",
			Help: "DaoVotingPeriod-u64",
		},
	)
	DaoVotingQuorum := prometheus.NewGauge(
		prometheus.GaugeOpts{
			Name: "bfc_dao_voting_quorum",
			Help: "DaoVotingQuorum-u64",
		},
	)
	DaoVotingActionDelay := prometheus.NewGauge(
		prometheus.GaugeOpts{
			Name: "bfc_dao_voting_action_delay",
			Help: "DaoVotingActionDelay-u64",
		},
	)

	DaoProposalCount := prometheus.NewGauge(
		prometheus.GaugeOpts{
			Name: "bfc_dao_proposal_count",
			Help: "DaoProposalCount-u64",
		},
	)

	DaoVotesCount := prometheus.NewGauge(
		prometheus.GaugeOpts{
			Name: "bfc_dao_votes_count",
			Help: "DaoVotesCount-u64",
		},
	)

	DaoVotingBFCCount := prometheus.NewGauge(
		prometheus.GaugeOpts{
			Name: "bfc_dao_voting_bfc_count",
			Help: "DaoVotingBFCCount-u64",
		},
	)

	VaultInfo := prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "vault_info",
			Help: "VaultInfo",
		},
		[]string{
			"vault_coinName",
			"vault_info",
		},
	)

	DexInfo := prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "Dex_info",
			Help: "DexInfo",
		},
		[]string{
			"dex_coinName",
			"dex_info",
		},
	)

	NftInfo := prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "Nft_info",
			Help: "NftInfo",
		},
		[]string{
			"nft",
			"type",
		},
	)

	StakingInfo := prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "Staking_info",
			Help: "StakingInfo",
		},
		[]string{
			"staking",
			"poolId",
			"objectId",
			"name",
			"type",
		},
	)

	// 使用我们自定义的注册表注册 gauge
	registry.MustRegister(DaoVotingDelay)
	registry.MustRegister(DaoVotingPeriod)
	registry.MustRegister(DaoVotingQuorum)
	registry.MustRegister(DaoVotingActionDelay)
	registry.MustRegister(VaultInfo)
	registry.MustRegister(NftInfo)
	registry.MustRegister(DexInfo)
	registry.Register(DaoProposalCount)
	registry.Register(DaoVotesCount)
	registry.Register(DaoVotingBFCCount)
	registry.Register(StakingInfo)

	bfcSystemMonitor.DaoVotingDelay = DaoVotingDelay
	bfcSystemMonitor.DaoVotingPeriod = DaoVotingPeriod
	bfcSystemMonitor.DaoVotingQuorum = DaoVotingQuorum
	bfcSystemMonitor.DaoVotingActionDelay = DaoVotingActionDelay
	bfcSystemMonitor.DaoProposalCount = DaoProposalCount
	bfcSystemMonitor.DaoVotesCount = DaoVotesCount
	bfcSystemMonitor.DaoVotingBFCCount = DaoVotingBFCCount
	bfcSystemMonitor.VaultInfo = VaultInfo
	bfcSystemMonitor.DexInfo = DexInfo
	bfcSystemMonitor.NftInfo = NftInfo
	bfcSystemMonitor.StakingInfo = StakingInfo
	bfcSystemMonitor.
		// 设置 gague 的值为 39
		DaoProposalCount.Set(3)

	go Start_httpservice()

	//add a timer to update the dao info

	// 暴露自定义指标
	http.Handle("/metrics", promhttp.HandlerFor(registry, promhttp.HandlerOpts{Registry: registry}))
	http.ListenAndServe(":9011", nil)
}

func GetBfcTreasureInfo() {
	//todo using time to update the treasure info
}

func getNFTStakeingInfo() {
	//todo
}

func Start_httpservice() {

	//Default返回一个默认的路由引擎
	r := gin.Default()
	handlers := func(c *gin.Context) {
		//输出json结果给调用方
		bfcSystemMonitor.DaoProposalCount.Set(3)
		c.JSON(
			200, gin.H{
				"message": "pong",
			},
		)
	}
	r.GET(
		"/ping", handlers,
	)

	r.GET(
		"/dao", func(c *gin.Context) {

			getDaoJob()
			c.JSON(
				200, gin.H{
					"message": "dao info",
				},
			)
		},
	)

	r.GET(
		"/vault", func(c *gin.Context) {
			getVaultJob()
			c.JSON(
				200, gin.H{
					"message": "vault info",
				},
			)
		},
	)

	r.GET(
		"/dex", func(c *gin.Context) {
			getDexJob()
			c.JSON(
				200, gin.H{
					"message": "dex info",
				},
			)
		},
	)

	r.GET(
		"/nft", func(c *gin.Context) {
			getNftJob()
			c.JSON(
				200, gin.H{
					"message": "nft info",
				},
			)
		},
	)

	r.GET(
		"/staking", func(c *gin.Context) {
			getStakingJob()
			c.JSON(
				200, gin.H{
					"message": "staking info",
				},
			)
		},
	)

	println("start httpservice")
	r.Run(":13344") // listen and serve on 0.0.0.0:8080

}
