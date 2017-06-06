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
// =========================================================================================
// 					Parts
// =========================================================================================
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
// Get Part Map (Récupération de la map des parts par ID)
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
		jsonResp2 = "{\"Error\":\"The following SN is Already taken, " + args + "\"}"
		return  errors.New(jsonResp2)
	} else if ( args != pt.SN ) { return nil }
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




// =========================================================================================
// 					Aircraft
// =========================================================================================
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
// Get Part Map (Récupération de la map des parts par SN)
//=====================
func getAircraftSnMap(stub shim.ChaincodeStubInterface)(map[string]Aircraft, error) {
	allPartMapAsByte, err := stub.GetState("allAircraftsSn")
	if(err !=nil){return nil,err}
	var partMap map[string]Aircraft 
	err = json.Unmarshal(allPartMapAsByte, &partMap)
	if(err !=nil){return nil,err}
	return partMap,err	
}
func findAircraftBySn(stub shim.ChaincodeStubInterface,sn string)(Aircraft, error){
	partMap,err:=getAircraftSnMap(stub)
	var part Aircraft
	if(err !=nil){return part,err}
	part=partMap[sn];
	return part,nil
}
//=====================
// Get Part Map (Récupération de la map des parts par AN)
//=====================
func getAircraftAnMap(stub shim.ChaincodeStubInterface)(map[string]Aircraft, error) {
	allPartMapAsByte, err := stub.GetState("allAircraftsAn")
	if(err !=nil){return nil,err}
	var partMap map[string]Aircraft 
	err = json.Unmarshal(allPartMapAsByte, &partMap)
	if(err !=nil){return nil,err}
	return partMap,err	
}
func findAircraftByAn(stub shim.ChaincodeStubInterface,an string)(Aircraft, error){
	partMap,err:=getAircraftAnMap(stub)
	var part Aircraft
	if(err !=nil){return part,err}
	part=partMap[an];
	return part,nil
}
// ===========================
// Check AN Availability For the aircraft
// ===========================
func checkAnAircraft (stub shim.ChaincodeStubInterface, args string) error {
	
	var err error
	var jsonResp2 string
	part,err:=findAircraftByAn(stub,args)
	if(err !=nil){return err}
	ptAS, _ := json.Marshal(part)
	var pt Aircraft
	err = json.Unmarshal(ptAS, &pt)
	if err != nil {return  errors.New("Failed to Unmarshal Part #" + args)}
	if ( args == pt.AN) { 
		jsonResp2 = "{\"Error\":\"The following PN is Already taken, " + args + "\"}"
		return  errors.New(jsonResp2)
	} else if (args != pt.AN) {return nil}
	
	return nil 
}
// ===========================
// Check SN Availability For the aircraft
// ===========================
func checkSnAircraft (stub shim.ChaincodeStubInterface, args string) error {

	var err error
	var jsonResp2 string
	part,err:=findAircraftBySn(stub,args)
	if(err !=nil){return err}
	ptAS, _ := json.Marshal(part)
	var pt Aircraft
	err = json.Unmarshal(ptAS, &pt)
	if err != nil {return  errors.New("Failed to Unmarshal Part #" + args)}
	if ( args == pt.SN) { 
		jsonResp2 = "{\"Error\":\"The following PN is Already taken, " + args + "\"}"
		return  errors.New(jsonResp2)
	} else if (args != pt.SN) {return nil}
	
	return nil 
}
// =========================================================================================
// 					Assembly
// =========================================================================================
//=====================
// Get Assembly Map (Récupération de la map des Assembly par ID)
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
//=====================
// Get Assembly Map (Récupération de la map des assemblies par SN)
//=====================
func getAssembliesSnMap(stub shim.ChaincodeStubInterface)(map[string]Assembly, error) {
	allPartMapAsByte, err := stub.GetState("allAssembliesSn")
	if(err !=nil){return nil,err}
	var partMap map[string]Assembly 
	err = json.Unmarshal(allPartMapAsByte, &partMap)
	if(err !=nil){return nil,err}
	return partMap,err	
}
func findAssembyBySn(stub shim.ChaincodeStubInterface,sn string)(Assembly, error){
	partMap,err:=getAssembliesSnMap(stub)
	var part Assembly
	if(err !=nil){return part,err}
	part=partMap[sn];
	return part,nil
}
//=====================
// Get Assembly Map (Récupération de la map des assemblies par AN)
//=====================
func getAssembliesAnMap(stub shim.ChaincodeStubInterface)(map[string]Assembly, error) {
	allPartMapAsByte, err := stub.GetState("allAssembliesAn")
	if(err !=nil){return nil,err}
	var partMap map[string]Assembly 
	err = json.Unmarshal(allPartMapAsByte, &partMap)
	if(err !=nil){return nil,err}
	return partMap,err	
}
func findAssembyByAn(stub shim.ChaincodeStubInterface,an string)(Assembly, error){
	partMap,err:=getAssembliesAnMap(stub)
	var part Assembly
	if(err !=nil){return part,err}
	part=partMap[an];
	return part,nil
}
// ===========================
// Check AN Availability for the assembly
// ===========================
func checkAnAssembly(stub shim.ChaincodeStubInterface, args string) error {

	var err error
	var jsonResp2 string
	part,err:=findAssembyByAn(stub,args)
	if(err !=nil){return err}
	ptAS, _ := json.Marshal(part)
	var pt Assembly
	err = json.Unmarshal(ptAS, &pt)
	if err != nil {return  errors.New("Failed to Unmarshal Part #" + args)}
	if ( args == pt.AN) { 
		jsonResp2 = "{\"Error\":\"The following PN is Already taken, " + args + "\"}"
		return  errors.New(jsonResp2)
	} else if (args != pt.AN) {return nil}
	
	return nil 
}
// ===========================
// Check SN Availability for the assembly
// ===========================
func checkSnAssembly(stub shim.ChaincodeStubInterface, args string) error {
	var err error
	var jsonResp2 string
	part,err:=findAssembyBySn(stub,args)
	if(err !=nil){return err}
	ptAS, _ := json.Marshal(part)
	var pt Assembly
	err = json.Unmarshal(ptAS, &pt)
	if err != nil {return  errors.New("Failed to Unmarshal Part #" + args)}
	if ( args == pt.SN) { 
		jsonResp2 = "{\"Error\":\"The following PN is Already taken, " + args + "\"}"
		return  errors.New(jsonResp2)
	} else if (args != pt.AN) {return nil}

	return nil 
}

