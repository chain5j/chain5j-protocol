// Package statetype
//
// @author: xwc1125
package statetype

import (
	"encoding/json"
	"fmt"
	"github.com/chain5j/chain5j-pkg/codec/rlp"
	"github.com/chain5j/chain5j-pkg/types"
	"github.com/chain5j/chain5j-pkg/util/hexutil"
	"io"
)

type Log struct {
	Address types.Address `json:"address" gencodec:"required"`
	Topics  []types.Hash  `json:"topics" gencodec:"required"`
	Data    []byte        `json:"data" gencodec:"required"`

	BlockHeight     uint64     `json:"block_height"`
	TransactionHash types.Hash `json:"tx_hash" gencodec:"required"`
	TxIndex         uint       `json:"tx_index" gencodec:"required"`
	BlockHash       types.Hash `json:"block_hash"`
	Index           uint       `json:"index" gencodec:"required"`

	Removed   bool   `json:"removed"`
	BlockTime uint64 `json:"block_time" gencodec:"required"`
}

type rlpLog struct {
	Address types.Address
	Topics  []types.Hash
	Data    []byte
}

func (l *Log) EncodeRLP(w io.Writer) error {
	return rlp.Encode(w, rlpLog{Address: l.Address, Topics: l.Topics, Data: l.Data})
}
func (l *Log) DecodeRLP(s *rlp.Stream) error {
	var dec rlpLog
	err := s.Decode(&dec)
	if err == nil {
		l.Address, l.Topics, l.Data = dec.Address, dec.Topics, dec.Data
	}
	return err
}

type Logs []*Log

func (logs Logs) Serialize() ([]byte, error) {
	return json.Marshal(logs)
}
func (logs Logs) Deserialize(d []byte) error {
	return json.Unmarshal(d, logs)
}

type rlpStorageLog struct {
	Address         types.Address
	Topics          []types.Hash
	Data            []byte
	BlockNumber     uint64
	TransactionHash types.Hash
	TxIndex         uint
	BlockHash       types.Hash
	Index           uint
}

// LogForStorage is a wrapper around a Log that flattens and parses the entire content of
// a log including non-consensus fields.
type LogForStorage Log

// EncodeRLP implements rlp.Encoder.
func (l *LogForStorage) EncodeRLP(w io.Writer) error {
	return rlp.Encode(w, rlpStorageLog{
		Address:         l.Address,
		Topics:          l.Topics,
		Data:            l.Data,
		BlockNumber:     l.BlockHeight,
		TransactionHash: l.TransactionHash,
		TxIndex:         l.TxIndex,
		BlockHash:       l.BlockHash,
		Index:           l.Index,
	})
}

// DecodeRLP implements rlp.Decoder.
func (l *LogForStorage) DecodeRLP(s *rlp.Stream) error {
	var dec rlpStorageLog
	err := s.Decode(&dec)
	if err == nil {
		*l = LogForStorage{
			Address:         dec.Address,
			Topics:          dec.Topics,
			Data:            dec.Data,
			BlockHeight:     dec.BlockNumber,
			TransactionHash: dec.TransactionHash,
			TxIndex:         dec.TxIndex,
			BlockHash:       dec.BlockHash,
			Index:           dec.Index,
		}
	}
	return err
}

func (l *Log) String() string {
	return fmt.Sprintf("{address=%s, data=%s, block_hash=%s, block_num=%d, tx_hash=%s, tx_index=%d, index=%d, topics=%v, block_time=%d}",
		l.Address.String(), hexutil.Encode(l.Data), l.BlockHash.String(), l.BlockHeight, l.TransactionHash.String(), l.TxIndex, l.Index, l.Topics, l.BlockTime)
}
