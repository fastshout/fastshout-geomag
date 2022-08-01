
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
	longitude float64
	height    float64
}

// NewLocationGeodetic returns a Location given an input latitude, longitude,
// and height specified in the Geodetic system.
//
// The Geodetic coordinate system is the usual latitude, longitude, and
// height above the WGS84 Reference Ellipsoid, i.e. as typically measured by a GPS receiver.
//
// Latitude and longitude are specified in decimal degrees and height in meters.
//
// Geodetic coordinates are the un-primed variables φ,λ,h in the WMM paper.
func NewLocationGeodetic(latitude, longitude, height float64) (loc Location) {
	return Location{
		latitude: latitude*Deg,
		longitude: longitude*Deg,
		height: height,
	}
}

// NewLocationMSL returns a Location given an input latitude, longitude, and height
// above mean sea level.
//
// The latitude and longitude are as specified in the Geodetic Coordinate System,
// and the height is the height above mean sea level, NOT above the WGS84 Reference Ellipsoid.
//
// Latitude and longitude are specified in decimal degrees and height in meters.
func NewLocationMSL(latitude, longitude, height float64) (loc Location, err error) {
	if len(egm96Grid)==0 {
		loadEGM96Grid()
	}

	nLng := int((longitude-egm96X0)/egm96DX) // Grid x just below desired x
	nLat := int((latitude-egm96Y0)/egm96DY) // Grid y just below desired y

	if nLng < 0 || nLng > egm96XN {
		return Location{}, fmt.Errorf("requested longitude %4.2f lies outside of EGM96 longitude range %4.1f to %4.1f",
			longitude, egm96X0, egm96X1)
	}
	if nLat < 0 || nLat > egm96YN {
		return Location{}, fmt.Errorf("requested latitude %4.2f lies outside of EGM96 latitude range %4.1f to %4.1f",
			latitude, egm96Y0, egm96Y1)
	}

	x := (longitude-egm96X0)/egm96DX-float64(nLng)
	y := (latitude-egm96Y0)/egm96DY-float64(nLat)
	h00 := egm96Grid[nLat*egm96XN+nLng]
	h10 := egm96Grid[nLat*egm96XN+nLng+1]
	h01 := egm96Grid[(nLat+1)*egm96XN+nLng]
	h11 := egm96Grid[(nLat+1)*egm96XN+nLng+1]


	return Location{
		latitude: latitude*Deg,
		longitude: longitude*Deg,
		height: height + ((1-x)*(1-y)*h00 + x*(1-y)*h10 + (1-x)*y*h01 + x*y*h11),
	}, nil
}

// Equals returns whether the latitude, longitude and height of the input location
// are equal to those of the caller.
func (l Location) Equals(ll Location) bool {
	return l.latitude==ll.latitude && l.longitude==ll.longitude && l.height==ll.height
}

// Geodetic returns the location's lat (latitude), lng (longitude), and h (height).
// lat and lng are in radians and r is in meters.
// Geodetic coordinates are the variables φ,λ,h in the WMM paper.
func (l Location) Geodetic() (phi, lambda, r float64) {
	return l.latitude, l.longitude, l.height
}

// Spherical returns the location's phi (φ', corresponding to latitude),
// lambda (λ, equal to geodetic longitude), and r (r, distance from center of
// WGS sphere).  phi and lambda are in radians and r is in meters.
// Spherical coordinates are the variables φ',λ,r in the WMM paper.
func (l Location) Spherical() (phi, lambda, r float64) {
	sinPhi := math.Sin(l.latitude)
	cosPhi := math.Cos(l.latitude)
	h := l.height
	rc := A/math.Sqrt(1-E2*sinPhi*sinPhi)
	p := (rc+h)*cosPhi
	z := (rc*(1-E2)+h)*sinPhi
	r = math.Sqrt(p*p+z*z)
	return math.Asin(z/r), l.longitude, r
}

// HeightAboveMSL calculates the height of the EGM96 geoid at the input Location,
// which corresponds to the height of MSL relative to the WGS84 reference ellipsoid.
// It then subtracts this height from the total height above the WGS84 reference
// ellipsoid at the input Location, giving the the height above MSL.
func (l Location) HeightAboveMSL() (h float64, err error) {
	if len(egm96Grid)==0 {
		loadEGM96Grid()
	}

	lng := l.longitude/Deg
	lat := l.latitude/Deg
	nLng := int((lng-egm96X0)/egm96DX) // Grid x just below desired x
	nLat := int((lat-egm96Y0)/egm96DY) // Grid y just below desired y

	if nLng < 0 || nLng > egm96XN {
		return 0, fmt.Errorf("requested longitude %4.2f lies outside of EGM96 longitude range %4.1f to %4.1f",
			lng, egm96X0, egm96X1)
	}
	if nLat < 0 || nLat > egm96YN {
		return 0, fmt.Errorf("requested latitude %4.2f lies outside of EGM96 latitude range %4.1f to %4.1f",
			lat, egm96Y0, egm96Y1)
	}

	x := (lng-egm96X0)/egm96DX-float64(nLng)
	y := (lat-egm96Y0)/egm96DY-float64(nLat)
	h00 := egm96Grid[nLat*egm96XN+nLng]
	h10 := egm96Grid[nLat*egm96XN+nLng+1]
	h01 := egm96Grid[(nLat+1)*egm96XN+nLng]