// =========================================================================================
// 					Fonctions génériques
// =========================================================================================

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

func (t *SimpleChaincode) getLogsList(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {
	Logs := []string { "CREATE", "OWNERSHIP_TRANSFER", "RESPONSIBILITY_TRANSFER", "ACTIVITIES_PERFORMED", "ADDED TO A/C: ", "ADDED TO ASSEMBLY: ", "REMOVED FROM A/C: ", "REMOVED FROM ASSEMBLY: ", "PART_REMOVAL", "PART_AFFILIATION", "REMOVED FROM ASSEMBLY: " }
return json.Marshal(Logs) 
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
//=====================
// Get Attributes 
//=====================
func getAttribute(stub shim.ChaincodeStubInterface, attributeName string) (string, error) {
	bytes, err := stub.ReadCertAttribute(attributeName)
	return string(bytes[:]), err
}

//=====================
// Check Assembly 
//=====================
func checkIfAlreadyGotAssembly(stub shim.ChaincodeStubInterface, args string) error {
	var err error
	var jsonResp2 string
	part,err:=findPartById(stub,args)
	if(err !=nil){return err}
	ptAS, _ := json.Marshal(part)
	var pt Part
	err = json.Unmarshal(ptAS, &pt)
	if err != nil {return  errors.New("Failed to Unmarshal Part #" + args)}
	if ( pt.Helicopter == "") {return nil
	} else if ( pt.Helicopter != "") {jsonResp2 = "{\"Error\":\"Ce part appartient déjà a un aircraft\"}"
		return  errors.New(jsonResp2) }
	return nil 
}

//=====================
// Check Aircraft 
//=====================
func checkIfAlreadyGotAircraft(stub shim.ChaincodeStubInterface, args string) error {
	var err error
	var jsonResp2 string
	part,err:=findPartById(stub,args)
	if(err !=nil){return err}
	ptAS, _ := json.Marshal(part)
	var pt Part
	err = json.Unmarshal(ptAS, &pt)
	if err != nil {return  errors.New("Failed to Unmarshal Part #" + args)}
	if ( pt.Helicopter == "") {return nil
	} else if ( pt.Helicopter != "") {jsonResp2 = "{\"Error\":\"Ce part appartient déjà a un aircraft\"}"
		return  errors.New(jsonResp2) }
	return nil 
}


func UpdateAircraft(stub shim.ChaincodeStubInterface, airc Aircraft) error {

// Début Partie Aircraft 
	//Update allAircraft 
		partzMap,err:=getAircraftMap(stub)
			partzMap[airc.Id_Aircraft] = airc
			allPAsBuytes, err := json.Marshal(partzMap)
			err=stub.PutState("allAircraft",allPAsBuytes)
			if err != nil {return   err}
	//Fin update allAircraft 
	//Update allAircraftsAn
	partzMap1,err:=getAircraftAnMap(stub)
		partzMap1[airc.AN] = airc
		allPAsBytes11, err := json.Marshal(partzMap1)
		err=stub.PutState("allAircraftsAn",allPAsBytes11)
		if err != nil {return  err}
	//Fin update allAircraftsAn
	//Update allAircraftsSn
	partzMap2,err:=getAircraftSnMap(stub)
		partzMap2[airc.SN] = airc
		allPAsBytes22, err := json.Marshal(partzMap2)
		err=stub.PutState("allAircraftsSn",allPAsBytes22)
		if err != nil {return  err}
	//Fin update allAircraftsSn	
// Fin Partie Aircraft 

	return  nil 
}

func UpdateAssembly(stub shim.ChaincodeStubInterface, assemb Assembly) error {
// Début Partie Assembly 
//Update allAssembly 
		partzMap,err:=getAssemblyMap(stub)
		partzMap[assemb.Id_Assembly] = assemb
		allPAsBuytes, err := json.Marshal(partzMap)
		err=stub.PutState("allAssembly",allPAsBuytes)
		if err != nil {return err}
//Fin update allAssembly 
//Update allAssembliesAn
	partzMap1,err:=getAssembliesAnMap(stub)
		partzMap1[assemb.AN] = assemb
		allPAsBytes1, err := json.Marshal(partzMap1)
		err=stub.PutState("allAssembliesAn",allPAsBytes1)
		if err != nil {return  err}
//Fin update allAssembliesAn
//Update allAssembliesSn
	partzMap2,err:=getAssembliesSnMap(stub)
		partzMap2[assemb.SN] = assemb
		allPAsBytes2, err := json.Marshal(partzMap2)
		err=stub.PutState("allAssembliesSn",allPAsBytes2)
		if err != nil {return  err}
//Fin update allAssembliesSn
// Fin Partie Assembly 

	return  nil 
}

func UpdatePart(stub shim.ChaincodeStubInterface, pt Part) error {
// Début Partie Part 
//Update allParts 
	partMap,err:=getPartsIdMap(stub)
		partMap[pt.Id] = pt
		allPAsBytes, err := json.Marshal(partMap)
		err=stub.PutState("allParts",allPAsBytes)
		if err != nil {return  err}
//Fin update allParts 
//Update allPartsPn
	partMap1,err:=getPartsPnMap(stub)
		partMap1[pt.PN] = pt
		allPAsBytes1, err := json.Marshal(partMap1)
		err=stub.PutState("allPartsPn",allPAsBytes1)
		if err != nil {return  err}
//Fin update allPartsPn
//Update allPartsSn
	partMap2,err:=getPartsSnMap(stub)
		partMap2[pt.SN] = pt
		allPAsBytes2, err := json.Marshal(partMap2)
		err=stub.PutState("allPartsSn",allPAsBytes2)
		if err != nil {return  err}
//Fin update allPartsSn
// Fin Partie Part 
	return  nil 
}