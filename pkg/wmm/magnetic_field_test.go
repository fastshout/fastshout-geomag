
package wmm

import (
	"bufio"
	"bytes"
	"io/ioutil"
	"strconv"
	"strings"
	"testing"

	"github.com/westphae/geomag/pkg/egm96"
)

const (
	epsM = 5e-5
	red = "\u001b[31m"
	green = "\u001b[32m"
	reset = "\u001b[0m"
)

func testDiff(name string, actual, expected float64, eps float64, t *testing.T) {
	if actual-expected >= -eps && actual-expected <= eps {
		t.Logf("%s%s correct: expected %6.4f, got %6.4f%s", green, name, expected, actual, reset)
		return
	}
	t.Errorf("%s%s incorrect: expected %6.4f, got %6.4f%s", red, name, expected, actual, reset)
}

func TestMagneticFieldFromPaperDetail(t *testing.T) {
	// Test values in paper are only for original version of WMM-2015
	_ = LoadWMMCOF("testdata/WMM2015v1.COF")
	tt := DecimalYear(2017.5)
	loc := egm96.NewLocationGeodetic(-80,240,100e3)

	lat, lng, hh := loc.Geodetic()
	testDiff("lambda", lng, 4.1887902048, epsM, t)