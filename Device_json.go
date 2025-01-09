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
	case field == `configuration_url`:
		Device_ptr1_json_unmarshal(iter, &(*out).ConfigurationURL)
		return true
	case field == `connections`:
		Device_array2_json_unmarshal(iter, &(*out).Connections)
		return true
	case field == `hw_version`:
		Device_ptr3_json_unmarshal(iter, &(*out).HardwareVersion)
		return true
	case field == `identifiers`:
		Device_array4_json_unmarshal(iter, &(*out).Identifiers)
		return true
	case field == `manufacturer`:
		Device_ptr5_json_unmarshal(iter, &(*out).Manufacturer)
		return true
	case field == `model`:
		Device_ptr6_json_unmarshal(iter, &(*out).Model)
		return true
	case field == `model_id`:
		Device_ptr7_json_unmarshal(iter, &(*out).ModelID)
		return true
	case field == `name`:
		Device_ptr8_json_unmarshal(iter, &(*out).Name)
		return true
	case field == `serial_number`:
		Device_ptr9_json_unmarshal(iter, &(*out).SerialNumber)
		return true
	case field == `suggested_area`:
		Device_ptr10_json_unmarshal(iter, &(*out).SuggestedArea)
		return true
	case field == `software_version`:
		Device_ptr11_json_unmarshal(iter, &(*out).SoftwareVersion)
		return true
	case field == `via_device`:
		Device_ptr12_json_unmarshal(iter, &(*out).ViaDevice)
		return true
	}
	return false
}
func Device_ptr1_json_unmarshal(iter *jsoniter.Iterator, out **string) {
	var val string
	iter.ReadString(&val)
	if iter.Error == nil {
		*out = &val
	}
}
func Device_array2_json_unmarshal(iter *jsoniter.Iterator, out *[]string) {
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
func Device_ptr3_json_unmarshal(iter *jsoniter.Iterator, out **string) {
	var val string
	iter.ReadString(&val)
	if iter.Error == nil {
		*out = &val
	}
}
func Device_array4_json_unmarshal(iter *jsoniter.Iterator, out *[]string) {
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
func Device_ptr5_json_unmarshal(iter *jsoniter.Iterator, out **string) {
	var val string
	iter.ReadString(&val)
	if iter.Error == nil {
		*out = &val
	}
}
func Device_ptr6_json_unmarshal(iter *jsoniter.Iterator, out **string) {
	var val string
	iter.ReadString(&val)
	if iter.Error == nil {
		*out = &val
	}
}
func Device_ptr7_json_unmarshal(iter *jsoniter.Iterator, out **string) {
	var val string
	iter.ReadString(&val)
	if iter.Error == nil {
		*out = &val
	}
}
func Device_ptr8_json_unmarshal(iter *jsoniter.Iterator, out **string) {
	var val string
	iter.ReadString(&val)
	if iter.Error == nil {
		*out = &val
	}
}
func Device_ptr9_json_unmarshal(iter *jsoniter.Iterator, out **string) {
	var val string
	iter.ReadString(&val)
	if iter.Error == nil {
		*out = &val
	}
}
func Device_ptr10_json_unmarshal(iter *jsoniter.Iterator, out **string) {
	var val string
	iter.ReadString(&val)
	if iter.Error == nil {
		*out = &val
	}
}
func Device_ptr11_json_unmarshal(iter *jsoniter.Iterator, out **string) {
	var val string
	iter.ReadString(&val)
	if iter.Error == nil {
		*out = &val
	}
}
func Device_ptr12_json_unmarshal(iter *jsoniter.Iterator, out **string) {
	var val string
	iter.ReadString(&val)
	if iter.Error == nil {
		*out = &val
	}
}
func Device_json_marshal(stream *jsoniter.Stream, val Device) {
	stream.WriteObjectHead()
	Device_json_marshal_field(stream, val)
	stream.WriteObjectTail()
}
func Device_json_marshal_field(stream *jsoniter.Stream, val Device) {
	stream.WriteObjectField(`configuration_url`)
	if val.ConfigurationURL == nil {
		stream.WriteNull()
	} else {
		stream.WriteString(*val.ConfigurationURL)
	}
	stream.WriteMore()
	stream.WriteObjectField(`connections`)
	Device_array13_json_marshal(stream, val.Connections)
	stream.WriteMore()
	stream.WriteObjectField(`hw_version`)
	if val.HardwareVersion == nil {
		stream.WriteNull()
	} else {
		stream.WriteString(*val.HardwareVersion)
	}
	stream.WriteMore()
	stream.WriteObjectField(`identifiers`)
	Device_array14_json_marshal(stream, val.Identifiers)
	stream.WriteMore()
	stream.WriteObjectField(`manufacturer`)
	if val.Manufacturer == nil {
		stream.WriteNull()
	} else {
		stream.WriteString(*val.Manufacturer)
	}
	stream.WriteMore()
	stream.WriteObjectField(`model`)
	if val.Model == nil {
		stream.WriteNull()
	} else {
		stream.WriteString(*val.Model)
	}
	stream.WriteMore()
	stream.WriteObjectField(`model_id`)
	if val.ModelID == nil {
		stream.WriteNull()
	} else {
		stream.WriteString(*val.ModelID)
	}
	stream.WriteMore()
	stream.WriteObjectField(`name`)
	if val.Name == nil {
		stream.WriteNull()
	} else {
		stream.WriteString(*val.Name)
	}
	stream.WriteMore()
	stream.WriteObjectField(`serial_number`)
	if val.SerialNumber == nil {
		stream.WriteNull()
	} else {
		stream.WriteString(*val.SerialNumber)
	}
	stream.WriteMore()
	stream.WriteObjectField(`suggested_area`)
	if val.SuggestedArea == nil {
		stream.WriteNull()
	} else {
		stream.WriteString(*val.SuggestedArea)
	}
	stream.WriteMore()
	stream.WriteObjectField(`software_version`)
	if val.SoftwareVersion == nil {
		stream.WriteNull()
	} else {
		stream.WriteString(*val.SoftwareVersion)
	}
	stream.WriteMore()
	stream.WriteObjectField(`via_device`)
	if val.ViaDevice == nil {
		stream.WriteNull()
	} else {
		stream.WriteString(*val.ViaDevice)
	}
	stream.WriteMore()
}
func Device_array13_json_marshal(stream *jsoniter.Stream, val []string) {
	if len(val) == 0 {
		stream.WriteEmptyArray()
	} else {
		stream.WriteArrayHead()
		for i, elem := range val {
			if i != 0 {
				stream.WriteMore()
			}
			stream.WriteString(elem)
		}
		stream.WriteArrayTail()
	}
}
func Device_array14_json_marshal(stream *jsoniter.Stream, val []string) {
	if len(val) == 0 {
		stream.WriteEmptyArray()
	} else {
		stream.WriteArrayHead()
		for i, elem := range val {
			if i != 0 {
				stream.WriteMore()
			}
			stream.WriteString(elem)
		}
		stream.WriteArrayTail()
	}
}
