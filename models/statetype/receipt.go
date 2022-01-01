// Package statetype
//
// @author: xwc1125
package statetype

import (
	"encoding/json"
	"errors"
	"io"
	"unsafe"

	"github.com/chain5j/chain5j-pkg/codec"
	"github.com/chain5j/chain5j-pkg/codec/rlp"
	"github.com/chain5j/chain5j-pkg/types"
	"github.com/chain5j/chain5j-pkg/util/hexutil"
)

var (
	receiptStatusFailedRLP     = []byte{}
	receiptStatusSuccessfulRLP = []byte{0x01}
)

const (
	// ReceiptStatusFailed is the status code of a transaction if execution failed.
	ReceiptStatusFailed = uint64(0)

	// ReceiptStatusSuccessful is the status code of a transaction if execution succeeded.
	ReceiptStatusSuccessful = uint64(1)
)

// Receipt represents the results of a transaction.
type Receipt struct {
	// Consensus fields
	Status            uint64 `json:"status"`
	CumulativeGasUsed uint64 `json:"cumulative_gas_used" gencodec:"required"`
	LogsBloom         Bloom  `json:"logs_bloom"         gencodec:"required"`
	Logs              []*Log `json:"logs"              gencodec:"required"`

	TransactionHash types.Hash    `json:"transaction_hash" gencodec:"required"`
	ContractAddress types.Address `json:"contract_address"`
	GasUsed         uint64        `json:"gas_used" gencodec:"required"`
}

// receiptRLP is the consensus encoding of a receipt.
type receiptRLP struct {
	Status            uint64
	CumulativeGasUsed uint64
	LogsBloom         Bloom
	Logs              []*Log
}

type receiptStorageRLP struct {
	Status            uint64
	CumulativeGasUsed uint64
	LogsBloom         Bloom
	TransactionHash   types.Hash
	ContractAddress   types.Address
	Logs              []*LogForStorage
	GasUsed           uint64
}

// NewReceipt creates a barebone transaction receipt, copying the init fields.
func NewReceipt(failed bool, cumulativeGasUsed uint64) *Receipt {
	r := &Receipt{CumulativeGasUsed: cumulativeGasUsed}
	if failed {
		r.Status = ReceiptStatusFailed
	} else {
		r.Status = ReceiptStatusSuccessful
	}
	return r
}

// EncodeRLP implements rlp.Encoder, and flattens the consensus fields of a receipt
// into an RLP stream. If no post state is present, byzantium fork is assumed.
func (r *Receipt) EncodeRLP(w io.Writer) error {
	return rlp.Encode(w, &receiptRLP{r.Status, r.CumulativeGasUsed, r.LogsBloom, r.Logs})
}

// DecodeRLP implements rlp.Decoder, and loads the consensus fields of a receipt
// from an RLP stream.
func (r *Receipt) DecodeRLP(s *rlp.Stream) error {
	var dec receiptRLP
	if err := s.Decode(&dec); err != nil {
		return err
	}

	r.Status, r.CumulativeGasUsed, r.LogsBloom, r.Logs = dec.Status, dec.CumulativeGasUsed, dec.LogsBloom, dec.Logs
	return nil
}

// Size returns the approximate memory used by all internal contents. It is used
// to approximate and limit the memory consumption of various caches.
func (r *Receipt) Size() types.StorageSize {
	size := types.StorageSize(unsafe.Sizeof(*r))

	size += types.StorageSize(len(r.Logs)) * types.StorageSize(unsafe.Sizeof(Log{}))
	for _, log := range r.Logs {
		size += types.StorageSize(len(log.Topics)*types.HashLength + len(log.Data))
	}
	return size
}

// ReceiptForStorage is a wrapper around a Receipt that flattens and parses the
// entire content of a receipt, as opposed to only the consensus fields originally.
type ReceiptForStorage Receipt

// EncodeRLP implements rlp.Encoder, and flattens all content fields of a receipt
// into an RLP stream.
func (r *ReceiptForStorage) EncodeRLP(w io.Writer) error {
	enc := &receiptStorageRLP{
		Status:            r.Status,
		CumulativeGasUsed: r.CumulativeGasUsed,
		LogsBloom:         r.LogsBloom,
		TransactionHash:   r.TransactionHash,
		ContractAddress:   r.ContractAddress,
		Logs:              make([]*LogForStorage, len(r.Logs)),
		GasUsed:           r.GasUsed,
	}
	for i, log := range r.Logs {
		enc.Logs[i] = (*LogForStorage)(log)
	}
	return rlp.Encode(w, enc)
}

