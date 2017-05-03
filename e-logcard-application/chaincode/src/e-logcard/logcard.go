package main
import (
	"errors"
	"fmt"
	"encoding/json"

	"github.com/hyperledger/fabric/core/chaincode/shim"
)
//==============================================================================================================================
//	 Structure Definitions
//==============================================================================================================================

//========================================================================================
//	Chaincode - A blank struct for use with Shim 
// (A HyperLedger included go file used for get/put state and other HyperLedger functions) 
//========================================================================================
type SimpleChaincode struct {
}
//==============================================
//	Part 
//==============================================
type Part struct { // Part et eLogcard sont regroupés dans cette première version
	PN			string 	`json:"pn"` 					// Part Number
	SN 			string 	`json:"sn"` 					// Serial Number
	Id   		string  `json:"id"` 					// Concaténation des deux PN et SN	
	PartName	string  `json:"partName"` 
	Type		string 	`json:"type"` 
	Owner       string  `json:"owner"` 
	Responsible string  `json:"responsible"` 			// Responsable à l'instant T de la pièce
	Helicopter	string  `json:"helicopter"` 
	Assembly    string  `json:"assembly"`
	Logs        []Log 	`json:"logs"` 					// changement sur la part  
}
//========================================================
//	AllParts 
//========================================================
type AllParts struct{
	Parts []string `json:"parts"`
}
//=========================================================
//	AllParts Details 
//=========================================================
type AllPartsDetails struct{
	Parts []Part `json:"parts"`
}
//================================================
//	Log - Defines the structure for a log object. 
//  It represents transactions for a part, states changes, maintenance tasks, etc..
//================================================
type Log struct { 
	Owner 		string 	`json:"owner"` 				// Owner of the part
	Responsible string  `json:"responsible"`        // Responsible of the part at the moment 
	VDate 		string   `json:"vDate"` 			// Date 
	Helicopter	string  `json:"helicopter"` 
	Assembly    string  `json:"assembly"`
	ModType     string   `json:"modType"` 			// Type de modifications 
	Description string   `json:"description"` 		// Description de la modification apportée 	
	LType 		string   `json:"ltype"` 			// Type of change
}

//=======================================================
//	Assembly 
//=======================================================
type Assembly struct{
	AssemblyName	string	`json:"assembly_name"` 
}
//=======================================================
//	AllAssemblies
//=======================================================
type AllAssemblies struct{
	Assemblies []string `json:"assemblies"`
}


//=======================================================
//	Aircraft 
//=======================================================
type Aircraft struct{
	AircraftName	string	`json:"aircraft_name"` 
	// Parts []		string `json:"parts"`
}
//=======================================================
//	AllAircrafts
//=======================================================
type AllAircrafts struct{
	Aircrafts []string `json:"aircrafts"`
}
//========================================================
//	AllAircraftsDetails
//========================================================
type AllAircraftDetails struct{
	Aircrafts []Aircraft `json:"aircrafts"`
}

//=======================================================
//	Owner  
//=======================================================
type Owner struct {
	OwnerName		 string	`json:"owner_name"` 
	// OwnerId   		string  `json:"owner_id"` 
	//PartId 			 string `json:"part_id"`
	Parts        	 []Part `json:"parts"`
	Statut        string	`json:"statut"`  
	
}
//========================================================
//	AllOwners  
//========================================================
type AllOwners struct{
	Owners []string `json:"owners"`
}
//==========================================================
//	Responsible   
//==========================================================
type Responsible struct {
	ResponsibleName		string	`json:"responsible_name"` 
	// responsibleId   		string  `json:"responsible_id"` 	
}
//===========================================================
//	AllResponsibles  
//===========================================================
type AllResponsibles struct{
	Responsibles []string `json:"responsibles"`
}



