// Package models
//
// @author: xwc1125
package models

import (
	"encoding/json"
	"github.com/chain5j/chain5j-pkg/codec"
	"github.com/chain5j/chain5j-pkg/codec/rlp"
	"github.com/chain5j/chain5j-pkg/crypto/hashalg"
	"github.com/chain5j/chain5j-pkg/crypto/signature"
	"github.com/chain5j/chain5j-pkg/types"
	"github.com/chain5j/chain5j-pkg/util/hexutil"
	"github.com/chain5j/chain5j-protocol/models/statetype"
	"github.com/chain5j/logger"
	"io"
	"sync/atomic"
)

// Header 区块header
type Header struct {
	ParentHash types.Hash `json:"parent_hash"` // 父Hash
	Height     uint64     `json:"height"`      // 区块高度
	Timestamp  uint64     `json:"timestamp"`   // 区块时间戳[毫秒]
	StateRoots []byte     `json:"state_roots"` // 状态根

	// 子区块内容
	SubBlockIndex uint64     `json:"sub_block_index"`           // 子区块顺序
	SubBlockRoot  types.Hash `json:"sub_block_root,omitempty"`  // 子区块根
	SubBlockCount uint64     `json:"sub_block_count,omitempty"` // 子区块个数

	// 交易内容
	TxsRoot      []types.Hash     `json:"txs_root"`                       // 交易树根数组(多种交易类型，存在多种交易根)
	TxsCount     uint64           `json:"txs_count"`                      // 交易个数
	ReceiptsRoot []types.Hash     `json:"receipts_root,omitempty"`        // 交易收据根
	LogsBloom    *statetype.Bloom `json:"logs_bloom,omitempty" rlp:"nil"` // logs的bloom

	// gas消耗内容
	GasUsed   uint64     `json:"gas_used"`                      // 总的区块消耗gas
	GasLimit  uint64     `json:"gas_limit"`                     // 区块最大gas
	Consensus *Consensus `json:"consensus,omitempty" rlp:"nil"` // 共识内容

	ArchiveHash *types.Hash           `json:"archive_hash,omitempty" rlp:"nil"` // 区块归档区块hash
	Extra       []byte                `json:"extra,omitempty"`                  // 扩展内容
	Memo        []byte                `json:"memo,omitempty"`                   // 备忘录
	Signature   *signature.SignResult `json:"signature" rlp:"nil"`              // 签名数据
}

// GenesisExtra 创世区块的扩展内容。此内容在创世块中，填充区块的Extra
type GenesisExtra struct {
	RawExtra     []byte `json:"raw_extra"`     // 原始扩展内容
	GenesisBytes []byte `json:"genesis_bytes"` // 创世文件内容
}

// Hash 区块头hash
func (h *Header) Hash() types.Hash {
	header := CopyHeader(h)
	if header.Signature == nil {
		logger.Crit("header sign is nil")
		return types.EmptyHash
	}
	// 共识时，需要将consensus的内容设置成不受限制
	header.Consensus = nil
	rlpHash, err := hashalg.RlpHash(header)
	if err != nil {
		logger.Error("rlp hash err", "err", err)
		return types.EmptyHash
	}
	return rlpHash
}

// HashNoSign 未签名的hash，用于需要签名时的处理
func (h *Header) HashNoSign() types.Hash {
	header := CopyHeader(h)
	// 共识时，需要将consensus的内容设置成不受限制
	header.Consensus = nil
	header.Signature = nil
	rlpHash, err := hashalg.RlpHash(header)
	if err != nil {
		logger.Error("rlp hash err", "err", err)
		return types.EmptyHash
	}
	return rlpHash
}

// Copy 拷贝
func (h *Header) Copy() *Header {
	return CopyHeader(h)
}

// CopyHeader 创建header的深拷贝，避免副作用修改header的值
func CopyHeader(h *Header) *Header {
	cpy := *h
	if len(h.StateRoots) > 0 {
		cpy.StateRoots = make([]byte, len(h.StateRoots))
		copy(cpy.StateRoots, h.StateRoots)
	}
	if h.ArchiveHash != nil {
		cpy.ArchiveHash = &(*h.ArchiveHash)
	}
	if h.Consensus != nil {
		cpy.Consensus = h.Consensus.Copy()
	}
	if len(h.Extra) > 0 {
		cpy.Extra = make([]byte, len(h.Extra))
		copy(cpy.Extra, h.Extra)
	}
	if len(h.Memo) > 0 {
		cpy.Memo = make([]byte, len(h.Memo))
		copy(cpy.Memo, h.Memo)
	}
	if h.Signature != nil {
		cpy.Signature = h.Signature.Copy()
	}
	return &cpy
}

