// Code generated by easyjson for marshaling/unmarshaling. DO NOT EDIT.

package entity

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

func easyjsonD2b7633eDecodeGitCanopsisNetCanopsisCanopsisCommunityCommunityGoEnginesCommunityLibApiEntity(in *jlexer.Lexer, out *BaseFilterRequest) {
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
		case "search":
			out.Search = string(in.String())
		case "filters":
			if in.IsNull() {
				in.Skip()
				out.Filters = nil
			} else {
				in.Delim('[')
				if out.Filters == nil {
					if !in.IsDelim(']') {
						out.Filters = make([]string, 0, 4)
					} else {
						out.Filters = []string{}
					}
				} else {
					out.Filters = (out.Filters)[:0]
				}
				for !in.IsDelim(']') {
					var v1 string
					v1 = string(in.String())
					out.Filters = append(out.Filters, v1)
					in.WantComma()
				}
				in.Delim(']')
			}
		case "category":
			out.Category = string(in.String())
		case "type":
			if in.IsNull() {
				in.Skip()
				out.Type = nil
			} else {
				in.Delim('[')
				if out.Type == nil {
					if !in.IsDelim(']') {
						out.Type = make([]string, 0, 4)
					} else {
						out.Type = []string{}
					}
				} else {
					out.Type = (out.Type)[:0]
				}
				for !in.IsDelim(']') {
					var v2 string
					v2 = string(in.String())
					out.Type = append(out.Type, v2)
					in.WantComma()
				}
				in.Delim(']')
			}
		case "no_events":
			out.NoEvents = bool(in.Bool())
		case "entity_pattern":
			out.EntityPattern = string(in.String())
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
func easyjsonD2b7633eEncodeGitCanopsisNetCanopsisCanopsisCommunityCommunityGoEnginesCommunityLibApiEntity(out *jwriter.Writer, in BaseFilterRequest) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"search\":"
		out.RawString(prefix[1:])
		out.String(string(in.Search))
	}
	{
		const prefix string = ",\"filters\":"
		out.RawString(prefix)
		if in.Filters == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
			out.RawString("null")
		} else {
			out.RawByte('[')
			for v3, v4 := range in.Filters {
				if v3 > 0 {
					out.RawByte(',')
				}
				out.String(string(v4))
			}
			out.RawByte(']')
		}
	}
	{
		const prefix string = ",\"category\":"
		out.RawString(prefix)
		out.String(string(in.Category))
	}
	{
		const prefix string = ",\"type\":"
		out.RawString(prefix)
		if in.Type == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
			out.RawString("null")
		} else {
			out.RawByte('[')
			for v5, v6 := range in.Type {
				if v5 > 0 {
					out.RawByte(',')
				}
				out.String(string(v6))
			}
			out.RawByte(']')
		}
	}
	{
		const prefix string = ",\"no_events\":"
		out.RawString(prefix)
		out.Bool(bool(in.NoEvents))
	}
	{
		const prefix string = ",\"entity_pattern\":"
		out.RawString(prefix)
		out.String(string(in.EntityPattern))
	}
	out.RawByte('}')
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v BaseFilterRequest) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonD2b7633eEncodeGitCanopsisNetCanopsisCanopsisCommunityCommunityGoEnginesCommunityLibApiEntity(w, v)
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *BaseFilterRequest) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonD2b7633eDecodeGitCanopsisNetCanopsisCanopsisCommunityCommunityGoEnginesCommunityLibApiEntity(l, v)
}
