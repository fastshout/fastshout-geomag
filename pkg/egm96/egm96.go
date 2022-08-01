
// Package egm96 provides a representation of the 1996 Earth Gravitational Model (EGM96),
// a geopotential model of the Earth.
//
// EGM96 is the geoid reference model component of the World Geodetic System (WGS84).
// It consists of n=m=360 spherical harmonic coefficients as published by the
// National Geospatial-Intelligence Agency (NGA).  The NGA also publishes a raster grid
// of the calculated heights which can be interpolated to approximate the geoid height
// at any location.
//
// In effect, this model provides the height of sea level above the WGS84 reference ellipsoid.
// It is used, for example, in GPS navigation to provide the height above sea level.
//
// This package is based on the NGA-provided 15'x15' resolution grid encoding
// the heights of the geopotential surface at each lat/long, and interpolates between grid
// points using a bilinear interpolation.
package egm96

import (
	"bufio"
	"bytes"
	"fmt"
	"math"
	"strconv"
	"strings"
)

// Constants defining the WGS84 reference ellipsoid
const (
	A  = 6378137         // Equatorial radius of WGS84 reference ellipsoid in meters
	F  = 1/298.257223563 // Flattening of WGS84 reference ellipsoid
	E2 = F*(2-F)         // Eccentricity squared of WGS84 reference ellipsoid
)

// Location is a type that represents a position in space as represented
// by a latitude, a longitude and a height.
type Location struct {
	latitude  float64