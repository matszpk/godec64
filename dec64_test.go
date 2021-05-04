/*
 * dec64_test.go - main fixed decimal int128 routines
 *
 * godec64 - go dec64 (for 64-bit decimal fixed point) library
 * Copyright (C) 2020  Mateusz Szpakowski
 *
 * This library is free software; you can redistribute it and/or
 * modify it under the terms of the GNU Lesser General Public
 * License as published by the Free Software Foundation; either
 * version 2.1 of the License, or (at your option) any later version.
 *
 * This library is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the GNU
 * Lesser General Public License for more details.
 *
 * You should have received a copy of the GNU Lesser General Public
 * License along with this library; if not, write to the Free Software
 * Foundation, Inc., 51 Franklin Street, Fifth Floor, Boston, MA  02110-1301  USA
 */

// Package to operate on 64-bit decimal fixed point
package godec64
 
import (
    "strconv"
    "testing"
)

type UDec64MulTC struct {
    a, b UDec64
    precision uint
    rounding bool
    expected UDec64
}

func TestUDec64Mul(t *testing.T) {
    testCases := []UDec64MulTC {
        UDec64MulTC{ 349884939232, 495983892892, 8, false, 1735372942245682 },
        UDec64MulTC{ 349884939234, 495983892793, 8, false, 1735372941909215 },
        UDec64MulTC{ 349884939232, 495983892892, 8, true, 1735372942245682 },
        UDec64MulTC{ 349884939234, 495983892793, 8, true, 1735372941909216 },
        UDec64MulTC{ 5489289235455, 443992839213, 10, false, 243720511291102 },
        UDec64MulTC{ 5489289235455, 443992839213, 10, true, 243720511291102 },
        UDec64MulTC{ 5489289235458, 443992839213, 10, false, 243720511291235 },
        UDec64MulTC{ 5489289235458, 443992839213, 10, true, 243720511291236 },
    }
    for i, tc := range testCases {
        a, b := tc.a, tc.b
        result := tc.a.Mul(tc.b, tc.precision, tc.rounding)
        if tc.expected!=result {
            t.Errorf("Result mismatch: %d: mul(%v,%v,%v,%v)->%v!=%v",
                     i, tc.a, tc.b, tc.precision, tc.rounding, tc.expected, result)
        }
        if tc.a!=a || tc.b!=b {
            t.Errorf("Argument has been modified: %d: %v,%v!=%v,%v",
                     i, a, b, tc.a, tc.b)
        }
    }
}

type UDec64DivTC struct {
    a, b UDec64
    precision uint
    expected UDec64
}

func TestUDec64Div(t *testing.T) {
    testCases := []UDec64DivTC {
        UDec64DivTC { 243720511291235, 443992839213, 10, 5489289235457 },
    }
    for i, tc := range testCases {
        a, b := tc.a, tc.b
        result := tc.a.Div(tc.b, tc.precision)
        if tc.expected!=result {
            t.Errorf("Result mismatch: %d: div(%v,%v,%v)->%v!=%v",
                     i, tc.a, tc.b, tc.precision, tc.expected, result)
        }
        if tc.a!=a || tc.b!=b {
            t.Errorf("Argument has been modified: %d: %v,%v!=%v,%v",
                     i, a, b, tc.a, tc.b)
        }
    }
}

type UDec64FmtTC struct {
    a UDec64
    precision uint
    trimZeroes bool
    expected string
}

type UDec64Fmt2TC struct {
    a UDec64
    precision uint
    dispPrecision uint
    trimZeroes bool
    expected string
}