// DecodeRLP implements rlp.Decoder, and loads both consensus and implementation
// fields of a receipt from an RLP stream.
func (r *ReceiptForStorage) DecodeRLP(s *rlp.Stream) error {
	var dec receiptStorageRLP
	if err := s.Decode(&dec); err != nil {
		return err
	}

	// Assign the consensus fields
	r.Status, r.CumulativeGasUsed, r.LogsBloom = dec.Status, dec.CumulativeGasUsed, dec.LogsBloom
	r.Logs = make([]*Log, len(dec.Logs))
	for i, log := range dec.Logs {
		r.Logs[i] = (*Log)(log)
	}
	// Assign the implementation fields
	r.TransactionHash, r.ContractAddress, r.GasUsed = dec.TransactionHash, dec.ContractAddress, dec.GasUsed
	return nil
}

// Receipts is a wrapper around a Receipt array to implement DerivableList.
type Receipts []*Receipt

// Len returns the number of receipts in this list.
func (r Receipts) Len() int { return len(r) }

// GetRlp returns the RLP encoding of one receipt from the list.
func (r Receipts) Item(i int) []byte {
	bytes, err := codec.Coder().Encode(r[i])
	if err != nil {
		panic(err)
	}
	return bytes
}
func (r Receipts) Key(i int) []byte {
	return r[i].TransactionHash.Bytes()
}

// MarshalJSON marshals as JSON.
func (r Receipt) MarshalJSON() ([]byte, error) {
	type Receipt struct {
		Status            hexutil.Uint64 `json:"status"`
		CumulativeGasUsed hexutil.Uint64 `json:"cumulative_gas_used" gencodec:"required"`
		LogsBloom         Bloom          `json:"logs_bloom"         gencodec:"required"`
		Logs              []*Log         `json:"logs"              gencodec:"required"`
		TransactionHash   types.Hash     `json:"transaction_hash" gencodec:"required"`
		ContractAddress   types.Address  `json:"contract_address"`
		GasUsed           hexutil.Uint64 `json:"gas_used" gencodec:"required"`
	}
	var enc Receipt
	enc.Status = hexutil.Uint64(r.Status)
	enc.CumulativeGasUsed = hexutil.Uint64(r.CumulativeGasUsed)
	enc.LogsBloom = r.LogsBloom
	enc.Logs = r.Logs
	enc.TransactionHash = r.TransactionHash
	enc.ContractAddress = r.ContractAddress
	enc.GasUsed = hexutil.Uint64(r.GasUsed)
	return json.Marshal(&enc)
}

// UnmarshalJSON unmarshals from JSON.
func (r *Receipt) UnmarshalJSON(input []byte) error {
	type Receipt struct {
		Status            *hexutil.Uint64 `json:"status"`
		CumulativeGasUsed *hexutil.Uint64 `json:"cumulative_gas_used" gencodec:"required"`
		LogsBloom         *Bloom          `json:"logs_bloom"         gencodec:"required"`
		Logs              []*Log          `json:"logs"              gencodec:"required"`
		TransactionHash   *types.Hash     `json:"transaction_hash" gencodec:"required"`
		ContractAddress   *types.Address  `json:"contract_address"`
		GasUsed           *hexutil.Uint64 `json:"gas_used" gencodec:"required"`
	}
	var dec Receipt
	if err := json.Unmarshal(input, &dec); err != nil {
		return err
	}
	if dec.Status != nil {
		r.Status = uint64(*dec.Status)
	}
	if dec.CumulativeGasUsed == nil {
		return errors.New("missing required field 'cumulativeGasUsed' for Receipt")
	}
	r.CumulativeGasUsed = uint64(*dec.CumulativeGasUsed)
	if dec.LogsBloom == nil {
		return errors.New("missing required field 'logsBloom' for Receipt")
	}
	r.LogsBloom = *dec.LogsBloom
	if dec.Logs == nil {
		return errors.New("missing required field 'logs' for Receipt")
	}
	r.Logs = dec.Logs
	if dec.TransactionHash == nil {
		return errors.New("missing required field 'transactionHash' for Receipt")
	}
	r.TransactionHash = *dec.TransactionHash
	if dec.ContractAddress != nil {
		r.ContractAddress = *dec.ContractAddress
	}
	if dec.GasUsed == nil {
		return errors.New("missing required field 'gasUsed' for Receipt")
	}
	r.GasUsed = uint64(*dec.GasUsed)
	return nil
}
