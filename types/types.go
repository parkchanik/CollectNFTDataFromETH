package types

import (
	//"math/big"

	"github.com/ethereum/go-ethereum/common"
)

type Attribute struct {
	Trait_type string
	Value      string
}

type TokenMetaData struct {
	Name     interface{}
	Image    string
	FileName string
	//Attributes []Attribute
}

type TransferInfo struct {
	Contractaddress string //common.Address
	Topic_0         string // signature
	Topic_1         string // from
	Topic_2         string // to
	Topic_3         string // token id
}

type TokenInfo struct {
	TransactionHash common.Hash
	Contractaddress common.Address
	ContractName    string
	Symbol          string
	TokenID         string
	FileName        string
}

type LogData struct {
	TransactionHash       common.Hash
	BlockTime             string
	EtherValue            int64
	MatchContractsAddress common.Address
	TokenInfos            []TokenInfo
}

type TokenMetaDataBase64 struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Image       string `json:"image"`
	Attributes  []struct {
		TraitType string      `json:"trait_type"`
		Value     interface{} `json:"value"`
	} `json:"attributes"`
}

type TokenInfoNew struct {
	TransactionHash common.Hash
	BlockTime       string
	Contractaddress common.Address
	ContractName    string
	Symbol          string
	TokenID         string
	ETHValue        string
	TransferSigCnt  int
}
