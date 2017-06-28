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
		tx.LType 		= "CREATION"
		tx.Description  = args[5] + " created this Part "
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

	e:= UpdatePart (stub, pt) 
		if e != nil { return nil, errors.New(e.Error())}
		
	return []byte("eLogcardlogcard created successfully"),err
	fmt.Println("eLogcardlogcard created successfully")	
	return nil, nil
}

// =========================================================================================
// 					ACTIVITIES ON PARTS
// =========================================================================================

// =========================
// Maintenance 
// =========================
// Vérifier Respo
func (t *SimpleChaincode) performActivities(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {

	var err error
	var key string 
	key = args[0]
	
	
	username, err := getAttribute(stub, "username")
		if(err !=nil){return nil,err}
	role, err := getAttribute(stub, "role")
		if(err !=nil){return nil,err}
	//if supplier or manufacturer or customer or maintenance user =>only my parts
	showOnlyMyPart := role=="supplier" || role == "manufacturer" || role == "customer" || role == "maintenance_user"

	
	part,err:=findPartById(stub,key)
		if err != nil {return nil, errors.New("Failed to get part #" + key)}
		ptAS, _ := json.Marshal(part)
	var pt Part
		err = json.Unmarshal(ptAS, &pt)
		if err != nil {return nil, errors.New("Failed to Unmarshal Part #" + key)}
	if (showOnlyMyPart && pt.Id == key && pt.Responsible == username) {	
	var tx Log
		tx.Owner 	= pt.Owner
		tx.Responsible 	= pt.Responsible
		tx.Description = args[2]
		tx.VDate 		= args[3]
		tx.LType 		= "ACTIVITIES_PERFORMED: " + args [1]
		pt.Logs = append(pt.Logs, tx)
		
		e:= UpdatePart (stub, pt) 
		if e != nil { return nil, errors.New(e.Error())}
		
		}
return nil, nil
}
// =========================
// Transfert de propriété 
// =========================
// Only registered suppliers, manufacturers, customers and maintenance_user can Transfer Ownership on a Part.
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
		tx.Responsible  = pt.Responsible
		tx.VDate 		= args[2]
		tx.LType 		= "OWNERNSHIP_TRANSFER"
		tx.Description  = "This part has been transfered to " + pt.Owner + ", the new Owner" 
		
	pt.Logs = append(pt.Logs, tx)
	
	e:= UpdatePart (stub, pt) 
		if e != nil { return nil, errors.New(e.Error())}
return nil, nil
}
// =============================
// Transfert de responsabilité 
// =============================
// A FAIRE Vérifier Respo

// Only registered suppliers, manufacturers, customers and maintenance_user can Transfer Responsibility on a Part.
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
		tx.Owner        = pt.Owner
		tx.VDate 		= args[2]
		tx.LType 		= "RESPONSIBILITY_TRANSFER"
		tx.Description  = "This part has been transfered to " + pt.Responsible + ", the new Responsible" 

		pt.Logs = append(pt.Logs, tx)

		
	e:= UpdatePart (stub, pt) 
		if e != nil { return nil, errors.New(e.Error())}

return nil, nil 
}

