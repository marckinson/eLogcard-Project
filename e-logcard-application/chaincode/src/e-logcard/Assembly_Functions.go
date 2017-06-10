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
// Creation of the Assembly 
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
	var tx LogAssembly
		tx.Owner 		= assemb.Owner
		tx.VDate 		= args[4]
		tx.LType 		= "CREATE"
	assemb.Logs = append(assemb.Logs, tx)
// If the PN or/and the SN is/are already used, a part can't be created.
	n:= checkAnAssembly(stub, args[0])
		if n != nil { return nil, errors.New(n.Error())}	
	o:= checkSnAssembly(stub, args[1])
		if o != nil { return nil, errors.New(o.Error())}
//Commit assembly to ledger
	ptAsBytes, _ := json.Marshal(assemb)
		err = stub.PutState(assemb.Id_Assembly, ptAsBytes)
		if err != nil {return nil, err}	
//Fin Commit assembly to ledger
	y:= UpdateAssembly (stub, assemb) 
		if y != nil { return nil, errors.New(y.Error())}
	
fmt.Println("Responsible created successfully")	
return nil, nil
}

// =========================================================================================
// 					ACTIVITIES ON ASSEMBLIES
// =========================================================================================
// ====================================================================
// addPartToAssemb
// ====================================================================
func (t *SimpleChaincode)addPartToAssemb(stub shim.ChaincodeStubInterface, args []string)([]byte, error) {

	key :=  args[0]
	idpart := args[1]
	
// Vérification
	test, err := findPartById (stub, idpart) 
		if(err !=nil){return nil,err}
	ptA, _ := json.Marshal(test)
	var ppart Part
		err = json.Unmarshal(ptA, &ppart)
		if err != nil {return nil, errors.New("Failed to Unmarshal Part #" + key)}
	if (ppart.Helicopter == "" && ppart.Assembly == "") {	
// Fin vérification

// Debut Partie Assembly 
	ac,err:=findAssemblyById(stub,key)
		if(err !=nil){return nil,err}
	ptAS1, _ := json.Marshal(ac)
	var assemb Assembly
		err = json.Unmarshal(ptAS1, &assemb)
		if err != nil {return nil, errors.New("Failed to Unmarshal Part #" + key)}
	var tx LogAssembly
		tx.Owner 		= assemb.Owner
		tx.LType 		= "PART_AFFILIATION: " + idpart
		tx.VDate		= args [2]
	assemb.Parts = append(assemb.Parts, idpart)	
	assemb.Logs = append(assemb.Logs, tx)
	y:= UpdateAssembly (stub, assemb) 
		if y != nil { return nil, errors.New(y.Error())}
// Fin Partie Assembly 

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
		tf.Responsible 	= pt.Responsible
		tf.Owner 		= pt.Owner
		tf.LType 		= "ADDED TO ASSEMBLY: " + key
		tf.VDate		= args [2]
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
	var tx LogAssembly
		tx.Owner 		= airc.Owner
		tx.LType 		= "PART_REMOVAL: " + idpart
		tx.VDate		= args [2]
		airc.Logs = append(airc.Logs, tx)
// Fin Partie Aircraft 

y:= UpdateAssembly (stub, airc) 
		if y != nil { return nil, errors.New(y.Error())}
		
// Debut Partie Part	
	part,err:=findPartById(stub,idpart)
		if err != nil {return nil, errors.New("Failed to get part #" + key)}
	ptAS, _ := json.Marshal(part)
		var pt Part
		err = json.Unmarshal(ptAS, &pt)
		if err != nil {return nil, errors.New("Failed to Unmarshal Part #" + key)}
		pt.Assembly = ""
		pt.Owner = "REMOVED_MANAGER"
		pt.Responsible = "REMOVED_MANAGER"
	var tf Log
		tf.Responsible	= pt.Responsible
		tf.Owner 		= pt.Owner
		tf.LType 		= "REMOVED FROM ASSEMBLY: " + key
		tf.VDate		= args [2]
	pt.Logs = append(pt.Logs, tf)
	
	e:= UpdatePart (stub, pt) 
		if e != nil { return nil, errors.New(e.Error())}
return nil, nil
}
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
	var tx LogAssembly
		tx.Owner 		= assemb.Owner
		tx.LType 		= "OWNERNSHIP_TRANSFER"
		tx.VDate		= args [2]
	assemb.Logs = append(assemb.Logs, tx)
	// Fin Partie Aircraft 

	y:= UpdateAssembly (stub, assemb) 
		if y != nil { return nil, errors.New(y.Error())}
	
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
			tx.Responsible	= pt.Responsible
			tx.Owner 		= pt.Owner
			tx.LType 		= "OWNERNSHIP_TRANSFER"
			tx.VDate  		= args [2]
			pt.Logs = append(pt.Logs, tx)
	
	e:= UpdatePart (stub, pt) 
		if e != nil { return nil, errors.New(e.Error())}

			i++
		}
return nil, nil
}

