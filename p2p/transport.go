package p2p 


// Peer is an interface that represents the remote node 
type Peer interface {}

// Transport is something that handles the communication 
// between nodes in a network. This can be of the form 
//(TCP, UDP, websockets)
type Transport interface {}