// =========================
// Scrapp a Part  
// =========================
func (t *SimpleChaincode) scrappPart(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {

	var err error
	var key string 
	key = args[0]
	part,err:=findPartById(stub,key)
		if err != nil {return nil, errors.New("Failed to get part #" + key)}
		ptAS, _ := json.Marshal(part)
	var pt Part
		err = json.Unmarshal(ptAS, &pt)
		if err != nil {return nil, errors.New("Failed to Unmarshal Part #" + key)}
		pt.Owner = "SCRAPPING_MANAGER"
		pt.Responsible = "SCRAPPING_MANAGER"
		pt.PN = ""
		pt.Helicopter = ""
		pt.Assembly = ""
	var tx Log
		tx.Owner 		= pt.Owner
		tx.Responsible 	= pt.Responsible
		tx.VDate 		=  args [1]
		tx.LType 		= "SCRAPPING"
		tx.Description  = "This part has been  scrapped and transfered to " + pt.Responsible + ", the new Owner & the new Responsible" 

	pt.Logs = append(pt.Logs, tx)
	
	e:= UpdatePart (stub, pt) 
		if e != nil { return nil, errors.New(e.Error())}

		
return nil, nil

}
// ====================================================================
// GET
// ====================================================================
// ====================================================================
// Obtenir tous les détails d'une part à partir de son id 
// Registered suppliers, manufacturers, customers and maintenance users can  display details on a specific part only if they own it.
// Auditor_authority and AH_Admin can see details on any specific part they want.
// ====================================================================
func (t *SimpleChaincode) getPartDetails(stub shim.ChaincodeStubInterface, args []string)([]byte, error) {
 
 // Vérifier Respo

	key :=  args[0]

	username, err := getAttribute(stub, "username")
		if(err !=nil){return nil,err}
	role, err := getAttribute(stub, "role")
		if(err !=nil){return nil,err}
	//if supplier or manufacturer or customer or maintenance user =>only my parts
	showOnlyMyPart := role=="supplier" || role == "manufacturer" || role == "customer" || role == "maintenance_user"

	part,err:=findPartById(stub,key)
		if(err !=nil){return nil,err}
	ptAS, _ := json.Marshal(part)
	var pt Part
		err = json.Unmarshal(ptAS, &pt)
		if err != nil {return nil, errors.New("Failed to Unmarshal Part #" + key)}
		if (showOnlyMyPart && pt.Id == key && pt.Owner == username) {
			return json.Marshal(part)  
		} else if (!showOnlyMyPart && pt.Id == key) {
			return json.Marshal(part)  }
	return nil, nil 
}
// ==================================================================
// Afficher toutes les parts créées en détail  
// Registered suppliers, manufacturers, customers and maintenance users can display details of all the parts they own.
// Auditor_authority and AH_Admin can display details of all the parts ever created.
//===================================================================
func (t *SimpleChaincode) getAllPartsDetails(stub shim.ChaincodeStubInterface, args []string)([]byte, error){
	
// A FAIRE: Vérifier Respo
	username, err := getAttribute(stub, "username")
		if(err !=nil){return nil,err}
	role, err := getAttribute(stub, "role")
		if(err !=nil){return nil,err}
	//if supplier or manufacturer or customer or maintenance user =>only my parts
	showOnlyMyPart := role=="supplier" || role == "manufacturer" || role == "customer" || role == "maintenance_user"
	
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
	
	return nil, nil 
}


// ==================================================================
// Afficher toutes les parts créées en détail  
// Registered suppliers, manufacturers, customers and maintenance users can display details of all the parts they own.
// Auditor_authority and AH_Admin can display details of all the parts ever created.
//===================================================================
func (t *SimpleChaincode) getAllPartsWithoutAssembly(stub shim.ChaincodeStubInterface, args []string)([]byte, error){
// A FAIRE: Vérifier Respo
	username, err := getAttribute(stub, "username")
		if(err !=nil){return nil,err}
	role, err := getAttribute(stub, "role")
		if(err !=nil){return nil,err}
	//if supplier or manufacturer or customer or maintenance user =>only my parts
	showOnlyMyPart := role=="supplier" || role == "manufacturer" || role == "customer" || role == "maintenance_user"
	
	partMap,err:=getPartsIdMap(stub)
		if(err !=nil){return nil,err}
	parts := make([]Part, len(partMap))
    idx := 0
    for  _, part := range partMap {
    	if(showOnlyMyPart &&  part.Owner == username && part.Assembly ==""){
    		parts[idx] = part
    		idx++
    	}
    }
    //si les deux longueurs sont differentes on slice
    if(showOnlyMyPart && len(partMap)!=idx){
    	parts=parts[0:idx]
    }
    return json.Marshal(parts) 
	
	return nil, nil 
}


func (t *SimpleChaincode) getAllPartsWithoutAircraft(stub shim.ChaincodeStubInterface, args []string)([]byte, error){
// A FAIRE: Vérifier Respo
	username, err := getAttribute(stub, "username")
		if(err !=nil){return nil,err}
	role, err := getAttribute(stub, "role")
		if(err !=nil){return nil,err}
	//if supplier or manufacturer or customer or maintenance user =>only my parts
	showOnlyMyPart := role=="supplier" || role == "manufacturer" || role == "customer" || role == "maintenance_user"
	
	partMap,err:=getPartsIdMap(stub)
		if(err !=nil){return nil,err}
	parts := make([]Part, len(partMap))
    idx := 0
    for  _, part := range partMap {
    	if(showOnlyMyPart &&  part.Owner == username && part.Helicopter == ""){
    		parts[idx] = part
    		idx++
    	}
    }
    //si les deux longueurs sont differentes on slice
    if(showOnlyMyPart && len(partMap)!=idx){
    	parts=parts[0:idx]
    }
    return json.Marshal(parts) 
	
	return nil, nil 
}
