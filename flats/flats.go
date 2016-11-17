package flats

import (
	"fmt"

	"github.com/dgraph-io/experiments/flats/fuids"
	flatbuffers "github.com/google/flatbuffers/go"
)

func ToAndFrom() {
}

func ToAndFromProto(uids []uint64) (error, int) {
	var ul UidList
	ul.Uid = make([]uint64, len(uids))
	copy(ul.Uid, uids)

	data, err := ul.Marshal()
	if err != nil {
		return err, 0
	}
	var nl UidList
	if err := nl.Unmarshal(data); err != nil {
		return err, 0
	}
	if len(nl.Uid) != len(ul.Uid) {
		return fmt.Errorf("Length doesn't match"), 0
	}
	for i := 0; i < len(uids); i++ {
		if nl.Uid[i] != uids[i] {
			return fmt.Errorf("ID doesn't match at index: %v", i), 0
		}
	}
	return nil, len(data)
}

func ToAndFromProtoAlt(uids []uint64) (error, int) {
	var ul UidListAlt
	ul.Uid = make([]uint64, len(uids))
	copy(ul.Uid, uids)

	data, err := ul.Marshal()
	if err != nil {
		return err, 0
	}
	var nl UidListAlt
	if err := nl.Unmarshal(data); err != nil {
		return err, 0
	}
	if len(nl.Uid) != len(ul.Uid) {
		return fmt.Errorf("Length doesn't match"), 0
	}
	for i := 0; i < len(uids); i++ {
		if nl.Uid[i] != uids[i] {
			return fmt.Errorf("ID doesn't match at index: %v", i), 0
		}
	}
	return nil, len(data)
}

func ToAndFromFlat(uids []uint64) (error, int) {
	b := flatbuffers.NewBuilder(0)
	fuids.UidListStartUidsVector(b, len(uids))
	for i := len(uids) - 1; i >= 0; i-- {
		b.PrependUint64(uids[i])
	}
	ve := b.EndVector(len(uids))
	fuids.UidListStart(b)
	fuids.UidListAddUids(b, ve)
	ue := fuids.UidListEnd(b)
	b.Finish(ue)
	data := b.FinishedBytes()

	nl := fuids.GetRootAsUidList(data, 0)
	if nl.UidsLength() != len(uids) {
		return fmt.Errorf("Length doesn't match"), 0
	}
	for i := 0; i < len(uids); i++ {
		if nl.Uids(i) != uids[i] {
			return fmt.Errorf("ID doesn't match at index: %v Expected: %v. Got: %v",
				i, uids[i], nl.Uids(i)), 0
		}
	}

	return nil, len(data)
}
