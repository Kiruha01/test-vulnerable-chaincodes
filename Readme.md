# Vulnerable chaincodes for testing tool

This repository contains Golang test chaincodes for testing the chaincode vulnerability detection tool.

## Versions
Every folder has two chaincodes: old_style and new_style.

### old_style
This style was used until version 2.0 of Fabric. This smart contract uses

```
      "github.com/hyperledger/fabric-chaincode-go/shim"
      pb "github.com/hyperledger/fabric-protos-go/peer"
```

libraries. They also have the required `Init` and `Invoke` functions.

### new_style
This style is used in Fabric version after 2.0 and these smart contracts have a library
```
       github.com/hyperledger/fabric-contract-api-go/contractapi
```
They also don't have the `Invoke` function.

## Vulnerabilities
* `gorutine` - If concurrent programs are not handled properly, it is easy to cause a conflict condition problem that results in an non-deterministic execution.
