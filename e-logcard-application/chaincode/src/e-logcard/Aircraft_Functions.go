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
// 					Aircraft
// =========================================================================================
// ===================================================================
// Creation of the Aircraft 
// Only registered suppliers and manufacturers can create Parts.  
// ===================================================================
func (t *SimpleChaincode) createAircraft(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {
	fmt.Println("Running createAircraft")

	var err error
	var air Aircraft
		air.AN = args [0] // Aircraft Number 
		air.SN = args [1] // Serial Number 
		air.Id_Aircraft = args[2] // Id of the Aircraft 
		air.Owner = args [3] // Owner of the Aircraft 
	var tx Log
		tx.Owner 		= air.Owner
		tx.VDate 		= args[4]
		tx.LType 		= "CREATE"
	air.Logs = append(air.Logs, tx)

//Commit part to ledger
	ptAsBytes, _ := json.Marshal(air)
		err = stub.PutState(air.Id_Aircraft, ptAsBytes)
		if err != nil {return nil, err}	
//Fin Commit part to ledger

//Update allAircraft 
		partzMap,err:=getAircraftMap(stub)
		partzMap[air.Id_Aircraft] = air
		allPAsBuytes, err := json.Marshal(partzMap)
		err=stub.PutState("allAircraft",allPAsBuytes)
		if err != nil {return nil, err}
//Fin update allAircraft 
		
fmt.Println("Responsible created successfully")	
return nil, nil
}
// ====================================================================
// addPartToAc 
// ====================================================================
func (t *SimpleChaincode)addPartToAc(stub shim.ChaincodeStubInterface, args []string)([]byte, error) {

	key :=  args[0]
	idpart := args[1]

	// Debut Partie Aircraft 
	ac,err:=findAircraftById(stub,key)
		if(err !=nil){return nil,err}
	ptAS1, _ := json.Marshal(ac)
	var airc Aircraft
		err = json.Unmarshal(ptAS1, &airc)
		if err != nil {return nil, errors.New("Failed to Unmarshal Part #" + key)}
	var tx Log
		tx.Owner 		= airc.Owner
		tx.LType 		= "ADD"
	
	airc.Parts = append(airc.Parts, idpart)	
	airc.Logs = append(airc.Logs, tx)
	// Fin Partie Aircraft 

	//Update allAircraft 
			partzMap,err:=getAircraftMap(stub)
			partzMap[airc.Id_Aircraft] = airc
			allPAsBuytes, err := json.Marshal(partzMap)
			err=stub.PutState("allAircraft",allPAsBuytes)
			if err != nil {return nil,  err}
	//Fin update allAircraft 
	
// Debut Partie Part	
	part,err:=findPartById(stub,idpart)
		if err != nil {return nil, errors.New("Failed to get part #" + key)}
	ptAS, _ := json.Marshal(part)
		var pt Part
		err = json.Unmarshal(ptAS, &pt)
		if err != nil {return nil, errors.New("Failed to Unmarshal Part #" + key)}
		pt.Helicopter = key
		pt.Owner = airc.Owner
	var tf Log
		tf.Owner 		= pt.Owner
		tf.LType 		= "added to A/C: " + key
	pt.Logs = append(pt.Logs, tf)
	
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
// fin Partie Part 

fmt.Println("Responsible created successfully")	
return nil, nil
}
// ====================================================================
// Remove Parts from Aircraft 
// ====================================================================
func (t *SimpleChaincode)RemovePartFromAc(stub shim.ChaincodeStubInterface, args []string)([]byte, error) {

	key :=  args[0]
	idpart := args[1]

// Debut Partie Aircraft 
	ac,err:=findAircraftById(stub,key)
		if(err !=nil){return nil,err}
	ptAS1, _ := json.Marshal(ac)
	var airc Aircraft
		err = json.Unmarshal(ptAS1, &airc)
		if err != nil {return nil, errors.New("Failed to Unmarshal Part #" + key)}
	
	for i, v := range airc.Parts{
			if v == idpart {
				airc.Parts = append(airc.Parts[:i], airc.Parts[i+1:]...)
			break
		}
			}
	var tx Log
		tx.Owner 		= airc.Owner
		tx.LType 		= "REMOVE"
		airc.Logs = append(airc.Logs, tx)
// Fin Partie Aircraft 

	//Update allAircraft 
			partzMap,err:=getAircraftMap(stub)
			partzMap[airc.Id_Aircraft] = airc
			allPAsBuytes, err := json.Marshal(partzMap)
			err=stub.PutState("allAircraft",allPAsBuytes)
			if err != nil {return nil,  err}
	//Fin update allAircraft
	
// Debut Partie Part	
	part,err:=findPartById(stub,idpart)
		if err != nil {return nil, errors.New("Failed to get part #" + key)}
	ptAS, _ := json.Marshal(part)
		var pt Part
		err = json.Unmarshal(ptAS, &pt)
		if err != nil {return nil, errors.New("Failed to Unmarshal Part #" + key)}
		pt.Helicopter = ""
	var tf Log
		tf.Owner 		= pt.Owner
		tf.LType 		= "removed from A/C: " + key
	pt.Logs = append(pt.Logs, tf)
	
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
// fin Partie Part 

return nil, nil
}
// ====================================================================
// Obtenir tous les détails d'un aircraft à partir de son id Ses Logs 
// ====================================================================
func (t *SimpleChaincode) getAcDetails(stub shim.ChaincodeStubInterface, args []string)([]byte, error) {

	key :=  args[0]
	part,err:=findAircraftById(stub,key)
		if(err !=nil){return nil,err}
		return json.Marshal(part)
	}
// ====================================================================
// Afficher la liste détailéles de toutes les parts composants un Aircraft donné à partir de son id
// ====================================================================
func (t *SimpleChaincode)AcPartsListing(stub shim.ChaincodeStubInterface, args []string)([]byte, error) {

key := args [0]
username, err := getAttribute(stub, "username")
		if(err !=nil){return nil,err}
	role, err := getAttribute(stub, "role")
		if(err !=nil){return nil,err}
	//if supplier or manufacturer or customer or maintenance user =>only my parts
	showOnlyMyPart := role=="supplier" || role == "manufacturer" || role == "customer" || role == "maintenance_user"

	partMap,err:=getPartsIdMap(stub)
	if err != nil {return nil, errors.New("Failed to get Part")}
	parts := make([]Part, len(partMap))
	idx := 0
    for  _, part := range partMap {
    	if(showOnlyMyPart && part.Helicopter == key && part.Owner == username){
    		parts[idx] = part
    		idx++
    	} else if (!showOnlyMyPart || part.Helicopter == key){
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
// ==================================================================
// Afficher toutes les Aircraft créées en détail  
//===================================================================
func (t *SimpleChaincode) getAllAircraftsDetails(stub shim.ChaincodeStubInterface, args []string)([]byte, error){
	
	username, err := getAttribute(stub, "username")
		if(err !=nil){return nil,err}
	role, err := getAttribute(stub, "role")
		if(err !=nil){return nil,err}
	//if supplier or manufacturer or customer or maintenance user =>only my parts
	showOnlyMyPart := role=="supplier" || role == "manufacturer" || role == "customer" || role == "maintenance_user"

	fmt.Println("Start find getAllPartsDetails ")
	fmt.Println("Looking for All Parts With Details ")
	
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
	return nil, nil 
}
// =========================================================================================
// 					ACTIVITIES 
// =========================================================================================
// =========================
// Transfert de propriété 
// =========================
func (t *SimpleChaincode) AcOwnershipTransfer(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {
	var err error
	key := args[0]
	
	// Debut Partie Aircraft 
	ac,err:=findAircraftById(stub,key)
		if(err !=nil){return nil,err}
	ptAS1, _ := json.Marshal(ac)
	var airc Aircraft
		err = json.Unmarshal(ptAS1, &airc)
		if err != nil {return nil, errors.New("Failed to Unmarshal Part #" + key)}
		airc.Owner = args[1] 
	var tx Log
		tx.Owner 		= airc.Owner
		tx.LType 		= "OWNERNSHIP_TRANSFER"
	airc.Logs = append(airc.Logs, tx)
	// Fin Partie Aircraft 

	//Update allAircraft 
			partzMap,err:=getAircraftMap(stub)
			partzMap[airc.Id_Aircraft] = airc
			allPAsBuytes, err := json.Marshal(partzMap)
			err=stub.PutState("allAircraft",allPAsBuytes)
			if err != nil {return nil,  err}
	//Fin update allAircraft 
	
	// Parts 
	
	for i := range airc.Parts{
		part,err:=findPartById(stub,airc.Parts[i])
			if err != nil {return nil, errors.New("Failed to get part #" + key)}
			ptAS, _ := json.Marshal(part)
		var pt Part
			err = json.Unmarshal(ptAS, &pt)
			if err != nil {return nil, errors.New("Failed to Unmarshal Part #" + key)}
		pt.Owner = args[1]
		var tx Log
			tx.Owner 		= pt.Owner
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
			i++
		}
return nil, nil
}