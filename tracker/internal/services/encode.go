package services

import (
	"encoding/json"
	"fmt"

	"github.com/depscloud/api/v1alpha/schema"
	"github.com/depscloud/api/v1alpha/store"
	"github.com/depscloud/depscloud/tracker/internal/types"
)

// Encode turns the provided schma type into the corresponding GraphItem
func Encode(msg interface{}) (*store.GraphItem, error) {
	var graphItemType string
	var k1 []byte
	var k2 []byte
	var k3 []byte

	switch msg.(type) {
	case *schema.Source:
		graphItemType = types.SourceType
		key := keyForSource(msg.(*schema.Source))
		k1 = key
		k2 = key
	case *schema.Manages:
		graphItemType = types.ManagesType
	case *schema.Module:
		graphItemType = types.ModuleType
		key := keyForModule(msg.(*schema.Module))
		k1 = key
		k2 = key
	case *schema.Depends:
		graphItemType = types.DependsType
	default:
		return nil, fmt.Errorf("unrecognized type")
	}

	graphItemData, err := json.Marshal(msg)
	if err != nil {
		return nil, fmt.Errorf("marshal failure")
	}

	return &store.GraphItem{
		GraphItemType: graphItemType,
		K1:            k1,
		K2:            k2,
		K3:            k3,
		Encoding:      store.GraphItemEncoding_JSON,
		GraphItemData: graphItemData,
	}, nil
}
