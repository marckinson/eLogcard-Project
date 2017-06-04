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

	// If the PN or/and the SN is/are already used, a part can't be created.
	n:= checkAnAircraft(stub, args[0])
		if n != nil { return nil, errors.New(n.Error())}	
	o:= checkSnAircraft(stub, args[1])
		if o != nil { return nil, errors.New(o.Error())}

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

//Update allAircraftsAn
	partMap1,err:=getAircraftAnMap(stub)
		partMap1[air.AN] = air
		allPAsBytes1, err := json.Marshal(partMap1)
		err=stub.PutState("allAircraftsAn",allPAsBytes1)
		if err != nil {return nil, err}
//Fin update allAircraftsAn

//Update allAircraftsSn
	partMap2,err:=getAircraftSnMap(stub)
		partMap2[air.SN] = air
		allPAsBytes2, err := json.Marshal(partMap2)
		err=stub.PutState("allAircraftsSn",allPAsBytes2)
		if err != nil {return nil, err}
//Fin update allAircraftsSn	
		
		
fmt.Println("Responsible created successfully")	
return nil, nil
}
// ====================================================================
// addPartToAc (Parts qui n'appartiennent à aucun Assembly )
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
		tx.LType 		= "PART_AFFILIATION"
	
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
	//Update allAircraftsAn
	partzMap1,err:=getAircraftAnMap(stub)
		partzMap1[airc.AN] = airc
		allPAsBytes11, err := json.Marshal(partzMap1)
		err=stub.PutState("allAircraftsAn",allPAsBytes11)
		if err != nil {return nil, err}
//Fin update allAircraftsAn
//Update allAircraftsSn
	partzMap2,err:=getAircraftSnMap(stub)
		partzMap2[airc.SN] = airc
		allPAsBytes22, err := json.Marshal(partzMap2)
		err=stub.PutState("allAircraftsSn",allPAsBytes22)
		if err != nil {return nil, err}
//Fin update allAircraftsSn	
	
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
		tf.LType 		= "ADDED TO A/C: " + key
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
		tx.LType 		= "PART_REMOVAL"
		airc.Logs = append(airc.Logs, tx)
// Fin Partie Aircraft 

	//Update allAircraft 
			partzMap,err:=getAircraftMap(stub)
			partzMap[airc.Id_Aircraft] = airc
			allPAsBuytes, err := json.Marshal(partzMap)
			err=stub.PutState("allAircraft",allPAsBuytes)
			if err != nil {return nil,  err}
	//Fin update allAircraft
//Update allAircraftsAn
	partzMap1,err:=getAircraftAnMap(stub)
		partzMap1[airc.AN] = airc
		allPAsBytes11, err := json.Marshal(partzMap1)
		err=stub.PutState("allAircraftsAn",allPAsBytes11)
		if err != nil {return nil, err}
//Fin update allAircraftsAn
//Update allAircraftsSn
	partzMap2,err:=getAircraftSnMap(stub)
		partzMap2[airc.SN] = airc
		allPAsBytes22, err := json.Marshal(partzMap2)
		err=stub.PutState("allAircraftsSn",allPAsBytes22)
		if err != nil {return nil, err}
//Fin update allAircraftsSn	
	
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
		tf.LType 		= "REMOVED FROM A/C: " + key
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
	//Update allAircraftsAn
	partzMap1,err:=getAircraftAnMap(stub)
		partzMap1[airc.AN] = airc
		allPAsBytes11, err := json.Marshal(partzMap1)
		err=stub.PutState("allAircraftsAn",allPAsBytes11)
		if err != nil {return nil, err}
//Fin update allAircraftsAn
//Update allAircraftsSn
	partzMap2,err:=getAircraftSnMap(stub)
		partzMap2[airc.SN] = airc
		allPAsBytes22, err := json.Marshal(partzMap2)
		err=stub.PutState("allAircraftsSn",allPAsBytes22)
		if err != nil {return nil, err}
//Fin update allAircraftsSn	
	
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


