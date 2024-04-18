# fastshout-geomag
fastshout-geomag is a Go-based implementation of the NOAA World Magnetic Model.
The World Magnetic Model home is at https://www.ngdc.noaa.gov/geomag/WMM/DoDWMM.shtml.
The coefficients for 2020-2024 can be downloaded at https://www.ngdc.noaa.gov/geomag/WMM/data/WMM2020/WMM2020COF.zip

## Commands
fastshout-geomag provides two command line programs, modeled after the command line programs in the official NOAA software.

`wmm_point` calculates magnetic field values for a single location and time.
The `wmm_grid` function (coming soon) will calculate magnetic field values for a grid of locations and/or times.

## Packages
This library provides two packages: `egm96` and `wmm`. `egm96` represents the 1996 Earth Gravitational Model (EGM96) and `wmm` represents the 2020 World Magnetic Model (WMM). T