// headerMarshal hash序列化对象
type headerMarshal struct {
	ParentHash types.Hash     `json:"parent_hash,omitempty"`
	Height     uint64         `json:"height,omitempty"`
	Timestamp  uint64         `json:"timestamp,omitempty"`
	StateRoots *hexutil.Bytes `json:"state_roots,omitempty"`

	SubBlockIndex uint64     `json:"sub_block_index,omitempty"`
	SubBlockRoot  types.Hash `json:"sub_block_root,omitempty"`
	SubBlockCount uint64     `json:"sub_block_count,omitempty"`

	TxsRoot      []types.Hash     `json:"txs_root,omitempty"`
	TxsCount     uint64           `json:"txs_count,omitempty"`
	ReceiptsRoot []types.Hash     `json:"receipts_root,omitempty"`
	LogsBloom    *statetype.Bloom `json:"logs_bloom,omitempty" rlp:"nil"`

	GasUsed   uint64     `json:"gas_used,omitempty"`
	GasLimit  uint64     `json:"gas_limit,omitempty"`
	Consensus *Consensus `json:"consensus,omitempty"`

	ArchiveHash *types.Hash           `json:"archive_hash,omitempty"`
	Extra       hexutil.Bytes         `json:"extra,omitempty"`
	Memo        hexutil.Bytes         `json:"memo,omitempty"`
	Signature   *signature.SignResult `json:"signature,omitempty"`
}

// MarshalJSON json序列化
func (h Header) MarshalJSON() ([]byte, error) {
	var enc headerMarshal
	enc.ParentHash = h.ParentHash
	enc.Height = h.Height
	enc.Timestamp = h.Timestamp
	if h.StateRoots != nil {
		stateRootBytes := hexutil.Bytes(h.StateRoots)
		enc.StateRoots = &stateRootBytes
	}

	enc.SubBlockIndex = h.SubBlockIndex
	enc.SubBlockCount = h.SubBlockCount
	enc.SubBlockRoot = h.SubBlockRoot

	enc.TxsRoot = h.TxsRoot
	enc.TxsCount = h.TxsCount
	enc.ReceiptsRoot = h.ReceiptsRoot
	enc.LogsBloom = h.LogsBloom

	enc.GasLimit = h.GasLimit
	enc.GasUsed = h.GasUsed
	enc.Consensus = h.Consensus

	enc.ArchiveHash = h.ArchiveHash
	enc.Extra = h.Extra
	enc.Memo = h.Memo
	enc.Signature = h.Signature
	return json.Marshal(&enc)
}

// UnmarshalJSON json反序列化
func (h *Header) UnmarshalJSON(input []byte) error {
	var dec headerMarshal
	if err := json.Unmarshal(input, &dec); err != nil {
		return err
	}
	h.ParentHash = dec.ParentHash
	h.Height = dec.Height
	h.Timestamp = dec.Timestamp
	if dec.StateRoots != nil {
		h.StateRoots = dec.StateRoots.Bytes()
	}

	h.SubBlockIndex = dec.SubBlockIndex
	h.SubBlockCount = dec.SubBlockCount
	h.SubBlockRoot = dec.SubBlockRoot

	h.TxsRoot = dec.TxsRoot
	h.TxsCount = dec.TxsCount
	h.ReceiptsRoot = dec.ReceiptsRoot
	h.LogsBloom = dec.LogsBloom

	h.GasUsed = dec.GasUsed
	h.GasLimit = dec.GasLimit
	h.Consensus = dec.Consensus

	h.ArchiveHash = dec.ArchiveHash
	h.Extra = dec.Extra
	h.Memo = dec.Memo
	h.Signature = dec.Signature
	return nil
}

// Block 区块
type Block struct {
	header       *Header      `json:"header"`       // 区块header
	transactions Transactions `json:"transactions"` // 交易内容
	subBlocks    *Blocks      `json:"blocks"`       // 子区块集

	hash       atomic.Value // header hash cache
	hashNoSign atomic.Value // header no sign hash cache
	size       atomic.Value // 区块 size cache
}

// BlockAbstract 区块的摘要
type BlockAbstract struct {
	Hash   types.Hash
	Height uint64
}

// Body 区块交易集合[可变的、不安全的]
type Body struct {
	Height uint64       `json:"height"`       // 区块高度
	Txs    Transactions `json:"transactions"` // 交易集合
}

