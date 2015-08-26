package main

// AUTO GENERATED - DO NOT EDIT

import (
	"bufio"
	"bytes"
	"encoding/json"
	C "github.com/glycerine/go-capnproto"
	"io"
)

type UidArrayCapn C.Struct

func NewUidArrayCapn(s *C.Segment) UidArrayCapn      { return UidArrayCapn(s.NewStruct(0, 1)) }
func NewRootUidArrayCapn(s *C.Segment) UidArrayCapn  { return UidArrayCapn(s.NewRootStruct(0, 1)) }
func AutoNewUidArrayCapn(s *C.Segment) UidArrayCapn  { return UidArrayCapn(s.NewStructAR(0, 1)) }
func ReadRootUidArrayCapn(s *C.Segment) UidArrayCapn { return UidArrayCapn(s.Root(0).ToStruct()) }
func (s UidArrayCapn) Uids() C.UInt64List            { return C.UInt64List(C.Struct(s).GetObject(0)) }
func (s UidArrayCapn) SetUids(v C.UInt64List)        { C.Struct(s).SetObject(0, C.Object(v)) }
func (s UidArrayCapn) WriteJSON(w io.Writer) error {
	b := bufio.NewWriter(w)
	var err error
	var buf []byte
	_ = buf
	err = b.WriteByte('{')
	if err != nil {
		return err
	}
	_, err = b.WriteString("\"uids\":")
	if err != nil {
		return err
	}
	{
		s := s.Uids()
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
	err = b.WriteByte('}')
	if err != nil {
		return err
	}
	err = b.Flush()
	return err
}
func (s UidArrayCapn) MarshalJSON() ([]byte, error) {
	b := bytes.Buffer{}
	err := s.WriteJSON(&b)
	return b.Bytes(), err
}
func (s UidArrayCapn) WriteCapLit(w io.Writer) error {
	b := bufio.NewWriter(w)
	var err error
	var buf []byte
	_ = buf
	err = b.WriteByte('(')
	if err != nil {
		return err
	}
	_, err = b.WriteString("uids = ")
	if err != nil {
		return err
	}
	{
		s := s.Uids()
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
	err = b.WriteByte(')')
	if err != nil {
		return err
	}
	err = b.Flush()
	return err
}
func (s UidArrayCapn) MarshalCapLit() ([]byte, error) {
	b := bytes.Buffer{}
	err := s.WriteCapLit(&b)
	return b.Bytes(), err
}

type UidArrayCapn_List C.PointerList

func NewUidArrayCapnList(s *C.Segment, sz int) UidArrayCapn_List {
	return UidArrayCapn_List(s.NewCompositeList(0, 1, sz))
}
func (s UidArrayCapn_List) Len() int { return C.PointerList(s).Len() }
func (s UidArrayCapn_List) At(i int) UidArrayCapn {
	return UidArrayCapn(C.PointerList(s).At(i).ToStruct())
}
func (s UidArrayCapn_List) ToArray() []UidArrayCapn {
	n := s.Len()
	a := make([]UidArrayCapn, n)
	for i := 0; i < n; i++ {
		a[i] = s.At(i)
	}
	return a
}
func (s UidArrayCapn_List) Set(i int, item UidArrayCapn) { C.PointerList(s).Set(i, C.Object(item)) }
