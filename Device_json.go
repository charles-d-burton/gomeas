package gomeas

import jsoniter "github.com/json-iterator/tinygo"

type Device_json struct {
}
func (json Device_json) Type() interface{} {
  var val Device
  return val
}
func (json Device_json) Unmarshal(iter *jsoniter.Iterator, out interface{}) {
  Device_json_unmarshal(iter, out.(*Device))
}
func (json Device_json) Marshal(stream *jsoniter.Stream, val interface{}) {
  Device_json_marshal(stream, val.(Device))
}
func Device_json_unmarshal(iter *jsoniter.Iterator, out *Device) {
  more := iter.ReadObjectHead()
  for more {
    field := iter.ReadObjectField()
    if !Device_json_unmarshal_field(iter, field, out) {
      iter.Skip()
    }
    more = iter.ReadObjectMore()
  }
}
func Device_json_unmarshal_field(iter *jsoniter.Iterator, field string, out *Device) bool {
  switch {
  case field == `cu`:
    iter.ReadString(&(*out).ConfigurationURL)
    return true
  case field == `cns`:
    Device_array1_json_unmarshal(iter, &(*out).Connections)
    return true
  case field == `hw`:
    iter.ReadString(&(*out).HardwareVersion)
    return true
  case field == `ids`:
    Device_array2_json_unmarshal(iter, &(*out).Identifiers)
    return true
  case field == `mf`:
    iter.ReadString(&(*out).Manufacturer)
    return true
  case field == `mdl`:
    iter.ReadString(&(*out).Model)
    return true
  case field == `mdl_id`:
    iter.ReadString(&(*out).ModelID)
    return true
  case field == `name`:
    iter.ReadString(&(*out).Name)
    return true
  case field == `sn`:
    iter.ReadString(&(*out).SerialNumber)
    return true
  case field == `sa`:
    iter.ReadString(&(*out).SuggestedArea)
    return true
  case field == `sw`:
    iter.ReadString(&(*out).SoftwareVersion)
    return true
  case field == `via_device`:
    iter.ReadString(&(*out).ViaDevice)
    return true
  }
  return false
}
func Device_array1_json_unmarshal (iter *jsoniter.Iterator, out *[]string) {
  i := 0
  val := *out
  more := iter.ReadArrayHead()
  for more {
    if i == len(val) {
      val = append(val, make([]string, 4)...)
    }
    iter.ReadString(&val[i])
    i++
    more = iter.ReadArrayMore()
  }
  if i == 0 {
    *out = []string{}
  } else {
    *out = val[:i]
  }
}
func Device_array2_json_unmarshal (iter *jsoniter.Iterator, out *[]string) {
  i := 0
  val := *out
  more := iter.ReadArrayHead()
  for more {
    if i == len(val) {
      val = append(val, make([]string, 4)...)
    }
    iter.ReadString(&val[i])
    i++
    more = iter.ReadArrayMore()
  }
  if i == 0 {
    *out = []string{}
  } else {
    *out = val[:i]
  }
}
func Device_json_marshal(stream *jsoniter.Stream, val Device) {
    stream.WriteObjectHead()
    Device_json_marshal_field(stream, val)
    stream.WriteObjectTail()
}
func Device_json_marshal_field(stream *jsoniter.Stream, val Device) {
    stream.WriteObjectField(`cu`)
    stream.WriteString(val.ConfigurationURL)
    stream.WriteMore()
    stream.WriteObjectField(`cns`)
    Device_array3_json_marshal(stream, val.Connections)
    stream.WriteMore()
    stream.WriteObjectField(`hw`)
    stream.WriteString(val.HardwareVersion)
    stream.WriteMore()
    stream.WriteObjectField(`ids`)
    Device_array4_json_marshal(stream, val.Identifiers)
    stream.WriteMore()
    stream.WriteObjectField(`mf`)
    stream.WriteString(val.Manufacturer)
    stream.WriteMore()
    stream.WriteObjectField(`mdl`)
    stream.WriteString(val.Model)
    stream.WriteMore()
    stream.WriteObjectField(`mdl_id`)
    stream.WriteString(val.ModelID)
    stream.WriteMore()
    stream.WriteObjectField(`name`)
    stream.WriteString(val.Name)
    stream.WriteMore()
    stream.WriteObjectField(`sn`)
    stream.WriteString(val.SerialNumber)
    stream.WriteMore()
    stream.WriteObjectField(`sa`)
    stream.WriteString(val.SuggestedArea)
    stream.WriteMore()
    stream.WriteObjectField(`sw`)
    stream.WriteString(val.SoftwareVersion)
    stream.WriteMore()
    stream.WriteObjectField(`via_device`)
    stream.WriteString(val.ViaDevice)
    stream.WriteMore()
}
func Device_array3_json_marshal (stream *jsoniter.Stream, val []string) {
  if len(val) == 0 {
    stream.WriteEmptyArray()
  } else {
    stream.WriteArrayHead()
    for i, elem := range val {
      if i != 0 { stream.WriteMore() }
    stream.WriteString(elem)
    }
    stream.WriteArrayTail()
  }
}
func Device_array4_json_marshal (stream *jsoniter.Stream, val []string) {
  if len(val) == 0 {
    stream.WriteEmptyArray()
  } else {
    stream.WriteArrayHead()
    for i, elem := range val {
      if i != 0 { stream.WriteMore() }
    stream.WriteString(elem)
    }
    stream.WriteArrayTail()
  }
}
