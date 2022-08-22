package egm96

// Factors for converting between radians and degrees
// and between meters and feet.
const (
	Deg = 1/57.29577951308232 // number of radians per degree
	Ft  = 0.3048              // number of meters per foot
)

// DMSToDegrees converts integral degrees d, minutes m and seconds s (al