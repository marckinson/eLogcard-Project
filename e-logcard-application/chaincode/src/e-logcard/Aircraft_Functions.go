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

	y:= UpdateAircraft (stub, air) 
		if y != nil { return nil, errors.New(y.Error())}

fmt.Println("Responsible created successfully")	
return nil, nil
}
// ====================================================================
// addPartToAc (Parts qui n'appartiennent à aucun Assembly )
// ====================================================================
func (t *SimpleChaincode)addPartToAc(stub shim.ChaincodeStubInterface, args []string)([]byte, error) {

	key :=  args[0]
	idpart := args[1]
	
	test, err := findPartById (stub, idpart) 
		if(err !=nil){return nil,err}
	ptA, _ := json.Marshal(test)
	var ppart Part
		err = json.Unmarshal(ptA, &ppart)
		if err != nil {return nil, errors.New("Failed to Unmarshal Part #" + key)}
	if (ppart.Helicopter == "" && ppart.Assembly == "") {	
	// Debut Partie Aircraft 
	ac,err:=findAircraftById(stub,key)
		if(err !=nil){return nil,err}
	ptAS1, _ := json.Marshal(ac)
	var airc Aircraft
		err = json.Unmarshal(ptAS1, &airc)
		if err != nil {return nil, errors.New("Failed to Unmarshal Part #" + key)}
	var tx Log
		tx.Owner 		= airc.Owner
		tx.LType 		= "PART_AFFILIATION: " + idpart
	
	airc.Parts = append(airc.Parts, idpart)	
	airc.Logs = append(airc.Logs, tx)
	// Fin Partie Aircraft 

	y:= UpdateAircraft (stub, airc) 
		if y != nil { return nil, errors.New(y.Error())}

// Debut Partie Part	
	part,err:=findPartById(stub,idpart)
		if err != nil {return nil, errors.New("Failed to get part #" + key)}
	ptAS, _ := json.Marshal(part)
		var pt Part
		err = json.Unmarshal(ptAS, &pt)
		if err != nil {return nil, errors.New("Failed to Unmarshal Part #" + key)}
		pt.Helicopter = key
		pt.Owner = airc.Owner   // Le champ Helicopter de la part prend la valeur de L'aircraft sur lequel elle est ajoutée
	var tf Log
		tf.Owner 		= pt.Owner
		tf.LType 		= "ADDED TO A/C: " + key
	pt.Logs = append(pt.Logs, tf)
	
	e:= UpdatePart (stub, pt) 
		if e != nil { return nil, errors.New(e.Error())}
	
	} else if (ppart.Helicopter != "" && ppart.Assembly != "") {
		return nil, errors.New ("Impossible") }

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
		tx.LType 		= "PART_REMOVAL: " + idpart
		airc.Logs = append(airc.Logs, tx)
// Fin Partie Aircraft 

	y:= UpdateAircraft (stub, airc) 
		if y != nil { return nil, errors.New(y.Error())}

// Debut Partie Part	
	part,err:=findPartById(stub,idpart)
		if err != nil {return nil, errors.New("Failed to get part #" + key)}
	ptAS, _ := json.Marshal(part)
		var pt Part
		err = json.Unmarshal(ptAS, &pt)
		if err != nil {return nil, errors.New("Failed to Unmarshal Part #" + key)}
		pt.Helicopter = "" // Le champ Helicopter de la part retirée de l'Helicopter revient à nul.
	var tf Log
		tf.Owner 		= pt.Owner
		tf.LType 		= "REMOVED FROM A/C: " + key
	pt.Logs = append(pt.Logs, tf)
	
	e:= UpdatePart (stub, pt) 
		if e != nil { return nil, errors.New(e.Error())}
		
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

	y:= UpdateAircraft (stub, airc) 
		if y != nil { return nil, errors.New(y.Error())}
	
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
		
	e:= UpdatePart (stub, pt) 
		if e != nil { return nil, errors.New(e.Error())}
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
		tx.LType 		= "PART_SUBSTITUTION : " + idpart1 +  " replace " + idpart
	airc.Logs = append(airc.Logs, tx)
	// Fin Partie Aircraft 

	y:= UpdateAircraft (stub, airc) 
		if y != nil { return nil, errors.New(y.Error())}
	
// Debut Partie Part	
	part,err:=findPartById(stub,idpart)
		if err != nil {return nil, errors.New("Failed to get part #" + key)}
	ptAS, _ := json.Marshal(part)
		var pt Part
		err = json.Unmarshal(ptAS, &pt)
		if err != nil {return nil, errors.New("Failed to Unmarshal Part #" + key)}
		pt.Helicopter = ""  // Le champ Helicopter de la part retirée de l'Helicopter revient à nul.
		pt.Owner = airc.Owner
	var tf Log
		tf.Owner 		= pt.Owner
		tf.LType 		= "REMOVED FROM A/C " + key + " SUBSTITUTED BY PART: " + idpart1
	pt.Logs = append(pt.Logs, tf)

	e:= UpdatePart (stub, pt) 
		if e != nil { return nil, errors.New(e.Error())}

partt,err:=findPartById(stub,idpart1)
		if err != nil {return nil, errors.New("Failed to get part #" + key)}
	ptASS, _ := json.Marshal(partt)
		var ptt Part
		err = json.Unmarshal(ptASS, &ptt)
		if err != nil {return nil, errors.New("Failed to Unmarshal Part #" + key)}
		ptt.Helicopter = key
		ptt.Owner = airc.Owner  // Le champ Helicopter de la part rajoutée à l'A/C prend la valeur A/C.
		ptt.Responsible = pt.Responsible
		ptt.Assembly = pt.Assembly
		ptt.PN = pt.PN

		var tff Log
		tff.Owner 		= ptt.Owner
		tff.LType 		= "ADDED TO A/C " + key + " AND SUBSTITUTES PART: " + idpart
	ptt.Logs = append(ptt.Logs, tff)

	r:= UpdatePart (stub, ptt) 
		if e != nil { return nil, errors.New(r.Error())}

fmt.Println("Responsible created successfully")	
return nil, nil

}
// =========================
// Add an assembly 
// =========================
func (t *SimpleChaincode)AddAssemblyToAc(stub shim.ChaincodeStubInterface, args []string)([]byte, error) {

	key :=  args[0]
	idassembly := args[1]
// Verification 
	test, err := findAssemblyById (stub, idassembly) 
		if(err !=nil){return nil,err}
	ptA, _ := json.Marshal(test)
	var ppart Assembly
		err = json.Unmarshal(ptA, &ppart)
		if err != nil {return nil, errors.New("Failed to Unmarshal Part #" + key)}
if (ppart.Helicopter == "") {  // Un assembly appartenant à un A/C ne peut pas être ajouté à un A/C

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
	
	t:= UpdateAircraft (stub, airc) 
		if t != nil { return nil, errors.New(t.Error())}
	
// Debut Partie Assembly	
	part,err:=findAssemblyById(stub,idassembly)
		if err != nil {return nil, errors.New("Failed to get part #" + key)}
	ptASS, _ := json.Marshal(part)
		var pt Assembly
		err = json.Unmarshal(ptASS, &pt)
		if err != nil {return nil, errors.New("Failed to Unmarshal Part #" + key)}
		pt.Helicopter = key
		pt.Owner = airc.Owner // Le champ Helicopter de l'Assembly ajouté à l'Helicopter prend la valeur A/C
	var tf Log
		tf.Owner 		= pt.Owner
		tf.LType 		= "ADDED TO A/C: " + key
	pt.Logs = append(pt.Logs, tf)
	
	
	// Début Partie Parts 

		for i := range pt.Parts{
		pppart,err:=findPartById(stub,pt.Parts[i])
			if err != nil {return nil, errors.New("Failed to get part #" + key)}
			ptASS, _ := json.Marshal(pppart)
		var ppt Part
			err = json.Unmarshal(ptASS, &ppt)
			if err != nil {return nil, errors.New("Failed to Unmarshal Part #" + key)}
		ppt.Helicopter = args[1]
		var txxx Log
			txxx.Owner 		= pt.Owner
			txxx.LType 		= "A/C Affiliation"
			ppt.Logs = append(ppt.Logs, txxx)
		}
// Fin Partie Parts 
	
	
	y:= UpdateAssembly (stub, pt) 
		if y != nil { return nil, errors.New(y.Error())}
	
} else if (ppart.Helicopter != "") {  // Un assembly appartenant à un A/C ne peut pas être ajouté à un A/C
		return nil, errors.New ("Impossible") }

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
		
	y:= UpdateAircraft (stub, airc) 
		if y != nil { return nil, errors.New(y.Error())}

// Debut Partie Assembly	
	part,err:=findAssemblyById(stub,idassembly)
		if err != nil {return nil, errors.New("Failed to get part #" + key)}
	ptASS, _ := json.Marshal(part)
		var pt Assembly
		err = json.Unmarshal(ptASS, &pt)
		if err != nil {return nil, errors.New("Failed to Unmarshal Part #" + key)}
		pt.Helicopter = "" // Le champ Helicopter de l'Assembly retirée de l'Helicopter revient à nul.
		pt.Owner = airc.Owner
	var tf Log
		tf.Owner 		= pt.Owner
		tf.LType 		= "REMOVED FROM A/C: " + key
	pt.Logs = append(pt.Logs, tf)

	e:= UpdateAssembly (stub, pt) 
		if e != nil { return nil, errors.New(e.Error())}

fmt.Println("Responsible created successfully")	
return nil, nil
}
