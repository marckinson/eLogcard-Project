
// =========================================================================================
// 					UTILITY FUNCTIONS
// =========================================================================================

package main
import (
	"errors"
	"fmt"
	"encoding/json"
	"github.com/hyperledger/fabric/core/chaincode/shim"
)
//=====================
// Get Attributes 
//=====================
func getAttribute(stub shim.ChaincodeStubInterface, attributeName string) (string, error) {
	bytes, err := stub.ReadCertAttribute(attributeName)
	return string(bytes[:]), err
}
//=====================
// Create Part Map 
//=====================
func createMap(stub shim.ChaincodeStubInterface, args string)(error) {
	var key string
	key = args 
	// map creation configuration
		var partMap=make(map[string]Part)
		json2AsBytes, err := json.Marshal(partMap)
		err=stub.PutState(key,json2AsBytes)
		if err != nil {return err}
		//END map creation configuration
	return nil
}
//=====================
// Get Part Map (Récupération de la map des parts par id)
//=====================
func getPartsIdMap(stub shim.ChaincodeStubInterface)(map[string]Part, error) {
	allPartMapAsByte, err := stub.GetState("allParts")
		if(err !=nil){return nil,err}
	var partMap map[string]Part 
	err = json.Unmarshal(allPartMapAsByte, &partMap)
	if(err !=nil){return nil,err}
	return partMap,err	
}
func findPartById(stub shim.ChaincodeStubInterface,id string)(Part, error){
	partMap,err:=getPartsIdMap(stub)
	var part Part
	if(err !=nil){return part,err}
	part=partMap[id];
	return part,nil
}
//=====================
// Get Part Map (Récupération de la map des parts par PN)
//=====================
func getPartsPnMap(stub shim.ChaincodeStubInterface)(map[string]Part, error) {
	allPartMapAsByte, err := stub.GetState("allPartsPn")
		if(err !=nil){return nil,err}
	var partMap map[string]Part 
	err = json.Unmarshal(allPartMapAsByte, &partMap)
	if(err !=nil){return nil,err}
	return partMap,err	
}
func findPartByPn(stub shim.ChaincodeStubInterface,pn string)(Part, error){
	partMap,err:=getPartsPnMap(stub)
	var part Part
	if(err !=nil){return part,err}
	part=partMap[pn];
	return part,nil
}
//=====================
// Get Part Map (Récupération de la map des parts par SN)
//=====================
func getPartsSnMap(stub shim.ChaincodeStubInterface)(map[string]Part, error) {
	allPartMapAsByte, err := stub.GetState("allPartsSn")
	if(err !=nil){return nil,err}
	var partMap map[string]Part 
	err = json.Unmarshal(allPartMapAsByte, &partMap)
	if(err !=nil){return nil,err}
	return partMap,err	
}
func findPartBySn(stub shim.ChaincodeStubInterface,sn string)(Part, error){
	partMap,err:=getPartsSnMap(stub)
	var part Part
	if(err !=nil){return part,err}
	part=partMap[sn];
	return part,nil
}
//=====================
// Get Aircraft Map (Récupération de la map des Aicraft par Id)
//=====================
func getAircraftMap(stub shim.ChaincodeStubInterface)(map[string]Aircraft, error) {
	allPartMapAsByte, err := stub.GetState("allAircraft")
	if(err !=nil){return nil,err}
	var partMap map[string]Aircraft 
	err = json.Unmarshal(allPartMapAsByte, &partMap)
	if(err !=nil){return nil,err}
	return partMap,err	
}
func findAircraftById(stub shim.ChaincodeStubInterface,id string)(Aircraft, error){
	partMap,err:=getAircraftMap(stub)
	var aircraft Aircraft
	if(err !=nil){return aircraft,err}
	aircraft=partMap[id];
	return aircraft,nil
}
//=====================
// Get Assembly Map (Récupération de la map des Assembly par SN)
//=====================
func getAssemblyMap(stub shim.ChaincodeStubInterface)(map[string]Assembly, error) {
	allPartMapAsByte, err := stub.GetState("allAssembly")
	if(err !=nil){return nil,err}
	var partMap map[string]Assembly 
	err = json.Unmarshal(allPartMapAsByte, &partMap)
	if(err !=nil){return nil,err}
	return partMap,err	
}
func findAssemblyById(stub shim.ChaincodeStubInterface,id string)(Assembly, error){
	partMap,err:=getAssemblyMap(stub)
	var assembly Assembly
	if(err !=nil){return assembly,err}
	assembly=partMap[id];
	return assembly,nil
}