// =========================
// Exchange a defective part with another 
// =========================
func (t *SimpleChaincode)Replace(stub shim.ChaincodeStubInterface, args []string)([]byte, error) {

	key :=  args[0]
	idpart := args[1]
	idpart1 := args[2]

	// Debut Partie Aircraft 
	ac,err:=findAircraftById(stub,key)
		if(err !=nil){return nil,err}
	ptAS1, _ := json.Marshal(ac)
	var airc Aircraft
		err = json.Unmarshal(ptAS1, &airc)
		if err != nil {return nil, errors.New("Failed to Unmarshal Part #" + key)}
		airc.Parts = append(airc.Parts, idpart1)	

	for i, v := range airc.Parts{
			if v == idpart {
				airc.Parts = append(airc.Parts[:i], airc.Parts[i+1:]...)
			break
		}
			}
	var tx Log
		tx.Owner 		= airc.Owner
		tx.LType 		= "PART_SUBSTITUTION"
	airc.Logs = append(airc.Logs, tx)
	// Fin Partie Aircraft 

	//Update allAircraft 
			partzMap,err:=getAircraftMap(stub)
			partzMap[airc.Id_Aircraft] = airc
			allPAsBuytes, err := json.Marshal(partzMap)
			err=stub.PutState("allAircraft",allPAsBuytes)
			if err != nil {return nil,  err}
	//Fin update allAircraft 
	//Update allAircraftsAn
		partzMap1,err:=getAircraftAnMap(stub)
		partzMap1[airc.AN] = airc
		allPAsBytes11, err := json.Marshal(partzMap1)
		err=stub.PutState("allAircraftsAn",allPAsBytes11)
		if err != nil {return nil, err}
	//Fin update allAircraftsAn
	//Update allAircraftsSn
		partzMap2,err:=getAircraftSnMap(stub)
		partzMap2[airc.SN] = airc
		allPAsBytes22, err := json.Marshal(partzMap2)
		err=stub.PutState("allAircraftsSn",allPAsBytes22)
		if err != nil {return nil, err}
	//Fin update allAircraftsSn	
	
	
// Debut Partie Part	
	part,err:=findPartById(stub,idpart)
		if err != nil {return nil, errors.New("Failed to get part #" + key)}
	ptAS, _ := json.Marshal(part)
		var pt Part
		err = json.Unmarshal(ptAS, &pt)
		if err != nil {return nil, errors.New("Failed to Unmarshal Part #" + key)}
		pt.Helicopter = ""
		pt.Owner = airc.Owner
	var tf Log
		tf.Owner 		= pt.Owner
		tf.LType 		= "REMOVED FROM A/C " + key + " SUBSTITUTED BY PART: " + idpart1
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


partt,err:=findPartById(stub,idpart1)
		if err != nil {return nil, errors.New("Failed to get part #" + key)}
	ptASS, _ := json.Marshal(partt)
		var ptt Part
		err = json.Unmarshal(ptASS, &ptt)
		if err != nil {return nil, errors.New("Failed to Unmarshal Part #" + key)}
		ptt.Helicopter = key
		ptt.Owner = airc.Owner
		ptt.Responsible = pt.Responsible
		ptt.Assembly = pt.Assembly
		ptt.PN = pt.PN

		var tff Log
		tff.Owner 		= ptt.Owner
		tff.LType 		= "ADDED TO A/C " + key + " AND SUBSTITUTES PART: " + idpart
	ptt.Logs = append(ptt.Logs, tff)
//Update allParts 
	partMapp,err:=getPartsIdMap(stub)
		partMapp[ptt.Id] = ptt
		allPAsBytess, err := json.Marshal(partMapp)
		err=stub.PutState("allParts",allPAsBytess)
	if err != nil {return nil, err}
//Fin update allParts 
//Update allPartsPn
	partMap11,err:=getPartsPnMap(stub)
		partMap11[ptt.PN] = ptt
		allPAsBytes111, err := json.Marshal(partMap11)
		err=stub.PutState("allPartsPn",allPAsBytes111)
		if err != nil {return nil, err}
//Fin update allPartsPn
//Update allPartsSn
	partMap22,err:=getPartsSnMap(stub)
		partMap22[ptt.SN] = ptt
		allPAsBytes222, err := json.Marshal(partMap22)
		err=stub.PutState("allPartsSn",allPAsBytes222)
		if err != nil {return nil, err}
//Fin update allPartsSn

// fin Partie Part 

fmt.Println("Responsible created successfully")	
return nil, nil

}

