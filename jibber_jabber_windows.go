// +build windows

package jibber_jabber

import (
	"errors"
	"syscall"
	"unsafe"
)

const LOCALE_NAME_MAX_LENGTH uint32 = 85

var SUPPORTED_LOCALES = map[uintptr]string{
	0x0401: "ar-SA",
	0x0402: "bg-BG",
	0x0403: "ca-ES",
	0x0404: "zh-TW",
	0x0405: "cs-CZ",
	0x0406: "da-DK",
	0x0407: "de-DE",
	0x0408: "el-GR",
	0x0409: "en-US",
	0x040B: "fi-FI",
	0x040C: "fr-FR",
	0x040D: "he-IL",
	0x040E: "hu-HU",
	0x040F: "is-IS",
	0x0410: "it-IT",
	0x0411: "ja-JP",
	0x0412: "ko-KR",
	0x0413: "nl-NL",
	0x0414: "nb-NO",
	0x0415: "pl-PL",
	0x0416: "pt-BR",
	0x0417: "rm-CH",
	0x0418: "ro-RO",
	0x0419: "ru-RU",
	0x041A: "hr-HR",
	0x041B: "sk-SK",
	0x041C: "sq-AL",
	0x041D: "sv-SE",
	0x041E: "th-TH",
	0x041F: "tr-TR",
	0x0420: "ur-PK",
	0x0421: "id-ID",
	0x0422: "uk-UA",
	0x0423: "be-BY",
	0x0424: "sl-SI",
	0x0425: "et-EE",
	0x0426: "lv-LV",
	0x0427: "lt-LT",
	0x0429: "fa-IR",
	0x042A: "vi-VN",
	0x042B: "hy-AM",
	0x042D: "eu-ES",
	0x042F: "mk-MK",
	0x0430: "st-ZA",
	0x0431: "ts-ZA",
	0x0432: "tn-ZA",
	0x0433: "ve-ZA",
	0x0434: "xh-ZA",
	0x0435: "zu-ZA",
	0x0436: "af-ZA",
	0x0437: "ka-GE",
	0x0438: "fo-FO",
	0x0439: "hi-IN",
	0x043A: "mt-MT",
	0x043B: "se-NO",
	0x043E: "ms-MY",
	0x043F: "kk-KZ",
	0x0440: "ky-KG",
	0x0441: "sw-KE",
	0x0442: "tk-TM",
	0x0444: "tt-RU",
	0x0445: "bn-IN",
	0x0446: "pa-IN",
	0x0447: "gu-IN",
	0x0448: "or-IN",
	0x0449: "ta-IN",
	0x044A: "te-IN",
	0x044B: "kn-IN",
	0x044C: "ml-IN",
	0x044D: "as-IN",
	0x044E: "mr-IN",
	0x044F: "sa-IN",
	0x0450: "mn-MN",
	0x0451: "bo-CN",
	0x0452: "cy-GB",
	0x0453: "km-KH",
	0x0454: "lo-LA",
	0x0455: "my-MM",
	0x0456: "gl-ES",
	0x045B: "si-LK",
	0x045E: "am-ET",
	0x0461: "ne-NP",
	0x0462: "fy-NL",
	0x0463: "ps-AF",
	0x0465: "dv-MV",
	0x046A: "yo-NG",
	0x046D: "ba-RU",
	0x046E: "lb-LU",
	0x046F: "kl-GL",
	0x0470: "ig-NG",
	0x0472: "om-ET",
	0x0473: "ti-ET",
	0x0474: "gn-PY",
	0x0477: "so-SO",
	0x0478: "ii-CN",
	0x047E: "br-FR",
	0x0480: "ug-CN",
	0x0481: "mi-NZ",
	0x0482: "oc-FR",
	0x0483: "co-FR",
	0x0487: "rw-RW",
	0x0488: "wo-SN",
	0x0491: "gd-GB",
	0x0801: "ar-IQ",
	0x0804: "zh-CN",
	0x0807: "de-CH",
	0x0809: "en-GB",
	0x080A: "es-MX",
	0x080C: "fr-BE",
	0x0810: "it-CH",
	0x0813: "nl-BE",
	0x0814: "nn-NO",
	0x0816: "pt-PT",
	0x0818: "ro-MD",
	0x0819: "ru-MD",
	0x081D: "sv-FI",
	0x0820: "ur-IN",
	0x0832: "tn-BW",
	0x083B: "se-SE",
	0x083C: "ga-IE",
	0x083E: "ms-BN",
	0x0845: "bn-BD",
	0x0849: "ta-LK",
	0x0861: "ne-IN",
	0x0873: "ti-ER",
	0x0C01: "ar-EG",
	0x0C04: "zh-HK",
	0x0C07: "de-AT",
	0x0C09: "en-AU",
	0x0C0A: "es-ES",
	0x0C0C: "fr-CA",
	0x0C3B: "se-FI",
	0x0C51: "dz-BT",
	0x1001: "ar-LY",
	0x1004: "zh-SG",
	0x1007: "de-LU",
	0x1009: "en-CA",
	0x100A: "es-GT",
	0x100C: "fr-CH",
	0x101A: "hr-BA",
	0x1401: "ar-DZ",
	0x1404: "zh-MO",
	0x1407: "de-LI",
	0x1409: "en-NZ",
	0x140A: "es-CR",
	0x140C: "fr-LU",
	0x1801: "ar-MA",
	0x1809: "en-IE",
	0x180A: "es-PA",
	0x180C: "fr-MC",
	0x1C01: "ar-TN",
	0x1C09: "en-ZA",
	0x1C0A: "es-DO",
	0x2001: "ar-OM",
	0x2009: "en-JM",
	0x200A: "es-VE",
	0x200C: "fr-RE",
	0x2401: "ar-YE",
	0x240A: "es-CO",
	0x240C: "fr-CD",
	0x2801: "ar-SY",
	0x2809: "en-BZ",
	0x280A: "es-PE",
	0x280C: "fr-SN",
	0x2C01: "ar-JO",
	0x2C09: "en-TT",
	0x2C0A: "es-AR",
	0x2C0C: "fr-CM",
	0x3001: "ar-LB",
	0x3009: "en-ZW",
	0x300A: "es-EC",
	0x300C: "fr-CI",
	0x3401: "ar-KW",
	0x3409: "en-PH",
	0x340A: "es-CL",
	0x340C: "fr-ML",
	0x3801: "ar-AE",
	0x380A: "es-UY",
	0x380C: "fr-MA",
	0x3c01: "ar-BH",
	0x3c09: "en-HK",
	0x3c0A: "es-PY",
	0x3c0C: "fr-HT",
	0x4001: "ar-QA",
	0x4009: "en-IN",
	0x400A: "es-BO",
	0x4409: "en-MY",
	0x440A: "es-SV",
	0x4809: "en-SG",
	0x480A: "es-HN",
	0x4C09: "en-AE",
	0x4C0A: "es-NI",
	0x500A: "es-PR",
	0x540A: "es-US",
	0x5C0A: "es-CU",
}

