package gomeas

import jsoniter "github.com/json-iterator/tinygo"

type Config_json struct {
}
func (json Config_json) Type() interface{} {
  var val Config
  return val
}
func (json Config_json) Unmarshal(iter *jsoniter.Iterator, out interface{}) {
  Config_json_unmarshal(iter, out.(*Config))
}
func (json Config_json) Marshal(stream *jsoniter.Stream, val interface{}) {
  Config_json_marshal(stream, val.(Config))
}
func Config_json_unmarshal(iter *jsoniter.Iterator, out *Config) {
  more := iter.ReadObjectHead()
  for more {
    field := iter.ReadObjectField()
    if !Config_json_unmarshal_field(iter, field, out) {
      iter.Skip()
    }
    more = iter.ReadObjectMore()
  }
}
func Config_json_unmarshal_field(iter *jsoniter.Iterator, field string, out *Config) bool {
  switch {
  case field == `name`:
    iter.ReadString(&(*out).Name)
    return true
  case field == `device_class`:
    iter.ReadString(&(*out).DeviceClass)
    return true
  case field == `stat_t`:
    Config_ptr1_json_unmarshal(iter, &(*out).StateTopic)
    return true
  case field == `cmd_t`:
    Config_ptr2_json_unmarshal(iter, &(*out).ComamandTopic)
    return true
  case field == `unique_id`:
    iter.ReadString(&(*out).UniqueID)
    return true
  case field == `dev`:
    Config_ptr3_json_unmarshal(iter, &(*out).Device)
    return true
  case field == `cmps`:
    Config_ptr4_json_unmarshal(iter, &(*out).Components)
    return true
  }
  return false
}
func Config_ptr1_json_unmarshal (iter *jsoniter.Iterator, out **string) {
    var val string
    iter.ReadString(&val)
    if iter.Error == nil {
      *out = &val
    }
}
func Config_ptr2_json_unmarshal (iter *jsoniter.Iterator, out **string) {
    var val string
    iter.ReadString(&val)
    if iter.Error == nil {
      *out = &val
    }
}
func Config_ptr3_json_unmarshal (iter *jsoniter.Iterator, out **Device) {
    var val Device
    Device_json_unmarshal(iter, &val)
    if iter.Error == nil {
      *out = &val
    }
}
func Config_ptr4_json_unmarshal (iter *jsoniter.Iterator, out **ComponentMap) {
    var val ComponentMap
    ComponentMap_json_unmarshal(iter, &val)
    if iter.Error == nil {
      *out = &val
    }
}
func Config_json_marshal(stream *jsoniter.Stream, val Config) {
    stream.WriteObjectHead()
    Config_json_marshal_field(stream, val)
    stream.WriteObjectTail()
}
func Config_json_marshal_field(stream *jsoniter.Stream, val Config) {
    stream.WriteObjectField(`name`)
    stream.WriteString(val.Name)
    stream.WriteMore()
    stream.WriteObjectField(`device_class`)
    stream.WriteString(val.DeviceClass)
    stream.WriteMore()
    stream.WriteObjectField(`stat_t`)
    if val.StateTopic == nil {
       stream.WriteNull()
    } else {
    stream.WriteString(*val.StateTopic)
    }
    stream.WriteMore()
    stream.WriteObjectField(`cmd_t`)
    if val.ComamandTopic == nil {
       stream.WriteNull()
    } else {
    stream.WriteString(*val.ComamandTopic)
    }
    stream.WriteMore()
    stream.WriteObjectField(`unique_id`)
    stream.WriteString(val.UniqueID)
    stream.WriteMore()
    stream.WriteObjectField(`dev`)
    if val.Device == nil {
       stream.WriteNull()
    } else {
    Device_json_marshal(stream, *val.Device)
    }
    stream.WriteMore()
    stream.WriteObjectField(`cmps`)
    if val.Components == nil {
       stream.WriteNull()
    } else {
    ComponentMap_json_marshal(stream, *val.Components)
    }
    stream.WriteMore()
}
