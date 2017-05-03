https://bitbucket.org/capblockchain/blockchain/src/9d71e245fdb7d7563fb0f33874fff2605580f71a/README.md?at=master&fileviewer=file-view-default



=================================================================================================================================================
Configuration
=================================================================================================================================================

//===============================================================================================================================================
// Pré Requis
//===============================================================================================================================================

=================================================================================================================================================
Source code 
=================================================================================================================================================
Peut-être penser à mettre mon code directement sur mon github ? pour pouvoir recuperer le code.
git clone https://[Nom_d'utilisateur]@bitbucket.org/capblockchain/blockchain.git
git clone https://Marckinson@bitbucket.org/capblockchain/blockchain.git

=================================================================================================================================================
Système d'exploitation 
=================================================================================================================================================
Windows 10 x64 

=================================================================================================================================================
Docker 
=================================================================================================================================================
Installation: https://docs.docker.com/docker-for-windows
Version de Docker : Docker version 17.03.0-ce, build 60ccb22

=================================================================================================================================================
Fabric 
=================================================================================================================================================
Installation: git clone -b v0.6 https://github.com/hyperledger/fabric.git 
Version de Fabric  v0.6
Documentation v0.6
http://hyperledger-fabric.readthedocs.io/en/v0.6/

=================================================================================================================================================
Node 
=================================================================================================================================================
Installation
version de node : 6.6.0 ( version idéale 6.5.9)

=================================================================================================================================================
Node-gyp 
=================================================================================================================================================
Installation : Commande à effectuer sous windows pour installer node-gyp, Windows you can now install all node-gyp dependencies with single command:
$ npm install --global --production windows-build-tools
$ npm install --global node-gyp
Attention: Peut-être nécessité d'installer Visual Studio C++ 

=================================================================================================================================================
Npm 
=================================================================================================================================================
Installation : La version installée avec node 6.6.0 n'est pas compatible. Taper: npm install -g npm@next pour installer la dernière version de npm
Version de npm recommandée: 4.5.0

=================================================================================================================================================
Docker-compose
=================================================================================================================================================
Installation: automatique avec docker for windows 
version de docker-compose : 1.11.2, build f963d76f

=================================================================================================================================================
Python
=================================================================================================================================================
Installation 
Version de pyhton recommandée : 2.7
Attention à bien rajouter les variables d'environnement notamment avec le path.

=================================================================================================================================================
Postman 
=================================================================================================================================================
Installation: https://www.getpostman.com/



//===============================================================================================================================================
// Installation de l'application
//===============================================================================================================================================


First, clone the repository :
git clone https://[yourUserName]@bitbucket.org/capblockchain/blockchain.git

I- Init Environment

universal-payment : 
In the folder universal-payment you can find a script initEnv.sh. You can launch this script localy to initialize your node context.
This will install vendor for our chaincode and run npm install in your node server automatically. 
This script is also used into our node container to initialize the node context.

II- Startup Peer

a - Tag the fabric image as latest version to be used inside container

# ./start.sh -t
Tag the latest version of fabric-baseimage
x86_64-0.1.0: Pulling from hyperledger/fabric-baseimage
Digest: sha256:ac6a2784cfd028ae62f5688f4436f95d7a60eeacd8506eb303c9c6335328c388
Status: Image is up to date for hyperledger/fabric-baseimage:x86_64-0.1.0

b- Then you can start containers
# ./start.sh
Rebuild composed containers
Building membersrvc
.
.
.

Verify that 2 containers named "docker_vp0" and "docker_membersrvc" are running :
docker ps -a



III- Configuration of Node App





