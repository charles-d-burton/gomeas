package gomeas

import jsoniter "github.com/json-iterator/tinygo"

type Availability_json struct {
}

func (json Availability_json) Type() interface{} {
	var val Availability
	return val
}
func (json Availability_json) Unmarshal(iter *jsoniter.Iterator, out interface{}) {
	Availability_json_unmarshal(iter, out.(*Availability))
}
func (json Availability_json) Marshal(stream *jsoniter.Stream, val interface{}) {
	Availability_json_marshal(stream, val.(Availability))
}
func Availability_json_unmarshal(iter *jsoniter.Iterator, out *Availability) {
	more := iter.ReadObjectHead()
	for more {
		field := iter.ReadObjectField()
		if !Availability_json_unmarshal_field(iter, field, out) {
			iter.Skip()
		}
		more = iter.ReadObjectMore()
	}
}
func Availability_json_unmarshal_field(iter *jsoniter.Iterator, field string, out *Availability) bool {
	switch {
	case field == `payload_available`:
		iter.ReadString(&(*out).PayloadAvailable)
		return true
	case field == `payload_not_available`:
		iter.ReadString(&(*out).PayloadNotAvailable)
		return true
	case field == `topic`:
		iter.ReadString(&(*out).Topic)
		return true
	case field == `value_template`:
		iter.ReadString(&(*out).ValueTemplate)
		return true
	}
	return false
}
func Availability_json_marshal(stream *jsoniter.Stream, val Availability) {
	stream.WriteObjectHead()
	Availability_json_marshal_field(stream, val)
	stream.WriteObjectTail()
}
func Availability_json_marshal_field(stream *jsoniter.Stream, val Availability) {
	stream.WriteObjectField(`payload_available`)
	stream.WriteString(val.PayloadAvailable)
	stream.WriteMore()
	stream.WriteObjectField(`payload_not_available`)
	stream.WriteString(val.PayloadNotAvailable)
	stream.WriteMore()
	stream.WriteObjectField(`topic`)
	stream.WriteString(val.Topic)
	stream.WriteMore()
	stream.WriteObjectField(`value_template`)
	stream.WriteString(val.ValueTemplate)
	stream.WriteMore()
}
