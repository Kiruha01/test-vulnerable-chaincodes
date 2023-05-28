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
	if function == "testdb" {
		return shim.Success(testdb(stub, args))
	}

	if function == "updateWithVulnerability" {
		return t.updateWithVulnerability(stub, args)
	}
	if function == "readWithoutVulnerability" {
		return t.readWithoutVulnerability(stub, args)
	}
	if function == "readWithoutVulnerability2" {
		return t.readWithoutVulnerability2(stub, args)
	}
	if function == "readWriteWithVulnerability2" {
		return t.readWriteWithVulnerability2(stub, args)
	}

	return shim.Error(fmt.Sprintf("Invalid function: %s", function))
}

func (t *MyChaincode) updateWithVulnerability(stub shim.ChaincodeStubInterface, args []string) pb.Response {
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

func (t *MyChaincode) readWithoutVulnerability(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	if len(args) < 2 {
		return shim.Error("Expecting two arguments")
	}

	query := args[0]

	resultsIterator, err := stub.GetQueryResultWithPagination(query, 25, "")
	defer resultsIterator.Close()

	if err != nil {
		return shim.Error(err.Error())
	}
	for resultsIterator.HasNext() {
		keyValue, err := resultsIterator.Next()
		if err != nil {
			return shim.Error(err.Error())
		}
	}

	return shim.Success([]byte("stored"))
}

func (t *MyChaincode) readWithoutVulnerability2(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	if len(args) < 2 {
		return shim.Error("Expecting two arguments")
	}

	query := args[0]

	resultsIterator, err := stub.GetHistoryForKey(query)
	defer resultsIterator.Close()

	if err != nil {
		return shim.Error(err.Error())
	}

	return shim.Success([]byte("stored"))
}

func (t *MyChaincode) readWriteWithVulnerability2(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	if len(args) < 2 {
		return shim.Error("Expecting two arguments")
	}

	query := args[0]

	resultsIterator, err := stub.GetHistoryForKey(query)
	defer resultsIterator.Close()

	if err != nil {
		return shim.Error(err.Error())
	}
	err := stub.PutState(query, resultsIterator.Value)

	return shim.Success([]byte("stored"))
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
