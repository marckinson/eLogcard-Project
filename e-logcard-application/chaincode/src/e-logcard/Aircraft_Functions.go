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
		air.AircraftName = args[3]
		air.Owner = args [4] // Owner of the Aircraft 
		air.Responsible = args [5] 
	var tx LogAircraft
		tx.Owner 		= air.Owner
		tx.Responsible  = air.Responsible
		tx.VDate 		= args[6]
		tx.LType 		= "CREATION"
		tx.Description  = args [4] + " Created This Aircraft" 
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

// =========================================================================================
// 					ACTIVITIES ON AIRCRAFTS
// =========================================================================================

// ====================================================================
// addPartToAc (Parts qui n'appartiennent à aucun Assembly )
// ====================================================================
func (t *SimpleChaincode)addPartToAc(stub shim.ChaincodeStubInterface, args []string)([]byte, error) {

	key :=  args[0]
	idpart := args[1]
// Verification
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
	var tx LogAircraft
		tx.Responsible  = airc.Responsible
		tx.Owner 		= airc.Owner
		tx.LType 		= "PART_AFFILIATION "
		tx.Description  =  idpart + " has been affiliated to this Aircraft "
		tx.VDate        = args [2] // Fonctionne 
		airc.Parts = append(airc.Parts, idpart)	
		airc.Logs = append(airc.Logs, tx)
	y:= UpdateAircraft (stub, airc) 
		if y != nil { return nil, errors.New(y.Error())}
// Fin Partie Aircraft 
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
		tf.Responsible 	= pt.Responsible
		tf.LType 		= "AIRCRAFT_AFFILIATION"
		tf.VDate        = args [2] // Fonctionne 
		tf.Description  = "This part has been attached to :" + key
		pt.Logs = append(pt.Logs, tf)
	e:= UpdatePart (stub, pt) 
		if e != nil { return nil, errors.New(e.Error())}
// Fin Partie Part
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
	var tx LogAircraft
		tx.Owner 		= airc.Owner
		tx.Responsible 		= airc.Responsible
		tx.LType 		= "PART_REMOVAL " 
		tx.VDate		= args [2]
		tx.Description  = idpart + " has been removed from this A/C "
		airc.Logs = append(airc.Logs, tx)
	y:= UpdateAircraft (stub, airc) 
		if y != nil { return nil, errors.New(y.Error())}
// Fin Partie Aircraft 
// Debut Partie Part	
		part,err:=findPartById(stub,idpart)
		if err != nil {return nil, errors.New("Failed to get part #" + key)}
		ptAS, _ := json.Marshal(part)
	var pt Part
		err = json.Unmarshal(ptAS, &pt)
		if err != nil {return nil, errors.New("Failed to Unmarshal Part #" + key)}
		pt.Helicopter = "" // Le champ Helicopter de la part retirée de l'Helicopter revient à nul.
		pt.Owner 	  = "REMOVAL_MANAGER"
		pt.Responsible = "REMOVAL_MANAGER"
	var tf Log
		tf.Owner 		= pt.Owner
		tf.Responsible	= pt.Responsible
		tf.LType 		= "AIRCRAFT_REMOVAL"
		tf.VDate		= args [2]
		tf.Description  = "This part has been removed from A/C: " + key + " This part has been transfered to " + pt.Responsible + ", the new Owner & the new Responsible."
		pt.Logs = append(pt.Logs, tf)
	e:= UpdatePart (stub, pt) 
		if e != nil { return nil, errors.New(e.Error())}
// Fin Partie Part
return nil, nil
}

