// Code generated by "stringer -type SpecialGlucose"; DO NOT EDIT.

package dexcom

import "strconv"

const (
	_SpecialGlucose_name_0 = "SensorNotActiveMinimalDeviationNoAntenna"
	_SpecialGlucose_name_1 = "SensorNotCalibratedCountDeviation"
	_SpecialGlucose_name_2 = "AbsoluteDeviationPowerDeviation"
	_SpecialGlucose_name_3 = "BadRF"
)

var (
	_SpecialGlucose_index_0 = [...]uint8{0, 15, 31, 40}
	_SpecialGlucose_index_1 = [...]uint8{0, 19, 33}
	_SpecialGlucose_index_2 = [...]uint8{0, 17, 31}
	_SpecialGlucose_index_3 = [...]uint8{0, 5}
)

func (i SpecialGlucose) String() string {
	switch {
	case 1 <= i && i <= 3:
		i -= 1
		return _SpecialGlucose_name_0[_SpecialGlucose_index_0[i]:_SpecialGlucose_index_0[i+1]]
	case 5 <= i && i <= 6:
		i -= 5
		return _SpecialGlucose_name_1[_SpecialGlucose_index_1[i]:_SpecialGlucose_index_1[i+1]]
	case 9 <= i && i <= 10:
		i -= 9
		return _SpecialGlucose_name_2[_SpecialGlucose_index_2[i]:_SpecialGlucose_index_2[i+1]]
	case i == 12:
		return _SpecialGlucose_name_3
	default:
		return "SpecialGlucose(" + strconv.FormatInt(int64(i), 10) + ")"
	}
}
