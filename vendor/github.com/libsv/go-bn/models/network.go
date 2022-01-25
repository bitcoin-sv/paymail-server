package models

// ParamsDisconnectNode model.
type ParamsDisconnectNode struct {
	Address string
	ID      uint64
}

// Args convert struct into optional positional arguments.
func (p *ParamsDisconnectNode) Args() []interface{} {
	aa := []interface{}{p.Address}
	if p.ID != 0 {
		aa = append(aa, p.ID)
	}

	return aa
}

// NodeInfo model.
type NodeInfo struct {
	Node      string `json:"node"`
	Connected bool   `json:"connected"`
	Addresses []struct {
		Address   string `json:"address"`
		Connected string `json:"connected"`
	} `json:"addresses"`
}

// OptsNodeInfo options.
type OptsNodeInfo struct {
	Node string
}

// Args convert struct into optional positional arguments.
func (o *OptsNodeInfo) Args() []interface{} {
	return []interface{}{o.Node}
}

// ExcessiveBlock model.
type ExcessiveBlock struct {
	ExcessiveBlockSize uint64 `json:"excessiveBlockSize"`
}

// NetworkTotals model.
type NetworkTotals struct {
	TotalBytesReceived uint64 `json:"totalbytesrecv"`
	TotalBytesSent     uint64 `json:"totalbytessent"`
	TimeMilliseconds   uint64 `json:"timemillis"`
	UploadTarget       struct {
		Timeframe             uint64 `json:"timeframe"`
		Target                uint64 `json:"target"`
		TargetReached         bool   `json:"target_reached"`
		ServeHistoricalBlocks bool   `json:"serve_historical_blocks"`
		BytesRemainingInCycle uint64 `json:"bytes_left_in_cycle"`
		TimeRemainingInCycle  uint64 `json:"time_left_in_cycle"`
	} `json:"uploadtarget"`
}

// NetworkInfo model.
type NetworkInfo struct {
	Version                uint64 `json:"version"`
	Subversion             string `json:"subversion"`
	ProtocolVersion        uint64 `json:"protocolversion"`
	LocalServices          string `json:"localservices"`
	LocalRelay             bool   `json:"localrelay"`
	TimeOffset             uint64 `json:"timeoffset"`
	TxPropagationFrequency uint64 `json:"txnpropagationfreq"`
	TxPropagationLength    uint64 `json:"txnpropagationlen"`
	NetworkActive          bool   `json:"networkactive"`
	Connections            uint64 `json:"connections"`
	AddressCount           uint64 `json:"addresscount"`
	StreamPolicies         string `json:"streampolicies"`
	Networks               []struct {
		Name                      string `json:"name"`
		Limited                   bool   `json:"limited"`
		Reachable                 bool   `json:"reachable"`
		Proxy                     string `json:"proxy"`
		ProxyRandomiseCredentials bool   `json:"proxy_randomize_credentials"` // nolint:misspell // in response
	} `json:"networks"`
	RelayFee                        float64 `json:"relayfee"`
	MinConsolidationFactor          uint64  `json:"minconsolidationfactor"`
	MaxConsolidationInputScriptSize uint64  `json:"maxconsolidationinputscriptsize"`
	MinConfConsolidationInput       uint64  `json:"minconfconsolidationinput"`
	MinConsolidationInputMaturity   uint64  `json:"minconsolidationinputmaturity"`
	AcceptNonStdConsolidationInput  bool    `json:"acceptnonstdconsolidationinput"`
	LocalAddresses                  []struct {
		Address string `json:"address"`
		Port    int    `json:"port"`
		Score   uint64 `json:"score"`
	} `json:"localaddresses"`
	Warnings string `json:"warnings"`
}

// PeerInfo model.
type PeerInfo struct {
	ID            int    `json:"id"`
	Addr          string `json:"addr"`
	Services      string `json:"services"`
	RelayTxs      bool   `json:"relaytxes"`
	LaStsend      int    `json:"lastsend"`
	LastReceived  int    `json:"lastrecv"`
	SendSize      int    `json:"sendsize"`
	ReceivedSize  int    `json:"recvsize"`
	PauseSend     bool   `json:"pausesend"`
	UnpauseSend   bool   `json:"unpausesend"`
	BytesSent     int    `json:"bytessent"`
	BytesReceived int    `json:"bytesrecv"`
	AvgReceivedBW int    `json:"avgrecvbw"`
	AssocID       string `json:"associd"`
	StreamPolicy  string `json:"streampolicy"`
	Streams       []struct {
		StreamType       string `json:"streamtype"`
		LastSend         int    `json:"lastsend"`
		LastReceived     int    `json:"lastrecv"`
		BytesSent        int    `json:"bytessent"`
		BytesReceived    int    `json:"bytesrecv"`
		SendSize         int    `json:"sendsize"`
		ReceivedSize     int    `json:"recvsize"`
		SpotReceivedBW   int    `json:"spotrecvbw"`
		MinuteReceivedBW int    `json:"minuterecvbw"`
		PauseReceive     bool   `json:"pauserecv"`
	} `json:"streams"`
	ConnTime        int      `json:"conntime"`
	TimeOffset      int64    `json:"timeoffset"`
	PingTime        float64  `json:"pingtime"`
	MinPing         float64  `json:"minping"`
	Version         int      `json:"version"`
	SubVer          string   `json:"subver"`
	Inbound         bool     `json:"inbound"`
	AddNode         bool     `json:"addnode"`
	StartingHeight  int      `json:"startingheight"`
	TxInvSize       int      `json:"txninvsize"`
	BanScore        int      `json:"banscore"`
	SyncedHeaders   int      `json:"synced_headers"`
	SyncedBlocks    int      `json:"synced_blocks"`
	Inflight        []uint64 `json:"inflight"`
	Whitelisted     bool     `json:"whitelisted"`
	BytesSentPerMsg struct {
		FeeFilter   int `json:"feefilter"`
		Headers     int `json:"headers"`
		Ping        int `json:"ping"`
		Pong        int `json:"pong"`
		Protoconf   int `json:"protoconf"`
		SendHeaders int `json:"sendheaders"`
		VerACK      int `json:"verack"`
		Version     int `json:"version"`
	} `json:"bytessent_per_msg"`
	BytesReceivedPerMsg struct {
		GetAddress  int `json:"getaddr"`
		GetHeaders  int `json:"getheaders"`
		Inv         int `json:"inv"`
		Ping        int `json:"ping"`
		Pong        int `json:"pong"`
		SendHeaders int `json:"sendheaders"`
		VerACK      int `json:"verack"`
		Version     int `json:"version"`
	} `json:"bytesrecv_per_msg"`
}

// BannedSubnet model.
type BannedSubnet struct {
	Address     string `json:"address"`
	BannedUntil uint64 `json:"banned_until"`
	BanCreated  uint64 `json:"ban_created"`
	BanReason   string `json:"ban_reason"`
}

// OptsSetBan options.
type OptsSetBan struct {
	BanTime  uint64
	Absolute bool
}

// Args convert struct into optional positional arguments.
func (o *OptsSetBan) Args() []interface{} {
	return []interface{}{o.BanTime, o.Absolute}
}
