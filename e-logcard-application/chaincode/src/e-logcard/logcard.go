//=============================================================================================================
//	 				e-LogCard CHAINCODE
//=============================================================================================================
package main
import (
	"errors"
	"fmt"
	"encoding/json"
	"github.com/hyperledger/fabric/core/chaincode/shim"
)
//=============================================================================================================
//	 Structure Definitions
//=============================================================================================================
//========================================================================================
//	Chaincode - A blank struct for use with Shim 
// (A HyperLedger included go file used for get/put state and other HyperLedger functions) 
//========================================================================================
type SimpleChaincode struct {
}
//=======================
//	Part 
//=======================
type Part struct { // Part et eLogcard sont regroupés dans cette première version
	PN string `json:"pn"` // Part Number
	SN string `json:"sn"` // Serial Number
	Id string `json:"id"` // Génération d'un UUID	
	PartName string `json:"partName"` // Nom de la Part 
	Type string `json:"type"` // Se renseigner sur les différents types de Parts 
	Owner string `json:"owner"` // Propriété portée par l'organisation
	Responsible string `json:"responsible"` // Responsable à l'instant T de la pièce (Portée par l'organisation)
	Helicopter	string `json:"helicopter"` // Aircraft
	Assembly string `json:"assembly"` // Assembly
	Logs []Log `json:"logs"` // Changements sur la part  + Transactions 
}
//================================================
//	Log - Defines the structure for a log object. 
//  It represents transactions for a part, states changes, maintenance tasks, etc..
//================================================
type Log struct { 
	LType string `json:"log_type"` // Type of change
	VDate string `json:"vDate"` // Date 
	Owner string `json:"owner"` // Owner of the part
	Responsible string `json:"responsible"` // Responsible of the part at the moment 
	ModType string `json:"modType"` // Type de modifications 
	Description string `json:"description"` // Description de la modification apportée 	
}
//================================================
// Aircraft
//================================================
type Aircraft struct { 
	Id_Aircraft string `json:"id_aircraft"` // Génération d'un UUID
	Owner string `json:"owner"` // Nom de la Part 
	Id_Parts []string  `json:"Id_parts"` // le faire composer d'id de part servant de clé 
}
//================================================
// Assembly 
//================================================ 
type Assembly struct { 
	Id_Assembly string `json:"id_assembly"` // Génération d'un UUID
	Owner string `json:"owner"` // Nom de la Part 
	Id_Parts []string  `json:"Id_parts"` // le faire composer d'id de part servant de clé 
}
// ============================================================================================================
// 					HYPERLEDGER FUNCTIONS
// ============================================================================================================
//============================================================
//	Init Function - Called when the user deploys the chaincode 
//============================================================
func (t *SimpleChaincode) Init(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {	
		
		n:= createMap(stub, "allParts")
			if n != nil { fmt.Println(n.Error()); return nil, errors.New(n.Error())}
		o:= createMap(stub, "allPartsPn")
			if o != nil { fmt.Println(o.Error()); return nil, errors.New(o.Error())}
		m:= createMap(stub, "allPartsSn")
			if m != nil { fmt.Println(m.Error()); return nil, errors.New(m.Error())}
		p:= createMap(stub, "allAircraft")
			if p != nil { fmt.Println(p.Error()); return nil, errors.New(p.Error())}
		q:= createMap(stub, "allAssembly")
			if q != nil { fmt.Println(q.Error()); return nil, errors.New(q.Error())}
	return nil, nil
}
// ========================================================
// Invoke is our entry point to invoke a chaincode function
// ========================================================
func (t *SimpleChaincode) Invoke(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {
    fmt.Println("invoke is running " + function)
	
	if function == "createPart" {
		role, err := getAttribute(stub, "role")
		if(role=="supplier" || role == "manufacturer"){ 
		return t.createPart(stub, args)
		}else { return []byte("You are not authorized"),err}}	
	if function == "ownershipTransfer" {
		role, err := getAttribute(stub, "role")
		if(role=="supplier" || role == "manufacturer" || role == "Customer" || role == "maintenance_user"){	
		return t.ownershipTransfer(stub, args)
		}else { return []byte("You are not authorized"),err}} 	
	if function == "responsibilityTransfer" {
		role, err := getAttribute(stub, "role")
		if(role=="supplier" || role == "manufacturer" || role == "Customer" || role == "maintenance_user"){	
		return t.responsibilityTransfer(stub, args)
		}else { return []byte("You are not authorized"),err}} 	
	if function == "performActivities" {
	/*	role, err := getAttribute(stub, "role")
		if(role=="supplier" || role == "manufacturer" || role == "Customer" || role == "maintenance_user"){	
			n:= checkResponsibility(stub, args[0])
			if n != nil {return []byte ("Vous n'êtes habilité à effectuer des activités sur cette part"), err }
			return t.performActivities(stub, args) 
		} else { return []byte("You are not authorized"),err}}  */
		return t.performActivities (stub,args)}
			
	fmt.Println("invoke did not find func: " + function)
	return nil, errors.New("Received unknown function invoke")
}
// =========================================================
// Query - read a variable from chaincode state - (aka read)  
// =========================================================
func (t *SimpleChaincode) Query(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {
    fmt.Println("query is running " + function)
	
// Audit functions 
	if function == "getPartDetails" {
		/*if len(args) != 1 {
		fmt.Println("Incorrect number of arguments. Expecting 1")
		return nil, errors.New("Incorrect number of arguments. Expecting 1: ID")}
		role, err := getAttribute(stub, "role")
		if(role=="supplier" || role == "manufacturer" || role == "Customer" || role == "maintenance_user"){	
		n:= checkOwnership(stub, args[0])
			if n != nil {return []byte("Vous n'êtes plus owner de cette part"),err }
		}	*/
	return t.getPartDetails (stub, args)}
	
	if function == "getAllPartsDetails" {
		return t.getAllPartsDetails (stub,args)}
				
	fmt.Println("query did not find func: " + function)
	return nil, errors.New("Received unknown function query")
}
// ============================================================================================================
// 					e-LogCard FUNCTIONS
// ============================================================================================================
// =========================================================================================
// 					PARTS
// =========================================================================================
// ===================================================================
// Creation of the Part (creation of the eLogcard)
// Only registered suppliers and manufacturers can create Parts.  
// ===================================================================

func (t *SimpleChaincode) createPart(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {
	fmt.Println("Running createPart")		
	var err error
// Part creation 
	var pt Part
	pt.PN = args[0]
	pt.SN = args [1]
	pt.Id = args[2]
	pt.PartName = args[3]
	pt.Type = args[4]
	pt.Owner = args[5]
	pt.Responsible = args[6]
	pt.Helicopter = args[7]
	pt.Assembly = args[8]
	var tx Log
	tx.Owner 		= pt.Owner
	tx.Responsible	= pt.Responsible
	tx.VDate 		= args[9]
	tx.LType 		= "CREATE"
	pt.Logs = append(pt.Logs, tx)
// If the PN or/and the SN is/are already used, a part can't be created.
	n:= checkPNavailibility(stub, args[0])
	if n != nil { return nil, errors.New(n.Error())}	
	o:= checkSNavailibility(stub, args[1])
	if o != nil { return nil, errors.New(o.Error())}
	//Commit part to ledger
		ptAsBytes, _ := json.Marshal(pt)
		err = stub.PutState(pt.SN, ptAsBytes)
		if err != nil {return nil, err}	
	//Update allParts 
		partMap,err:=getPartsIdMap(stub)
		partMap[pt.Id] = pt
		allPAsBytes, err := json.Marshal(partMap)
		err=stub.PutState("allParts",allPAsBytes)
		if err != nil {return nil, err}
	//Fin update allParts 
	//Update allPartsPn
		partMap1,err:=getPartsPnMap(stub)
		partMap1[pt.PN] = pt
		allPAsBytes1, err := json.Marshal(partMap1)
		err=stub.PutState("allPartsPn",allPAsBytes1)
		if err != nil {return nil, err}
	//Fin update allPartsPn
	//Update allPartsSn
		partMap2,err:=getPartsSnMap(stub)
		partMap2[pt.SN] = pt
		allPAsBytes2, err := json.Marshal(partMap2)
		err=stub.PutState("allPartsSn",allPAsBytes2)
		if err != nil {return nil, err}
	//Fin update allPartsSn
// Debut creation Aircraft
	u:= createAircraft(stub, args[7], args [5], args[2])
	if u != nil {
	fmt.Println(u.Error())
	return nil, errors.New(u.Error())}	
// Fin creation Aircraft
// Debut creation Assembly
	if ( pt.Assembly != "") {	
	d:= createAssembly(stub, args[8], args [5], args[2])
	if d != nil {
	fmt.Println(d.Error())
	return nil, errors.New(d.Error())}}
// Fin creation Assembly

	return []byte("eLogcardlogcard created successfully"),err
	fmt.Println("eLogcardlogcard created successfully")	
	return nil, nil
}

// ====================================================================
// Obtenir tous les détails d'une part à partir de son id 
// Registered suppliers, manufacturers, customers and maintenance users can  display details on a specific part only if they own it.
// Auditor_authority and AH_Admin can see details on any specific part they want.
// ====================================================================
func (t *SimpleChaincode) getPartDetails(stub shim.ChaincodeStubInterface, args []string)([]byte, error) {

// var key string
// var typ string 
typ := args[0]
key :=  args[1]

if ( typ == "Part") {
	part,err:=findPartById(stub,key)
	if(err !=nil){return nil,err}
	return json.Marshal(part)  	
} else if ( typ == "Aircraft"){	
	partMap,err:=getPartsIdMap(stub)
	if err != nil {return nil, errors.New("Failed to get Part")}
	parts := make([]Part, len(partMap))
	idx := 0
    for  _, part := range partMap {
    	if(part.Helicopter == key){
    		parts[idx] = part
    		idx++
    	}
    }
    //si les deux longueurs sont differentes on slice
    if(len(partMap)!=idx){
    	parts=parts[0:idx]
		}	 
    return json.Marshal(parts)
	
} else if ( typ == "Assembly") {	
	partMap,err:=getPartsIdMap(stub)
	if err != nil {return nil, errors.New("Failed to get Part")}
	parts := make([]Part, len(partMap))
	idx := 0
    for  _, part := range partMap {
		if (part.Assembly == key){
    		parts[idx] = part
    		idx++
    	}
    }
    //si les deux longueurs sont differentes on slice
    if(len(partMap)!=idx){
    	parts=parts[0:idx]
		}	 
    return json.Marshal(parts)
	} 
	return nil, nil 
}

// ==================================================================
// Afficher toutes les parts créées en détail  
// Registered suppliers, manufacturers, customers and maintenance users can display details of all the parts they own.
// Auditor_authority and AH_Admin can display details of all the parts ever created.
//===================================================================

func (t *SimpleChaincode) getAllPartsDetails(stub shim.ChaincodeStubInterface, args []string)([]byte, error){

var typ string 
typ = args[0]

if ( typ == "Part") {
// Parts
	fmt.Println("Start find getAllPartsDetails ")
	fmt.Println("Looking for All Parts With Details ")
	username, err := getAttribute(stub, "username")
	if(err !=nil){return nil,err}
	role, err := getAttribute(stub, "role")
	if(err !=nil){return nil,err}
	//if supplier or manufacturer or customer or maintenance user =>only my parts
	showOnlyMyPart := role=="supplier" || role == "manufacturer" || role == "Customer" || role == "maintenance_user"
	partMap,err:=getPartsIdMap(stub)
	parts := make([]Part, len(partMap))
    idx := 0
    for  _, part := range partMap {
    	if(!showOnlyMyPart ||  part.Owner == username){
    		parts[idx] = part
    		idx++
    	}
    }
    //si les deux longueurs sont differentes on slice
    if(showOnlyMyPart && len(partMap)!=idx){
    	parts=parts[0:idx]
    }
    return json.Marshal(parts) 
} else if ( typ == "Aircraft"){
// Aircrafts
	fmt.Println("Start find getAllAircraftDetails ")
	fmt.Println("Looking for All Aircrafts With Details ")
	
	partMap,err:=getAircraftMap(stub)
		if(err !=nil){return nil,err}

	parts := make([]Aircraft, len(partMap))
    idx := 0
    for  _, part := range partMap {
    		parts[idx] = part
    		idx++    	
    }
    //si les deux longueurs sont differentes on slice
    if(len(partMap)!=idx){
    	parts=parts[0:idx]
    }
    return json.Marshal(parts)

} else if ( typ == "Assembly") {
// Assemblies
	fmt.Println("Start find getAllAircraftDetails ")
	fmt.Println("Looking for All Aircrafts With Details ")
	
	partMap,err:=getAssemblyMap(stub)
		if(err !=nil){return nil,err}

	parts := make([]Assembly, len(partMap))
    idx := 0
    for  _, part := range partMap {
    		parts[idx] = part
    		idx++    	
    }
    //si les deux longueurs sont differentes on slice
    if(len(partMap)!=idx){
    	parts=parts[0:idx]
    }
    return json.Marshal(parts)
}
	return nil, nil 
}

// =========================================================================================
// 					ACTIVITIES 
// =========================================================================================
// =========================
// Transfert de propriété 
// =========================
// Only registered suppliers, manufacturers, Customers and maintenance_user can Transfer Ownership on a Part.
// Provided that they are currently owner of this part.

func (t *SimpleChaincode) ownershipTransfer(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {
	var err error
	var key string 
	key = args[0]
	part,err:=findPartById(stub,key)
	if err != nil {return nil, errors.New("Failed to get part #" + key)}
	ptAS, _ := json.Marshal(part)
	var pt Part
	err = json.Unmarshal(ptAS, &pt)
	if err != nil {return nil, errors.New("Failed to Unmarshal Part #" + key)}
	pt.Owner = args[1]
	var tx Log
	tx.Owner 		= pt.Owner
	tx.VDate 		= args[2]
	tx.LType 		= "OWNERNSHIP_TRANSFER"
	pt.Logs = append(pt.Logs, tx)
	
	//Update allParts 
		partMap,err:=getPartsIdMap(stub)
		partMap[pt.Id] = pt
		allPAsBytes, err := json.Marshal(partMap)
		err=stub.PutState("allParts",allPAsBytes)
	if err != nil {return nil, err}
	//Fin update allParts 
	//Update allPartsPn
		partMap1,err:=getPartsPnMap(stub)
		partMap1[pt.PN] = pt
		allPAsBytes1, err := json.Marshal(partMap1)
		err=stub.PutState("allPartsPn",allPAsBytes1)
		if err != nil {return nil, err}
	//Fin update allPartsPn
	//Update allPartsSn
		partMap2,err:=getPartsSnMap(stub)
		partMap2[pt.SN] = pt
		allPAsBytes2, err := json.Marshal(partMap2)
		err=stub.PutState("allPartsSn",allPAsBytes2)
		if err != nil {return nil, err}
	//Fin update allPartsSn
	
	/*
	// Update allAircraft
	key = pt.Helicopter
	aircraft, err := findAircraftById(stub,key)
	if err != nil {return nil, errors.New("Failed to get part #" + key)}
	ptA1, _ := json.Marshal(aircraft)
	var aire Aircraft
	err = json.Unmarshal(ptA1, &aire)
	aire.Id_Parts = append(aire.Id_Parts, pt)
	partzMap,err:=getAircraftMap(stub)
	partzMap[aire.Id_Aircraft] = aire
	allPAsBuytes, err := json.Marshal(partzMap)
	err=stub.PutState("allAircraft",allPAsBuytes)
	if err != nil {return nil, err}
	//Fin update allParts 	
	*/

	return nil, nil
}
// =============================
// Transfert de responsabilité 
// =============================
// Only registered suppliers, manufacturers, Customers and maintenance_user can Transfer Responsibility on a Part.
// Provided that they are currently owner of this part.

func (t *SimpleChaincode) responsibilityTransfer(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {
	var err error
	var key string 
	key = args[0]
	part,err:=findPartById(stub,key)
	if err != nil {return nil, errors.New("Failed to get part #" + key)}
	ptAS, _ := json.Marshal(part)
	var pt Part
	err = json.Unmarshal(ptAS, &pt)
	if err != nil {return nil, errors.New("Failed to Unmarshal Part #" + key)}
	pt.Responsible = args[1]
	var tx Log
	tx.Responsible 	= pt.Responsible
	tx.VDate 		= args[2]
	tx.LType 		= "RESPONSIBILITY_TRANSFER"
	pt.Logs = append(pt.Logs, tx)
	
	//Update allParts 
		partMap,err:=getPartsIdMap(stub)
		partMap[pt.Id] = pt
		allPAsBytes, err := json.Marshal(partMap)
		err=stub.PutState("allParts",allPAsBytes)
		if err != nil {return nil, err}
	//Fin update allParts 
	//Update allPartsPn
		partMap1,err:=getPartsPnMap(stub)
		partMap1[pt.PN] = pt
		allPAsBytes1, err := json.Marshal(partMap1)
		err=stub.PutState("allPartsPn",allPAsBytes1)
		if err != nil {return nil, err}
	//Fin update allPartsPn
	//Update allPartsSn
		partMap2,err:=getPartsSnMap(stub)
		partMap2[pt.SN] = pt
		allPAsBytes2, err := json.Marshal(partMap2)
		err=stub.PutState("allPartsSn",allPAsBytes2)
		if err != nil {return nil, err}
	//Fin update allPartsSn
	
	return nil, nil
}

// =========================
// Acitivités sur la part 
// =========================
func (t *SimpleChaincode) performActivities(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {

	var err error
	var key string 
	key = args[0]
	part,err:=findPartById(stub,key)
	if err != nil {return nil, errors.New("Failed to get part #" + key)}
	ptAS, _ := json.Marshal(part)
	var pt Part
	err = json.Unmarshal(ptAS, &pt)
	if err != nil {return nil, errors.New("Failed to Unmarshal Part #" + key)}
	var tx Log
	tx.Owner 	= pt.Owner
	tx.Responsible 	= pt.Responsible
	tx.ModType = args[1]
	tx.Description = args[2]
	tx.VDate 		= args[3]
	tx.LType 		= "ACTIVITIES"
	pt.Logs = append(pt.Logs, tx)
	
	//Update allParts 
		partMap,err:=getPartsIdMap(stub)
		partMap[pt.Id] = pt
		allPAsBytes, err := json.Marshal(partMap)
		err=stub.PutState("allParts",allPAsBytes)
		if err != nil {return nil, err}
	//Fin update allParts 
	//Update allPartsPn
		partMap1,err:=getPartsPnMap(stub)
		partMap1[pt.PN] = pt
		allPAsBytes1, err := json.Marshal(partMap1)
		err=stub.PutState("allPartsPn",allPAsBytes1)
		if err != nil {return nil, err}
	//Fin update allPartsPn
	//Update allPartsSn
		partMap2,err:=getPartsSnMap(stub)
		partMap2[pt.SN] = pt
		allPAsBytes2, err := json.Marshal(partMap2)
		err=stub.PutState("allPartsSn",allPAsBytes2)
		if err != nil {return nil, err}
	//Fin update allPartsSn

	return nil, nil
}




// =========================================================================================
// 					UTILITY FUNCTIONS
// =========================================================================================
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
// Get Part Map 
// recuperation de la map des parts par id
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
// Get Part Map 
// recuperation de la map des parts par PN
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
// Get Part Map 
// recuperation de la map des parts par SN
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
// Get Aircraft Map 
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
// Get Assembly Map 
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

//=============================================================================================================
// Main - main - Starts up the chaincode  
//=============================================================================================================
func main() {
	fmt.Println("Welcome to eLogcard System!")
	err := shim.Start(new(SimpleChaincode))
	if err != nil {
		fmt.Printf("Error starting Simple chaincode: %s", err)
	}
}