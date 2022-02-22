package main

import (
	"context"
	"encoding/json"
	"errors"
	"flag"
	"io/ioutil"
	"math/big"
	"runtime"

	"bytes"
	"strings"

	"fmt"
	"io"
	"log"
	"os"

	"encoding/base64"

	"strconv"
	//"math"
	//"encoding/hex"
	"time"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"

	//"github.com/ethereum/go-ethereum/core/types"

	//"github.com/ethereum/go-ethereum/crypto"
	//"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/ethclient"

	erc20 "CollectNFTDataETH/contracts/output/ERC20"

	erc721 "CollectNFTDataETH/contracts/output/ERC721"

	wyvern "CollectNFTDataETH/contracts/output/WYVERN"

	//erc1155 "ETHCollectTrans/contracts/output/ERC1155"

	. "CollectNFTDataETH/types"

	logger "CollectNFTDataETH/logger"

	"CollectNFTDataETH/config"

	"net/http"
)

var client *ethclient.Client = nil

var IMAGE_PATH string = "../CollectNFT/Images/"

func main() {

	runtime.GOMAXPROCS(runtime.NumCPU())

	fmt.Println("GOMAXPROCS : ", runtime.GOMAXPROCS(0))

	fromNum := flag.Int64("fromblock", 0, "FromBlockNumber")
	toNum := flag.Int64("toblock", 0, "ToBlockNumber")

	flag.Parse()

	configData, err := config.LoadConfigration("config.json")
	if err != nil {
		log.Fatal("LoadConfigration :", err)
	}

	url := configData.URL
	logger.LoggerInit()

	/*
		logOrderMatchedSigHash := common.HexToHash("0xc4109843e0b7d514e4c093114b863f8e7d8d9a458c372cd51bfe526b588006c9") // ordermatch

		logTransferSigHash := common.HexToHash("0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef") // transfer ERC 721 일때

		logTransferSingleSigHash := common.HexToHash("0xc3d58168c5ae7397731d063d5bbf3d657854427343f4c083240f7aacaa2d0f62") //transferSingle  ERC1155 일때
	*/
	ethdial, err := ethclient.Dial(url)
	if err != nil {
		log.Fatal("ethclient.Dial ", err)
	}

	client = ethdial

	//https://etherscan.io/tx/0xd0aba0ee33bbc5fb725630a8e9f2bab44a8c9fbffdfcb3213690b283cac0a47a
	//https://etherscan.io/tx/0x7d7218d30264344363e4cdc208090bdba3eea622fe854df94b75fc9c12c65f4d

	//contractAddress := common.HexToAddress("0x7be8076f4ea4a4ad08075c2508e481d6c946d12b")
	// 컨트랙트 0x7be8076f4ea4a4ad08075c2508e481d6c946d12b opensea 의 topic[0] eventSignature
	//0xc4109843e0b7d514e4c093114b863f8e7d8d9a458c372cd51bfe526b588006c9
	//OrdersMatched (bytes32 buyHash, bytes32 sellHash, index_topic_1 address maker, index_topic_2 address taker, uint256 price, index_topic_3 bytes32 metadata)View Source
	//topicEventHashOrderMatch := common.HexToHash("0xc4109843e0b7d514e4c093114b863f8e7d8d9a458c372cd51bfe526b588006c9")
	//topicEventHashTransfer := common.HexToHash("0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef")

	// block number 13916166 (Jan-01-2022 12:00:03 AM +UTC)

	// block number 13717847 (Dec-01-2021 12:00:45 AM +UTC)
	// block number 13717846 (Nov-30-2021 11:59:50 PM +UTC)

	// block number 13717846 (Nov-30-2021 11:59:50 PM +UTC)
	// block number 13527859 (Nov-01-2021 12:00:07 AM +UTC)
	// block number 13527858 (Oct-31-2021 11:59:20 PM +UTC)
	// block number 13330090 (Oct-01-2021 12:00:00 AM +UTC)
	// block number 13330089 (Sep-30-2021 11:59:56 PM +UTC)

	// block number 12936340 (Aug-01-2021 12:00:17 AM +UTC)

	var fromBlockNumber int64 = 12936340 //13347221
	var toBlockNumber int64 = 13330190   //13347221

	if *fromNum != 0 {
		fromBlockNumber = *fromNum
		toBlockNumber = *toNum
	}

	logger.InfoLog("-----Start fromBlockNumber :  %d , toBlockNumber : %d", fromBlockNumber, toBlockNumber)

	for fromBlockNumber <= toBlockNumber {

		divideToBlockNumber := fromBlockNumber + 100 //1000
		CollectTrxProcess(fromBlockNumber, divideToBlockNumber)
		fromBlockNumber = divideToBlockNumber + 1

	}

}

