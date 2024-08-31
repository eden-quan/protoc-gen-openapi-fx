package generator

import (
	"google.golang.org/protobuf/compiler/protogen"
	"google.golang.org/protobuf/proto"

	"github.com/eden-quan/protoc-gen-openapi-fx/generator/meta"
)

type FlattenInfo struct {
	Rule    *meta.FlattenRules
	Message *protogen.Message
	Field   *protogen.Field
}

func ExtractFlatten(message *protogen.Message, field *protogen.Field) (*FlattenInfo, bool) {
	flatten := proto.GetExtension(field.Desc.Options(), meta.E_Flatten).(bool)
	ext := proto.GetExtension(field.Desc.Options(), meta.E_FlattenRule).(*meta.FlattenRules)

	if ext == nil {
		// create default ext
		m := int32(0)
		ext = &meta.FlattenRules{Reserved: &meta.Reserved{
			Min: &m, // useless now
			Max: &m, // useless now
		}}
	} else {
		flatten = true
	}

	return &FlattenInfo{
		Rule:    ext,
		Message: message,
		Field:   field,
	}, flatten
}
