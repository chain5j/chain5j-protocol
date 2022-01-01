// Package models
//
// @author: xwc1125
package models

import (
	"encoding/json"
	"math/big"

	"github.com/chain5j/chain5j-pkg/types"
)

// AccountRef 合约对象
type AccountRef types.Address

// Address 合约地址
func (ar AccountRef) Address() types.Address { return (types.Address)(ar) }

// VmMessage 合约调用对象
type VmMessage interface {
	From() string       // 发送者
	To() string         // 接收者
	Nonce() uint64      // 唯一码
	GasLimit() uint64   // gasLimit
	GasPrice() *big.Int // gasPrice
	Value() *big.Int    // value
	Input() []byte      // 合约请求的数据
	CheckNonce() bool   // 是否检测nonce
}

type Message struct {
	msgData
}

type msgData struct {
	From       types.Address  `json:"from"`
	To         *types.Address `json:"to"`
	Nonce      uint64         `json:"nonce"`
	Amount     *big.Int       `json:"amount"`
	GasLimit   uint64         `json:"gas_limit"`
	GasPrice   *big.Int       `json:"gas_price"`
	Data       []byte         `json:"data"`
	CheckNonce bool           `json:"check_nonce"`
}

func NewEvmMessage(from types.Address, to *types.Address, nonce uint64, amount *big.Int, gasLimit uint64, gasPrice *big.Int, data []byte, checkNonce bool) Message {
	return Message{msgData{
		From:       from,
		To:         to,
		Nonce:      nonce,
		Amount:     amount,
		GasLimit:   gasLimit,
		GasPrice:   gasPrice,
		Data:       data,
		CheckNonce: checkNonce,
	}}
}

func (m Message) From() string { return m.msgData.From.Hex() }
func (m Message) To() string {
	if m.msgData.To == nil {
		return ""
	}
	return m.msgData.To.Hex()
}
func (m Message) GasPrice() *big.Int { return m.msgData.GasPrice }
func (m Message) SetPrice(gasPrice *big.Int) {
	m.msgData.GasPrice = gasPrice
}
func (m Message) Value() *big.Int { return m.msgData.Amount }
func (m Message) SetValue(value *big.Int) {
	m.msgData.Amount = value
}
func (m Message) GasLimit() uint64 { return m.msgData.GasLimit }
func (m Message) Nonce() uint64    { return m.msgData.Nonce }
func (m Message) Input() []byte    { return m.msgData.Data }
func (m Message) CheckNonce() bool { return m.msgData.CheckNonce }

func (m *Message) MarshalJSON() ([]byte, error) {
	return json.Marshal(m.msgData)
}

func (m *Message) UnmarshalJSON(input []byte) error {
	return json.Unmarshal(input, &m.msgData)
}
