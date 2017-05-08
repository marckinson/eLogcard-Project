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


#### Secured methods

- #### CreatePart

Create parts with specific criteria (see Body line)

| Criteria      | Value                                                                                                                   |
| --------------| ------------------------------------------------------------------------------------------------------------------------------|
| Host 			| http://localhost:3000                               				      						  				            |
| Method		| POST   			   				            |
| URI		    | /logcard/parts			   				            |
| Header   	 	| 					            |
| Parameters   	|         		                |
| Body		    | { "pn":"", "sn":"", "partName":"", "type":"", "responsible":"", "helicopter":"", "assembly":""}              		                |

- #### Transfer Ownership 

Transfer Ownership on part by informing the new owner name 

| Criteria      | Value                                                                                                                   |
| --------------| ------------------------------------------------------------------------------------------------------------------------------|
| Host 			| http://localhost:3000                               				      						  				            |
| Method		| PUT   			   				            |
| URI		    | /parts/:id/:action				   				            |
| Header   	 	| 					            |
| Parameters   	|  action : OwnerTransfer        		                |
| Body		    | { "owner":"" }             		                |

- #### Transfer Responsibility 

Transfer Responsibility on part by informing the new responsible name 

| Criteria      | Value                                                                                                                   |
| --------------| ------------------------------------------------------------------------------------------------------------------------------|
| Host 			| http://localhost:3000                               				      						  				            |
| Method		| PUT   			   				            |
| URI		    | /parts/:id/:action				   				            |
| Header   	 	| 					            |
| Parameters   	|  action: RespoTransfer       		                |
| Body		    | {	"responsible":"" }              		                |

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

- #### Display part Historic 

Display the lifecycle of a part by informing its ID

| Criteria      | Value                                                                                                                   |
| --------------| ------------------------------------------------------------------------------------------------------------------------------|
| Host 			| http://localhost:3000                               				      						  				            || Method		    			   				            |
| Method		| GET  			   				            |
| URI		    | /parts/:id/historic				   				            |
| Header   	 	| 					            |
| Parameters   	|      		                |
| Body		    |               		                |


- #### Display all parts 

Display the lifecycle of all created parts 

| Criteria      | Value                                                                                                                   |
| --------------| ------------------------------------------------------------------------------------------------------------------------------|
| Host 			| http://localhost:3000                               				      						  				            |
| Method		| GET    			   				            |
| URI		    | /logcard/parts				   				            |
| Header   	 	| 					            |
| Parameters   	|         		                |
| Body		    |               		                |