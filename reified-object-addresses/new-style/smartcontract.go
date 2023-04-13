package main

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/hyperledger/fabric-contract-api-go/contractapi"
)

type Asset struct {
	ID   string `json:"ID"`
	Name string `json:"Name"`
}

type MyChaincode struct {
	contractapi.Contract
}

func (t *MyChaincode) CreateAsset(ctx contractapi.TransactionContextInterface, id string, name string) error {
	asset1 := &Asset{ID: id, Name: name}
	assetJson, err := json.Marshal(asset1)
	if err != nil {
		return fmt.Errorf("failed to marshal asset: %s", err.Error())
	}

	err = ctx.GetStub().PutState(name, assetJson)
	if err != nil {
		return fmt.Errorf("failed to put state: %s", err.Error())
	}

	return nil
}

func (t *MyChaincode) CreateAsset2(ctx contractapi.TransactionContextInterface, id string, name string) error {
	asset1 := new(Asset{ID: id, Name: name})
	assetJson, err := json.Marshal(asset1)
	if err != nil {
		return fmt.Errorf("failed to marshal asset: %s", err.Error())
	}

	err = ctx.GetStub().PutState(name, assetJson)
	if err != nil {
		return fmt.Errorf("failed to put state: %s", err.Error())
	}

	return nil
}

func (t *MyChaincode) CreateAsset3(ctx contractapi.TransactionContextInterface, id string, name string) error {
	var asset1 *Asset = new(Asset{id, name})
	assetJson, err := json.Marshal(asset1)
	if err != nil {
		return fmt.Errorf("failed to marshal asset: %s", err.Error())
	}

	err = ctx.GetStub().PutState(name, assetJson)
	if err != nil {
		return fmt.Errorf("failed to put state: %s", err.Error())
	}

	return nil
}

func main() {
	cc, err := contractapi.NewChaincode(new(MyChaincode))
	if err != nil {
		log.Panicf("Error creating chaincode: %v", err)
	}

	if err := cc.Start(); err != nil {
		log.Panicf("Error starting chaincode: %v", err)
	}
}