func TestUDec64Format(t *testing.T) {
    testCases := []UDec64FmtTC {
        UDec64FmtTC{ 425143693331510191, 15, false, "425.143693331510191" },
        UDec64FmtTC{ 425143693331510191, 10, false, "42514369.3331510191" },
        UDec64FmtTC{ 425143693331510200, 15, false, "425.143693331510200" },
        UDec64FmtTC{ 425143693331510200, 15, true, "425.1436933315102" },
        UDec64FmtTC{ 425143693331510000, 15, false, "425.143693331510000" },
        UDec64FmtTC{ 425143693331510000, 15, true, "425.14369333151" },
        UDec64FmtTC{ 1984593924556, 15, false, "0.001984593924556" },
        UDec64FmtTC{ 1984593924560, 15, false, "0.001984593924560" },
        UDec64FmtTC{ 1984593924560, 15, true, "0.00198459392456" },
        UDec64FmtTC{ 1984593924000, 15, false, "0.001984593924000" },
        UDec64FmtTC{ 1984593924000, 15, true, "0.001984593924" },
        UDec64FmtTC{ 0, 15, true, "0.0" },
        UDec64FmtTC{ 1, 15, false, "0.000000000000001" },
        UDec64FmtTC{ 3211984593924556, 15, false, "3.211984593924556" },
        UDec64FmtTC{ 33000000000000000, 15, false, "33.000000000000000" },
        UDec64FmtTC{ 33000000000000000, 15, true, "33.0" },
        UDec64FmtTC{ 33400000000000000, 15, true, "33.4" },
        UDec64FmtTC{ 33000400000000000, 15, true, "33.0004" },
        // zero digits after comma
        UDec64FmtTC{ 425143693331510191, 0, false, "425143693331510191" },
    }
    for i, tc := range testCases {
        a := tc.a
        result := tc.a.Format(tc.precision, tc.trimZeroes)
        if tc.expected!=result {
            t.Errorf("Result mismatch: %d: fmt(%v)->%v!=%v",
                     i, tc.a, tc.expected, result)
        }
        if tc.a!=a {
            t.Errorf("Argument has been modified: %d: %v!=%v", i, a, tc.a)
        }
        resultBytes := tc.a.FormatBytes(tc.precision, tc.trimZeroes)
        if tc.expected!=string(resultBytes) {
            t.Errorf("Result mismatch: %d: fmtBytes(%v)->%v!=%v",
                     i, tc.a, tc.expected, string(resultBytes))
        }
        if tc.a!=a {
            t.Errorf("Argument has been modified: %d: %v!=%v", i, a, tc.a)
        }
    }
    
    testCases2 := []UDec64Fmt2TC {
        UDec64Fmt2TC{ 425143693331510191, 15, 17, false, "425.14369333151019100" },
        UDec64Fmt2TC{ 425143693331510000, 15, 17, false, "425.14369333151000000" },
        UDec64Fmt2TC{ 425143693331510191, 15, 12, false, "425.143693331510" },
        UDec64Fmt2TC{ 425143693331510191, 15, 10, false, "425.1436933315" },
        UDec64Fmt2TC{ 425143693331510191, 15, 17, true, "425.143693331510191" },
        UDec64Fmt2TC{ 425143693331510191, 15, 10, true, "425.1436933315" },
        UDec64Fmt2TC{ 425143693331510191, 15, 12, true, "425.14369333151" },
        UDec64Fmt2TC{ 425143693331510200, 15, 12, true, "425.14369333151" },
        UDec64Fmt2TC{ 425143693331510000, 15, 13, true, "425.14369333151" },
    }
    for i, tc := range testCases2 {
        a := tc.a
        result := tc.a.FormatNew(tc.precision, tc.dispPrecision, tc.trimZeroes)
        if tc.expected!=result {
            t.Errorf("Result mismatch: %d: fmt(%v)->%v!=%v",
                     i, tc.a, tc.expected, result)
        }
        if tc.a!=a {
            t.Errorf("Argument has been modified: %d: %v!=%v", i, a, tc.a)
        }
        resultBytes := tc.a.FormatNewBytes(tc.precision, tc.dispPrecision, tc.trimZeroes)
        if tc.expected!=string(resultBytes) {
            t.Errorf("Result mismatch: %d: fmtBytes(%v)->%v!=%v",
                     i, tc.a, tc.expected, string(resultBytes))
        }
        if tc.a!=a {
            t.Errorf("Argument has been modified: %d: %v!=%v", i, a, tc.a)
        }
    }
}

type UDec64ParseTC struct {
    str string
    precision uint
    rounding bool
    expected UDec64
    expError error
}

