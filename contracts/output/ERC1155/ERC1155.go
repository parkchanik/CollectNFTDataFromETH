// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package erc1155

import (
	"errors"
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// Erc1155MetaData contains all meta data concerning the Erc1155 contract.
var Erc1155MetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"string\",\"name\":\"uri_\",\"type\":\"string\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"ApprovalForAll\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256[]\",\"name\":\"ids\",\"type\":\"uint256[]\"},{\"indexed\":false,\"internalType\":\"uint256[]\",\"name\":\"values\",\"type\":\"uint256[]\"}],\"name\":\"TransferBatch\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"TransferSingle\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"value\",\"type\":\"string\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"URI\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"accounts\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"ids\",\"type\":\"uint256[]\"}],\"name\":\"balanceOfBatch\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"isApprovedForAll\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256[]\",\"name\":\"ids\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"amounts\",\"type\":\"uint256[]\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"safeBatchTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"setApprovalForAll\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"uri\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x60806040523480156200001157600080fd5b5060405162002910380380620029108339818101604052810190620000379190620002b8565b62000048816200004f60201b60201c565b506200036e565b8060029080519060200190620000679291906200006b565b5050565b828054620000799062000338565b90600052602060002090601f0160209004810192826200009d5760008555620000e9565b82601f10620000b857805160ff1916838001178555620000e9565b82800160010185558215620000e9579182015b82811115620000e8578251825591602001919060010190620000cb565b5b509050620000f89190620000fc565b5090565b5b8082111562000117576000816000905550600101620000fd565b5090565b6000604051905090565b600080fd5b600080fd5b600080fd5b600080fd5b6000601f19601f8301169050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b620001848262000139565b810181811067ffffffffffffffff82111715620001a657620001a56200014a565b5b80604052505050565b6000620001bb6200011b565b9050620001c9828262000179565b919050565b600067ffffffffffffffff821115620001ec57620001eb6200014a565b5b620001f78262000139565b9050602081019050919050565b60005b838110156200022457808201518184015260208101905062000207565b8381111562000234576000848401525b50505050565b6000620002516200024b84620001ce565b620001af565b90508281526020810184848401111562000270576200026f62000134565b5b6200027d84828562000204565b509392505050565b600082601f8301126200029d576200029c6200012f565b5b8151620002af8482602086016200023a565b91505092915050565b600060208284031215620002d157620002d062000125565b5b600082015167ffffffffffffffff811115620002f257620002f16200012a565b5b620003008482850162000285565b91505092915050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b600060028204905060018216806200035157607f821691505b6020821081141562000368576200036762000309565b5b50919050565b612592806200037e6000396000f3fe608060405234801561001057600080fd5b50600436106100875760003560e01c80634e1273f41161005b5780634e1273f414610138578063a22cb46514610168578063e985e9c514610184578063f242432a146101b457610087565b8062fdd58e1461008c57806301ffc9a7146100bc5780630e89341c146100ec5780632eb2c2d61461011c575b600080fd5b6100a660048036038101906100a19190611386565b6101d0565b6040516100b391906113d5565b60405180910390f35b6100d660048036038101906100d19190611448565b610299565b6040516100e39190611490565b60405180910390f35b610106600480360381019061010191906114ab565b61037b565b6040516101139190611571565b60405180910390f35b61013660048036038101906101319190611790565b61040f565b005b610152600480360381019061014d9190611922565b6104b0565b60405161015f9190611a58565b60405180910390f35b610182600480360381019061017d9190611aa6565b6105c9565b005b61019e60048036038101906101999190611ae6565b6105df565b6040516101ab9190611490565b60405180910390f35b6101ce60048036038101906101c99190611b26565b610673565b005b60008073ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff161415610241576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161023890611c2f565b60405180910390fd5b60008083815260200190815260200160002060008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002054905092915050565b60007fd9b67a26000000000000000000000000000000000000000000000000000000007bffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916827bffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916148061036457507f0e89341c000000000000000000000000000000000000000000000000000000007bffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916827bffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916145b80610374575061037382610714565b5b9050919050565b60606002805461038a90611c7e565b80601f01602080910402602001604051908101604052809291908181526020018280546103b690611c7e565b80156104035780601f106103d857610100808354040283529160200191610403565b820191906000526020600020905b8154815290600101906020018083116103e657829003601f168201915b50505050509050919050565b61041761077e565b73ffffffffffffffffffffffffffffffffffffffff168573ffffffffffffffffffffffffffffffffffffffff16148061045d575061045c8561045761077e565b6105df565b5b61049c576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161049390611d22565b60405180910390fd5b6104a98585858585610786565b5050505050565b606081518351146104f6576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016104ed90611db4565b60405180910390fd5b6000835167ffffffffffffffff81111561051357610512611598565b5b6040519080825280602002602001820160405280156105415781602001602082028036833780820191505090505b50905060005b84518110156105be5761058e85828151811061056657610565611dd4565b5b602002602001015185838151811061058157610580611dd4565b5b60200260200101516101d0565b8282815181106105a1576105a0611dd4565b5b602002602001018181525050806105b790611e32565b9050610547565b508091505092915050565b6105db6105d461077e565b8383610a9a565b5050565b6000600160008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900460ff16905092915050565b61067b61077e565b73ffffffffffffffffffffffffffffffffffffffff168573ffffffffffffffffffffffffffffffffffffffff1614806106c157506106c0856106bb61077e565b6105df565b5b610700576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016106f790611eed565b60405180910390fd5b61070d8585858585610c07565b5050505050565b60007f01ffc9a7000000000000000000000000000000000000000000000000000000007bffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916827bffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916149050919050565b600033905090565b81518351146107ca576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016107c190611f7f565b60405180910390fd5b600073ffffffffffffffffffffffffffffffffffffffff168473ffffffffffffffffffffffffffffffffffffffff16141561083a576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161083190612011565b60405180910390fd5b600061084461077e565b9050610854818787878787610e89565b60005b8451811015610a0557600085828151811061087557610874611dd4565b5b60200260200101519050600085838151811061089457610893611dd4565b5b60200260200101519050600080600084815260200190815260200160002060008b73ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002054905081811015610935576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161092c906120a3565b60405180910390fd5b81810360008085815260200190815260200160002060008c73ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020819055508160008085815260200190815260200160002060008b73ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008282546109ea91906120c3565b92505081905550505050806109fe90611e32565b9050610857565b508473ffffffffffffffffffffffffffffffffffffffff168673ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff167f4a39dc06d4c0dbc64b70af90fd698a233a518aa5d07e595d983b8c0526c8f7fb8787604051610a7c929190612119565b60405180910390a4610a92818787878787610e91565b505050505050565b8173ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff161415610b09576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610b00906121c2565b60405180910390fd5b80600160008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548160ff0219169083151502179055508173ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff167f17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c3183604051610bfa9190611490565b60405180910390a3505050565b600073ffffffffffffffffffffffffffffffffffffffff168473ffffffffffffffffffffffffffffffffffffffff161415610c77576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610c6e90612011565b60405180910390fd5b6000610c8161077e565b9050610ca1818787610c9288611069565b610c9b88611069565b87610e89565b600080600086815260200190815260200160002060008873ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002054905083811015610d38576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610d2f906120a3565b60405180910390fd5b83810360008087815260200190815260200160002060008973ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020819055508360008087815260200190815260200160002060008873ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206000828254610ded91906120c3565b925050819055508573ffffffffffffffffffffffffffffffffffffffff168773ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff167fc3d58168c5ae7397731d063d5bbf3d657854427343f4c083240f7aacaa2d0f628888604051610e6a9291906121e2565b60405180910390a4610e808288888888886110e3565b50505050505050565b505050505050565b610eb08473ffffffffffffffffffffffffffffffffffffffff166112bb565b15611061578373ffffffffffffffffffffffffffffffffffffffff1663bc197c8187878686866040518663ffffffff1660e01b8152600401610ef695949392919061226f565b6020604051808303816000875af1925050508015610f3257506040513d601f19601f82011682018060405250810190610f2f91906122ec565b60015b610fd857610f3e612326565b806308c379a01415610f9b5750610f53612348565b80610f5e5750610f9d565b806040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610f929190611571565b60405180910390fd5b505b6040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610fcf90612450565b60405180910390fd5b63bc197c8160e01b7bffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916817bffffffffffffffffffffffffffffffffffffffffffffffffffffffff19161461105f576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401611056906124e2565b60405180910390fd5b505b505050505050565b60606000600167ffffffffffffffff81111561108857611087611598565b5b6040519080825280602002602001820160405280156110b65781602001602082028036833780820191505090505b50905082816000815181106110ce576110cd611dd4565b5b60200260200101818152505080915050919050565b6111028473ffffffffffffffffffffffffffffffffffffffff166112bb565b156112b3578373ffffffffffffffffffffffffffffffffffffffff1663f23a6e6187878686866040518663ffffffff1660e01b8152600401611148959493929190612502565b6020604051808303816000875af192505050801561118457506040513d601f19601f8201168201806040525081019061118191906122ec565b60015b61122a57611190612326565b806308c379a014156111ed57506111a5612348565b806111b057506111ef565b806040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016111e49190611571565b60405180910390fd5b505b6040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161122190612450565b60405180910390fd5b63f23a6e6160e01b7bffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916817bffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916146112b1576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016112a8906124e2565b60405180910390fd5b505b505050505050565b6000808273ffffffffffffffffffffffffffffffffffffffff163b119050919050565b6000604051905090565b600080fd5b600080fd5b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b600061131d826112f2565b9050919050565b61132d81611312565b811461133857600080fd5b50565b60008135905061134a81611324565b92915050565b6000819050919050565b61136381611350565b811461136e57600080fd5b50565b6000813590506113808161135a565b92915050565b6000806040838503121561139d5761139c6112e8565b5b60006113ab8582860161133b565b92505060206113bc85828601611371565b9150509250929050565b6113cf81611350565b82525050565b60006020820190506113ea60008301846113c6565b92915050565b60007fffffffff0000000000000000000000000000000000000000000000000000000082169050919050565b611425816113f0565b811461143057600080fd5b50565b6000813590506114428161141c565b92915050565b60006020828403121561145e5761145d6112e8565b5b600061146c84828501611433565b91505092915050565b60008115159050919050565b61148a81611475565b82525050565b60006020820190506114a56000830184611481565b92915050565b6000602082840312156114c1576114c06112e8565b5b60006114cf84828501611371565b91505092915050565b600081519050919050565b600082825260208201905092915050565b60005b838110156115125780820151818401526020810190506114f7565b83811115611521576000848401525b50505050565b6000601f19601f8301169050919050565b6000611543826114d8565b61154d81856114e3565b935061155d8185602086016114f4565b61156681611527565b840191505092915050565b6000602082019050818103600083015261158b8184611538565b905092915050565b600080fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b6115d082611527565b810181811067ffffffffffffffff821117156115ef576115ee611598565b5b80604052505050565b60006116026112de565b905061160e82826115c7565b919050565b600067ffffffffffffffff82111561162e5761162d611598565b5b602082029050602081019050919050565b600080fd5b600061165761165284611613565b6115f8565b9050808382526020820190506020840283018581111561167a5761167961163f565b5b835b818110156116a3578061168f8882611371565b84526020840193505060208101905061167c565b5050509392505050565b600082601f8301126116c2576116c1611593565b5b81356116d2848260208601611644565b91505092915050565b600080fd5b600067ffffffffffffffff8211156116fb576116fa611598565b5b61170482611527565b9050602081019050919050565b82818337600083830152505050565b600061173361172e846116e0565b6115f8565b90508281526020810184848401111561174f5761174e6116db565b5b61175a848285611711565b509392505050565b600082601f83011261177757611776611593565b5b8135611787848260208601611720565b91505092915050565b600080600080600060a086880312156117ac576117ab6112e8565b5b60006117ba8882890161133b565b95505060206117cb8882890161133b565b945050604086013567ffffffffffffffff8111156117ec576117eb6112ed565b5b6117f8888289016116ad565b935050606086013567ffffffffffffffff811115611819576118186112ed565b5b611825888289016116ad565b925050608086013567ffffffffffffffff811115611846576118456112ed565b5b61185288828901611762565b9150509295509295909350565b600067ffffffffffffffff82111561187a57611879611598565b5b602082029050602081019050919050565b600061189e6118998461185f565b6115f8565b905080838252602082019050602084028301858111156118c1576118c061163f565b5b835b818110156118ea57806118d6888261133b565b8452602084019350506020810190506118c3565b5050509392505050565b600082601f83011261190957611908611593565b5b813561191984826020860161188b565b91505092915050565b60008060408385031215611939576119386112e8565b5b600083013567ffffffffffffffff811115611957576119566112ed565b5b611963858286016118f4565b925050602083013567ffffffffffffffff811115611984576119836112ed565b5b611990858286016116ad565b9150509250929050565b600081519050919050565b600082825260208201905092915050565b6000819050602082019050919050565b6119cf81611350565b82525050565b60006119e183836119c6565b60208301905092915050565b6000602082019050919050565b6000611a058261199a565b611a0f81856119a5565b9350611a1a836119b6565b8060005b83811015611a4b578151611a3288826119d5565b9750611a3d836119ed565b925050600181019050611a1e565b5085935050505092915050565b60006020820190508181036000830152611a7281846119fa565b905092915050565b611a8381611475565b8114611a8e57600080fd5b50565b600081359050611aa081611a7a565b92915050565b60008060408385031215611abd57611abc6112e8565b5b6000611acb8582860161133b565b9250506020611adc85828601611a91565b9150509250929050565b60008060408385031215611afd57611afc6112e8565b5b6000611b0b8582860161133b565b9250506020611b1c8582860161133b565b9150509250929050565b600080600080600060a08688031215611b4257611b416112e8565b5b6000611b508882890161133b565b9550506020611b618882890161133b565b9450506040611b7288828901611371565b9350506060611b8388828901611371565b925050608086013567ffffffffffffffff811115611ba457611ba36112ed565b5b611bb088828901611762565b9150509295509295909350565b7f455243313135353a2062616c616e636520717565727920666f7220746865207a60008201527f65726f2061646472657373000000000000000000000000000000000000000000602082015250565b6000611c19602b836114e3565b9150611c2482611bbd565b604082019050919050565b60006020820190508181036000830152611c4881611c0c565b9050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b60006002820490506001821680611c9657607f821691505b60208210811415611caa57611ca9611c4f565b5b50919050565b7f455243313135353a207472616e736665722063616c6c6572206973206e6f742060008201527f6f776e6572206e6f7220617070726f7665640000000000000000000000000000602082015250565b6000611d0c6032836114e3565b9150611d1782611cb0565b604082019050919050565b60006020820190508181036000830152611d3b81611cff565b9050919050565b7f455243313135353a206163636f756e747320616e6420696473206c656e67746860008201527f206d69736d617463680000000000000000000000000000000000000000000000602082015250565b6000611d9e6029836114e3565b9150611da982611d42565b604082019050919050565b60006020820190508181036000830152611dcd81611d91565b9050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b6000611e3d82611350565b91507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff821415611e7057611e6f611e03565b5b600182019050919050565b7f455243313135353a2063616c6c6572206973206e6f74206f776e6572206e6f7260008201527f20617070726f7665640000000000000000000000000000000000000000000000602082015250565b6000611ed76029836114e3565b9150611ee282611e7b565b604082019050919050565b60006020820190508181036000830152611f0681611eca565b9050919050565b7f455243313135353a2069647320616e6420616d6f756e7473206c656e6774682060008201527f6d69736d61746368000000000000000000000000000000000000000000000000602082015250565b6000611f696028836114e3565b9150611f7482611f0d565b604082019050919050565b60006020820190508181036000830152611f9881611f5c565b9050919050565b7f455243313135353a207472616e7366657220746f20746865207a65726f20616460008201527f6472657373000000000000000000000000000000000000000000000000000000602082015250565b6000611ffb6025836114e3565b915061200682611f9f565b604082019050919050565b6000602082019050818103600083015261202a81611fee565b9050919050565b7f455243313135353a20696e73756666696369656e742062616c616e636520666f60008201527f72207472616e7366657200000000000000000000000000000000000000000000602082015250565b600061208d602a836114e3565b915061209882612031565b604082019050919050565b600060208201905081810360008301526120bc81612080565b9050919050565b60006120ce82611350565b91506120d983611350565b9250827fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0382111561210e5761210d611e03565b5b828201905092915050565b6000604082019050818103600083015261213381856119fa565b9050818103602083015261214781846119fa565b90509392505050565b7f455243313135353a2073657474696e6720617070726f76616c2073746174757360008201527f20666f722073656c660000000000000000000000000000000000000000000000602082015250565b60006121ac6029836114e3565b91506121b782612150565b604082019050919050565b600060208201905081810360008301526121db8161219f565b9050919050565b60006040820190506121f760008301856113c6565b61220460208301846113c6565b9392505050565b61221481611312565b82525050565b600081519050919050565b600082825260208201905092915050565b60006122418261221a565b61224b8185612225565b935061225b8185602086016114f4565b61226481611527565b840191505092915050565b600060a082019050612284600083018861220b565b612291602083018761220b565b81810360408301526122a381866119fa565b905081810360608301526122b781856119fa565b905081810360808301526122cb8184612236565b90509695505050505050565b6000815190506122e68161141c565b92915050565b600060208284031215612302576123016112e8565b5b6000612310848285016122d7565b91505092915050565b60008160e01c9050919050565b600060033d11156123455760046000803e612342600051612319565b90505b90565b600060443d1015612358576123db565b6123606112de565b60043d036004823e80513d602482011167ffffffffffffffff821117156123885750506123db565b808201805167ffffffffffffffff8111156123a657505050506123db565b80602083010160043d0385018111156123c35750505050506123db565b6123d2826020018501866115c7565b82955050505050505b90565b7f455243313135353a207472616e7366657220746f206e6f6e204552433131353560008201527f526563656976657220696d706c656d656e746572000000000000000000000000602082015250565b600061243a6034836114e3565b9150612445826123de565b604082019050919050565b600060208201905081810360008301526124698161242d565b9050919050565b7f455243313135353a204552433131353552656365697665722072656a6563746560008201527f6420746f6b656e73000000000000000000000000000000000000000000000000602082015250565b60006124cc6028836114e3565b91506124d782612470565b604082019050919050565b600060208201905081810360008301526124fb816124bf565b9050919050565b600060a082019050612517600083018861220b565b612524602083018761220b565b61253160408301866113c6565b61253e60608301856113c6565b81810360808301526125508184612236565b9050969550505050505056fea2646970667358221220d47404d71696ac22d0710101e1892f17549ee0ad464224bfce63226dbc448fcb64736f6c634300080a0033",
}