func CollectTrxProcess(fromBlockNumber, toBlockNumber int64) {

	var minETHValue int = 2000000000 // 10000000000000000000 가 10 ether 인데 10개 뺀다

	address := "0x7be8076f4ea4a4ad08075c2508e481d6c946d12b" //opensea Project Wyvern Exchange contract address

	WyvernContractAddress := common.HexToAddress(address)

	//logApprovalSigHash := common.HexToHash("0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925") // approval topic[0]

	logOrderMatchedSigHash := common.HexToHash("0xc4109843e0b7d514e4c093114b863f8e7d8d9a458c372cd51bfe526b588006c9") // ordermatch  topic[0]

	logTransferSigHash := common.HexToHash("0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef") // transfer ERC 721 일때

	logTransferSingleSigHash := common.HexToHash("0xc3d58168c5ae7397731d063d5bbf3d657854427343f4c083240f7aacaa2d0f62") //transferSingle  ERC1155 일때

	logger.InfoLog("-----Start CollectTrxProcess filterQuery fromBlockNumber[%d] , toBlockNumber[%d]", fromBlockNumber, toBlockNumber)

	query := ethereum.FilterQuery{
		FromBlock: big.NewInt(fromBlockNumber),
		ToBlock:   big.NewInt(toBlockNumber),
		Addresses: []common.Address{
			WyvernContractAddress,
		},
		Topics: [][]common.Hash{
			{logOrderMatchedSigHash},
		},
	}

	logs, err := client.FilterLogs(context.Background(), query)
	if err != nil {
		log.Fatal(err)
	}

	for _, m := range logs {

		if m.Topics[0].Hex() == "0xc4109843e0b7d514e4c093114b863f8e7d8d9a458c372cd51bfe526b588006c9" { //wyvern  contract의 orderMatchHash

			wyverninstance, err := wyvern.NewWyvern(m.Address, client)
			if err != nil {
				logger.InfoLog("-------------------Error NewWyvern TxHash[%s] , err[%s]\n", m.TxHash.Hex(), err.Error())
				continue
			}

			wyverOrdersMatch, err := wyverninstance.ParseOrdersMatched(m)
			if err != nil {
				logger.InfoLog("-------------------Error NewWyvern ParseOrdersMatched TxHash[%s] , err[%s]\n", m.TxHash.Hex(), err.Error())
				continue

			}

			ETHString := wyverOrdersMatch.Price.String()

			if len(ETHString) < 18 {
				continue
			}

			ETHint := ChangeETHValue(ETHString)

			if ETHint >= minETHValue {
				//특정 eth 이상만 체크

				logger.InfoLog("--------------------------------------------------------------------------------------------------------\n")
				logger.InfoLog("-------Order Match Value than %d -------------------ETHint >= minETHValue-------------------\n", minETHValue)
				blocknum := m.BlockNumber
				blocknumNew := big.NewInt(int64(blocknum))

				txhash := m.TxHash

				ETHLast := fmt.Sprintf("%f", float64(ETHint)/100000000)

				block, err := client.BlockByNumber(context.Background(), blocknumNew)
				if err != nil {
					logger.InfoLog("!!!!!Error BlockByHash Hash Get Error BlockByNumber[%d] , err[%s]\n", blocknumNew.Int64(), err.Error())

				}

				coinType := ""

				transaction := block.Transaction(txhash)

				trxValue := transaction.Value().String()

				logger.InfoLog("block Trx Value  TxHash[%s] trxValue[%s]\n", txhash, trxValue)

				if trxValue != "0" { // 트랜잭션 value 에 값이 있으면 ETH로 거래 한 내용이라고 처리
					coinType = "ETH"
				}

				blocktime := int64(block.Time())
				blocktimestring := time.Unix(blocktime, 0).Format("2006-01-02 15:04:05")

				// 해당 트랜잭션의 영수증
				rept, err := client.TransactionReceipt(context.Background(), txhash)
				if err != nil {
					logger.InfoLog("!!!!TransactionReceiptt Error vLog.TxHash[%s] , err[%s]\n", txhash, err.Error())
					continue
				}

				transferSigCount := 0
				orderMatchSig := 0
				transferSingleSigCount := 0
				//0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef transfer
				// transfer 함수가 없으면 패스 하는걸로 한다
				// 첫번째 토픽 OrdersMatched 로그가 아니

				//////////////////////
				// ETH WETH 외에 다른 토큰으로도 거래가 가능하다
				// https://etherscan.io/tx/0x734279e0043eb3583467c3ad46e8064d293867e9e9b25734a07dc3c22574b67b SAND토큰으로 거래한 예제
				// https://etherscan.io/tx/0x113d7a881afa660fd063681022ee13560b7414731a0e1304fba71bc047d113e3 ETH 로만 거래한 예제
				// https://etherscan.io/tx/0x9193b913d9041603b3311324ffd6ebd00f0efdf736f570bf1c73480b03e1a6fc WETH 로 거래 한 예제

				// https://etherscan.io/tx/0xacd9ec7f7ea4630dded33752ce6bf18c6be3ff1cd560c261b5665ec9f0332f89 ETH 거래 인데 transfer 가 이상한 애들은 따 로처리 해야 할듯

				orderMatchContractAddress := ""
				for _, m := range rept.Logs {

					if m.Topics[0] == logTransferSigHash { //WETH 주소 transfer 라면 transfer count  에 넣지 않는다
						transferSigCount = transferSigCount + 1
					}

					if m.Topics[0] == logTransferSingleSigHash {
						transferSingleSigCount = transferSingleSigCount + 1
					}

					if m.Topics[0] == logOrderMatchedSigHash {
						orderMatchSig = orderMatchSig + 1

						orderMatchContractAddress = m.Address.Hex()
					}

				}

				// 아래와 같은 경우가있었으니 나중에 참고
				// 			!!!!!!!!!!transferSigCount[2] , transferSingleSigCount[0] ,  orderMatchSig[0] txs.Hash[0xfbf1f28b04325a4ade7edbe3efcf7a85f8f4a58da4c201b0527f65bc07f76323]

				// !!!!!!!!!!transferSigCount[2] , transferSingleSigCount[0] ,  orderMatchSig[0] txs.Hash[0x243a5b11e99e1edfc25065dd3f0aa0230a62fcb40ed1bcfbc57fea429f29967a]

				// !!!!!!!!!!transferSigCount[2] , transferSingleSigCount[0] ,  orderMatchSig[0] txs.Hash[0x3441db76d0221145ea77416fa91d5f2bf67d526e4eecc1c9451545d68da9f989]

				if orderMatchSig == 0 {
					//logger.InfoLog("!!!!!!!!!!|| orderMatchSig == 0 txs.Hash[%s]\n", txhash)
					continue
				}

				//if transferSigCount == 0 {
				if transferSingleSigCount > 0 {
					logger.InfoLog("!!!!!!!!!!!!transferSigCount ==0 Not ERC-721 pass !! transferSigCount[%d] , transferSingleSigCount[%d] ,  orderMatchSig[%d] txs.Hash[%s]\n", transferSigCount, transferSingleSigCount, orderMatchSig, txhash)
					continue
				}

				logger.InfoLog("--transferSigCount[%d] transferSingleSigCount[%d] orderMatchSig[%d] orderMatchContractAddress[%s] txs.Hash[%s]\n", transferSigCount, transferSingleSigCount, orderMatchSig, orderMatchContractAddress, txhash)

				//  m.Address.String() != "0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2" WETH address

				var tokenAddress common.Address
				tokenName := ""
				tokenSymbol := ""
				// 첫번째 event log 가 뭐냐에 따라서 달라진다
				// topic[0] =  0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925 Approval 이면 일반 ETH 거래
				//꼭 첫 log 가 approval 이 아니더라도 ETH 거래인경우가 있더라 Transaction 에서 체크 한다 주석처리
				// if rept.Logs[0].Topics[0] == logApprovalSigHash {
				// 	// type ETH
				// 	coinType = "ETH"
				// }

				if coinType == "ETH" {

					tokenName = coinType
					tokenSymbol = coinType

				} else { // 타입이 ETH가 아니면 첫 transfer 에 token 정보를 읽는다
					if rept.Logs[0].Topics[0] == logTransferSigHash { //

						// topic[0] 이 logTransferSigHash 인데 그 주소가  WETH 라면
						// if rept.Logs[0].Address == common.HexToAddress("0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2") {
						// 	coinType = "WETH"
						// }

						erc20, err := erc20.NewErc20(rept.Logs[0].Address, client)
						if err != nil {
							logger.InfoLog("!!! NewErc20 Error vLog.TxHash[%s] ,rept.Logs[0].Address[%s] err[%s]\n", txhash, rept.Logs[0].Address.Hex(), err.Error())
						}

						tokenAddress = rept.Logs[0].Address

						name, err := erc20.Name(&bind.CallOpts{})
						if err != nil {
							logger.InfoLog("!!! NewErc20 Name Error vLog.TxHash[%s] ,rept.Logs[0].Address[%s] err[%s]\n", txhash, rept.Logs[0].Address.Hash(), err.Error())
						}

						tokenName = name

						symbol, err := erc20.Symbol(&bind.CallOpts{})
						if err != nil {
							logger.InfoLog("!!! NewErc20 Symbol Error vLog.TxHash[%s] ,rept.Logs[0].Address[%s] err[%s]\n", txhash, rept.Logs[0].Address.Hash(), err.Error())
						}

						coinType = "TOKEN"
						tokenSymbol = symbol
					}
				}

				if coinType == "" { // ETH 도 아니고 WETH 도 아니면 어쩔꺼냐
					logger.InfoLog("------CoinType is blank !!!! why ?? vLog.TxHash[%s] ,rept.Logs[0].Address[%s]\n", txhash, rept.Logs[0].Address.Hash())
				}

				logger.InfoLog("-- OrdersMatched Price TxHash[%s] BlockTime[%s] CoinType[%s] TokenName[%s] TokenSymbol[%s] PriceInt[%d] PriceString[%s] ETHLast[%s] Len[%d]\n", m.TxHash.Hex(), blocktimestring, coinType, tokenName, tokenSymbol, wyverOrdersMatch.Price.Int64(), wyverOrdersMatch.Price.String(), ETHLast, len(wyverOrdersMatch.Price.String()))

				transferAlready := false

				contractName := ""
				var contractAddress common.Address
				contractSymbol := ""
				var tokenID int64 = 0
				var transferCountERC721 int = 0
				for _, z := range rept.Logs {

					//Transfer //토큰 erc20 주소 transfer 라면 transfer 처리 하지 않는다
					if z.Address == tokenAddress {
						logger.InfoLog("-- z.Address != tokenAddress z.Address[%s] tokenAddress[%s]\n", z.Address.Hex(), tokenAddress.Hex())
						continue
					}

					if z.Topics[0] == logTransferSigHash {

						if transferAlready == false { // erc721 의 첫 transfer 만 저장 한다 나머지는 카운트만 센다
							instance, err := erc721.NewErc721(z.Address, client)
							if err != nil {
								logger.InfoLog("!!! Error GetDataERC721 NewErc721  error[%s] ", err.Error())
								continue
							}

							name, err := instance.Name(&bind.CallOpts{})
							if err != nil {
								logger.InfoLog("!!! Error GetDataERC721 instance.Name error[%s] ", err.Error())

							}

							symbol, err := instance.Symbol(&bind.CallOpts{})
							if err != nil {
								logger.InfoLog("!!! Error GetDataERC721 instance.Symbol error[%s] ", err.Error())

							}

							//0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef transfer
							erc721transfer, err := instance.ParseTransfer(*z)
							if err != nil {
								logger.InfoLog("!!! Error GetDataERC721 instance.ParseTransfer Log Addresss[%s] error[%s] ", z.Address.Hex(), err.Error())
								continue
							}

							contractName = name
							contractSymbol = symbol
							tokenID = erc721transfer.TokenId.Int64()

							contractAddress = z.Address

							transferAlready = true
						}

						transferCountERC721 = transferCountERC721 + 1

					}
				}

				tokeninfo := &TokenInfoNew{}

				tokeninfo.BlockTime = blocktimestring
				tokeninfo.TransactionHash = txhash
				tokeninfo.ContractName = contractName
				tokeninfo.Contractaddress = contractAddress
				tokeninfo.ContractSymbol = contractSymbol

				tokeninfo.CoinType = coinType
				tokeninfo.Tokenaddress = tokenAddress
				tokeninfo.TokenName = tokenName
				tokeninfo.TokenSymbol = tokenSymbol

				tokenIDStr := fmt.Sprintf("%d", tokenID)
				tokeninfo.Value = ETHLast

				tokeninfo.TokenID = tokenIDStr
				tokeninfo.TransferCountERC721 = transferCountERC721
				//transferSigCount
				PrintTokenDataNew(tokeninfo)

			}

			// wyverOrdersMatch.Price.Int64() 로 하면 특정 수를 넘어 가면 overflow 가 나서 이상한 값으로 오는 듯 하다
			//OrdersMatched Price TxHash[0xeb3a9351c34094fc568d2b25946b724d32b7cf8679509d33c7385a4c2edcd04c] , PriceInt[9106511852580896768]  ,PriceStringp[46000000000000000000]
			// 46000000000000000000 -> 46 ether
			//logger.InfoLog("------OrdersMatched Price TxHash[%s] , PriceInt[%d]  ,PriceStringp[%s] Len[%d]\n", m.TxHash.Hex(), wyverOrdersMatch.Price.//Int64(), wyverOrdersMatch.Price.String(), len(wyverOrdersMatch.Price.String()))

		}

	}

}

func ChangeETHValue(ValueString string) int {

	wKlayString := ValueString

	//logger.InfoLog("----wKlayString[%s]\n", wKlayString)

	wKlayrune := []rune(wKlayString)

	//logger.InfoLog("----wKlayrune[%s]\n", wKlayrune)
	rune10length := len(wKlayrune) - 10 // 전체 길이에서 10을 뺀다
	//logger.InfoLog("----rune10length[%d]\n", rune10length)

	wklayMinimal := string(wKlayrune[:rune10length])

	//logger.InfoLog("----wklayMinimal[%s]\n", wklayMinimal)

	wklayint, err := strconv.Atoi(wklayMinimal)
	if err != nil {
		return -1
	}

	return wklayint

}

func PrintTokenDataNew(logdata *TokenInfoNew) {

	transaction := logdata.TransactionHash.Hex()
	blockTime := logdata.BlockTime[:10]
	contractAddress := logdata.Contractaddress.Hex()
	contractName := logdata.ContractName
	contractSymbol := logdata.ContractSymbol
	tokenID := logdata.TokenID
	ETHValue := logdata.Value //float64

	transferCountERC721 := strconv.Itoa(logdata.TransferCountERC721)

	coinType := logdata.CoinType
	tokenAddress := logdata.Tokenaddress.Hex()
	tokenName := logdata.TokenName
	tokenSymbol := logdata.TokenSymbol

	var b bytes.Buffer

	b.WriteString(blockTime)
	b.WriteString(",")
	b.WriteString(transaction)
	b.WriteString(",")
	b.WriteString(contractAddress)
	b.WriteString(",")
	b.WriteString(contractName)
	b.WriteString(",")
	b.WriteString(contractSymbol)
	b.WriteString(",")
	b.WriteString(tokenID)
	b.WriteString(",")
	b.WriteString(coinType)
	b.WriteString(",")
	b.WriteString(tokenAddress)
	b.WriteString(",")
	b.WriteString(tokenName)
	b.WriteString(",")
	b.WriteString(tokenSymbol)
	b.WriteString(",")
	b.WriteString(ETHValue)
	b.WriteString(",")
	b.WriteString(transferCountERC721)

	logger.TokenLog(b.String())

}

func PrintTokenData(logdata *TokenInfo) {

	transaction := logdata.TransactionHash.Hex()
	contractAddress := logdata.Contractaddress.Hex()
	contractName := logdata.ContractName
	contractSymbol := logdata.Symbol
	tokenID := logdata.TokenID

	var b bytes.Buffer

	b.WriteString(transaction)
	b.WriteString(",")
	b.WriteString(contractAddress)
	b.WriteString(",")
	b.WriteString(contractName)
	b.WriteString(",")
	b.WriteString(contractSymbol)
	b.WriteString(",")
	b.WriteString(tokenID)

	logger.TokenLog(b.String())

}

func PrintTrxData(logdata *LogData) {

	transaction := logdata.TransactionHash.Hex()

	timestring := logdata.BlockTime[:10]

	etherstring := fmt.Sprintf("%f", float64(logdata.EtherValue)/1000000000000000000)

	var b bytes.Buffer

	b.WriteString(transaction)
	b.WriteString(",")
	b.WriteString(timestring)
	b.WriteString(",")
	b.WriteString(etherstring)

	logger.TrxLog(b.String())

}

func getTokenImageUri(tokenuri string) (string, error) {

	//ipfs:// 로 시작하면 변경해줘야 한다

	// https://ipfs.io/ipfs/QmSTtv3w1jqcv5AKRRYVR5NN7fkTuuL9sNrkxRNL9e3fUo/4744 이런식으로

	if strings.Contains(tokenuri, "ipfs://") == true {

		tokenuri = strings.ReplaceAll(tokenuri, "ipfs://", "https://ipfs.io/ipfs/")

	}

	res, err := http.Get(tokenuri)
	if err != nil {
		return "", err
		//fmt.Printf("http Get Error Transaction[%s] , Tokenuri[%s] Error[%s]\n ", vLog.TxHash, tokenuri, err.Error())
	}

	defer res.Body.Close()

	data, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return "", err
		//fmt.Printf("res.Body error  Transaction[%s] , Tokenuri[%s] Error[%s]\n ", vLog.TxHash, tokenuri, err.Error())

	}
	metadata := TokenMetaData{}

	err = json.Unmarshal(data, &metadata)
	if err != nil {
		return "", err
		//fmt.Printf("metadata unmarshal error Transaction[%s] , Tokenuri[%s] Error[%s]\n ", vLog.TxHash, tokenuri, err.Error())

	}

	return metadata.Image, nil

}

