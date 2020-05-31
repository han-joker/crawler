package module

type SummaryStruct struct {
	ID        MID
	Called    uint64      `json:"id"`
	Accepted  uint64      `json:"called"`
	Completed uint64      `json:"completed"`
	Handing   uint64      `json:"handding"`
	Extra     interface{} `json:"extra,omitempty"`
}
