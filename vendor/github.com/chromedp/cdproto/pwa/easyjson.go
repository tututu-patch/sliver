// Code generated by easyjson for marshaling/unmarshaling. DO NOT EDIT.

package pwa

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

func easyjsonC5a4559bDecodeGithubComChromedpCdprotoPwa(in *jlexer.Lexer, out *GetOsAppStateReturns) {
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
		case "badgeCount":
			out.BadgeCount = int64(in.Int64())
		case "fileHandlers":
			if in.IsNull() {
				in.Skip()
				out.FileHandlers = nil
			} else {
				in.Delim('[')
				if out.FileHandlers == nil {
					if !in.IsDelim(']') {
						out.FileHandlers = make([]*FileHandler, 0, 8)
					} else {
						out.FileHandlers = []*FileHandler{}
					}
				} else {
					out.FileHandlers = (out.FileHandlers)[:0]
				}
				for !in.IsDelim(']') {
					var v1 *FileHandler
					if in.IsNull() {
						in.Skip()
						v1 = nil
					} else {
						if v1 == nil {
							v1 = new(FileHandler)
						}
						(*v1).UnmarshalEasyJSON(in)
					}
					out.FileHandlers = append(out.FileHandlers, v1)
					in.WantComma()
				}
				in.Delim(']')
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
func easyjsonC5a4559bEncodeGithubComChromedpCdprotoPwa(out *jwriter.Writer, in GetOsAppStateReturns) {
	out.RawByte('{')
	first := true
	_ = first
	if in.BadgeCount != 0 {
		const prefix string = ",\"badgeCount\":"
		first = false
		out.RawString(prefix[1:])
		out.Int64(int64(in.BadgeCount))
	}
	if len(in.FileHandlers) != 0 {
		const prefix string = ",\"fileHandlers\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		{
			out.RawByte('[')
			for v2, v3 := range in.FileHandlers {
				if v2 > 0 {
					out.RawByte(',')
				}
				if v3 == nil {
					out.RawString("null")
				} else {
					(*v3).MarshalEasyJSON(out)
				}
			}
			out.RawByte(']')
		}
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v GetOsAppStateReturns) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonC5a4559bEncodeGithubComChromedpCdprotoPwa(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v GetOsAppStateReturns) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonC5a4559bEncodeGithubComChromedpCdprotoPwa(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *GetOsAppStateReturns) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonC5a4559bDecodeGithubComChromedpCdprotoPwa(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *GetOsAppStateReturns) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonC5a4559bDecodeGithubComChromedpCdprotoPwa(l, v)
}
func easyjsonC5a4559bDecodeGithubComChromedpCdprotoPwa1(in *jlexer.Lexer, out *GetOsAppStateParams) {
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
		case "manifestId":
			out.ManifestID = string(in.String())
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
func easyjsonC5a4559bEncodeGithubComChromedpCdprotoPwa1(out *jwriter.Writer, in GetOsAppStateParams) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"manifestId\":"
		out.RawString(prefix[1:])
		out.String(string(in.ManifestID))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v GetOsAppStateParams) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonC5a4559bEncodeGithubComChromedpCdprotoPwa1(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v GetOsAppStateParams) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonC5a4559bEncodeGithubComChromedpCdprotoPwa1(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *GetOsAppStateParams) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonC5a4559bDecodeGithubComChromedpCdprotoPwa1(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *GetOsAppStateParams) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonC5a4559bDecodeGithubComChromedpCdprotoPwa1(l, v)
}
func easyjsonC5a4559bDecodeGithubComChromedpCdprotoPwa2(in *jlexer.Lexer, out *FileHandlerAccept) {
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
		case "mediaType":
			out.MediaType = string(in.String())
		case "fileExtensions":
			if in.IsNull() {
				in.Skip()
				out.FileExtensions = nil
			} else {
				in.Delim('[')
				if out.FileExtensions == nil {
					if !in.IsDelim(']') {
						out.FileExtensions = make([]string, 0, 4)
					} else {
						out.FileExtensions = []string{}
					}
				} else {
					out.FileExtensions = (out.FileExtensions)[:0]
				}
				for !in.IsDelim(']') {
					var v4 string
					v4 = string(in.String())
					out.FileExtensions = append(out.FileExtensions, v4)
					in.WantComma()
				}
				in.Delim(']')
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
func easyjsonC5a4559bEncodeGithubComChromedpCdprotoPwa2(out *jwriter.Writer, in FileHandlerAccept) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"mediaType\":"
		out.RawString(prefix[1:])
		out.String(string(in.MediaType))
	}
	{
		const prefix string = ",\"fileExtensions\":"
		out.RawString(prefix)
		if in.FileExtensions == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
			out.RawString("null")
		} else {
			out.RawByte('[')
			for v5, v6 := range in.FileExtensions {
				if v5 > 0 {
					out.RawByte(',')
				}
				out.String(string(v6))
			}
			out.RawByte(']')
		}
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v FileHandlerAccept) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonC5a4559bEncodeGithubComChromedpCdprotoPwa2(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v FileHandlerAccept) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonC5a4559bEncodeGithubComChromedpCdprotoPwa2(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *FileHandlerAccept) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonC5a4559bDecodeGithubComChromedpCdprotoPwa2(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *FileHandlerAccept) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonC5a4559bDecodeGithubComChromedpCdprotoPwa2(l, v)
}
func easyjsonC5a4559bDecodeGithubComChromedpCdprotoPwa3(in *jlexer.Lexer, out *FileHandler) {
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
		case "action":
			out.Action = string(in.String())
		case "accepts":
			if in.IsNull() {
				in.Skip()
				out.Accepts = nil
			} else {
				in.Delim('[')
				if out.Accepts == nil {
					if !in.IsDelim(']') {
						out.Accepts = make([]*FileHandlerAccept, 0, 8)
					} else {
						out.Accepts = []*FileHandlerAccept{}
					}
				} else {
					out.Accepts = (out.Accepts)[:0]
				}
				for !in.IsDelim(']') {
					var v7 *FileHandlerAccept
					if in.IsNull() {
						in.Skip()
						v7 = nil
					} else {
						if v7 == nil {
							v7 = new(FileHandlerAccept)
						}
						(*v7).UnmarshalEasyJSON(in)
					}
					out.Accepts = append(out.Accepts, v7)
					in.WantComma()
				}
				in.Delim(']')
			}
		case "displayName":
			out.DisplayName = string(in.String())
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
func easyjsonC5a4559bEncodeGithubComChromedpCdprotoPwa3(out *jwriter.Writer, in FileHandler) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"action\":"
		out.RawString(prefix[1:])
		out.String(string(in.Action))
	}
	{
		const prefix string = ",\"accepts\":"
		out.RawString(prefix)
		if in.Accepts == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
			out.RawString("null")
		} else {
			out.RawByte('[')
			for v8, v9 := range in.Accepts {
				if v8 > 0 {
					out.RawByte(',')
				}
				if v9 == nil {
					out.RawString("null")
				} else {
					(*v9).MarshalEasyJSON(out)
				}
			}
			out.RawByte(']')
		}
	}
	{
		const prefix string = ",\"displayName\":"
		out.RawString(prefix)
		out.String(string(in.DisplayName))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v FileHandler) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonC5a4559bEncodeGithubComChromedpCdprotoPwa3(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v FileHandler) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonC5a4559bEncodeGithubComChromedpCdprotoPwa3(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *FileHandler) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonC5a4559bDecodeGithubComChromedpCdprotoPwa3(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *FileHandler) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonC5a4559bDecodeGithubComChromedpCdprotoPwa3(l, v)
}
