// Code generated by MockGen. DO NOT EDIT.
// Source: backend.go
//
// Generated by this command:
//
//	mockgen -source=backend.go -destination=mock_backend.go -package=ethapi
//

// Package ethapi is a generated GoMock package.
package ethapi

import (
	context "context"
	iter "iter"
	big "math/big"
	reflect "reflect"
	time "time"

	evmcore "github.com/0xsoniclabs/sonic/evmcore"
	inter "github.com/0xsoniclabs/sonic/inter"
	iblockproc "github.com/0xsoniclabs/sonic/inter/iblockproc"
	state "github.com/0xsoniclabs/sonic/inter/state"
	scc "github.com/0xsoniclabs/sonic/scc"
	cert "github.com/0xsoniclabs/sonic/scc/cert"
	result "github.com/0xsoniclabs/sonic/utils/result"
	hash "github.com/Fantom-foundation/lachesis-base/hash"
	idx "github.com/Fantom-foundation/lachesis-base/inter/idx"
	accounts "github.com/ethereum/go-ethereum/accounts"
	common "github.com/ethereum/go-ethereum/common"
	core "github.com/ethereum/go-ethereum/core"
	types "github.com/ethereum/go-ethereum/core/types"
	vm "github.com/ethereum/go-ethereum/core/vm"
	event "github.com/ethereum/go-ethereum/event"
	params "github.com/ethereum/go-ethereum/params"
	rpc "github.com/ethereum/go-ethereum/rpc"
	gomock "go.uber.org/mock/gomock"
)

// MockBackend is a mock of Backend interface.
type MockBackend struct {
	ctrl     *gomock.Controller
	recorder *MockBackendMockRecorder
}

// MockBackendMockRecorder is the mock recorder for MockBackend.
type MockBackendMockRecorder struct {
	mock *MockBackend
}

// NewMockBackend creates a new mock instance.
func NewMockBackend(ctrl *gomock.Controller) *MockBackend {
	mock := &MockBackend{ctrl: ctrl}
	mock.recorder = &MockBackendMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockBackend) EXPECT() *MockBackendMockRecorder {
	return m.recorder
}

// AccountManager mocks base method.
func (m *MockBackend) AccountManager() *accounts.Manager {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AccountManager")
	ret0, _ := ret[0].(*accounts.Manager)
	return ret0
}

// AccountManager indicates an expected call of AccountManager.
func (mr *MockBackendMockRecorder) AccountManager() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AccountManager", reflect.TypeOf((*MockBackend)(nil).AccountManager))
}

// BlockByHash mocks base method.
func (m *MockBackend) BlockByHash(ctx context.Context, hash common.Hash) (*evmcore.EvmBlock, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "BlockByHash", ctx, hash)
	ret0, _ := ret[0].(*evmcore.EvmBlock)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// BlockByHash indicates an expected call of BlockByHash.
func (mr *MockBackendMockRecorder) BlockByHash(ctx, hash any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BlockByHash", reflect.TypeOf((*MockBackend)(nil).BlockByHash), ctx, hash)
}

// BlockByNumber mocks base method.
func (m *MockBackend) BlockByNumber(ctx context.Context, number rpc.BlockNumber) (*evmcore.EvmBlock, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "BlockByNumber", ctx, number)
	ret0, _ := ret[0].(*evmcore.EvmBlock)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// BlockByNumber indicates an expected call of BlockByNumber.
func (mr *MockBackendMockRecorder) BlockByNumber(ctx, number any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BlockByNumber", reflect.TypeOf((*MockBackend)(nil).BlockByNumber), ctx, number)
}

// CalcBlockExtApi mocks base method.
func (m *MockBackend) CalcBlockExtApi() bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CalcBlockExtApi")
	ret0, _ := ret[0].(bool)
	return ret0
}

// CalcBlockExtApi indicates an expected call of CalcBlockExtApi.
func (mr *MockBackendMockRecorder) CalcBlockExtApi() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CalcBlockExtApi", reflect.TypeOf((*MockBackend)(nil).CalcBlockExtApi))
}

// ChainConfig mocks base method.
func (m *MockBackend) ChainConfig() *params.ChainConfig {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ChainConfig")
	ret0, _ := ret[0].(*params.ChainConfig)
	return ret0
}

