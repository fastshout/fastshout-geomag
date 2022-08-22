package egm96

// Factors for converting between radians and degrees
// and between meters and feet.
const (
	Deg = 1/57.29577951308232 // number of radians per degree
	Ft  = 0.3048              // number of meters per foot
)

// DMSToDegrees converts integral degrees d, minutes m and seconds s (all of type float64)
// to a float-valued degrees amount.
//
// If d<0 then must pass m>0 and s>0;
// if d==0 and m<0 then must pass s>0.
func DMSToDegrees(d, m, s float64) (dd float64) {
	var sgn float64 = 1
	if d<0 {
		sgn = -