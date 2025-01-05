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
    iter.ReadString(&(*out).StateTopic)
    return true
  case field == `cmd_t`:
    iter.ReadString(&(*out).ComamandTopic)
    return true
  case field == `unique_id`:
    iter.ReadString(&(*out).UniqueID)
    return true
  case field == `dev`:
    Device_json_unmarshal(iter, &(*out).Device)
    return true
  case field == `cmps`:
    Config_map1_json_unmarshal(iter, &(*out).Components)
    return true
  }
  return false
}
func Config_map1_json_unmarshal (iter *jsoniter.Iterator, out *map[string]Component) {
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
    stream.WriteString(val.StateTopic)
    stream.WriteMore()
    stream.WriteObjectField(`cmd_t`)
    stream.WriteString(val.ComamandTopic)
    stream.WriteMore()
    stream.WriteObjectField(`unique_id`)
    stream.WriteString(val.UniqueID)
    stream.WriteMore()
    stream.WriteObjectField(`dev`)
    Device_json_marshal(stream, val.Device)
    stream.WriteMore()
    stream.WriteObjectField(`cmps`)
    Config_map2_json_marshal(stream, val.Components)
    stream.WriteMore()
}
func Config_map2_json_marshal (stream *jsoniter.Stream, val map[string]Component) {
  stream.WriteObjectHead()
  for k, v := range val {
    stream.WriteObjectField(k)
    Component_json_marshal(stream, v)
    stream.WriteMore()
  }
  stream.WriteObjectTail()
}
