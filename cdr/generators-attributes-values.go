package cdr

import (
	"fmt"
	"math/rand"
	"strconv"
	"strings"
	"time"

	"github.com/bxcodec/faker"
	"github.com/bxcodec/faker/support/slice"
)

type CdrValues struct {
	AcctStatusType int
	ServiceType    int
	ResponseCode   string
	Method         string
	EventTimestamp time.Time
	FromTag        string
	ToTag          string
	AcctSessionId  string
	MsDuration     int
	SetupTime      int
	CallerId       string
	CalleeId       string
	DstNumber      string
}

// random ResponseCode in a collection
func ResponseCode() string {
	codes := []string{
		"200",
		"480",
		"503",
	}
	return codes[rand.Int()%len(codes)]
}

// generate brazilian phone number on default E164
func PhoneNumberBrazil() string {
	out := ""
	box_numbers := []string{
		"11",
		"21",
		"31",
		"51",
		"66",
	}
	ints, _ := faker.RandomInt(1, 8)

	for _, v := range slice.IntToString(ints) {
		out += string(v)
	}
	return fmt.Sprintf("55%s9%s", box_numbers[rand.Int()%len(box_numbers)], strings.Join(slice.IntToString(ints), ""))
}

// generate ms_duration, setuptime based on sip_code
func CdrTimers(c int) (int, int) {
	st := rand.Intn(30)

	if c != 200 {
		return 0, st
	}

	rand.Seed(time.Now().UnixNano())
	min := 100000
	max := 900000
	ms := min + rand.Intn(max-min+1)
	return ms, st
}

// random Addresses IPV4 in a collection
func Addresses() (string, string) {
	s := []string{
		"200.200.200.200",
		"250.250.250.250",
	}
	d := []string{
		"100.100.100.100",
		"130.130.130.130",
		"150.150.150.150",
	}
	return s[rand.Int()%len(s)], d[rand.Int()%len(d)]
}

// create and set all struct CdrValues with generated data
func FillCdr() *CdrValues {
	src_ip, dst_ip := Addresses()
	r := ResponseCode()
	ri, _ := strconv.Atoi(r)
	ms, st := CdrTimers(ri)
	dr := PhoneNumberBrazil()
	de := PhoneNumberBrazil()
	return &CdrValues{
		AcctStatusType: 4,
		ServiceType:    15,
		ResponseCode:   r,
		Method:         "INVITE",
		EventTimestamp: time.Now(),
		FromTag:        faker.GetIdentifier().Digit()[:16],
		ToTag:          faker.GetIdentifier().Digit()[:8],
		AcctSessionId:  faker.GetIdentifier().Digit()[:12] + "@" + src_ip,
		MsDuration:     ms,
		SetupTime:      st,
		CallerId:       "sip:" + dr + "@" + src_ip + ":5077",
		CalleeId:       "sip:" + de + "@" + dst_ip + ":5060",
		DstNumber:      de,
	}
}
