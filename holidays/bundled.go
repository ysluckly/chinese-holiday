//go:generate statik -src=./data
package holidays

import (
	"encoding/json"
	"os"

	"github.com/rakyll/statik/fs"
	_ "github.com/ysluckly/chinese-holiday/statik" // load data
)

func loadData() ([]event, error) {
	statikFS, err := fs.New()
	if err != nil {
		return nil, err
	}
	var events []event
	err = fs.Walk(statikFS, "/", func(path string, info os.FileInfo, err error) error {
		if path == "/" {
			return err
		}
		b, err := fs.ReadFile(statikFS, path)
		if err != nil {
			return nil
		}
		e, err := parseEvents(b)
		events = append(events, e...)

		return err
	})
	if err != nil {
		return nil, err
	}
	return events, nil
}

func parseEvents(b []byte) ([]event, error) {
	var events []event
	err := json.Unmarshal(b, &events)
	if err != nil {
		return nil, err
	}

	return events, nil
}
