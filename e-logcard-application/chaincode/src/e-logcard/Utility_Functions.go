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

// ========================================
// Listes finies & Récupération 
// ========================================

func (t *SimpleChaincode) getRolesList (stub shim.ChaincodeStubInterface, args []string)([]byte, error) {
	roles := []string { "Auditor_authority", "AH_admin", "supplier", "manufacturer", "customer", "maintenance_user" }
return json.Marshal(roles) 
}

func (t *SimpleChaincode) getActionsList (stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {
	Modifications := []string { "SB", "Monte", "Demonte"}
return json.Marshal(Modifications)
}

func (t *SimpleChaincode) getAircraftTypesList(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {
	Type := []string { "Defense" }
return json.Marshal(Type) 
}

func (t *SimpleChaincode) getAircraftsList(stub shim.ChaincodeStubInterface, args []string) ([]byte, error){
	partMap,err:=getAircraftMap(stub)
	if(err !=nil){return nil, nil}
	parts := make([]string, len(partMap))
    idx := 0
    for  _, part := range partMap {
    		parts[idx] = part.Id_Aircraft
    		idx++    
    }
    //si les deux longueurs sont differentes on slice
    if(len(partMap)!=idx){
    	parts=parts[0:idx]
    }
return json.Marshal(parts)
}

func (t *SimpleChaincode) getAssembliesList(stub shim.ChaincodeStubInterface, args []string) ([]byte, error){
	partMap,err:=getAssemblyMap(stub)
	if(err !=nil){return nil, nil}
	parts := make([]string, len(partMap))
    idx := 0
    for  _, part := range partMap {
    		parts[idx] = part.Id_Assembly
    		idx++    
    }
    //si les deux longueurs sont differentes on slice
    if(len(partMap)!=idx){
    	parts=parts[0:idx]
    }
return json.Marshal(parts)
}

func (t *SimpleChaincode) getPartsList (stub shim.ChaincodeStubInterface, args []string)([]byte, error){
	
	partMap,err:=getPartsIdMap(stub)
		if(err !=nil){return nil, nil}
	parts := make([]string, len(partMap))
    idx := 0
    for  _, part := range partMap {
    		parts[idx] = part.Id
    		idx++
    }
    //si les deux longueurs sont differentes on slice
    if(len(partMap)!=idx){
    	parts=parts[0:idx]
    }
return json.Marshal(parts)
}


func (t *SimpleChaincode) getList(stub shim.ChaincodeStubInterface, args []string) ([]byte, error)  {
typ := args [0]

if (typ == "List_Roles") {
roles := []string { "Auditor_authority", "AH_admin", "supplier", "manufacturer", "customer", "maintenance_user" }
	return json.Marshal(roles)
} else if ( typ == "Modification_Types") {
Modifications := []string { "SB", "Monte", "Demonte"}
return json.Marshal(Modifications)
} else if ( typ == "Asset_Types") {
Type := []string { "Defense" }
return json.Marshal(Type) 

} else if (typ == "List_Aircrafts"){
	partMap,err:=getAircraftMap(stub)
	if(err !=nil){return nil, nil}
	parts := make([]string, len(partMap))
    idx := 0
    for  _, part := range partMap {
    		parts[idx] = part.Id_Aircraft
    		idx++    
    }
    //si les deux longueurs sont differentes on slice
    if(len(partMap)!=idx){
    	parts=parts[0:idx]
    }
return json.Marshal(parts)
} else if (typ == "List_Assemblies") {
	partMap,err:=getAssemblyMap(stub)
	if(err !=nil){return nil, nil}
	parts := make([]string, len(partMap))
    idx := 0
    for  _, part := range partMap {
    		parts[idx] = part.Id_Assembly
    		idx++    
    }
    //si les deux longueurs sont differentes on slice
    if(len(partMap)!=idx){
    	parts=parts[0:idx]
    }
return json.Marshal(parts)
} else if (typ == "List_Parts") {
	partMap,err:=getPartsIdMap(stub)
		if(err !=nil){return nil, nil}
	parts := make([]string, len(partMap))
    idx := 0
    for  _, part := range partMap {
    		parts[idx] = part.Id
    		idx++
    }
    //si les deux longueurs sont differentes on slice
    if(len(partMap)!=idx){
    	parts=parts[0:idx]
    }
return json.Marshal(parts)
}
	return nil,nil
}
