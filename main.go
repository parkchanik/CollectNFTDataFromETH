package main

import (
	"context"
	"encoding/json"
	"errors"
	"flag"
	"io/ioutil"
	"math/big"
	"runtime"
	"time"

	"bytes"
	"strings"

	"fmt"
	"io"
	"log"
	"os"

	"encoding/base64"

	//"math"
	//"encoding/hex"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"

	//"github.com/ethereum/go-ethereum/core/types"

	//"github.com/ethereum/go-ethereum/crypto"
	//"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/ethclient"

	erc721 "CollectNFTDataETH/contracts/output/ERC721"

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

	logOrderMatchedSigHash := common.HexToHash("0xc4109843e0b7d514e4c093114b863f8e7d8d9a458c372cd51bfe526b588006c9") // ordermatch

	logTransferSigHash := common.HexToHash("0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef") // transfer ERC 721 일때

	logTransferSingleSigHash := common.HexToHash("0xc3d58168c5ae7397731d063d5bbf3d657854427343f4c083240f7aacaa2d0f62") //transferSingle  ERC1155 일때

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

	var fromBlockNumber int64 = 13330303 //13347221
	var toBlockNumber int64 = 13330303   //13347221

	if *fromNum != 0 {
		fromBlockNumber = *fromNum
		toBlockNumber = *toNum
	}

	logger.InfoLog("-----Start fromBlockNumber :  %d , toBlockNumber : %d", fromBlockNumber, toBlockNumber)

	var minETHValue int64 = 8000000000000000000
	//var minETHValue int64 = 0

	i := fromBlockNumber

	for i <= toBlockNumber {

		logger.InfoLog("----- Block Num :  %d , Time : %s", i, time.Now())

		blockNum := big.NewInt(i)

		block, err := client.BlockByNumber(context.Background(), blockNum)
		if err != nil {
			logger.InfoLog("!!!!!!!!!!!!!!!!!!!!!!!!!!BlockByHash Hash Get Error BlockByNumber[%d] , err[%s]\n", blockNum.Int64(), err.Error())
			log.Fatal(err)
		}

		blocktime := int64(block.Time())
		blocktimestring := time.Unix(blocktime, 0).Format("2006-01-02 15:04:05")

		for _, txs := range block.Transactions() {

			etherint64 := txs.Value().Int64()

			if etherint64 < minETHValue {
				continue
			}

			txhash := txs.Hash()
			// 해당 트랜잭션의 영수증
			rept, err := client.TransactionReceipt(context.Background(), txhash)
			if err != nil {
				logger.InfoLog("!!!!!!!!!!!!!!!!!!!!!!!!!!TransactionReceiptt Error vLog.TxHash[%s] , err[%s]\n", txhash, err.Error())
				continue
			}

			if len(rept.Logs) == 0 { //event log 가없으면 일반 거래일것이다
				continue
			}

			// 아래는 테스트 트랜잭션만 처리 하기 위해 추가
			// 0x7c5125feedc5cf4dd447bde160a6e13a089c1a0ac5431267c5eabcc7321d1ca0 -- erc1155
			// 0xa8f5f098526f577d544f874bed744ec84b7eada669836a18cb82e4540e436b10 -- erc721
			// if txhash.Hex() != "0xf0179b678809acff8535ad89338bc7fa8a87d28cc10f07c7e595ef823b0e4690" {
			if txhash.Hex() != "0x385c223e7183b5a0ed07edd5f72564f0df1874da428ffa1f325b4496cca631b6" {
				continue
			}

			transferSigCount := 0
			orderMatchSig := 0
			transferSingleSigCount := 0
			//0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef transfer
			// transfer 함수가 없으면 패스 하는걸로 한다
			// 첫번째 토픽 OrdersMatched 로그가 아니

			orderMatchContractAddress := ""
			for _, m := range rept.Logs {
				if m.Topics[0] == logTransferSigHash {
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

			if transferSigCount == 0 {
				logger.InfoLog("!!!!!transferSigCount ==0 Not ERC-721 pass !! transferSigCount[%d] , transferSingleSigCount[%d] ,  orderMatchSig[%d] txs.Hash[%s]\n", transferSigCount, transferSingleSigCount, orderMatchSig, txhash)
				continue
			}

			logger.InfoLog("--------------------------------------------------------------------------------------------------------\n")
			logger.InfoLog("!!!!!!!!!!transferSigCount[%d] , transferSingleSigCount[%d] ,  orderMatchSig[%d] , orderMatchContractAddress[%s] , txs.Hash[%s]\n", transferSigCount, transferSingleSigCount, orderMatchSig, orderMatchContractAddress, txhash)

			//////////////////////////////

			logdata := &LogData{}
			logdata.TransactionHash = txs.Hash()
			logdata.BlockTime = blocktimestring

			logdata.EtherValue = etherint64

			PrintTrxData(logdata)

			for _, z := range rept.Logs {
				if z.Topics[0] == logTransferSigHash { //Transfer

					contractAddr, name, symbol, tokenid, err := GetERC721Data(*z)
					if err != nil {
						logger.InfoLog("-------Error--getDataERC721 txs.Hash[%s] , error[%s] ", txhash, err.Error())
						continue
					}

					tokeninfo := &TokenInfo{}

					tokeninfo.TransactionHash = txs.Hash()
					tokeninfo.Contractaddress = contractAddr

					tokeninfo.ContractName = name

					tokeninfo.Symbol = symbol

					tokeninfo.TokenID = tokenid

					PrintTokenData(tokeninfo)
					//replacer := strings.NewReplacer(" ", "_", ":", "", "?", "", "*", "", "<", "", ">", "", "|", "", "\"", "", "/", "")
					//contractNameFilter := replacer.Replace(tokeninfo.ContractName)

				}
			}

		}

		i = i + 1

	}

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

//if z.Topics[0] == logTransferSingleSigHash { // TransferSingle erc1155

// erc1155instance, err := erc1155.NewErc1155(contractAddress, client)
// if err != nil {
// 	logger.InfoLog("!!!!!!!!!!NewErc1155 erc !!!! == 0 txs.Hash[%s] , error[%s]\n", txhash, err.Error())

// }

// logger.InfoLog("erc1155instance [%v]\n", erc1155instance)

// //arg0 := big.NewInt(1)

// erc1155transfersingle, err := erc1155instance.ParseTransferSingle(*z)
// if err != nil {
// 	logger.InfoLog("!!!!!!!!!!erc1155transfersingle ParseTransferSingle txs.Hash[%s] , error[%s]\n", txhash, err.Error())
// }

// logger.InfoLog("erc1155transfersingle ParseTransferSingle ID [%s] , Value[%d] \n", erc1155transfersingle.Id.String(), erc1155transfersingle.Value.Int64())

// erc1155uri, err := erc1155instance.Uri(&bind.CallOpts{}, erc1155transfersingle.Value)
// if err != nil {
// 	logger.InfoLog("!!!!!!!!!!erc1155instance.Uri txs.Hash[%s] , error[%s]\n", txhash, err.Error())
// }

// erc1155json, err := erc1155transfersingle.Raw.MarshalJSON()
// if err != nil {
// 	logger.InfoLog("!!!!!!!erc1155transfersingle.Raw.MarshalJSON() txs.Hash[%s] , error[%s]\n", txhash, err.Error())
// }
// fmt.Println(" erc1155transfersingle json ", string(erc1155json))

// erc1155uri, err := erc1155instance.Uri(&bind.CallOpts{}, arg0)
// if err != nil {
// 	logger.InfoLog("!!!!!!!!!!erc1155instance.Uri error txs.Hash[%s] , error[%s]\n", txhash, err.Error())
// }

//logger.InfoLog("erc1155instance uri [%s]\n", erc1155uri)

//}
