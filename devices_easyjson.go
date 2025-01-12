// Code generated by easyjson for marshaling/unmarshaling. DO NOT EDIT.

package gomeas

import (
	json "encoding/json"
	easyjson "github.com/mailru/easyjson"
	jlexer "github.com/mailru/easyjson/jlexer"
	jwriter "github.com/mailru/easyjson/jwriter"
)

// suppress unused package warning
var (
	_ *json.RawMessage
	_ *jlexer.Lexer
	_ *jwriter.Writer
	_ easyjson.Marshaler
)

func easyjson65411fd3DecodeGithubComCharlesDBurtonGomeas(in *jlexer.Lexer, out *Device) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeFieldName(false)
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "configuration_url":
			if in.IsNull() {
				in.Skip()
				out.ConfigurationURL = nil
			} else {
				if out.ConfigurationURL == nil {
					out.ConfigurationURL = new(string)
				}
				*out.ConfigurationURL = string(in.String())
			}
		case "connections":
			if in.IsNull() {
				in.Skip()
				out.Connections = nil
			} else {
				in.Delim('[')
				if out.Connections == nil {
					if !in.IsDelim(']') {
						out.Connections = make([]string, 0, 4)
					} else {
						out.Connections = []string{}
					}
				} else {
					out.Connections = (out.Connections)[:0]
				}
				for !in.IsDelim(']') {
					var v1 string
					v1 = string(in.String())
					out.Connections = append(out.Connections, v1)
					in.WantComma()
				}
				in.Delim(']')
			}
		case "hw_version":
			if in.IsNull() {
				in.Skip()
				out.HardwareVersion = nil
			} else {
				if out.HardwareVersion == nil {
					out.HardwareVersion = new(string)
				}
				*out.HardwareVersion = string(in.String())
			}
		case "identifiers":
			if in.IsNull() {
				in.Skip()
				out.Identifiers = nil
			} else {
				in.Delim('[')
				if out.Identifiers == nil {
					if !in.IsDelim(']') {
						out.Identifiers = make([]string, 0, 4)
					} else {
						out.Identifiers = []string{}
					}
				} else {
					out.Identifiers = (out.Identifiers)[:0]
				}
				for !in.IsDelim(']') {
					var v2 string
					v2 = string(in.String())
					out.Identifiers = append(out.Identifiers, v2)
					in.WantComma()
				}
				in.Delim(']')
			}
		case "manufacturer":
			if in.IsNull() {
				in.Skip()
				out.Manufacturer = nil
			} else {
				if out.Manufacturer == nil {
					out.Manufacturer = new(string)
				}
				*out.Manufacturer = string(in.String())
			}
		case "model":
			if in.IsNull() {
				in.Skip()
				out.Model = nil
			} else {
				if out.Model == nil {
					out.Model = new(string)
				}
				*out.Model = string(in.String())
			}
		case "model_id":
			if in.IsNull() {
				in.Skip()
				out.ModelID = nil
			} else {
				if out.ModelID == nil {
					out.ModelID = new(string)
				}
				*out.ModelID = string(in.String())
			}
		case "name":
			if in.IsNull() {
				in.Skip()
				out.Name = nil
			} else {
				if out.Name == nil {
					out.Name = new(string)
				}
				*out.Name = string(in.String())
			}
		case "serial_number":
			if in.IsNull() {
				in.Skip()
				out.SerialNumber = nil
			} else {
				if out.SerialNumber == nil {
					out.SerialNumber = new(string)
				}
				*out.SerialNumber = string(in.String())
			}
		case "suggested_area":
			if in.IsNull() {
				in.Skip()
				out.SuggestedArea = nil
			} else {
				if out.SuggestedArea == nil {
					out.SuggestedArea = new(string)
				}
				*out.SuggestedArea = string(in.String())
			}
		case "software_version":
			if in.IsNull() {
				in.Skip()
				out.SoftwareVersion = nil
			} else {
				if out.SoftwareVersion == nil {
					out.SoftwareVersion = new(string)
				}
				*out.SoftwareVersion = string(in.String())
			}
		case "via_device":
			if in.IsNull() {
				in.Skip()
				out.ViaDevice = nil
			} else {
				if out.ViaDevice == nil {
					out.ViaDevice = new(string)
				}
				*out.ViaDevice = string(in.String())
			}
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}
func easyjson65411fd3EncodeGithubComCharlesDBurtonGomeas(out *jwriter.Writer, in Device) {
	out.RawByte('{')
	first := true
	_ = first
	if in.ConfigurationURL != nil {
		const prefix string = ",\"configuration_url\":"
		first = false
		out.RawString(prefix[1:])
		out.String(string(*in.ConfigurationURL))
	}
	if len(in.Connections) != 0 {
		const prefix string = ",\"connections\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		{
			out.RawByte('[')
			for v3, v4 := range in.Connections {
				if v3 > 0 {
					out.RawByte(',')
				}
				out.String(string(v4))
			}
			out.RawByte(']')
		}
	}
	if in.HardwareVersion != nil {
		const prefix string = ",\"hw_version\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(*in.HardwareVersion))
	}
	if len(in.Identifiers) != 0 {
		const prefix string = ",\"identifiers\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		{
			out.RawByte('[')
			for v5, v6 := range in.Identifiers {
				if v5 > 0 {
					out.RawByte(',')
				}
				out.String(string(v6))
			}
			out.RawByte(']')
		}
	}
	if in.Manufacturer != nil {
		const prefix string = ",\"manufacturer\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(*in.Manufacturer))
	}
	if in.Model != nil {
		const prefix string = ",\"model\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(*in.Model))
	}
	if in.ModelID != nil {
		const prefix string = ",\"model_id\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(*in.ModelID))
	}
	if in.Name != nil {
		const prefix string = ",\"name\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(*in.Name))
	}
	if in.SerialNumber != nil {
		const prefix string = ",\"serial_number\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(*in.SerialNumber))
	}
	if in.SuggestedArea != nil {
		const prefix string = ",\"suggested_area\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(*in.SuggestedArea))
	}
	if in.SoftwareVersion != nil {
		const prefix string = ",\"software_version\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(*in.SoftwareVersion))
	}
	if in.ViaDevice != nil {
		const prefix string = ",\"via_device\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(*in.ViaDevice))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v Device) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson65411fd3EncodeGithubComCharlesDBurtonGomeas(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v Device) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson65411fd3EncodeGithubComCharlesDBurtonGomeas(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *Device) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson65411fd3DecodeGithubComCharlesDBurtonGomeas(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *Device) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson65411fd3DecodeGithubComCharlesDBurtonGomeas(l, v)
}
func easyjson65411fd3DecodeGithubComCharlesDBurtonGomeas1(in *jlexer.Lexer, out *Config) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeFieldName(false)
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "availability":
			if in.IsNull() {
				in.Skip()
				out.Availability = nil
			} else {
				in.Delim('[')
				if out.Availability == nil {
					if !in.IsDelim(']') {
						out.Availability = make([]*Availability, 0, 8)
					} else {
						out.Availability = []*Availability{}
					}
				} else {
					out.Availability = (out.Availability)[:0]
				}
				for !in.IsDelim(']') {
					var v7 *Availability
					if in.IsNull() {
						in.Skip()
						v7 = nil
					} else {
						if v7 == nil {
							v7 = new(Availability)
						}
						(*v7).UnmarshalEasyJSON(in)
					}
					out.Availability = append(out.Availability, v7)
					in.WantComma()
				}
				in.Delim(']')
			}
		case "availibility_mode":
			if in.IsNull() {
				in.Skip()
				out.AvailibilityMode = nil
			} else {
				if out.AvailibilityMode == nil {
					out.AvailibilityMode = new(string)
				}
				*out.AvailibilityMode = string(in.String())
			}
		case "availability_template":
			if in.IsNull() {
				in.Skip()
				out.AvailabilityTemplate = nil
			} else {
				if out.AvailabilityTemplate == nil {
					out.AvailabilityTemplate = new(string)
				}
				*out.AvailabilityTemplate = string(in.String())
			}
		case "availability_topic":
			if in.IsNull() {
				in.Skip()
				out.AvailabilityTopic = nil
			} else {
				if out.AvailabilityTopic == nil {
					out.AvailabilityTopic = new(string)
				}
				*out.AvailabilityTopic = string(in.String())
			}
		case "name":
			out.Name = string(in.String())
		case "device_class":
			out.DeviceClass = string(in.String())
		case "state_topic":
			if in.IsNull() {
				in.Skip()
				out.StateTopic = nil
			} else {
				if out.StateTopic == nil {
					out.StateTopic = new(string)
				}
				*out.StateTopic = string(in.String())
			}
		case "command_topic":
			if in.IsNull() {
				in.Skip()
				out.CommandTopic = nil
			} else {
				if out.CommandTopic == nil {
					out.CommandTopic = new(string)
				}
				*out.CommandTopic = string(in.String())
			}
		case "unique_id":
			out.UniqueID = string(in.String())
		case "device":
			if in.IsNull() {
				in.Skip()
				out.Device = nil
			} else {
				if out.Device == nil {
					out.Device = new(Device)
				}
				(*out.Device).UnmarshalEasyJSON(in)
			}
		case "components":
			if in.IsNull() {
				in.Skip()
			} else {
				in.Delim('{')
				if !in.IsDelim('}') {
					out.Components = make(map[string]Component)
				} else {
					out.Components = nil
				}
				for !in.IsDelim('}') {
					key := string(in.String())
					in.WantColon()
					var v8 Component
					(v8).UnmarshalEasyJSON(in)
					(out.Components)[key] = v8
					in.WantComma()
				}
				in.Delim('}')
			}
		case "enabled_by_default":
			out.EnabledByDefault = bool(in.Bool())
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}
func easyjson65411fd3EncodeGithubComCharlesDBurtonGomeas1(out *jwriter.Writer, in Config) {
	out.RawByte('{')
	first := true
	_ = first
	if len(in.Availability) != 0 {
		const prefix string = ",\"availability\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		{
			out.RawByte('[')
			for v9, v10 := range in.Availability {
				if v9 > 0 {
					out.RawByte(',')
				}
				if v10 == nil {
					out.RawString("null")
				} else {
					(*v10).MarshalEasyJSON(out)
				}
			}
			out.RawByte(']')
		}
	}
	if in.AvailibilityMode != nil {
		const prefix string = ",\"availibility_mode\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(*in.AvailibilityMode))
	}
	if in.AvailabilityTemplate != nil {
		const prefix string = ",\"availability_template\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(*in.AvailabilityTemplate))
	}
	if in.AvailabilityTopic != nil {
		const prefix string = ",\"availability_topic\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(*in.AvailabilityTopic))
	}
	if in.Name != "" {
		const prefix string = ",\"name\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.Name))
	}
	if in.DeviceClass != "" {
		const prefix string = ",\"device_class\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.DeviceClass))
	}
	if in.StateTopic != nil {
		const prefix string = ",\"state_topic\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(*in.StateTopic))
	}
	if in.CommandTopic != nil {
		const prefix string = ",\"command_topic\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(*in.CommandTopic))
	}
	if in.UniqueID != "" {
		const prefix string = ",\"unique_id\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.UniqueID))
	}
	if in.Device != nil {
		const prefix string = ",\"device\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		(*in.Device).MarshalEasyJSON(out)
	}
	if len(in.Components) != 0 {
		const prefix string = ",\"components\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		{
			out.RawByte('{')
			v11First := true
			for v11Name, v11Value := range in.Components {
				if v11First {
					v11First = false
				} else {
					out.RawByte(',')
				}
				out.String(string(v11Name))
				out.RawByte(':')
				(v11Value).MarshalEasyJSON(out)
			}
			out.RawByte('}')
		}
	}
	if in.EnabledByDefault {
		const prefix string = ",\"enabled_by_default\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Bool(bool(in.EnabledByDefault))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v Config) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson65411fd3EncodeGithubComCharlesDBurtonGomeas1(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v Config) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson65411fd3EncodeGithubComCharlesDBurtonGomeas1(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *Config) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson65411fd3DecodeGithubComCharlesDBurtonGomeas1(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *Config) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson65411fd3DecodeGithubComCharlesDBurtonGomeas1(l, v)
}
func easyjson65411fd3DecodeGithubComCharlesDBurtonGomeas2(in *jlexer.Lexer, out *Component) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeFieldName(false)
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "p":
			out.P = string(in.String())
		case "device_class":
			out.DeviceClass = string(in.String())
		case "unit_of_measurement":
			out.UnitOfMeasurement = string(in.String())
		case "value_template":
			out.ValueTemplate = string(in.String())
		case "unique_id":
			out.UniqueID = string(in.String())
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}
func easyjson65411fd3EncodeGithubComCharlesDBurtonGomeas2(out *jwriter.Writer, in Component) {
	out.RawByte('{')
	first := true
	_ = first
	if in.P != "" {
		const prefix string = ",\"p\":"
		first = false
		out.RawString(prefix[1:])
		out.String(string(in.P))
	}
	if in.DeviceClass != "" {
		const prefix string = ",\"device_class\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.DeviceClass))
	}
	if in.UnitOfMeasurement != "" {
		const prefix string = ",\"unit_of_measurement\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.UnitOfMeasurement))
	}
	if in.ValueTemplate != "" {
		const prefix string = ",\"value_template\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.ValueTemplate))
	}
	if in.UniqueID != "" {
		const prefix string = ",\"unique_id\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.UniqueID))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v Component) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson65411fd3EncodeGithubComCharlesDBurtonGomeas2(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v Component) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson65411fd3EncodeGithubComCharlesDBurtonGomeas2(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *Component) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson65411fd3DecodeGithubComCharlesDBurtonGomeas2(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *Component) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson65411fd3DecodeGithubComCharlesDBurtonGomeas2(l, v)
}
func easyjson65411fd3DecodeGithubComCharlesDBurtonGomeas3(in *jlexer.Lexer, out *Availability) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeFieldName(false)
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "payload_available":
			out.PayloadAvailable = string(in.String())
		case "payload_not_available":
			out.PayloadNotAvailable = string(in.String())
		case "topic":
			out.Topic = string(in.String())
		case "value_template":
			out.ValueTemplate = string(in.String())
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}
func easyjson65411fd3EncodeGithubComCharlesDBurtonGomeas3(out *jwriter.Writer, in Availability) {
	out.RawByte('{')
	first := true
	_ = first
	if in.PayloadAvailable != "" {
		const prefix string = ",\"payload_available\":"
		first = false
		out.RawString(prefix[1:])
		out.String(string(in.PayloadAvailable))
	}
	if in.PayloadNotAvailable != "" {
		const prefix string = ",\"payload_not_available\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.PayloadNotAvailable))
	}
	if in.Topic != "" {
		const prefix string = ",\"topic\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.Topic))
	}
	if in.ValueTemplate != "" {
		const prefix string = ",\"value_template\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.ValueTemplate))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v Availability) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson65411fd3EncodeGithubComCharlesDBurtonGomeas3(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v Availability) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson65411fd3EncodeGithubComCharlesDBurtonGomeas3(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *Availability) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson65411fd3DecodeGithubComCharlesDBurtonGomeas3(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *Availability) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson65411fd3DecodeGithubComCharlesDBurtonGomeas3(l, v)
}
