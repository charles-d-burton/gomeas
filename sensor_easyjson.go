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

func easyjsonCf4941d2DecodeGithubComCharlesDBurtonGomeas(in *jlexer.Lexer, out *Sensor) {
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
		case "platform":
			out.Platform = string(in.String())
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
					var v1 *Availability
					if in.IsNull() {
						in.Skip()
						v1 = nil
					} else {
						if v1 == nil {
							v1 = new(Availability)
						}
						(*v1).UnmarshalEasyJSON(in)
					}
					out.Availability = append(out.Availability, v1)
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
					var v2 Component
					(v2).UnmarshalEasyJSON(in)
					(out.Components)[key] = v2
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
func easyjsonCf4941d2EncodeGithubComCharlesDBurtonGomeas(out *jwriter.Writer, in Sensor) {
	out.RawByte('{')
	first := true
	_ = first
	if in.Platform != "" {
		const prefix string = ",\"platform\":"
		first = false
		out.RawString(prefix[1:])
		out.String(string(in.Platform))
	}
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
			for v3, v4 := range in.Availability {
				if v3 > 0 {
					out.RawByte(',')
				}
				if v4 == nil {
					out.RawString("null")
				} else {
					(*v4).MarshalEasyJSON(out)
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
			v5First := true
			for v5Name, v5Value := range in.Components {
				if v5First {
					v5First = false
				} else {
					out.RawByte(',')
				}
				out.String(string(v5Name))
				out.RawByte(':')
				(v5Value).MarshalEasyJSON(out)
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
func (v Sensor) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonCf4941d2EncodeGithubComCharlesDBurtonGomeas(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v Sensor) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonCf4941d2EncodeGithubComCharlesDBurtonGomeas(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *Sensor) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonCf4941d2DecodeGithubComCharlesDBurtonGomeas(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *Sensor) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonCf4941d2DecodeGithubComCharlesDBurtonGomeas(l, v)
}
