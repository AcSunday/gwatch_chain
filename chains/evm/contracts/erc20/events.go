package erc20

import (
	"github.com/ethereum/go-ethereum/crypto"
	"gwatch_chain/chains/evm/contracts/abs"
)

// TransferEvent
//
// Transfer(address indexed from, address indexed to, uint256 value);
func TransferEvent() abs.Event {
	return abs.Event(crypto.Keccak256Hash([]byte("Transfer(address,address,uint256)")).Hex())
}

// ApprovalEvent
//
// Approval(address indexed owner, address indexed spender, uint256 value);
func ApprovalEvent() abs.Event {
	return abs.Event(crypto.Keccak256Hash([]byte("Approval(address,address,uint256)")).Hex())
}