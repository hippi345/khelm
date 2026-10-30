package config

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestLoaderConfigChartVersion(t *testing.T) {
	require.Equal(t, "1.2.3", (&LoaderConfig{Version: "1.2.3"}).ChartVersion())
	require.Equal(t, "", (&LoaderConfig{}).ChartVersion())
	require.Equal(t, ">0.0.0-0", (&LoaderConfig{Devel: true}).ChartVersion())
	require.Equal(t, "1.0.0", (&LoaderConfig{Version: " 1.0.0 ", Devel: true}).ChartVersion())
}
