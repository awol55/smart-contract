package spvnode

import (
	"context"
	"errors"

	"github.com/tokenized/smart-contract/pkg/wire"
)

// TXHandler exists to handle the Ping command.
type TXHandler struct {
	Config       Config
	BlockService *BlockService
	Listener     Listener
}

// NewTXHandler returns a new TXHandler with the given Config.
func NewTXHandler(config Config,
	blockService *BlockService,
	listener Listener) TXHandler {

	return TXHandler{
		Config:       config,
		BlockService: blockService,
		Listener:     listener,
	}
}

// Handle implments the Handler interface.
//
// This function handles type conversion and delegates the the contrete
// handler.
func (h TXHandler) Handle(ctx context.Context,
	m wire.Message) ([]wire.Message, error) {

	msg, ok := m.(*wire.MsgTx)
	if !ok {
		return nil, errors.New("Could not assert as *wire.MsgTx")
	}

	return h.handle(ctx, msg)
}

// handle processes the MsgTxn.
//
// There is no response for this handler.
func (h TXHandler) handle(ctx context.Context,
	tx *wire.MsgTx) ([]wire.Message, error) {

	if h.Listener != nil {
		// notify the listener
		h.Listener.Handle(ctx, tx)
	}

	return nil, nil
}
