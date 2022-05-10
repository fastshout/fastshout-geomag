// wmm_point estimates the strength and direction of Earth's main Magnetic field for a given point/area.
//
// Usage is
//  wmm_point --cof_file=WMM2020.COF --spherical [latitude] [longitude] [altitude] [date]
//
// The World Magnetic Model (WMM) for 2020
// is a model of Earth's main Magnetic field.  The WMM
// is recomputed every five (5) years, in years divisible by
// five (i.e. 2010, 2015, 2020).
//
// Information on the model is available at https://www.ngdc.noaa.gov/geomag/WMM/DoDWMM.shtml
//
// Input required is the location in geodetic latitude and
// longitude (positive for northern latitudes and eastern
// longitudes), geodetic altitude in meters, and the date of
// interest in years.
//
// The program computes the estimated Magnetic Declination
// (Decl) which is sometimes called MagneticVAR, Inclination (Incl), Total
// Intensity (F or TI), Horizontal Intensity (H or HI), Vertical
// Intensity (Z), and Grid Variation (GV). Declination and Grid
// Variation are measured in units of degrees and are considered
// positive when east or north.  Inclination is measured in units
// of degrees and is considered positive when pointing down (into
// the Earth).  The WMM is referenced to the WGS-84 ellipsoid and
// is valid for 5 years after the base epoch. Uncertainties for the
// WMM are one standard deviation uncertainties averaged over the globe.
// We represent the uncertainty as constant values in Incl, F, H, X,
// Y, and Z. Uncertainty in Declination varies depending on the strength
// of the horizontal field.  For more information see the WMM Technical
// Report.
//
// It is very important to note that a  degree and  order 12 model,
// such as WMM, describes only the long  wavelength spatial Magnetic
// fluctuations due to  Earth's core.  Not included in the WMM series
// models are intermediate and short wavelength spatial fluctuations
// that originate in Earth's mantle and crust. Consequently, isolated
// angular errors at various  positions on the surface (primarily over
// land, along continental margins and  over oceanic sea-mounts, ridges and
// trenches) of several degrees may be expected.  Also not included in
// the model are temporal fluctuations of magnetospheric and ionospheric
// origin. On the days during and immediately following Magnetic storms,
// temporal fluctuations can cause substantial deviations of the Geomagnetic
// field  from model  values.  If the required  declination accuracy  is
// more stringent than the WMM  series of models provide, the user is
// advised to request special (regional or local) surveys be performed
// and models prepared. The World Magnetic Model is a joint product of
// the United States’ National Geospatial-Intelligence Agency (NGA) and
// the United Kingdom’s Defence Geographic Centre (DGC). The WMM was
// developed jointly by the National Centers for Environmental Information (NCEI, Boulder
// CO, USA) and the British Geological Survey (BGS, Edinburgh, Scotland).
//
// Sample output:
//  Results For
//  
//  Latitude:       30.00N
//  Longitude:      88.51W
//  Altitude:        0.010 kilometers above mean sea level
//  Date:           2019.5
//  
//         Main Field             Secular Change
//         F    =  46944.3 nT ± 152.0 nT  -118.8 nT/yr
//         H    =  24074.6 nT ± 133.0 nT    -6.8 nT/yr
//         X    =  24060.2 nT ± 138.0 nT    -8.0 nT/yr
//         Y    =   -831.0 nT ±  89.0 nT   -36.3 nT/yr
//         Z    =  40301.2 nT ± 165.0 nT  -134.3 nT/yr
//         Decl =     -1º 59' ± 19'         -5.2'/yr
//         Incl =     59º  9' ± 13'         -4.6'/yr
//  
//         Grid Variation =  -1º 59'
package main

import (
	"bufio"
	"errors"
	