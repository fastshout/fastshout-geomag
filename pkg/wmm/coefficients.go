
package wmm

import (
	"bufio"
	"bytes"
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
	"time"
)

const (
	MaxLegendreOrder = 12
)

var (
	Epoch     DecimalYear // The Epoch of the loaded coefficients file, e.g. 2015.0
	COFName   string      // The filename of the loaded COF file
	ValidDate time.Time   // The beginning valid date of the loaded COF file
	cGnm      [][]float64
	cHnm      [][]float64
	cDGnm     [][]float64
	cDHnm     [][]float64
)

// GetWMMCoefficients calculates the spherical harmonic coefficients G(n,m), H(n,m)
// and their rates of change dG(n,m), dH(n,m) at the input time.
//
// If the request n,m are invalid or the requested time is outside of the range
// of validity of the loaded coefficients file, it will return an error.
func GetWMMCoefficients(n, m int, t time.Time) (gnm, hnm, dgnm, dhnm float64, err error) {
	if Epoch==0 {
		_ = LoadWMMCOF("")
	}
	if n<0 || n>MaxLegendreOrder || m<0 || m>MaxLegendreOrder {
		return 0, 0, 0, 0, fmt.Errorf("n, m = (%d,%d) must be between 0 and %d",
			n, m, MaxLegendreOrder)
	}
	if m>n {
		return 0, 0, 0, 0, fmt.Errorf("m=%d must be less than n=%d", m, n)
	}
	if t.Sub(ValidDate) < 0 || TimeToDecimalYears(t)>Epoch+5 {
		err = fmt.Errorf("requested date %v is outside of validity period beginning %v of WMM.COF file",
				t, ValidDate)
	}
	dt := float64(TimeToDecimalYears(t)- Epoch)
	gnm = cGnm[n][m] + dt*cDGnm[n][m]
	hnm = cHnm[n][m] + dt*cDHnm[n][m]
	dgnm = cDGnm[n][m]
	dhnm = cDHnm[n][m]
	return gnm, hnm, dgnm, dhnm, err
}

// LoadWMMCOF loads the specified coefficients file.
//
// It populates the internal coefficient values representing G(n,m), H(n,m), DG(n,m), DH(n,m),
// Epoch, COFName, and ValidDate.
// If the passed filename is "", it loads the default (current) coefficients file.
//
// The default coefficients file is currently WMM2020.COF, valid from
// 12/10/2019 until 12/31/2024.
func LoadWMMCOF(fn string) (err error) {
	var (
		data []byte
		epoch float64
		n, m  int
	)

	if fn=="" {
		data, err = getAsset("WMM.COF")
	} else {
		data, err = ioutil.ReadFile(fn)
	}
	if err != nil {
		return err
	}

	scanner := bufio.NewScanner(bytes.NewReader(data))
	// Read and parse header
	if !scanner.Scan() {
		return fmt.Errorf("Could not read header line in WMM coefficient file %s", fn)
	}
	dat := strings.Fields(scanner.Text())
	if epoch, err = strconv.ParseFloat(dat[0], 64); err != nil {
		return fmt.Errorf("bad header epoch date in WMM coefficient file %s", fn)
	}
	Epoch = DecimalYear(epoch)
	COFName = dat[1]
	if ValidDate, err = time.Parse("01/02/2006", dat[2]); err != nil {
		return fmt.Errorf("bad header valid date in WMM coefficient file %s", fn)
	}

	cGnm = make([][]float64, MaxLegendreOrder+1)
	cGnm[0] = []float64{0}
	cHnm = make([][]float64, MaxLegendreOrder+1)
	cHnm[0] = []float64{0}
	cDGnm = make([][]float64, MaxLegendreOrder+1)
	cDGnm[0] = []float64{0}
	cDHnm = make([][]float64, MaxLegendreOrder+1)
	cDHnm[0] = []float64{0}

	// Read and parse testdata
	curN := 0
	for scanner.Scan() {
		s := strings.Fields(scanner.Text())
		if len(s)<6 {
			continue
		}
		if n, err = strconv.Atoi(s[0]); err!=nil {
			return fmt.Errorf("bad n value in WMM coefficient file %s", fn)
		}
		if m, err = strconv.Atoi(s[1]); err!=nil {
			return fmt.Errorf("bad m value in WMM coefficient file %s", fn)
		}
		if n>curN {
			cGnm[n] = make([]float64, n+1)
			cHnm[n] = make([]float64, n+1)
			cDGnm[n] = make([]float64, n+1)
			cDHnm[n] = make([]float64, n+1)
			curN = n
		}
		if cGnm[n][m], err = strconv.ParseFloat(s[2], 64); err != nil {
			return fmt.Errorf("bad Gnm value in WMM coefficient file %s", fn)
		}
		if cHnm[n][m], err = strconv.ParseFloat(s[3], 64); err != nil {
			return fmt.Errorf("bad Hnm value in WMM coefficient file %s", fn)
		}
		if cDGnm[n][m], err = strconv.ParseFloat(s[4], 64); err != nil {
			return fmt.Errorf("bad DGnm value in WMM coefficient file %s", fn)
		}
		if cDHnm[n][m], err = strconv.ParseFloat(s[5], 64); err != nil {
			return fmt.Errorf("bad DHnm value in WMM coefficient file %s", fn)
		}
	}

	if err := scanner.Err(); err != nil {
		return err
	}
	return nil
}