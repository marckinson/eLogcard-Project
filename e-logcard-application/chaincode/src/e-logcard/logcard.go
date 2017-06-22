//=============================================================================================================
//	 				e-LogCard CHAINCODE
//=============================================================================================================
package main
import (
	"errors"
	"fmt"
	"github.com/hyperledger/fabric/core/chaincode/shim"

)

// tout uniformiser 


//=============================================================================================================
//	 Structure Definitions
//=============================================================================================================
//========================================================================================
//	Chaincode - A blank struct for use with Shim 
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
	Responsible string `json:"responsible"` // Responsabilité (Portée par l'organisation) à l'instant t de la pièce 
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
	VDate string `json:"vDate"` // TimeStamp 
	Owner string `json:"owner"` // Owner of the part
	Responsible string `json:"responsible"` // Responsible of the part at the moment 
	Description string `json:"description"` // Description of the modification  	
}
type LogAssembly struct { 
	LType string `json:"log_type"` // Type of change
	VDate string `json:"vDate"` // TimeStamp 
	Owner string `json:"owner"` // Owner of the part
	Description string `json:"description"` // Description of the modification  	
	}	
type LogAircraft struct { 
	LType string `json:"log_type"` // Type of change
	VDate string `json:"vDate"` // TimeStamp 
	Owner string `json:"owner"` // Owner of the part
	Description string `json:"description"` // Description of the modification  	
	}
	
