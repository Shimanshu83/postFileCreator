package application

type Record struct {
	Value  string
	End    int
	Length int
}

func HeaderDeepCopy() map[string]Record {

	// key will be name of the field of struct from data base.
	// Value will be empty for a some field
	header := map[string]Record{
		"DPID": {
			Value:  "IN307865",
			End:    34,
			Length: 43,
		},
		"BatchNo": {
			Value:  "IN307865",
			End:    34,
			Length: 43,
		},
		"BrachCode": {
			Value:  "IN307865",
			End:    34,
			Length: 43,
		},

		"DetailType": {
			Value:  "IN307865",
			End:    34,
			Length: 43,
		},
		"RecordLength": {
			Value:  "IN307865",
			End:    34,
			Length: 43,
		},
	}

	// return by value not by object so it will work for me and all of us.
	return header
}