func TestUDec64Parse(t *testing.T) {
    testCases := []UDec64ParseTC {
        UDec64ParseTC{ "425.143693331510191", 15, false, 425143693331510191, nil },
        UDec64ParseTC{ "425.1436933315101915", 15, false, 425143693331510191, nil },
        UDec64ParseTC{ "425.143693331510191999", 15, false, 425143693331510191, nil },
        UDec64ParseTC{ "425.1436933315101915", 15, true, 425143693331510192, nil },
        UDec64ParseTC{ "425.1436933315101", 15, false, 425143693331510100, nil },
        UDec64ParseTC{ "4592112", 10, false, 45921120000000000, nil },
        UDec64ParseTC{ "18446744073709551616", 11, false, 0, strconv.ErrRange },
        UDec64ParseTC{ "0.001984593924556", 15, false, 1984593924556, nil },
        UDec64ParseTC{ ".0019845939245565", 15, false, 1984593924556, nil },
        UDec64ParseTC{ ".0019845939245565", 15, true, 1984593924557, nil },
        UDec64ParseTC{ "0.001984593924560", 15, false, 1984593924560, nil },
        UDec64ParseTC{ ".001984593924560", 15, false, 1984593924560, nil },
        UDec64ParseTC{ "0.00198459392456", 15, false, 1984593924560, nil },
        UDec64ParseTC{ ".00198459392456", 15, false, 1984593924560, nil },
        UDec64ParseTC{ ".001984593924", 15, false, 1984593924000, nil },
        UDec64ParseTC{ "0.201984593924556", 15, false, 201984593924556, nil },
        UDec64ParseTC{ ".30198459392456", 15, false, 301984593924560, nil },
        UDec64ParseTC{ "0.0", 10, false, 0, nil },
        UDec64ParseTC{ "0", 10, false, 0, nil },
        UDec64ParseTC{ "0.", 10, false, 0, nil },
        UDec64ParseTC{ ".0", 10, false, 0, nil },
        UDec64ParseTC{ "425.143693331510191e0", 15, false, 425143693331510191, nil },
        UDec64ParseTC{ "42.5143693331510191e1", 15, false, 425143693331510191, nil },
        UDec64ParseTC{ "42.5143693331510191E1", 15, false, 425143693331510191, nil },
        UDec64ParseTC{ "42.5143693331510191ee1", 15, false, 0, strconv.ErrSyntax },
        UDec64ParseTC{ "4.25143693331510191e2", 15, false, 425143693331510191, nil },
        UDec64ParseTC{ "244.194251436933315e5", 10, false, 244194251436933315, nil },
        UDec64ParseTC{ "4251.43693331510191e-1", 15, false, 425143693331510191, nil },
        UDec64ParseTC{ "4251.43693331510191E-1", 15, false, 425143693331510191, nil },
        UDec64ParseTC{ "425143693.331510190e-6", 15, false, 425143693331510190, nil },
        UDec64ParseTC{ "0.01984593924556e-1", 15, false, 1984593924556, nil },
        UDec64ParseTC{ "2.14231E-1", 6, false, 214231, nil },
        UDec64ParseTC{ ".01984593924556e-1", 15, false, 1984593924556, nil },
        UDec64ParseTC{ "0.1984593924556e-2", 15, false, 1984593924556, nil },
        UDec64ParseTC{ "00.1984593924556e-2", 15, false, 1984593924556, nil },
        UDec64ParseTC{ ".1984593924556e-2", 15, false, 1984593924556, nil },
        UDec64ParseTC{ "1.984593924556e-3", 15, false, 1984593924556, nil },
        UDec64ParseTC{ "12e3", 15, false, 12000000000000000000, nil },
        UDec64ParseTC{ "12.e3", 15, false, 12000000000000000000, nil },
        UDec64ParseTC{ "12.77e3", 15, false, 12770000000000000000, nil },
        UDec64ParseTC{ "0.0e0", 10, false, 0, nil },
        UDec64ParseTC{ "0.0e1", 10, false, 0, nil },
        UDec64ParseTC{ ".0e1", 10, false, 0, nil },
        UDec64ParseTC{ "0.e1", 10, false, 0, nil },
        UDec64ParseTC{ "0.0e3", 10, false, 0, nil },
        UDec64ParseTC{ ".0e3", 10, false, 0, nil },
        UDec64ParseTC{ "0.e3", 10, false, 0, nil },
        UDec64ParseTC{ "0.0e-1", 10, false, 0, nil },
        UDec64ParseTC{ ".0e-1", 10, false, 0, nil },
        UDec64ParseTC{ "0.e-1", 10, false, 0, nil },
        UDec64ParseTC{ "0.0e-3", 10, false, 0, nil },
        UDec64ParseTC{ ".0e-3", 10, false, 0, nil },
        UDec64ParseTC{ "0.e-3", 10, false, 0, nil },
        UDec64ParseTC{ "12344", 0, false, 12344, nil },
        UDec64ParseTC{ "12344.", 0, false, 12344, nil },
        UDec64ParseTC{ "12344.0000", 0, false, 12344, nil },
        UDec64ParseTC{ "12344.7000", 0, false, 12344, nil },
        UDec64ParseTC{ "12344.7000", 0, true, 12345, nil },
    }
    for i, tc := range testCases {
        result, err := ParseUDec64(tc.str, tc.precision, tc.rounding)
        if tc.expected!=result || tc.expError!=err {
            t.Errorf("Result mismatch: %d: parse(%v)->%v,%v!=%v,%v",
                     i, tc.str, tc.expected, tc.expError, result, err)
        }
        result, err = ParseUDec64Bytes([]byte(tc.str), tc.precision, tc.rounding)
        if tc.expected!=result || tc.expError!=err {
            t.Errorf("Result mismatch: %d: parse(%v)->%v,%v!=%v,%v",
                     i, tc.str, tc.expected, tc.expError, result, err)
        }
    }
}