//================================================
// Aircraft
//================================================
type Aircraft struct { 
	AN string `json:"an"` // Part Number
	SN string `json:"sn"` // Serial Number
	Id_Aircraft string `json:"id_aircraft"` // Génération d'un UUID
	// AircraftName string `json:"aircraftName"` 
	Owner string `json:"owner"` // Nom de la Part 
	Parts []string `json:"parts"` // Parts 
	Assemblies [] string `json:"assemblies"` // Parts 
	Logs []LogAircraft `json:"logs"` // Changements sur la part  + Transactions 
}
//================================================
// Assembly 
//================================================ 
type Assembly struct { 
	AN string `json:"an"` // Part Number
	SN string `json:"sn"` // Serial Number
	Id_Assembly string `json:"id_assembly"` // Génération d'un UUID
// 	AssemblyName string `json:"assemblyName"` 
	Helicopter	string `json:"helicopter"` // Aircraft
	Owner string `json:"owner"` // Nom de la Part 
	Parts []string `json:"parts"` // Parts 
	Logs []LogAssembly `json:"logs"` // Changements sur la part  + Transactions 
}
// ============================================================================================================
// 					HYPERLEDGER FUNCTIONS
// ============================================================================================================
//============================================================
//	Init Function - Called when the user deploys the chaincode 
//============================================================
func (t *SimpleChaincode) Init(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {	
		
// Parts		
		n:= createMap(stub, "allParts")
			if n != nil { fmt.Println(n.Error()); return nil, errors.New(n.Error())}
		o:= createMap(stub, "allPartsPn")
			if o != nil { fmt.Println(o.Error()); return nil, errors.New(o.Error())}
		m:= createMap(stub, "allPartsSn")
			if m != nil { fmt.Println(m.Error()); return nil, errors.New(m.Error())}
// Aircrafts
		p:= createMap(stub, "allAircraft")
			if p != nil { fmt.Println(p.Error()); return nil, errors.New(p.Error())}
		a:= createMap(stub, "allAircraftsAn")
			if a != nil { fmt.Println(a.Error()); return nil, errors.New(a.Error())}
		c:= createMap(stub, "allAircraftsSn")
			if c != nil { fmt.Println(c.Error()); return nil, errors.New(c.Error())}
// Assembly		
		q:= createMap(stub, "allAssembly")
			if q != nil { fmt.Println(q.Error()); return nil, errors.New(q.Error())}
		h:= createMap(stub, "allAssembliesAn")
			if h != nil { fmt.Println(h.Error()); return nil, errors.New(h.Error())}
		u:= createMap(stub, "allAssembliesSn")
			if u != nil { fmt.Println(u.Error()); return nil, errors.New(u.Error())}
	return nil, nil
}
// ========================================================
// Invoke is our entry point to invoke a chaincode function
// ========================================================
func (t *SimpleChaincode) Invoke(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {
    fmt.Println("invoke is running " + function)
	
// Parts 
	if function == "createPart" {
		role, err := getAttribute(stub, "role")
		if(role=="supplier" || role == "manufacturer"){ 
		return t.createPart(stub, args)
		}else { return []byte("You are not authorized"),err}}	
	if function == "ownershipTransfer" {
		role, err := getAttribute(stub, "role")
		if(role=="supplier" || role == "manufacturer" || role == "customer" || role == "maintenance_user"){	
		return t.ownershipTransfer(stub, args)
		}else { return []byte("You are not authorized"),err}} 	
	if function == "responsibilityTransfer" {
		role, err := getAttribute(stub, "role")
		if(role=="supplier" || role == "manufacturer" || role == "customer" || role == "maintenance_user"){	
		return t.responsibilityTransfer(stub, args)
		}else { return []byte("You are not authorized"),err}} 	
	if function == "performActivities" {
		role, err := getAttribute(stub, "role")
		if(role=="supplier" || role == "manufacturer" || role == "customer" || role == "maintenance_user"){	
			return t.performActivities(stub, args) 
		} else { return []byte("You are not authorized"),err}}  
	if function == "scrappPart" {
		role, err := getAttribute(stub, "role")
		if(role=="supplier" || role == "manufacturer" || role == "customer" || role == "maintenance_user"){	
			return t.scrappPart(stub, args) 
		} else { return []byte("You are not authorized"),err}}  
// Aircrafts 
	if function == "createAircraft" {
		role, err := getAttribute(stub, "role")
		if(role=="supplier" || role == "manufacturer"){ 
		return t.createAircraft(stub, args)
		}else { return []byte("You are not authorized"),err}}
	if function == "addPartToAc" {
		role, err := getAttribute(stub, "role")
		if(role=="supplier" || role == "manufacturer" || role == "customer" || role == "maintenance_user"){	
			return t.addPartToAc(stub, args) 
		} else { return []byte("You are not authorized"),err}} 
	if function == "RemovePartFromAc" {
		role, err := getAttribute(stub, "role")
		if(role=="supplier" || role == "manufacturer" || role == "customer" || role == "maintenance_user"){	
			return t.RemovePartFromAc(stub, args) 
		} else { return []byte("You are not authorized"),err}} 	
	if function == "AcOwnershipTransfer" {
		role, err := getAttribute(stub, "role")
		if(role=="supplier" || role == "manufacturer" || role == "customer" || role == "maintenance_user"){	
			return t.AcOwnershipTransfer(stub, args) 
		} else { return []byte("You are not authorized"),err}}
	if function == "ReplacePartOnAircraft" {
		role, err := getAttribute(stub, "role")
		if(role=="supplier" || role == "manufacturer" || role == "customer" || role == "maintenance_user"){	
			return t.ReplacePartOnAircraft(stub, args) 
		} else { return []byte("You are not authorized"),err}}	
	if function == "RemoveAssemblyFromAc" {
		role, err := getAttribute(stub, "role")
		if(role=="supplier" || role == "manufacturer" || role == "customer" || role == "maintenance_user"){	
			return t.RemoveAssemblyFromAc(stub, args) 
		} else { return []byte("You are not authorized"),err}}
	if function == "AddAssemblyToAc" {
		role, err := getAttribute(stub, "role")
		if(role=="supplier" || role == "manufacturer" || role == "customer" || role == "maintenance_user"){	
			return t.AddAssemblyToAc(stub, args) 
		} else { return []byte("You are not authorized"),err}}		
	if function == "scrappAircraft" {
		role, err := getAttribute(stub, "role")
		if(role=="supplier" || role == "manufacturer" || role == "customer" || role == "maintenance_user"){	
			return t.scrappAircraft(stub, args) 
		} else { return []byte("You are not authorized"),err}}  	
	if function == "replaceAssemblyOnAC" {
		role, err := getAttribute(stub, "role")
		if(role=="supplier" || role == "manufacturer" || role == "customer" || role == "maintenance_user"){	
			return t.replaceAssemblyOnAC(stub, args) 
		} else { return []byte("You are not authorized"),err}}  		
// Assembly 
	if function == "createAssembly" {
		role, err := getAttribute(stub, "role")
		if(role=="supplier" || role == "manufacturer"){ 
		return t.createAssembly(stub, args)
		}else { return []byte("You are not authorized"),err}}
	if function == "addPartToAssemb" {
		role, err := getAttribute(stub, "role")
		if(role=="supplier" || role == "manufacturer" || role == "customer" || role == "maintenance_user"){	
			return t.addPartToAssemb(stub, args) 
		} else { return []byte("You are not authorized"),err}} 
	if function == "AssembOwnershipTransfer" {
		role, err := getAttribute(stub, "role")
		if(role=="supplier" || role == "manufacturer" || role == "customer" || role == "maintenance_user"){	
			return t.AssembOwnershipTransfer(stub, args) 
		} else { return []byte("You are not authorized"),err}}	
	if function == "RemovePartFromAs" {
		role, err := getAttribute(stub, "role")
		if(role=="supplier" || role == "manufacturer" || role == "customer" || role == "maintenance_user"){	
			return t.RemovePartFromAs(stub, args) 
		} else { return []byte("You are not authorized"),err}}
	if function == "ReplacePartOnAssembly" {
		role, err := getAttribute(stub, "role")
		if(role=="supplier" || role == "manufacturer" || role == "customer" || role == "maintenance_user"){	
			return t.ReplacePartOnAssembly(stub, args) 
		} else { return []byte("You are not authorized"),err}}		
	if function == "scrappAssembly" {
		role, err := getAttribute(stub, "role")
		if(role=="supplier" || role == "manufacturer" || role == "customer" || role == "maintenance_user"){	
			return t.scrappAssembly(stub, args) 
		} else { return []byte("You are not authorized"),err}}  		
	
	fmt.Println("invoke did not find func: " + function)
	return nil, errors.New("Received unknown function invoke")
}
// =========================================================
// Query - read a variable from chaincode state - (aka read)  
// =========================================================
func (t *SimpleChaincode) Query(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {
    fmt.Println("query is running " + function)

// Parts
	if function == "getPartDetails" {
		return t.getPartDetails (stub, args)}
	if function == "getAllPartsDetails" {
		return t.getAllPartsDetails (stub,args)}
	if function == "getAllPartsWithoutAssembly" {
		return t.getAllPartsWithoutAssembly (stub,args)}
	if function == "getAllPartsWithoutAircraft" {
		return t.getAllPartsWithoutAircraft (stub,args)}

// Aircrafts 
	if function == "getAcDetails" {
		return t.getAcDetails (stub, args)}
	if function == "AcPartsListing" {
		return t.AcPartsListing (stub,args)}
	if function == "getAllAircraftsDetails" {
		return t.getAllAircraftsDetails (stub,args)}
	if function == "AcAssembliesListing" {
		return t.AcAssembliesListing (stub,args)}	
		
// Assemblies 
	if function == "getAssembDetails" {
		return t.getAssembDetails (stub, args)}
	if function == "AssembPartsListing" {
		return t.AssembPartsListing (stub, args)}
	if function == "getAllAssembliesDetails" {
		return t.getAllAssembliesDetails (stub, args)}
		
// Lists 
	if function == "getList" {
		return t.getList (stub, args)}
	if function == "getAircraftsList" {
		return t.getAircraftsList (stub, args)}
	if function == "getRolesList" {
		return t.getRolesList (stub, args)}
	if function == "getActionsList" {
		return t.getActionsList (stub, args)}
	if function == "getAircraftTypesList" {
		return t.getAircraftTypesList (stub, args)}
	if function == "getAssembliesList" {
		return t.getAssembliesList (stub, args)}
	if function == "getPartsList" {
		return t.getPartsList (stub, args)}
	if function == "getAircraftTypesList" {
		return t.getAircraftTypesList (stub, args)}
	if function == "getLogsList" {
		return t.getLogsList (stub, args)}
			
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