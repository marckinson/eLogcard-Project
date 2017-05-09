# e-Logcard Application 

## Description of the APIs

#### Non-secured methods

- #### Registration 

Register new users by informing specific criteria (see body line)

| Criteria      | Value                                                                                                                   |
| --------------| ------------------------------------------------------------------------------------------------------------------------------|
| Host 			| http://localhost:3000                               				      						  				            |
| Method		| POST   			   				            |
| URI		    | /registration			   				            |
| Header   	 	| 					            |
| Parameters   	|         		                |
| Body		    | { "username":"", "password":"", "role":"" }              		                |
| Response   	 	| 	Token 				            |


- #### Login 

Log registered users into the application by informing specific criteria (see body line)

| Criteria      | Value                                                                                                                   |
| --------------| ------------------------------------------------------------------------------------------------------------------------------|
| Host 			| http://localhost:3000                               				      						  				            |
| Method		| POST  			   				            |
| URI		    | /login				   				            |
| Header   	 	| 					            |
| Parameters   	|         		                |
| Body		    | {	"username":"","password":"" }              		                |
| Response   	 	| 	Token 				            |

#### Registration and Login methods return a token. This token has to be passed in every secured methods (necessary) to be able to use them. In the header the  Authorization field has to be fulfilled like this : "Bearer + token"

#### Secured methods

- #### CreatePart

Create parts with specific criteria (see Body line)

| Criteria      | Value                                                                                                                   |
| --------------| ------------------------------------------------------------------------------------------------------------------------------|
| Host 			| http://localhost:3000                               				      						  				            |
| Method		| POST   			   				            |
| URI		    | /logcard/parts			   				            |
| Header   	 	| Bearer + token					            |
| Parameters   	|         		                |
| Body		    | { "pn":"", "sn":"", "partName":"", "type":"", "responsible":"", "helicopter":"", "assembly":""}              		                |
| Response   	 	| 	{ "pn": "", "sn": "", "id": "", "partName": "", "type": "", "owner": "", "responsible": "", "helicopter": "", "assembly": "", "vDate": ""}		     |

- #### Transfer Ownership 

Transfer Ownership on part by informing the new owner name 

| Criteria      | Value                                                                                                                   |
| --------------| ------------------------------------------------------------------------------------------------------------------------------|
| Host 			| http://localhost:3000                               				      						  				            |
| Method		| PUT   			   				            |
| URI		    | /parts/:id/:action				   				            |
| Header   	 	| Bearer + token					            |
| Parameters   	|  action : OwnerTransfer        		                |
| Body		    | { "owner":"" }             		                |
| Response   	 	|  true			            |

- #### Transfer Responsibility 

Transfer Responsibility on part by informing the new responsible name 

| Criteria      | Value                                                                                                                   |
| --------------| ------------------------------------------------------------------------------------------------------------------------------|
| Host 			| http://localhost:3000                               				      						  				            |
| Method		| PUT   			   				            |
| URI		    | /parts/:id/:action				   				            |
| Header   	 	| Bearer + token					            |
| Parameters   	|  action: RespoTransfer       		                |
| Body		    | {	"responsible":"" }              		                |
| Response   	 	| 	true 				            |

- #### Perform Activities 

Perform acts on parts by informing specific criteria

| Criteria      | Value                                                                                                                   |
| --------------| ------------------------------------------------------------------------------------------------------------------------------|
| Host 			| http://localhost:3000                               				      						  				            |
| Method		| PUT   			   				            |
| URI		    | /parts/:id/:action				   				            |
| Header   	 	| 					            |
| Parameters   	|  action:PerformActs       		                |
| Body		    | {	"modType":"", "description": "" }           		                |
| Response   	 	| 		true			            |

- #### Display part Historic 

Display the lifecycle of a part by informing its ID

| Criteria      | Value                                                                                                                   |
| --------------| ------------------------------------------------------------------------------------------------------------------------------|
| Host 			| http://localhost:3000                               				      						  				            || Method		    			   				            |
| Method		| GET  			   				            |
| URI		    | /parts/:id/historic				   				            |
| Header   	 	| Bearer + token					            |
| Parameters   	|      		                |
| Body		    |               		                |
| Response   	| 	[{"pn": "", "sn": "", "id": "", "partName": "", "type": "", "owner": "", "responsible": "", "helicopter": "", "assembly": "","logs": [{"log_type": "", "vDate": "", "owner": "", "responsible": "", "modType": "", "description": ""}] }]				            |


- #### Display all parts 

Display the lifecycle of all created parts 

| Criteria      | Value                                                                                                                   |
| --------------| ------------------------------------------------------------------------------------------------------------------------------|
| Host 			| http://localhost:3000                               				      						  				            |
| Method		| GET    			   				            |
| URI		    | /logcard/parts				   				            |
| Header   	 	| Bearer + token					            |
| Parameters   	|         		                |
| Body		    |               		                |
| Response   	| 	[{"pn": "", "sn": "", "id": "", "partName": "", "type": "", "owner": "", "responsible": "", "helicopter": "", "assembly": "","logs": [{"log_type": "", "vDate": "", "owner": "", "responsible": "", "modType": "", "description": ""}] }]				            |
