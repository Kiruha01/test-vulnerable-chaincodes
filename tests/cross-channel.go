package main

import (
	"fmt"
	"github.com/hyperledger/fabric-chaincode-go/shim"
	pb "github.com/hyperledger/fabric-protos-go/peer"
)

type MyChaincode struct{}

func (t *MyChaincode) Init(stub shim.ChaincodeStubInterface) pb.Response {
	return shim.Success(nil)
}

func (t *MyChaincode) Invoke(stub shim.ChaincodeStubInterface) pb.Response {
	function, args := stub.GetFunctionAndParameters()
	if function == "testdb" {
		return shim.Success(testdb(stub, args))
	}

	if function == "invokeOtherChaincode" {
		return t.invokeOtherChaincode(stub, args)
	}
	if function == "invokeOtherChaincodeInCurrentChannel" {
		return t.invokeOtherChaincodeInCurrentChannel(stub, args)
	}
	if function == "invokeOtherChaincodeInHardcodeChannel" {
		return t.invokeOtherChaincodeInHardcodeChannel(stub, args)
	}

	return shim.Error("Invalid function name")
}

func (t *MyChaincode) invokeOtherChaincode(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	if len(args) < 2 {
		return shim.Error("Expecting two arguments")
	}

	channelID := args[0]
	chaincodeName := args[1]

	argsAsBytes := [][]byte{}
	for _, arg := range args[2:] {
		argsAsBytes = append(argsAsBytes, []byte(arg))
	}

	response := stub.InvokeChaincode(chaincodeName, argsAsBytes, channelID)
	if response.Status != shim.OK {
		return shim.Error(fmt.Sprintf("Failed to invoke chaincode '%s' on channel '%s'", chaincodeName, channelID))
	}

	return shim.Success(response.Payload)
}

func (t *MyChaincode) invokeOtherChaincodeInCurrentChannel(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	if len(args) < 2 {
		return shim.Error("Expecting two arguments")
	}

	chaincodeName := args[0]

	argsAsBytes := [][]byte{}
	for _, arg := range args[2:] {
		argsAsBytes = append(argsAsBytes, []byte(arg))
	}

	response := stub.InvokeChaincode(chaincodeName, argsAsBytes, "")
	if response.Status != shim.OK {
		return shim.Error(fmt.Sprintf("Failed to invoke chaincode '%s' on channel '%s'", chaincodeName, channelID))
	}

	return shim.Success(response.Payload)
}

func (t *MyChaincode) invokeOtherChaincodeInHardcodeChannel(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	if len(args) < 2 {
		return shim.Error("Expecting two arguments")
	}

	chaincodeName := args[0]

	argsAsBytes := [][]byte{}
	for _, arg := range args[2:] {
		argsAsBytes = append(argsAsBytes, []byte(arg))
	}

	response := stub.InvokeChaincode(chaincodeName, argsAsBytes, "myChannel")
	if response.Status != shim.OK {
		return shim.Error(fmt.Sprintf("Failed to invoke chaincode '%s' on channel '%s'", chaincodeName, channelID))
	}

	return shim.Success(response.Payload)
}

func main() {
	err := shim.Start(new(MyChaincode))
	if err != nil {
		fmt.Printf("Error starting Simple chaincode: %s", err)
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
