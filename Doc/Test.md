

=================================================================================================================================================
														Test 
=================================================================================================================================================


=================================================================================================================================================
Docker
=================================================================================================================================================
Se placer dans le dossier Docker:

C:\Users\mjeancha\Desktop\Hyperledger\blockchain\Docker

docker-compose up  -d: pour cr�er et lancer les containers 
docker ps -a : pour voir la liste des containers 
docker-compose down : pour supprimer les containers

container dev-vp0-6217ef3790b39d6ff1ed19b3e576da85e301846af41e548e6a855fdf3c83198c : Container cr�e par le peer lors du deploiement du chaincode to execute la chaincode.


=================================================================================================================================================
Application Node Js
=================================================================================================================================================
Pour d�ployer la Chaincode: 
Commandes: 
start_node.sh -s : pour lancer l'appli node Js 
Node app.js pour lancer l'appli node Js ==> on obtient ensuite un Hash a utiliser dans Postman 
localhost:5000 : pour pouvoir interragir avec leur interface web.


=================================================================================================================================================
Test de la Chaincode
=================================================================================================================================================
Attention:  
Modification du chaincode :  
We have to keep it unless you delete the container docker_membersrvc and docker_vp0. 
Sometimes you may encounter problems, and you need to delete KeyValStore, chaincodeID.txt (you can use ./start-node.sh -c) and re-run start.sh (run ./start.sh -r -a before to force all your containers to stop/remove).
Supprimer le fichier KeyValStore : You can find all authentication information in KeyValStore.
Supprimer le container cr�e par le peer. 
Supprimer les fichier chaincodeID.txt. 
Supprimer les deux containers cr�er au debut 

Conseil : Trouver une plateforme d'execution pour tester le code ou du moins des bouts de code sans pour autant refaire toutes ses manipulations.

//=================================================================================================================================================
//	 Application 
//=================================================================================================================================================

T�ches: �crire un script permette d'effectuer le d�marrage de l'application et le d�marrage des container,
ainsi que le d�marrage en dur du container cr�er par le peer. 

Chose � rectifier dans l'installation 
Dans les fichiers default 
vp0 ==> localhost
membersrvc_ ==> localhost 

Composition du dossier:

=================================================================================================================================================
Docker 
=================================================================================================================================================
start.sh : sert � lancer les containers. Attention ne lance pas le container cr�er par le peer.
membersrvc.yml:  Te servira si tu souhaites disposer d�une configuration particuli�re pour la gestion des utilisateur/role etc. 
Nous l�avons faiblement modifi� par rapport � celui du repo git fabric.

=================================================================================================================================================
Postman
=================================================================================================================================================
Cr�er des collections personnalis�e 

=================================================================================================================================================
Universal Payment 
=================================================================================================================================================
Universal-payment est le serveur node.js qui va r�aliser les �changes via HFC avec Hyperledger
start-node.sh: sert � lancer l'application 