// =========================
// Exchange a defective part with another 
// =========================
func (t *SimpleChaincode)ReplacePartOnAircraft(stub shim.ChaincodeStubInterface, args []string)([]byte, error) {
	key :=  args[0]  // L'id de l'A/C
	idpart := args[1] // L'id de l'ancien Part 
	idpart1 := args[2] // L'id du nouveau part 
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
	var tx LogAircraft
		tx.Responsible  = airc.Responsible
		tx.Owner 		= airc.Owner
		tx.LType 		= "PART_SUBSTITUTION " 
		tx.VDate		= args [3]
		tx.Description  =  idpart1 +  " replace " + idpart + " on this A/C "
		airc.Logs = append(airc.Logs, tx)
	y:= UpdateAircraft (stub, airc) 
		if y != nil { return nil, errors.New(y.Error())}
// Fin Partie Aircraft 
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
		ptt.Helicopter = key
		ptt.Owner = airc.Owner  // Le champ Helicopter de la part rajoutée à l'A/C prend la valeur A/C.
		ptt.Responsible = pt.Responsible
		ptt.Assembly = pt.Assembly
		ptt.PN = pt.PN
	var tff Log
		tff.Responsible = ptt.Responsible
		tff.Owner 		= ptt.Owner
		tff.LType 		= "AIRCRAFT_AFFILIATION"
		tff.VDate 		= args [3]
		tff.Description = "AFFILIATED TO A/C " + key + " AND SUBSTITUTES PART: " + idpart
		ptt.Logs = append(ptt.Logs, tff)
	r:= UpdatePart (stub, ptt) 
		if r != nil { return nil, errors.New(r.Error())}
		
		
		pt.Helicopter = ""  // Le champ Helicopter de la part retirée de l'Helicopter revient à nul.
		pt.Owner = "REMOVAL_MANAGER"  
		pt.Responsible = "REMOVAL_MANAGER"
	var tf Log
		tf.Responsible  = pt.Responsible
		tf.Owner 		= pt.Owner
		tf.LType 		= "AIRCRAFT_REMOVAL"
		tf.VDate 		= args [3]
		tf.Description  = "REMOVED FROM A/C " + key + " SUBSTITUTED BY PART: " + idpart1 + " This part has been transfered to " + pt.Responsible + ", the new Owner & the new Responsible."
		pt.Logs = append(pt.Logs, tf)
	e:= UpdatePart (stub, pt) 
		if e != nil { return nil, errors.New(e.Error())}
	
// Fin Partie Part 

fmt.Println("Responsible created successfully")	
return nil, nil
}
// =========================
// Add an assembly to An Aircraft 
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
	var tx LogAircraft
		tx.Responsible  = airc.Responsible
		tx.Owner 		= airc.Owner
		tx.LType 		= "ASSEMBLY_AFFILIATION "
		tx.VDate		= args [2]
		tx.Description  =  idassembly + " has been affiliated to this Aircraft "
		airc.Assemblies = append(airc.Assemblies, idassembly)	
		airc.Logs = append(airc.Logs, tx)
	d:= UpdateAircraft (stub, airc) 
		if d != nil { return nil, errors.New(d.Error())}
// Fin partie Aircraft 

// Debut Partie Assembly	
		part,err:=findAssemblyById(stub,idassembly)
		if err != nil {return nil, errors.New("Failed to get part #" + key)}
		ptASS, _ := json.Marshal(part)
	var pt Assembly
		err = json.Unmarshal(ptASS, &pt)
		if err != nil {return nil, errors.New("Failed to Unmarshal Part #" + key)}
		pt.Helicopter = key
		pt.Owner = airc.Owner // Le champ Helicopter de l'Assembly ajouté à l'Helicopter prend la valeur A/C
	var tf LogAssembly
		tf.Responsible 		= pt.Responsible
		tf.Owner 		= pt.Owner
		tf.LType 		= "AIRCRAFT_AFFILIATION "
		tf.VDate		= args [2]
		tf.Description  =  "This Assembly has been affiliated to Aircraft: " + key
		pt.Logs = append(pt.Logs, tf)
	y:= UpdateAssembly (stub, pt) 
		if y != nil { return nil, errors.New(y.Error())}
// Fin Partie Assembly 

// Debut Partie Part 	
		for i := range pt.Parts{
			part,err:=findPartById(stub,pt.Parts[i])
			if err != nil {return nil, errors.New("Failed to get part #" + key)}
			ptAS, _ := json.Marshal(part)
		var pt1 Part
			err = json.Unmarshal(ptAS, &pt1)
			if err != nil {return nil, errors.New("Failed to Unmarshal Part #" + key)}
			pt1.Helicopter = key
			pt1.Owner = airc.Owner
		var tx1 Log
			tx1.Owner		= pt1.Owner
			tx1.Owner 		= pt1.Responsible
			tx1.LType 		= "AIRCRAFT AFFILIATION: " 
			tx1.Description = "This part has been affiliated to aircraft: " + key 
			tx1.VDate		= args [2]
			pt1.Logs = append(pt1.Logs, tx1)
			
		e:= UpdatePart (stub, pt1) 
		if e != nil { return nil, errors.New(e.Error())}
			i++
		} 	
} else if (ppart.Helicopter != "") {  // Un assembly appartenant à un A/C ne peut pas être ajouté à un A/C
		return nil, errors.New ("Impossible") }

