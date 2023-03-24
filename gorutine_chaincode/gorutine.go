// peer chaincode query -C mychannel -n gor -c '{"Args":["executeParallel", "[\"a\"]"]}'

package main

import (
    "fmt"
    "sync"
    "github.com/hyperledger/fabric-chaincode-go/shim"
    pb "github.com/hyperledger/fabric-protos-go/peer"
)

type SimpleChaincode struct {
}

func main() {
    err := shim.Start(new(SimpleChaincode))
    if err != nil {
        fmt.Printf("Error starting Simple chaincode: %s", err)
    }
}

func (t *SimpleChaincode) Init(stub shim.ChaincodeStubInterface) pb.Response {
    fmt.Println("MyChaincode Init")
    return shim.Success(nil)
}

func (t *SimpleChaincode) Invoke(stub shim.ChaincodeStubInterface) pb.Response {
    function, args := stub.GetFunctionAndParameters()
    fmt.Println("invoke is running " + function)

    if function == "executeParallel" {
        go t.executeParallel(stub, args)
        return shim.Success(nil)
    }

    fmt.Println("invoke did not find func: " + function)
    return shim.Error("Received unknown function invocation")
}


func (t *SimpleChaincode) executeParallel(stub shim.ChaincodeStubInterface, args []string) {
    var wg sync.WaitGroup
    for _, arg := range args {
        fmt.Println(arg)
        wg.Add(1)
        go func(arg string) {
            defer wg.Done()
            // do some heavy computation with arg
            // ...
        }(arg)
    }
    wg.Wait()
    fmt.Println("All goroutines finished")
}
