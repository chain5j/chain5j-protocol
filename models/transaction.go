// Package models
//
// @author: xwc1125
package models

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"math/big"
	"reflect"
	"sort"
	"sync"

	"github.com/chain5j/chain5j-pkg/codec"
	"github.com/chain5j/chain5j-pkg/codec/rlp"
	"github.com/chain5j/chain5j-pkg/crypto/hashalg"
	"github.com/chain5j/chain5j-pkg/types"
)

var (
	typeRegistry = make(map[types.TxType]reflect.Type)
	lock         sync.RWMutex
)

// RegisterTransaction 注册交易交易模型
func RegisterTransaction(txInf interface{}) {
	typeOf := reflect.TypeOf(txInf)
	if typeOf.Kind() != reflect.Ptr {
		log().Crit("register transaction must pointer")
	}
	tx, ok := txInf.(Transaction)
	if !ok {
		log().Crit("register invalid transaction")
	}
	lock.Lock()
	defer lock.Unlock()
	typeRegistry[tx.TxType()] = reflect.TypeOf(txInf)
}

// NewTransaction 创建一个新的交易对象
func NewTransaction(t types.TxType) (Transaction, error) {
	rt, ok := typeRegistry[t]
	if !ok {
		log().Error("txType is not register", "txType", t)
		return nil, nil
	}

	v := reflect.New(rt.Elem()).Interface()
	return v.(Transaction), nil
}

// Transaction 交易接口
type Transaction interface {
	TxType() types.TxType      // 交易类型
	ChainId() string           // 链ID
	Hash() types.Hash          // 交易Hash
	Less(tx2 Transaction) bool // 是否小于tx2.用于交易排序
	Size() types.StorageSize   // 交易大小
	codec.Serializer           // 交易序列化
	codec.Deserializer         // 交易反序列化
	// Sign(prvKey *ecdsa.PrivateKey) ([]byte, error) // 签名
}

// StateTransaction 状态交易
type StateTransaction interface {
	Transaction

	From() string                   // 发送者
	To() string                     // 接收者
	GasLimit() uint64               // gasLimit
	Value() *big.Int                // value
	Input() []byte                  // 合约请求的数据
	GasPrice() uint64               // gasPrice
	Nonce() uint64                  // 唯一码
	Signer() (types.Address, error) // 签名者
	Cost() *big.Int                 // 预估balance消耗

}

type txTemp struct {
	TxType types.TxType
	Data   []byte
}

// TxEncode 交易编码
func TxEncode(tx Transaction) ([]byte, error) {
	bytes, err := tx.Serialize()
	if err != nil {
		log().Info("transaction serialize err", "txType", tx.TxType(), "hash", tx.Hash(), "err", err)
		return nil, err
	}
	jsonTx := &txTemp{
		TxType: tx.TxType(),
		Data:   bytes,
	}
	return codec.Coder().Encode(jsonTx)
}

// TxDecode 交易解码
func TxDecode(bytes []byte) (Transaction, error) {
	var jsonTx txTemp
	err := codec.Coder().Decode(bytes, &jsonTx)
	if err != nil {
		log().Info("transaction coder decode err", "err", err)
		return nil, err
	}
	tx, err := NewTransaction(jsonTx.TxType)
	if err != nil {
		log().Error("NewTransaction error", "txType", jsonTx.TxType, "err", err)
		return nil, err
	}

	err = tx.Deserialize(jsonTx.Data)
	if err != nil {
		log().Error("transaction deserialize err", "txType", jsonTx.TxType, "err", err)
		return nil, err
	}
	return tx, nil
}

// TxAsVmMessage 将transaction转换为vmMessage
func TxAsVmMessage(tx StateTransaction) (VmMessage, error) {
	var (
		to   *types.Address
		from types.Address
	)

	from = types.HexToAddress(tx.From())
	if from.Nil() {
		return nil, errors.New("tx from is nil")
	}

	if tx.To() != "" {
		addr := types.HexToAddress(tx.To())
		to = &addr
	}

	return NewEvmMessage(from, to, tx.Nonce(), tx.Value(), tx.GasLimit(), new(big.Int).SetUint64(tx.GasPrice()), tx.Input(), true), nil
}

// Transactions 交易集合
type Transactions []TransactionSortedList

// NewTransactions 创建一个新的交易集合
func NewTransactions(txs []TransactionSortedList) Transactions {
	if txs == nil || len(txs) == 0 {
		return nil
	}
	return txs
}

// DeepCopy 交易集合深度拷贝
func (txs Transactions) DeepCopy() Transactions {
	var newTxs Transactions
	newTxs = make([]TransactionSortedList, txs.Len())
	copy(newTxs, txs)
	return newTxs
}

// Hashes 交易集合的hash集
func (txs Transactions) Hashes() map[types.TxType][]types.Hash {
	if txs == nil {
		return nil
	}
	hashes := make(map[types.TxType][]types.Hash, len(txs))
	for _, datas := range txs {
		if datas.Len() > 0 {
			hashes[datas[0].TxType()] = datas.Hashes()
		}
	}
	return hashes
}

