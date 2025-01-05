package gomeas

import jsoniter "github.com/json-iterator/tinygo"

type Component_json struct {
}
func (json Component_json) Type() interface{} {
  var val Component
  return val
}
func (json Component_json) Unmarshal(iter *jsoniter.Iterator, out interface{}) {
  Component_json_unmarshal(iter, out.(*Component))
}
func (json Component_json) Marshal(stream *jsoniter.Stream, val interface{}) {
  Component_json_marshal(stream, val.(Component))
}
func Component_json_unmarshal(iter *jsoniter.Iterator, out *Component) {
  more := iter.ReadObjectHead()
  for more {
    field := iter.ReadObjectField()
    if !Component_json_unmarshal_field(iter, field, out) {
      iter.Skip()
    }
    more = iter.ReadObjectMore()
  }
}
func Component_json_unmarshal_field(iter *jsoniter.Iterator, field string, out *Component) bool {
  switch {
  case field == `p`:
    iter.ReadString(&(*out).P)
    return true
  case field == `device_class`:
    iter.ReadString(&(*out).DeviceClass)
    return true
  case field == `unit_of_measurement`:
    iter.ReadString(&(*out).UnitOfMeasurement)
    return true
  case field == `value_template`:
    iter.ReadString(&(*out).ValueTemplate)
    return true
  case field == `unique_id`:
    iter.ReadString(&(*out).UniqueID)
    return true
  }
  return false
}
func Component_json_marshal(stream *jsoniter.Stream, val Component) {
    stream.WriteObjectHead()
    Component_json_marshal_field(stream, val)
    stream.WriteObjectTail()
}
func Component_json_marshal_field(stream *jsoniter.Stream, val Component) {
    stream.WriteObjectField(`p`)
    stream.WriteString(val.P)
    stream.WriteMore()
    stream.WriteObjectField(`device_class`)
    stream.WriteString(val.DeviceClass)
    stream.WriteMore()
    stream.WriteObjectField(`unit_of_measurement`)
    stream.WriteString(val.UnitOfMeasurement)
    stream.WriteMore()
    stream.WriteObjectField(`value_template`)
    stream.WriteString(val.ValueTemplate)
    stream.WriteMore()
    stream.WriteObjectField(`unique_id`)
    stream.WriteString(val.UniqueID)
    stream.WriteMore()
}