// ChainConfig indicates an expected call of ChainConfig.
func (mr *MockBackendMockRecorder) ChainConfig() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ChainConfig", reflect.TypeOf((*MockBackend)(nil).ChainConfig))
}

// CurrentBlock mocks base method.
func (m *MockBackend) CurrentBlock() *evmcore.EvmBlock {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CurrentBlock")
	ret0, _ := ret[0].(*evmcore.EvmBlock)
	return ret0
}

// CurrentBlock indicates an expected call of CurrentBlock.
func (mr *MockBackendMockRecorder) CurrentBlock() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CurrentBlock", reflect.TypeOf((*MockBackend)(nil).CurrentBlock))
}

// CurrentEpoch mocks base method.
func (m *MockBackend) CurrentEpoch(ctx context.Context) idx.Epoch {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CurrentEpoch", ctx)
	ret0, _ := ret[0].(idx.Epoch)
	return ret0
}

// CurrentEpoch indicates an expected call of CurrentEpoch.
func (mr *MockBackendMockRecorder) CurrentEpoch(ctx any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CurrentEpoch", reflect.TypeOf((*MockBackend)(nil).CurrentEpoch), ctx)
}

// EnumerateBlockCertificates mocks base method.
func (m *MockBackend) EnumerateBlockCertificates(first idx.Block) iter.Seq[result.T[cert.BlockCertificate]] {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "EnumerateBlockCertificates", first)
	ret0, _ := ret[0].(iter.Seq[result.T[cert.BlockCertificate]])
	return ret0
}

// EnumerateBlockCertificates indicates an expected call of EnumerateBlockCertificates.
func (mr *MockBackendMockRecorder) EnumerateBlockCertificates(first any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "EnumerateBlockCertificates", reflect.TypeOf((*MockBackend)(nil).EnumerateBlockCertificates), first)
}

// EnumerateCommitteeCertificates mocks base method.
func (m *MockBackend) EnumerateCommitteeCertificates(first scc.Period) iter.Seq[result.T[cert.CommitteeCertificate]] {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "EnumerateCommitteeCertificates", first)
	ret0, _ := ret[0].(iter.Seq[result.T[cert.CommitteeCertificate]])
	return ret0
}

// EnumerateCommitteeCertificates indicates an expected call of EnumerateCommitteeCertificates.
func (mr *MockBackendMockRecorder) EnumerateCommitteeCertificates(first any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "EnumerateCommitteeCertificates", reflect.TypeOf((*MockBackend)(nil).EnumerateCommitteeCertificates), first)
}

// ExtRPCEnabled mocks base method.
func (m *MockBackend) ExtRPCEnabled() bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ExtRPCEnabled")
	ret0, _ := ret[0].(bool)
	return ret0
}

// ExtRPCEnabled indicates an expected call of ExtRPCEnabled.
func (mr *MockBackendMockRecorder) ExtRPCEnabled() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ExtRPCEnabled", reflect.TypeOf((*MockBackend)(nil).ExtRPCEnabled))
}