// Data 交易集合的交易数组
func (txs Transactions) Data() []TransactionSortedList {
	return txs
}

// GetTx 获取交易集合中指定位置的交易
func (txs Transactions) GetTx(txType types.TxType, i uint) Transaction {
	for _, txList := range txs {
		if txList[0].TxType() == txType {
			if int(i) >= txList.Len() {
				log().Warn("txs getIndex out of bounds")
				return nil
			}
			return txList[i]
		}
	}
	return nil
}
func (txs Transactions) AllLen() int {
	count := 0
	for _, list := range txs {
		count = count + list.Len()
	}
	return count
}

func (txs Transactions) Print() string {
	var buf bytes.Buffer
	for _, list := range txs {
		buf.WriteString(fmt.Sprintf("tx_type=%s,len=%d", list[0].TxType(), list.Len()))
		buf.WriteString(";")
	}
	s := buf.String()
	return s[:len(s)-1]
}
func (txs Transactions) Len() int {
	return len(txs)
}
func (txs Transactions) Less(i, j int) bool {
	return txs[i][0].TxType() < txs[j][0].TxType()
}
func (txs Transactions) Swap(i, j int) {
	txs[i], txs[j] = txs[j], txs[i]
}

type txsTemp struct {
	TxType types.TxType
	Data   [][]byte
}

// EncodeRLP rlp编码
func (txs Transactions) EncodeRLP(w io.Writer) error {
	var tpTxs []*txsTemp
	sort.Sort(txs)
	for _, txList := range txs {
		if len(txList) == 0 {
			continue
		}
		data := make([][]byte, txList.Len())
		for i, tx := range txList {
			bytes, err := rlp.EncodeToBytes(tx)
			if err != nil {
				log().Info("txs EncodeTxsRLP err", "err", err)
				break
			}
			data[i] = bytes
		}
		tpTx := &txsTemp{
			TxType: txList[0].TxType(),
			Data:   data,
		}
		tpTxs = append(tpTxs, tpTx)
	}
	return rlp.Encode(w, tpTxs)
}

// DecodeRLP rlp解码
func (txs *Transactions) DecodeRLP(s *rlp.Stream) error {
	var tpTxs []txsTemp
	err := s.Decode(&tpTxs)
	if err != nil {
		log().Error("txs decodeTxsRLP err", "err", err)
		return err
	}
	txsData := make([]TransactionSortedList, 0, len(tpTxs))
	for _, tpTxList := range tpTxs {
		txSortList := make([]Transaction, 0, len(tpTxList.Data))
		for _, tpTx := range tpTxList.Data {
			tx, err := NewTransaction(tpTxList.TxType)
			if err != nil {
				log().Error("rlp decode new transaction error", "err", err)
				continue
			}
			err = rlp.DecodeBytes(tpTx, tx)
			if err != nil {
				log().Error("rlp decode txBytes error", "err", err)
				continue
			}
			txSortList = append(txSortList, tx)
		}
		txsData = append(txsData, txSortList)
	}
	*txs = txsData
	return nil
}

// TxsRoot 所有交易的hash
func (txs Transactions) TxsRoot() []types.Hash {
	// txSet := bmt.WriteSet{}
	// for _, tx := range txs {
	//	data, err := tx.Serialize()
	//	if err != nil {
	//		return types.Hash{}
	//	}
	//	txSet[tx.Hash().Hex()] = data
	// }
	// root, _ := bmt.Hash(txSet)
	// return root
	sort.Sort(txs)
	txsRoot := make([]types.Hash, txs.Len())
	for i, list := range txs {
		txsRoot[i] = hashalg.RootHash(list)
	}
	return txsRoot
}

// Add 添加单笔交易
func (txs Transactions) Add(tx Transaction) {
	for _, list := range txs {
		if list[0].TxType() == tx.TxType() {
			list = append(list, tx)
			return
		}
	}
	return
}

// TransactionSortedList nonce递增排序，price递减排序
type TransactionSortedList []Transaction

func NewTransactionSortedList(txs []Transaction) TransactionSortedList {
	return txs
}
func (txList TransactionSortedList) Item(i int) []byte {
	if i >= txList.Len() {
		log().Warn("txs getIndex out of bounds")
		return nil
	}
	bytes, _ := txList[i].Serialize()
	return bytes
}
func (txList TransactionSortedList) Key(i int) []byte {
	if i >= txList.Len() {
		log().Warn("txs getIndex out of bounds")
		return nil
	}
	return txList[i].Hash().Bytes()
}
func (txList TransactionSortedList) Len() int {
	return len(txList)
}
func (txList TransactionSortedList) Less(i, j int) bool {
	return txList[i].Less(txList[j])
}
func (txList TransactionSortedList) Swap(i, j int) {
	txList[i], txList[j] = txList[j], txList[i]
}

func (txList TransactionSortedList) Hashes() []types.Hash {
	arr := make([]types.Hash, txList.Len())
	for i, tx := range txList {
		arr[i] = tx.Hash()
	}
	return arr
}
