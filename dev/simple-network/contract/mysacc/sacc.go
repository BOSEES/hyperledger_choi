/*
 * Copyright IBM Corp All Rights Reserved
 *
 * SPDX-License-Identifier: Apache-2.0
 */

package main

import (
	"encoding/json"
	"fmt"

	"github.com/hyperledger/fabric/core/chaincode/shim"
	"github.com/hyperledger/fabric/protos/peer"
)

// SimpleAsset은 자산을 관리하기 위한 간단한 체인 코드를 구현함
type SimpleAsset struct {

}

type Data struct {
	Key string `json:"key"`
	Value string `json:"value"`
}

// 체인코드 인스턴스화 중에 초기화하기 위해 Init를 호출함
// 데이터. 체인 코드 업그레이드에서는 이 기능도 재설정이라고 함에 유의하십시오.
// 또는 데이터를 마이그레이션한다.
func (t *SimpleAsset) Init(stub shim.ChaincodeStubInterface) peer.Response {
	return shim.Success(nil)
}
// 체인코드에서 트랜잭션별로 호출된다. 각각의 거래는
// Init 함수에 의해 생성된 자산의 'get' 또는 'set' 중 하나. 세트
// 새로운 키-값 쌍을 지정하여 새 자산을 생성할 수 있다.
func (t *SimpleAsset) Invoke(stub shim.ChaincodeStubInterface) peer.Response {
	// 거래 제안서에서 함수 및 아그 추출
	fn, args := stub.GetFunctionAndParameters()
	
	var result string
	var err error
	if fn == "set" {
		result, err = set(stub, args);
	} else if fn == "get" {
		result, err = get(stub, args);
	} else if fn == "getAllKeys" {
		result, err = getAllKeys(stub);
	} else {
		return shim.Error("not supported Chaingcode function.")
	}
	
	if err != nil {
		return shim.Error(err.Error())
	}
	return shim.Success([]byte(result))
}

// 대장에 있는 자산(키와 값 모두)을 저장한다. 키가 존재한다면
// 값을 새 값으로 재정의함


func set(stub shim.ChaincodeStubInterface, args []string) (string, error) {
	if len(args) != 2 {
		return "", fmt.Errorf("Incorrect arguments. Expecting a key")
	}
	//json 반환
	var data = Data{Key: args[0], Value: args[1]}
	dataAsBytes, _ := json.Marshal(data)
	
	err :=stub.PutState(args[0], dataAsBytes)
	if err != nil {
		return "", fmt.Errorf("Failed to set asset: %s", args[0])
	}
	return string(dataAsBytes), nil
}

func get(stub shim.ChaincodeStubInterface, args []string) (string, error) {
	if len(args) != 1 {
		return "", fmt.Errorf("incorrect arguments. Expecting a key")
	}

	value, err := stub.GetState(args[0])
	if err != nil {
		return "", fmt.Errorf("Failed to get asset: %s with error: %s", args[0], err)
	}
	if value == nil {
		return "", fmt.Errorf("Asset not found: %s", args[0])
	}
	return string(value), nil
}

// 지정된 자산 키의 값 가져오기
func getAllKeys(stub shim.ChaincodeStubInterface) (string, error) {
	iter, err := stub.GetStateByRange("a","z")
	if err != nil {
		return "", fmt.Errorf("Faile to get all keys with error: %s", err)
	}
	defer iter.Close()

	var buffer string
	buffer = "["

	comma := false
	for iter.HasNext() {
		res, err := iter.Next()
		if err != nil {
			return "", fmt.Errorf("%s",err)
		}
		if comma == true {
			buffer += ","
		}
		buffer += string(res.Value)
		comma = true
	}
	buffer += "]"

	fmt.Println(buffer)

	return string(buffer), nil
}

// 인스턴스화하는 동안 컨테이너의 체인 코드를 시작하는 주 기능

func main() {
	if err := shim.Start(new(SimpleAsset)); err != nil {
		fmt.Printf("Error starting SimpleAsset chaincode: %s", err)
	}
}
