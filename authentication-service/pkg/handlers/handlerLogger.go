package handlers

import (
	"log"
	"net/http"
	"net/rpc"
)

// logItemViaRPC logs an item by making an RPC call to the logger microservice
func (rep *Repository) LogItemViaRPC(w http.ResponseWriter, l RPCPayload) error {
	// Dial up RPC server
	client, err := rpc.Dial("tcp", "logger-service:5001")
	if err != nil {
		log.Println("Client failed to dial")
		// app.errorJSON(w, err)
		return err
	}
	// Prepare payload to be sent to logger-service
	rpcPayload := RPCPayload{
		Collection: "authentication",
		Name:       l.Name,
		Data:       l.Data,
	}

	// Call RPC server to run LogInfo and return result
	var result string
	err = client.Call("RPCServer.LogInfo", rpcPayload, &result)
	if err != nil {
		log.Println("error is here", err)
		// app.errorJSON(w, err)
		return err
	}

	// payload is the response to be sent back
	// payload := jsonResponse{
	// 	Error:   false,
	// 	Message: result,
	// }
	// app.writeJSON(w, http.StatusAccepted, payload, "")

	log.Println("Mongo logged authentication successfully", result)
	return nil
}