fmt.Println("Responsible created successfully")	
return nil, nil
}
// =========================
// Remove an Assembly From An Aircraft
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
	var tx LogAircraft
		tx.Responsible  = airc.Responsible
		tx.Owner 		= airc.Owner
		tx.LType 		= "ASSEMBLY_REMOVAL:"
		tx.Description  =  idassembly + "has been removed from this A/C"
		tx.VDate		= args [2]
		airc.Logs = append(airc.Logs, tx)
	y:= UpdateAircraft (stub, airc) 
		if y != nil { return nil, errors.New(y.Error())}
// Fin Partie Aircraft
// Debut Partie Assembly	
		part,err:=findAssemblyById(stub,idassembly)
		if err != nil {return nil, errors.New("Failed to get part #" + key)}
		ptASS, _ := json.Marshal(part)
	var pt Assembly
		err = json.Unmarshal(ptASS, &pt)
		if err != nil {return nil, errors.New("Failed to Unmarshal Part #" + key)}
		pt.Helicopter = "" // Le champ Helicopter de l'Assembly retirée de l'Helicopter revient à nul.
		pt.Owner = "REMOVAL_MANAGER"
		pt.Responsible = "REMOVAL_MANAGER"
	var tf LogAssembly
		tf.Responsible 		= pt.Responsible
		tf.Owner 		= pt.Owner
		tf.LType 		= "AIRCRAFT REMOVAL"
		tf.Description  = "REMOVED FROM A/C: " + key
		tf.VDate		= args [2]
		pt.Logs = append(pt.Logs, tf)
	e:= UpdateAssembly (stub, pt) 
		if e != nil { return nil, errors.New(e.Error())}
// Fin Partie Assembly 
// Debut Partie Part 	
		for i := range pt.Parts{
			part,err:=findPartById(stub,pt.Parts[i])
			if err != nil {return nil, errors.New("Failed to get part #" + key)}
			ptAS, _ := json.Marshal(part)
		var pt1 Part
			err = json.Unmarshal(ptAS, &pt1)
			if err != nil {return nil, errors.New("Failed to Unmarshal Part #" + key)}
			pt1.Helicopter = ""
			pt1.Owner = "REMOVAL_MANAGER"
			pt1.Responsible = "REMOVAL_MANAGER"
		var tx1 Log
			tx1.Owner		= pt1.Owner
			tx1.Responsible = pt1.Responsible
			tx1.LType 		= "AIRCRAFT REMOVAL"
			tx1.Description = "This part has been removed from Aircraft: " + key + " This part has been transfered to " + pt1.Responsible + ", the new Owner & the new Responsible."
			tx1.VDate		= args [2]
			pt1.Logs = append(pt1.Logs, tx1)
		e:= UpdatePart (stub, pt1) 
		if e != nil { return nil, errors.New(e.Error())}
			i++
		} 	

fmt.Println("Responsible created successfully")	
return nil, nil
}
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
	var tx LogAircraft
		tx.Responsible  = airc.Responsible
		tx.Owner 		= airc.Owner
		tx.LType 		= "OWNERNSHIP_TRANSFER"
		tx.Description  = " This Aircraft has been transfered to: " + airc.Owner
		tx.VDate		= args[2]
		airc.Logs = append(airc.Logs, tx)
	y:= UpdateAircraft (stub, airc) 
		if y != nil { return nil, errors.New(y.Error())}
// Fin Partie Aircraft 
// Debut Partie Part 	
		for i := range airc.Parts{
			part,err:=findPartById(stub,airc.Parts[i])
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
			tx.Description  = "This part has been transfered to " + pt.Owner + ", the new Owner." 
			tx.VDate		= args[2]
			pt.Logs = append(pt.Logs, tx)
		e:= UpdatePart (stub, pt) 
		if e != nil { return nil, errors.New(e.Error())}
			i++
		}
// Fin PArtie Part 

