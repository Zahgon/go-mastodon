package mastodon

type ID string

func (id *ID) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

// Compare compares the Mastodon IDs i and j.
// Compare returns:
//
//	-1 if i is less than j,
//	 0 if i equals j,
//	-1 if j is greater than i.
//
// Compare can be used as an argument of [slices.SortFunc]:
//
//	slices.SortFunc([]mastodon.ID{id1, id2}, mastodon.ID.Compare)
func (i ID) Compare(j ID) int { _ = "STUB: not implemented"; return 0 }

func (i ID) u64() uint64 { _ = "STUB: not implemented"; return 0 }

type Sbool bool

func (s *Sbool) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }
