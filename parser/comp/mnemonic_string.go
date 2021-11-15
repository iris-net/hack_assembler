// Code generated by "stringer -type=Mnemonic"; DO NOT EDIT.

package comp

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[Zero-0]
	_ = x[One-1]
	_ = x[MinusOne-2]
	_ = x[D-3]
	_ = x[A-4]
	_ = x[NotD-5]
	_ = x[NotA-6]
	_ = x[MinusD-7]
	_ = x[MinusA-8]
	_ = x[DPlusOne-9]
	_ = x[APlusOne-10]
	_ = x[DMinusOne-11]
	_ = x[AMinusOne-12]
	_ = x[DPlusA-13]
	_ = x[DMinusA-14]
	_ = x[AMinusD-15]
	_ = x[DAndA-16]
	_ = x[DOrA-17]
	_ = x[M-18]
	_ = x[NotM-19]
	_ = x[MinusM-20]
	_ = x[MPlusOne-21]
	_ = x[MMinusOne-22]
	_ = x[DPlusM-23]
	_ = x[DMinusM-24]
	_ = x[MMinusD-25]
	_ = x[DAndM-26]
	_ = x[DOrM-27]
	_ = x[Unknown-28]
}

const _Mnemonic_name = "ZeroOneMinusOneDANotDNotAMinusDMinusADPlusOneAPlusOneDMinusOneAMinusOneDPlusADMinusAAMinusDDAndADOrAMNotMMinusMMPlusOneMMinusOneDPlusMDMinusMMMinusDDAndMDOrMUnknown"

var _Mnemonic_index = [...]uint8{0, 4, 7, 15, 16, 17, 21, 25, 31, 37, 45, 53, 62, 71, 77, 84, 91, 96, 100, 101, 105, 111, 119, 128, 134, 141, 148, 153, 157, 164}

func (i Mnemonic) String() string {
	if i < 0 || i >= Mnemonic(len(_Mnemonic_index)-1) {
		return "Mnemonic(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _Mnemonic_name[_Mnemonic_index[i]:_Mnemonic_index[i+1]]
}
