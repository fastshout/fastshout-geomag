
// Package wmm provides a representation of the World Magnetic Model (WMM),
// a mathematical model of the magnetic field produced by the Earth's core and
// its variation over time.
//
// WMM is the magnetic model component of the World Geodetic System (WGS84).
// It consists of n=m=12 spherical harmonic coefficients as published by the
// National Geospatial-Intelligence Agency (NGA).
//
// This model evaluates all magnetic field components and their rates of change
// for any location on the Earth's surface.  These field components include the
// X, Y, and Z values and the declination D and inclination I.
// The Declination is used, for example, in correcting a Magnetic Heading to a
// True Heading.
package wmm

import (
	"math"
	"time"

	"github.com/westphae/geomag/pkg/egm96"
	"github.com/westphae/geomag/pkg/polynomial"
)

const (
	AGeo  = 6371200 // Geomagnetic Reference Radius
	errX  = 131     // WMM global average X error, nT
	errY  = 94      // WMM global average Y error, nT
	errZ  = 157     // WMM global average Z error, nT
	errH  = 128     // WMM global average H error, nT
	errF  = 148     // WMM global average F error, nT
	errI  = 0.21    // WMM global average I error, º
	errDA = 0.26    // WMM rough global average D error away from poles, º
	errDB = 5625    // WMM average H uncertainty scale near the poles, nT
)

// MagneticField represents a geomagnetic field and its rate of change.
type MagneticField struct {
	l          egm96.Location
	x, y, z    float64
	dx, dy, dz float64
}

// Ellipsoidal returns the magnetic field in ellipsoidal coordinate axes.
//
// The Ellipsoidal axes are the most commonly desired axes, in which the
// horizontal directions are parallel to the WGS84 ellipsoid.
//
// Field strengths are in nT and field strength changes in nT/Year.
func (m MagneticField) Ellipsoidal() (x, y, z, dx, dy, dz float64) {
	latS, _, _ := m.l.Spherical()
	latG, _, _ := m.l.Geodetic()
	cosDPhi := math.Cos(latS - latG)
	sinDPhi := math.Sin(latS - latG)
	x = m.x*cosDPhi - m.z*sinDPhi
	y = m.y
	z = m.x*sinDPhi + m.z*cosDPhi
	dx = m.dx*cosDPhi - m.dz*sinDPhi
	dy = m.dy
	dz = m.dx*sinDPhi + m.dz*cosDPhi
	return x, y, z, dx, dy, dz
}

// Spherical returns the magnetic field in spherical coordinate axes.
//
// The spherical axes are centered on the Earth's center of mass.
// These axes won't typically be used for navigation on or near the
// Earth's surface, but might be used in space.
//
// Field strengths are in nT and field strength changes in nT/Year.
func (m MagneticField) Spherical() (x, y, z, dx, dy, dz float64) {
	return m.x, m.y, m.z, m.dx, m.dy, m.dz
}

// H returns the strength of the magnetic field in the horizontal
// direction, i.e. the component parallel to the WGS84 ellipsoid.
//
// The return value is in nT.
func (m MagneticField) H() (h float64) {
	x, y, _, _, _, _ := m.Ellipsoidal()
	return math.Sqrt(x*x + y*y)
}

// F returns the total strength of the magnetic field.