// =========================
// Scrapp an Assembly  
// =========================
func (t *SimpleChaincode) scrappAssembly(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {

	var err error
	var key string 
	key = args[0]
	part,err:=findAssemblyById(stub,key)
		if err != nil {return nil, errors.New("Failed to get part #" + key)}
		ptAS, _ := json.Marshal(part)
	var assemb Assembly
		err = json.Unmarshal(ptAS, &assemb)
		if err != nil {return nil, errors.New("Failed to Unmarshal Part #" + key)}
		assemb.Owner = "SCRAPPING_MANAGER"
		assemb.AN = ""
		assemb.Helicopter = ""
	var tx LogAssembly
		tx.Owner 		= assemb.Owner
		tx.VDate 		= args[1]
		tx.LType 		= "SCRAPPING"
		assemb.Logs = append(assemb.Logs, tx)
	e:= UpdateAssembly (stub, assemb) 
		if e != nil { return nil, errors.New(e.Error())}

		
// Debut Partie Part 	
			for i := range assemb.Parts{
			part3,err:=findPartById(stub,assemb.Parts[i])
			if err != nil {return nil, errors.New("Failed to get part3 #" + key)}
			ptAS3, _ := json.Marshal(part3)
		var pt1 Part
			err = json.Unmarshal(ptAS3, &pt1)
			if err != nil {return nil, errors.New("Failed to Unmarshal Part #" + key)}
			pt1.Owner = "SCRAPPING_MANAGER"
			pt1.Responsible = "SCRAPPING_MANAGER"
			pt1.PN = ""
			pt1.Helicopter = ""
			pt1.Assembly = ""
		var tx3 Log
			tx3.Owner 		= pt1.Owner
			tx3.Responsible = pt1.Responsible
			tx3.VDate 		= args [1]
			tx3.LType 		= "SCRAPPING"
			pt1.Logs = append(pt1.Logs, tx3)
		f:= UpdatePart (stub, pt1) 
			if f != nil { return nil, errors.New(f.Error())}
			i++
		}
		
return nil, nil

}

// =========================
// Exchange a defective part with another 
// =========================
func (t *SimpleChaincode)ReplacePartOnAssembly(stub shim.ChaincodeStubInterface, args []string)([]byte, error) {
	key :=  args[0]  // L'id de l'Assembly
	idpart := args[1] // L'id de l'ancien Part 
	idpart1 := args[2] // L'id du nouveau part 
// Debut Partie Assembly 
		ac,err:=findAssemblyById(stub,key)
		if(err !=nil){return nil,err}
		ptAS1, _ := json.Marshal(ac)
	var airc Assembly
		err = json.Unmarshal(ptAS1, &airc)
		if err != nil {return nil, errors.New("Failed to Unmarshal Part #" + key)}
		airc.Parts = append(airc.Parts, idpart1)	
	for i, v := range airc.Parts{
			if v == idpart {
				airc.Parts = append(airc.Parts[:i], airc.Parts[i+1:]...)
			break
		}
			}
	var tx LogAssembly
		tx.Owner 		= airc.Owner
		tx.LType 		= "PART_SUBSTITUTION : " + idpart1 +  " replace " + idpart
		tx.VDate		= args [3]
		airc.Logs = append(airc.Logs, tx)
	y:= UpdateAssembly (stub, airc) 
		if y != nil { return nil, errors.New(y.Error())}
// Fin Partie Assembly 
// Debut Partie Part	
		part,err:=findPartById(stub,idpart)
		if err != nil {return nil, errors.New("Failed to get part #" + key)}
		ptAS, _ := json.Marshal(part)
	var pt Part
		err = json.Unmarshal(ptAS, &pt)
		if err != nil {return nil, errors.New("Failed to Unmarshal Part #" + key)}
		
	partt,err:=findPartById(stub,idpart1)
		if err != nil {return nil, errors.New("Failed to get part #" + key)}
		ptASS, _ := json.Marshal(partt)
	var ptt Part
		err = json.Unmarshal(ptASS, &ptt)
		if err != nil {return nil, errors.New("Failed to Unmarshal Part #" + key)}
		ptt.Assembly = key
		ptt.Owner = airc.Owner  // Le champ Helicopter de la part rajoutée à l'A/C prend la valeur A/C.
		ptt.Responsible = pt.Responsible
		ptt.Helicopter = pt.Helicopter
		ptt.PN = pt.PN
	var tff Log
		tff.Responsible = ptt.Responsible
		tff.Owner 		= ptt.Owner
		tff.LType 		= "ADDED TO A/C " + key + " AND SUBSTITUTES PART: " + idpart
		tff.VDate 		= args [3]
		ptt.Logs = append(ptt.Logs, tff)
	r:= UpdatePart (stub, ptt) 
		if r != nil { return nil, errors.New(r.Error())}
		
		pt.Assembly = ""  // Le champ Assembly de la part retirée de l'Assembly revient à nul.
		pt.Owner = "REMOVED_MANAGER"  
		pt.Responsible = "REMOVED_MANAGER"
	var tf Log
		tf.Responsible  = pt.Responsible
		tf.Owner 		= pt.Owner
		tf.LType 		= "REMOVED FROM A/C " + key + " SUBSTITUTED BY PART: " + idpart1
		tf.VDate 		= args [3]
		pt.Logs = append(pt.Logs, tf)
	e:= UpdatePart (stub, pt) 
		if e != nil { return nil, errors.New(e.Error())}
		
// Fin Partie Part 

fmt.Println("Responsible created successfully")	
return nil, nil
}



// =========================
// GET  
// =========================

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
