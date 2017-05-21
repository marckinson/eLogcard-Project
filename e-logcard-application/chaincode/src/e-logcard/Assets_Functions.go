// ============================================================================================================
// 					e-LogCard FUNCTIONS
// ============================================================================================================
package main
import (
	"errors"
	"fmt"
	"encoding/json"
	"github.com/hyperledger/fabric/core/chaincode/shim"
)
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
 
typ := args[0]
key :=  args[1]

	username, err := getAttribute(stub, "username")
		if(err !=nil){return nil,err}
	role, err := getAttribute(stub, "role")
		if(err !=nil){return nil,err}
	//if supplier or manufacturer or customer or maintenance user =>only my parts
	showOnlyMyPart := role=="supplier" || role == "manufacturer" || role == "Customer" || role == "maintenance_user"

if ( typ == "Part") {
	part,err:=findPartById(stub,key)
	if(err !=nil){return nil,err}
	ptAS, _ := json.Marshal(part)
	var pt Part
	err = json.Unmarshal(ptAS, &pt)
	if err != nil {return nil, errors.New("Failed to Unmarshal Part #" + key)}
	if (showOnlyMyPart && pt.Id == key && pt.Owner == username) {
	return json.Marshal(part)  }
	
} else if ( typ == "Aircraft"){	
	partMap,err:=getPartsIdMap(stub)
	if err != nil {return nil, errors.New("Failed to get Part")}
	parts := make([]Part, len(partMap))
	idx := 0
    for  _, part := range partMap {
    	if(!showOnlyMyPart || part.Helicopter == key || part.Owner == username){
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
		if (!showOnlyMyPart || part.Assembly == key || part.Owner == username){
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
	
	username, err := getAttribute(stub, "username")
		if(err !=nil){return nil,err}
	role, err := getAttribute(stub, "role")
		if(err !=nil){return nil,err}
	//if supplier or manufacturer or customer or maintenance user =>only my parts
	showOnlyMyPart := role=="supplier" || role == "manufacturer" || role == "Customer" || role == "maintenance_user"
	
if ( typ == "Part") {
// Parts
	fmt.Println("Start find getAllPartsDetails ")
	fmt.Println("Looking for All Parts With Details ")
	
	partMap,err:=getPartsIdMap(stub)
		if(err !=nil){return nil,err}

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
		if(!showOnlyMyPart || part.Owner == username){
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
// Assemblies
	fmt.Println("Start find getAllAircraftDetails ")
	fmt.Println("Looking for All Aircrafts With Details ")
	
	partMap,err:=getAssemblyMap(stub)
		if(err !=nil){return nil,err}

	parts := make([]Assembly, len(partMap))
    idx := 0
    for  _, part := range partMap {
		if(!showOnlyMyPart || part.Owner == username){
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