// NewBlock 创建一个新的block。传入的数据都是通过深度copy进行赋值。
// 外部对header，txs进行修改不会影响到block
// block未被签名，需要进行签名
func NewBlock(header *Header, txs Transactions, subBlocks *Blocks) *Block {
	b := &Block{header: CopyHeader(header)}
	if txs == nil || txs.Len() == 0 {
		var emptyTxs Transactions
		b.header.TxsRoot = emptyTxs.TxsRoot()
		b.header.TxsCount = 0
		b.transactions = emptyTxs
	} else {
		b.header.TxsRoot = txs.TxsRoot()
		b.header.TxsCount = uint64(txs.Len())
		b.transactions = txs.DeepCopy()
	}
	if subBlocks == nil || subBlocks.Len() == 0 {
		emptyBlocks := new(Blocks)
		b.header.SubBlockRoot = emptyBlocks.BlocksRoot()
		b.header.SubBlockCount = 0
		b.subBlocks = emptyBlocks
	} else {
		b.header.SubBlockRoot = subBlocks.BlocksRoot()
		b.header.SubBlockCount = uint64(subBlocks.Len())
		b.subBlocks = subBlocks.DeepCopy()
	}

	return b
}

// WithSeal 创建一个新的block，新的block使用传入的header替换，txs使用对象的。
// 被consensus seal
func (b *Block) WithSeal(header *Header) *Block {
	cpy := *header

	return &Block{
		header:       &cpy,
		transactions: b.transactions.DeepCopy(),
		subBlocks:    b.subBlocks.DeepCopy(),
	}
}

// Hash 获取区块hash【已签名】
func (b *Block) Hash() types.Hash {
	if hash := b.hash.Load(); hash != nil {
		return hash.(types.Hash)
	}
	hash := b.header.Hash()
	b.hash.Store(hash)
	return hash
}

// HashNoSign 获取未签名的hash【未签名】
func (b *Block) HashNoSign() types.Hash {
	if hash := b.hashNoSign.Load(); hash != nil {
		return hash.(types.Hash)
	}
	hash := b.header.HashNoSign()
	b.hashNoSign.Store(hash)
	return hash
}

// Transaction 根据hash从block的交易中获取交易对象
func (b *Block) Transaction(hash types.Hash) Transaction {
	for _, txs := range b.transactions.Data() {
		for _, tx := range txs {
			if tx.Hash() == hash {
				return tx
			}
		}
	}
	return nil
}
func (b *Block) Block(hash types.Hash) *Block {
	for _, block := range b.subBlocks.Data() {
		if block.Hash() == hash {
			return block
		}
	}
	return nil
}
func (b *Block) ParentHash() types.Hash   { return b.header.ParentHash }
func (b *Block) Height() uint64           { return b.header.Height }
func (b *Block) StateRoots() []byte       { return b.header.StateRoots }
func (b *Block) TxsRoot() []types.Hash    { return b.header.TxsRoot }
func (b *Block) TxsCount() uint64         { return b.header.TxsCount }
func (b *Block) Timestamp() uint64        { return b.header.Timestamp }
func (b *Block) GasUsed() uint64          { return b.header.GasUsed }
func (b *Block) GasLimit() uint64         { return b.header.GasLimit }
func (b *Block) Consensus() *Consensus    { return b.header.Consensus.Copy() }
func (b *Block) SubBlockIndex() uint64    { return b.header.SubBlockIndex }
func (b *Block) SubBlockCount() uint64    { return b.header.SubBlockCount }
func (b *Block) SubBlockRoot() types.Hash { return b.header.SubBlockRoot }
func (b *Block) Extra() []byte            { return hexutil.CopyBytes(b.header.Extra) }
func (b *Block) Memo() []byte             { return hexutil.CopyBytes(b.header.Memo) }
func (b *Block) Signature() *signature.SignResult {
	if b.header.Signature != nil {
		signature := &signature.SignResult{
			Name: b.header.Signature.Name,
		}
		if !b.header.Signature.PubKey.Nil() {
			signature.PubKey = make([]byte, len(b.header.Signature.PubKey))
			copy(signature.PubKey, b.header.Signature.PubKey)
		}
		if len(b.header.Signature.Signature) > 0 {
			signature.Signature = make([]byte, len(b.header.Signature.Signature))
			copy(signature.Signature, b.header.Signature.Signature)
		}
		return signature
	}
	return nil
}
func (b *Block) Header() *Header            { return CopyHeader(b.header) }
func (b *Block) Transactions() Transactions { return b.transactions }
func (b *Block) SubBlocks() *Blocks         { return b.subBlocks }
func (b *Block) Body() *Body                { return &Body{b.Height(), b.transactions} }
func (b *Block) Size() types.StorageSize {
	if size := b.size.Load(); size != nil {
		return size.(types.StorageSize)
	}
	c := types.WriteCounter(0)
	bytes, _ := codec.Coder().Encode(b)
	c.Write(bytes)
	b.size.Store(types.StorageSize(c))
	return types.StorageSize(c)
}
func (b *Block) SetStateRoots(stateRoots []byte) {
	b.header.StateRoots = stateRoots
}
func (b *Block) SetTransactions(transactions Transactions) {
	b.transactions = transactions
}
func (b *Block) SetConsensus(consensus *Consensus) {
	b.header.Consensus = consensus
}

