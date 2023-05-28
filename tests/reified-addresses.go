package main

import (
	"encoding/json"
	"fmt"
	"github.com/hyperledger/fabric-chaincode-go/shim"
	pb "github.com/hyperledger/fabric-protos-go/peer"
)

type Asset struct {
	ID   string `json:"ID"`
	Name string `json:"Name"`
}

type MyChaincode struct{}

func (t *MyChaincode) Init(stub shim.ChaincodeStubInterface) pb.Response {
	return shim.Success(nil)
}

func (t *MyChaincode) Invoke(stub shim.ChaincodeStubInterface) pb.Response {
	function, args := stub.GetFunctionAndParameters()
	if function == "testdb" {
		return shim.Success(testdb(stub, args))
	}
	if function == "testdb" {
		return shim.Success(testdb(stub, args))
	}

	if function == "createAsset" {
		return t.createAsset(stub, args)
	}
	if function == "createAsset2" {
		return t.createAsset2(stub, args)
	}
	if function == "createAsset3" {
		return t.createAsset3(stub, args)
	}

	return shim.Error("Invalid function name")
}

func (t *MyChaincode) createAsset(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	id := args[0]
	name := args[1]

	asset1 := &Asset{id, name}
	assetJson, err := json.Marshal(asset1)
	if err != nil {
		return shim.Error(err.Error())
	}

	err = stub.PutState(name, assetJson)
	if err != nil {
		return shim.Error(err.Error())
	}

	return shim.Success(nil)
}

func (t *MyChaincode) createAsset2(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	id := args[0]
	name := args[1]

	asset1 := new(Asset{id, name})
	assetJson, err := json.Marshal(asset1)
	if err != nil {
		return shim.Error(err.Error())
	}

	err = stub.PutState(name, assetJson)
	if err != nil {
		return shim.Error(err.Error())
	}

	return shim.Success(nil)
}

func (t *MyChaincode) createAsset3(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	id := args[0]
	name := args[1]

	var asset1 *Asset = new(Asset{id, name})
	assetJson, err := json.Marshal(asset1)
	if err != nil {
		return shim.Error(err.Error())
	}

	err = stub.PutState(name, assetJson)
	if err != nil {
		return shim.Error(err.Error())
	}

	return shim.Success(nil)
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
