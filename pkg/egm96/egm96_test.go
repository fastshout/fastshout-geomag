package egm96

import (
	"fmt"
	"testing"
)

const eps = 1e-6

func testDiff(name string, actual, expected float64, eps float64, t *testing.T) {
	if actual - expected > -eps && actual - expected < eps {
		t.Logf("%s correct: expected %8.4f, got %8.4f", name, expected, actual)
		return
	}
	t.Errorf("%s incorrect: expected %8.4f, got %8.4f", name, expected, actual)
}

func TestEGM96GridLookup(t *testing.T) {
	lats := []float64{38, -12.25, -84.75, 26, 0}
	lngs := []float64{270, 82.75, 180.5, 279.5, 0}
	hts  := []float64{-30.262, -67.347, -40.254, -26.621, 17.162}

	for i:=0; i<len(lats); i++ {
		p, _ := NewLocationGeodetic(lats[i],lngs[i],0).NearestEGM96GridPoint()

		testDiff("latitude", p.latitude/Deg, lats[i], eps, t)
		testDiff("longitude", p.longitude/Deg, lngs[i], eps, t)
		testDiff("height", p.height, hts[i], eps, t)
	}
}

func TestEGM96GridInterpolationAgainstKnown(t *testing.T) {
	lats := []float64{38, -12.25, 0, 38.628155, -14.621217, 46.874319, -23.617446, 38.625473, -0.466744}
	lngs := []float64{270, 82.75, 0, 269.779155, 305.021114, 102.448729, 133.874712, 359.999500, 0.002300}
	hts  := []float64{-30.262, -67.347, 17.162,