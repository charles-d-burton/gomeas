package gomeas

import jsoniter "github.com/json-iterator/tinygo"

type BinarySensorState_json struct {
}

func (json BinarySensorState_json) Type() interface{} {
	var val BinarySensorState
	return val
}
func (json BinarySensorState_json) Unmarshal(iter *jsoniter.Iterator, out interface{}) {
	BinarySensorState_json_unmarshal(iter, out.(*BinarySensorState))
}
func (json BinarySensorState_json) Marshal(stream *jsoniter.Stream, val interface{}) {
	BinarySensorState_json_marshal(stream, val.(BinarySensorState))
}
func BinarySensorState_json_unmarshal(iter *jsoniter.Iterator, out *BinarySensorState) {
	more := iter.ReadObjectHead()
	for more {
		field := iter.ReadObjectField()
		if !BinarySensorState_json_unmarshal_field(iter, field, out) {
			iter.Skip()
		}
		more = iter.ReadObjectMore()
	}
}
func BinarySensorState_json_unmarshal_field(iter *jsoniter.Iterator, field string, out *BinarySensorState) bool {
	switch {
	case field == `state`:
		iter.ReadString(&(*out).State)
		return true
	}
	return false
}
func BinarySensorState_json_marshal(stream *jsoniter.Stream, val BinarySensorState) {
	stream.WriteObjectHead()
	BinarySensorState_json_marshal_field(stream, val)
	stream.WriteObjectTail()
}
func BinarySensorState_json_marshal_field(stream *jsoniter.Stream, val BinarySensorState) {
	stream.WriteObjectField(`state`)
	stream.WriteString(val.State)
	stream.WriteMore()
}
