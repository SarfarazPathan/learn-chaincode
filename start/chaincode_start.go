/*
Copyright IBM Corp 2016 All Rights Reserved.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

		 http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package main

import (
	"errors"
	"fmt"
	"encoding/json"

	"github.com/hyperledger/fabric/core/chaincode/shim"
)

// SimpleChaincode example simple Chaincode implementation
type SimpleChaincode struct {
}

type Provider struct {
	id        string  `json:"id"`
	name      string  `json:"name"`
	add       string  `json:"add"`
}

// ============================================================================================================================
// Main
// ============================================================================================================================
func main() {
	err := shim.Start(new(SimpleChaincode))
	if err != nil {
		fmt.Printf("Error starting Simple chaincode: %s", err)
	}
}

// Init resets all the things
//func (t *SimpleChaincode) Init(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {
func (t *SimpleChaincode) Init(stub *shim.ChaincodeStub, function string, args []string) ([]byte, error){
	if len(args) != 1 {
		return nil, errors.New("Incorrect number of arguments. Expecting 1")
	}
	
	var v Provider
	
	id := "\"id\":\"1001\", "
	name := "\"name\":\"Sarfaraz\", "
	add := "\"add\":\"Pune\", "

	v_json := "{"+id+name+add+"}" 

	json.Unmarshal([]byte(v_json), &v) 
	//if err != nil { return nil, errors.New("Invalid JSON object") }
	

	bytes, _ := json.Marshal(v)

	
	_ = stub.PutState(v.id, bytes)

	


	return nil, nil
}

// Invoke is our entry point to invoke a chaincode function
//func (t *SimpleChaincode) Invoke(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {
func (t *SimpleChaincode) Invoke(stub *shim.ChaincodeStub, function string, args []string) ([]byte, error) {
	fmt.Println("invoke is running " + function)

	// Handle different functions
	if function == "init" {													//initialize the chaincode state, used as reset
		return t.Init(stub, "init", args)
	}else if function == "write" {
		return t.write(stub, args)
	}	
	fmt.Println("invoke did not find func: " + function)					//error

	return nil, errors.New("Received unknown function invocation: " + function)
}

// Query is our entry point for queries
//func (t *SimpleChaincode) Query(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {
func (t *SimpleChaincode) Query(stub *shim.ChaincodeStub, function string, args []string) ([]byte, error) {
	fmt.Println("query is running " + function)
	
	// Handle different functions
	if function == "read" {											//read a variable
		fmt.Println("hi there " + function)						//error
		return t.read(stub, args)
	
	}
	fmt.Println("query did not find func: " + function)						//error

	return nil, errors.New("Received unknown function query: " + function)
}

func (t *SimpleChaincode) write(stub *shim.ChaincodeStub, args []string) ([]byte, error) {
    var key, value string
    var err error
    fmt.Println("running write()")

    if len(args) != 2 {
        return nil, errors.New("Incorrect number of arguments. Expecting 2. name of the key and value to set")
    }

    key = args[0]                            //rename for fun
    value = args[1]
    err = stub.PutState(key, []byte(value))  //write the variable into the chaincode state
    if err != nil {
        return nil, err
    }
    return nil, nil
}

func (t *SimpleChaincode) read(stub *shim.ChaincodeStub, args []string) ([]byte, error) {
fmt.Println("In read")

   var v Provider

	bytes, _ := stub.GetState("1001");

	//if err != nil {	fmt.Printf("RETRIEVE_V5C: Failed to invoke vehicle_code: %s", err); return v, errors.New("RETRIEVE_V5C: Error retrieving vehicle with v5cID = " + v5cID) }

	_ = json.Unmarshal(bytes, &v);

        //if err != nil {	fmt.Printf("RETRIEVE_V5C: Corrupt vehicle record "+string(bytes)+": %s", err); return v, errors.New("RETRIEVE_V5C: Corrupt vehicle record"+string(bytes))	}

	return bytes, nil
}