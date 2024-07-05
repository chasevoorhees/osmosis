package domain

import "time"

type Transaction struct {
	Height             uint64        `json:"height"`
	BlockTime          time.Time     `json:"timestamp"`
	TransactionType    string        `json:"tx_type"`
	TransactionHash    string        `json:"tx_hash"`
	TransactionIndexId int           `json:"tx_index_id"`
	EventIndexId       int           `json:"event_index_id"`
	Events             []interface{} `json:"events"`
	IngestedAt         time.Time     `json:"ingested_at"`
}
