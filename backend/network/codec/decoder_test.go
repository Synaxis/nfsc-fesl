package codec

import (
	"fmt"
	"strconv"
)

type userStats struct {
	OwnerID   int
	OwnerType int
	Stats     []entry
}

type entry struct {
	Key        string      `fesl:"k"`  // Example: c_ltp
	PointType  []int       `fesl:"pt"` // Example: "0,2" or just 0
	Text       string      `fesl:"t"`  // Example: ""
	UpdateType int         `fesl:"ut"` // Example: enum, 0 or 3
	Value      interface{} `fesl:"v"`  // Example: 9025.0000 or just 100
	Precision  int
}

func parseObject(d Fields) []userStats {
	uv := d["u.[]"]
	uc, _ := strconv.Atoi(uv)
	conts := make([]userStats, uc)
	for u := 0; u < uc; u++ {
		conts[u] = parseUser(d, u)
	}
	return conts
}

func parseUser(d Fields, u int) userStats {
	prefix := fmt.Sprintf("u.%d.", u)

	ownerID, _ := d.IntVal(prefix + "o")
	ownerType, _ := d.IntVal(prefix + "ot")

	sc, _ := d.IntVal(prefix + "s.[]")
	stats := make([]entry, sc)

	for s := 0; s < sc; s++ {
		stats[s] = parseStat(d, u, s)
	}

	c := userStats{
		Stats:     stats,
		OwnerType: ownerType,
		OwnerID:   ownerID,
	}

	return c
}

const (
	_ = iota
	utInsert
	_
	_
	utUpdate
)

func parseStat(d Fields, u, s int) entry {
	prefix := fmt.Sprintf("u.%d.s.%d.", u, s)

	ut, _ := d.IntVal(prefix + "ut")

	pt := d.IntArr(prefix+"pt", ",")
	val, prec, _ := d.FloatAsInt(prefix + "v")

	p := entry{
		Key:        d[prefix+"k"],
		Text:       d[prefix+"t"],
		PointType:  pt,
		UpdateType: ut,
		Value:      val,
		Precision:  prec,
	}
	return p
}

func ExampleDecodeStats() {
	d := DecodeFESL([]byte(`TXN=UpdateStats
u.0.o=729
u.0.ot=1
u.0.s.[]=6
u.0.s.0.ut=3
u.0.s.0.k=cpc
u.0.s.0.v=1.0000
u.0.s.0.t=
u.0.s.0.pt=0
u.0.s.1.ut=3
u.0.s.1.k=ct
u.0.s.1.v=1787.8120
u.0.s.1.t=
u.0.s.1.pt=0
u.0.s.2.ut=3
u.0.s.2.k=dfv5
u.0.s.2.v=8.0000
u.0.s.2.t=
u.0.s.2.pt=0
u.0.s.3.ut=3
u.0.s.3.k=dfw3201
u.0.s.3.v=1.0000
u.0.s.3.t=
u.0.s.3.pt=0
u.0.s.4.ut=3
u.0.s.4.k=m_ct12
u.0.s.4.v=108.9059
u.0.s.4.t=
u.0.s.4.pt=0
u.0.s.5.ut=3
u.0.s.5.k=rs
u.0.s.5.v=1141.0000
u.0.s.5.t=
u.0.s.5.pt="0,2"
u.[]=1
`))
	us := parseObject(d)

	for _, u := range us {
		fmt.Printf("OwnerID: %d, OwnerType: %d\n", u.OwnerID, u.OwnerType)
		for _, s := range u.Stats {
			fmt.Println(s)
		}
		fmt.Println()
	}

	// Output:
	// OwnerID: 729, OwnerType: 1
	// {cpc [0]  3 10000 4}
	// {ct [0]  3 17878120 4}
	// {dfv5 [0]  3 80000 4}
	// {dfw3201 [0]  3 10000 4}
	// {m_ct12 [0]  3 1089059 4}
	// {rs [0 2]  3 11410000 4}
	//
}
