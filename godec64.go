/*
 * dec64.go - main fixed decimal int64 routines
 *
 * godec128 - go dec128 (for 12-bit decimal fixed point) library
 * Copyright (C) 2021  Mateusz Szpakowski
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
    "bytes"
    "math/bits"
    "strconv"
    "strings"
)

type UDec64 uint64

var uint64_powers []uint64 = []uint64{
    1,
    10,
    100,
    1000,
    10000,
    100000,
    1000000,
    10000000,
    100000000,
    1000000000,
    10000000000,
    100000000000,
    1000000000000,
    10000000000000,
    100000000000000,
    1000000000000000,
    10000000000000000,
    100000000000000000,
    1000000000000000000,
}

func (a UDec64) Mul(b UDec64, precision uint, rounding bool) UDec64 {
    chi, clo := bits.Mul64(uint64(a), uint64(b))
    quo, rem := bits.Div64(chi, clo, uint64_powers[precision])
    if rounding && rem>=(uint64_powers[precision]>>1) {
        quo++
    }
    return UDec64(quo)
}

func (a UDec64) MulFull(b UDec64) (UDec64, UDec64) {
    chi, clo := bits.Mul64(uint64(a), uint64(b))
    return UDec64(chi), UDec64(clo)
}

func (a UDec64) Div(b UDec64, precision uint) UDec64 {
    // multiply by precisioners
    chi, clo := bits.Mul64(uint64(a), uint64_powers[precision])
    q, _ := bits.Div64(chi, clo, uint64(b))
    return UDec64(q)
}

func DivFull(hi, lo, b UDec64) (UDec64, UDec64) {
    quo, rem := bits.Div64(uint64(hi), uint64(lo), uint64(b))
    return UDec64(quo), UDec64(rem)
}

var zeroPart []byte = []byte("0.000000000000000000000000000")

func (a UDec64) FormatNew(precision, displayPrecision uint, trimZeroes bool) string {
    if a==0 { return "0.0" }
    if precision==0 { return strconv.FormatUint(uint64(a), 10) }
    str := strconv.AppendUint(nil, uint64(a), 10)
    slen := len(str)
    i := slen
    if slen <= int(precision) {
        if trimZeroes {
            for i--; i>=0; i-- {
                if str[i]!='0' { break }
            }
            i++
        }
        var os strings.Builder
        os.Write(zeroPart[:2+int(precision)-slen])
        os.Write(str[:i])
        return os.String()
    }
    if trimZeroes {
        for i--; i>=slen-int(precision); i-- {
            if str[i]!='0' { break }
        }
        i++
    }
    var os strings.Builder
    os.Grow(i)
    os.Write(str[:slen-int(precision)])
    dotPos := os.Len()
    os.WriteByte('.')
    if (trimZeroes && precision<displayPrecision) || precision==displayPrecision {
        os.Write(str[slen-int(precision):i])
    } else if precision>displayPrecision {
        x := slen-int(precision)+int(displayPrecision) // fix
        if trimZeroes {
            if x>i { x = i }
            for x--; str[x]=='0'; x-- { }
            x++
        }
        os.Write(str[slen-int(precision):x])
    } else {
        os.Write(str[slen-int(precision):i])
        for i:=0; i < int(displayPrecision-precision); i++ {
            os.WriteByte('0')
        }
    }
    if os.Len()==dotPos+1 {
        os.WriteByte('0')
    }
    return os.String()
}

// format number
func (a UDec64) Format(precision uint, trimZeroes bool) string {
    return a.FormatNew(precision, precision, trimZeroes)
}

// new format routine with additional displayPrecision argument. Format to bytes
func (a UDec64) FormatNewBytes(precision, displayPrecision uint,
                                trimZeroes bool) []byte {
    if a==0 { return zeroPart[:3] }
    if precision==0 { return strconv.AppendUint(nil, uint64(a), 10) }
    str := strconv.AppendUint(nil, uint64(a), 10)
    slen := len(str)
    i := slen
    if slen <= int(precision) {
        if trimZeroes {
            for i--; i>=0; i-- {
                if str[i]!='0' { break }
            }
            i++
        }
        l := 2+int(precision)-slen
        os := make([]byte, l+i)
        copy(os[:l], zeroPart[:l])
        copy(os[l:], str[:i])
        return os
    }
    if trimZeroes {
        for i--; i>=slen-int(precision); i-- {
            if str[i]!='0' { break }
        }
        i++
    }
    os := make([]byte, i+1)
    l := slen-int(precision)
    copy(os[:l], str[:l])
    dotPos := l
    os[l] = '.'
    copy(os[l+1:], str[slen-int(precision):i])
    if !(trimZeroes && precision<displayPrecision) && precision!=displayPrecision {
        if precision>displayPrecision {
            x := slen+1-int(precision)+int(displayPrecision) // fix
            if trimZeroes {
                if x>i+1 { x = i+1 }
                for x--; os[x]=='0'; x-- { }
                x++
            }
            os = os[:x]
        } else {
            copy(os[l+1:], str[slen-int(precision):i])
            for i:=0; i < int(displayPrecision-precision); i++ {
                os = append(os, '0')
            }
        }
    }
    if len(os)==dotPos+1 {
        os = append(os, '0')
    }
    return os
}

// format number to bytes
func (a UDec64) FormatBytes(precision uint, trimZeroes bool) []byte {
    return a.FormatNewBytes(precision, precision, trimZeroes)
}

// parse number from string
func ParseUDec64(str string, precision uint, rounding bool) (UDec64, error) {
    slen := len(str)
    epos := strings.LastIndexByte(str, 'e')
    if epos!=-1 {
        // parse exponent
        if epos+1==slen {
            return 0, strconv.ErrSyntax
        }
        // sign of exponent
        endOfMantisa := epos
        epos++
        exponent, err := strconv.ParseInt(str[epos:], 10, 8)
        if err!=nil { return 0, err }
        
        if exponent!=0 {
            mantisa := str[:endOfMantisa]
            commaPos := strings.IndexByte(mantisa, '.')
            // move comma
            if commaPos==-1 { commaPos = endOfMantisa }
            
            newCommaPos := commaPos + int(exponent)
            
            i := 0
            for ; str[i]=='0' || str[i]=='.'; i++ {
                if str[i]=='.' { continue }
                newCommaPos--
            } // skip first zero
            //fmt.Println("NewCommaPos:", newCommaPos, commaPos, exponent, i)
            var sb strings.Builder
            // add zeroes
            if newCommaPos<0 {
                sb.WriteRune('.')
                for ; newCommaPos<0;  newCommaPos++ {
                    sb.WriteRune('0')
                }
            } else if newCommaPos>0 {
                for ; newCommaPos>0;  newCommaPos-- {
                    if str[i]=='.' { i++ }
                    if i<endOfMantisa {
                        sb.WriteByte(str[i])
                        i++
                    } else {
                        sb.WriteRune('0')
                    }
                }
                if i<endOfMantisa {
                    sb.WriteRune('.') // append new comma
                }
            }
            // to end of mantisa
            for ; i<endOfMantisa; i++ {
                if str[i]=='.' { i++ }
                if i<endOfMantisa { sb.WriteByte(str[i]) }
            }
            
            //fmt.Println("new str:", sb.String())
            str = sb.String()
            slen = len(str)
            if slen==0 { return 0, nil }
        } else {
            str = str[:endOfMantisa]
            slen = len(str)
        }
    }
    
    commaIdx := strings.LastIndexByte(str, '.')
    if commaIdx==-1 {
        // comma not found
        v, err := strconv.ParseUint(str, 10, 64)
        if err!=nil { return UDec64(v), err }
        chi, clo := bits.Mul64(v, uint64_powers[precision])
        if chi!=0 {
            return 0, strconv.ErrRange
        }
        return UDec64(clo), nil
    }
    if slen-(commaIdx+1) >= int(precision) {
        //  more than in fraction
        realSlen := commaIdx+1+int(precision)
        s2 := str[:commaIdx] + str[commaIdx+1:realSlen]
        v, err := strconv.ParseUint(s2, 10, 64)
        if err!=nil { return 0, err }
        // rounding
        if rounding && realSlen!=slen && str[realSlen]>='5' {
            v++ // add rounding
        }
        // check last part of string
        for i:=realSlen; i<slen; i++ {
            if str[i]<'0' || str[i]>'9' {
                return 0, strconv.ErrSyntax
            }
        }
        return UDec64(v), nil
    } else {
        // less than in fraction
        s2 := str[:commaIdx] + str[commaIdx+1:]
        v, err := strconv.ParseUint(s2, 10, 64)
        if err!=nil { return 0, err }
        pow10ForVal := int(precision) - (slen-(commaIdx+1))
        chi, clo := bits.Mul64(v, uint64_powers[pow10ForVal])
        if chi!=0 {
            return 0, strconv.ErrRange
        }
        return UDec64(clo), nil
    }
    return 0, nil
}

func ParseUIntDecBytes(s []byte, bits int) (uint64, error) {
    maxVal := uint64(1<<bits)-1
    v := uint64(0)
    for _, c := range s {
        vp := v
        if c>='0' && c<='9' {
            v = v*10 + uint64(c-'0')
        } else {
            return 0, strconv.ErrSyntax
        }
        v &= maxVal
        if vp > v || v>maxVal {
            return 0, strconv.ErrRange
        }
    }
    return v, nil
}

// parse number from bytes
func ParseUDec64Bytes(str []byte, precision uint, rounding bool) (UDec64, error) {
    slen := len(str)
    epos := bytes.LastIndexByte(str, 'e')
    if epos!=-1 {
        // parse exponent
        if epos+1==slen {
            return 0, strconv.ErrSyntax
        }
        // sign of exponent
        endOfMantisa := epos
        epos++
        exponent, err := strconv.ParseInt(string(str[epos:]), 10, 8)
        if err!=nil { return 0, err }
        
        if exponent!=0 {
            mantisa := str[:endOfMantisa]
            commaPos := bytes.IndexByte(mantisa, '.')
            // move comma
            if commaPos==-1 { commaPos = endOfMantisa }
            
            newCommaPos := commaPos + int(exponent)
            
            i := 0
            for ; str[i]=='0' || str[i]=='.'; i++ {
                if str[i]=='.' { continue }
                newCommaPos--
            } // skip first zero
            //fmt.Println("NewCommaPos:", newCommaPos, commaPos, exponent, i)
            var sb bytes.Buffer
            // add zeroes
            if newCommaPos<0 {
                sb.WriteRune('.')
                for ; newCommaPos<0;  newCommaPos++ {
                    sb.WriteRune('0')
                }
            } else if newCommaPos>0 {
                for ; newCommaPos>0;  newCommaPos-- {
                    if str[i]=='.' { i++ }
                    if i<endOfMantisa {
                        sb.WriteByte(str[i])
                        i++
                    } else {
                        sb.WriteRune('0')
                    }
                }
                if i<endOfMantisa {
                    sb.WriteRune('.') // append new comma
                }
            }
            // to end of mantisa
            for ; i<endOfMantisa; i++ {
                if str[i]=='.' { i++ }
                if i<endOfMantisa { sb.WriteByte(str[i]) }
            }
            
            //fmt.Println("new str:", sb.String())
            str = sb.Bytes()
            slen = len(str)
            if slen==0 { return 0, nil }
        } else {
            str = str[:endOfMantisa]
            slen = len(str)
        }
    }
    
    commaIdx := bytes.LastIndexByte(str, '.')
    if commaIdx==-1 {
        // comma not found
        v, err := ParseUIntDecBytes(str, 64)
        if err!=nil { return UDec64(v), err }
        chi, clo := bits.Mul64(v, uint64_powers[precision])
        if chi!=0 {
            return 0, strconv.ErrRange
        }
        return UDec64(clo), nil
    }
    if slen-(commaIdx+1) >= int(precision) {
        //  more than in fraction
        realSlen := commaIdx+1+int(precision)
        s2 := make([]byte, realSlen-1)
        copy(s2[:commaIdx], str[:commaIdx])
        copy(s2[commaIdx:], str[commaIdx+1:])
        v, err := ParseUIntDecBytes(s2, 64)
        if err!=nil { return 0, err }
        // rounding
        if rounding && realSlen!=slen && str[realSlen]>='5' {
            v++ // add rounding
        }
        // check last part of string
        for i:=realSlen; i<slen; i++ {
            if str[i]<'0' || str[i]>'9' {
                return 0, strconv.ErrSyntax
            }
        }
        return UDec64(v), nil
    } else {
        // less than in fraction
        s2 := make([]byte, slen-1)
        copy(s2[:commaIdx], str[:commaIdx])
        copy(s2[commaIdx:], str[commaIdx+1:])
        v, err := ParseUIntDecBytes(s2, 64)
        if err!=nil { return 0, err }
        pow10ForVal := int(precision) - (slen-(commaIdx+1))
        chi, clo := bits.Mul64(v, uint64_powers[pow10ForVal])
        if chi!=0 {
            return 0, strconv.ErrRange
        }
        return UDec64(clo), nil
    }
    return 0, nil
}
