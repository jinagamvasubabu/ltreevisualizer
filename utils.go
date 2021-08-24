package ltreevisualizer

import (
	"context"
	log "github.com/sirupsen/logrus"
	"strings"
	"time"
)

//CalculateTimeTaken calculates the time taken to complete the execution of a method
func CalculateTimeTaken(ctx context.Context, start time.Time, name string) {
	logger := log.WithContext(ctx).WithFields(log.Fields{"Method": "CalculateTimeTaken"})
	elapsed := time.Since(start)
	logger.Debugf("%s took %s", name, elapsed)
}

//Contains Check if a search term is available in a slice, returns bool
func Contains(list []string, searchTerm string) bool {
	for _, s := range list {
		if strings.ToLower(s) == strings.ToLower(searchTerm) {
			return true
		}
	}
	return false
}