// Début Partie Assembly 
	for i := range airc.Assemblies{
			part2,err:=findAssemblyById(stub,airc.Assemblies[i])
			if err != nil {return nil, errors.New("Failed to get part2 #" + key)}
			ptAS2, _ := json.Marshal(part2)
		var assemb Assembly
			err = json.Unmarshal(ptAS2, &assemb)
			if err != nil {return nil, errors.New("Failed to Unmarshal Part #" + key)}
			assemb.Owner = args[1]
		var tx2 LogAssembly
			tx2.Responsible 		= assemb.Responsible
			tx2.Owner 		= assemb.Owner
			tx2.VDate 		=  args [2]
			tx2.LType 		= "OWNERNSHIP_TRANSFER"
			tx2.Description  = "This assembly has been transfered to " + assemb.Owner + ", the new Owner." 

			assemb.Logs = append(assemb.Logs, tx2)
		e:= UpdateAssembly (stub, assemb) 
		if e != nil { return nil, errors.New(e.Error())}
		
		// Debut Partie Part affilié à Assembly	
		for i := range assemb.Parts{
			part3,err:=findPartById(stub,assemb.Parts[i])
			if err != nil {return nil, errors.New("Failed to get part3 #" + key)}
			ptAS3, _ := json.Marshal(part3)
		var pt1 Part
			err = json.Unmarshal(ptAS3, &pt1)
			if err != nil {return nil, errors.New("Failed to Unmarshal Part #" + key)}
			pt1.Owner = args [1]
		var tx3 Log
			tx3.Owner 		= pt1.Owner
			tx3.Responsible = pt1.Responsible
			tx3.VDate 		=  args [2]
			tx3.LType 		= "OWNERNSHIP_TRANSFER"
			pt1.Logs = append(pt1.Logs, tx3)
		f:= UpdatePart (stub, pt1) 
			if f != nil { return nil, errors.New(f.Error())}
			i++
		}
			i++
		}	
// Fin Partie Assembly 
		
	return nil, nil
}

