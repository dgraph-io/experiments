package main

import (
  capn "github.com/glycerine/go-capnproto"
  "io"
)




func (s *UidArray) Save(w io.Writer) error {
  	seg := capn.NewBuffer(nil)
  	UidArrayGoToCapn(seg, s)
    _, err := seg.WriteTo(w)
    return err
}
 


func (s *UidArray) Load(r io.Reader) error {
  	capMsg, err := capn.ReadFromStream(r, nil)
  	if err != nil {
  		//panic(fmt.Errorf("capn.ReadFromStream error: %s", err))
        return err
  	}
  	z := ReadRootUidArrayCapn(capMsg)
      UidArrayCapnToGo(z, s)
   return nil
}



func UidArrayCapnToGo(src UidArrayCapn, dest *UidArray) *UidArray {
  if dest == nil {
    dest = &UidArray{}
  }

  var n int

    // Uids
	n = src.Uids().Len()
	dest.Uids = make([]uint64, n)
	for i := 0; i < n; i++ {
        dest.Uids[i] = uint64(src.Uids().At(i))
    }


  return dest
}



func UidArrayGoToCapn(seg *capn.Segment, src *UidArray) UidArrayCapn {
  dest := AutoNewUidArrayCapn(seg)


  mylist1 := seg.NewUInt64List(len(src.Uids))
  for i := range src.Uids {
     mylist1.Set(i, uint64(src.Uids[i]))
  }
  dest.SetUids(mylist1)

  return dest
}



func SliceUint64ToUInt64List(seg *capn.Segment, m []uint64) capn.UInt64List {
	lst := seg.NewUInt64List(len(m))
	for i := range m {
		lst.Set(i, uint64(m[i]))
	}
	return lst
}



func UInt64ListToSliceUint64(p capn.UInt64List) []uint64 {
	v := make([]uint64, p.Len())
	for i := range v {
        v[i] = uint64(p.At(i))
	}
	return v
}
