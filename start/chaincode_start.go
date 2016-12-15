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


//==============================================================================================================================
//	User_and_eCert - Struct for storing the JSON of a user and their ecert
//==============================================================================================================================

type User_and_eCert struct {
	Identity string `json:"identity"`
	eCert string `json:"ecert"`
}


type PreAuthRequest struct {
    /*Id               string `json:"id"`
    Dor              string `json:"dor"`
    Name             string `json:"name"`
    Location         string `json:"location"`
    ProviderNM       string `json:"providerNM"`
    PrimayInsurance  string `json:"primaryInsurance"`
    Status           string `json:"status"`
    User             string `json:"user"`
   */
    PreAuthReqId	     string `json:"preAuthReqId"`

    ProviderId	     string `json:"providerId"`
    ProviderName     string `json:"providerName"`
    ProviderAddr     string `json:"providerAddr"`
    ProviderCity     string `json:"providerCity"`
    ProviderZip      string `json:"providerZip"`
    ProviderPhone    string `json:"providerPhone"`
    ProviderFax      string `json:"providerFax"`
    ProviderContactPerson      string `json:"providerContactPerson"`
    
    PatientName      string `json:"patientName"`
    PatientID        string `json:"patientID"`
    PatientDOB        string `json:"patientDOB"`
    PatientDOR        string `json:"patientDOR"`

    SrvRequested        string `json:"srvRequested"`
    SrvDOS		string `json:"srvDOS"`
    SrvCPTCode		string `json:"srvCPTCode"`
    SrvICDCode		string `json:"srvICDCode"`
    SrvProviderFacility		string `json:"srvProviderFacility"`
    SrvPhone		string `json:"srvPhone"`
    SrvAddr		string `json:"srvAddr"`
    SrvCity		string `json:"srvCity"`
    SrvZip		string `json:"srvZip"`

    ClinicalInfo	string `json:"clinicalInfo"`

    PrimayInsurance  string `json:"primaryInsurance"`
    Status           string `json:"status"`
    User             string `json:"user"`

}
//type PreAuthRequests []*PreAuthRequest



type ProviderInformation struct {
	ProviderID	string `json:"providerId"`
	Name 		string `json:"providerName"`	
	Address 	string `json:"address"`
	City	 	string `json:"city"`
	ZipCode 	string `json:"zipCode"`
	Phone 		string `json:"phone"`
	Fax 		string `json:"fax"`
	ContactPerson	string `json:"contactPerson"`
}
type ProvidersInformation []*ProviderInformation

type PatientInformation struct {
	Name 		string `json:"memberName"`
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
	Phone		 string `json:"serPhone"`
	Address		 string `json:"serAddress"`
	City		 string `json:"serCity"`
	ZipCode		 string `json:"serZipCode"`	
}
type ServicesRequested []*ServiceRequested

type ClientInformation struct {
	ClinicalInformationDoc 		string `json:"clinicalInformationDoc"`
	Attachements 	string `json:"attachements"`
}
type ClientsInformation []*ClientInformation

var temp []byte

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

	preAuthRequests := []PreAuthRequest{}
	//preAuthRequest := PreAuthRequest{"MQ001","01/01/2017", "SARFARAZ","INPATIENT HOSPITAL","JOHN ROBERT","TRICARE 	SOUTH","REJECTED"}
	//preAuthRequests = append(preAuthRequests,preAuthRequest)
        preAuthReqJson, _ := json.Marshal(preAuthRequests)
	fmt.Println(string(preAuthReqJson))
	stub.PutState("lst", preAuthReqJson)

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

	
	//for i:=0; i < len(randomVal); i=i+2 {
	//	t.add_ecert(stub, randomVal[i], randomVal[i+1])
	//}


	return nil, nil

}


//==============================================================================================================================
//	 General Functions
//==============================================================================================================================
//	 get_ecert - Takes the name passed and calls out to the REST API for HyperLedger to retrieve the ecert
//				 for that user. Returns the ecert as retrived including html encoding.
//==============================================================================================================================
func (t *SimpleChaincode) get_ecert(stub shim.ChaincodeStubInterface, name string) ([]byte, error) {

	ecert, err := stub.GetState(name)

	if err != nil { return nil, errors.New("Couldn't retrieve ecert for user " + name) }

	return ecert, nil
}

