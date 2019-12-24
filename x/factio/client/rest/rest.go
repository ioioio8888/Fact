package rest

import (
	"fmt"

	"github.com/cosmos/cosmos-sdk/client/context"

	"github.com/gorilla/mux"
)

const (
	restName = "factio"
)

// RegisterRoutes - Central function to define routes that get registered by the main application
func RegisterRoutes(cliCtx context.CLIContext, r *mux.Router, storeName string) {
	r.HandleFunc(fmt.Sprintf("/%s/factlist", storeName), getFactListHandler(cliCtx, storeName)).Methods("GET")
	r.HandleFunc(fmt.Sprintf("/%s/{%s}/getfact", storeName, restName), getFactHandler(cliCtx, storeName)).Methods("GET")

	r.HandleFunc(fmt.Sprintf("/%s/createfact", storeName), CreateFactHandler(cliCtx)).Methods("POST")
	r.HandleFunc(fmt.Sprintf("/%s/editfact", storeName), EditFactHandler(cliCtx)).Methods("PUT")
	r.HandleFunc(fmt.Sprintf("/%s/{%s}/getvotelist", storeName, restName), getVoteListHandler(cliCtx, storeName)).Methods("GET")

	r.HandleFunc(fmt.Sprintf("/%s/{%s}/getvotepower", storeName, restName), getVotePowerHandler(cliCtx, storeName)).Methods("GET")
	r.HandleFunc(fmt.Sprintf("/%s/votefact", storeName), VoteHandler(cliCtx)).Methods("POST")
	r.HandleFunc(fmt.Sprintf("/%s/unvotefact", storeName), UnVoteHandler(cliCtx)).Methods("POST")
}
