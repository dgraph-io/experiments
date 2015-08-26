package plist

// AUTO GENERATED - DO NOT EDIT

import (
	"bufio"
	"bytes"
	"encoding/json"
	C "github.com/glycerine/go-capnproto"
	"io"
)

type PostingList C.Struct

func NewPostingList(s *C.Segment) PostingList      { return PostingList(s.NewStruct(0, 2)) }
func NewRootPostingList(s *C.Segment) PostingList  { return PostingList(s.NewRootStruct(0, 2)) }
func AutoNewPostingList(s *C.Segment) PostingList  { return PostingList(s.NewStructAR(0, 2)) }
func ReadRootPostingList(s *C.Segment) PostingList { return PostingList(s.Root(0).ToStruct()) }
func (s PostingList) Ids() C.UInt64List            { return C.UInt64List(C.Struct(s).GetObject(0)) }
func (s PostingList) SetIds(v C.UInt64List)        { C.Struct(s).SetObject(0, C.Object(v)) }
func (s PostingList) Title() string                { return C.Struct(s).GetObject(1).ToText() }
func (s PostingList) SetTitle(v string)            { C.Struct(s).SetObject(1, s.Segment.NewText(v)) }
func (s PostingList) WriteJSON(w io.Writer) error {
	b := bufio.NewWriter(w)
	var err error
	var buf []byte
	_ = buf
	err = b.WriteByte('{')
	if err != nil {
		return err
	}
	_, err = b.WriteString("\"ids\":")
	if err != nil {
		return err
	}
	{
		s := s.Ids()
		{
			err = b.WriteByte('[')
			if err != nil {
				return err
			}
			for i, s := range s.ToArray() {
				if i != 0 {
					_, err = b.WriteString(", ")
				}
				if err != nil {
					return err
				}
				buf, err = json.Marshal(s)
				if err != nil {
					return err
				}
				_, err = b.Write(buf)
				if err != nil {
					return err
				}
			}
			err = b.WriteByte(']')
		}
		if err != nil {
			return err
		}
	}
	err = b.WriteByte(',')
	if err != nil {
		return err
	}
	_, err = b.WriteString("\"title\":")
	if err != nil {
		return err
	}
	{
		s := s.Title()
		buf, err = json.Marshal(s)
		if err != nil {
			return err
		}
		_, err = b.Write(buf)
		if err != nil {
			return err
		}
	}
	err = b.WriteByte('}')
	if err != nil {
		return err
	}
	err = b.Flush()
	return err
}
func (s PostingList) MarshalJSON() ([]byte, error) {
	b := bytes.Buffer{}
	err := s.WriteJSON(&b)
	return b.Bytes(), err
}
func (s PostingList) WriteCapLit(w io.Writer) error {
	b := bufio.NewWriter(w)
	var err error
	var buf []byte
	_ = buf
	err = b.WriteByte('(')
	if err != nil {
		return err
	}
	_, err = b.WriteString("ids = ")
	if err != nil {
		return err
	}
	{
		s := s.Ids()
		{
			err = b.WriteByte('[')
			if err != nil {
				return err
			}
			for i, s := range s.ToArray() {
				if i != 0 {
					_, err = b.WriteString(", ")
				}
				if err != nil {
					return err
				}
				buf, err = json.Marshal(s)
				if err != nil {
					return err
				}
				_, err = b.Write(buf)
				if err != nil {
					return err
				}
			}
			err = b.WriteByte(']')
		}
		if err != nil {
			return err
		}
	}
	_, err = b.WriteString(", ")
	if err != nil {
		return err
	}
	_, err = b.WriteString("title = ")
	if err != nil {
		return err
	}
	{
		s := s.Title()
		buf, err = json.Marshal(s)
		if err != nil {
			return err
		}
		_, err = b.Write(buf)
		if err != nil {
			return err
		}
	}
	err = b.WriteByte(')')
	if err != nil {
		return err
	}
	err = b.Flush()
	return err
}
func (s PostingList) MarshalCapLit() ([]byte, error) {
	b := bytes.Buffer{}
	err := s.WriteCapLit(&b)
	return b.Bytes(), err
}

type PostingList_List C.PointerList

func NewPostingListList(s *C.Segment, sz int) PostingList_List {
	return PostingList_List(s.NewCompositeList(0, 2, sz))
}
func (s PostingList_List) Len() int             { return C.PointerList(s).Len() }
func (s PostingList_List) At(i int) PostingList { return PostingList(C.PointerList(s).At(i).ToStruct()) }
func (s PostingList_List) ToArray() []PostingList {
	n := s.Len()
	a := make([]PostingList, n)
	for i := 0; i < n; i++ {
		a[i] = s.At(i)
	}
	return a
}
func (s PostingList_List) Set(i int, item PostingList) { C.PointerList(s).Set(i, C.Object(item)) }
