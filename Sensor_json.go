package gomeas

import jsoniter "github.com/json-iterator/tinygo"

type Sensor_json struct {
}

func (json Sensor_json) Type() interface{} {
	var val Sensor
	return val
}
func (json Sensor_json) Unmarshal(iter *jsoniter.Iterator, out interface{}) {
	Sensor_json_unmarshal(iter, out.(*Sensor))
}
func (json Sensor_json) Marshal(stream *jsoniter.Stream, val interface{}) {
	Sensor_json_marshal(stream, val.(Sensor))
}
func Sensor_json_unmarshal(iter *jsoniter.Iterator, out *Sensor) {
	more := iter.ReadObjectHead()
	for more {
		field := iter.ReadObjectField()
		if !Sensor_json_unmarshal_field(iter, field, out) {
			iter.Skip()
		}
		more = iter.ReadObjectMore()
	}
}
func Sensor_json_unmarshal_field(iter *jsoniter.Iterator, field string, out *Sensor) bool {
	if Config_json_unmarshal_field(iter, field, &out.Config) {
		return true
	}
	switch {
	case field == `platform`:
		iter.ReadString(&(*out).Platform)
		return true
	}
	return false
}
func Sensor_json_marshal(stream *jsoniter.Stream, val Sensor) {
	stream.WriteObjectHead()
	Sensor_json_marshal_field(stream, val)
	stream.WriteObjectTail()
}
func Sensor_json_marshal_field(stream *jsoniter.Stream, val Sensor) {
	Config_json_marshal_field(stream, val.Config)
	stream.WriteObjectField(`platform`)
	stream.WriteString(val.Platform)
	stream.WriteMore()
}
