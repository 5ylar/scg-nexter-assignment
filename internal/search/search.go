package search

import "strings"

type search struct {
}

func New() *search {
	return &search{}
}

func (s search) SearchMultipleKeysFromStringDataset(keys []string, dataset []string, caseSensitive bool) []ResultFromStringDataset {
	var rs []ResultFromStringDataset

	if len(dataset) == 0 {
		return rs
	}

	for _, k := range keys {
		r := ResultFromStringDataset{
			Key: k,
		}

		for i, d := range dataset {
			if !caseSensitive {
				d = strings.ToLower(d)
				k = strings.ToLower(k)
			}

			if d == k {
				r.Positions = append(r.Positions, i)
			}
		}

		rs = append(rs, r)
	}

	return rs
}

type ResultFromStringDataset struct {
	Key       string `json:"key"`
	Positions []int  `json:"positions"` // start with 0
}
