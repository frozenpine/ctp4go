package ctp4go

import (
	"errors"
	"io"
	"log/slog"
	"math"
	"os"
	"path/filepath"
	"time"

	"gopkg.in/natefinch/lumberjack.v2"
)

const (
	DBL_MAX = math.MaxFloat64

	DEFAULT_LOG_ROTATE_MB   = 300
	DEFAULT_LOG_ROTATE_KEEP = 7
)

var CST *time.Location

type logCfg struct {
	dst     []io.Writer
	verbose int
	jsonFmt bool
}

type logOpt func(*logCfg) error

type LogOptions []logOpt

type logFileCfg struct {
	size, age, keep int
	utc, nocompress bool
}

type logFileOpt func(*logFileCfg)

func WithLogFileSize(size int) logFileOpt {
	return func(lfc *logFileCfg) {
		lfc.size = size
	}
}

func WithLogFileAge(age int) logFileOpt {
	return func(lfc *logFileCfg) {
		lfc.age = age
	}
}

func WithLogFileKeep(keep int) logFileOpt {
	return func(lfc *logFileCfg) {
		lfc.keep = keep
	}
}

func WithLogFileUTC() logFileOpt {
	return func(lfc *logFileCfg) {
		lfc.utc = true
	}
}

func WithLogFileNoCompress() logFileOpt {
	return func(lfc *logFileCfg) {
		lfc.nocompress = true
	}
}

func WithLogFile(filePath string, options ...logFileOpt) logOpt {
	return func(lc *logCfg) error {
		if filePath == "" {
			return errors.New("invalid log file path")
		}

		logDir := filepath.Dir(filePath)
		if info, err := os.Stat(logDir); err != nil {
			if errors.Is(err, os.ErrNotExist) {
				err = os.MkdirAll(logDir, os.ModePerm)
			}

			if err != nil {
				return err
			}
		} else if !info.IsDir() {
			return errors.New("invalid log file path")
		}

		cfg := logFileCfg{
			size: DEFAULT_LOG_ROTATE_MB,
			age:  DEFAULT_LOG_ROTATE_KEEP,
		}

		for _, opt := range options {
			if opt == nil {
				continue
			}

			opt(&cfg)
		}

		lc.dst = append(lc.dst, &lumberjack.Logger{
			Filename:   filePath,
			MaxSize:    cfg.size,
			MaxAge:     cfg.age,
			MaxBackups: cfg.keep,
			LocalTime:  !cfg.utc,
			Compress:   !cfg.nocompress,
		})

		return nil
	}
}

func WithLogConsole() logOpt {
	return func(lc *logCfg) error {
		lc.dst = append(lc.dst, os.Stderr)
		return nil
	}
}

func WithLogVerbose(v int) logOpt {
	return func(lc *logCfg) error {
		lc.verbose = v
		return nil
	}
}

func WithLogJson() logOpt {
	return func(lc *logCfg) error {
		lc.jsonFmt = true
		return nil
	}
}

func NewLogger(
	name string, options ...logOpt,
) (*slog.Logger, error) {
	var (
		logWr     io.Writer
		level     slog.Level = slog.LevelInfo
		addSource            = false

		logger *slog.Logger
	)

	var cfg logCfg

	for _, opt := range options {
		if opt == nil {
			continue
		}

		if err := opt(&cfg); err != nil {
			return nil, err
		}
	}

	if len(cfg.dst) < 1 {
		return nil, errors.New("no log dst specified")
	}

	if len(cfg.dst) > 1 {
		logWr = io.MultiWriter(cfg.dst...)
	} else {
		logWr = cfg.dst[0]
	}

	if cfg.verbose > 0 {
		if cfg.verbose > 1 {
			addSource = true
		}

		level = slog.LevelDebug - slog.Level(cfg.verbose-1)
	}

	if cfg.jsonFmt {
		logger = slog.New(slog.NewJSONHandler(
			logWr, &slog.HandlerOptions{
				AddSource: addSource,
				Level:     level,
			},
		))
	} else {
		logger = slog.New(slog.NewTextHandler(
			logWr, &slog.HandlerOptions{
				AddSource: addSource,
				Level:     level,
			},
		))
	}

	return logger.WithGroup(name), nil
}

func init() {
	var err error

	CST, err = time.LoadLocation("Asia/Shanghai")
	if err != nil {
		CST = time.FixedZone("Asia/Shanghai", 8*3600)
	}
}
