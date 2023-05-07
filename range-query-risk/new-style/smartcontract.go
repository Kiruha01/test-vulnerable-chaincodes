package main

import (
	"fmt"
	"github.com/hyperledger/fabric-contract-api-go/contractapi"
)

type MyChaincode struct {
	contractapi.Contract
}

func (cc *MyChaincode) Init(ctx contractapi.TransactionContextInterface) error {
	fmt.Println("MyChaincode Init")
	return nil
}

func (cc *MyChaincode) Update(ctx contractapi.TransactionContextInterface, startKey string, endKey string) error {
	stub = ctx.GetStub()
	resultsIterator, err := stub.GetStateByRange(startKey, endKey)
	if err != nil {
		return err
	}
	defer resultsIterator.Close()

	for resultsIterator.HasNext() {
		keyValue, err := resultsIterator.Next()
		if err != nil {
			return err
		}
		err = stub.PutState(keyValue.Key, keyValue.Value)
		if err != nil {
			return err
		}
	}

	return nil
}

func main() {
	err := shim.Start(new(MyChaincode))
	if err != nil {
		fmt.Printf("Error starting MyChaincode chaincode: %s", err.Error())
	}
}
