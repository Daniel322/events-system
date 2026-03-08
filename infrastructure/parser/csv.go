package parsers

import (
	"context"
	"encoding/csv"
	"events-system/internal/application/commands"
	"events-system/pkg/utils"
	"io"
	"strings"
	"time"
)

type ParseOptions struct {
	UserId string
	AccId  string
}

func ParseCsv(ctx context.Context, reader io.ReadCloser, options ParseOptions) (*[]commands.CreateEventData, error) {
	csvReader := csv.NewReader(reader)
	csvReader.LazyQuotes = true

	result := make([]commands.CreateEventData, 0)

	for {
		record, err := csvReader.Read()

		if err == io.EOF {
			break
		}

		rows := strings.Split(record[0], "\n")

		for index, row := range rows {
			if index == 0 {
				continue
			}

			rowParts := strings.Split(row, ";")

			date, err := time.Parse("\"2006-01-02\"", rowParts[1])

			if err != nil {
				return nil, utils.GenerateError("CSVParser", err.Error())
			}

			result = append(
				result,
				commands.CreateEventData{
					Info:         rowParts[0],
					Date:         date,
					Providers:    []string{"telegram"},
					NotifyLevels: []string{"today", "tomorrow", "week", "month"},
					UserId:       options.UserId,
					AccId:        options.AccId,
				})
		}
	}

	return &result, nil
}