// Erc1155ABI is the input ABI used to generate the binding from.
// Deprecated: Use Erc1155MetaData.ABI instead.
var Erc1155ABI = Erc1155MetaData.ABI

// Erc1155Bin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use Erc1155MetaData.Bin instead.
var Erc1155Bin = Erc1155MetaData.Bin

// DeployErc1155 deploys a new Ethereum contract, binding an instance of Erc1155 to it.
func DeployErc1155(auth *bind.TransactOpts, backend bind.ContractBackend, uri_ string) (common.Address, *types.Transaction, *Erc1155, error) {
	parsed, err := Erc1155MetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(Erc1155Bin), backend, uri_)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Erc1155{Erc1155Caller: Erc1155Caller{contract: contract}, Erc1155Transactor: Erc1155Transactor{contract: contract}, Erc1155Filterer: Erc1155Filterer{contract: contract}}, nil
}

// Erc1155 is an auto generated Go binding around an Ethereum contract.
type Erc1155 struct {
	Erc1155Caller     // Read-only binding to the contract
	Erc1155Transactor // Write-only binding to the contract
	Erc1155Filterer   // Log filterer for contract events
}

// Erc1155Caller is an auto generated read-only Go binding around an Ethereum contract.
type Erc1155Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Erc1155Transactor is an auto generated write-only Go binding around an Ethereum contract.
type Erc1155Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Erc1155Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type Erc1155Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Erc1155Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type Erc1155Session struct {
	Contract     *Erc1155          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// Erc1155CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type Erc1155CallerSession struct {
	Contract *Erc1155Caller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// Erc1155TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type Erc1155TransactorSession struct {
	Contract     *Erc1155Transactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// Erc1155Raw is an auto generated low-level Go binding around an Ethereum contract.
type Erc1155Raw struct {
	Contract *Erc1155 // Generic contract binding to access the raw methods on
}

// Erc1155CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type Erc1155CallerRaw struct {
	Contract *Erc1155Caller // Generic read-only contract binding to access the raw methods on
}

