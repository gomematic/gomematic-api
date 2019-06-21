package logger

import (
	"database/sql/driver"
	"fmt"
	"reflect"
	"regexp"
	"time"

	"github.com/rs/zerolog/log"
)

var (
	sql = regexp.MustCompile(`(\$\d+)|\?`)
)

// Logger defines a Gorm compatible logger.
type Logger struct{}

// Print is used by Gorm for SQL logging.
func (l *Logger) Print(values ...interface{}) {
	if len(values) > 1 {
		entry := log.With().
			Str("component", "db").
			Str("source", values[1].(string)).
			Logger()

		if values[0] == "sql" {
			var formattedValues []interface{}

			for _, value := range values[4].([]interface{}) {
				indirectValue := reflect.Indirect(reflect.ValueOf(value))

				if indirectValue.IsValid() {
					value = indirectValue.Interface()

					if t, ok := value.(time.Time); ok {
						formattedValues = append(formattedValues, fmt.Sprintf("'%v'", t.Format(time.RFC3339)))
					} else if b, ok := value.([]byte); ok {
						formattedValues = append(formattedValues, fmt.Sprintf("'%v'", string(b)))
					} else if r, ok := value.(driver.Valuer); ok {
						if value, err := r.Value(); err == nil && value != nil {
							formattedValues = append(formattedValues, fmt.Sprintf("'%v'", value))
						} else {
							formattedValues = append(formattedValues, "NULL")
						}
					} else {
						formattedValues = append(formattedValues, fmt.Sprintf("'%v'", value))
					}
				} else {
					formattedValues = append(formattedValues, fmt.Sprintf("'%v'", value))
				}
			}

			entry.Debug().
				Dur("duration", values[2].(time.Duration)).
				Msg(fmt.Sprintf(sql.ReplaceAllString(values[3].(string), "%v"), formattedValues...))

		} else {
			entry.Error().Msg(fmt.Sprintf("%s", values[2:]...))
		}
	} else {
		log.Error().Str("component", "db").Msg(fmt.Sprintf("%s", values...))
	}
}

// New prepares a Gorm compatible logger.
func New() *Logger {
	return &Logger{}
}
