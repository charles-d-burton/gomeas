package gomeas

import jsoniter "github.com/json-iterator/tinygo"

type ComponentMap_json struct {
}
func (json ComponentMap_json) Type() interface{} {
  var val ComponentMap
  return val
}
func (json ComponentMap_json) Unmarshal(iter *jsoniter.Iterator, out interface{}) {
  ComponentMap_json_unmarshal(iter, out.(*ComponentMap))
}
func (json ComponentMap_json) Marshal(stream *jsoniter.Stream, val interface{}) {
  ComponentMap_json_marshal(stream, val.(ComponentMap))
}
func ComponentMap_json_unmarshal(iter *jsoniter.Iterator, out *ComponentMap) {
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
func ComponentMap_json_unmarshal_field(iter *jsoniter.Iterator, field string, out *ComponentMap) bool {
  if field == "ComponentMap" {
    ComponentMap_json_unmarshal(iter, out)
    return true
  }
  return false
}
func ComponentMap_json_marshal(stream *jsoniter.Stream, val ComponentMap) {
  stream.WriteObjectHead()
  for k, v := range val {
    stream.WriteObjectField(k)
    Component_json_marshal(stream, v)
    stream.WriteMore()
  }
  stream.WriteObjectTail()
}
func ComponentMap_json_marshal_field(stream *jsoniter.Stream, val ComponentMap) {
    stream.WriteObjectField("ComponentMap")
    ComponentMap_json_marshal(stream, val)
    stream.WriteMore()
}