// GetDowntime mocks base method.
func (m *MockBackend) GetDowntime(ctx context.Context, vid idx.ValidatorID) (idx.Block, inter.Timestamp, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetDowntime", ctx, vid)
	ret0, _ := ret[0].(idx.Block)
	ret1, _ := ret[1].(inter.Timestamp)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// GetDowntime indicates an expected call of GetDowntime.
func (mr *MockBackendMockRecorder) GetDowntime(ctx, vid any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetDowntime", reflect.TypeOf((*MockBackend)(nil).GetDowntime), ctx, vid)
}

// GetEVM mocks base method.
func (m *MockBackend) GetEVM(ctx context.Context, msg *core.Message, state vm.StateDB, header *evmcore.EvmHeader, vmConfig *vm.Config, blockContext *vm.BlockContext) (*vm.EVM, func() error, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetEVM", ctx, msg, state, header, vmConfig, blockContext)
	ret0, _ := ret[0].(*vm.EVM)
	ret1, _ := ret[1].(func() error)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// GetEVM indicates an expected call of GetEVM.
func (mr *MockBackendMockRecorder) GetEVM(ctx, msg, state, header, vmConfig, blockContext any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetEVM", reflect.TypeOf((*MockBackend)(nil).GetEVM), ctx, msg, state, header, vmConfig, blockContext)
}

// GetEpochBlockState mocks base method.
func (m *MockBackend) GetEpochBlockState(ctx context.Context, epoch rpc.BlockNumber) (*iblockproc.BlockState, *iblockproc.EpochState, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetEpochBlockState", ctx, epoch)
	ret0, _ := ret[0].(*iblockproc.BlockState)
	ret1, _ := ret[1].(*iblockproc.EpochState)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// GetEpochBlockState indicates an expected call of GetEpochBlockState.
func (mr *MockBackendMockRecorder) GetEpochBlockState(ctx, epoch any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetEpochBlockState", reflect.TypeOf((*MockBackend)(nil).GetEpochBlockState), ctx, epoch)
}

// GetEvent mocks base method.
func (m *MockBackend) GetEvent(ctx context.Context, shortEventID string) (*inter.Event, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetEvent", ctx, shortEventID)
	ret0, _ := ret[0].(*inter.Event)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetEvent indicates an expected call of GetEvent.
func (mr *MockBackendMockRecorder) GetEvent(ctx, shortEventID any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetEvent", reflect.TypeOf((*MockBackend)(nil).GetEvent), ctx, shortEventID)
}

// GetEventPayload mocks base method.
func (m *MockBackend) GetEventPayload(ctx context.Context, shortEventID string) (*inter.EventPayload, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetEventPayload", ctx, shortEventID)
	ret0, _ := ret[0].(*inter.EventPayload)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetEventPayload indicates an expected call of GetEventPayload.
func (mr *MockBackendMockRecorder) GetEventPayload(ctx, shortEventID any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetEventPayload", reflect.TypeOf((*MockBackend)(nil).GetEventPayload), ctx, shortEventID)
}

// GetHeads mocks base method.
func (m *MockBackend) GetHeads(ctx context.Context, epoch rpc.BlockNumber) (hash.Events, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetHeads", ctx, epoch)
	ret0, _ := ret[0].(hash.Events)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetHeads indicates an expected call of GetHeads.
func (mr *MockBackendMockRecorder) GetHeads(ctx, epoch any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetHeads", reflect.TypeOf((*MockBackend)(nil).GetHeads), ctx, epoch)
}

// GetLatestBlockCertificate mocks base method.
func (m *MockBackend) GetLatestBlockCertificate() (cert.BlockCertificate, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetLatestBlockCertificate")
	ret0, _ := ret[0].(cert.BlockCertificate)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetLatestBlockCertificate indicates an expected call of GetLatestBlockCertificate.
func (mr *MockBackendMockRecorder) GetLatestBlockCertificate() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetLatestBlockCertificate", reflect.TypeOf((*MockBackend)(nil).GetLatestBlockCertificate))
}

// GetLatestCommitteeCertificate mocks base method.
func (m *MockBackend) GetLatestCommitteeCertificate() (cert.CommitteeCertificate, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetLatestCommitteeCertificate")
	ret0, _ := ret[0].(cert.CommitteeCertificate)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetLatestCommitteeCertificate indicates an expected call of GetLatestCommitteeCertificate.
func (mr *MockBackendMockRecorder) GetLatestCommitteeCertificate() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetLatestCommitteeCertificate", reflect.TypeOf((*MockBackend)(nil).GetLatestCommitteeCertificate))
}

// GetOriginatedFee mocks base method.
func (m *MockBackend) GetOriginatedFee(ctx context.Context, vid idx.ValidatorID) (*big.Int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetOriginatedFee", ctx, vid)
	ret0, _ := ret[0].(*big.Int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetOriginatedFee indicates an expected call of GetOriginatedFee.
func (mr *MockBackendMockRecorder) GetOriginatedFee(ctx, vid any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetOriginatedFee", reflect.TypeOf((*MockBackend)(nil).GetOriginatedFee), ctx, vid)
}

// GetPoolNonce mocks base method.
func (m *MockBackend) GetPoolNonce(ctx context.Context, addr common.Address) (uint64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetPoolNonce", ctx, addr)
	ret0, _ := ret[0].(uint64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetPoolNonce indicates an expected call of GetPoolNonce.
func (mr *MockBackendMockRecorder) GetPoolNonce(ctx, addr any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPoolNonce", reflect.TypeOf((*MockBackend)(nil).GetPoolNonce), ctx, addr)
}

// GetPoolTransaction mocks base method.
func (m *MockBackend) GetPoolTransaction(txHash common.Hash) *types.Transaction {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetPoolTransaction", txHash)
	ret0, _ := ret[0].(*types.Transaction)
	return ret0
}

// GetPoolTransaction indicates an expected call of GetPoolTransaction.
func (mr *MockBackendMockRecorder) GetPoolTransaction(txHash any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPoolTransaction", reflect.TypeOf((*MockBackend)(nil).GetPoolTransaction), txHash)
}

// GetPoolTransactions mocks base method.
func (m *MockBackend) GetPoolTransactions() (types.Transactions, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetPoolTransactions")
	ret0, _ := ret[0].(types.Transactions)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetPoolTransactions indicates an expected call of GetPoolTransactions.
func (mr *MockBackendMockRecorder) GetPoolTransactions() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPoolTransactions", reflect.TypeOf((*MockBackend)(nil).GetPoolTransactions))
}

// GetReceiptsByNumber mocks base method.
func (m *MockBackend) GetReceiptsByNumber(ctx context.Context, number rpc.BlockNumber) (types.Receipts, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetReceiptsByNumber", ctx, number)
	ret0, _ := ret[0].(types.Receipts)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetReceiptsByNumber indicates an expected call of GetReceiptsByNumber.
func (mr *MockBackendMockRecorder) GetReceiptsByNumber(ctx, number any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetReceiptsByNumber", reflect.TypeOf((*MockBackend)(nil).GetReceiptsByNumber), ctx, number)
}

// GetTransaction mocks base method.
func (m *MockBackend) GetTransaction(ctx context.Context, txHash common.Hash) (*types.Transaction, uint64, uint64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetTransaction", ctx, txHash)
	ret0, _ := ret[0].(*types.Transaction)
	ret1, _ := ret[1].(uint64)
	ret2, _ := ret[2].(uint64)
	ret3, _ := ret[3].(error)
	return ret0, ret1, ret2, ret3
}

// GetTransaction indicates an expected call of GetTransaction.
func (mr *MockBackendMockRecorder) GetTransaction(ctx, txHash any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTransaction", reflect.TypeOf((*MockBackend)(nil).GetTransaction), ctx, txHash)
}

// GetUptime mocks base method.
func (m *MockBackend) GetUptime(ctx context.Context, vid idx.ValidatorID) (*big.Int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUptime", ctx, vid)
	ret0, _ := ret[0].(*big.Int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUptime indicates an expected call of GetUptime.
func (mr *MockBackendMockRecorder) GetUptime(ctx, vid any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUptime", reflect.TypeOf((*MockBackend)(nil).GetUptime), ctx, vid)
}

// HeaderByHash mocks base method.
func (m *MockBackend) HeaderByHash(ctx context.Context, hash common.Hash) (*evmcore.EvmHeader, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "HeaderByHash", ctx, hash)
	ret0, _ := ret[0].(*evmcore.EvmHeader)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// HeaderByHash indicates an expected call of HeaderByHash.
func (mr *MockBackendMockRecorder) HeaderByHash(ctx, hash any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "HeaderByHash", reflect.TypeOf((*MockBackend)(nil).HeaderByHash), ctx, hash)
}

// HeaderByNumber mocks base method.
func (m *MockBackend) HeaderByNumber(ctx context.Context, number rpc.BlockNumber) (*evmcore.EvmHeader, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "HeaderByNumber", ctx, number)
	ret0, _ := ret[0].(*evmcore.EvmHeader)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// HeaderByNumber indicates an expected call of HeaderByNumber.
func (mr *MockBackendMockRecorder) HeaderByNumber(ctx, number any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "HeaderByNumber", reflect.TypeOf((*MockBackend)(nil).HeaderByNumber), ctx, number)
}

// MaxGasLimit mocks base method.
func (m *MockBackend) MaxGasLimit() uint64 {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "MaxGasLimit")
	ret0, _ := ret[0].(uint64)
	return ret0
}

// MaxGasLimit indicates an expected call of MaxGasLimit.
func (mr *MockBackendMockRecorder) MaxGasLimit() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "MaxGasLimit", reflect.TypeOf((*MockBackend)(nil).MaxGasLimit))
}

// MinGasPrice mocks base method.
func (m *MockBackend) MinGasPrice() *big.Int {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "MinGasPrice")
	ret0, _ := ret[0].(*big.Int)
	return ret0
}

// MinGasPrice indicates an expected call of MinGasPrice.
func (mr *MockBackendMockRecorder) MinGasPrice() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "MinGasPrice", reflect.TypeOf((*MockBackend)(nil).MinGasPrice))
}

// Progress mocks base method.
func (m *MockBackend) Progress() PeerProgress {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Progress")
	ret0, _ := ret[0].(PeerProgress)
	return ret0
}

// Progress indicates an expected call of Progress.
func (mr *MockBackendMockRecorder) Progress() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Progress", reflect.TypeOf((*MockBackend)(nil).Progress))
}

// RPCEVMTimeout mocks base method.
func (m *MockBackend) RPCEVMTimeout() time.Duration {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RPCEVMTimeout")
	ret0, _ := ret[0].(time.Duration)
	return ret0
}

// RPCEVMTimeout indicates an expected call of RPCEVMTimeout.
func (mr *MockBackendMockRecorder) RPCEVMTimeout() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RPCEVMTimeout", reflect.TypeOf((*MockBackend)(nil).RPCEVMTimeout))
}

// RPCGasCap mocks base method.
func (m *MockBackend) RPCGasCap() uint64 {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RPCGasCap")
	ret0, _ := ret[0].(uint64)
	return ret0
}

// RPCGasCap indicates an expected call of RPCGasCap.
func (mr *MockBackendMockRecorder) RPCGasCap() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RPCGasCap", reflect.TypeOf((*MockBackend)(nil).RPCGasCap))
}

// RPCTxFeeCap mocks base method.
func (m *MockBackend) RPCTxFeeCap() float64 {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RPCTxFeeCap")
	ret0, _ := ret[0].(float64)
	return ret0
}

// RPCTxFeeCap indicates an expected call of RPCTxFeeCap.
func (mr *MockBackendMockRecorder) RPCTxFeeCap() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RPCTxFeeCap", reflect.TypeOf((*MockBackend)(nil).RPCTxFeeCap))
}

// ResolveRpcBlockNumberOrHash mocks base method.
func (m *MockBackend) ResolveRpcBlockNumberOrHash(ctx context.Context, blockNrOrHash rpc.BlockNumberOrHash) (idx.Block, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ResolveRpcBlockNumberOrHash", ctx, blockNrOrHash)
	ret0, _ := ret[0].(idx.Block)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ResolveRpcBlockNumberOrHash indicates an expected call of ResolveRpcBlockNumberOrHash.
func (mr *MockBackendMockRecorder) ResolveRpcBlockNumberOrHash(ctx, blockNrOrHash any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ResolveRpcBlockNumberOrHash", reflect.TypeOf((*MockBackend)(nil).ResolveRpcBlockNumberOrHash), ctx, blockNrOrHash)
}

// SealedEpochTiming mocks base method.
func (m *MockBackend) SealedEpochTiming(ctx context.Context) (inter.Timestamp, inter.Timestamp) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SealedEpochTiming", ctx)
	ret0, _ := ret[0].(inter.Timestamp)
	ret1, _ := ret[1].(inter.Timestamp)
	return ret0, ret1
}

// SealedEpochTiming indicates an expected call of SealedEpochTiming.
func (mr *MockBackendMockRecorder) SealedEpochTiming(ctx any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SealedEpochTiming", reflect.TypeOf((*MockBackend)(nil).SealedEpochTiming), ctx)
}

// SendTx mocks base method.
func (m *MockBackend) SendTx(ctx context.Context, signedTx *types.Transaction) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SendTx", ctx, signedTx)
	ret0, _ := ret[0].(error)
	return ret0
}

