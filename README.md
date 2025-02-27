# fastshout-geomag
fastshout-geomag is a Go-based implementation of the NOAA World Magnetic Model.
The World Magnetic Model home is at https://www.ngdc.noaa.gov/geomag/WMM/DoDWMM.shtml.
The coefficients for 2020-2024 can be downloaded at https://www.ngdc.noaa.gov/geomag/WMM/data/WMM2020/WMM2020COF.zip

## Commands
fastshout-geomag provides two command line programs, modeled after the command line programs in the official NOAA software.

`wmm_point` calculates magnetic field values for a single location and time.
The `wmm_grid` function (coming soon) will calculate magnetic field values for a grid of locations and/or times.

## Packages
This library provides two packages: `egm96` and `wmm`. `egm96` represents the 1996 Earth Gravitational Model (EGM96) and `wmm` represents the 2020 World Magnetic Model (WMM). These packages offer capabilities of representing geopotential model of the Earth and magnetic field produced by the Earth's core respectively.

## Validation
All library code is fully tested, covering all test values provided with the official NOAA WMM, along with the detailed example in the WMM technical paper. Please submit any issues on GitHub if you notice anomalies.

## Updating Coefficients
Coefficients are updated via go-bindata conversion of coefficient files; the `bindata.go` file should be edited to reflect updated coefficient data.

The WMM source code originates from the public domain and is not protected by copyright: https://www.ngdc.noaa.gov/geomag/WMM/license.shtml.