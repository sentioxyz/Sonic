package scc

import (
	"fmt"
	"slices"
	"strings"

	"github.com/0xsoniclabs/sonic/scc/bls"
)

// Committee is a group of Members that can sign statements to produce
// certificates.
//
// On the Sonic Certification Chain (SCC), time is divided into periods. For
// each period, a committee is authorized to sign the statement for that period.
// Among those statements is the confirmation of the succeeding period's
// committee.
//
// Committees are composed of an ordered list of Members, each weighted with a
// non-zero voting power. Members are identified by their public keys, for which
// they are required to provide a proof of possession.
type Committee struct {
	members []Member
}

// MaxCommitteeSize is the maximum number of members that can be in a committee.
// The number needs to be limited to prevent Committee certificates from becoming
// too large.
const MaxCommitteeSize = 512

// MemberId is used to identify a member in a committee by its position in the
// ordered list of members.
type MemberId uint16

// NewCommittee creates a new committee with the provided members. There must be
// at least one member, all members must be valid, and no duplicates are allowed.
func NewCommittee(members ...Member) (Committee, error) {
	res := Committee{members: slices.Clone(members)}
	if err := res.Validate(); err != nil {
		return Committee{}, err
	}
	return res, nil
}

// Members returns a copy of the members in the committee.
func (c Committee) Members() []Member {
	return slices.Clone(c.members)
}

// GetMember returns the member with the given id. If the id is out of bounds,
// the second return value is false.
func (c Committee) GetMember(id MemberId) (Member, bool) {
	if int(id) >= len(c.members) {
		return Member{}, false
	}
	return c.members[id], true
}

// GetMemberId returns the id of the member with the given public key. If the
// public key is not found, the second return value is false.
func (c Committee) GetMemberId(publicKey bls.PublicKey) (MemberId, bool) {
	for id, m := range c.members {
		if m.PublicKey == publicKey {
			return MemberId(id), true
		}
	}
	return 0, false
}

// Validate checks that the committee is well-formed. For a committee to be well
// formed, the following properties need to be satisfied:
// - The committee must have at least one member.
// - The committee size must not exceed the maximum committee size.
// - All members must be valid.
// - No two members can have the same public key.
// - The sum of the voting power must not exceed 2^64 - 1.
// If any of these properties are violated, an error is returned.
func (c Committee) Validate() error {
	if len(c.members) == 0 {
		return fmt.Errorf("committee must have at least one member")
	}

	if len(c.members) > MaxCommitteeSize {
		return fmt.Errorf("committee size exceeds the maximum of %d", MaxCommitteeSize)
	}

	for _, m := range c.members {
		if err := m.Validate(); err != nil {
			return fmt.Errorf("invalid member %v, %w", m, err)
		}
	}

	for i, a := range c.members {
		for j, b := range c.members {
			if i != j && a.PublicKey == b.PublicKey {
				return fmt.Errorf("duplicate members at indexes %d and %d", i, j)
			}
		}
	}

	sum := uint64(0)
	for _, m := range c.members {
		next := sum + m.VotingPower
		if next < sum {
			return fmt.Errorf("voting power overflow")
		}
		sum = next
	}

	return nil
}

// String produces a human-readable summary of the Committee information mainly
// for debugging purposes. The output is not sufficient to reconstruct the
// committee.
func (c Committee) String() string {
	result := strings.Builder{}
	result.WriteString("Committee{")
	for i, m := range c.Members() {
		key := m.PublicKey.Serialize()
		result.WriteString(
			fmt.Sprintf(
				"%d: 0x%x..%x -> %d, ",
				i, key[:2], key[len(key)-2:],
				m.VotingPower,
			),
		)
	}
	result.WriteString(fmt.Sprintf("Valid: %t}", c.Validate() == nil))
	return result.String()
}

// Serialize serializes the committee into a byte slice. The serialization format
// is a concatenation of the serialized members. Note that also invalid committees
// can be serialized.
func (c Committee) Serialize() []byte {
	if len(c.members) == 0 {
		return nil
	}
	res := make([]byte, 0, len(c.members)*len(EncodedMember{}))
	for _, m := range c.Members() {
		cur := m.Serialize()
		res = append(res, cur[:]...)
	}
	return res
}

// DeserializeCommittee deserializes a committee from a byte slice. An error is
// returned if the provided data does not contain a valid encoding of a
// committee. Note, this function does not validate members nor the committee.
// Thus, it is possible to deserialize invalid committees.
func DeserializeCommittee(data []byte) (Committee, error) {
	if len(data) == 0 {
		return Committee{}, nil
	}
	if len(data)%len(EncodedMember{}) != 0 {
		return Committee{}, fmt.Errorf("invalid committee data length")
	}

	members := make([]Member, 0, len(data)/len(EncodedMember{}))
	for len(data) > 0 {
		var m Member
		m, err := DeserializeMember(*(*EncodedMember)(data))
		if err != nil {
			return Committee{}, fmt.Errorf("invalid member, %w", err)
		}
		members = append(members, m)
		data = data[len(EncodedMember{}):]
	}

	return Committee{members}, nil
}