// SendTx indicates an expected call of SendTx.
func (mr *MockBackendMockRecorder) SendTx(ctx, signedTx any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SendTx", reflect.TypeOf((*MockBackend)(nil).SendTx), ctx, signedTx)
}

// StateAndHeaderByNumberOrHash mocks base method.
func (m *MockBackend) StateAndHeaderByNumberOrHash(ctx context.Context, blockNrOrHash rpc.BlockNumberOrHash) (state.StateDB, *evmcore.EvmHeader, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "StateAndHeaderByNumberOrHash", ctx, blockNrOrHash)
	ret0, _ := ret[0].(state.StateDB)
	ret1, _ := ret[1].(*evmcore.EvmHeader)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// StateAndHeaderByNumberOrHash indicates an expected call of StateAndHeaderByNumberOrHash.
func (mr *MockBackendMockRecorder) StateAndHeaderByNumberOrHash(ctx, blockNrOrHash any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StateAndHeaderByNumberOrHash", reflect.TypeOf((*MockBackend)(nil).StateAndHeaderByNumberOrHash), ctx, blockNrOrHash)
}

// Stats mocks base method.
func (m *MockBackend) Stats() (int, int) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Stats")
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(int)
	return ret0, ret1
}

// Stats indicates an expected call of Stats.
func (mr *MockBackendMockRecorder) Stats() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Stats", reflect.TypeOf((*MockBackend)(nil).Stats))
}

