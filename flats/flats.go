package flats

import "fmt"

func ToAndFromProto(uids []uint64) error {
	var ul UidList
	ul.Uid = make([]uint64, len(uids))
	copy(ul.Uid, uids)

	data, err := ul.Marshal()
	if err != nil {
		return err
	}
	var nl UidList
	if err := nl.Unmarshal(data); err != nil {
		return err
	}
	if len(nl.Uid) != len(ul.Uid) {
		return fmt.Errorf("Length doesn't match")
	}
	return nil
}

func ToAndFromFlat(uids []uint64) error {
	return nil
}