func getTokenMetaData(tokenuri string) (TokenMetaData, error) {

	metadata := TokenMetaData{}

	//ipfs:// 로 시작하면 변경해줘야 한다
	// https://ipfs.io/ipfs/QmSTtv3w1jqcv5AKRRYVR5NN7fkTuuL9sNrkxRNL9e3fUo/4744 이런식으로

	// tokenuri
	//ipfs://QmWS694ViHvkTms9UkKqocv1kWDm2MTQqYEJeYi6LsJbxK 이런 경우가있고
	//ipfs://ipfs/QmWS694ViHvkTms9UkKqocv1kWDm2MTQqYEJeYi6LsJbxK 이런 경우도 있다 이놈때문에이렇게 바꿔존다
	// https://ipfs.io/ipfs/QmWS694ViHvkTms9UkKqocv1kWDm2MTQqYEJeYi6LsJbxK 이렇게 바꾼다

	logger.InfoLog("-------tokenuri before : %s", tokenuri)

	r := strings.NewReplacer("ipfs://ipfs/", "https://ipfs.io/ipfs/", "ipfs://", "https://ipfs.io/ipfs/")

	tokenuri = r.Replace(tokenuri)

	logger.InfoLog("-------tokenuri after  %s", tokenuri)

	res, err := http.Get(tokenuri)
	if err != nil {
		logger.InfoLog("-------getTokenMetaData http.Get(tokenuri) tokenuri[%s] error[%s] ", tokenuri, err.Error())
		return metadata, err

	}

	defer res.Body.Close()

	data, err := ioutil.ReadAll(res.Body)
	if err != nil {
		logger.InfoLog("-------getTokenMetaData ioutil.ReadAll tokenuri[%s] error[%s] ", tokenuri, err.Error())
		return metadata, err

	}

	err = json.Unmarshal(data, &metadata)
	if err != nil {
		logger.InfoLog("-------getTokenMetaData  json.Unmarshal(data, &metadata)  data[%s] error[%s] ", string(data), err.Error())
		return metadata, err

	}

	return metadata, nil

}