// =========================
// Add an assembly 
// =========================
func (t *SimpleChaincode)AddAssemblyToAc(stub shim.ChaincodeStubInterface, args []string)([]byte, error) {

	key :=  args[0]
	idassembly := args[1]

// Debut Partie Aircraft 
	ac,err:=findAircraftById(stub,key)
		if(err !=nil){return nil,err}
	ptAS1, _ := json.Marshal(ac)
	var airc Aircraft
		err = json.Unmarshal(ptAS1, &airc)
		if err != nil {return nil, errors.New("Failed to Unmarshal Part #" + key)}
	var tx Log
		tx.Owner 		= airc.Owner
		tx.LType 		= "ASSEMBLY_AFFILIATION: " + idassembly
	airc.Assemblies = append(airc.Assemblies, idassembly)	
	airc.Logs = append(airc.Logs, tx)
	//Update allAircraft 
			partzMap,err:=getAircraftMap(stub)
			partzMap[airc.Id_Aircraft] = airc
			allPAsBuytes, err := json.Marshal(partzMap)
			err=stub.PutState("allAircraft",allPAsBuytes)
			if err != nil {return nil,  err}
	//Fin update allAircraft 
	//Update allAircraftsAn
	partzMap1,err:=getAircraftAnMap(stub)
		partzMap1[airc.AN] = airc
		allPAsBytes11, err := json.Marshal(partzMap1)
		err=stub.PutState("allAircraftsAn",allPAsBytes11)
		if err != nil {return nil, err}
	//Fin update allAircraftsAn
	//Update allAircraftsSn
	partzMap2,err:=getAircraftSnMap(stub)
		partzMap2[airc.SN] = airc
		allPAsBytes22, err := json.Marshal(partzMap2)
		err=stub.PutState("allAircraftsSn",allPAsBytes22)
		if err != nil {return nil, err}
	//Fin update allAircraftsSn	
// Fin Partie Aircraft 


// Debut Partie Assembly	
	part,err:=findAssemblyById(stub,idassembly)
		if err != nil {return nil, errors.New("Failed to get part #" + key)}
	ptASS, _ := json.Marshal(part)
		var pt Assembly
		err = json.Unmarshal(ptASS, &pt)
		if err != nil {return nil, errors.New("Failed to Unmarshal Part #" + key)}
		pt.Helicopter = key
		pt.Owner = airc.Owner
	var tf Log
		tf.Owner 		= pt.Owner
		tf.LType 		= "ADDED TO A/C: " + key
	pt.Logs = append(pt.Logs, tf)
	
//Update allParts 
	partMap,err:=getAssemblyMap(stub)
		partMap[pt.Id_Assembly] = pt
		allPAsBytes, err := json.Marshal(partMap)
		err=stub.PutState("allAssembly",allPAsBytes)
	if err != nil {return nil, err}
//Fin update allParts 
//Update allPartsPn
	partMap1,err:=getAssembliesAnMap(stub)
		partMap1[pt.AN] = pt
		allPAsBytes1, err := json.Marshal(partMap1)
		err=stub.PutState("allAssembliesAn",allPAsBytes1)
		if err != nil {return nil, err}
//Fin update allPartsPn
//Update allPartsSn
	partMap2,err:=getAssembliesSnMap(stub)
		partMap2[pt.SN] = pt
		allPAsBytes2, err := json.Marshal(partMap2)
		err=stub.PutState("allAssembliesSn",allPAsBytes2)
		if err != nil {return nil, err}
//Fin update allPartsSn
// fin Partie Part 

fmt.Println("Responsible created successfully")	
return nil, nil
}


