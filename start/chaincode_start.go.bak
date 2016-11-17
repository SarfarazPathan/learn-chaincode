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
    Id               string `json:"id"`
    Name             string `json:"name"`
    Location         string `json:"location"`
    ProviderNM       string `json:"providerNM"`
    PrimayInsurance  string `json:"primaryInsurance"`
    Status           string `json:"status"`
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

	  provider1 := Provider{"MQ001", "PETE","INPATIENT HOSPITAL","JOHN ROBERT","TRICARE SOUTH","REJECTED"}
	  provider2 := Provider{"MQ002", "SRIKANTH","INPATIENT HOSPITAL","JOHN ROBERT","TRICARE SOUTH","PENDING"}
	  provider3 := Provider{"MQ003", "SARFARAZ","INPATIENT HOSPITAL","JOHN ROBERT","TRICARE SOUTH","APPROVED"}
	  provider4 := Provider{"MQ004", "TESTER","INPATIENT HOSPITAL","JOHN ROBERT","TRICARE SOUTH","APPROVED"}
	  provider5 := Provider{"MQ005", "PROVIDER","INPATIENT HOSPITAL","JOHN ROBERT","TRICARE SOUTH","APPROVED"}
	  provider6 := Provider{"MQ006", "PAYER","INPATIENT HOSPITAL","JOHN ROBERT","TRICARE SOUTH","PENDING"}
	  provider7 := Provider{"MQ007", "MEMBER","INPATIENT HOSPITAL","JOHN ROBERT","TRICARE SOUTH","REJECTED"}
	  provider8 := Provider{"MQ008", "JEANNE","INPATIENT HOSPITAL","JOHN ROBERT","TRICARE SOUTH","APPROVED"}
	  provider9 := Provider{"MQ009", "JING","INPATIENT HOSPITAL","JOHN ROBERT","TRICARE SOUTH","PENDING"}
	  provider10 := Provider{"MQ010", "PETER","INPATIENT HOSPITAL","JOHN ROBERT","TRICARE SOUTH","APPROVED"}
	 
	   theJson1, _ := json.Marshal(provider1)
	    theJson2, _ := json.Marshal(provider2)
	     theJson3, _ := json.Marshal(provider3)
	      theJson4, _ := json.Marshal(provider4)
	       theJson5, _ := json.Marshal(provider5)
	        theJson6, _ := json.Marshal(provider6)
		 theJson7, _ := json.Marshal(provider7)
		  theJson8, _ := json.Marshal(provider8)
		   theJson9, _ := json.Marshal(provider9)
		    theJson10, _ := json.Marshal(provider10)

	fmt.Printf("%+v\n", string(theJson1))
	fmt.Printf("%+v\n", string(theJson10))

	stub.PutState("MQ001", theJson1)
	stub.PutState("MQ002", theJson2)
	stub.PutState("MQ003", theJson3)
	stub.PutState("MQ004", theJson4)
	stub.PutState("MQ005", theJson5)
	stub.PutState("MQ006", theJson6)
	stub.PutState("MQ007", theJson7)
	stub.PutState("MQ008", theJson8)
	stub.PutState("MQ009", theJson9)
	stub.PutState("MQ010", theJson10)



	//  provider := Provider{"1001", "SarfarazPathan"}
	//  theJson, _ := json.Marshal(provider)

	// fmt.Printf("%+v\n", string(theJson))

	
	//stub.PutState("1001", theJson)
	
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
		//return nil, nil;
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
    var key, jsonResp string
    var err  error

    if len(args) != 1 {
        return nil, errors.New("Incorrect number of arguments. Expecting name of the key to query")
    }

   
    key = args[0]
    valAsbytes, err := stub.GetState(key)
    if err != nil {
        jsonResp = "{\"Error\":\"Failed to get state for " + key + "\"}"
        return nil, errors.New(jsonResp)
    }
	

	var p Provider	
	
    json.Unmarshal(valAsbytes, &p)

    fmt.Printf("%+v\n", p)

    return valAsbytes, nil
}

func (p *Provider) UnmarshalJSON(buf []byte) error {
	tmp := []interface{}{&p.Id, &p.Name}
	wantLen := len(tmp)
	if err := json.Unmarshal(buf, &tmp); err != nil {
		return err
	}
	if g, e := len(tmp), wantLen; g != e {
		return fmt.Errorf("wrong number of fields in Notification: %d != %d", g, e)
	}
	return nil
}