func downloadFile(URL, fileName string) error {
	//Get the response bytes from the url

	logger.InfoLog("start download image uri : %s , fileName : %s \n", URL, fileName)

	response, err := http.Get(URL)
	if err != nil {
		return err
	}
	defer response.Body.Close()

	if response.StatusCode != 200 {
		logger.InfoLog("-------downloadFile status code is not 200  URL[%s] fileName[%s] , code[%d]", URL, fileName, response.StatusCode)
		return errors.New("Received non 200 response code")
	}
	//Create a empty file
	file, err := os.Create(fileName)
	if err != nil {
		return err
	}
	defer file.Close()

	//Write the bytes to the fiel
	_, err = io.Copy(file, response.Body)
	if err != nil {
		return err
	}

	return nil
}

func GetERC721Data(eventlog types.Log) (ContractAddr common.Address, Name string, Symbol string, TokenID string, err error) {

	err = nil

	ContractAddr = eventlog.Address
	Name = ""
	Symbol = ""
	TokenID = ""

	instance, err := erc721.NewErc721(eventlog.Address, client)
	if err != nil {
		logger.InfoLog("GetDataERC721 NewErc721 contractAddressHex[%s] , error[%s] ", ContractAddr.Hex(), err.Error())
		return
	}

	Name, err = instance.Name(&bind.CallOpts{})
	if err != nil {
		logger.InfoLog("GetDataERC721 instance.Name error[%s] ", err.Error())

	}

	Symbol, err = instance.Symbol(&bind.CallOpts{})
	if err != nil {
		logger.InfoLog("GetDataERC721 instance.Symbol error[%s] ", err.Error())

	}

	//0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef transfer
	erc721transfer, err := instance.ParseTransfer(eventlog)
	if err != nil {
		logger.InfoLog("GetDataERC721 instance.ParseTransfer  error[%s] ", err.Error())
		return
	}

	TokenID = fmt.Sprintf("%s", erc721transfer.TokenId)

	logger.InfoLog("GetDataERC721  From[%s] , To[%s]  , TokenID[%d]", erc721transfer.From.Hex(), erc721transfer.To.Hex(), erc721transfer.TokenId.Int64())

	return

}

