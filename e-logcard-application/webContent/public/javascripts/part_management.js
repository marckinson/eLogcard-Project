var hyperledger=hyperledger||{};


class PartManager{
	
	constructor(){
		this.init();
	}
	
	init(){
		this.registerPartCreationListener();
		this.refreshPartsDisplay();
	}
	
	registerPartCreationListener(){
		$("#part_creation").click(this.onPartCreation);
	}
	
	refreshPartsDisplay(){
		$.getJSON('/logcard/parts',function(data){
			var dataSet = data.parts;
			console.log(data.parts);
			console.log(dataSet);
		    $('#example').DataTable( {	
				
		    data: dataSet,	
		        columns: [
		            { data: "sn" ,title:"Serial Number" },
		            { data: "pn" ,title:"Part Number"},
		            { data: "id" ,title:"Id"},
					{ data: "partName" ,title:"Name"},
					{ data: "type" ,title:"Type"},
					{ data: "owner" ,title:"Owner"},
					{ data: "responsible" ,title:"Responsible"},
					{ data: "helicopter" ,title:"Helicopter"},
					{ data: "assembly" ,title:"Assembly"},
					{ data: "id" ,title:"Actions"},

		        ], 
				
				/*
				columnDefs: [  
					 { targets: 3, data: "sn",title:"Serial Number", defaultContent: "<button>Click!</button>",visible:false} 
				]	
				
				*/
		    } ); 
		});
		
		
	
	}
	onPartCreation(){
		var part={};
		part.pn=$("#createPartPn").val();
		part.sn=$("#createPartSn").val();
		part.id=$("#createPartId").val();
		part.partName=$("#createPartPartName").val();
		part.type=$("#createPartType").val();
		part.owner=$("#createPartOwner").val();
		part.responsible=$("#createPartResponsible").val();
		part.helicopter=$("#createPartHelicopter").val();
		part.assembly=$("#createPartAssembly").val();
		console.log("Creation de la part "+ part);
		
		$.post("/logcard/parts",part,function(data){
			
			console.log("creation succeed");
		    $("#part-creation-dialog" ).dialog({
		        modal: true,
		        buttons: {
		          Ok: function() {
		        	hyperledger.partManager.refreshPartsDisplay();
		            $( this ).dialog( "close" );
		          }
		        }
		      });
		});
	}
};







$(function(){
	
	hyperledger.partManager=new PartManager();
    console.log( $("#part-creation-dialog" ).length);
	
});



