package gomeas

import jsoniter "github.com/json-iterator/tinygo"

type Components_json struct {
}

func (json Components_json) Type() interface{} {
	var val Components
	return val
}
func (json Components_json) Unmarshal(iter *jsoniter.Iterator, out interface{}) {
	Components_json_unmarshal(iter, out.(*Components))
}
func (json Components_json) Marshal(stream *jsoniter.Stream, val interface{}) {
	Components_json_marshal(stream, val.(Components))
}
func Components_json_unmarshal(iter *jsoniter.Iterator, out *Components) {
	more := iter.ReadObjectHead()
	if *out == nil && iter.Error == nil {
		*out = make(map[string]Component)
	}
	for more {
		field := iter.ReadObjectField()
		var value Component
		var key string
		var err error
		key = field
		Component_json_unmarshal(iter, &value)
		if err != nil {
			iter.ReportError("read map key", err.Error())
		} else {
			(*out)[key] = value
		}
		more = iter.ReadObjectMore()
	}
}
func Components_json_unmarshal_field(iter *jsoniter.Iterator, field string, out *Components) bool {
	if field == "Components" {
		Components_json_unmarshal(iter, out)
		return true
	}
	return false
}
func Components_json_marshal(stream *jsoniter.Stream, val Components) {
	stream.WriteObjectHead()
	for k, v := range val {
		stream.WriteObjectField(k)
		Component_json_marshal(stream, v)
		stream.WriteMore()
	}
	stream.WriteObjectTail()
}
func Components_json_marshal_field(stream *jsoniter.Stream, val Components) {
	stream.WriteObjectField("Components")
	Components_json_marshal(stream, val)
	stream.WriteMore()
}