func GetDataERC721(eventlog types.Log) (ContractAddr string, Name string, Symbol string, TokenID string, TokenURI string, err error) {

	err = nil
	ContractAddr = ""
	Name = ""
	Symbol = ""
	TokenID = ""
	TokenURI = ""
	ContractAddr = eventlog.Address.Hex()

	instance, err := erc721.NewErc721(eventlog.Address, client)
	if err != nil {
		logger.InfoLog("GetDataERC721 NewErc721 contractAddressHex[%s] , error[%s] ", ContractAddr, err.Error())
		return
	}

	Name, err = instance.Name(&bind.CallOpts{})
	if err != nil {
		logger.InfoLog("GetDataERC721 instance.Name error[%s] ", err.Error())
		return
	}

	Symbol, err = instance.Symbol(&bind.CallOpts{})
	if err != nil {
		logger.InfoLog("GetDataERC721 instance.Symbol error[%s] ", err.Error())
		return
	}

	//0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef transfer
	erc721transfer, err := instance.ParseTransfer(eventlog)
	if err != nil {
		logger.InfoLog("GetDataERC721 instance.ParseTransfer  error[%s] ", err.Error())
		return
	}

	TokenID = fmt.Sprintf("%s", erc721transfer.TokenId)

	logger.InfoLog("GetDataERC721  From[%s] , To[%s]  , TokenID[%d]", erc721transfer.From.Hex(), erc721transfer.To.Hex(), erc721transfer.TokenId.Int64())

	TokenURI, err = instance.TokenURI(&bind.CallOpts{}, erc721transfer.TokenId)
	if err != nil {
		logger.InfoLog("GetDataERC721 Token URI : tokenid[%d] , error[%s] ", erc721transfer.TokenId.Int64(), err.Error())
		return
	}

	return

}

