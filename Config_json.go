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
	case field == `availability`:
		Config_array1_json_unmarshal(iter, &(*out).Availability)
		return true
	case field == `availibility_mode`:
		Config_ptr3_json_unmarshal(iter, &(*out).AvailibilityMode)
		return true
	case field == `availability_template`:
		Config_ptr4_json_unmarshal(iter, &(*out).AvailabilityTemplate)
		return true
	case field == `availability_topic`:
		Config_ptr5_json_unmarshal(iter, &(*out).AvailabilityTopic)
		return true
	case field == `name`:
		iter.ReadString(&(*out).Name)
		return true
	case field == `device_class`:
		iter.ReadString(&(*out).DeviceClass)
		return true
	case field == `state_topic`:
		Config_ptr6_json_unmarshal(iter, &(*out).StateTopic)
		return true
	case field == `command_topic`:
		Config_ptr7_json_unmarshal(iter, &(*out).ComamandTopic)
		return true
	case field == `unique_id`:
		iter.ReadString(&(*out).UniqueID)
		return true
	case field == `device`:
		Config_ptr8_json_unmarshal(iter, &(*out).Device)
		return true
	case field == `components`:
		Components_json_unmarshal(iter, &(*out).Components)
		return true
	case field == `enabled_by_default`:
		iter.ReadBool(&(*out).EnabledByDefault)
		return true
	}
	return false
}
func Config_ptr2_json_unmarshal(iter *jsoniter.Iterator, out **Availability) {
	var val Availability
	Availability_json_unmarshal(iter, &val)
	if iter.Error == nil {
		*out = &val
	}
}
func Config_array1_json_unmarshal(iter *jsoniter.Iterator, out *[]*Availability) {
	i := 0
	val := *out
	more := iter.ReadArrayHead()
	for more {
		if i == len(val) {
			val = append(val, make([]*Availability, 4)...)
		}
		Config_ptr2_json_unmarshal(iter, &val[i])
		i++
		more = iter.ReadArrayMore()
	}
	if i == 0 {
		*out = []*Availability{}
	} else {
		*out = val[:i]
	}
}
func Config_ptr3_json_unmarshal(iter *jsoniter.Iterator, out **string) {
	var val string
	iter.ReadString(&val)
	if iter.Error == nil {
		*out = &val
	}
}
func Config_ptr4_json_unmarshal(iter *jsoniter.Iterator, out **string) {
	var val string
	iter.ReadString(&val)
	if iter.Error == nil {
		*out = &val
	}
}
func Config_ptr5_json_unmarshal(iter *jsoniter.Iterator, out **string) {
	var val string
	iter.ReadString(&val)
	if iter.Error == nil {
		*out = &val
	}
}
func Config_ptr6_json_unmarshal(iter *jsoniter.Iterator, out **string) {
	var val string
	iter.ReadString(&val)
	if iter.Error == nil {
		*out = &val
	}
}
func Config_ptr7_json_unmarshal(iter *jsoniter.Iterator, out **string) {
	var val string
	iter.ReadString(&val)
	if iter.Error == nil {
		*out = &val
	}
}
func Config_ptr8_json_unmarshal(iter *jsoniter.Iterator, out **Device) {
	var val Device
	Device_json_unmarshal(iter, &val)
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
	stream.WriteObjectField(`availability`)
	Config_array9_json_marshal(stream, val.Availability)
	stream.WriteMore()
	stream.WriteObjectField(`availibility_mode`)
	if val.AvailibilityMode == nil {
		stream.WriteNull()
	} else {
		stream.WriteString(*val.AvailibilityMode)
	}
	stream.WriteMore()
	stream.WriteObjectField(`availability_template`)
	if val.AvailabilityTemplate == nil {
		stream.WriteNull()
	} else {
		stream.WriteString(*val.AvailabilityTemplate)
	}
	stream.WriteMore()
	stream.WriteObjectField(`availability_topic`)
	if val.AvailabilityTopic == nil {
		stream.WriteNull()
	} else {
		stream.WriteString(*val.AvailabilityTopic)
	}
	stream.WriteMore()
	stream.WriteObjectField(`name`)
	stream.WriteString(val.Name)
	stream.WriteMore()
	stream.WriteObjectField(`device_class`)
	stream.WriteString(val.DeviceClass)
	stream.WriteMore()
	stream.WriteObjectField(`state_topic`)
	if val.StateTopic == nil {
		stream.WriteNull()
	} else {
		stream.WriteString(*val.StateTopic)
	}
	stream.WriteMore()
	stream.WriteObjectField(`command_topic`)
	if val.ComamandTopic == nil {
		stream.WriteNull()
	} else {
		stream.WriteString(*val.ComamandTopic)
	}
	stream.WriteMore()
	stream.WriteObjectField(`unique_id`)
	stream.WriteString(val.UniqueID)
	stream.WriteMore()
	stream.WriteObjectField(`device`)
	if val.Device == nil {
		stream.WriteNull()
	} else {
		Device_json_marshal(stream, *val.Device)
	}
	stream.WriteMore()
	stream.WriteObjectField(`components`)
	Components_json_marshal(stream, val.Components)
	stream.WriteMore()
	stream.WriteObjectField(`enabled_by_default`)
	stream.WriteBool(val.EnabledByDefault)
	stream.WriteMore()
}
func Config_array9_json_marshal(stream *jsoniter.Stream, val []*Availability) {
	if len(val) == 0 {
		stream.WriteEmptyArray()
	} else {
		stream.WriteArrayHead()
		for i, elem := range val {
			if i != 0 {
				stream.WriteMore()
			}
			if elem == nil {
				stream.WriteNull()
			} else {
				Availability_json_marshal(stream, *elem)
			}
		}
		stream.WriteArrayTail()
	}
}
