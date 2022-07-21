# EGM96

This package egm96 provides a Go representation of the EGM96 geopotential model of the Earth.
It calculates the geoid height of the 1996 Earth Gravitational Model (EGM96) for a given latitude and longitude.

## The EGM96 Model
The EGM96 model is a component of the 1984 World Geodetic System (WGS84).
The EGM96 homepage is at https://cddis.nasa.gov/926/egm96/egm96.html.

WGS84 defines a datum surface which is an ellipsoid whose center coincides with the Earth's center of mass.
EGM96 defines a "geoid," a gravitational equipotential surface, relative to this datum surface.
As an equipotential surface, the geoid