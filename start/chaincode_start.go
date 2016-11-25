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
	"encoding/json"
	"errors"
	"fmt"

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

type Providers []*Provider



type ProviderInformation struct {
	ProviderID	string `json:"providerId"`
	Name 		string `json:"name"`	
	Address 	string `json:"address"`
	City	 	string `json:"city"`
	ZipCode 	string `json:"zipCode"`
	Phone 		string `json:"phone"`
	Fax 		string `json:"fax"`
	ContactPerson	string `json:"contactPerson"`
}
type ProvidersInformation []*ProviderInformation

type PatientInformation struct {
	Name 		string `json:"name"`
	MemberID 	string `json:"memberID"`
	DOB 		string `json:"dob"`
	DateOfRequest 	string `json:"dateOfRequest"`
}
type PatientsInformation []*PatientInformation

type ServiceRequested struct {
	ServiceRequested string `json:"serviceRequested"`	
	DOS		 string `json:"dos"`
	CPTCode		 string `json:"cptCode"`
	ICD10Code	 string `json:"icd10Code"`
	ProviderFacility string `json:"providerFacility"`
	Phone		 string `json:"phone"`
	Address		 string `json:"address"`
	City		 string `json:"city"`
	ZipCode		 string `json:"zipCode"`	
}
type ServicesRequested []*ServiceRequested

