package tfl

import (
	"cmp"
	"encoding/json"
	"errors"
	"fmt"
	"ltt/output"
	"net/http"
	"slices"
	"strings"

	"github.com/fatih/color"
)

var reallyBadStatuses = map[string]struct{}{
	"Closed":          {},
	"No Service":      {},
	"Planned Closure": {},
	"Severe Delays":   {},
	"Not Running":     {},
	"Suspended":       {},
}

func lineStatusUrl(lines []string) (string, error) {
	if len(lines) == 0 {
		return "", errors.New("no lines provided")
	}
	return fmt.Sprintf("%s/Line/%s/Status", BaseUrl, strings.Join(lines, ",")), nil
}

type LineStatus struct {
	Severity    uint8   `json:"statusseverity"`
	Description string  `json:"statusSeveritydescription"`
	Reason      *string `json:"reason,omitempty"`
}

func (status LineStatus) severityColor() *color.Color {
	// https://api.tfl.gov.uk/Line/Meta/Severity
	var key string
	if status.Description == "Good Service" {
		key = "green"
	} else if _, ok := reallyBadStatuses[status.Description]; ok {
		key = "red"
	} else {
		key = "yellow"
	}
	rgb, ok := safetyColors[key]
	if ok {
		return rgb.Add(color.Bold)
	} else {
		return nil
	}
}

type Line struct {
	Id       string        `json:"id"`
	Name     string        `json:"name"`
	Mode     string        `json:"modeName"`
	Statuses []*LineStatus `json:"lineStatuses"`
}

func (line Line) mostSevereStatus() (*LineStatus, error) {
	if len(line.Statuses) == 0 {
		return nil, errors.New("no statuses found")
	}
	mostSevere := slices.MinFunc(line.Statuses, func(a, b *LineStatus) int {
		return cmp.Compare(a.Severity, b.Severity)
	})
	return mostSevere, nil
}

func (line Line) lineColor() *color.Color {
	var lineColor *color.Color
	var ok bool
	lineColor, ok = lineColors[line.Id]
	if !ok {
		lineColor, ok = modeColors[line.Mode]
	}
	if ok {
		return lineColor
	} else {
		return nil
	}
}

func (line Line) ToRow(withLineColor bool) (output.FormattedRow, error) {
	lineCol := output.FormattedText{}
	statusCol := output.FormattedText{}
	row := output.FormattedRow{}
	mostSevere, err := line.mostSevereStatus()
	if err != nil {
		return row, err
	}
	lineColor := line.lineColor()
	severityColor := mostSevere.severityColor()
	if lineColor != nil && withLineColor {
		lineCol.Add("    ", lineColor)
		lineCol.Add(" ", nil)
	}
	lineCol.Add(line.Name, nil)
	statusCol.Add(mostSevere.Description, severityColor)
	row.AddCol(lineCol)
	row.AddCol(statusCol)
	return row, nil
}

func GetLineStatuses(lineIds []string) ([]Line, error) {
	url, err := lineStatusUrl(lineIds)
	if err != nil {
		return nil, err
	}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("User-Agent", "ltt")
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	var lines []Line
	if err := json.NewDecoder(resp.Body).Decode(&lines); err != nil {
		return nil, err
	}
	return lines, nil
}