// =========================
// Check PN Availability 
// =========================
func checkPNavailibility(stub shim.ChaincodeStubInterface, args string) error {
	fmt.Println("Running checkSNavailibility")
	var err error
	var jsonResp2 string
	part,err:=findPartByPn(stub,args)
	if(err !=nil){return err}
	ptAS, _ := json.Marshal(part)
	var pt Part
	err = json.Unmarshal(ptAS, &pt)
	if err != nil {return  errors.New("Failed to Unmarshal Part #" + args)}
	if ( args == pt.PN) { 
		jsonResp2 = "{\"Error\":\"The following PN is Already taken, " + args + "\"}"
		return  errors.New(jsonResp2)
	} else if (args != pt.PN) {return nil}
	return nil 
}
// ===========================
// Check SN Availability 
// ===========================
func checkSNavailibility(stub shim.ChaincodeStubInterface, args string) error {
	fmt.Println("Running checkSNavailibility")
	var err error
	var jsonResp2 string
	part,err:=findPartBySn(stub,args)
	if(err !=nil){return err}
	ptAS, _ := json.Marshal(part)
	var pt Part
	err = json.Unmarshal(ptAS, &pt)
	if err != nil {return  errors.New("Failed to Unmarshal Part #" + args)}
	if ( args == pt.SN) { 
		jsonResp2 = "{\"Error\":\"The following PN is Already taken, " + args + "\"}"
		return  errors.New(jsonResp2)
	} else if ( args != pt.SN ) { return nil }
	return nil 
}
// ===========================
// Check Id Availability For the aircraft
// ===========================
func checkIdavailibility(stub shim.ChaincodeStubInterface, args string) error {
	fmt.Println("Running checkIdavailibility")
	var err error
	var jsonResp2 string
	part,err:=findAircraftById(stub,args)
	if(err !=nil){return err}
	ptAS, _ := json.Marshal(part)
	var pt Aircraft
	err = json.Unmarshal(ptAS, &pt)
	if err != nil {return  errors.New("Failed to Unmarshal Part #" + args)}
	if ( args == pt.Id_Aircraft) { 
		jsonResp2 = "{\"Error\":\"The following PN is Already taken, " + args + "\"}"
		return  errors.New(jsonResp2)
	} else if ( args != pt.Id_Aircraft ) { return nil }
	return nil 
}
// ===========================
// Check Id Availability for the assembly
// ===========================
func checkIdAssavailibility(stub shim.ChaincodeStubInterface, args string) error {
	fmt.Println("Running checkIdavailibility")
	var err error
	var jsonResp2 string
	part,err:=findAircraftById(stub,args)
	if(err !=nil){return err}
	ptAS, _ := json.Marshal(part)
	var pt Assembly
	err = json.Unmarshal(ptAS, &pt)
	if err != nil {return  errors.New("Failed to Unmarshal Part #" + args)}
	if ( args == pt.Id_Assembly) { 
		jsonResp2 = "{\"Error\":\"The following PN is Already taken, " + args + "\"}"
		return  errors.New(jsonResp2)
	} else if ( args != pt.Id_Assembly ) { return nil }
	return nil 
}
// ==================================================================================
// Check Ownership On Part 
// ==================================================================================
func checkOwnership(stub shim.ChaincodeStubInterface, args string) error {
	fmt.Println("Running CheckOwnership")
	
	var err error
	var key string // Id de la part 
	key = args 
	var jsonResp2 string
	username, err := getAttribute(stub, "username")
	
	part,err:=findPartById(stub,key)
	if err != nil {return  errors.New("Failed to get part #" + key)}
	ptAS, _ := json.Marshal(part)
	var pt Part
	err = json.Unmarshal(ptAS, &pt)
	if err != nil {return errors.New("Failed to Unmarshal Part #" + key)}
	if ( username != pt.Owner) { 
		jsonResp2 = "{\"Error\":\"You are not owner of this part, " + key + "\"}"
		return  errors.New(jsonResp2)
	} else if (username == pt.Owner) {return nil}
	return nil 
}
// ==================================================================================
// Check Responsibiity On Part 
// ==================================================================================
func checkResponsibility(stub shim.ChaincodeStubInterface, args string) error {
	fmt.Println("Running checkResponsibility")
	
	var err error
	var key string // Id de la part 
	key = args 
	var jsonResp2 string
	username, err := getAttribute(stub, "username")
	part,err:=findPartById(stub,key)
	if err != nil {return  errors.New("Failed to get part #" + key)}
	ptAS, _ := json.Marshal(part)
	var pt Part
	err = json.Unmarshal(ptAS, &pt)
	if err != nil {return errors.New("Failed to Unmarshal Part #" + key)}
	if ( username != pt.Responsible) { 
		jsonResp2 = "{\"Error\":\"You are not Responsible of this part, " + key + "\"}"
		return  errors.New(jsonResp2)
	} else if (username == pt.Responsible) {return nil}
	return nil 
}
// =========================================================================
// Creation of the Aircraft  
// =========================================================================
func createAircraft(stub shim.ChaincodeStubInterface, args string, args1 string, args2 string) error {
	fmt.Println("Running createAircraft")

	var err error
	e:= checkIdavailibility(stub, args)
	if e != nil { 
		var key string 
		key = args
		aircraft,err:=findAircraftById(stub,key)
			if err != nil {return  errors.New("Failed to get aircraft #" + key)}
			ptAS, _ := json.Marshal(aircraft)
		var airc Aircraft
			err = json.Unmarshal(ptAS, &airc)
			if err != nil {return  errors.New("Failed to Unmarshal Part #" + key)}
			airc.Id_Parts = append(airc.Id_Parts, args2)
		//Update allAircraft 
			partzMap,err:=getAircraftMap(stub)
			partzMap[airc.Id_Aircraft] = airc
			allPAsBuytes, err := json.Marshal(partzMap)
			err=stub.PutState("allAircraft",allPAsBuytes)
			if err != nil {return  err}
		//Fin update allParts 
	} else {	
		var air Aircraft 
			air.Id_Aircraft = args
			air.Owner = args1
			air.Id_Parts = append(air.Id_Parts, args2)
			
		//Commit aircfaft to ledger
			ptAsBytees, _ := json.Marshal(air)
				err = stub.PutState(air.Id_Aircraft, ptAsBytees)
				if err != nil {return  err}
		//Update allAircraft 
			partzMap,err:=getAircraftMap(stub)
			partzMap[air.Id_Aircraft] = air
			allPAsBuytes, err := json.Marshal(partzMap)
				err=stub.PutState("allAircraft",allPAsBuytes)
				if err != nil {return  err}
		//Fin update allAircraft 
	}
	fmt.Println("Responsible created successfully")	
return nil
}

