//=============================================================================================================
//	 											e-LogCard CHAINCODE
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
	LType string `json:"log_type"` // Type of change
	VDate string `json:"vDate"` // Date 
	Owner string `json:"owner"` // Owner of the part
	Responsible string `json:"responsible"` // Responsible of the part at the moment 
	ModType string `json:"modType"` // Type de modifications 
	Description string `json:"description"` // Description de la modification apportée 	
}
// ============================================================================================================
// 												HYPERLEDGER FUNCTIONS
// ============================================================================================================
//============================================================
//	Init Function - Called when the user deploys the chaincode 
//============================================================
func (t *SimpleChaincode) Init(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {
	var err error
	var parts AllParts
	jsonAsBytes, _ := json.Marshal(parts)
	err = stub.PutState("allParts", jsonAsBytes)
	if err != nil {return nil, err}
	return nil, nil
}
// ========================================================
// Invoke is our entry point to invoke a chaincode function
// ========================================================
func (t *SimpleChaincode) Invoke(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {
    fmt.Println("invoke is running " + function)
   
// Users functions 
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
		return t.performActivities(stub, args)}
		
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
		if len(args) != 1 {
		fmt.Println("Incorrect number of arguments. Expecting 1")
		return nil, errors.New("Incorrect number of arguments. Expecting 1: ID")}
		return t.getPartDetails (stub,args)}
	if function == "getAllPartsDetails" {
		return t.getAllPartsDetails (stub,args)}
		
		fmt.Println("query did not find func: " + function)
		return nil, errors.New("Received unknown function query")
}
// ============================================================================================================
// 												PARTS
// ============================================================================================================
// ===================================================================
// Creation of the Part (creation of the eLogcard) 
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
	
// Check 
	n:= checkPNavailibility(stub, args[0])
	if n != nil { fmt.Println(n.Error()); return nil, errors.New(n.Error())}	
	o:= checkSNavailibility(stub, args[1])
	if o != nil {fmt.Println(o.Error()); return nil, errors.New(o.Error())}

//Commit part to ledger
	ptAsBytes, _ := json.Marshal(pt)
	err = stub.PutState(pt.Id, ptAsBytes)	
	err = stub.PutState(pt.PN, ptAsBytes)	
	err = stub.PutState(pt.SN, ptAsBytes)
	if err != nil {return nil, err}
		
// Update AllParts Array
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
// ==================================================================
// Afficher toutes les parts créées en détail  
//===================================================================
func (t *SimpleChaincode) getAllPartsDetails(stub shim.ChaincodeStubInterface, args []string)([]byte, error){
	
	fmt.Println("Start find getAllPartsDetails ")
	fmt.Println("Looking for All Parts With Details ")
	
//Get the AllParts index
	allPAsBytes, err := stub.GetState("allParts")
	if err != nil {return nil, errors.New("Failed to get all Parts")}
	var res AllParts
	err = json.Unmarshal(allPAsBytes, &res)
	if err != nil {return nil, errors.New("Failed to Unmarshal all Parts")}
	
// Display all the parts 
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
// ============================================================================================================
// 												ACTIVITIES 
// ============================================================================================================
// =========================
// Transfert de propriété 
// =========================
func (t *SimpleChaincode) ownershipTransfer(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {
	var err error
	var key string 
	key = args[0]
		
//Update Part owner
	valAsbytes, err := stub.GetState(key)
	if err != nil {return nil, errors.New("Failed to get part #" + key)}
	var pt Part
	err = json.Unmarshal(valAsbytes, &pt)
	if err != nil {return nil, errors.New("Failed to Unmarshal Part #" + key)}
	pt.Owner = args[1]
	var tx Log
	tx.Owner 		= pt.Owner
	tx.VDate 		= args[2]
	tx.LType 		= "OWNERNSHIP_TRANSFER"
	pt.Logs = append(pt.Logs, tx)
	
//Commit updates part to ledger
	ptAsBytes, _ := json.Marshal(pt)
	err = stub.PutState(pt.Id, ptAsBytes)	
	if err != nil {return nil, err}
	return nil, nil
}
// =============================
// Transfert de responsabilité 
// =============================
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
	tx.VDate 		= args[2]
	tx.LType 		= "RESPONSIBILITY_TRANSFER"
	pt.Logs = append(pt.Logs, tx)

//Commit updates batch to ledger
	ptAsBytes, _ := json.Marshal(pt)
	err = stub.PutState(pt.Id, ptAsBytes)	
	if err != nil {return nil, err}
	return nil, nil
}

// =========================
// Acitivités sur la part 
// =========================
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
	
// A COMPLETER AVEC INFOS COMPLETE
	if (args[1] == "Monte" || args[1] == "Demonte" || args[1] == "Scrapping" || args[1] == "SB" ) {
		tx.Description  = args[2]
		tx.VDate 		= args[3]
		tx.LType 		= "ACTIVITIES"
		pt.Logs = append(pt.Logs, tx)
	
		//Commit updates part to ledger
		ptAsBytes, _ := json.Marshal(pt)
		err = stub.PutState(pt.Id, ptAsBytes)	
		if err != nil {return nil, err}
	}
	return nil, nil
}
// ============================================================================================================
// 												UTILITY FUNCTIONS
// ============================================================================================================
//=====================
// Get Attributes 
//=====================
func getAttribute(stub shim.ChaincodeStubInterface, attributeName string) (string, error) {
	bytes, err := stub.ReadCertAttribute(attributeName)
	return string(bytes[:]), err
}
// =========================
// Check PN Availability 
// =========================
func checkPNavailibility(stub shim.ChaincodeStubInterface, args string) error {

	fmt.Println("Running checkIDavailibility")
	var err error
	var key, jsonResp, jsonResp2 string
	key = args
	partAsBytes, err := stub.GetState(args)
	if err != nil {
		jsonResp = "{\"Error\":\"Failed to get state for " + key + "\"}"
		return  errors.New(jsonResp)}
	 if partAsBytes != nil {
		jsonResp2 = "{\"Error\":\"The following PN is Already taken, " + key + "\"}"
		return  errors.New(jsonResp2)}
	fmt.Println("PN checked successfully")	
	return nil
}
// ===========================
// Check SN Availability 
// ===========================
func checkSNavailibility(stub shim.ChaincodeStubInterface, args string) error {
	fmt.Println("Running checkSNavailibility")
	var err error
	var key, jsonResp, jsonResp2 string
	key = args
	partAsBytes, err := stub.GetState(args)
	if err != nil {
		jsonResp = "{\"Error\":\"Failed to get state for " + key + "\"}"
		return  errors.New(jsonResp)}
	 if partAsBytes != nil {
		jsonResp2 = "{\"Error\":\"The following SN is Already taken, " + key + "\"}"
		return  errors.New(jsonResp2)}
	fmt.Println("SN checked successfully")	
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