//==============================================================================================================================
//	 add_ecert - Adds a new ecert and user pair to the table of ecerts
//==============================================================================================================================

func (t *SimpleChaincode) add_ecert(stub shim.ChaincodeStubInterface, name string, ecert string) ([]byte, error) {


	err := stub.PutState(name, []byte(ecert))

	if err == nil {
		return nil, errors.New("Error storing eCert for user " + name + " identity: " + ecert)
	}

	return nil, nil

}

//==============================================================================================================================
//	 get_caller - Retrieves the username of the user who invoked the chaincode.
//				  Returns the username as a string.
//==============================================================================================================================

func (t *SimpleChaincode) get_username(stub shim.ChaincodeStubInterface) (string, error) {

    username, err := stub.ReadCertAttribute("username");
	if err != nil { return "", errors.New("Couldn't get attribute 'username'. Error: " + err.Error()) }
	return string(username), nil
}






// Invoke is our entry point to invoke a chaincode function
func (t *SimpleChaincode) Invoke(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {
	fmt.Println("invoke is running " + function)

	// Handle different functions
	if function == "init" {													//initialize the chaincode state, used as reset
		return t.Init(stub, "init", args)
	}else if function == "write" {
		return t.write(stub, args)
	}else if function == "writePreAuth" {
		return t.writePreAuth(stub, args)
	}else if function == "statusChange" {
		return t.statusChange(stub, args)
	}

	
	fmt.Println("invoke did not find func: " + function)					//error

	return nil, errors.New("Received unknown function invocation: " + function)
}


func (t *SimpleChaincode) write(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {
	  // var key, value string
	//var key , key1 , key2 , key3 , key4 , key5, key6 , key7 string
	//var key8 , key9 , key10 , key11 , key12 , key13, key14 , key15 string
	//var key16 , key17 , key18 , key19 , key20 , key21, key22 , key23 string

    var err error
    fmt.Println("running write()")

    if len(args) != 6 {
        return nil, errors.New("Incorrect number of arguments. Expecting 2. name of the key and value to set")
    }

	//key = args[0]                            //rename for fun
	
	//key1 = args[1]
	//key2 = args[2]
	//key3 = args[3]
	//key4 = args[4]
	//key5 = args[5]
	//key6 = args[6]

	//fmt.Println(key+" : "+key1+" : "+key2+" : "+key3+" : "+key4+" : "+key5+" : "+key6)

	//provider := PreAuthRequest{key, key1,key2,key3,key4,key5,key6}
	
    /*
	  key = args[0]                            //rename for fun
	  key1 = args[12]
	  key2 = args[10]
	  key3 = args[3] + " , " +args[4]
	  key4 = args[2]
	  key5 = "Blue Cross"
	  key6 = "PENDING"
	  key7 = args[23]

	  provider := PreAuthRequest{key, key1,key2,key3,key4,key5,key6,key7} 

	  key = args[0]                        
	  key1 = args[12]
	  key2 = args[10]
	  key3 = args[3] + " , " +args[4]
	  key4 = args[2]
	  key5 = "Blue Cross"
	  key6 = "PENDING"
	  key7 = args[23]

	  provider := PreAuthRequest{key, key1,key2,key3,key4,key5,key6,key7}

	theJson, err := json.Marshal(provider)
	fmt.Printf("%+v\n", string(theJson))
	stub.PutState(key, theJson)
     */
    
    //err = stub.PutState(key, []byte(value))  //write the variable into the chaincode state
    if err != nil {
       return nil, err
    }

    return nil, nil
}

//////////////////////////////

func (t *SimpleChaincode) writePreAuth(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {
   // var key , key1 , key2 , key3 , key4 , key5, key6 , key7 string
   // var key8 , key9 , key10 , key11 , key12 , key13, key14 , key15 string
   // var key16 , key17 , key18 , key19 , key20 , key21, key22 , key23 string

      var preAuthReqId , 	providerId , providerName ,providerAddr , providerCity , providerZip , providerPhone ,providerFax string	
      var providerContactPerson , patientName , 	patientID , patientDOB , patientDOR , srvRequested , srvDOS , srvCPTCode string
      var srvICDCode , srvProviderFacility , srvPhone , srvAddr , srvCity , srvZip , clinicalInfo  string
      var primaryInsurance , status , user string 

    var err error
    fmt.Println("running writePreAuth >>>>> ")

   // if len(args) != 6 {
     //   return nil, errors.New("Incorrect number of arguments. Expecting 2. name of the key and value to set")
   // }

	 valAsbytes , _ := t.read(stub,args)

	 fmt.Println("writePreAuth valAsbytes >>>>> "+string(valAsbytes))

	 for i:=0; i < len(args); i++ {
		fmt.Println(args[i])
		fmt.Println(i)
	}

	// var oldPreAuthReq *[]PreAuthRequest
	// json.Unmarshal(valAsbytes, &oldPreAuthReq)

	 oldPreAuthReq := []PreAuthRequest{}
	 json.Unmarshal(valAsbytes, &oldPreAuthReq)

	fmt.Println("<<<<<<<< PreAuthRequest >>>>>")

	 /* key = args[0]                            //rename for fun
	  key4 = args[2]
	  key3 = args[3] + " , " +args[4]
	  key2 = args[9]
	  key1 = args[12]
	  key7 = args[23]

	  key5 = "Blue Cross"
	  key6 = "PENDING"

	 fmt.Println(key+" : "+key1+" : "+key2+" : "+key3+" : "+key4+" : "+key5+" : "+key6+" : "+key7) */

	        preAuthReqId = 	args[0]				//// 0
		providerId = 	args[1]				//// 1
		providerName =  args[2]				//// 2
		providerAddr = 	args[3]				//// 3 
		providerCity = 	args[4]				//// 4 	
		providerZip = 	args[5]				//// 5 	
		providerPhone = args[6]				//// 6 
		   
		providerFax =  	args[7]				//// 7
		providerContactPerson = args[8]			//// 8	
		patientName =  	args[9]				//// 9
		patientID =  	args[10]				//// 10	
		patientDOB =  	args[11]				//// 11
		patientDOR =  	args[12]				//// 12
		srvRequested = 	args[13]				//// 13
		   
		srvDOS =  	args[14]				//// 14
		srvCPTCode =  	args[15]				//// 15
		srvICDCode =  	args[16]				//// 16	
		srvProviderFacility =  args[17]			//// 17	
		srvPhone =  	args[18]				//// 18
		srvAddr =  	args[19]				//// 19
		srvCity = 	args[20]				//// 20
		   
		srvZip =  	args[21]				//// 21
		clinicalInfo = 	args[22]				//// 22
		user = 	args[23]					//// 23

		primaryInsurance = args[24]	
		//primaryInsurance = "Blue Cross"
	        status = "PENDING"

	// fmt.Println(key+" : "+key1+" : "+key2+" : "+key3+" : "+key4+" : "+key5+" : "+key6+" : "+key7)

         //newPreAuthReq := PreAuthRequest{key, key1,key2,key3,key4,key5,key6,key7}
	 newPreAuthReq := PreAuthRequest{preAuthReqId , providerId , providerName ,providerAddr , providerCity , providerZip , providerPhone ,providerFax , providerContactPerson , patientName , patientID , patientDOB , patientDOR , srvRequested , srvDOS , srvCPTCode , srvICDCode , srvProviderFacility , srvPhone , srvAddr , srvCity , srvZip , clinicalInfo ,primaryInsurance  ,status  , user}
	 oldPreAuthReq = append(oldPreAuthReq,newPreAuthReq)
	
	// oldPreAuthReq := PreAuthRequests{&newPreAuthReq}
 	//*oldPreAuthReq = append(*oldPreAuthReq,newPreAuthReq)
	//fmt.Println(*oldPreAuthReq)

	theJson, err := json.Marshal(oldPreAuthReq)
	fmt.Printf("oldPreAuthReq %+v\n", string(theJson))
	stub.PutState("lst", theJson)

    
    //err = stub.PutState(key, []byte(value))  //write the variable into the chaincode state
    if err != nil {
       return nil, err
    }

    return nil, nil
}







///////////

func (t *SimpleChaincode) writeProviderInformation(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {
	  // var key, value string
   // var key , key1 , key2 , key3 , key4 , key5  , key6   , key7 string
    var err error
    fmt.Println("running write()")

    if len(args) != 6 {
        return nil, errors.New("Incorrect number of arguments. Expecting 2. name of the key and value to set")
    }
/*
	key = args[0]                            //rename for fun
	//value = args[1]
	
	key1 = args[1]
	key2 = args[2]
	key3 = args[3]
	key4 = args[4]
	key5 = args[5]
	key6 = args[6]
        key7 = args[7]
	
	fmt.Println(key+" : "+key1+" : "+key2+" : "+key3+" : "+key4+" : "+key5+" : "+key6)
	
	provider := PreAuthRequest{key, key1,key2,key3,key4,key5,key6,key7}
	theJson, err := json.Marshal(provider)
	fmt.Printf("%+v\n", string(theJson))
	stub.PutState(key, theJson)
*/
    
    //err = stub.PutState(key, []byte(value))  //write the variable into the chaincode state
    if err != nil {
       return nil, err
    }

    return nil, nil
}

func (t *SimpleChaincode) writePatientsInformation(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {
	  // var key, value string
  //  var key , key1 , key2 , key3 , key4 , key5  , key6  , key7 string
    var err error
    fmt.Println("running write()")

    if len(args) != 6 {
        return nil, errors.New("Incorrect number of arguments. Expecting 2. name of the key and value to set")
    }
/*
	key = args[0]                            //rename for fun
	//value = args[1]
	
	key1 = args[1]
	key2 = args[2]
	key3 = args[3]
	key4 = args[4]
	key5 = args[5]
	key6 = args[6]
	key7 = args[7]
	
	fmt.Println(key+" : "+key1+" : "+key2+" : "+key3+" : "+key4+" : "+key5+" : "+key6)
 
	provider := PreAuthRequest{key, key1,key2,key3,key4,key5,key6,key7}
	theJson, err := json.Marshal(provider)
	fmt.Printf("%+v\n", string(theJson))
	stub.PutState(key, theJson)
*/
    
    //err = stub.PutState(key, []byte(value))  //write the variable into the chaincode state
    if err != nil {
       return nil, err
    }

    return nil, nil
}


func (t *SimpleChaincode) writeServicesRequested(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {
	  // var key, value string
  //  var key , key1 , key2 , key3 , key4 , key5  , key6   , key7 string
    var err error
    fmt.Println("running write()")

    if len(args) != 6 {
        return nil, errors.New("Incorrect number of arguments. Expecting 2. name of the key and value to set")
    }
/*
	key = args[0]                            //rename for fun
	//value = args[1]
	
	key1 = args[1]
	key2 = args[2]
	key3 = args[3]
	key4 = args[4]
	key5 = args[5]
	key6 = args[6]
	key7 = args[7]

	fmt.Println(key+" : "+key1+" : "+key2+" : "+key3+" : "+key4+" : "+key5+" : "+key6)

	provider := PreAuthRequest{key, key1,key2,key3,key4,key5,key6,key7}
	theJson, err := json.Marshal(provider)
	fmt.Printf("%+v\n", string(theJson))
	stub.PutState(key, theJson)
*/
    
    //err = stub.PutState(key, []byte(value))  //write the variable into the chaincode state
    if err != nil {
       return nil, err
    }

    return nil, nil
}

func (t *SimpleChaincode) writeClientsInformation(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {
	  // var key, value string
 //   var key , key1 , key2 , key3 , key4 , key5  , key6   , key7 string
    var err error
    fmt.Println("running write()")
/*
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
	key6 = args[6]
	key7 = args[7]

	fmt.Println(key+" : "+key1+" : "+key2+" : "+key3+" : "+key4+" : "+key5+" : "+key6)

	provider := PreAuthRequest{key, key1,key2,key3,key4,key5,key6,key7}
	theJson, err := json.Marshal(provider)
	fmt.Printf("%+v\n", string(theJson))
	stub.PutState(key, theJson)
*/
    
    //err = stub.PutState(key, []byte(value))  //write the variable into the chaincode state
    if err != nil {
       return nil, err
    }

    return nil, nil
}





func (t *SimpleChaincode) read(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {
 fmt.Println(" In read ...")
    var key, jsonResp string
    var err  error
	
	user, _ := t.get_username(stub)
	fmt.Println("USER : >> "+user)

  
    fmt.Println(" In read 1")
    valAsbytes, err := stub.GetState("lst")
    fmt.Println(" In read 2")
    if err != nil {
        jsonResp = "{\"Error\":\"Failed to get state for " + key + "\"}"
        return nil, errors.New(jsonResp)
    }
    fmt.Println(" In read 3")
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

	if function == "readPreAuth" {											//read a variable
		fmt.Println("hi there " + function)						//error
		return t.readPreAuth(stub, args)
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


		/////////// Read PreAuth  
func (t *SimpleChaincode) readPreAuth(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {
    fmt.Println(" In readPreAuth ")
    var key, jsonResp  string
    var err  error

    if len(args) != 1 {
        return nil, errors.New("Incorrect number of arguments. Expecting name of the key to query")
    }

    fmt.Println(" In readPreAuth 1")
    valAsbytes, err := stub.GetState("lst")
    if err != nil {
        jsonResp = "{\"Error\":\"Failed to get state for " + key + "\"}"
        return nil, errors.New(jsonResp)
    }
   
    fmt.Println(" In readPreAuth 2")
    preAuthReq := []PreAuthRequest{}
    json.Unmarshal(valAsbytes, &preAuthReq)
    fmt.Println("PreAuth >>>>> "+string(valAsbytes))

    key = args[0]
    var preAuthJsonOut []byte
    fmt.Println(" In readPreAuth 3")

     for _, preAuth := range preAuthReq {
        fmt.Println(preAuth.PreAuthReqId +" : "+ preAuth.ProviderName)
        fmt.Println("")
	if key == preAuth.PreAuthReqId {
		preAuthJson, _ := json.Marshal(preAuth)	
		preAuthJsonOut = preAuthJson
		fmt.Println("preAuthJson >>>>> "+string(preAuthJsonOut))
		//temp = append(temp,preAuthJson...)
	}else {
		fmt.Println("No Match Found...")
	}
    }
	fmt.Println(string(preAuthJsonOut))
	return preAuthJsonOut, nil	
      //return preAuthJson, nil	
    //return valAsbytes, nil
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


		/////////// Status Changes   
func (t *SimpleChaincode) statusChange(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {
    var preAuthReqId, jsonResp, approve  string
    var err  error

     valAsbytes, err := stub.GetState("lst")
    if err != nil {
        jsonResp = "{\"Error\":\"Failed to get state for " + preAuthReqId + "\"}"
        return nil, errors.New(jsonResp)
    }
   
    preAuthReqId = args[0]
    approve = args[1]

    fmt.Println(" In statusChange 2"+preAuthReqId)
    newPreAuthReq := []PreAuthRequest{}
    oldPreAuthReq  := []PreAuthRequest{}

    json.Unmarshal(valAsbytes, &oldPreAuthReq )
    
    fmt.Println("statusChange >>>>> "+string(valAsbytes))

	for _, preAuth := range oldPreAuthReq {
 	    fmt.Println(preAuth.PreAuthReqId +" : "+ preAuth.ProviderName)
		if preAuthReqId == preAuth.PreAuthReqId {	
			preAuth.Status = approve
		}
	  newPreAuthReq = append(newPreAuthReq,preAuth)	
	}
	
	newPreAuthReqJson, _ := json.Marshal(newPreAuthReq)
	fmt.Println(string(newPreAuthReqJson))
	
	stub.PutState("lst", newPreAuthReqJson)

    return nil, nil
}