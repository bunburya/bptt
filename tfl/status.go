package tfl

import (
	"cmp"
	"errors"
	"fmt"
	"ptt/output"
	"slices"
	"strings"

	"github.com/fatih/color"
)

// The TfL API assigns numerical codes to each status severity, and it seems like you could *mostly* get by just
// treating lower numbers as being more severe, but that may not necessarily work in all cases. So below is a list of
// most severity descriptions observed at https://api.tfl.gov.uk/Line/Meta/Severity (excluding some that seem clearly
// only intended for use with stations, rather than lines), ordered roughly in the order of severity.
var severityOrder = []string{
	"Closed",
	"No Service",
	"Not Running",
	"Planned Closure",
	"Suspended",
	"Part Closure",
	"Part Closed",
	"Part Suspended",
	"Severe Delays",
	// Special Service is used differently on different lines and can mean anything from minor delays to suspended.
	"Special Service",
	"Reduced Service",
	"Bus Service",
	"Change of frequency",
	"Diverted",
	"Issues Reported",
	"Minor Delays",
	"Information",
	"No Issues",
	"Good Service",
}

func lineStatusUrl(lines []string) (string, error) {
	if len(lines) == 0 {
		return "", errors.New("no lines provided")
	}
	return fmt.Sprintf("%s/Line/%s/Status", BaseUrl, strings.Join(lines, ",")), nil
}

func modeStatusUrl(modes []string) (string, error) {
	if len(modes) == 0 {
		return "", errors.New("no modes provided")
	}
	return fmt.Sprintf("%s/Line/Mode/%s/Status", BaseUrl, strings.Join(modes, ",")), nil
}

type lineStatus struct {
	Description string  `json:"statusSeverityDescription"`
	Reason      *string `json:"reason,omitempty"`

	// `severity` is the internal value that we assign to the status, based on the position of the description in the
	// `severityOrder` slice. It is different to the numerical value assigned by the TfL API. `severityInit` describes
	// whether we have already calculated and cached the result.
	severity     int
	severityInit bool
}

func (status *lineStatus) Severity() int {
	if !status.severityInit {
		status.severity = slices.Index(severityOrder, status.Description)
		status.severityInit = true
	}
	return status.severity
}

func (status *lineStatus) severityColor() *color.Color {
	// https://api.tfl.gov.uk/Line/Meta/Severity
	var key string
	s := status.Severity()
	if s <= 8 {
		key = "red"
	} else if s <= 16 {
		key = "yellow"
	} else {
		key = "green"
	}
	rgb, ok := safetyColors[key]
	if ok {
		return rgb.Add(color.Bold)
	} else {
		return nil
	}
}

// lineWithStatuses represents a single TfL line/route, with the currently applicable statuses.
type lineWithStatuses struct {
	Id       string        `json:"id"`
	Name     string        `json:"name"`
	Mode     string        `json:"modeName"`
	Statuses []*lineStatus `json:"lineStatuses"`
}

func (line *lineWithStatuses) mostSevereStatus() (*lineStatus, error) {
	if len(line.Statuses) == 0 {
		return nil, errors.New("no statuses found")
	}
	mostSevere := slices.MinFunc(line.Statuses, func(a, b *lineStatus) int {
		return cmp.Compare(a.Severity(), b.Severity())
	})
	return mostSevere, nil
}

func (line *lineWithStatuses) lineColor() *color.Color {
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

func (line *lineWithStatuses) toRow(withColor bool) (output.Row, error) {
	lineCell := output.Cell{}
	statusCell := output.Cell{}
	row := output.Row{}
	mostSevere, err := line.mostSevereStatus()
	if err != nil {
		return row, err
	}
	lineColor := line.lineColor()
	severityColor := mostSevere.severityColor()
	if lineColor != nil && withColor {
		lineCell.AddText(" ", lineColor)
		lineCell.AddText(" ", nil)
	}
	lineCell.AddText(line.Name, nil)
	statusCell.AddText(mostSevere.Description, severityColor)
	row.AddCell(lineCell)
	row.AddCell(statusCell)
	return row, nil
}

func getLineStatuses(url string, apiKey string) ([]lineWithStatuses, error) {
	lines, err := request[[]lineWithStatuses](url, apiKey)
	if err != nil {
		return nil, err
	}
	return lines, nil
}

func statusTable(
	url string,
	apiKey string,
	options output.Options,
) (output.Table, error) {
	table := output.Table{}
	lines, err := getLineStatuses(url, apiKey)
	if err != nil {
		return table, err
	}
	if options.Header {
		table.SetHeader(output.NewRow(
			output.NewCell("Line", color.New(color.Bold)),
			output.NewCell("Status", color.New(color.Bold)),
		))
	}
	for _, line := range lines {
		row, err := line.toRow(options.Color)
		if err != nil {
			return table, err
		}
		table.AddRow(row)
	}
	if options.Timestamp {
		table.Timestamp()
	}
	return table, nil
}

func LineStatusTable(lines []string, apiKey string, options output.Options) (output.Table, error) {
	url, err := lineStatusUrl(lines)
	if err != nil {
		return output.Table{}, err
	}
	return statusTable(url, apiKey, options)
}

func ModeStatusTable(modes []string, apiKey string, options output.Options) (output.Table, error) {
	url, err := modeStatusUrl(modes)
	if err != nil {
		return output.Table{}, err
	}
	return statusTable(url, apiKey, options)
}