type ClientInformation struct {
	ClinicalInformationDoc 		string `json:"clinicalInformationDoc"`
	Attachements 	string `json:"attachements"`
}
type ClientsInformation []*ClientInformation

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
func (t *SimpleChaincode) Init(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {
	if len(args) != 1 {
		return nil, errors.New("Incorrect number of arguments. Expecting 1")
	}

	  provider1 := Provider{"MQ001", "SARFARAZ","INPATIENT HOSPITAL","JOHN ROBERT","TRICARE SOUTH","REJECTED"}
	  provider2 := Provider{"MQ002", "JOHN","INPATIENT HOSPITAL","JOHN ROBERT","TRICARE SOUTH","PENDING"}
	  provider3 := Provider{"MQ003", "PETE","INPATIENT HOSPITAL","JOHN ROBERT","TRICARE SOUTH","APPROVED"}
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

	    provBytes := Providers{&provider1, &provider2,&provider3,&provider4,&provider5,&provider6,&provider7,&provider8}

	    b, _ := json.Marshal(provBytes)
	    fmt.Println(string(b))
	
	stub.PutState("lst", b)


		///////////// Default ProvidersInformation Stored
	provInfo := ProviderInformation{"PI001", "SarfarazPathan","NIBM ROAD","PUNE","441802","01010101010","FAX10101","Khan"}
	provIBytes := ProvidersInformation{&provInfo}
	provInfoJson, _ := json.Marshal(provIBytes)
	fmt.Println(string(provInfoJson))
	stub.PutState("provInfoJson1", provInfoJson)

		///////////// Default PatientInformation  Stored
	patInfo := PatientInformation{"MI001", "SarfarazPathan","02-20-2020","02-20-2121"}
	patBytes := PatientsInformation{&patInfo}
	patInfoJson, _ := json.Marshal(patBytes)
	fmt.Println(string(patInfoJson))
	stub.PutState("patInfoJson1", patInfoJson)

		///////////// Default PatientInformation  Stored
	serReq := ServiceRequested {"EyeTransplant", "DOS","CPT001","ICD10001","YES","01010101010","OZON 2 Hadpsar","PUNE","441803"}
	serReqBytes := ServicesRequested{&serReq}
	serReqJson, _ := json.Marshal(serReqBytes)
	fmt.Println(string(serReqJson))
	stub.PutState("serReqJson1", serReqJson)
	
		///////////// Default ClientInformation   Stored
	cliInfo := ClientInformation {"Blood Report,ECG", "YES"}
	cliInfoBytes := ClientsInformation{&cliInfo}
	cliInfoJson, _ := json.Marshal(cliInfoBytes)
	fmt.Println(string(cliInfoJson))
	stub.PutState("cliInfoJson1", cliInfoJson)

	return nil, nil

}


// Invoke is our entry point to invoke a chaincode function
func (t *SimpleChaincode) Invoke(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {
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


func (t *SimpleChaincode) write(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {
	  // var key, value string
    var key , key1 , key2 , key3 , key4 , key5 string
    var err error
    fmt.Println("running write()")

    if len(args) != 6 {
        return nil, errors.New("Incorrect number of arguments. Expecting 2. name of the key and value to set")
    }

	key = args[0]                            //rename for fun
	//value = args[1]
	
	key1 = args[1]
	key2 = args[2]
	key3 = args[3]
	key4 = args[4]
	key5 = args[5]

	fmt.Println(key+" : "+key1+" : "+key2+" : "+key3+" : "+key4+" : "+key5)

	provider := Provider{key, key1,key2,key3,key4,key5}
	theJson, err := json.Marshal(provider)
	fmt.Printf("%+v\n", string(theJson))
	stub.PutState(key, theJson)

    
    //err = stub.PutState(key, []byte(value))  //write the variable into the chaincode state
    if err != nil {
       return nil, err
    }

    return nil, nil
}


///////////

func (t *SimpleChaincode) writeProviderInformation(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {
	  // var key, value string
    var key , key1 , key2 , key3 , key4 , key5 string
    var err error
    fmt.Println("running write()")

    if len(args) != 6 {
        return nil, errors.New("Incorrect number of arguments. Expecting 2. name of the key and value to set")
    }

	key = args[0]                            //rename for fun
	//value = args[1]
	
	key1 = args[1]
	key2 = args[2]
	key3 = args[3]
	key4 = args[4]
	key5 = args[5]

	fmt.Println(key+" : "+key1+" : "+key2+" : "+key3+" : "+key4+" : "+key5)

	provider := Provider{key, key1,key2,key3,key4,key5}
	theJson, err := json.Marshal(provider)
	fmt.Printf("%+v\n", string(theJson))
	stub.PutState(key, theJson)

    
    //err = stub.PutState(key, []byte(value))  //write the variable into the chaincode state
    if err != nil {
       return nil, err
    }

    return nil, nil
}

func (t *SimpleChaincode) writePatientsInformation(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {
	  // var key, value string
    var key , key1 , key2 , key3 , key4 , key5 string
    var err error
    fmt.Println("running write()")

    if len(args) != 6 {
        return nil, errors.New("Incorrect number of arguments. Expecting 2. name of the key and value to set")
    }

	key = args[0]                            //rename for fun
	//value = args[1]
	
	key1 = args[1]
	key2 = args[2]
	key3 = args[3]
	key4 = args[4]
	key5 = args[5]

	fmt.Println(key+" : "+key1+" : "+key2+" : "+key3+" : "+key4+" : "+key5)

	provider := Provider{key, key1,key2,key3,key4,key5}
	theJson, err := json.Marshal(provider)
	fmt.Printf("%+v\n", string(theJson))
	stub.PutState(key, theJson)

    
    //err = stub.PutState(key, []byte(value))  //write the variable into the chaincode state
    if err != nil {
       return nil, err
    }

    return nil, nil
}


func (t *SimpleChaincode) writeServicesRequested(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {
	  // var key, value string
    var key , key1 , key2 , key3 , key4 , key5 string
    var err error
    fmt.Println("running write()")

    if len(args) != 6 {
        return nil, errors.New("Incorrect number of arguments. Expecting 2. name of the key and value to set")
    }

	key = args[0]                            //rename for fun
	//value = args[1]
	
	key1 = args[1]
	key2 = args[2]
	key3 = args[3]
	key4 = args[4]
	key5 = args[5]

	fmt.Println(key+" : "+key1+" : "+key2+" : "+key3+" : "+key4+" : "+key5)

	provider := Provider{key, key1,key2,key3,key4,key5}
	theJson, err := json.Marshal(provider)
	fmt.Printf("%+v\n", string(theJson))
	stub.PutState(key, theJson)

    
    //err = stub.PutState(key, []byte(value))  //write the variable into the chaincode state
    if err != nil {
       return nil, err
    }

    return nil, nil
}

func (t *SimpleChaincode) writeClientsInformation(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {
	  // var key, value string
    var key , key1 , key2 , key3 , key4 , key5 string
    var err error
    fmt.Println("running write()")

    if len(args) != 6 {
        return nil, errors.New("Incorrect number of arguments. Expecting 2. name of the key and value to set")
    }

	key = args[0]                            //rename for fun
	//value = args[1]
	
	key1 = args[1]
	key2 = args[2]
	key3 = args[3]
	key4 = args[4]
	key5 = args[5]

	fmt.Println(key+" : "+key1+" : "+key2+" : "+key3+" : "+key4+" : "+key5)

	provider := Provider{key, key1,key2,key3,key4,key5}
	theJson, err := json.Marshal(provider)
	fmt.Printf("%+v\n", string(theJson))
	stub.PutState(key, theJson)

    
    //err = stub.PutState(key, []byte(value))  //write the variable into the chaincode state
    if err != nil {
       return nil, err
    }

    return nil, nil
}





func (t *SimpleChaincode) read(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {
	  var key, jsonResp string
    var err  error

    if len(args) != 1 {
        return nil, errors.New("Incorrect number of arguments. Expecting name of the key to query")
    }

   
    //key = args[0]
   // valAsbytes, err := stub.GetState(key)
    //if err != nil {
     //   jsonResp = "{\"Error\":\"Failed to get state for " + key + "\"}"
      //  return nil, errors.New(jsonResp)
    //}
//    var p Provider	
//    json.Unmarshal(valAsbytes, &p)
//    fmt.Printf("%+v\n", p)
	
 //   valAsbytes, _ :=  t.get_providers(stub)
	
//	fmt.Println("hi valAsbytes " + string(valAsbytes))

    valAsbytes, err := stub.GetState("lst")
    if err != nil {
        jsonResp = "{\"Error\":\"Failed to get state for " + key + "\"}"
        return nil, errors.New(jsonResp)
    }
   
var p Providers
    json.Unmarshal(valAsbytes, &p)
   fmt.Println("hi >>>>> "+string(valAsbytes))

    return valAsbytes, nil

}

// Query is our entry point for queries
func (t *SimpleChaincode) Query(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {
	fmt.Println("query is running " + function)

	// Handle different functions
	if function == "read" {											//read a variable
		fmt.Println("hi there " + function)						//error
		return t.read(stub, args)
	}

	if function == "pro" {											//read a variable
		fmt.Println("hi there " + function)						//error
		return t.readProviderInformation(stub, args)
	}

	if function == "mem" {											//read a variable
		fmt.Println("hi there " + function)						//error
		return t.readPatientInformation(stub, args)
	}

	if function == "ser" {											//read a variable
		fmt.Println("hi there " + function)						//error
		return t.readServiceRequested(stub, args)
	}
	
	if function == "cli" {											//read a variable
		fmt.Println("hi there " + function)						//error
		return t.readClientInformation(stub, args)
	}

	fmt.Println("query did not find func: " + function)						//error

	return nil, errors.New("Received unknown function query: " + function)

}


		/////////// Read ProviderInformation 
func (t *SimpleChaincode) readProviderInformation(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {
    var key, jsonResp string
    var err  error

    if len(args) != 1 {
        return nil, errors.New("Incorrect number of arguments. Expecting name of the key to query")
    }

    valAsbytes, err := stub.GetState("provInfoJson1")
    if err != nil {
        jsonResp = "{\"Error\":\"Failed to get state for " + key + "\"}"
        return nil, errors.New(jsonResp)
    }
   
    var p ProvidersInformation
    json.Unmarshal(valAsbytes, &p)
    fmt.Println("ProvidersInformation >>>>> "+string(valAsbytes))

    return valAsbytes, nil
}

		/////////// Read PatientInformation  
func (t *SimpleChaincode) readPatientInformation(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {
    var key, jsonResp string
    var err  error

    if len(args) != 1 {
        return nil, errors.New("Incorrect number of arguments. Expecting name of the key to query")
    }

    valAsbytes, err := stub.GetState("patInfoJson1")
    if err != nil {
        jsonResp = "{\"Error\":\"Failed to get state for " + key + "\"}"
        return nil, errors.New(jsonResp)
    }
   
    var p PatientsInformation
    json.Unmarshal(valAsbytes, &p)
    fmt.Println("PatientInformation >>>>> "+string(valAsbytes))

    return valAsbytes, nil
}

		/////////// Read ServiceRequested  
func (t *SimpleChaincode) readServiceRequested(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {
    var key, jsonResp string
    var err  error

    if len(args) != 1 {
        return nil, errors.New("Incorrect number of arguments. Expecting name of the key to query")
    }

    valAsbytes, err := stub.GetState("serReqJson1")
    if err != nil {
        jsonResp = "{\"Error\":\"Failed to get state for " + key + "\"}"
        return nil, errors.New(jsonResp)
    }
   
    var s ServicesRequested
    json.Unmarshal(valAsbytes, &s)
    fmt.Println("ServiceRequested >>>>> "+string(valAsbytes))

    return valAsbytes, nil
}


		/////////// Read ClientInformation   
func (t *SimpleChaincode) readClientInformation(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {
    var key, jsonResp string
    var err  error

    if len(args) != 1 {
        return nil, errors.New("Incorrect number of arguments. Expecting name of the key to query")
    }

    valAsbytes, err := stub.GetState("cliInfoJson1")
    if err != nil {
        jsonResp = "{\"Error\":\"Failed to get state for " + key + "\"}"
        return nil, errors.New(jsonResp)
    }
   
    var c ClientsInformation
    json.Unmarshal(valAsbytes, &c)
    fmt.Println("ClientInformation  >>>>> "+string(valAsbytes))

    return valAsbytes, nil
}