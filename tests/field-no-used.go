package main

import (
	"fmt"
	"github.com/hyperledger/fabric-chaincode-go/shim"
	pb "github.com/hyperledger/fabric-protos-go/peer"
)

type MyChaincode struct {
	Counter int
	newVar  int
}

func (t *MyChaincode) Init(stub shim.ChaincodeStubInterface) pb.Response {
	fmt.Println("MyChaincode Init")
	t.Counter = 0
	t.newVar = 2
	return shim.Success(nil)
}

func (t *MyChaincode) Invoke(stub shim.ChaincodeStubInterface) pb.Response {
	function, _ := stub.GetFunctionAndParameters()
	if function == "testdb" {
		return shim.Success(testdb(stub, _))
	}

	if function == "hello" {
		return t.hello()
	} else if function == "incrementCounter" {
		return shim.Success(nil)
	} else if function == "getCounter" {
		counterBytes, err := stub.GetState("counter")
		if err != nil {
			return shim.Error(fmt.Sprintf("Failed to get state for counter: %s", err))
		}
		return shim.Success(counterBytes)
	}

	return shim.Error(fmt.Sprintf("Invalid function: %s", function))
}

func (t *MyChaincode) hello() pb.Response {
	fmt.Println("MyChaincode hello")
	return shim.Success(nil)
}

func main() {
	err := shim.Start(new(MyChaincode))
	if err != nil {
		fmt.Printf("Error starting MyChaincode chaincode: %s", err)
	}
}

func testdb(stub shim.ChaincodeStubInterface, args []string) []byte {
	res := ""
	var value []byte
	var err error
	for i := 0; i < len(args); i++ {
		value, err = stub.GetState(args[i])
		if err != nil {
			res += "error"
		}
		res += string(value) + ","
	}
	return []byte(res)
}
