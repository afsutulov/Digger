// SPDX-License-Identifier: GPL-3.0-only
package game

import (
	_ "embed"
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

//go:embed assets/DIGGER.HSC
var initialHSC []byte

func decodeHSC(b []byte) ([]scoreRecord, error) {
	if len(b) != 512 || b[0] != 's' {
		return nil, fmt.Errorf("invalid DIGGER.HSC")
	}
	a := make([]scoreRecord, 10)
	for i := range a {
		row := b[1+i*11 : 1+(i+1)*11]
		n, err := strconv.Atoi(strings.TrimSpace(string(row[3:])))
		if err != nil || n < 0 || n > 999999 {
			return nil, fmt.Errorf("invalid HSC row %d", i+1)
		}
		a[i] = scoreRecord{cleanInitials(string(row[:3])), n}
	}
	return a, nil
}

// ImportHSC converts the original table without executing the DOS program.
func ImportHSC(source, destination string) error {
	b, err := os.ReadFile(source)
	if err != nil {
		return err
	}
	if len(b) != 512 || b[0] != 's' {
		return fmt.Errorf("invalid DIGGER.HSC: expected 512 bytes and marker s")
	}
	a := make([]scoreRecord, 10)
	for i := range a {
		row := b[1+i*11 : 1+(i+1)*11]
		n, err := strconv.Atoi(strings.TrimSpace(string(row[3:])))
		if err != nil || n < 0 || n > 999999 {
			return fmt.Errorf("invalid score in HSC row %d", i+1)
		}
		a[i] = scoreRecord{cleanInitials(string(row[:3])), n}
	}
	d, err := json.MarshalIndent(a, "", "  ")
	if err != nil {
		return err
	}
	if err = os.MkdirAll(filepath.Dir(destination), 0755); err != nil {
		return err
	}
	return os.WriteFile(destination, d, 0644)
}
