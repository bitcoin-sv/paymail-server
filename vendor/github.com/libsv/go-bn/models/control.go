package models

// ZMQNotification model.
type ZMQNotification struct {
	Notification string `json:"notification"`
	Address      string `json:"address"`
}

// Info model.
type Info struct {
	Version                      uint32  `json:"version"`
	ProtocolVersion              uint32  `json:"protocolversion"`
	Wallet                       uint32  `json:"wallet"`
	Balance                      float64 `json:"balance"`
	Blocks                       uint64  `json:"blocks"`
	TimeOffset                   uint32  `json:"timeoffset"`
	Connections                  uint32  `json:"connections"`
	Proxy                        string  `json:"proxy"`
	Difficulty                   float64 `json:"difficulty"`
	Testnet                      bool    `json:"testnet"`
	Stn                          bool    `json:"stn"`
	KeypoolOldest                uint32  `json:"keypoololdest"`
	KeypoolSize                  uint32  `json:"keypoolsize"`
	PayTxFee                     float64 `json:"paytxfee"`
	RelayFee                     float64 `json:"relayfee"`
	Errors                       string  `json:"errors"`
	MaxBlockSize                 uint64  `json:"maxblocksize"`
	MaxMinedBlockSize            uint64  `json:"maxminedblocksize"`
	MaxStackMemoryUsageConsensus uint64  `json:"maxstackmemoryusageconsensus"`
}

// MemoryInfo model.
type MemoryInfo struct {
	Locked struct {
		Used       uint32 `json:"used"`
		Free       uint32 `json:"free"`
		Total      uint32 `json:"total"`
		Locked     uint32 `json:"locked"`
		ChunksUsed uint32 `json:"chunks_used"`
		ChunksFree uint32 `json:"chunks_free"`
	} `json:"locked"`
	Preloading struct {
		ChainStateCached float64 `json:"chainStateCached"`
	} `json:"preloading"`
}

// Settings model.
type Settings struct {
	ExcessiveBlockSize              uint32  `json:"excessiveblocksize"`
	BlockMaxsIze                    uint32  `json:"blockmaxsize"`
	MaxTxSize                       uint32  `json:"maxtxsizepolicy"`
	MaxOrphanTxSize                 uint32  `json:"maxorphantxsize"`
	DataCarrierSize                 uint32  `json:"datacarriersize"`
	MaxScriptSize                   uint32  `json:"maxscriptsizepolicy"`
	MaxOpsPerScript                 uint32  `json:"maxopsperscriptpolicy"`
	MaxScriptNumLength              uint32  `json:"maxscriptnumlengthpolicy"`
	MaxPubKeysPerMultiSig           uint32  `json:"maxpubkeyspermultisigpolicy"`
	MaxTxSigOpsCounts               uint32  `json:"maxtxsigopscountspolicy"`
	MaxStackMemoryUsage             uint32  `json:"maxstackmemoryusagepolicy"`
	MaxStackMemoryUsageConsensus    uint32  `json:"maxstackmemoryusageconsensus"`
	LimitAncestorCount              uint32  `json:"limitancestorcount"`
	LimitCPFPGroupMembersCount      uint32  `json:"limitcpfpgroupmemberscount"`
	MaxMempool                      uint32  `json:"maxmempool"`
	MaxMempoolSizeDisk              uint32  `json:"maxmempoolsizedisk"`
	MempoolMaxPercentCPFP           uint32  `json:"mempoolmaxpercentcpfp"`
	AcceptNonStdOutputs             bool    `json:"acceptnonstdoutputs"`
	DataCarrier                     bool    `json:"datacarrier"`
	MinRelayTxFee                   float64 `json:"minrelaytxfee"`
	DustRelayFee                    float64 `json:"dustrelayfee"`
	DustLimitFactor                 uint32  `json:"dustlimitfactor"`
	BlockMinTxFee                   float64 `json:"blockmintxfee"`
	MaxStdTxValidationDuration      uint32  `json:"maxstdtxvalidationduration"`
	MaxNonStdTxValidationDuration   uint32  `json:"maxnonstdtxvalidationduration"`
	MaxTxChainValidationBudget      uint32  `json:"maxtxchainvalidationbudget"`
	ValidationClockCPU              bool    `json:"validationclockcpu"`
	MinConsolidationFactor          uint32  `json:"minconsolidationfactor"`
	MaxConsolidationInputScriptSize uint32  `json:"maxconsolidationinputscriptsize"`
	MinConfConsolidationInput       uint32  `json:"minconfconsolidationinput"`
	MinConsolidationInputMaturity   uint32  `json:"minconsolidationinputmaturity"`
	AcceptNonStdConsolidationInput  bool    `json:"acceptnonstdconsolidationinput"`
}
