package logger_test

import (
	"bytes"
	"encoding/json"
	"github.com/ranog/api-study-in-go/config/logger"
	"os"
	"testing"
	"time"

	"log/slog"
)

func TestInitLogger_StartingAPI(t *testing.T) {
	timeNow := time.Now()
	r, w, _ := os.Pipe()
	originalStdout := os.Stdout
	os.Stdout = w

	logger.InitLogger()
	slog.Info("Starting API", slog.Time("time", timeNow))

	w.Close()
	os.Stdout = originalStdout

	var buf bytes.Buffer
	buf.ReadFrom(r)
	var resultMap map[string]interface{}
	json.Unmarshal([]byte(buf.String()), &resultMap)

	result, _ := json.Marshal(resultMap)

	payload := map[string]interface{}{
		"time":  timeNow.Format("2006-01-02T15:04:05.999999999-07:00"),
		"level": "INFO",
		"msg":   "Starting API",
	}

	expectedResult, _ := json.Marshal(payload)

	if !bytes.Equal(result, expectedResult) {
		t.Fatalf("\nCurrent:   %s\nExpected: %s", expectedResult, result)
	}
}