type UDec64ToFloat64TC struct {
    value UDec64
    precision uint
    expected float64
}

func TestUDec64ToFloat64(t *testing.T) {
    testCases := []UDec64ToFloat64TC{
        UDec64ToFloat64TC{ 0, 11, 0.0 },
        UDec64ToFloat64TC{ 1, 11, 1.0*1e-11 },
        UDec64ToFloat64TC{ 54930201, 11, 54930201.0*1e-11 },
        UDec64ToFloat64TC{ 85959028918918968, 0, 85959028918918968.0 },
        UDec64ToFloat64TC{ 85959028918918968, 11, 85959028918918968.0*1e-11 },
        UDec64ToFloat64TC{ 85959028918918968, 17, 0.8595902891891898 },
        UDec64ToFloat64TC{ 0xffffffffffffffff, 11, 18446744073709551615.0*1e-11 },
    }
    for i, tc := range testCases {
        result := tc.value.ToFloat64(tc.precision)
        if tc.expected!=result {
            t.Errorf("Result mismatch: %d: tofloat64(%v,%v)->%v!=%v",
                     i, tc.value, tc.precision, tc.expected, result)
        }
    }
}

type Float64ToUDec64TC struct {
    value float64
    precision uint
    expected UDec64
    expError error
}

func TestFloat64ToUDec64(t *testing.T) {
    testCases := []Float64ToUDec64TC{
        Float64ToUDec64TC{ 0.0, 0, 0, nil },
        Float64ToUDec64TC{ 1.0, 0, 1, nil },
        Float64ToUDec64TC{ 1.7, 0, 1, nil },
        Float64ToUDec64TC{ 145645677.18, 0, 145645677, nil },
        Float64ToUDec64TC{ 3145645677.778, 0, 3145645677, nil },
        Float64ToUDec64TC{ 187923786919586921.0, 0, 187923786919586912, nil },
        Float64ToUDec64TC{ 11792378691958692154.0, 0, 11792378691958691840, nil },
        Float64ToUDec64TC{ 145645677.18, 3, 145645677180, nil },
        Float64ToUDec64TC{ 58590303.45539292211, 11, 0x514f750e8a1a8c00, nil },
        Float64ToUDec64TC{ -1.0, 0, 0, strconv.ErrRange },
        Float64ToUDec64TC{ 18446744073709551616.0, 0, 0, strconv.ErrRange },
        Float64ToUDec64TC{ 18446744073709551617.0, 0, 0, strconv.ErrRange },
    }
    for i, tc := range testCases {
        result, err := Float64ToUDec64(tc.value, tc.precision)
        if tc.expected!=result || tc.expError!=err {
            t.Errorf("Result mismatch: %d: toudec128(%v)->%v,%v!=%v,%v",
                     i, tc.value, tc.expected, tc.expError, result, err)
        }
    }
}

type ConvertUDec64TC struct {
    value UDec64
    srcPrecision uint
    destPrecision uint
    rounding bool
    expected UDec64
}

func TestConvert(t *testing.T) {
    testCases := []ConvertUDec64TC{
        ConvertUDec64TC{ 145556, 1, 3, false, 14555600 },
        ConvertUDec64TC{ 145556, 1, 3, true, 14555600 },
        ConvertUDec64TC{ 145556, 1, 1, false, 145556 },
        ConvertUDec64TC{ 145556, 1, 1, true, 145556 },
        ConvertUDec64TC{ 145556, 3, 1, false, 1455 },
        ConvertUDec64TC{ 145556, 3, 1, true, 1456 },
    }
    for i, tc := range testCases {
        result := tc.value.Convert(tc.srcPrecision, tc.destPrecision, tc.rounding)
        if tc.expected!=result {
            t.Errorf("Result mismatch: %d: convert(%v,%v,%v,%v)->%v!=%v",
                     i, tc.value, tc.srcPrecision, tc.destPrecision, tc.rounding,
                     tc.expected, result)
        }
    }
}