// ============================================================================================================================================
// HYPERLEDGER FUNCTIONS
// ============================================================================================================================================
//===============================================================
//	Init Function - Called when the user deploys the chaincode 
//===============================================================
func (t *SimpleChaincode) Init(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {
	var err error
	var parts AllParts
	jsonAsBytes, _ := json.Marshal(parts)
	err = stub.PutState("allParts", jsonAsBytes)
	if err != nil {return nil, err}
	
	var owners AllOwners
	jsonAsaBytes, _ := json.Marshal(owners)
	err = stub.PutState("allOwners", jsonAsaBytes)
	if err != nil {return nil, err}
	
	var responsibles AllResponsibles 
	jsonAsaaBytes, _ := json.Marshal(responsibles)
	err = stub.PutState("allResponsibles", jsonAsaaBytes)
	if err != nil {return nil, err}
	
	var aircrafts AllAircrafts
	jsonAsairBytes, _ := json.Marshal(aircrafts)
	err = stub.PutState("allAircrafts", jsonAsairBytes)
	if err != nil {return nil, err}
	
	var assemblies AllAssemblies 
	jsonAssBytes, _ := json.Marshal(assemblies)
	err = stub.PutState("allAssemblies", jsonAssBytes)
	if err != nil {return nil, err}
	
	// créer une structure AllOrganizations 
	// append directement une orga appelée easa 
	// test_user2
	return nil, nil
}
// ================================================================
// Invoke is our entry point to invoke a chaincode function
// ================================================================
func (t *SimpleChaincode) Invoke(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {
    fmt.Println("invoke is running " + function)
   
	// Users functions 
   if function == "createPart" {
		if len(args) != 10 {
		fmt.Println("Incorrect number of arguments. Expecting 10")
		return nil, errors.New("Incorrect number of arguments. Expecting 10: PN, SN, ID, PartName, Type, Owner, Responsible, Helicopter, Assembly, Date")}
		role, err := getAttribute(stub, "role")
		if(role=="supplier" || role == "manufacturer"){
		return t.createPart(stub, args)
		}else { return []byte("You are not authorized"),err
		}
		}	
	if function == "ownershipTransfer" {
		if len(args) != 3 {
		fmt.Println("Incorrect number of arguments. Expecting 3")
		return nil, errors.New("Incorrect number of arguments. Expecting 3: ID, New Owner, Date")}
		role, err := getAttribute(stub, "role")
		if(role=="supplier" || role == "manufacturer" || role == "Customer" || role == "maintenance_user"){	
		return t.ownershipTransfer(stub, args)
		}else { return []byte("You are not authorized"),err
		}
		} 	
	if function == "responsibilityTransfer" {
		if len(args) != 3 {
		fmt.Println("Incorrect number of arguments. Expecting 3")
		return nil, errors.New("Incorrect number of arguments. Expecting 3: ID, New responsible, Date")}
		role, err := getAttribute(stub, "role")
		if(role=="supplier" || role == "manufacturer" || role == "Customer" || role == "maintenance_user"){	
		return t.responsibilityTransfer(stub, args)
		}else { return []byte("You are not authorized"),err
		}
		} 	
	if function == "assemblyTransfer" {
		if len(args) != 3 {
		fmt.Println("Incorrect number of arguments. Expecting 3")
		return nil, errors.New("Incorrect number of arguments. Expecting 3: ID, New responsible, Date")}
		role, err := getAttribute(stub, "role")
		if(role == "maintenance_user"){	
		return t.assemblyTransfer(stub, args)
		}else { return []byte("You are not authorized"),err
		}
		} 
	if function == "helicoTransfer"{
		if len(args) != 3 {
		fmt.Println("Incorrect number of arguments. Expecting 3")
		return nil, errors.New("Incorrect number of arguments. Expecting 3: ID, New responsible, Date")}
		role, err := getAttribute(stub, "role")
		if(role == "maintenance_user"){	
		return t.helicoTransfer(stub, args)
		}else { return []byte("You are not authorized"),err
		}
		}
	if function == "performActivities" {
		if len(args) != 4 {
		fmt.Println("Incorrect number of arguments. Expecting 4")
		return nil, errors.New("Incorrect number of arguments. Expecting 4: ID, ModType, Description, Vdate")}
		// role, err := getAttribute(stub, "role")
		// if(role=="maintenance_user"){
		return t.performActivities(stub, args)
		// }else{
		// return []byte("You are not authorized"),err
		// }
		}
	fmt.Println("invoke did not find func: " + function)
	return nil, errors.New("Received unknown function invoke")
}
// ===============================================================
// Query - read a variable from chaincode state - (aka read)  
// ===============================================================
func (t *SimpleChaincode) Query(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {
    fmt.Println("query is running " + function)
	
	// Audit functions 
	if function == "getPartDetails" {
		if len(args) != 1 {
		fmt.Println("Incorrect number of arguments. Expecting 1")
		return nil, errors.New("Incorrect number of arguments. Expecting 1: ID")}
		return t.getPartDetails (stub,args) 
		}
	if function == "getAllParts" {
		return t.getAllParts (stub,args) 
		}
	if function == "getAllPartsDetails" {
		return t.getAllPartsDetails (stub,args) 
		}
	if function == "getAllAssemblies" {
		return t.getAllAssemblies (stub,args) 
		}
	if function == "getAllOwners" {
		return t.getAllOwners (stub,args) 
		}
	if function == "getOwnerDetails" {
		return t.getOwnerDetails (stub,args) 
		}
	if function == "getAllPartsOfThisOwner" {
		if len(args) != 1 {
		fmt.Println("Incorrect number of arguments. Expecting 1")
		return nil, errors.New("Incorrect number of arguments. Expecting 1: Owner")}
		return t.getAllPartsOfThisOwner (stub,args) 
		}
	if function == "getAllPartsOfThisOwnerInDetails" {
		if len(args) != 1 {
		fmt.Println("Incorrect number of arguments. Expecting 1")
		return nil, errors.New("Incorrect number of arguments. Expecting 1: Owner")}
		return t.getAllPartsOfThisOwnerInDetails (stub,args) 
		}
	if function == "getAllResponsibles" {
		return t.getAllResponsibles (stub,args) 
		}
	if function == "getAllHelico" {
		return t.getAllHelico (stub,args) 
		}
	if function == "getAllPartsOfThisResponsible" {
		if len(args) != 1 {
		fmt.Println("Incorrect number of arguments. Expecting 1")
		return nil, errors.New("Incorrect number of arguments. Expecting 1: Responsible")}
		return t.getAllPartsOfThisResponsible (stub,args) 
		}
	if function == "getAllPartsOfThisResponsibleInDetails" {
		if len(args) != 1 {
		fmt.Println("Incorrect number of arguments. Expecting 1")
		return nil, errors.New("Incorrect number of arguments. Expecting 1: Responsible")}
		return t.getAllPartsOfThisResponsibleInDetails (stub,args) 
		}
		
	// Users functions 
	if function == "getAllPartsOfThisAssembly" {
		if len(args) != 1 {
		fmt.Println("Incorrect number of arguments. Expecting 1")
		return nil, errors.New("Incorrect number of arguments. Expecting 1: Helicopter")}
		return t.getAllPartsOfThisAssembly (stub,args) }
	if function == "getAllPartsOfThisAssemblyInDetails" {
		if len(args) != 1 {
		fmt.Println("Incorrect number of arguments. Expecting 1")
		return nil, errors.New("Incorrect number of arguments. Expecting 1: Helicopter")}
		return t.getAllPartsOfThisAssemblyInDetails (stub,args) }
	if function == "getAllPartsOfThisHelico" {
		if len(args) != 1 {
		fmt.Println("Incorrect number of arguments. Expecting 1")
		return nil, errors.New("Incorrect number of arguments. Expecting 1: Helicopter")}
		return t.getAllPartsOfThisHelico (stub,args) }
	if function == "getAllPartsOfThisHelicoInDetails" {
		if len(args) != 1 {
		fmt.Println("Incorrect number of arguments. Expecting 1")
		return nil, errors.New("Incorrect number of arguments. Expecting 1: Helicopter")}
		return t.getAllPartsOfThisHelicoInDetails (stub,args) }
		
	// Test functions 
	if function == "hello" {return t.hello (stub,args) }
	if function =="user_name" {
		role, err := getAttribute(stub, "role")
		if(role=="supplier"){
			return []byte("Welcome, please click on the link"),err
		} else{
			return []byte("You are not authorized"),err
		}
		
		} // return stub.ReadCertAttribute("username")
		fmt.Println("query did not find func: " + function)
		return nil, errors.New("Received unknown function query")
		}
	


// ============================================================================================================================================
// PARTS
// ============================================================================================================================================
// ===================================================================
// Creation of the Part (creation of the eLogcard) 
// ===================================================================
func (t *SimpleChaincode) createPart(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {
	fmt.Println("Running createPart")		
	var err error
	// Parts
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
	tx.Helicopter   = pt.Helicopter
	tx.Assembly     = pt.Assembly
	tx.VDate 		= args[9]
	tx.LType 		= "CREATE"
	pt.Logs = append(pt.Logs, tx)
		
	n:= checkPNavailibility(stub, args[0])
	if n != nil {
		fmt.Println(n.Error())
		return nil, errors.New(n.Error())}	
	o:= checkSNavailibility(stub, args[1])
	if o != nil {
		fmt.Println(o.Error())
		return nil, errors.New(o.Error())}
	m:= checkIDavailibility(stub, args[2])
	if m != nil {
		fmt.Println(m.Error())
		return nil, errors.New(m.Error())}

	//Commit part to ledger
	ptAsBytes, _ := json.Marshal(pt)
	err = stub.PutState(pt.Id, ptAsBytes)	
	err = stub.PutState(pt.PN, ptAsBytes)	
	err = stub.PutState(pt.SN, ptAsBytes)
	// err = stub.PutState(pt.Helicopter, ptAsBytes)
	if err != nil {return nil, err}
	
		
	// All Parts - Update AllParts Array
	allPAsBytes, err := stub.GetState("allParts")
	if err != nil {return nil, errors.New("Failed to get all Parts")}
	var allpt AllParts
	err = json.Unmarshal(allPAsBytes, &allpt)
	if err != nil {return nil, errors.New("Failed to Unmarshal all Parts")}
	allpt.Parts = append(allpt.Parts,pt.Id)
	
	//Commit AllParts to ledger	
	allPuAsBytes, _ := json.Marshal(allpt)
	err = stub.PutState("allParts", allPuAsBytes)	
	if err != nil {return nil, err}
	
	z:= createOwner(stub, args[5] /*, args[2] */, pt)
	if z != nil {
		fmt.Println(z.Error())
		return nil, errors.New(z.Error())}	
	
	y:= createResponsible(stub, args[6])
	if y != nil {
		fmt.Println(y.Error())
		return nil, errors.New(y.Error())}	
	
	r:= createAircraft(stub, args[7])
	if r != nil {
		fmt.Println(r.Error())
		return nil, errors.New(r.Error())}
		
		
	if (args[8] != "") {
	w:= createAssembly(stub, args[8])
	if w != nil {
		fmt.Println(w.Error())
		return nil, errors.New(w.Error())}
		}
	return []byte("eLogcardlogcard created successfully"),err
	fmt.Println("eLogcardlogcard created successfully")	
	return nil, nil
	}
// ====================================================================
// Obtenir tous les détails d'une part à partir de son id 
// ====================================================================
func (t *SimpleChaincode) getPartDetails(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {
	var key, jsonResp string
	var err error
	key = args[0]
	valAsbytes, err := stub.GetState(key)
	if err != nil {
		jsonResp = "{\"Error\":\"Failed to get state for " + key + "\"}"
		return nil, errors.New(jsonResp)
	}
	return valAsbytes, nil 
}
// ===================================================================
// Afficher toutes les parts créées   
//====================================================================
func (t *SimpleChaincode) getAllParts(stub shim.ChaincodeStubInterface, args []string)([]byte, error){
	
	fmt.Println("Start find getAllParts ")
	fmt.Println("Looking for All Parts ")
	//get the AllParts index
	allPAsBytes, err := stub.GetState("allParts")
	if err != nil {return nil, errors.New("Failed to get all Parts")}

	var res AllParts
	err = json.Unmarshal(allPAsBytes, &res)
	if err != nil {return nil, errors.New("Failed to Unmarshal all Parts")}
	var rap AllParts
	for i := range res.Parts{
		spAsBytes, err := stub.GetState(res.Parts[i])
		if err != nil {return nil, errors.New("Failed to get Part")}
		var sp Part
		json.Unmarshal(spAsBytes, &sp)
		rap.Parts = append(rap.Parts,sp.Id); 
	}
	rapAsBytes, _ := json.Marshal(rap)
	return rapAsBytes, nil
}
// ==================================================================
// Afficher toutes les parts créées en détail  
//===================================================================
func (t *SimpleChaincode) getAllPartsDetails(stub shim.ChaincodeStubInterface, args []string)([]byte, error){
	
	fmt.Println("Start find getAllPartsDetails ")
	fmt.Println("Looking for All Parts With Details ")
	
	//get the AllParts index
	allPAsBytes, err := stub.GetState("allParts")
	if err != nil {return nil, errors.New("Failed to get all Parts")}
	
	var res AllParts
	err = json.Unmarshal(allPAsBytes, &res)
	if err != nil {return nil, errors.New("Failed to Unmarshal all Parts")}
	
	var rap AllPartsDetails
	
	for i := range res.Parts{

	spAsBytes, err := stub.GetState(res.Parts[i])
		if err != nil {return nil, errors.New("Failed to get Part")}
		var sp Part
		json.Unmarshal(spAsBytes, &sp)
		
		rap.Parts = append(rap.Parts,sp); 
	}
	rapAsBytes, _ := json.Marshal(rap)
	return rapAsBytes, nil
}
// ===========================================================================================================================================
// PROPRIETAIRES
// ============================================================================================================================================
// ================================================================
// Afficher toutes les propriétaires 
//=================================================================
func (t *SimpleChaincode) getAllOwners(stub shim.ChaincodeStubInterface, args []string)([]byte, error){
	
	fmt.Println("Start find getAllOwners ")
	fmt.Println("Looking for All Owners ")
	
	//get the AllOwners index
	allOAsBytes, err := stub.GetState("allOwners")
	if err != nil {return nil, errors.New("Failed to get all Owners")}
	var res AllOwners
	err = json.Unmarshal(allOAsBytes, &res)
	if err != nil {return nil, errors.New("Failed to Unmarshal all Owners")}
	var rap AllOwners
	for i := range res.Owners{
		soAsBytes, err := stub.GetState(res.Owners[i])
		if err != nil {return nil, errors.New("Failed to get Owner")}
		var so Owner
		json.Unmarshal(soAsBytes, &so)
		rap.Owners = append(rap.Owners,so.OwnerName)
		rap.Owners = append(rap.Owners,so.Statut);
	}
	rapAsBytes, _ := json.Marshal(rap)
	return rapAsBytes, nil	
}

// ====================================================================
// Obtenir tous les détails d'un Owner à partir de son id 
// ====================================================================
func (t *SimpleChaincode) getOwnerDetails(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {
	var key, jsonResp string
	var err error
	key = args[0]
	valAsbytes, err := stub.GetState(key)
	if err != nil {
		jsonResp = "{\"Error\":\"Failed to get state for " + key + "\"}"
		return nil, errors.New(jsonResp)
	}
	return valAsbytes, nil 
}
// ===================================================================
// Afficher toutes les parts d'un propriétaire 
//====================================================================
func (t *SimpleChaincode) getAllPartsOfThisOwner(stub shim.ChaincodeStubInterface, args [] string)([]byte, error){
			
	fmt.Println("Start find getAllParts ")
	fmt.Println("Looking for All Parts ")
	
	var key string 
	key = args[0]

	//get the AllParts index
	allPAsBytes, err := stub.GetState("allParts")
	if err != nil {return nil, errors.New("Failed to get all Parts")}

	var res AllParts
	err = json.Unmarshal(allPAsBytes, &res)
	if err != nil {return nil, errors.New("Failed to Unmarshal all Parts")}

	var rap AllParts

	for i := range res.Parts{
		spAsBytes, err := stub.GetState(res.Parts[i])
		if err != nil {return nil, errors.New("Failed to get Part")}
		var sp Part
		json.Unmarshal(spAsBytes, &sp) 
		if(sp.Owner == key) { rap.Parts = append(rap.Parts,sp.Id); 
		}
	}
	rabAsBytes, _ := json.Marshal(rap)
	return rabAsBytes, nil
}
// =====================================================================
// Afficher toutes les parts d'un propriétaire en détails 
//======================================================================
func (t *SimpleChaincode) getAllPartsOfThisOwnerInDetails(stub shim.ChaincodeStubInterface, args []string)([]byte, error){
	fmt.Println("Start find getAllPartsOfThisOwnerInDetails ")
	fmt.Println("Looking for All Parts of this Owner in Details ")
	
	var key string 
	key = args[0]
	
	//get the AllParts index
	allPAsBytes, err := stub.GetState("allParts")
	if err != nil {return nil, errors.New("Failed to get all Parts")}
	
	var res AllParts
	err = json.Unmarshal(allPAsBytes, &res)
	if err != nil {return nil, errors.New("Failed to Unmarshal all Parts")}
	
	//var res1 AllParts
	//err = json.Unmarshal(allPAsBytes, &res1)
	//if err != nil {return nil, errors.New("Failed to Unmarshal all Parts")}
	
	var rap AllPartsDetails
	//var rap1 AllParts

	for i := range res.Parts{
	spAsBytes, err := stub.GetState(res.Parts[i])
		if err != nil {return nil, errors.New("Failed to get Part")}
		var sp Part
		json.Unmarshal(spAsBytes, &sp)
		if(sp.Owner == key) { rap.Parts = append(rap.Parts,sp); 
		} 
		
		/* 
		else { 
		for i := range res1.Parts{
		spAsBytes1, err := stub.GetState(res1.Parts[i])
		if err != nil {return nil, errors.New("Failed to get Part")}
		var sp1 Part
		json.Unmarshal(spAsBytes1, &sp1) 
		rap1.Parts = append(rap1.Parts,"Cette organization n'est pas propriétaire de parts");
		}
		}
		*/
}
	rapAsBytes, _ := json.Marshal(rap)
	//rapAsBytes1, _ := json.Marshal(rap1)
	return rapAsBytes, nil
	//return rapAsBytes1, nil
}
// ============================================================================================================================================
// RESPONSABLES
// ============================================================================================================================================
// ========================================================================
// Afficher toutes les responsables  
//=========================================================================
func (t *SimpleChaincode) getAllResponsibles(stub shim.ChaincodeStubInterface, args []string)([]byte, error){

	//get the getAllResponsibles index
	allRAsBytes, err := stub.GetState("allResponsibles")
	if err != nil {return nil, errors.New("Failed to get all Responsibles")}

	var res AllResponsibles
	err = json.Unmarshal(allRAsBytes, &res)
	if err != nil {return nil, errors.New("Failed to Unmarshal all Responsibles")}

	var rap AllResponsibles
	for i := range res.Responsibles{
		srAsBytes, err := stub.GetState(res.Responsibles[i])
		if err != nil {return nil, errors.New("Failed to get Responsible")}
		var sr Responsible
		json.Unmarshal(srAsBytes, &sr)
		rap.Responsibles = append(rap.Responsibles,sr.ResponsibleName);
	}
	rapAsBytes, _ := json.Marshal(rap)
	return rapAsBytes, nil
}
// ==========================================================================
// Afficher toutes les parts d'un responsable  
//===========================================================================
func (t *SimpleChaincode) getAllPartsOfThisResponsible(stub shim.ChaincodeStubInterface, args [] string)([]byte, error){
	fmt.Println("Start find getAllParts ")
	fmt.Println("Looking for All Parts ")
	
	var key string 
	key = args[0]

	//get the AllParts index
	allPAsBytes, err := stub.GetState("allParts")
	if err != nil {return nil, errors.New("Failed to get all Parts")}

	var res AllParts
	err = json.Unmarshal(allPAsBytes, &res)
	if err != nil {return nil, errors.New("Failed to Unmarshal all Parts")}

	var rap AllParts

	for i := range res.Parts{

		spAsBytes, err := stub.GetState(res.Parts[i])
		if err != nil {return nil, errors.New("Failed to get Part")}
		
		var sp Part
		json.Unmarshal(spAsBytes, &sp)
		if(sp.Responsible == key) { rap.Parts = append(rap.Parts,sp.Id); 		
		} else {  rap.Parts = append(rap.Parts, " Cette organization n'est responsable d'aucune part"); 
		}
	}
	rabAsBytes, _ := json.Marshal(rap)
	return rabAsBytes, nil
}
// ==========================================================================
// Afficher toutes les parts d'un responsable en détail  
//===========================================================================
func (t *SimpleChaincode) getAllPartsOfThisResponsibleInDetails(stub shim.ChaincodeStubInterface, args []string)([]byte, error){
	fmt.Println("Start find getAllPartsOfThisResponsibleInDetails ")
	fmt.Println("Looking for All Parts of this Responsible in Details ")
	
	var key string 
	key = args[0]
	
	//get the AllParts index
	allPAsBytes, err := stub.GetState("allParts")
	if err != nil {return nil, errors.New("Failed to get all Parts")}
	
	var res AllParts
	err = json.Unmarshal(allPAsBytes, &res)
	if err != nil {return nil, errors.New("Failed to Unmarshal all Parts")}
	
	var rap AllPartsDetails
	
	for i := range res.Parts{

	spAsBytes, err := stub.GetState(res.Parts[i])
		if err != nil {return nil, errors.New("Failed to get Part")}
		var sp Part
		json.Unmarshal(spAsBytes, &sp)
		if(sp.Responsible == key) { rap.Parts = append(rap.Parts,sp); 
		} 
		//else { rap.Parts = append(rap.Parts, " Cette organization n'est  responsable d'aucune part"); 
		//}
	}
	rapAsBytes, _ := json.Marshal(rap)
	return rapAsBytes, nil
}
// ===========================================================================================================================================
// Helicopter
// ============================================================================================================================================
// ===========================================================
// Afficher tous les hélicos   
//============================================================
func (t *SimpleChaincode) getAllHelico(stub shim.ChaincodeStubInterface, args []string)([]byte, error){
	
	fmt.Println("Start find getAllOwners ")
	fmt.Println("Looking for All Owners ")

	//get the AllOwners index
	allOAsBytes, err := stub.GetState("allAircrafts")
	if err != nil {return nil, errors.New("Failed to get all Owners")}

	var res AllAircrafts
	err = json.Unmarshal(allOAsBytes, &res)
	if err != nil {return nil, errors.New("Failed to Unmarshal all Owners")}

	var rap AllAircrafts
	for i := range res.Aircrafts{
		soAsBytes, err := stub.GetState(res.Aircrafts[i])
		if err != nil {return nil, errors.New("Failed to get Owner")}
		var so Aircraft
		json.Unmarshal(soAsBytes, &so)
		rap.Aircrafts = append(rap.Aircrafts,so.AircraftName); 
	}
	rapAsBytes, _ := json.Marshal(rap)
	return rapAsBytes, nil
}
// ==========================================================
// Afficher toutes les parts d'un Hélicoptère
//===========================================================
func (t *SimpleChaincode) getAllPartsOfThisHelico(stub shim.ChaincodeStubInterface, args [] string)([]byte, error){
	fmt.Println("Start find getAllParts ")
	fmt.Println("Looking for All Parts ")
	
	var key string 
	key = args[0]

	//get the AllParts index
	allPAsBytes, err := stub.GetState("allParts")
	if err != nil {return nil, errors.New("Failed to get all Parts")}

	var res AllParts
	err = json.Unmarshal(allPAsBytes, &res)
	if err != nil {return nil, errors.New("Failed to Unmarshal all Parts")}

	var rap AllParts

	for i := range res.Parts{

		spAsBytes, err := stub.GetState(res.Parts[i])
		if err != nil {return nil, errors.New("Failed to get Part")}
		
		var sp Part
		json.Unmarshal(spAsBytes, &sp)
		if(sp.Helicopter == key) {rap.Parts = append(rap.Parts,sp.Id); 
		} else { rap.Parts = append(rap.Parts," Cet Helicoptère n'est pas composé de part ")
		}
	}
	rabAsBytes, _ := json.Marshal(rap)
	return rabAsBytes, nil
}
// ==========================================================
// Afficher toutes les parts d'un Hélicoptère en détail
//===========================================================
func (t *SimpleChaincode) getAllPartsOfThisHelicoInDetails(stub shim.ChaincodeStubInterface, args [] string)([]byte, error){
	fmt.Println("Start find getAllParts ")
	fmt.Println("Looking for All Parts ")
	
	var key string 
	key = args[0]

	//get the AllParts index
	allPAsBytes, err := stub.GetState("allParts")
	if err != nil {return nil, errors.New("Failed to get all Parts")}

	var res AllParts
	err = json.Unmarshal(allPAsBytes, &res)
	if err != nil {return nil, errors.New("Failed to Unmarshal all Parts")}

	var rap AllPartsDetails

	for i := range res.Parts{

		spAsBytes, err := stub.GetState(res.Parts[i])
		if err != nil {return nil, errors.New("Failed to get Part")}
		
		var sp Part
		json.Unmarshal(spAsBytes, &sp)
		if(sp.Helicopter == key) {rap.Parts = append(rap.Parts,sp); } 
		//else { rap.Parts = append(rap.Parts," Cet Helicopter n'est pas composé de parts ")
		//}
	}
	rabAsBytes, _ := json.Marshal(rap)
	return rabAsBytes, nil
}
// ===========================================================================================================================================
// Assembly
// ============================================================================================================================================
// ====================================================
// Afficher toutes les Assembly  
//=====================================================
func (t *SimpleChaincode) getAllAssemblies(stub shim.ChaincodeStubInterface, args []string)([]byte, error){
	
	fmt.Println("Start find getAllAssemblies ")
	fmt.Println("Looking for All Responsibles ")

	//get the getAllAssemblies index
	allRAsBytes, err := stub.GetState("allAssemblies")
	if err != nil {return nil, errors.New("Failed to get all Responsibles")}

	var res AllAssemblies
	err = json.Unmarshal(allRAsBytes, &res)
	if err != nil {return nil, errors.New("Failed to Unmarshal all Responsibles")}

	var rap AllAssemblies

	for i := range res.Assemblies{

		srAsBytes, err := stub.GetState(res.Assemblies[i])
		if err != nil {return nil, errors.New("Failed to get Assembly")}
		var sr Assembly
		json.Unmarshal(srAsBytes, &sr)
		
		rap.Assemblies = append(rap.Assemblies,sr.AssemblyName); 
	}
	rapAsBytes, _ := json.Marshal(rap)
	return rapAsBytes, nil
}
// ========================================================
// Afficher toutes les parts d'un Assembly
//=========================================================
func (t *SimpleChaincode) getAllPartsOfThisAssembly(stub shim.ChaincodeStubInterface, args [] string)([]byte, error){
	fmt.Println("Start find getAllParts ")
	fmt.Println("Looking for All Parts ")
	
	var key string 
	key = args[0]

	//get the AllParts index
	allPAsBytes, err := stub.GetState("allParts")
	if err != nil {return nil, errors.New("Failed to get all Parts")}

	var res AllParts
	err = json.Unmarshal(allPAsBytes, &res)
	if err != nil {return nil, errors.New("Failed to Unmarshal all Parts")}

	var rap AllParts

	for i := range res.Parts{
		spAsBytes, err := stub.GetState(res.Parts[i])
		if err != nil {return nil, errors.New("Failed to get Part")}
		
		var sp Part
		json.Unmarshal(spAsBytes, &sp)
		if(sp.Assembly == key) { rap.Parts = append(rap.Parts,sp.Id); 
		} else { rap.Parts = append(rap.Parts," Cet assembly n'est pas composé de parts")
		}
	}
	rabAsBytes, _ := json.Marshal(rap)
	return rabAsBytes, nil
}
// ==========================================================
// Afficher toutes les parts d'un Assembly en détail
//===========================================================
func (t *SimpleChaincode) getAllPartsOfThisAssemblyInDetails(stub shim.ChaincodeStubInterface, args [] string)([]byte, error){
	fmt.Println("Start find getAllParts ")
	fmt.Println("Looking for All Parts ")
	
	var key string 
	key = args[0]

	//get the AllParts index
	allPAsBytes, err := stub.GetState("allParts")
	if err != nil {return nil, errors.New("Failed to get all Parts")}

	var res AllParts
	err = json.Unmarshal(allPAsBytes, &res)
	if err != nil {return nil, errors.New("Failed to Unmarshal all Parts")}

	var rap AllPartsDetails

	for i := range res.Parts{
		spAsBytes, err := stub.GetState(res.Parts[i])
		if err != nil {return nil, errors.New("Failed to get Part")}
		
		var sp Part
		json.Unmarshal(spAsBytes, &sp)
		if(sp.Assembly == key) { rap.Parts = append(rap.Parts,sp); 
		} 
		//else { rap.Parts = append(rap.Parts,"Cet assembly n'est pas composé de parts")
		//}
	}
	rabAsBytes, _ := json.Marshal(rap)
	return rabAsBytes, nil
}
// ==============================================================================================================================
// ACTIVITIES 
// ============================================================================================================================
// ============================================================
// Transfert de propriété 
// ============================================================
func (t *SimpleChaincode) ownershipTransfer(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {
	var err error
	var key string 
	key = args[0]
	
	n:= checkOwnership(stub, args[0])
	if n != nil {
		fmt.Println(n.Error())
		return nil, errors.New(n.Error())}	
	
	
	//Update Part owner
	valAsbytes, err := stub.GetState(key)
	if err != nil {return nil, errors.New("Failed to get part #" + key)}
	var pt Part
	err = json.Unmarshal(valAsbytes, &pt)
	if err != nil {return nil, errors.New("Failed to Unmarshal Part #" + key)}
	pt.Owner = args[1]
	var tx Log
	tx.Owner 		= pt.Owner
	tx.Responsible  = pt.Responsible
	tx.Helicopter   = pt.Helicopter
	tx.Assembly     = pt.Assembly
	tx.VDate 		= args[2]
	tx.LType 		= "OWNERNSHIP_TRANSFER"
	pt.Logs = append(pt.Logs, tx)
	//Commit updates part to ledger
	ptAsBytes, _ := json.Marshal(pt)
	err = stub.PutState(pt.Id, ptAsBytes)	
	if err != nil {return nil, err}

	z:= createOwner(stub, args[1], /*args[0],*/ pt)
	if z != nil {
		fmt.Println(z.Error())
		return nil, errors.New(z.Error())
	}
	return nil, nil
}
// =========================================================
// Transfert de responsabilité -
// =========================================================
func (t *SimpleChaincode) responsibilityTransfer(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {
	var err error
	var key string 
	key = args[0]
	
	//Update Part Responsible
	valAsbytes, err := stub.GetState(key)
	if err != nil {return nil, errors.New("Failed to get part #" + key)}
	
	var pt Part
	err = json.Unmarshal(valAsbytes, &pt)
	if err != nil {return nil, errors.New("Failed to Unmarshal Part #" + key)}
	
	pt.Responsible = args[1]
	var tx Log
	tx.Responsible	= pt.Responsible
	tx.Owner 		= pt.Owner
	tx.Helicopter   = pt.Helicopter
	tx.Assembly     = pt.Assembly
	tx.VDate 		= args[2]
	tx.LType 		= "RESPONSIBILITY_TRANSFER"
	pt.Logs = append(pt.Logs, tx)

	//Commit updates batch to ledger
	ptAsBytes, _ := json.Marshal(pt)
	err = stub.PutState(pt.Id, ptAsBytes)	
	
	
	if err != nil {return nil, err}
	
	y:= createResponsible(stub, args[1])
	if y != nil {
		fmt.Println(y.Error())
		return nil, errors.New(y.Error())
	}
	return nil, nil
}
// ========================================================
// Acitivités sur la part - 
// ========================================================
func (t *SimpleChaincode) performActivities(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {

	var err error
	var key string 
	key = args[0]
	
	//Update Part Responsible
	valAsbytes, err := stub.GetState(key)
	if err != nil {return nil, errors.New("Failed to get part #" + key)}

	var pt Part
	err = json.Unmarshal(valAsbytes, &pt)
	if err != nil {return nil, errors.New("Failed to Unmarshal Part #" + key)}
	
	var tx Log
	tx.Responsible	= pt.Responsible
	tx.Owner 		= pt.Owner
	tx.ModType      = args[1]
	
	if (args[1] == "Monte" || args[1] == "Demonte" || args[1] == "Scrapping" || args[1] == "SB" ) { // A COMPLETER AVEC INFOS COMPLETE
	tx.Description  = args[2]
	tx.Helicopter   = pt.Helicopter
	tx.Assembly     = pt.Assembly
	tx.VDate 		= args[3]
	tx.LType 		= "ACTIVITIES"
	pt.Logs = append(pt.Logs, tx)
	
		//Commit updates batch to ledger
	
	ptAsBytes, _ := json.Marshal(pt)
	err = stub.PutState(pt.Id, ptAsBytes)	
	
	if err != nil {return nil, err}

	}
	return nil, nil
}
// ============================================================
// Transfert d'Assembly  
// ============================================================
func (t *SimpleChaincode) assemblyTransfer(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {
	var err error
	var key string 
	key = args[0]
	
	//Update Part Responsible
	valAsbytes, err := stub.GetState(key)
	if err != nil {return nil, errors.New("Failed to get part #" + key)}
	
	var pt Part
	err = json.Unmarshal(valAsbytes, &pt)
	if err != nil {return nil, errors.New("Failed to Unmarshal Part #" + key)}
	
	pt.Assembly = args[1]
	var tx Log
	tx.Responsible	= pt.Responsible
	tx.Owner 		= pt.Owner
	tx.Helicopter   = pt.Helicopter
	tx.Assembly     = pt.Assembly
	tx.VDate 		= args[2]
	tx.LType 		= "ASSEMBLY_TRANSFER"
	pt.Logs = append(pt.Logs, tx)

	//Commit updates batch to ledger
	ptAsBytes, _ := json.Marshal(pt)
	err = stub.PutState(pt.Id, ptAsBytes)	
	
	if err != nil {return nil, err}
	
	y:= createAssembly(stub, args[1])
	if y != nil {
		fmt.Println(y.Error())
		return nil, errors.New(y.Error())
	}
	return nil, nil
}
// ============================================================
// Transfert d'Helicoptère  
// ============================================================
func (t *SimpleChaincode) helicoTransfer(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {

	var err error
	var key string 
	key = args[0]
	
	//Update Part Responsible
	valAsbytes, err := stub.GetState(key)
	if err != nil {return nil, errors.New("Failed to get part #" + key)}
	
	var pt Part
	err = json.Unmarshal(valAsbytes, &pt)
	if err != nil {return nil, errors.New("Failed to Unmarshal Part #" + key)}
	
	pt.Helicopter = args[1]
	var tx Log
	tx.Responsible	= pt.Responsible
	tx.Owner 		= pt.Owner
	tx.Helicopter   = pt.Helicopter
	tx.Assembly     = pt.Assembly
	tx.VDate 		= args[2]
	tx.LType 		= "HELICO_TRANSFER"
	pt.Logs = append(pt.Logs, tx)

	//Commit updates batch to ledger
	ptAsBytes, _ := json.Marshal(pt)
	err = stub.PutState(pt.Id, ptAsBytes)	
	
	if err != nil {return nil, err}
	
	y:= createAircraft(stub, args[1])
	if y != nil {
		fmt.Println(y.Error())
		return nil, errors.New(y.Error())
	}
	return nil, nil
}
// ================================================================================================================================
// UTILITY FUNCTIONS
// ===============================================================================================================================

//=====================
//	 Get Attributes 
//=====================
func getAttribute(stub shim.ChaincodeStubInterface, attributeName string) (string, error) {
	bytes, err := stub.ReadCertAttribute(attributeName)
	return string(bytes[:]), err
}
// ========================================================================
// Creation of the Owner   
// ========================================================================
func createOwner(stub shim.ChaincodeStubInterface, args string /*, args1 string */, args2 Part) error {
	fmt.Println("Running createOwner")	
	var err error
	//owner 
	
	var own Owner
	own.OwnerName = args
	// own.PartId = args1
	own.Parts = append(own.Parts, args2)
	own.Statut = "Current"
	
	var  jsonResp3 string
	
	owneAsBytes, err := stub.GetState(own.OwnerName)
	if err != nil {
		jsonResp3 = "{\"Error\":\"Failed to get state for " + own.OwnerName + "\"}"
		errors.New(jsonResp3)
				return nil
	}
	 if owneAsBytes != nil {
		
		y:= ownerHistoric (stub, args, args2)
		if y != nil {
		fmt.Println(y.Error())
		return nil // , errors.New(y.Error())
		}
		return  nil
	
	} else {	
	
	//Commit owner to ledger
	ownAsBytes, _ := json.Marshal(own)
	err = stub.PutState(own.OwnerName, ownAsBytes)	
	if err != nil {return err}
		
	// All Owners - Update AllOwners Array
	allOAsBytes, err := stub.GetState("allOwners")
	if err != nil {return errors.New("Failed to get all Owners")}
	var allown AllOwners
	err = json.Unmarshal(allOAsBytes, &allown)
	if err != nil {return errors.New("Failed to Unmarshal all Owners")}
	allown.Owners = append(allown.Owners,own.OwnerName)
	
	//Commit AllOwners to ledger	
	allOuAsBytes, _ := json.Marshal(allown)
	err = stub.PutState("allOwners", allOuAsBytes)	
	if err != nil {return  err}
	}
	
	fmt.Println("Owner created successfully")	
	return nil
}
// ===========================================================================
// Creation of the Responsible  
// ===========================================================================
func createResponsible(stub shim.ChaincodeStubInterface, args string) error {
	fmt.Println("Running createResponsible")
	var err error
	
	// Responsible 
	var resp Responsible 
	resp.ResponsibleName = args
	
	var  jsonResp3, jsonResp4 string
	
	owneAsBytes, err := stub.GetState(resp.ResponsibleName)
	if err != nil {
		jsonResp3 = "{\"Error\":\"Failed to get state for " + resp.ResponsibleName + "\"}"
		errors.New(jsonResp3)
				return nil
	}
	 if owneAsBytes != nil {
		jsonResp4 = "{\"Error\":\"The following Owner  Already exists, " + resp.ResponsibleName + "\"}"
		errors.New(jsonResp4)
		return  nil
	} else {	
	//Commit responsible to ledger
	respAsBytes, _ := json.Marshal(resp)
	err = stub.PutState(resp.ResponsibleName, respAsBytes)	
	if err != nil {return err}
	
	// AllResponsibles- Update AllResponsibles Array
	allRAsBytes, err := stub.GetState("allResponsibles")
	if err != nil {return errors.New("Failed to get all Responsibles")}
	var allresp AllResponsibles
	err = json.Unmarshal(allRAsBytes, &allresp)
	if err != nil {return errors.New("Failed to Unmarshal all Responsibles")}
	allresp.Responsibles = append(allresp.Responsibles,resp.ResponsibleName)
	
	//Commit AllResponsibles to ledger	
	allRuAsBytes, _ := json.Marshal(allresp)
	err = stub.PutState("allResponsibles", allRuAsBytes)	
	if err != nil {return err}	
}

fmt.Println("Responsible created successfully")	
	return nil
}
// =========================================================================
// Creation of the Aircraft  
// =========================================================================
func createAircraft(stub shim.ChaincodeStubInterface, args string) error {
	fmt.Println("Running createAircraft")
	var err error
	
	// Responsible 
	var resp Aircraft 
	resp.AircraftName = args
	
	var  jsonResp3, jsonResp4 string
	
	owneAsBytes, err := stub.GetState(resp.AircraftName)
	if err != nil {
		jsonResp3 = "{\"Error\":\"Failed to get state for " + resp.AircraftName + "\"}"
		errors.New(jsonResp3)
				return nil
	}
	 if owneAsBytes != nil {
		jsonResp4 = "{\"Error\":\"The following Owner  Already exists, " + resp.AircraftName + "\"}"
		errors.New(jsonResp4)
		return  nil
	} else {	
	//Commit responsible to ledger
	respAsBytes, _ := json.Marshal(resp)
	err = stub.PutState(resp.AircraftName, respAsBytes)	
	if err != nil {return err}
	
	// AllResponsibles- Update AllResponsibles Array
	allRAsBytes, err := stub.GetState("allAircrafts")
	if err != nil {return errors.New("Failed to get all Responsibles")}
	var allresp AllAircrafts
	err = json.Unmarshal(allRAsBytes, &allresp)
	if err != nil {return errors.New("Failed to Unmarshal all Responsibles")}
	allresp.Aircrafts = append(allresp.Aircrafts,resp.AircraftName)
	
	//Commit AllResponsibles to ledger	
	allRuAsBytes, _ := json.Marshal(allresp)
	err = stub.PutState("allAircrafts", allRuAsBytes)	
	if err != nil {return err}	
}

fmt.Println("Responsible created successfully")	
	return nil
}
// =======================================================================
// Creation of the Assembly  
// =======================================================================
func createAssembly(stub shim.ChaincodeStubInterface, args string) error {
	fmt.Println("Running createAssembly")
	var err error
	
	// Responsible 
	var resp Assembly
	resp.AssemblyName = args
	
	var  jsonResp3, jsonResp4 string
	
	owneAsBytes, err := stub.GetState(resp.AssemblyName)
	if err != nil {
		jsonResp3 = "{\"Error\":\"Failed to get state for " + resp.AssemblyName + "\"}"
		errors.New(jsonResp3)
				return nil
	}
	 if owneAsBytes != nil {
		jsonResp4 = "{\"Error\":\"The following Owner  Already exists, " + resp.AssemblyName + "\"}"
		errors.New(jsonResp4)
		return  nil
	} else {	
	
	//Commit responsible to ledger
	respAsBytes, _ := json.Marshal(resp)
	err = stub.PutState(resp.AssemblyName, respAsBytes)	
	if err != nil {return err}
	
	// AllResponsibles- Update AllResponsibles Array
	allRAsBytes, err := stub.GetState("allAssemblies")
	if err != nil {return errors.New("Failed to get all Responsibles")}
	
	var allresp AllAssemblies
	err = json.Unmarshal(allRAsBytes, &allresp)
	if err != nil {return errors.New("Failed to Unmarshal all Responsibles")}
	allresp.Assemblies = append(allresp.Assemblies,resp.AssemblyName)
	
	//Commit AllResponsibles to ledger	
	allRuAsBytes, _ := json.Marshal(allresp)
	err = stub.PutState("allAssemblies", allRuAsBytes)	
	if err != nil {return err}	
}

fmt.Println("Responsible created successfully")	
	return nil
}
// ============================================================================
// check Id Availability -
// ============================================================================
func checkIDavailibility(stub shim.ChaincodeStubInterface, args string) error {

	fmt.Println("Running checkIDavailibility")
	
	var err error
	var key, jsonResp, jsonResp2 string
	key = args
	partAsBytes, err := stub.GetState(args)
	if err != nil {
		jsonResp = "{\"Error\":\"Failed to get state for " + key + "\"}"
		return  errors.New(jsonResp)
	}
	 if partAsBytes != nil {
		jsonResp2 = "{\"Error\":\"The following ID is Already taken, " + key + "\"}"
		return  errors.New(jsonResp2)	
	}
	fmt.Println("ID checked successfully")	
	return nil
}
// =============================================================================
// check PN Availability -
// =============================================================================
func checkPNavailibility(stub shim.ChaincodeStubInterface, args string) error {

	fmt.Println("Running checkIDavailibility")
	
	var err error
	var key, jsonResp, jsonResp2 string
	key = args
	partAsBytes, err := stub.GetState(args)
	if err != nil {
		jsonResp = "{\"Error\":\"Failed to get state for " + key + "\"}"
		return  errors.New(jsonResp)
	}
	 if partAsBytes != nil {
		jsonResp2 = "{\"Error\":\"The following PN is Already taken, " + key + "\"}"
		return  errors.New(jsonResp2)
	}
	fmt.Println("PN checked successfully")	
	return nil
}
// =============================================================================
// check SN Availability -
// =============================================================================
func checkSNavailibility(stub shim.ChaincodeStubInterface, args string) error {

	fmt.Println("Running checkIDavailibility")
	
	var err error
	var key, jsonResp, jsonResp2 string
	key = args
	partAsBytes, err := stub.GetState(args)
	if err != nil {
		jsonResp = "{\"Error\":\"Failed to get state for " + key + "\"}"
		return  errors.New(jsonResp)
	}
	 if partAsBytes != nil {
		jsonResp2 = "{\"Error\":\"The following SN is Already taken, " + key + "\"}"
		return  errors.New(jsonResp2)
	}
	fmt.Println("SN checked successfully")	
	return nil
}
// ==================================================================================
// check Ownership On Part -
// ==================================================================================
func checkOwnership(stub shim.ChaincodeStubInterface, args string) error {

	fmt.Println("Running CheckOwnership")
	
	var err error
	var key string // Id de la part 
	key = args 
	var jsonResp2 string

	username, err := getAttribute(stub, "username")
	partAsBytes, err := stub.GetState(key)
	if err != nil {return  errors.New("Failed to get Parts")}
	var pt Part
	err = json.Unmarshal(partAsBytes, &pt)
	
	if ( username != pt.Owner) { 
		jsonResp2 = "{\"Error\":\"You are not owner of this part, " + key + "\"}"
		return  errors.New(jsonResp2)
	}
	fmt.Println("PN checked successfully")	
	return nil
}
// =============================
// Display Owner's Historic Part Listing 
//==============================
func ownerHistoric(stub shim.ChaincodeStubInterface, args string, args1 Part) error {

	fmt.Println("Running ownerHistoric")
	
	var err error
	var key string
	key = args
	
	partAsBytes, err := stub.GetState(key)
	
	var own Owner
	err = json.Unmarshal(partAsBytes, &own)
	if err != nil {return nil } //,errors.New("Failed to Unmarshal Owner #" + key )}
	
	own.Parts = append(own.Parts, args1)
	
	ptAsBytes, _ := json.Marshal(own)
	err = stub.PutState(own.OwnerName, ptAsBytes)
	
	if err != nil {return nil } //, err}
	return nil
}

// ================================================================================================================================
// Test Functions 
// ================================================================================================================================
// =================================================
// hello - query function to read key/value pair
// =================================================
func (t *SimpleChaincode) hello(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {
	return []byte("Welcolme to the eLogcard Application") , nil
}


//=================================================================================================================================
//	 Main - main - Starts up the chaincode  
//=================================================================================================================================
func main() {
	fmt.Println("Welcome to eLogcard System!")
	err := shim.Start(new(SimpleChaincode))
	if err != nil {
		fmt.Printf("Error starting Simple chaincode: %s", err)
	}
}

// pour l'historique: l'owner est composé de part donc a chaque création on ajoute la part à son champ []parts 
// pour le statut: 




	