// Erc1155TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type Erc1155TransactorRaw struct {
	Contract *Erc1155Transactor // Generic write-only contract binding to access the raw methods on
}

// NewErc1155 creates a new instance of Erc1155, bound to a specific deployed contract.
func NewErc1155(address common.Address, backend bind.ContractBackend) (*Erc1155, error) {
	contract, err := bindErc1155(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Erc1155{Erc1155Caller: Erc1155Caller{contract: contract}, Erc1155Transactor: Erc1155Transactor{contract: contract}, Erc1155Filterer: Erc1155Filterer{contract: contract}}, nil
}

// NewErc1155Caller creates a new read-only instance of Erc1155, bound to a specific deployed contract.
func NewErc1155Caller(address common.Address, caller bind.ContractCaller) (*Erc1155Caller, error) {
	contract, err := bindErc1155(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &Erc1155Caller{contract: contract}, nil
}

// NewErc1155Transactor creates a new write-only instance of Erc1155, bound to a specific deployed contract.
func NewErc1155Transactor(address common.Address, transactor bind.ContractTransactor) (*Erc1155Transactor, error) {
	contract, err := bindErc1155(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &Erc1155Transactor{contract: contract}, nil
}

// NewErc1155Filterer creates a new log filterer instance of Erc1155, bound to a specific deployed contract.
func NewErc1155Filterer(address common.Address, filterer bind.ContractFilterer) (*Erc1155Filterer, error) {
	contract, err := bindErc1155(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &Erc1155Filterer{contract: contract}, nil
}

// bindErc1155 binds a generic wrapper to an already deployed contract.
func bindErc1155(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(Erc1155ABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Erc1155 *Erc1155Raw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Erc1155.Contract.Erc1155Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Erc1155 *Erc1155Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Erc1155.Contract.Erc1155Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Erc1155 *Erc1155Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Erc1155.Contract.Erc1155Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Erc1155 *Erc1155CallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Erc1155.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Erc1155 *Erc1155TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Erc1155.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Erc1155 *Erc1155TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Erc1155.Contract.contract.Transact(opts, method, params...)
}

// BalanceOf is a free data retrieval call binding the contract method 0x00fdd58e.
//
// Solidity: function balanceOf(address account, uint256 id) view returns(uint256)
func (_Erc1155 *Erc1155Caller) BalanceOf(opts *bind.CallOpts, account common.Address, id *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Erc1155.contract.Call(opts, &out, "balanceOf", account, id)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x00fdd58e.
//
// Solidity: function balanceOf(address account, uint256 id) view returns(uint256)
func (_Erc1155 *Erc1155Session) BalanceOf(account common.Address, id *big.Int) (*big.Int, error) {
	return _Erc1155.Contract.BalanceOf(&_Erc1155.CallOpts, account, id)
}

// BalanceOf is a free data retrieval call binding the contract method 0x00fdd58e.
//
// Solidity: function balanceOf(address account, uint256 id) view returns(uint256)
func (_Erc1155 *Erc1155CallerSession) BalanceOf(account common.Address, id *big.Int) (*big.Int, error) {
	return _Erc1155.Contract.BalanceOf(&_Erc1155.CallOpts, account, id)
}

// BalanceOfBatch is a free data retrieval call binding the contract method 0x4e1273f4.
//
// Solidity: function balanceOfBatch(address[] accounts, uint256[] ids) view returns(uint256[])
func (_Erc1155 *Erc1155Caller) BalanceOfBatch(opts *bind.CallOpts, accounts []common.Address, ids []*big.Int) ([]*big.Int, error) {
	var out []interface{}
	err := _Erc1155.contract.Call(opts, &out, "balanceOfBatch", accounts, ids)

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// BalanceOfBatch is a free data retrieval call binding the contract method 0x4e1273f4.
//
// Solidity: function balanceOfBatch(address[] accounts, uint256[] ids) view returns(uint256[])
func (_Erc1155 *Erc1155Session) BalanceOfBatch(accounts []common.Address, ids []*big.Int) ([]*big.Int, error) {
	return _Erc1155.Contract.BalanceOfBatch(&_Erc1155.CallOpts, accounts, ids)
}

// BalanceOfBatch is a free data retrieval call binding the contract method 0x4e1273f4.
//
// Solidity: function balanceOfBatch(address[] accounts, uint256[] ids) view returns(uint256[])
func (_Erc1155 *Erc1155CallerSession) BalanceOfBatch(accounts []common.Address, ids []*big.Int) ([]*big.Int, error) {
	return _Erc1155.Contract.BalanceOfBatch(&_Erc1155.CallOpts, accounts, ids)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address account, address operator) view returns(bool)
func (_Erc1155 *Erc1155Caller) IsApprovedForAll(opts *bind.CallOpts, account common.Address, operator common.Address) (bool, error) {
	var out []interface{}
	err := _Erc1155.contract.Call(opts, &out, "isApprovedForAll", account, operator)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address account, address operator) view returns(bool)
func (_Erc1155 *Erc1155Session) IsApprovedForAll(account common.Address, operator common.Address) (bool, error) {
	return _Erc1155.Contract.IsApprovedForAll(&_Erc1155.CallOpts, account, operator)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address account, address operator) view returns(bool)
func (_Erc1155 *Erc1155CallerSession) IsApprovedForAll(account common.Address, operator common.Address) (bool, error) {
	return _Erc1155.Contract.IsApprovedForAll(&_Erc1155.CallOpts, account, operator)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Erc1155 *Erc1155Caller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _Erc1155.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Erc1155 *Erc1155Session) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _Erc1155.Contract.SupportsInterface(&_Erc1155.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Erc1155 *Erc1155CallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _Erc1155.Contract.SupportsInterface(&_Erc1155.CallOpts, interfaceId)
}

// Uri is a free data retrieval call binding the contract method 0x0e89341c.
//
// Solidity: function uri(uint256 ) view returns(string)
func (_Erc1155 *Erc1155Caller) Uri(opts *bind.CallOpts, arg0 *big.Int) (string, error) {
	var out []interface{}
	err := _Erc1155.contract.Call(opts, &out, "uri", arg0)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Uri is a free data retrieval call binding the contract method 0x0e89341c.
//
// Solidity: function uri(uint256 ) view returns(string)
func (_Erc1155 *Erc1155Session) Uri(arg0 *big.Int) (string, error) {
	return _Erc1155.Contract.Uri(&_Erc1155.CallOpts, arg0)
}

// Uri is a free data retrieval call binding the contract method 0x0e89341c.
//
// Solidity: function uri(uint256 ) view returns(string)
func (_Erc1155 *Erc1155CallerSession) Uri(arg0 *big.Int) (string, error) {
	return _Erc1155.Contract.Uri(&_Erc1155.CallOpts, arg0)
}

// SafeBatchTransferFrom is a paid mutator transaction binding the contract method 0x2eb2c2d6.
//
// Solidity: function safeBatchTransferFrom(address from, address to, uint256[] ids, uint256[] amounts, bytes data) returns()
func (_Erc1155 *Erc1155Transactor) SafeBatchTransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, ids []*big.Int, amounts []*big.Int, data []byte) (*types.Transaction, error) {
	return _Erc1155.contract.Transact(opts, "safeBatchTransferFrom", from, to, ids, amounts, data)
}

// SafeBatchTransferFrom is a paid mutator transaction binding the contract method 0x2eb2c2d6.
//
// Solidity: function safeBatchTransferFrom(address from, address to, uint256[] ids, uint256[] amounts, bytes data) returns()
func (_Erc1155 *Erc1155Session) SafeBatchTransferFrom(from common.Address, to common.Address, ids []*big.Int, amounts []*big.Int, data []byte) (*types.Transaction, error) {
	return _Erc1155.Contract.SafeBatchTransferFrom(&_Erc1155.TransactOpts, from, to, ids, amounts, data)
}

// SafeBatchTransferFrom is a paid mutator transaction binding the contract method 0x2eb2c2d6.
//
// Solidity: function safeBatchTransferFrom(address from, address to, uint256[] ids, uint256[] amounts, bytes data) returns()
func (_Erc1155 *Erc1155TransactorSession) SafeBatchTransferFrom(from common.Address, to common.Address, ids []*big.Int, amounts []*big.Int, data []byte) (*types.Transaction, error) {
	return _Erc1155.Contract.SafeBatchTransferFrom(&_Erc1155.TransactOpts, from, to, ids, amounts, data)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0xf242432a.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 id, uint256 amount, bytes data) returns()
func (_Erc1155 *Erc1155Transactor) SafeTransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, id *big.Int, amount *big.Int, data []byte) (*types.Transaction, error) {
	return _Erc1155.contract.Transact(opts, "safeTransferFrom", from, to, id, amount, data)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0xf242432a.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 id, uint256 amount, bytes data) returns()
func (_Erc1155 *Erc1155Session) SafeTransferFrom(from common.Address, to common.Address, id *big.Int, amount *big.Int, data []byte) (*types.Transaction, error) {
	return _Erc1155.Contract.SafeTransferFrom(&_Erc1155.TransactOpts, from, to, id, amount, data)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0xf242432a.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 id, uint256 amount, bytes data) returns()
func (_Erc1155 *Erc1155TransactorSession) SafeTransferFrom(from common.Address, to common.Address, id *big.Int, amount *big.Int, data []byte) (*types.Transaction, error) {
	return _Erc1155.Contract.SafeTransferFrom(&_Erc1155.TransactOpts, from, to, id, amount, data)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_Erc1155 *Erc1155Transactor) SetApprovalForAll(opts *bind.TransactOpts, operator common.Address, approved bool) (*types.Transaction, error) {
	return _Erc1155.contract.Transact(opts, "setApprovalForAll", operator, approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_Erc1155 *Erc1155Session) SetApprovalForAll(operator common.Address, approved bool) (*types.Transaction, error) {
	return _Erc1155.Contract.SetApprovalForAll(&_Erc1155.TransactOpts, operator, approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_Erc1155 *Erc1155TransactorSession) SetApprovalForAll(operator common.Address, approved bool) (*types.Transaction, error) {
	return _Erc1155.Contract.SetApprovalForAll(&_Erc1155.TransactOpts, operator, approved)
}

// Erc1155ApprovalForAllIterator is returned from FilterApprovalForAll and is used to iterate over the raw logs and unpacked data for ApprovalForAll events raised by the Erc1155 contract.
type Erc1155ApprovalForAllIterator struct {
	Event *Erc1155ApprovalForAll // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *Erc1155ApprovalForAllIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Erc1155ApprovalForAll)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(Erc1155ApprovalForAll)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *Erc1155ApprovalForAllIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Erc1155ApprovalForAllIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Erc1155ApprovalForAll represents a ApprovalForAll event raised by the Erc1155 contract.
type Erc1155ApprovalForAll struct {
	Account  common.Address
	Operator common.Address
	Approved bool
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApprovalForAll is a free log retrieval operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed account, address indexed operator, bool approved)
func (_Erc1155 *Erc1155Filterer) FilterApprovalForAll(opts *bind.FilterOpts, account []common.Address, operator []common.Address) (*Erc1155ApprovalForAllIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _Erc1155.contract.FilterLogs(opts, "ApprovalForAll", accountRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return &Erc1155ApprovalForAllIterator{contract: _Erc1155.contract, event: "ApprovalForAll", logs: logs, sub: sub}, nil
}

// WatchApprovalForAll is a free log subscription operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed account, address indexed operator, bool approved)
func (_Erc1155 *Erc1155Filterer) WatchApprovalForAll(opts *bind.WatchOpts, sink chan<- *Erc1155ApprovalForAll, account []common.Address, operator []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _Erc1155.contract.WatchLogs(opts, "ApprovalForAll", accountRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Erc1155ApprovalForAll)
				if err := _Erc1155.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseApprovalForAll is a log parse operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed account, address indexed operator, bool approved)
func (_Erc1155 *Erc1155Filterer) ParseApprovalForAll(log types.Log) (*Erc1155ApprovalForAll, error) {
	event := new(Erc1155ApprovalForAll)
	if err := _Erc1155.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Erc1155TransferBatchIterator is returned from FilterTransferBatch and is used to iterate over the raw logs and unpacked data for TransferBatch events raised by the Erc1155 contract.
type Erc1155TransferBatchIterator struct {
	Event *Erc1155TransferBatch // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *Erc1155TransferBatchIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Erc1155TransferBatch)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(Erc1155TransferBatch)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *Erc1155TransferBatchIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Erc1155TransferBatchIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Erc1155TransferBatch represents a TransferBatch event raised by the Erc1155 contract.
type Erc1155TransferBatch struct {
	Operator common.Address
	From     common.Address
	To       common.Address
	Ids      []*big.Int
	Values   []*big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterTransferBatch is a free log retrieval operation binding the contract event 0x4a39dc06d4c0dbc64b70af90fd698a233a518aa5d07e595d983b8c0526c8f7fb.
//
// Solidity: event TransferBatch(address indexed operator, address indexed from, address indexed to, uint256[] ids, uint256[] values)
func (_Erc1155 *Erc1155Filterer) FilterTransferBatch(opts *bind.FilterOpts, operator []common.Address, from []common.Address, to []common.Address) (*Erc1155TransferBatchIterator, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}
	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Erc1155.contract.FilterLogs(opts, "TransferBatch", operatorRule, fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &Erc1155TransferBatchIterator{contract: _Erc1155.contract, event: "TransferBatch", logs: logs, sub: sub}, nil
}

// WatchTransferBatch is a free log subscription operation binding the contract event 0x4a39dc06d4c0dbc64b70af90fd698a233a518aa5d07e595d983b8c0526c8f7fb.
//
// Solidity: event TransferBatch(address indexed operator, address indexed from, address indexed to, uint256[] ids, uint256[] values)
func (_Erc1155 *Erc1155Filterer) WatchTransferBatch(opts *bind.WatchOpts, sink chan<- *Erc1155TransferBatch, operator []common.Address, from []common.Address, to []common.Address) (event.Subscription, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}
	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Erc1155.contract.WatchLogs(opts, "TransferBatch", operatorRule, fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Erc1155TransferBatch)
				if err := _Erc1155.contract.UnpackLog(event, "TransferBatch", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseTransferBatch is a log parse operation binding the contract event 0x4a39dc06d4c0dbc64b70af90fd698a233a518aa5d07e595d983b8c0526c8f7fb.
//
// Solidity: event TransferBatch(address indexed operator, address indexed from, address indexed to, uint256[] ids, uint256[] values)
func (_Erc1155 *Erc1155Filterer) ParseTransferBatch(log types.Log) (*Erc1155TransferBatch, error) {
	event := new(Erc1155TransferBatch)
	if err := _Erc1155.contract.UnpackLog(event, "TransferBatch", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Erc1155TransferSingleIterator is returned from FilterTransferSingle and is used to iterate over the raw logs and unpacked data for TransferSingle events raised by the Erc1155 contract.
type Erc1155TransferSingleIterator struct {
	Event *Erc1155TransferSingle // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *Erc1155TransferSingleIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Erc1155TransferSingle)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(Erc1155TransferSingle)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *Erc1155TransferSingleIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Erc1155TransferSingleIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Erc1155TransferSingle represents a TransferSingle event raised by the Erc1155 contract.
type Erc1155TransferSingle struct {
	Operator common.Address
	From     common.Address
	To       common.Address
	Id       *big.Int
	Value    *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterTransferSingle is a free log retrieval operation binding the contract event 0xc3d58168c5ae7397731d063d5bbf3d657854427343f4c083240f7aacaa2d0f62.
//
// Solidity: event TransferSingle(address indexed operator, address indexed from, address indexed to, uint256 id, uint256 value)
func (_Erc1155 *Erc1155Filterer) FilterTransferSingle(opts *bind.FilterOpts, operator []common.Address, from []common.Address, to []common.Address) (*Erc1155TransferSingleIterator, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}
	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Erc1155.contract.FilterLogs(opts, "TransferSingle", operatorRule, fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &Erc1155TransferSingleIterator{contract: _Erc1155.contract, event: "TransferSingle", logs: logs, sub: sub}, nil
}

// WatchTransferSingle is a free log subscription operation binding the contract event 0xc3d58168c5ae7397731d063d5bbf3d657854427343f4c083240f7aacaa2d0f62.
//
// Solidity: event TransferSingle(address indexed operator, address indexed from, address indexed to, uint256 id, uint256 value)
func (_Erc1155 *Erc1155Filterer) WatchTransferSingle(opts *bind.WatchOpts, sink chan<- *Erc1155TransferSingle, operator []common.Address, from []common.Address, to []common.Address) (event.Subscription, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}
	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Erc1155.contract.WatchLogs(opts, "TransferSingle", operatorRule, fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Erc1155TransferSingle)
				if err := _Erc1155.contract.UnpackLog(event, "TransferSingle", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseTransferSingle is a log parse operation binding the contract event 0xc3d58168c5ae7397731d063d5bbf3d657854427343f4c083240f7aacaa2d0f62.
//
// Solidity: event TransferSingle(address indexed operator, address indexed from, address indexed to, uint256 id, uint256 value)
func (_Erc1155 *Erc1155Filterer) ParseTransferSingle(log types.Log) (*Erc1155TransferSingle, error) {
	event := new(Erc1155TransferSingle)
	if err := _Erc1155.contract.UnpackLog(event, "TransferSingle", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Erc1155URIIterator is returned from FilterURI and is used to iterate over the raw logs and unpacked data for URI events raised by the Erc1155 contract.
type Erc1155URIIterator struct {
	Event *Erc1155URI // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *Erc1155URIIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Erc1155URI)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(Erc1155URI)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *Erc1155URIIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Erc1155URIIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Erc1155URI represents a URI event raised by the Erc1155 contract.
type Erc1155URI struct {
	Value string
	Id    *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterURI is a free log retrieval operation binding the contract event 0x6bb7ff708619ba0610cba295a58592e0451dee2622938c8755667688daf3529b.
//
// Solidity: event URI(string value, uint256 indexed id)
func (_Erc1155 *Erc1155Filterer) FilterURI(opts *bind.FilterOpts, id []*big.Int) (*Erc1155URIIterator, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _Erc1155.contract.FilterLogs(opts, "URI", idRule)
	if err != nil {
		return nil, err
	}
	return &Erc1155URIIterator{contract: _Erc1155.contract, event: "URI", logs: logs, sub: sub}, nil
}

// WatchURI is a free log subscription operation binding the contract event 0x6bb7ff708619ba0610cba295a58592e0451dee2622938c8755667688daf3529b.
//
// Solidity: event URI(string value, uint256 indexed id)
func (_Erc1155 *Erc1155Filterer) WatchURI(opts *bind.WatchOpts, sink chan<- *Erc1155URI, id []*big.Int) (event.Subscription, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _Erc1155.contract.WatchLogs(opts, "URI", idRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Erc1155URI)
				if err := _Erc1155.contract.UnpackLog(event, "URI", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseURI is a log parse operation binding the contract event 0x6bb7ff708619ba0610cba295a58592e0451dee2622938c8755667688daf3529b.
//
// Solidity: event URI(string value, uint256 indexed id)
func (_Erc1155 *Erc1155Filterer) ParseURI(log types.Log) (*Erc1155URI, error) {
	event := new(Erc1155URI)
	if err := _Erc1155.contract.UnpackLog(event, "URI", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
