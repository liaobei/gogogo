package uc

import "testing"

type ucTest struct {
	in, out string
}

var ucTests = []ucTest{
	ucTest{"abc", "ABC"},
	ucTest{"cvo-az", "CVO-AZ"},
	ucTest{"Antwrep", "ANTWREP"},
}

func testUC(t *testing.T) {
	for _, ut := range ucTests {
		uc := UpperCase(ut.in)
		if uc != ut.out {
			t.Error("UpperCase(%s) = %s,ms=yst be %s", ut.in, uc, ut.out)
		}
	}
}
