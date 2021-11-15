package cht

import (
	_ "embed"
	"encoding/xml"
	"os"
	"os/exec"

	"github.com/go-faster/errors"
)

type Logger struct {
	Level   string `xml:"level"`
	Console int    `xml:"console,omitempty"`
}

type UsersXML struct {
	Path string `xml:"path"`
}

type UserDir struct {
	UsersXML UsersXML `xml:"users_xml"`
}

type Config struct {
	XMLName xml.Name `xml:"clickhouse"`
	Logger  Logger   `xml:"logger"`
	HTTP    int      `xml:"http_port"`
	TCP     int              `xml:"tcp_port"`
	Host    string           `xml:"host"`

	Path            string                    `xml:"path"`
	TempPath        string                    `xml:"tmp_path"`
	UserFilesPath   string  `xml:"user_files_path"`
	UserDirectories UserDir `xml:"user_directories"`

	MarkCacheSize int `xml:"mark_cache_size"`
	MMAPCacheSize int `xml:"mmap_cache_size"`
}

// EnvBin is environmental variable that sets paths to current
// ClickHouse binary.
const EnvBin = "CLICKHOUSE_BIN"

//go:embed clickhouse.users.xml
var usersCfg []byte

// Bin returns path to current ClickHouse binary.
func Bin() (string, error) {
	v, ok := os.LookupEnv(EnvBin)
	if !ok {
		// Fallback to default binary name.
		// Should be in $PATH.
		v = "clickhouse-server"
	}
	p, err := exec.LookPath(v)
	if err != nil {
		return "", errors.Wrap(err, "lookup")
	}
	return p, nil
}