// =========================
// Transfert de responsabilité  
// =========================
func (t *SimpleChaincode) AcRespoTransfer(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {
	var err error
	key := args[0]
// Debut Partie Aircraft 
		ac,err:=findAircraftById(stub,key)
		if(err !=nil){return nil,err}
		ptAS1, _ := json.Marshal(ac)
	var airc Aircraft
		err = json.Unmarshal(ptAS1, &airc)
		if err != nil {return nil, errors.New("Failed to Unmarshal Part #" + key)}
		airc.Responsible = args[1] 
	var tx LogAircraft
		tx.Responsible 		= airc.Responsible
		tx.LType 		= "RESPONSIBLE_TRANSFER"
		tx.Description  = " This Aircraft has been transfered to: " + airc.Responsible
		tx.VDate		= args[2]
		airc.Logs = append(airc.Logs, tx)
	y:= UpdateAircraft (stub, airc) 
		if y != nil { return nil, errors.New(y.Error())}
// Fin Partie Aircraft 
// Debut Partie Part 	
		for i := range airc.Parts{
			part,err:=findPartById(stub,airc.Parts[i])
			if err != nil {return nil, errors.New("Failed to get part #" + key)}
			ptAS, _ := json.Marshal(part)
		var pt Part
			err = json.Unmarshal(ptAS, &pt)
			if err != nil {return nil, errors.New("Failed to Unmarshal Part #" + key)}
			pt.Responsible = args[1]
		var tx Log
			tx.Responsible	= pt.Responsible
			tx.Owner 		= pt.Owner
			tx.LType 		= "RESPONSIBLE_TRANSFER"
			tx.Description  = "This part has been transfered to " + pt.Responsible + ", the new Responsible." 
			tx.VDate		= args[2]
			pt.Logs = append(pt.Logs, tx)
		e:= UpdatePart (stub, pt) 
		if e != nil { return nil, errors.New(e.Error())}
			i++
		}
// Fin PArtie Part 

// Début Partie Assembly 
	for i := range airc.Assemblies{
			part2,err:=findAssemblyById(stub,airc.Assemblies[i])
			if err != nil {return nil, errors.New("Failed to get part2 #" + key)}
			ptAS2, _ := json.Marshal(part2)
		var assemb Assembly
			err = json.Unmarshal(ptAS2, &assemb)
			if err != nil {return nil, errors.New("Failed to Unmarshal Part #" + key)}
			assemb.Responsible = args[1]
		var tx2 LogAssembly
			tx2.Responsible = assemb.Responsible
			tx2.VDate 		=  args [2]
			tx2.LType 		= "RESPONSIBLE_TRANSFER"
			tx2.Description  = "This assembly has been transfered to " + assemb.Responsible + ", the new Responsible." 

			assemb.Logs = append(assemb.Logs, tx2)
		e:= UpdateAssembly (stub, assemb) 
		if e != nil { return nil, errors.New(e.Error())}
		
		// Debut Partie Part affilié à Assembly	
		for i := range assemb.Parts{
			part3,err:=findPartById(stub,assemb.Parts[i])
			if err != nil {return nil, errors.New("Failed to get part3 #" + key)}
			ptAS3, _ := json.Marshal(part3)
		var pt1 Part
			err = json.Unmarshal(ptAS3, &pt1)
			if err != nil {return nil, errors.New("Failed to Unmarshal Part #" + key)}
			pt1.Responsible = args [1]
		var tx3 Log
			tx3.Owner 		= pt1.Owner
			tx3.Responsible = pt1.Responsible
			tx3.VDate 		=  args [2]
			tx3.LType 		= "RESPONSIBLE_TRANSFER"
			pt1.Logs = append(pt1.Logs, tx3)
		f:= UpdatePart (stub, pt1) 
			if f != nil { return nil, errors.New(f.Error())}
			i++
		}
			i++
		}	
// Fin Partie Assembly 
		
	return nil, nil
}
// =========================
// Scrapp an Aircraft  
// =========================
func (t *SimpleChaincode) scrappAircraft(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {

	var err error
	var key string 
	key = args[0]
	part,err:=findAircraftById(stub,key)
		if err != nil {return nil, errors.New("Failed to get part #" + key)}
		ptAS, _ := json.Marshal(part)
	var airc Aircraft
		err = json.Unmarshal(ptAS, &airc)
		if err != nil {return nil, errors.New("Failed to Unmarshal Part #" + key)}
		airc.Owner = "SCRAPPING_MANAGER"
		airc.Responsible = "SCRAPPING_MANAGER"
		airc.AN = ""
	var tx LogAircraft
		tx.Responsible  = airc.Responsible
		tx.Owner 		= airc.Owner
		tx.VDate 		=  args [1]
		tx.LType 		= "SCRAPPING"
		tx.Description  = "This Aircraft has been scrapped and has been transfered to " + airc.Owner + ", the new Owner." 

		airc.Logs = append(airc.Logs, tx)
	
	e:= UpdateAircraft (stub, airc) 
		if e != nil { return nil, errors.New(e.Error())}

	// Parts 
	for i := range airc.Parts{
		part1,err:=findPartById(stub,airc.Parts[i])
			if err != nil {return nil, errors.New("Failed to get part1 #" + key)}
			ptAS1, _ := json.Marshal(part1)
		var pt Part
			err = json.Unmarshal(ptAS1, &pt)
			if err != nil {return nil, errors.New("Failed to Unmarshal Part #" + key)}
			pt.Owner = "SCRAPPING_MANAGER"
			pt.Responsible = "SCRAPPING_MANAGER"
			pt.PN = ""
			pt.Helicopter = ""
			pt.Assembly = ""
		var tx1 Log
			tx1.Owner 		= pt.Owner
			tx1.Responsible = pt.Responsible
			tx1.VDate 		=  args [1]
			tx1.LType 		= "SCRAPPING"
			tx1.Description = "This part has been scrapped and has been transfered to " + pt.Responsible + ", the new Owner & Responsible." 

			pt.Logs = append(pt.Logs, tx1)
		e:= UpdatePart (stub, pt) 
			if e != nil { return nil, errors.New(e.Error())}
				i++
		}
	
	// Assembly
		for i := range airc.Assemblies{
			part2,err:=findAssemblyById(stub,airc.Assemblies[i])
			if err != nil {return nil, errors.New("Failed to get part2 #" + key)}
			ptAS2, _ := json.Marshal(part2)
		var assemb Assembly
			err = json.Unmarshal(ptAS2, &assemb)
			if err != nil {return nil, errors.New("Failed to Unmarshal Part #" + key)}
			assemb.Owner = "SCRAPPING_MANAGER"
			assemb.Responsible = "SCRAPPING_MANAGER"
			assemb.AN = ""
			assemb.Helicopter = ""
		var tx2 LogAssembly
			tx2.Responsible = assemb.Responsible
			tx2.Owner 		= assemb.Owner
			tx2.VDate 		=  args [1]
			tx2.LType 		= "SCRAPPING"
			tx2.Description  = "This Assembly has been scrapped and has been transfered to " + assemb.Owner + ", the new Owner." 

			assemb.Logs = append(assemb.Logs, tx2)
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
			tx3.Responsible 	= pt1.Responsible
			tx3.VDate 		=  args [1]
			tx3.LType 		= "SCRAPPING"
			tx3.Description = "This part has been scrapped and has been transfered to " + pt1.Responsible + ", the new Owner & the new Responsible." 
			pt1.Logs = append(pt1.Logs, tx3)
		f:= UpdatePart (stub, pt1) 
			if f != nil { return nil, errors.New(f.Error())}
			
			i++
		}
		
		i++
	}

	return nil, nil
	
}

// ====================================================================
// Replace Assembly on Part  
// ====================================================================
func (t *SimpleChaincode) replaceAssemblyOnAC(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {

	key :=  args[0] 			     // L'id de l'A/C
	idassembly := args[1] 			 // L'id de l'ancien Assembly 
	idassembly1 := args[2] 			 // L'id du nouvel Assembly 

// Debut Partie Aircraft 
		ac,err:=findAircraftById(stub,key)
		if(err !=nil){return nil,err}
		ptAS1, _ := json.Marshal(ac)
	var airc Aircraft
		err = json.Unmarshal(ptAS1, &airc)
		if err != nil {return nil, errors.New("Failed to Unmarshal Part #" + key)}
		airc.Assemblies = append(airc.Assemblies, idassembly1)	
	for i, v := range airc.Assemblies{
			if v == idassembly {
				airc.Assemblies = append(airc.Assemblies[:i], airc.Assemblies[i+1:]...)
			break
		}
			}
	var tx LogAircraft
		tx.Owner 		= airc.Owner
		tx.Responsible  = airc.Responsible
		tx.LType 		= "ASSEMBLY_SUBSTITUTION"
		tx.VDate		= args [3]
		tx.Description  = idassembly1 +  " replace " + idassembly
		airc.Logs = append(airc.Logs, tx)
	y:= UpdateAircraft (stub, airc) 
		if y != nil { return nil, errors.New(y.Error())}
// Fin Partie Aircraft 


// Debut Partie Part		

// Assembly retiré 
	part,err:=findAssemblyById(stub,idassembly)
	if err != nil {return nil, errors.New("Failed to get part #" + key)}
	ptAS, _ := json.Marshal(part)	
	var pt Assembly
		err = json.Unmarshal(ptAS, &pt)
		if err != nil {return nil, errors.New("Failed to Unmarshal Part #" + key)}
// Assembly rajouté 
	partt,err:=findAssemblyById(stub,idassembly1)
	if err != nil {return nil, errors.New("Failed to get part #" + key)}
	ptASS, _ := json.Marshal(partt)
	var ptt Assembly
		err = json.Unmarshal(ptASS, &ptt)
		if err != nil {return nil, errors.New("Failed to Unmarshal Part #" + key)}
		ptt.Helicopter = key
		ptt.Owner = pt.Owner  
		ptt.AN = pt.AN
	var tff LogAssembly
		tff.Owner 		= ptt.Owner
		tff.LType 		= "AIRCRAFT_AFFILIATION"
		tff.Description = "AFFILIATED TO A/C " + key + " AND SUBSTITUTES ASSEMBLY: " + idassembly
		tff.VDate 		= args [3]
		ptt.Logs = append(ptt.Logs, tff)
	r:= UpdateAssembly (stub, ptt) 
		if r != nil { return nil, errors.New(r.Error())}
// Parts affiliés à l'Assembly rajouté 
	for i := range ptt.Parts{
		part,err:=findPartById(stub,ptt.Parts[i])
			if err != nil {return nil, errors.New("Failed to get part #" + key)}
			ptAS, _ := json.Marshal(part)
		var pt2 Part
			err = json.Unmarshal(ptAS, &pt2)
			if err != nil {return nil, errors.New("Failed to Unmarshal Part #" + key)}
		pt2.Owner = ptt.Owner
		pt2.Helicopter = key
		var tx Log
			tx.Responsible	= pt2.Responsible
			tx.Owner 		= pt2.Owner
			tx.LType 		= "AIRCRAFT_AFFILIATION"
			tx.Description  = "AFFILIATED TO A/C " + key + " BECAUSE " + idassembly1 + " HAS BEEN ADDED TO A/C " + key 
			tx.VDate  		= args [3]
			pt2.Logs = append(pt2.Logs, tx)
	e:= UpdatePart (stub, pt2) 
		if e != nil { return nil, errors.New(e.Error())}
			i++
		}
		
	
// Assembly retiré  
		pt.Helicopter = ""  // Le champ Helicopter de la part retirée de l'Helicopter revient à nul.
		pt.Owner = "REMOVAL_MANAGER"  
		pt.Responsible = "REMOVAL_MANAGER"  
	var tf LogAssembly
		tf.Owner 		= pt.Owner
		tf.LType 		= "AIRCRAFT_REMOVAL"
		tf.Description  = "REMOVED FROM A/C " + key + " SUBSTITUTED BY ASSEMBLY: " + idassembly1
		tf.VDate 		= args [3]
		pt.Logs = append(pt.Logs, tf)
	e:= UpdateAssembly (stub, pt) 
		if e != nil { return nil, errors.New(e.Error())}
// Parts affiliés à l'Assembly retiré 
	for i := range pt.Parts{
		part,err:=findPartById(stub,pt.Parts[i])
			if err != nil {return nil, errors.New("Failed to get part #" + key)}
			ptAS, _ := json.Marshal(part)
		var pt2 Part
			err = json.Unmarshal(ptAS, &pt2)
			if err != nil {return nil, errors.New("Failed to Unmarshal Part #" + key)}
		pt2.Owner = "REMOVAL_MANAGER" 
		pt2.Responsible = "REMOVAL_MANAGER"  
		pt2.Helicopter = "" 
		var tx Log
			tx.Responsible	= pt2.Responsible
			tx.Owner 		= pt2.Owner
			tx.LType 		= "AIRCRAFT_REMOVAL"
			tx.Description  = "REMOVED FROM A/C " + key + " BECAUSE ASSEMBLY " + idassembly + " HAS BEEN REMOVED FROM A/C " + key +  " This assembly has been transfered to " + pt2.Owner + ", the new Owner "
			tx.VDate  		= args [3]
			pt2.Logs = append(pt2.Logs, tx)
	e:= UpdatePart (stub, pt2) 
		if e != nil { return nil, errors.New(e.Error())}
			i++
		}
// Fin Assembly retiré

fmt.Println("Responsible created successfully")	
return nil, nil
}








// ====================================================================
// GET 
// ====================================================================

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
		} else if(showOnlyMyPart && part.Helicopter == key && part.Responsible == username){
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

// ====================================================================
// Afficher la liste détailéles de tous les assmemblies composants un Aircraft donné à partir de son id
// ====================================================================
func (t *SimpleChaincode)AcAssembliesListing(stub shim.ChaincodeStubInterface, args []string)([]byte, error) {

	key := args [0]
	username, err := getAttribute(stub, "username")
	if(err !=nil){return nil,err}
	
	role, err := getAttribute(stub, "role")
		if(err !=nil){return nil,err}
	//if supplier or manufacturer or customer or maintenance user =>only my parts
	showOnlyMyPart := role=="supplier" || role == "manufacturer" || role == "customer" || role == "maintenance_user"

	partMap,err:=getAssemblyMap(stub)
	if err != nil {return nil, errors.New("Failed to get Part")}
	parts := make([]Assembly, len(partMap))
	idx := 0
    for  _, part := range partMap {
    	if(showOnlyMyPart && part.Helicopter == key && part.Owner == username){
    		parts[idx] = part
    		idx++
		} else if(showOnlyMyPart && part.Helicopter == key && part.Responsible == username){
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

	partMap,err:=getAircraftMap(stub)
	if(err !=nil){return nil,err}
	parts := make([]Aircraft, len(partMap))
    idx := 0
    for  _, part := range partMap {
		if(!showOnlyMyPart || part.Owner == username || part.Responsible == username){
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