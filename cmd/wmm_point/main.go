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
// longitude (positive for northern latitudes an