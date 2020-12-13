package v1beta

import (
	"crypto/sha256"
	"encoding/binary"
	"hash/crc32"

	"github.com/depscloud/api/v1beta"
	"github.com/depscloud/api/v1beta/graphstore"

	"github.com/golang/protobuf/proto"
	"github.com/golang/protobuf/ptypes"
)

var moduleKind string
var sourceKind string

func init() {
	m, err := ptypes.MarshalAny(&v1beta.Module{})
	if err != nil {
		panic(err)
	}
	moduleKind = m.GetTypeUrl()

	s, err := ptypes.MarshalAny(&v1beta.Source{})
	if err != nil {
		panic(err)
	}
	sourceKind = s.GetTypeUrl()
}

var sep = "---"
var sepData = []byte(sep)

func generateKey(parts ...string) []byte {
	hash := sha256.New()

	for _, part := range parts {
		data := []byte(part)

		checksum := make([]byte, 4)
		binary.BigEndian.PutUint32(checksum, crc32.ChecksumIEEE(data))

		hash.Write(sepData)
		hash.Write(checksum)
		hash.Write(data)
	}

	return hash.Sum(nil)
}

func newNode(msg proto.Message) (*graphstore.Node, error) {
	any, err := ptypes.MarshalAny(msg)
	if err != nil {
		return nil, err
	}

	var key []byte
	if source, ok := msg.(*v1beta.Source); ok {
		key = generateKey(
			any.GetTypeUrl(),
			source.GetUrl(),
		)
	} else if module, ok := msg.(*v1beta.Module); ok {
		key = generateKey(
			any.GetTypeUrl(),
			module.GetLanguage(),
			module.GetName(),
		)
	}

	return &graphstore.Node{
		Key:  key,
		Body: any,
	}, nil
}

func newEdge(msg proto.Message) (*graphstore.Edge, error) {
	any, err := ptypes.MarshalAny(msg)
	if err != nil {
		return nil, err
	}

	return &graphstore.Edge{
		Body: any,
	}, nil
}
