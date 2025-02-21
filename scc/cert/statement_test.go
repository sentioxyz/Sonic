package cert

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"testing"

	"github.com/0xsoniclabs/sonic/scc"
	"github.com/0xsoniclabs/sonic/scc/bls"
	"github.com/ethereum/go-ethereum/common"
	"github.com/stretchr/testify/require"
)

func TestStatement_getSignData_EncodesChainIdAndDocumentType(t *testing.T) {
	tests := map[string]struct {
		chainId  uint64
		docType  string
		expected []byte
	}{
		"zero chain id no doc": {
			chainId:  0,
			docType:  "",
			expected: concat(byte(0), "scc_", uint64(0)),
		},
		"chain id 12 no doc": {
			chainId:  12,
			docType:  "",
			expected: concat(byte(0), "scc_", uint64(12)),
		},
		"chain id 12 with doc": {
			chainId:  12,
			docType:  "test",
			expected: concat(byte(0), "scc_test", uint64(12)),
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			s := statement{ChainId: test.chainId}
			got := s.getDataToSign(test.docType)
			require.Equal(t, test.expected, got)
		})
	}

}

func TestBlockStatement_GetSignData_HasStatementPrefix(t *testing.T) {
	statement := BlockStatement{}
	got := statement.GetDataToSign()
	prefix := concat(byte(StatementEncodingVersion), "scc_bs", uint64(0))
	require.True(t, bytes.HasPrefix(got, prefix))
}

func TestBlockStatement_GetSignData_EncodesBlockInformation(t *testing.T) {
	statement := BlockStatement{
		Number:    123,
		Hash:      [32]byte{1, 2, 3},
		StateRoot: [32]byte{4, 5, 6},
	}
	got := statement.GetDataToSign()
	want := concat(
		byte(0), "scc_bs", uint64(0),
		statement.Number,
		statement.Hash,
		statement.StateRoot,
	)
	require.Equal(t, want, got)
}

func TestCommitteeStatement_GetSignData_HasStatementPrefix(t *testing.T) {
	statement := CommitteeStatement{}
	got := statement.GetDataToSign()
	prefix := concat(byte(StatementEncodingVersion), "scc_cs", uint64(0))
	require.True(t, bytes.HasPrefix(got, prefix))
}

func TestCommitteeStatement_GetSignData_EncodesCommitteeInformation(t *testing.T) {
	key := bls.NewPrivateKey()
	committee, err := scc.NewCommittee(scc.Member{
		PublicKey:         key.PublicKey(),
		ProofOfPossession: key.GetProofOfPossession(),
		VotingPower:       12,
	})
	require.NoError(t, err)

	statement := CommitteeStatement{
		Period:    123,
		Committee: committee,
	}
	got := statement.GetDataToSign()
	want := concat(
		byte(0), "scc_cs", uint64(0),
		statement.Period,
		statement.Committee.Serialize(),
	)
	require.Equal(t, want, got)
}

func concat(elements ...any) []byte {
	res := []byte{}
	for _, e := range elements {
		switch v := e.(type) {
		case byte:
			res = append(res, v)
		case uint64:
			res = binary.BigEndian.AppendUint64(res, v)
		case string:
			res = append(res, []byte(v)...)
		case common.Hash:
			res = append(res, v.Bytes()...)
		case []byte:
			res = append(res, v...)
		default:
			panic(fmt.Sprintf("unsupported type: %T", v))
		}
	}
	return res
}
