package common

import (
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"strings"
)

func GetKey(chainID uint64, address string) string {
	return fmt.Sprintf("%d_%s", chainID, address)
}

func GetABI(abiJSON string) (*abi.ABI, error) {
	wrapABI, err := abi.JSON(strings.NewReader(abiJSON))
	if err != nil {
		return nil, err
	}
	return &wrapABI, nil
}