// =========================
// Remove an Assembly 
// =========================
func (t *SimpleChaincode)RemoveAssemblyFromAc(stub shim.ChaincodeStubInterface, args []string)([]byte, error) {

	key :=  args[0]
	idassembly := args[1]

// Debut Partie Aircraft 
	ac,err:=findAircraftById(stub,key)
		if(err !=nil){return nil,err}
	ptAS1, _ := json.Marshal(ac)
	var airc Aircraft
		err = json.Unmarshal(ptAS1, &airc)
		if err != nil {return nil, errors.New("Failed to Unmarshal Part #" + key)}
	
	for i, v := range airc.Assemblies{
			if v == idassembly {
				airc.Assemblies = append(airc.Assemblies[:i], airc.Assemblies[i+1:]...)
			break
		}
			}
	var tx Log
		tx.Owner 		= airc.Owner
		tx.LType 		= "ASSEMBLY_REMOVAL: " + idassembly
		airc.Logs = append(airc.Logs, tx)
	//Update allAircraft 
			partzMap,err:=getAircraftMap(stub)
			partzMap[airc.Id_Aircraft] = airc
			allPAsBuytes, err := json.Marshal(partzMap)
			err=stub.PutState("allAircraft",allPAsBuytes)
			if err != nil {return nil,  err}
	//Fin update allAircraft
	//Update allAircraftsAn
	partzMap1,err:=getAircraftAnMap(stub)
		partzMap1[airc.AN] = airc
		allPAsBytes11, err := json.Marshal(partzMap1)
		err=stub.PutState("allAircraftsAn",allPAsBytes11)
		if err != nil {return nil, err}
	//Fin update allAircraftsAn
	//Update allAircraftsSn
	partzMap2,err:=getAircraftSnMap(stub)
		partzMap2[airc.SN] = airc
		allPAsBytes22, err := json.Marshal(partzMap2)
		err=stub.PutState("allAircraftsSn",allPAsBytes22)
		if err != nil {return nil, err}
	//Fin update allAircraftsSn	
// Fin Partie Aircraft 

// Debut Partie Assembly	
	part,err:=findAssemblyById(stub,idassembly)
		if err != nil {return nil, errors.New("Failed to get part #" + key)}
	ptASS, _ := json.Marshal(part)
		var pt Assembly
		err = json.Unmarshal(ptASS, &pt)
		if err != nil {return nil, errors.New("Failed to Unmarshal Part #" + key)}
		pt.Helicopter = ""
		pt.Owner = airc.Owner
	var tf Log
		tf.Owner 		= pt.Owner
		tf.LType 		= "REMOVED FROM A/C: " + key
	pt.Logs = append(pt.Logs, tf)
	
//Update allParts 
	partMap,err:=getAssemblyMap(stub)
		partMap[pt.Id_Assembly] = pt
		allPAsBytes, err := json.Marshal(partMap)
		err=stub.PutState("allAssembly",allPAsBytes)
	if err != nil {return nil, err}
//Fin update allParts 
//Update allPartsPn
	partMap1,err:=getAssembliesAnMap(stub)
		partMap1[pt.AN] = pt
		allPAsBytes1, err := json.Marshal(partMap1)
		err=stub.PutState("allAssembliesAn",allPAsBytes1)
		if err != nil {return nil, err}
//Fin update allPartsPn
//Update allPartsSn
	partMap2,err:=getAssembliesSnMap(stub)
		partMap2[pt.SN] = pt
		allPAsBytes2, err := json.Marshal(partMap2)
		err=stub.PutState("allAssembliesSn",allPAsBytes2)
		if err != nil {return nil, err}
//Fin update allPartsSn
// fin Partie Part 

fmt.Println("Responsible created successfully")	
return nil, nil
}
