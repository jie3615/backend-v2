package bounty

import "github.com/ethereum/go-ethereum/ethclient"

// bountyFactory合约create事件topic
// const EventCreated = "0x813d49ce84071dda2c63fce419cce1843e73c55512e307c889d031b57eda73c7"
const EventCreated = "0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0"

type ContractInfo struct {
	ABI         string
	Address     string
	CreatedHash string
	Client      *ethclient.Client
}
