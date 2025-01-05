package gomeas

import jsoniter "github.com/json-iterator/tinygo"

type BinarySensor_json struct {
}
func (json BinarySensor_json) Type() interface{} {
  var val BinarySensor
  return val
}
func (json BinarySensor_json) Unmarshal(iter *jsoniter.Iterator, out interface{}) {
  BinarySensor_json_unmarshal(iter, out.(*BinarySensor))
}
func (json BinarySensor_json) Marshal(stream *jsoniter.Stream, val interface{}) {
  BinarySensor_json_marshal(stream, val.(BinarySensor))
}
func BinarySensor_json_unmarshal(iter *jsoniter.Iterator, out *BinarySensor) {
  more := iter.ReadObjectHead()
  for more {
    field := iter.ReadObjectField()
    if !BinarySensor_json_unmarshal_field(iter, field, out) {
      iter.Skip()
    }
    more = iter.ReadObjectMore()
  }
}
func BinarySensor_json_unmarshal_field(iter *jsoniter.Iterator, field string, out *BinarySensor) bool {
  switch {
  case field == `device`:
    Device_json_unmarshal(iter, &(*out).Device)
    return true
  case field == `availability`:
    BinarySensor_array1_json_unmarshal(iter, &(*out).Availability)
    return true
  case field == `availibility_mode`:
    iter.ReadString(&(*out).AvailibilityMode)
    return true
  case field == `availability_template`:
    iter.ReadString(&(*out).AvailabilityTemplate)
    return true
  case field == `availability_topic`:
    iter.ReadString(&(*out).AvailabilityTopic)
    return true
  case field == `device_class`:
    BinarySensor_ptr3_json_unmarshal(iter, &(*out).DeviceClass)
    return true
  case field == `enabled_by_default`:
    iter.ReadBool(&(*out).EnabledByDefault)
    return true
  case field == `encoding`:
    iter.ReadString(&(*out).Encoding)
    return true
  case field == `entity_category`:
    iter.ReadString(&(*out).EntityCategory)
    return true
  case field == `entity_picture`:
    iter.ReadString(&(*out).EntityPicture)
    return true
  case field == `expire_after`:
    iter.ReadInt(&(*out).ExpireAfter)
    return true
  case field == `force_update`:
    iter.ReadBool(&(*out).ForceUpdate)
    return true
  case field == `icon`:
    iter.ReadString(&(*out).Icon)
    return true
  case field == `json_attributes_template`:
    iter.ReadString(&(*out).JsonAttributesTemplate)
    return true
  case field == `json_attributes_topic`:
    iter.ReadString(&(*out).JsonAttributesTopic)
    return true
  case field == `name`:
    iter.ReadString(&(*out).Name)
    return true
  case field == `object_id`:
    iter.ReadString(&(*out).ObjectID)
    return true
  case field == `off_delay`:
    iter.ReadInt(&(*out).OffDelay)
    return true
  case field == `payload_available`:
    iter.ReadString(&(*out).PayloadAvailable)
    return true
  case field == `payload_not_available`:
    iter.ReadString(&(*out).PayloadNotAvailable)
    return true
  case field == `payload_off`:
    iter.ReadString(&(*out).PayloadOff)
    return true
  case field == `payload_on`:
    iter.ReadString(&(*out).PayloadOn)
    return true
  case field == `platform`:
    iter.ReadString(&(*out).Platform)
    return true
  case field == `qos`:
    iter.ReadInt(&(*out).Qos)
    return true
  case field == `state_topic`:
    iter.ReadString(&(*out).StateTopic)
    return true
  case field == `unique_id`:
    iter.ReadString(&(*out).UniqueID)
    return true
  case field == `value_template`:
    iter.ReadString(&(*out).ValueTemplate)
    return true
  }
  return false
}
func BinarySensor_struct2_json_unmarshal_field (iter *jsoniter.Iterator, field string, out *struct {
	PayloadAvailable	string	`json:"payload_available,omitempty"`
	PayloadNotAvailable	string	`json:"payload_not_available,omitempty"`
	Topic			string	`json:"topic"`
	ValueTemplate		string	`json:"value_template,omitempty"`
}) bool {
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
func BinarySensor_struct2_json_unmarshal (iter *jsoniter.Iterator, out *struct {
	PayloadAvailable	string	`json:"payload_available,omitempty"`
	PayloadNotAvailable	string	`json:"payload_not_available,omitempty"`
	Topic			string	`json:"topic"`
	ValueTemplate		string	`json:"value_template,omitempty"`
}) {
  more := iter.ReadObjectHead()
  for more {
    field := iter.ReadObjectField()
    if !BinarySensor_struct2_json_unmarshal_field(iter, field, out) {
      iter.Skip()
    }
    more = iter.ReadObjectMore()
  }
}
func BinarySensor_array1_json_unmarshal (iter *jsoniter.Iterator, out *[]struct {
	PayloadAvailable	string	`json:"payload_available,omitempty"`
	PayloadNotAvailable	string	`json:"payload_not_available,omitempty"`
	Topic			string	`json:"topic"`
	ValueTemplate		string	`json:"value_template,omitempty"`
}) {
  i := 0
  val := *out
  more := iter.ReadArrayHead()
  for more {
    if i == len(val) {
      val = append(val, make([]struct {
	PayloadAvailable	string	`json:"payload_available,omitempty"`
	PayloadNotAvailable	string	`json:"payload_not_available,omitempty"`
	Topic			string	`json:"topic"`
	ValueTemplate		string	`json:"value_template,omitempty"`
}, 4)...)
    }
    BinarySensor_struct2_json_unmarshal(iter, &val[i])
    i++
    more = iter.ReadArrayMore()
  }
  if i == 0 {
    *out = []struct {
	PayloadAvailable	string	`json:"payload_available,omitempty"`
	PayloadNotAvailable	string	`json:"payload_not_available,omitempty"`
	Topic			string	`json:"topic"`
	ValueTemplate		string	`json:"value_template,omitempty"`
}{}
  } else {
    *out = val[:i]
  }
}
func BinarySensor_ptr3_json_unmarshal (iter *jsoniter.Iterator, out **string) {
    var val string
    iter.ReadString(&val)
    if iter.Error == nil {
      *out = &val
    }
}
func BinarySensor_json_marshal(stream *jsoniter.Stream, val BinarySensor) {
    stream.WriteObjectHead()
    BinarySensor_json_marshal_field(stream, val)
    stream.WriteObjectTail()
}
func BinarySensor_json_marshal_field(stream *jsoniter.Stream, val BinarySensor) {
    stream.WriteObjectField(`device`)
    Device_json_marshal(stream, val.Device)
    stream.WriteMore()
    stream.WriteObjectField(`availability`)
    BinarySensor_array4_json_marshal(stream, val.Availability)
    stream.WriteMore()
    stream.WriteObjectField(`availibility_mode`)
    stream.WriteString(val.AvailibilityMode)
    stream.WriteMore()
    stream.WriteObjectField(`availability_template`)
    stream.WriteString(val.AvailabilityTemplate)
    stream.WriteMore()
    stream.WriteObjectField(`availability_topic`)
    stream.WriteString(val.AvailabilityTopic)
    stream.WriteMore()
    stream.WriteObjectField(`device_class`)
    if val.DeviceClass == nil {
       stream.WriteNull()
    } else {
    stream.WriteString(*val.DeviceClass)
    }
    stream.WriteMore()
    stream.WriteObjectField(`enabled_by_default`)
    stream.WriteBool(val.EnabledByDefault)
    stream.WriteMore()
    stream.WriteObjectField(`encoding`)
    stream.WriteString(val.Encoding)
    stream.WriteMore()
    stream.WriteObjectField(`entity_category`)
    stream.WriteString(val.EntityCategory)
    stream.WriteMore()
    stream.WriteObjectField(`entity_picture`)
    stream.WriteString(val.EntityPicture)
    stream.WriteMore()
    stream.WriteObjectField(`expire_after`)
    stream.WriteInt(val.ExpireAfter)
    stream.WriteMore()
    stream.WriteObjectField(`force_update`)
    stream.WriteBool(val.ForceUpdate)
    stream.WriteMore()
    stream.WriteObjectField(`icon`)
    stream.WriteString(val.Icon)
    stream.WriteMore()
    stream.WriteObjectField(`json_attributes_template`)
    stream.WriteString(val.JsonAttributesTemplate)
    stream.WriteMore()
    stream.WriteObjectField(`json_attributes_topic`)
    stream.WriteString(val.JsonAttributesTopic)
    stream.WriteMore()
    stream.WriteObjectField(`name`)
    stream.WriteString(val.Name)
    stream.WriteMore()
    stream.WriteObjectField(`object_id`)
    stream.WriteString(val.ObjectID)
    stream.WriteMore()
    stream.WriteObjectField(`off_delay`)
    stream.WriteInt(val.OffDelay)
    stream.WriteMore()
    stream.WriteObjectField(`payload_available`)
    stream.WriteString(val.PayloadAvailable)
    stream.WriteMore()
    stream.WriteObjectField(`payload_not_available`)
    stream.WriteString(val.PayloadNotAvailable)
    stream.WriteMore()
    stream.WriteObjectField(`payload_off`)
    stream.WriteString(val.PayloadOff)
    stream.WriteMore()
    stream.WriteObjectField(`payload_on`)
    stream.WriteString(val.PayloadOn)
    stream.WriteMore()
    stream.WriteObjectField(`platform`)
    stream.WriteString(val.Platform)
    stream.WriteMore()
    stream.WriteObjectField(`qos`)
    stream.WriteInt(val.Qos)
    stream.WriteMore()
    stream.WriteObjectField(`state_topic`)
    stream.WriteString(val.StateTopic)
    stream.WriteMore()
    stream.WriteObjectField(`unique_id`)
    stream.WriteString(val.UniqueID)
    stream.WriteMore()
    stream.WriteObjectField(`value_template`)
    stream.WriteString(val.ValueTemplate)
    stream.WriteMore()
}
func BinarySensor_struct5_json_marshal_field (stream *jsoniter.Stream, val struct {
	PayloadAvailable	string	`json:"payload_available,omitempty"`
	PayloadNotAvailable	string	`json:"payload_not_available,omitempty"`
	Topic			string	`json:"topic"`
	ValueTemplate		string	`json:"value_template,omitempty"`
}) {
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
func BinarySensor_struct5_json_marshal (stream *jsoniter.Stream, val struct {
	PayloadAvailable	string	`json:"payload_available,omitempty"`
	PayloadNotAvailable	string	`json:"payload_not_available,omitempty"`
	Topic			string	`json:"topic"`
	ValueTemplate		string	`json:"value_template,omitempty"`
}) {
    stream.WriteObjectHead()
    BinarySensor_struct5_json_marshal_field(stream, val)
    stream.WriteObjectTail()
}
func BinarySensor_array4_json_marshal (stream *jsoniter.Stream, val []struct {
	PayloadAvailable	string	`json:"payload_available,omitempty"`
	PayloadNotAvailable	string	`json:"payload_not_available,omitempty"`
	Topic			string	`json:"topic"`
	ValueTemplate		string	`json:"value_template,omitempty"`
}) {
  if len(val) == 0 {
    stream.WriteEmptyArray()
  } else {
    stream.WriteArrayHead()
    for i, elem := range val {
      if i != 0 { stream.WriteMore() }
    BinarySensor_struct5_json_marshal(stream, elem)
    }
    stream.WriteArrayTail()
  }
}