// SubscribeNewTxsNotify mocks base method.
func (m *MockBackend) SubscribeNewTxsNotify(arg0 chan<- evmcore.NewTxsNotify) event.Subscription {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SubscribeNewTxsNotify", arg0)
	ret0, _ := ret[0].(event.Subscription)
	return ret0
}

// SubscribeNewTxsNotify indicates an expected call of SubscribeNewTxsNotify.
func (mr *MockBackendMockRecorder) SubscribeNewTxsNotify(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SubscribeNewTxsNotify", reflect.TypeOf((*MockBackend)(nil).SubscribeNewTxsNotify), arg0)
}

// SuggestGasTipCap mocks base method.
func (m *MockBackend) SuggestGasTipCap(ctx context.Context, certainty uint64) *big.Int {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SuggestGasTipCap", ctx, certainty)
	ret0, _ := ret[0].(*big.Int)
	return ret0
}

// SuggestGasTipCap indicates an expected call of SuggestGasTipCap.
func (mr *MockBackendMockRecorder) SuggestGasTipCap(ctx, certainty any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SuggestGasTipCap", reflect.TypeOf((*MockBackend)(nil).SuggestGasTipCap), ctx, certainty)
}

// TxPoolContent mocks base method.
func (m *MockBackend) TxPoolContent() (map[common.Address]types.Transactions, map[common.Address]types.Transactions) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "TxPoolContent")
	ret0, _ := ret[0].(map[common.Address]types.Transactions)
	ret1, _ := ret[1].(map[common.Address]types.Transactions)
	return ret0, ret1
}

