// Copyright (C) 2018 Storj Labs, Inc.
// See LICENSE for copying information.

package memory

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

// base 2 and base 10 sizes
const (
	B Size = 1 << (10 * iota)
	KiB
	MiB
	GiB
	TiB
	PiB
	EiB

	KB Size = 1e3
	MB Size = 1e6
	GB Size = 1e9
	TB Size = 1e12
	PB Size = 1e15
	EB Size = 1e18
)

// Size implements flag.Value for collecting memory size in bytes
type Size int64

// Int returns bytes size as int
func (size Size) Int() int { return int(size) }

// Int32 returns bytes size as int32
func (size Size) Int32() int32 { return int32(size) }

// Int64 returns bytes size as int64
func (size Size) Int64() int64 { return int64(size) }

// Float64 returns bytes size as float64
func (size Size) Float64() float64 { return float64(size) }

// KiB returns size in kibibytes
func (size Size) KiB() float64 { return size.Float64() / KiB.Float64() }

// MiB returns size in mebibytes
func (size Size) MiB() float64 { return size.Float64() / MiB.Float64() }

// GiB returns size in gibibytes
func (size Size) GiB() float64 { return size.Float64() / GiB.Float64() }

// TiB returns size in tebibytes
func (size Size) TiB() float64 { return size.Float64() / TiB.Float64() }

// PiB returns size in pebibytes
func (size Size) PiB() float64 { return size.Float64() / PiB.Float64() }

// EiB returns size in exbibytes
func (size Size) EiB() float64 { return size.Float64() / EiB.Float64() }

// KB returns size in kilobytes
func (size Size) KB() float64 { return size.Float64() / KB.Float64() }

// MB returns size in megabytes
func (size Size) MB() float64 { return size.Float64() / MB.Float64() }

// GB returns size in gigabytes
func (size Size) GB() float64 { return size.Float64() / GB.Float64() }

// TB returns size in terabytes
func (size Size) TB() float64 { return size.Float64() / TB.Float64() }

// PB returns size in petabytes
func (size Size) PB() float64 { return size.Float64() / PB.Float64() }

// EB returns size in exabytes
func (size Size) EB() float64 { return size.Float64() / EB.Float64() }

// String converts size to a string
func (size Size) String() string {
	if size == 0 {
		return "0"
	}

	switch {
	case size >= EiB*2/3:
		return fmt.Sprintf("%.1f EiB", size.EiB())
	case size >= PiB*2/3:
		return fmt.Sprintf("%.1f PiB", size.PiB())
	case size >= TiB*2/3:
		return fmt.Sprintf("%.1f TiB", size.TiB())
	case size >= GiB*2/3:
		return fmt.Sprintf("%.1f GiB", size.GiB())
	case size >= MiB*2/3:
		return fmt.Sprintf("%.1f MiB", size.MiB())
	case size >= KiB*2/3:
		return fmt.Sprintf("%.1f KiB", size.KiB())
	}

	return strconv.Itoa(size.Int()) + " B"
}

// Base10String converts size to a string
func (size Size) Base10String() string {
	if size == 0 {
		return "0"
	}

	switch {
	case size >= EB*2/3:
		return fmt.Sprintf("%.1f EB", size.EB())
	case size >= PB*2/3:
		return fmt.Sprintf("%.1f PB", size.PB())
	case size >= TB*2/3:
		return fmt.Sprintf("%.1f TB", size.TB())
	case size >= GB*2/3:
		return fmt.Sprintf("%.1f GB", size.GB())
	case size >= MB*2/3:
		return fmt.Sprintf("%.1f MB", size.MB())
	case size >= KB*2/3:
		return fmt.Sprintf("%.1f KB", size.KB())
	}

	return strconv.Itoa(size.Int()) + " B"
}

func isLetter(b byte) bool {
	return ('a' <= b && b <= 'z') || ('A' <= b && b <= 'Z')
}

// Set updates value from string
func (size *Size) Set(s string) error {
	if s == "" {
		return errors.New("empty size")
	}

	p := len(s)
	for isLetter(s[p-1]) {
		p--

		if p < 0 {
			return errors.New("p out of bounds")
		}
	}

	value, suffix := s[:p], s[p:]
	suffix = strings.ToUpper(suffix)
	if suffix == "" || suffix[len(suffix)-1] != 'B' {
		suffix += "B"
	}

	value = strings.TrimSpace(value)
	v, err := strconv.ParseFloat(value, 64)
	if err != nil {
		return err
	}

	switch suffix {
	case "EB":
		*size = Size(v * EB.Float64())
	case "EIB":
		*size = Size(v * EiB.Float64())
	case "PB":
		*size = Size(v * PB.Float64())
	case "PIB":
		*size = Size(v * PiB.Float64())
	case "TB":
		*size = Size(v * TB.Float64())
	case "TIB":
		*size = Size(v * TiB.Float64())
	case "GB":
		*size = Size(v * GB.Float64())
	case "GIB":
		*size = Size(v * GiB.Float64())
	case "MB":
		*size = Size(v * MB.Float64())
	case "MIB":
		*size = Size(v * MiB.Float64())
	case "KB":
		*size = Size(v * KB.Float64())
	case "KIB":
		*size = Size(v * KiB.Float64())
	case "B", "":
		*size = Size(v)
	default:
		return fmt.Errorf("unknown suffix %q", suffix)
	}

	return nil
}

// Type implements pflag.Value
func (Size) Type() string { return "memory.Size" }