func getWindowsLocaleFrom(sysCall string) (locale string, err error) {
	buffer := make([]uint16, LOCALE_NAME_MAX_LENGTH)

	dll := syscall.MustLoadDLL("kernel32")
	proc := dll.MustFindProc(sysCall)
	r, _, dllError := proc.Call(uintptr(unsafe.Pointer(&buffer[0])), uintptr(LOCALE_NAME_MAX_LENGTH))
	if r == 0 {
		err = errors.New(COULD_NOT_DETECT_PACKAGE_ERROR_MESSAGE + ":\n" + dllError.Error())
		return
	}

	locale = syscall.UTF16ToString(buffer)

	return
}

func getAllWindowsLocaleFrom(sysCall string) (string, error) {
	dll, err := syscall.LoadDLL("kernel32")
	if err != nil {
		return "", errors.New("Could not find kernel32 dll")
	}

	proc, err := dll.FindProc(sysCall)
	if err != nil {
		return "", err
	}

	locale, _, dllError := proc.Call()
	if locale == 0 {
		return "", errors.New(COULD_NOT_DETECT_PACKAGE_ERROR_MESSAGE + ":\n" + dllError.Error())
	}

	return SUPPORTED_LOCALES[locale], nil
}

func getWindowsLocale() (locale string, err error) {
	dll, err := syscall.LoadDLL("kernel32")
	if err != nil {
		return "", errors.New("Could not find kernel32 dll")
	}

	proc, err := dll.FindProc("GetVersion")
	if err != nil {
		return "", err
	}

	v, _, _ := proc.Call()
	windowsVersion := byte(v)
	isVistaOrGreater := (windowsVersion >= 6)

	if isVistaOrGreater {
		locale, err = getWindowsLocaleFrom("GetUserDefaultLocaleName")
		if err != nil {
			locale, err = getWindowsLocaleFrom("GetSystemDefaultLocaleName")
		}
	} else if !isVistaOrGreater {
		locale, err = getAllWindowsLocaleFrom("GetUserDefaultLCID")
		if err != nil {
			locale, err = getAllWindowsLocaleFrom("GetSystemDefaultLCID")
		}
	} else {
		panic(v)
	}
	return
}
func DetectIETF() (locale string, err error) {
	locale, err = getWindowsLocale()
	return
}

func DetectLanguage() (language string, err error) {
	windows_locale, err := getWindowsLocale()
	if err == nil {
		language, _ = splitLocale(windows_locale)
	}

	return
}

func DetectTerritory() (territory string, err error) {
	windows_locale, err := getWindowsLocale()
	if err == nil {
		_, territory = splitLocale(windows_locale)
	}

	return
}