// TxPoolContent indicates an expected call of TxPoolContent.
func (mr *MockBackendMockRecorder) TxPoolContent() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "TxPoolContent", reflect.TypeOf((*MockBackend)(nil).TxPoolContent))
}

// TxPoolContentFrom mocks base method.
func (m *MockBackend) TxPoolContentFrom(addr common.Address) (types.Transactions, types.Transactions) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "TxPoolContentFrom", addr)
	ret0, _ := ret[0].(types.Transactions)
	ret1, _ := ret[1].(types.Transactions)
	return ret0, ret1
}

// TxPoolContentFrom indicates an expected call of TxPoolContentFrom.
func (mr *MockBackendMockRecorder) TxPoolContentFrom(addr any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "TxPoolContentFrom", reflect.TypeOf((*MockBackend)(nil).TxPoolContentFrom), addr)
}

// UnprotectedAllowed mocks base method.
func (m *MockBackend) UnprotectedAllowed() bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UnprotectedAllowed")
	ret0, _ := ret[0].(bool)
	return ret0
}

// UnprotectedAllowed indicates an expected call of UnprotectedAllowed.
func (mr *MockBackendMockRecorder) UnprotectedAllowed() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UnprotectedAllowed", reflect.TypeOf((*MockBackend)(nil).UnprotectedAllowed))
}
