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
// 					Assembly
// =========================================================================================
// ===================================================================
// Creation of the Aircraft 
// Only registered suppliers and manufacturers can create Parts.  
// ===================================================================
func (t *SimpleChaincode) createAssembly(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {
	fmt.Println("Running createAssembly")

	var err error
	var assemb Assembly
		assemb.AN = args [0] // Assembly Number 
		assemb.SN = args [1] // Serial Number 
		assemb.Id_Assembly = args[2] // Id of the Assembly 
		assemb.Owner = args [3] // Owner of the Assembly 
	var tx Log
		tx.Owner 		= assemb.Owner
		tx.VDate 		= args[4]
		tx.LType 		= "CREATE"
	assemb.Logs = append(assemb.Logs, tx)

//Commit part to ledger
	ptAsBytes, _ := json.Marshal(assemb)
		err = stub.PutState(assemb.Id_Assembly, ptAsBytes)
		if err != nil {return nil, err}	
//Fin Commit part to ledger

//Update allAssembly 
		partzMap,err:=getAssemblyMap(stub)
		partzMap[assemb.Id_Assembly] = assemb
		allPAsBuytes, err := json.Marshal(partzMap)
		err=stub.PutState("allAssembly",allPAsBuytes)
		if err != nil {return nil, err}
//Fin update allAssembly 
		
fmt.Println("Responsible created successfully")	
return nil, nil
}
// ====================================================================
// addPartToAssemb
// ====================================================================
func (t *SimpleChaincode)addPartToAssemb(stub shim.ChaincodeStubInterface, args []string)([]byte, error) {

	key :=  args[0]
	idpart := args[1]

	// Debut Partie Assembly 
	ac,err:=findAssemblyById(stub,key)
		if(err !=nil){return nil,err}
	ptAS1, _ := json.Marshal(ac)
	var assemb Assembly
		err = json.Unmarshal(ptAS1, &assemb)
		if err != nil {return nil, errors.New("Failed to Unmarshal Part #" + key)}
	var tx Log
		tx.Owner 		= assemb.Owner
		tx.LType 		= "PART_AFFILIATION"
	
	assemb.Parts = append(assemb.Parts, idpart)	
	assemb.Logs = append(assemb.Logs, tx)
	// Fin Partie Assembly 

//Update allAssembly 
		partzMap,err:=getAssemblyMap(stub)
		partzMap[assemb.Id_Assembly] = assemb
		allPAsBuytes, err := json.Marshal(partzMap)
		err=stub.PutState("allAssembly",allPAsBuytes)
		if err != nil {return nil, err}
//Fin update allAssembly
	
// Debut Partie Part	
	part,err:=findPartById(stub,idpart)
		if err != nil {return nil, errors.New("Failed to get part #" + key)}
	ptAS, _ := json.Marshal(part)
		var pt Part
		err = json.Unmarshal(ptAS, &pt)
		if err != nil {return nil, errors.New("Failed to Unmarshal Part #" + key)}
		pt.Assembly = key
		pt.Owner = assemb.Owner
	var tf Log
		tf.Owner 		= pt.Owner
		tf.LType 		= "added to Assemb: " + key
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
func (t *SimpleChaincode)RemovePartFromAs(stub shim.ChaincodeStubInterface, args []string)([]byte, error) {
key :=  args[0]
	idpart := args[1]

// Debut Partie Aircraft 
	ac,err:=findAssemblyById(stub,key)
		if(err !=nil){return nil,err}
	ptAS1, _ := json.Marshal(ac)
	var airc Assembly
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

//Update allAssembly 
		partzMap,err:=getAssemblyMap(stub)
		partzMap[airc.Id_Assembly] = airc
		allPAsBuytes, err := json.Marshal(partzMap)
		err=stub.PutState("allAssembly",allPAsBuytes)
		if err != nil {return nil, err}
//Fin update allAssembly
	
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
		tf.LType 		= "removed from Assemb: " + key
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
// Obtenir tous les détails d'un aircraft à partir de son id 
// ====================================================================
func (t *SimpleChaincode) getAssembDetails(stub shim.ChaincodeStubInterface, args []string)([]byte, error) {

	key :=  args[0]
	part,err:=findAssemblyById(stub,key)
		if(err !=nil){return nil,err}
		return json.Marshal(part)
	}
// ====================================================================
// Afficher la liste détailéles de toutes les parts composants un Assembly donné à partir de son id
// ====================================================================
func (t *SimpleChaincode)AssembPartsListing(stub shim.ChaincodeStubInterface, args []string)([]byte, error) {

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
    	if(showOnlyMyPart && part.Assembly == key && part.Owner == username){
    		parts[idx] = part
    		idx++
    	} else if (!showOnlyMyPart || part.Assembly == key){
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
// Afficher toutes les Assembly créées en détail  
//===================================================================
func (t *SimpleChaincode) getAllAssembliesDetails(stub shim.ChaincodeStubInterface, args []string)([]byte, error){
	
	username, err := getAttribute(stub, "username")
		if(err !=nil){return nil,err}
	role, err := getAttribute(stub, "role")
		if(err !=nil){return nil,err}
	//if supplier or manufacturer or customer or maintenance user =>only my parts
	showOnlyMyPart := role=="supplier" || role == "manufacturer" || role == "customer" || role == "maintenance_user"

	fmt.Println("Start find getAllPartsDetails ")
	fmt.Println("Looking for All Parts With Details ")
	
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
	return nil, nil 
}
// =========================================================================================
// 					ACTIVITIES 
// =========================================================================================
// =========================
// Transfert de propriété 
// =========================
func (t *SimpleChaincode) AssembOwnershipTransfer(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {
	var err error
	key := args[0]
	
	// Debut Partie Assembly 
	ac,err:=findAssemblyById(stub,key)
		if(err !=nil){return nil,err}
	ptAS1, _ := json.Marshal(ac)
	var assemb Assembly
		err = json.Unmarshal(ptAS1, &assemb)
		if err != nil {return nil, errors.New("Failed to Unmarshal Part #" + key)}
		assemb.Owner = args[1] 
	var tx Log
		tx.Owner 		= assemb.Owner
		tx.LType 		= "OWNERNSHIP_TRANSFER"
	assemb.Logs = append(assemb.Logs, tx)
	// Fin Partie Aircraft 

//Update allAssembly 
		partzMap,err:=getAssemblyMap(stub)
		partzMap[assemb.Id_Assembly] = assemb
		allPAsBuytes, err := json.Marshal(partzMap)
		err=stub.PutState("allAssembly",allPAsBuytes)
		if err != nil {return nil, err}
//Fin update allAssembly
	
	// Parts 
	
	for i := range assemb.Parts{
		part,err:=findPartById(stub,assemb.Parts[i])
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