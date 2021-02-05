/*
 * locale_test.go - tests for int128 routines
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

package godec64

import (
    "strconv"
    "testing"
)

type UDec64LocTC struct {
    lang string
    noSep1000 bool
    a UDec64
    precision uint
    trimZeroes bool
    expected string
}

func TestUDec64LocaleFormat(t *testing.T) {
    testCases := []UDec64LocTC {
        UDec64LocTC{ "af", false, 0xab54a98ceb1f0ad3,
                10, false, "1 234 567 890,1234567891" },
        UDec64LocTC{ "af", false, 0xab54a98ceb1f0ad2,
                10, false, "1 234 567 890,1234567890" },
        UDec64LocTC{ "af", false, 0xab54a98ceb1f0ad2,
                10, true, "1 234 567 890,123456789" },
        UDec64LocTC{ "am", false, 0xab54a98ceb1f0ad3,
                10, false, "1,234,567,890.1234567891" },
        UDec64LocTC{ "ar", false, 0xab54a98ceb1f0ad3,
                10, false, "١٬٢٣٤٬٥٦٧٬٨٩٠٫١٢٣٤٥٦٧٨٩١" },
        UDec64LocTC{ "az", false, 0xab54a98ceb1f0ad3,
                10, false, "1.234.567.890,1234567891" },
        UDec64LocTC{ "bg", false, 0xab54a98ceb1f0ad3,
                10, false, "1 234 567 890,1234567891" },
        UDec64LocTC{ "bn", false, 0xab54a98ceb1f0ad3,
                10, false, "১,২৩,৪৫,৬৭,৮৯০.১২৩৪৫৬৭৮৯১" },
        UDec64LocTC{ "ca", false, 0xab54a98ceb1f0ad3,
                10, false, "1.234.567.890,1234567891" },
        UDec64LocTC{ "cs", false, 0xab54a98ceb1f0ad3,
                10, false, "1 234 567 890,1234567891" },
        UDec64LocTC{ "da", false, 0xab54a98ceb1f0ad3,
                10, false, "1.234.567.890,1234567891" },
        UDec64LocTC{ "de", false, 0xab54a98ceb1f0ad3,
                10, false, "1.234.567.890,1234567891" },
        UDec64LocTC{ "el", false, 0xab54a98ceb1f0ad3,
                10, false, "1.234.567.890,1234567891" },
        UDec64LocTC{ "en", false, 0xab54a98ceb1f0ad3,
                10, false, "1,234,567,890.1234567891" },
        UDec64LocTC{ "es", false, 0xab54a98ceb1f0ad3,
                10, false, "1.234.567.890,1234567891" },
        UDec64LocTC{ "et", false, 0xab54a98ceb1f0ad3,
                10, false, "1 234 567 890,1234567891" },
        UDec64LocTC{ "fa", false, 0xab54a98ceb1f0ad3,
                10, false, "۱٬۲۳۴٬۵۶۷٬۸۹۰٫۱۲۳۴۵۶۷۸۹۱" },
        UDec64LocTC{ "fi", false, 0xab54a98ceb1f0ad3,
                10, false, "1 234 567 890,1234567891" },
        UDec64LocTC{ "fil", false, 0xab54a98ceb1f0ad3,
                10, false, "1,234,567,890.1234567891" },
        UDec64LocTC{ "fr", false, 0xab54a98ceb1f0ad3,
                10, false, "1 234 567 890,1234567891" },
        UDec64LocTC{ "gu", false, 0xab54a98ceb1f0ad3,
                10, false, "1,23,45,67,890.1234567891" },
        UDec64LocTC{ "he", false, 0xab54a98ceb1f0ad3,
                10, false, "1,234,567,890.1234567891" },
        UDec64LocTC{ "hi", false, 0xab54a98ceb1f0ad3,
                10, false, "1,23,45,67,890.1234567891" },
        UDec64LocTC{ "hr", false, 0xab54a98ceb1f0ad3,
                10, false, "1.234.567.890,1234567891" },
        UDec64LocTC{ "hu", false, 0xab54a98ceb1f0ad3,
                10, false, "1 234 567 890,1234567891" },
        UDec64LocTC{ "hy", false, 0xab54a98ceb1f0ad3,
                10, false, "1 234 567 890,1234567891" },
        UDec64LocTC{ "id", false, 0xab54a98ceb1f0ad3,
                10, false, "1.234.567.890,1234567891" },
        UDec64LocTC{ "is", false, 0xab54a98ceb1f0ad3,
                10, false, "1.234.567.890,1234567891" },
        UDec64LocTC{ "it", false, 0xab54a98ceb1f0ad3,
                10, false, "1.234.567.890,1234567891" },
        UDec64LocTC{ "ja", false, 0xab54a98ceb1f0ad3,
                10, false, "1,234,567,890.1234567891" },
        UDec64LocTC{ "ka", false, 0xab54a98ceb1f0ad3,
                10, false, "1 234 567 890,1234567891" },
        UDec64LocTC{ "kk", false, 0xab54a98ceb1f0ad3,
                10, false, "1 234 567 890,1234567891" },
        UDec64LocTC{ "km", false, 0xab54a98ceb1f0ad3,
                10, false, "1.234.567.890,1234567891" },
        UDec64LocTC{ "kn", false, 0xab54a98ceb1f0ad3,
                10, false, "1,234,567,890.1234567891" },
        UDec64LocTC{ "ko", false, 0xab54a98ceb1f0ad3,
                10, false, "1,234,567,890.1234567891" },
        UDec64LocTC{ "ky", false, 0xab54a98ceb1f0ad3,
                10, false, "1 234 567 890,1234567891" },
        UDec64LocTC{ "lo", false, 0xab54a98ceb1f0ad3,
                10, false, "1.234.567.890,1234567891" },
        UDec64LocTC{ "lt", false, 0xab54a98ceb1f0ad3,
                10, false, "1 234 567 890,1234567891" },
        UDec64LocTC{ "lv", false, 0xab54a98ceb1f0ad3,
                10, false, "1 234 567 890,1234567891" },
        UDec64LocTC{ "mk", false, 0xab54a98ceb1f0ad3,
                10, false, "1.234.567.890,1234567891" },
        UDec64LocTC{ "ml", false, 0xab54a98ceb1f0ad3,
                10, false, "1,23,45,67,890.1234567891" },
        UDec64LocTC{ "mn", false, 0xab54a98ceb1f0ad3,
                10, false, "1,234,567,890.1234567891" },
        UDec64LocTC{ "mo", false, 0xab54a98ceb1f0ad3,
                10, false, "1.234.567.890,1234567891" },
        UDec64LocTC{ "mr", false, 0xab54a98ceb1f0ad3,
                10, false, "१,२३,४५,६७,८९०.१२३४५६७८९१" },
        UDec64LocTC{ "ms", false, 0xab54a98ceb1f0ad3,
                10, false, "1,234,567,890.1234567891" },
        UDec64LocTC{ "mul", false, 0xab54a98ceb1f0ad3,
                10, false, "1,234,567,890.1234567891" },
        UDec64LocTC{ "my", false, 0xab54a98ceb1f0ad3,
                10, false, "၁,၂၃၄,၅၆၇,၈၉၀.၁၂၃၄၅၆၇၈၉၁" },
        UDec64LocTC{ "nb", false, 0xab54a98ceb1f0ad3,
                10, false, "1 234 567 890,1234567891" },
        UDec64LocTC{ "ne", false, 0xab54a98ceb1f0ad3,
                10, false, "१,२३४,५६७,८९०.१२३४५६७८९१" },
        UDec64LocTC{ "nl", false, 0xab54a98ceb1f0ad3,
                10, false, "1.234.567.890,1234567891" },
        UDec64LocTC{ "no", false, 0xab54a98ceb1f0ad3,
                10, false, "1,234,567,890.1234567891" },
        UDec64LocTC{ "pa", false, 0xab54a98ceb1f0ad3,
                10, false, "1,23,45,67,890.1234567891" },
        UDec64LocTC{ "pl", false, 0xab54a98ceb1f0ad3,
                10, false, "1 234 567 890,1234567891" },
        UDec64LocTC{ "pt", false, 0xab54a98ceb1f0ad3,
                10, false, "1.234.567.890,1234567891" },
        UDec64LocTC{ "ro", false, 0xab54a98ceb1f0ad3,
                10, false, "1.234.567.890,1234567891" },
        UDec64LocTC{ "ru", false, 0xab54a98ceb1f0ad3,
                10, false, "1 234 567 890,1234567891" },
        UDec64LocTC{ "sh", false, 0xab54a98ceb1f0ad3,
                10, false, "1.234.567.890,1234567891" },
        UDec64LocTC{ "si", false, 0xab54a98ceb1f0ad3,
                10, false, "1,234,567,890.1234567891" },
        UDec64LocTC{ "sk", false, 0xab54a98ceb1f0ad3,
                10, false, "1 234 567 890,1234567891" },
        UDec64LocTC{ "sl", false, 0xab54a98ceb1f0ad3,
                10, false, "1.234.567.890,1234567891" },
        UDec64LocTC{ "sq", false, 0xab54a98ceb1f0ad3,
                10, false, "1 234 567 890,1234567891" },
        UDec64LocTC{ "sr", false, 0xab54a98ceb1f0ad3,
                10, false, "1.234.567.890,1234567891" },
        UDec64LocTC{ "sv", false, 0xab54a98ceb1f0ad3,
                10, false, "1 234 567 890,1234567891" },
        UDec64LocTC{ "sw", false, 0xab54a98ceb1f0ad3,
                10, false, "1,234,567,890.1234567891" },
        UDec64LocTC{ "ta", false, 0xab54a98ceb1f0ad3,
                10, false, "1,23,45,67,890.1234567891" },
        UDec64LocTC{ "te", false, 0xab54a98ceb1f0ad3,
                10, false, "1,234,567,890.1234567891" },
        UDec64LocTC{ "th", false, 0xab54a98ceb1f0ad3,
                10, false, "1,234,567,890.1234567891" },
        UDec64LocTC{ "tl", false, 0xab54a98ceb1f0ad3,
                10, false, "1,234,567,890.1234567891" },
        UDec64LocTC{ "tn", false, 0xab54a98ceb1f0ad3,
                10, false, "1,234,567,890.1234567891" },
        UDec64LocTC{ "tr", false, 0xab54a98ceb1f0ad3,
                10, false, "1.234.567.890,1234567891" },
        UDec64LocTC{ "uk", false, 0xab54a98ceb1f0ad3,
                10, false, "1 234 567 890,1234567891" },
        UDec64LocTC{ "ur", false, 0xab54a98ceb1f0ad3,
                10, false, "1,234,567,890.1234567891" },
        UDec64LocTC{ "uz", false, 0xab54a98ceb1f0ad3,
                10, false, "1 234 567 890,1234567891" },
        UDec64LocTC{ "vi", false, 0xab54a98ceb1f0ad3,
                10, false, "1.234.567.890,1234567891" },
        UDec64LocTC{ "zh", false, 0xab54a98ceb1f0ad3,
                10, false, "1,234,567,890.1234567891" },
        UDec64LocTC{ "zu", false, 0xab54a98ceb1f0ad3,
                10, false, "1,234,567,890.1234567891" },
        UDec64LocTC{ "", false, 0xab54a98ceb1f0ad3, 10,
                false, "1,234,567,890.1234567891" },
        UDec64LocTC{ "C", false, 0xab54a98ceb1f0ad3,
                10, false, "1,234,567,890.1234567891" },
        UDec64LocTC{ "pl-PL", false, 0xab54a98ceb1f0ad3,
                10, false, "1 234 567 890,1234567891" },
        UDec64LocTC{ "pl_PL", false, 0xab54a98ceb1f0ad3,
                10, false, "1 234 567 890,1234567891" },
        UDec64LocTC{ "pl_PL.UTF-8", false, 0xab54a98ceb1f0ad3,
                10, false, "1 234 567 890,1234567891" },
        // no separator 1000
        UDec64LocTC{ "af", true, 0xab54a98ceb1f0ad3,
                10, false, "1234567890,1234567891" },
        UDec64LocTC{ "am", true, 0xab54a98ceb1f0ad3,
                10, false, "1234567890.1234567891" },
        UDec64LocTC{ "ar", true, 0xab54a98ceb1f0ad3,
                10, false, "١٢٣٤٥٦٧٨٩٠٫١٢٣٤٥٦٧٨٩١" },
        UDec64LocTC{ "az", true, 0xab54a98ceb1f0ad3,
                10, false, "1234567890,1234567891" },
        UDec64LocTC{ "bg", true, 0xab54a98ceb1f0ad3,
                10, false, "1234567890,1234567891" },
        UDec64LocTC{ "bn", true, 0xab54a98ceb1f0ad3,
                10, false, "১২৩৪৫৬৭৮৯০.১২৩৪৫৬৭৮৯১" },
        UDec64LocTC{ "ca", true, 0xab54a98ceb1f0ad3,
                10, false, "1234567890,1234567891" },
    }
    for i, tc := range testCases {
        a := tc.a
        result := tc.a.LocaleFormat(tc.lang, tc.precision, tc.trimZeroes, tc.noSep1000)
        if tc.expected!=result {
            t.Errorf("Result mismatch: %d: fmt(%v,%s,%v,%v)->%v!=%v",
                     i, tc.a, tc.lang, tc.precision, tc.trimZeroes, tc.expected, result)
        }
        if tc.a!=a {
            t.Errorf("Argument has been modified: %d %s: %v!=%v", i, tc.lang, a, tc.a)
        }
        resultBytes := tc.a.LocaleFormatBytes(tc.lang, tc.precision,
                        tc.trimZeroes, tc.noSep1000)
        if tc.expected!=string(resultBytes) {
            t.Errorf("Result mismatch: %d: fmtBytes(%v,%s,%v,%v)->%v!=%v",
                     i, tc.a, tc.lang, tc.precision, tc.trimZeroes, tc.expected, result)
        }
        if tc.a!=a {
            t.Errorf("Argument has been modified: %d %s: %v!=%v", i, tc.lang, a, tc.a)
        }
    }
}

type UDec64LocParseTC struct {
    lang string
    str string
    precision uint
    rounding bool
    expected UDec64
    expError error
}

func TestUDec64LocaleParse(t *testing.T) {
    testCases := []UDec64LocParseTC {
        UDec64LocParseTC{ "en", "", 10, false, 0, strconv.ErrSyntax },
        UDec64LocParseTC{ "en", "1,234,567,890.1234567891", 10, false,
                0xab54a98ceb1f0ad3, nil },
        UDec64LocParseTC{ "en", "1,234,567,890.12345678915", 10, false,
                0xab54a98ceb1f0ad3, nil },
        UDec64LocParseTC{ "en", "1,234,567,890.12345678915", 10, true,
                0xab54a98ceb1f0ad4, nil },
        UDec64LocParseTC{ "en", "1,234,567890.1234567891", 10, false,
                0xab54a98ceb1f0ad3, nil },
        UDec64LocParseTC{ "pl", "1 234 567 890,1234567891", 10, false,
                0xab54a98ceb1f0ad3, nil },
        UDec64LocParseTC{ "pl", "1 234 567 890,1234567891", 10, false,
                0xab54a98ceb1f0ad3, nil },
        UDec64LocParseTC{ "bn", "১,২৩,৪৫,৬৭,৮৯০.১২৩৪৫৬৭৮৯১", 10, false,
                0xab54a98ceb1f0ad3, nil },
        UDec64LocParseTC{ "bn", "1,234,567890.1234567891", 10, false,
                0xab54a98ceb1f0ad3, nil },
    }
    for i, tc := range testCases {
        result, err := LocaleParseUDec64(tc.lang, tc.str, tc.precision, tc.rounding)
        if tc.expected!=result || tc.expError!=err {
            t.Errorf("Result mismatch: %d: parse(%v,%v,%v,%v)->%v,%v!=%v,%v",
                     i, tc.lang, tc.str, tc.precision, tc.rounding,
                     tc.expected, tc.expError, result, err)
        }
        result, err = LocaleParseUDec64Bytes(tc.lang, []byte(tc.str),
                                tc.precision, tc.rounding)
        if tc.expected!=result || tc.expError!=err {
            t.Errorf("Result mismatch: %d: parseBytes(%v,%v,%v,%v)->%v,%v!=%v,%v",
                     i, tc.lang, tc.str, tc.precision, tc.rounding,
                     tc.expected, tc.expError, result, err)
        }
    }
}
