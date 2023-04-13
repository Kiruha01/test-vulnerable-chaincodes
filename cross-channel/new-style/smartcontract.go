package main

import (
	"fmt"

	"github.com/hyperledger/fabric-contract-api-go/contractapi"
)

type MyChaincode struct {
	contractapi.Contract
}

func (t *MyChaincode) InvokeOtherChaincode(ctx contractapi.TransactionContextInterface, channelID, chaincodeName string, args ...string) ([]byte, error) {
	if len(args) < 1 {
		return nil, fmt.Errorf("expecting at least one argument")
	}

	// Convert args to a byte array
	argsAsBytes := make([][]byte, len(args))
	for i, arg := range args {
		argsAsBytes[i] = []byte(arg)
	}

	response, err := ctx.GetStub().InvokeChaincode(chaincodeName, argsAsBytes, channelID)
	if err != nil {
		return nil, fmt.Errorf("failed to invoke chaincode '%s' on channel '%s': %v", chaincodeName, channelID, err)
	}
	if response.GetStatus() != 200 {
		return nil, fmt.Errorf("failed to invoke chaincode '%s' on channel '%s': %s", chaincodeName, channelID, response.GetMessage())
	}

	return response.Payload, nil
}

func main() {
	chaincode, err := contractapi.NewChaincode(new(MyChaincode))
	if err != nil {
		fmt.Printf("Error creating MyChaincode chaincode: %s", err.Error())
		return
	}

	if err := chaincode.Start(); err != nil {
		fmt.Printf("Error starting MyChaincode chaincode: %s", err.Error())
	}
}
