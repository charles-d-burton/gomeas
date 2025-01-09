package gomeas

import jsoniter "github.com/json-iterator/tinygo"

type Availabilities_json struct {
}
func (json Availabilities_json) Type() interface{} {
  var val Availabilities
  return val
}
func (json Availabilities_json) Unmarshal(iter *jsoniter.Iterator, out interface{}) {
  Availabilities_json_unmarshal(iter, out.(*Availabilities))
}
func (json Availabilities_json) Marshal(stream *jsoniter.Stream, val interface{}) {
  Availabilities_json_marshal(stream, val.(Availabilities))
}
func Availabilities_json_unmarshal(iter *jsoniter.Iterator, out *Availabilities) {
  i := 0
  val := *out
  more := iter.ReadArrayHead()
  for more {
    if i == len(val) {
      val = append(val, make(Availabilities, 4)...)
    }
    Availabilities_ptr1_json_unmarshal(iter, &val[i])
    i++
    more = iter.ReadArrayMore()
  }
  if i == 0 {
    *out = Availabilities{}
  } else {
    *out = val[:i]
  }
}
func Availabilities_json_unmarshal_field(iter *jsoniter.Iterator, field string, out *Availabilities) bool {
  if field == "Availabilities" {
    Availabilities_json_unmarshal(iter, out)
    return true
  }
  return false
}
func Availabilities_ptr1_json_unmarshal (iter *jsoniter.Iterator, out **Availability) {
    var val Availability
    Availability_json_unmarshal(iter, &val)
    if iter.Error == nil {
      *out = &val
    }
}
func Availabilities_json_marshal(stream *jsoniter.Stream, val Availabilities) {
  if len(val) == 0 {
    stream.WriteEmptyArray()
  } else {
    stream.WriteArrayHead()
    for i, elem := range val {
      if i != 0 { stream.WriteMore() }
    if elem == nil {
       stream.WriteNull()
    } else {
    Availability_json_marshal(stream, *elem)
    }
    }
    stream.WriteArrayTail()
  }
}
func Availabilities_json_marshal_field(stream *jsoniter.Stream, val Availabilities) {
    stream.WriteObjectField("Availabilities")
    Availabilities_json_marshal(stream, val)
    stream.WriteMore()
}
