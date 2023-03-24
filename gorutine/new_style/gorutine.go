// peer chaincode query -C mychannel -n gor2 -c '{"Args":["ExecuteParallel", "[\"a\",\"b\"]"]}'

package main

import (
    "fmt"
    "sync"

    "github.com/hyperledger/fabric-contract-api-go/contractapi"
)

type MyChaincode struct {
    contractapi.Contract
}

func (t *MyChaincode) ExecuteParallel(ctx contractapi.TransactionContextInterface, args []string) error {
    var wg sync.WaitGroup
    for _, arg := range args {
        wg.Add(1)
        go func(arg string) {
            defer wg.Done()
            // do some heavy computation with arg
            // ...
        }(arg)
    }
    wg.Wait()
    fmt.Println("All goroutines finished")
    return nil
}

func main() {
    cc, err := contractapi.NewChaincode(&MyChaincode{})
    if err != nil {
        fmt.Printf("Error creating chaincode: %s", err.Error())
        return
    }

    if err := cc.Start(); err != nil {
        fmt.Printf("Error starting chaincode: %s", err.Error())
    }
}