// =========================================================================
// Creation of the Assembly  
// =========================================================================
func createAssembly(stub shim.ChaincodeStubInterface, args string, args1 string, args2 string) error {
	fmt.Println("Running createAssembly")

	var err error
	e:= checkIdAssavailibility(stub, args)
	if e != nil { 
		var key string 
		key = args
		aircraft,err:=findAssemblyById(stub,key)
			if err != nil {return  errors.New("Failed to get aircraft #" + key)}
			ptAS, _ := json.Marshal(aircraft)
		var airc Assembly
			err = json.Unmarshal(ptAS, &airc)
			if err != nil {return  errors.New("Failed to Unmarshal Part #" + key)}
			airc.Id_Parts = append(airc.Id_Parts, args2)
	
		//Update allAssembly 
		partzMap,err:=getAssemblyMap(stub)
		partzMap[airc.Id_Assembly] = airc
		allPAsBuytes, err := json.Marshal(partzMap)
			err=stub.PutState("allAssembly",allPAsBuytes)
			if err != nil {return  err}
		//Fin update allAssembly 
	} else {	
		var air Assembly 
			air.Id_Assembly = args
			air.Owner = args1
			air.Id_Parts = append(air.Id_Parts, args2)
		//Commit aircfaft to ledger
			ptAsBytees, _ := json.Marshal(air)
				err = stub.PutState(air.Id_Assembly, ptAsBytees)
				if err != nil {return  err}
		//Update allAssembly 
			partzMap,err:=getAssemblyMap(stub)
			partzMap[air.Id_Assembly] = air
			allPAsBuytes, err := json.Marshal(partzMap)
				err=stub.PutState("allAssembly",allPAsBuytes)
				if err != nil {return  err}
		//Fin update allAssembly 
	}
	fmt.Println("Responsible created successfully")	
return nil
}


