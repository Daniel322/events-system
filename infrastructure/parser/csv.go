package parsers

import (
	"context"
	"encoding/csv"
	"events-system/internal/application/commands"
	"fmt"
	"io"
)

func ParseCsv(ctx context.Context, reader io.ReadCloser) (*[]commands.CreateEventData, error) {
	csvReader := csv.NewReader(reader)
	csvReader.LazyQuotes = true

	for {
		record, err := csvReader.Read()

		if err == io.EOF {
			break
		}
		fmt.Println(record)
	}

	return &[]commands.CreateEventData{}, nil
}
