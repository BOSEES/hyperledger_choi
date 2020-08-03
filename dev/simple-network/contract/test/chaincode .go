package main

import (
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/hyperledger/fabric/core/chaincode/shim"
	sc "github.com/hyperledger/fabric/protos/peer"
)

type SmartContract struct {

}

type UserRating struct {
	User string `json:"user"`
	Average float64 `json:"average"`
	Rates []Rate `json: "rates"`
}

type Rate struct {
	ProjectTitle string `json:"projecttitle"`
	Score float64 `json: "score"`
}

func(s *SmartContract) Init(APIstup shim.ChaincodeStubInterface) sc.Response {
	return shim.Success(nil)
}

func(s *SmartContract) Invoke(APIstup shim.ChaincodeStubInterface) sc.Response {
	function, args := APIstup.GetFunctionAndParameters()

	if function == "addUser" {
		return s.addUser(APIstup, args);
	} else if function == "addRating" {
		return s.addRating(APIstup, args);
	} else if function == "readRating" {
		retrun s.readRating(APIstup, args);
	}
	return shim.Error("invalid Smart Contract function name.")
}

func (s *SmartContract) addUser(APIstup shim.ChaincodeStubInterface, args []string) sc.Response {
	if len(args) !=
}