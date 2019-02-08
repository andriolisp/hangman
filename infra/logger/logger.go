package logger

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"io"

	"github.com/Sirupsen/logrus"
	"github.com/andriolisp/hangman/infra/config"
)

// Logger is the default application logger compatible with the echo.Logger interface
type Logger struct {
	cfg *config.Config
	*logrus.Entry
}

// InfoWriter returns the io.Writer for info level
func (l *Logger) InfoWriter() io.Writer {
	return l.WriterLevel(logrus.InfoLevel)
}

// ErrorWriter returns the io.Writer for error level
func (l *Logger) ErrorWriter() io.Writer {
	return l.WriterLevel(logrus.ErrorLevel)
}

type statusWriter struct {
	http.ResponseWriter
	status int
	length int
}

func (w *statusWriter) WriteHeader(status int) {
	w.status = status
	w.ResponseWriter.WriteHeader(status)
}

func (w *statusWriter) Write(b []byte) (int, error) {
	if w.status == 0 {
		w.status = 200
	}
	n, err := w.ResponseWriter.Write(b)
	w.length += n
	return n, err
}

// Middleware will add intercept all requests and log
func (l *Logger) Middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		sw := statusWriter{ResponseWriter: w}

		next.ServeHTTP(w, r)

		duration := time.Now().Sub(start)
		fields := logrus.Fields{}
		fields["host"] = r.Host
		fields["remote_addr"] = r.RemoteAddr
		fields["remote_uri"] = r.RequestURI
		fields["proto"] = r.Proto
		fields["status"] = sw.status
		fields["length"] = sw.length
		fields["method"] = r.Method
		fields["user_agent"] = r.Header.Get("User-Agent")
		fields["duration"] = duration

		l.WithFields(fields).Print(r.URL.String())
	})
}

// New returns a new Logger instance
func New(cfg *config.Config) (*Logger, error) {
	if cfg.Log.LogToFile {
		file, err := os.Create(cfg.Log.Path)
		if err != nil {
			fmt.Printf("Error to create log file for library: %s\n", err.Error())
			panic(err)
		}
		logrus.SetOutput(file)
	}

	logrus.SetFormatter(&logrus.JSONFormatter{})

	if cfg.App.Debug {
		logrus.SetLevel(logrus.DebugLevel)
	} else {
		logrus.SetLevel(logrus.InfoLevel)
	}

	entry := logrus.WithFields(logrus.Fields{
		"app":   cfg.App.Name,
		"debug": cfg.App.Debug,
	})

	return &Logger{cfg, entry}, nil
}