// extBlock 序列化header("external")
type extBlock struct {
	Header       *Header      `json:"header"`                           // 区块header
	Transactions Transactions `json:"transactions,omitempty" rlp:"nil"` // 交易内容
	SubBlocks    *Blocks      `json:"blocks,omitempty" rlp:"nil"`       // 子区块集
}

func (b *Block) EncodeRLP(w io.Writer) error {
	return rlp.Encode(w, extBlock{
		Header:       b.header,
		Transactions: b.transactions,
		SubBlocks:    b.subBlocks,
	})
}
func (b *Block) DecodeRLP(s *rlp.Stream) error {
	var eb extBlock
	_, size, _ := s.Kind()
	if err := s.Decode(&eb); err != nil {
		return err
	}
	b.header = eb.Header
	b.transactions = eb.Transactions
	b.subBlocks = eb.SubBlocks
	//if b.transactions == nil {
	//	b.transactions = new(Transactions)
	//}
	if b.subBlocks == nil {
		b.subBlocks = new(Blocks)
	}
	b.size.Store(types.StorageSize(rlp.ListSize(size)))
	return nil
}
func (b *Block) MarshalJSON() ([]byte, error) {
	bJson := extBlock{
		Header: b.header.Copy(),
	}
	if b.transactions != nil {
		bJson.Transactions = b.transactions.DeepCopy()
	}
	if b.subBlocks != nil {
		bJson.SubBlocks = b.subBlocks.DeepCopy()
	}
	return json.Marshal(bJson)
}
func (b *Block) UnmarshalJSON(input []byte) error {
	var bJson extBlock
	err := json.Unmarshal(input, &bJson)
	if err != nil {
		return err
	}
	block := NewBlock(bJson.Header, bJson.Transactions, bJson.SubBlocks)
	*b = *block
	return nil
}

// Blocks 交易集合
type Blocks struct {
	data []*Block
}

func NewBlocks(blocks []*Block) *Blocks {
	return &Blocks{
		data: blocks,
	}
}
func (bs *Blocks) DeepCopy() *Blocks {
	var newBlocks Blocks
	newBlocks.data = make([]*Block, bs.Len())
	copy(newBlocks.data, bs.data)
	return &newBlocks
}
func (bs *Blocks) Hashes() []string {
	if bs == nil {
		return nil
	}
	hashes := make([]string, len(bs.data))
	for i, tx := range bs.data {
		hashes[i] = tx.Hash().Hex()
	}
	return hashes
}
func (bs *Blocks) Data() []*Block {
	return bs.data
}
func (bs *Blocks) Get(i int) *Block {
	if i >= bs.Len() || i < 0 {
		log().Warn("blocks getIndex out of bounds")
		return nil
	}
	return bs.data[i]
}
func (bs *Blocks) Key(i int) []byte {
	if i >= bs.Len() || i < 0 {
		log().Warn("blocks getIndex out of bounds")
		return nil
	}
	return bs.data[i].Hash().Bytes()
}
func (bs *Blocks) Item(i int) []byte {
	if i >= bs.Len() || i < 0 {
		log().Warn("blocks getIndex out of bounds")
		return nil
	}
	bytes, _ := codec.Coder().Encode(bs.data[i])
	return bytes
}
func (bs *Blocks) BlocksRoot() types.Hash {
	return hashalg.RootHash(bs)
}
func (bs *Blocks) Len() int {
	return len(bs.data)
}
func (bs *Blocks) Less(i, j int) bool {
	if bs.data[i].Height() < bs.data[j].Height() {
		return true
	}
	if bs.data[i].SubBlockIndex() < bs.data[j].SubBlockIndex() {
		return true
	} else {
		return false
	}
}
func (bs *Blocks) Swap(i, j int) {
	bs.data[i], bs.data[j] = bs.data[j], bs.data[i]
}
func (bs *Blocks) EncodeRLP(w io.Writer) error {
	var temp = make([][]byte, 0, bs.Len())
	for _, b := range bs.data {
		bytes, err := rlp.EncodeToBytes(b)
		if err != nil {
			log().Info("block rlp encode err", "err", err)
			break
		}
		temp = append(temp, bytes)
	}

	return rlp.Encode(w, temp)
}
func (bs *Blocks) DecodeRLP(s *rlp.Stream) error {
	var temp = make([][]byte, 0)
	err := s.Decode(&temp)
	if err != nil {
		log().Error("blocks rlp decode err", "err", err)
		return err
	}
	var block = new(Block)
	for _, bBytes := range temp {
		err = rlp.DecodeBytes(bBytes, block)
		if err != nil {
			log().Error("rlp decode blockBytes error", "err", err)
			continue
		}
		bs.data = append(bs.data, block)
	}

	return nil
}
