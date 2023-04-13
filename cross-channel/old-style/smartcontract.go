// peer chaincode invoke -o localhost:7050 --ordererTLSHostnameOverride orderer.example.com --tls --cafile "${PWD}/organizations/ordererOrganizations/example.com/orderers/orderer.example.com/msp/tlscacerts/tlsca.example.com-cert.pem" -C ch2 -n cross --peerAddresses localhost:7051 --tlsRootCertFiles "${PWD}/organizations/peerOrganizations/org1.example.com/peers/peer0.org1.example.com/tls/ca.crt" --peerAddresses localhost:9051 --tlsRootCertFiles "${PWD}/organizations/peerOrganizations/org2.example.com/peers/peer0.org2.example.com/tls/ca.crt" -c '{"function":"invokeOtherChaincode","Args":["ch1", "gor", "executeParallel", "a"]}'

package main

import (
	"fmt"
	"github.com/hyperledger/fabric-chaincode-go/shim"
	pb "github.com/hyperledger/fabric-protos-go/peer"
)

type MyChaincode struct {}

func (t *MyChaincode) Init(stub shim.ChaincodeStubInterface) pb.Response {
	return shim.Success(nil)
}

func (t *MyChaincode) Invoke(stub shim.ChaincodeStubInterface) pb.Response {
	function, args := stub.GetFunctionAndParameters()

	if function == "invokeOtherChaincode" {
		return t.invokeOtherChaincode(stub, args)
	}

	return shim.Error("Invalid function name")
}


func (t *MyChaincode) invokeOtherChaincode(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	if len(args) < 2 {
		return shim.Error("Expecting two arguments")
	}

	channelID := args[0]
	chaincodeName := args[1]
	
	// Convert args to a byte array
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

func main() {
	err := shim.Start(new(MyChaincode))
	if err != nil {
		fmt.Printf("Error starting Simple chaincode: %s", err)
	}
}