package startup

import (
	"github.com/ethereum/go-ethereum/ethclient"
)

const EventCreated = "0x822c16987e5c88fd1ec8ce2935c0b5daf646231496d234745d143a9b62673973"

type ContractInfo struct {
	ABI         string
	Address     string
	CreatedHash string
	Client      *ethclient.Client
}
