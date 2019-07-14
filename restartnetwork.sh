docker exec -it cli bash
CORE_PEER_MSPCONFIGPATH=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/upm.tietonetworks.com/users/Admin@upm.tietonetworks.com/msp
CORE_PEER_LOCALMSPID="UpmMSP"
CORE_PEER_TLS_ROOTCERT_FILE=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/upm.tietonetworks.com/peers/peer1.upm.tietonetworks.com/tls/ca.crt
CORE_PEER_TLS_ENABLED=true
CORE_PEER_ADDRESS=peer1.upm.tietonetworks.com:7051
CORE_PEER_CHAINCODELISTENADDRESS=peer1.upm.tietonetworks.com:7052
export CHANNEL_NAME=upmorgchannel
export ORDERER_CA=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/ordererOrganizations/tietonetworks.com/orderers/orderer.tietonetworks.com/msp/tlscacerts/tlsca.tietonetworks.com-cert.pem
export CC_SRC_PATH="github.com/chaincode/"

peer chaincode install -n doctors_chain -v 0 -l golang -p  ${CC_SRC_PATH}
