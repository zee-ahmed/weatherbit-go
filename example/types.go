package main

// Nowcast struct holds relevant information for each observation, can be saved to a slice or a map
type Nowcast struct {
	lat, lon                    float64
	dhi, windspeed, temperature float64
	cloudcover                  float64
	observationage              int64
	city                        string
	// here can be extended to have additional metrics (policies and weighting)
	// for other currently unmeasurable energy sources (i.e hydroelectric)
}

// Pair required to move from a map to a slice for sorting
// Golang-nuts by Andrew Gerrand
// https://stackoverflow.com/questions/18695346/how-to-sort-a-mapstringint-by-its-values
type Pair struct {
	Key   string
	Value Nowcast
}

// A data structure is partially persistent if all versions can be accessed but only the newest version can be modified.

// PairList required to move from a map to a slice for sorting
type PairList []Pair

func (p PairList) Len() int           { return len(p) }
func (p PairList) Less(i, j int) bool { return p[i].Value.dhi < p[j].Value.dhi } // extensible
func (p PairList) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }

func tidyresults(body interface{}, granularity string) interface{} {

	granularity = "what"

	return body
}
