package git

import (
	"encoding/hex"
	"time"
)

// Hash represents a git commit hash (20 bytes) - compatible with plumbing.Hash
type Hash [20]byte

// String returns the hex string representation of the hash
func (h Hash) String() string {
	return hex.EncodeToString(h[:])
}

// NewHash creates a Hash from a hex string
func NewHash(s string) (Hash, error) {
	var h Hash
	bytes, err := hex.DecodeString(s)
	if err != nil {
		return h, err
	}
	copy(h[:], bytes)
	return h, nil
}

// MustHash creates a Hash from a hex string, panicking on error
func MustHash(s string) Hash {
	h, err := NewHash(s)
	if err != nil {
		panic(err)
	}
	return h
}

// Commit represents a git commit - compatible with object.Commit
type Commit struct {
	Hash      Hash
	Message   string
	Author    Signature
	Committer Signature
}

// Signature represents a git signature
type Signature struct {
	Name  string
	Email string
	When  time.Time
}

// Reference represents a git reference - compatible with plumbing.Reference
type Reference struct {
	name string
	hash Hash
}

// Name returns the reference name
func (r *Reference) Name() string {
	return r.name
}

// Hash returns the reference hash
func (r *Reference) Hash() Hash {
	return r.hash
}

// NewReference creates a new Reference
func NewReference(name string, hash Hash) *Reference {
	return &Reference{name: name, hash: hash}
}
