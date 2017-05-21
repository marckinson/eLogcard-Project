//=============================================================================================================
//	 				e-LogCard CHAINCODE
//=============================================================================================================
package main
import (
	"errors"
	"fmt"
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