/*
 * locale.go - locale
 *
 * godec128 - go dec64 (for 64-bit decimal fixed point) library
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
    "strings"
    "strconv"
    "unicode/utf8"
)

// locale formatting info
type LocFmt struct {
    Comma, Sep1000, Sep1000_2 rune
    Sep100and1000 bool
    Digits []rune
}

var normalDigits []rune = []rune("0123456789")
var arDigits []rune = []rune("٠١٢٣٤٥٦٧٨٩")
var faDigits []rune = []rune("۰۱۲۳۴۵۶۷۸۹")
var bnDigits []rune = []rune("০১২৩৪৫৬৭৮৯")
var mrDigits []rune = []rune("०१२३४५६७८९")
var myDigits []rune = []rune("၀၁၂၃၄၅၆၇၈၉")

var defaultLocaleFormat LocFmt = LocFmt{ '.', ',', ',', false, normalDigits }

var localeFormats map[string]LocFmt = map[string]LocFmt {
    "af": LocFmt{ ',', ' ', ' ', false, normalDigits },
    "am": LocFmt{ '.', ',', ',', false, normalDigits },
    "ar": LocFmt{ '٫', '٬', '٬', false, arDigits },
    "az": LocFmt{ ',', '.', '.', false, normalDigits },
    "bg": LocFmt{ ',', ' ', ' ', false, normalDigits },
    "bn": LocFmt{ '.', ',', ',', true, bnDigits },
    "ca": LocFmt{ ',', '.', '.',false, normalDigits },
    "cs": LocFmt{ ',', ' ', ' ', false, normalDigits },
    "da": LocFmt{ ',', '.', '.', false, normalDigits },
    "de": LocFmt{ ',', '.', '.', false, normalDigits },
    "el": LocFmt{ ',', '.', '.', false, normalDigits },
    "en": LocFmt{ '.', ',', ',', false, normalDigits },
    "es": LocFmt{ ',', '.', '.', false, normalDigits },
    "et": LocFmt{ ',', ' ', ' ', false, normalDigits },
    "fa": LocFmt{ '٫', '٬', '٬', false, faDigits },
    "fi": LocFmt{ ',', ' ', ' ', false, normalDigits },
    "fil": LocFmt{ '.', ',', ',', false, normalDigits },
    "fr": LocFmt{ ',', ' ', ' ', false, normalDigits },
    "gu": LocFmt{ '.', ',', ',', true, normalDigits },
    "he": LocFmt{ '.', ',', ',', false, normalDigits },
    "hi": LocFmt{ '.', ',', ',', true, normalDigits },
    "hr": LocFmt{ ',', '.', '.', false, normalDigits },
    "hu": LocFmt{ ',', ' ', ' ', false, normalDigits },
    "hy": LocFmt{ ',', ' ', ' ', false, normalDigits },
    "id": LocFmt{ ',', '.', '.', false, normalDigits },
    "is": LocFmt{ ',', '.', '.', false, normalDigits },
    "it": LocFmt{ ',', '.', '.', false, normalDigits },
    "ja": LocFmt{ '.', ',', ',', false, normalDigits },
    "ka": LocFmt{ ',', ' ', ' ', false, normalDigits },
    "kk": LocFmt{ ',', ' ', ' ', false, normalDigits },
    "km": LocFmt{ ',', '.', '.', false, normalDigits },
    "kn": LocFmt{ '.', ',', ',', false, normalDigits },
    "ko": LocFmt{ '.', ',', ',', false, normalDigits },
    "ky": LocFmt{ ',', ' ', ' ', false, normalDigits },
    "lo": LocFmt{ ',', '.', '.', false, normalDigits },
    "lt": LocFmt{ ',', ' ', ' ', false, normalDigits },
    "lv": LocFmt{ ',', ' ', ' ', false, normalDigits },
    "mk": LocFmt{ ',', '.', '.', false, normalDigits },
    "ml": LocFmt{ '.', ',', ',', true, normalDigits },
    "mn": LocFmt{ '.', ',', ',', false, normalDigits },
    "mo": LocFmt{ ',', '.', '.', false, normalDigits },
    "mr": LocFmt{ '.', ',', ',', true, mrDigits },
    "ms": LocFmt{ '.', ',', ',', false, normalDigits },
    "mul": LocFmt{ '.', ',', ',', false, normalDigits },
    "my": LocFmt{ '.', ',', ',', false, myDigits },
    "nb": LocFmt{ ',', ' ', ' ', false, normalDigits },
    "ne": LocFmt{ '.', ',', ',', false, mrDigits },
    "nl": LocFmt{ ',', '.', '.', false, normalDigits },
    "no": LocFmt{ '.', ',', ',', false, normalDigits },
    "pa": LocFmt{ '.', ',', ',', true, normalDigits },
    "pl": LocFmt{ ',', ' ', ' ', false, normalDigits },
    "pt": LocFmt{ ',', '.', '.', false, normalDigits },
    "ro": LocFmt{ ',', '.', '.', false, normalDigits },
    "ru": LocFmt{ ',', ' ', ' ', false, normalDigits },
    "sh": LocFmt{ ',', '.', '.', false, normalDigits },
    "si": LocFmt{ '.', ',', ',', false, normalDigits },
    "sk": LocFmt{ ',', ' ', ' ', false, normalDigits },
    "sl": LocFmt{ ',', '.', '.', false, normalDigits },
    "sq": LocFmt{ ',', ' ', ' ', false, normalDigits },
    "sr": LocFmt{ ',', '.', '.', false, normalDigits },
    "sv": LocFmt{ ',', ' ', ' ', false, normalDigits },
    "sw": LocFmt{ '.', ',', ',', false, normalDigits },
    "ta": LocFmt{ '.', ',', ',', true, normalDigits },
    "te": LocFmt{ '.', ',', ',', false, normalDigits },
    "th": LocFmt{ '.', ',', ',', false, normalDigits },
    "tl": LocFmt{ '.', ',', ',', false, normalDigits },
    "tn": LocFmt{ '.', ',', ',', false, normalDigits },
    "tr": LocFmt{ ',', '.', '.', false, normalDigits },
    "uk": LocFmt{ ',', ' ', ' ', false, normalDigits },
    "ur": LocFmt{ '.', ',', ',', false, normalDigits },
    "uz": LocFmt{ ',', ' ', ' ', false, normalDigits },
    "vi": LocFmt{ ',', '.', '.', false, normalDigits },
    "zh": LocFmt{ '.', ',', ',', false, normalDigits },
    "zu": LocFmt{ '.', ',', ',', false, normalDigits },
}

// get locale formating info
func GetLocFmt(lang string) *LocFmt {
    outLang := lang
    langSlen := len(lang)
    if langSlen>=3 && (lang[2]=='_' || lang[2]=='-') {
        outLang = lang[0:2]
    } else if langSlen>=4 && (lang[3]=='_' || lang[3]=='-') {
        outLang = lang[0:3]
    }
    l, ok := localeFormats[outLang]
    if !ok { l = defaultLocaleFormat }
    return &l
}

// format 64-bit decimal fixed point including locale
func (a UDec64) LocaleFormatNewBytes(lang string, precision, displayPrecision uint,
                                trimZeroes, noSep1000 bool) []byte {
    l := GetLocFmt(lang)
    s := a.FormatNewBytes(precision, displayPrecision, trimZeroes)
    slen := len(s)
    os := make([]byte, slen<<1) // optimization
    oslen := 0
    commaIdx := bytes.LastIndexByte(s, '.')
    if commaIdx==-1 {
        commaIdx = slen
    }
    ti := commaIdx
    i := commaIdx
    if !l.Sep100and1000 {
        ti = (commaIdx)%3
        if ti==0 { ti=3 }
    }
    for k:=0; k < commaIdx; k++ {
        r := s[k]
        if r>='0' && r<='9' {
            if oslen+4 >= len(os) {
                os = append(os, 0,0,0,0)
            }
            oslen += utf8.EncodeRune(os[oslen:], l.Digits[r-'0'])
        }
        if !noSep1000 && i!=1 {
            if !l.Sep100and1000 || ti<=3 {
                ti--
                if ti==0 {
                    if oslen+4 >= len(os) {
                        os = append(os, 0,0,0,0)
                    }
                    oslen += utf8.EncodeRune(os[oslen:], l.Sep1000)
                    ti = 3
                }
            } else {
                ti--
                if (ti-3)&1==0 {
                    if oslen+4 >= len(os) {
                        os = append(os, 0,0,0,0)
                    }
                    oslen += utf8.EncodeRune(os[oslen:], l.Sep1000)
                }
            }
        }
        i--
    }
    // comma
    if commaIdx!=slen {
        if oslen+4 >= len(os) {
            os = append(os, 0,0,0,0)
        }
        oslen += utf8.EncodeRune(os[oslen:], l.Comma)
        for i = commaIdx+1; i < slen; i++ {
            if oslen+4 >= len(os) {
                os = append(os, 0,0,0,0)
            }
            oslen += utf8.EncodeRune(os[oslen:], l.Digits[s[i]-'0'])
        }
    }
    return os[:oslen]
}

func (a UDec64) LocaleFormatBytes(lang string, precision uint,
                                trimZeroes, noSep1000 bool) []byte {
    return a.LocaleFormatNewBytes(lang, precision, precision, trimZeroes, noSep1000)
}

// format 64-bit decimal fixed point including locale
func (a UDec64) LocaleFormatNew(lang string, precision, displayPrecision uint,
                            trimZeroes, noSep1000 bool) string {
    l := GetLocFmt(lang)
    s := a.FormatNewBytes(precision, displayPrecision, trimZeroes)
    var os strings.Builder
    slen := len(s)
    os.Grow(slen*3)
    commaIdx := bytes.LastIndexByte(s, '.')
    if commaIdx==-1 {
        commaIdx = slen
    }
    ti := commaIdx
    i := commaIdx
    if !l.Sep100and1000 {
        ti = (commaIdx)%3
        if ti==0 { ti=3 }
    }
    for k:=0; k < commaIdx; k++ {
        r := s[k]
        if r>='0' && r<='9' {
            os.WriteRune(l.Digits[r-'0'])
        }
        if !noSep1000 && i!=1 {
            if !l.Sep100and1000 || ti<=3 {
                ti--
                if ti==0 {
                    os.WriteRune(l.Sep1000)
                    ti = 3
                }
            } else {
                ti--
                if (ti-3)&1==0 {
                    os.WriteRune(l.Sep1000)
                }
            }
        }
        i--
    }
    // comma
    if commaIdx!=slen {
        os.WriteRune(l.Comma)
        for i = commaIdx+1; i < slen; i++ {
            os.WriteRune(l.Digits[s[i]-'0'])
        }
    }
    return os.String()
}

func (a UDec64) LocaleFormat(lang string, precision uint,
                            trimZeroes, noSep1000 bool) string {
    return a.LocaleFormatNew(lang, precision, precision, trimZeroes, noSep1000)
}

// parse decimal fixed point from string and return value and error (nil if no error)
func LocaleParseUDec64(lang, str string, precision uint, rounding bool) (UDec64, error) {
    l := GetLocFmt(lang)
    if len(str)==0 { return 0, strconv.ErrSyntax }
    
    os := make([]byte, 0, len(str))
    for _, r := range str {
        if r>='0' && r<='9' {
            // if standard digits
            os = append(os, byte(r))
        } else if r!=l.Sep1000 && r!=l.Sep1000_2 && r!=l.Comma {
            // if non-standard digit
            dig:=0
            found := false
            for ; dig<=9; dig++ {
                if l.Digits[dig]==r {
                    found = true
                    break
                }
            }
            if !found { return 0, strconv.ErrSyntax }
            os = append(os, '0'+byte(dig))
        } else if r==l.Comma {
            os = append(os, '.')
        }
        // otherwise skip sep1000
    }
    return ParseUDec64Bytes(os, precision, rounding)
}

// parse decimal fixed point from string and return value and error (nil if no error)
func LocaleParseUDec64Bytes(lang string, strInput []byte,
                             precision uint, rounding bool) (UDec64, error) {
    l := GetLocFmt(lang)
    if len(strInput)==0 { return 0, strconv.ErrSyntax }
    
    os := make([]byte, 0, len(strInput))
    str := strInput
    for len(str)>0 {
        r, size := utf8.DecodeRune(str)
        if r>='0' && r<='9' {
            // if standard digits
            os = append(os, byte(r))
        } else if r!=l.Sep1000 && r!=l.Sep1000_2 && r!=l.Comma {
            // if non-standard digit
            dig:=0
            found := false
            for ; dig<=9; dig++ {
                if l.Digits[dig]==r {
                    found = true
                    break
                }
            }
            if !found { return 0, strconv.ErrSyntax }
            os = append(os, '0'+byte(dig))
        } else if r==l.Comma {
            os = append(os, '.')
        }
        // otherwise skip sep1000
        str = str[size:]
    }
    return ParseUDec64Bytes(os, precision, rounding)
}
