package main

import (
	"fmt"
	"github.com/hyperledger/fabric-chaincode-go/shim"
	pb "github.com/hyperledger/fabric-protos-go/peer"
)

type MyChaincode struct {
}

func (t *MyChaincode) Init(stub shim.ChaincodeStubInterface) pb.Response {
	fmt.Println("MyChaincode Init")
	return shim.Success(nil)
}

func (t *MyChaincode) Invoke(stub shim.ChaincodeStubInterface) pb.Response {
	function, args := stub.GetFunctionAndParameters()

	if function == "update" {
		return t.update(stub, args)
	}

	return shim.Error(fmt.Sprintf("Invalid function: %s", function))
}

func (t *MyChaincode) update(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	if len(args) < 2 {
		return shim.Error("Expecting two arguments")
	}
	
	startKey := args[0]
	endKey := args[1]
	
	resultsIterator, err := stub.GetStateByRange(startKey, endKey)
	defer resultsIterator.Close()
	
	if err != nil {
		return shim.Error(err.Error())
	}
	for resultsIterator.HasNext() {
		keyValue, err := resultsIterator.Next()
		if err != nil {
			return shim.Error(err.Error())
		}
		err := stub.PutState(keyValue.Key, keyValue.Value)
	}

	return shim.Success([]byte("stored"))
}

func main() {
	err := shim.Start(new(MyChaincode))
	if err != nil {
		fmt.Printf("Error starting MyChaincode chaincode: %s", err)
	}
}
