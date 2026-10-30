package helm

import (
	"testing"

	chartv2 "helm.sh/helm/v4/pkg/chart/v2"
	repo "helm.sh/helm/v4/pkg/repo/v1"

	"github.com/mgoltzsche/khelm/v2/pkg/config"
	"github.com/stretchr/testify/require"
)

func TestChartVersionDevelSelectsPrerelease(t *testing.T) {
	idx := repo.NewIndexFile()
	idx.Entries = map[string]repo.ChartVersions{
		"mychart": {
			{
				Metadata: &chartv2.Metadata{Name: "mychart", Version: "2.0.0-beta.1"},
				URLs:     []string{"http://example/mychart-2.0.0-beta.1.tgz"},
			},
			{
				Metadata: &chartv2.Metadata{Name: "mychart", Version: "1.0.0"},
				URLs:     []string{"http://example/mychart-1.0.0.tgz"},
			},
		},
	}

	cv, err := idx.Get("mychart", "")
	require.NoError(t, err)
	require.Equal(t, "1.0.0", cv.Version)

	cfg := &config.LoaderConfig{Devel: true}
	cv, err = idx.Get("mychart", cfg.ChartVersion())
	require.NoError(t, err)
	require.Equal(t, "2.0.0-beta.1", cv.Version)
}
