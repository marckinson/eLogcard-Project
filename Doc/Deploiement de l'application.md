https://bitbucket.org/capblockchain/blockchain/src/9d71e245fdb7d7563fb0f33874fff2605580f71a/README.md?at=master&fileviewer=file-view-default


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