func GetImageFromDataApplicationJson(tokenuri, pathandfilename string) string {

	logger.InfoLog("------- tokenuri uri [%s]\n", "data:application/json........")

	//logger.InfoLog("token uri data:json : imageuri uri %s\n", tokenuri)

	tokenuriarr := strings.Split(tokenuri, ",")

	tokenMetaData := TokenMetaDataBase64{}

	if strings.Trim(tokenuriarr[0], " ") == "data:application/json;utf8" {

		//logger.InfoLog("token uri data:json : strings.Replace(tokenuri, data:application/json;utf8 uri %s\n", strings.Replace(tokenuri, "data:application/json;utf8,", "", 1))

		data := strings.Replace(tokenuri, "data:application/json;utf8,", "", 1)

		//logger.InfoLog("------- tokenuri uri [%s]\n", tokenuriarr[0])
		err := json.Unmarshal([]byte(data), &tokenMetaData)
		if err != nil {
			logger.InfoLog(" tokenMetaData utf8 Unmarshal Error : ", err)
			logger.InfoLog("token string [%s]\n", tokenuriarr[1])
			return ""
		}

	} else if strings.Trim(tokenuriarr[0], " ") == "data:application/json;base64" {

		logger.InfoLog("------- tokenuri uri [%s]\n", tokenuriarr[0])

		data, err := base64.StdEncoding.DecodeString(tokenuriarr[1])
		if err != nil {
			logger.InfoLog(" tokenMetaData base64.StdEncoding.DecodeString Error : ", err)
			return ""
		}

		//fmt.Printf("test data : %s\n", string(data))

		err = json.Unmarshal(data, &tokenMetaData)
		if err != nil {
			logger.InfoLog(" tokenMetaData base64 Unmarshal Error : ", err)
			logger.InfoLog("token DecodeString [%s]\n", string(data))
			return ""
		}

	} else {

		logger.InfoLog("------- tokenuri uri not  data:application/json;utf8 and  data:application/json;base64 [%s]\n", tokenuriarr[0])
		return ""
	}

	//logger.InfoLog("token uri data:json : imageuri tokenuriarr[1]  ---- uri [%s]\n", tokenuriarr[1])

	imagearr := strings.Split(tokenMetaData.Image, ",")

	file, err := os.Create(pathandfilename)
	if err != nil {
		logger.InfoLog("getImageFromDataApplicationJson os.Create Error : ", err)
		return ""
	}

	defer file.Close()

	//logger.InfoLog("tokenMetaData.Image[%s]\n", tokenMetaData.Image)

	if strings.Trim(imagearr[0], " ") == "data:image/svg+xml;utf8" {

		//logger.InfoLog("data:image/svg+xml;utf8 imagearr[1][%s]\n", imagearr[2])

		imageUTF8 := strings.Replace(tokenMetaData.Image, "data:image/svg+xml;utf8,", "", 1)

		cnt, err := file.WriteString(imageUTF8)
		if err != nil {
			logger.InfoLog("getImageFromDataApplicationJson data:image/svg+xml;utf8 file.WriteString Error : ", err)
			return ""
		}

		logger.InfoLog("file.WriteString data:image/svg+xml;utf8 cnt %d ", cnt)

		return "OK"

	} else if strings.Trim(imagearr[0], " ") == "data:image/svg+xml;base64" { // svg , base64 로 인코딩 되어있는 경우 svg 를 파일로
		imgdata, err := base64.StdEncoding.DecodeString(imagearr[1])
		if err != nil {
			logger.InfoLog("base64.StdEncoding.DecodeString(imagearr Error : ", err)
			return ""
		}

		//logger.InfoLog("base64.StdEncoding.DecodeString  %s\n", imgdata)

		cnt, err := file.WriteString(string(imgdata))
		if err != nil {
			logger.InfoLog("getImageFromDataApplicationJson data:image/svg+xml;base64 file.WriteString Error : ", err)
			return ""
		}

		logger.InfoLog("file.WriteString data:image/svg+xml;base64 cnt %d ", cnt)

		return "OK"
	}

	return ""
}

func GetTokenURIData(tokenuri, tokenid, contractName string) string {

	replacer := strings.NewReplacer(" ", "_", ":", "", "?", "", "*", "", "<", "", ">", "", "|", "", "\"", "", "/", "")
	contractNameFilter := replacer.Replace(contractName)

	rtn := ""
	if strings.Contains(tokenuri, "data:application/json") == true {

		filename := fmt.Sprintf("%s_%s.svg", contractNameFilter, tokenid)
		pathandfilename := fmt.Sprintf("%s%s", IMAGE_PATH, filename)
		result := GetImageFromDataApplicationJson(tokenuri, pathandfilename)

		rtn = filename

		if result == "OK" {

		} else {
			logger.InfoLog("GetImageFromDataApplicationJson Result Not OK Tokenuri[%s] , FileName[%s] \n ", tokenuri, filename)
		}

	} else {

		logger.InfoLog("------- tokenuri uri [%s]\n", tokenuri)

		tokenMetaData, err := getTokenMetaData(tokenuri)
		if err != nil {
			logger.InfoLog("--------------------------getTokenImageUri , Tokenuri[%s] Error[%s]\n ", tokenuri, err.Error())
		} else {

			imageuri := tokenMetaData.Image

			filename := fmt.Sprintf("%s_%s.png", contractNameFilter, tokenid)
			pathandfilename := fmt.Sprintf("%s%s", IMAGE_PATH, filename)

			rtn = filename

			if strings.Contains(imageuri, "ipfs://") == true {
				imageuri = strings.ReplaceAll(imageuri, "ipfs://", "https://ipfs.io/ipfs/")
			}

			if strings.Contains(imageuri, "ipfs") == true { /// 20220116 ipfs 에서 image 다운로드가 너무 오래걸린다  받아 지지도 않음 download pas

				logger.InfoLog("------ipfs image url!! Tokenuri[%s] FileName[%s] ,  ImageURL[%s]\n ", tokenuri, filename, imageuri)
			} else {

				err = downloadFile(imageuri, pathandfilename)
				if err != nil {
					logger.InfoLog("--------------------------downloadfile error Transaction[%s] , Image[%s] , FileName[%s] , Error[%s]\n ", imageuri, filename, err.Error())

				}
			}
		}

	}

	return rtn

}
