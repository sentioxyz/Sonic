package proxy

import (
	"github.com/sirupsen/logrus"

	"github.com/Fantom-foundation/go-lachesis/src/inter"
	"github.com/Fantom-foundation/go-lachesis/src/poset"
)

// inmemAppProxy implements the AppProxy interface.
type inmemAppProxy struct {
	logger           *logrus.Logger
	handler          App
	submitCh         chan []byte
	submitInternalCh chan inter.InternalTransaction
}

// NewInmemAppProxy instantiates an InmemProxy from a set of handlers.
func NewInmemAppProxy(handler App, logger *logrus.Logger) AppProxy {
	if logger == nil {
		logger = logrus.New()
		logger.Level = logrus.DebugLevel
	}

	return &inmemAppProxy{
		logger:           logger,
		handler:          handler,
		submitCh:         make(chan []byte),
		submitInternalCh: make(chan inter.InternalTransaction),
	}
}

/*
 * AppProxy implementation:
 */

func (p *inmemAppProxy) Close() {
}

func (p *inmemAppProxy) SubmitCh() chan []byte {
	return p.submitCh
}

func (p *inmemAppProxy) SubmitInternalCh() chan inter.InternalTransaction {
	return p.submitInternalCh
}

func (p *inmemAppProxy) CommitBlock(block poset.Block) ([]byte, error) {
	stateHash, err := p.handler.CommitHandler(block)
	p.logger.WithFields(logrus.Fields{
		"round_received": block.RoundReceived(),
		"txs":            len(block.Transactions()),
		"state_hash":     stateHash,
		"err":            err,
	}).Debug("inmemAppProxy.CommitBlock")
	return stateHash, err
}

func (p *inmemAppProxy) GetSnapshot(blockIndex int64) ([]byte, error) {
	snapshot, err := p.handler.SnapshotHandler(blockIndex)
	p.logger.WithFields(logrus.Fields{
		"block":    blockIndex,
		"snapshot": snapshot,
		"err":      err,
	}).Debug("inmemAppProxy.GetSnapshot")
	return snapshot, err
}

func (p *inmemAppProxy) Restore(snapshot []byte) error {
	stateHash, err := p.handler.RestoreHandler(snapshot)
	p.logger.WithFields(logrus.Fields{
		"state_hash": stateHash,
		"err":        err,
	}).Debug("inmemAppProxy.Restore")
	return err
}

/*
 * staff:
 */

// SubmitTx is called by the App to submit a transaction to Lachesis
func (p *inmemAppProxy) SubmitTx(tx []byte) {
	//have to make a copy, or the tx will be garbage collected and weird stuff
	//happens in transaction pool
	t := make([]byte, len(tx))
	copy(t, tx)
	p.submitCh <- t
}
