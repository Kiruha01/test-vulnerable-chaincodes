package main

import (
    "fmt"
    "github.com/hyperledger/fabric-contract-api-go/contractapi"
)

type MyChaincode struct {
    contractapi.Contract
    Counter int
}

func (t *MyChaincode) Init(ctx contractapi.TransactionContextInterface) error {
    fmt.Println("MyChaincode Init")
    t.Counter = 0
    return nil
}

func (t *MyChaincode) Hello(ctx contractapi.TransactionContextInterface) error {
    fmt.Println("MyChaincode Hello")
    return nil
}

func (t *MyChaincode) IncrementCounter(ctx contractapi.TransactionContextInterface) error {
    t.Counter++ // Vulnerable code: direct access to public field
    return nil
}

func (t *MyChaincode) GetCounter(ctx contractapi.TransactionContextInterface) (int, error) {
    return t.Counter, nil
}

func main() {
    chaincode, err := contractapi.NewChaincode(&MyChaincode{})
    if err != nil {
        fmt.Printf("Error creating MyChaincode chaincode: %s", err.Error())
        return
    }

    if err := chaincode.Start(); err != nil {
        fmt.Printf("Error starting MyChaincode chaincode: %s", err.Error())
    }
}
