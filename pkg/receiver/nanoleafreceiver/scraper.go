package nanoleafreceiver

import (
	"context"
	"time"

	"github.com/Elfo404/home-otel-collector/pkg/receiver/nanoleafreceiver/internal/metadata"
	"github.com/adnanbrq/nanoleaf"
	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/pdata/pcommon"
	"go.opentelemetry.io/collector/pdata/pmetric"
	"go.opentelemetry.io/collector/receiver"
	"go.opentelemetry.io/collector/receiver/scrapererror"
)

type nanoleafScraper struct {
	settings       component.TelemetrySettings
	mb             *metadata.MetricsBuilder
	cfg            *Config
	nanoleafClient *nanoleaf.Nanoleaf
}

func newNanoleafScraper(settings receiver.CreateSettings, cfg *Config) *nanoleafScraper {

	nanoleafClient := nanoleaf.NewNanoleaf(cfg.Endpoint)
	nanoleafClient.SetToken(cfg.Token)

	return &nanoleafScraper{
		cfg:            cfg,
		settings:       settings.TelemetrySettings,
		mb:             metadata.NewMetricsBuilder(cfg.MetricsBuilderConfig, settings),
		nanoleafClient: nanoleafClient,
	}
}

func (r *nanoleafScraper) scrape(ctx context.Context) (pmetric.Metrics, error) {
	r.settings.Logger.Debug("Scraping...")

	now := pcommon.NewTimestampFromTime(time.Now())
	errs := &scrapererror.ScrapeErrors{}

	r.scrapeStatus(now, errs)
	r.scrapeBrighness(now, errs)
	r.scrapeColorMode(now, errs)
	r.scrapeSaturation(now, errs)
	r.scrapeHue(now, errs)
	r.scrapeColorTemperature(now, errs)

	return r.mb.Emit(), errs.Combine()
}

func (r *nanoleafScraper) scrapeBrighness(now pcommon.Timestamp, errs *scrapererror.ScrapeErrors) {
	brighness, err := r.nanoleafClient.State.GetBrightness()

	if err != nil {
		errs.Add(err)
		return
	}
	r.mb.RecordBrightnessDataPoint(now, float64(brighness.Value)/float64(brighness.Max))
}

func (r *nanoleafScraper) scrapeSaturation(now pcommon.Timestamp, errs *scrapererror.ScrapeErrors) {
	saturation, err := r.nanoleafClient.State.GetSaturation()

	if err != nil {
		errs.Add(err)
		return
	}
	r.mb.RecordSaturationDataPoint(now, float64(saturation.Value)/float64(saturation.Max))
}

func (r *nanoleafScraper) scrapeStatus(now pcommon.Timestamp, errs *scrapererror.ScrapeErrors) {
	isOn, err := r.nanoleafClient.State.IsOn()

	value := 0
	if isOn {
		value = 1
	}

	if err != nil {
		errs.Add(err)
		return
	}
	r.mb.RecordStatusDataPoint(now, int64(value))
}

func (r *nanoleafScraper) scrapeHue(now pcommon.Timestamp, errs *scrapererror.ScrapeErrors) {
	hue, err := r.nanoleafClient.State.GetHue()

	if err != nil {
		errs.Add(err)
		return
	}
	r.mb.RecordHueDataPoint(now, int64(hue.Value))

}

func (r *nanoleafScraper) scrapeColorTemperature(now pcommon.Timestamp, errs *scrapererror.ScrapeErrors) {
	temp, err := r.nanoleafClient.State.GetColorTemp()

	if err != nil {
		errs.Add(err)
		return
	}
	r.mb.RecordColorTemperatureDataPoint(now, int64(temp.Value))
}

type ColorMode int64

const (
	colorModeHueSaturation    ColorMode = 0
	colorModeColorTemperature           = 1
	colorModeEffect                     = 2
)

var colorModeMap = map[string]ColorMode{
	"\"hs\"":     colorModeHueSaturation,
	"\"ct\"":     colorModeColorTemperature,
	"\"effect\"": colorModeEffect,
}

func (r *nanoleafScraper) scrapeColorMode(now pcommon.Timestamp, errs *scrapererror.ScrapeErrors) {
	mode, err := r.nanoleafClient.State.GetColorMode()
	if err != nil {
		errs.Add(err)
		return
	}

	r.mb.RecordColorModeDataPoint(now, int64(colorModeMap[mode]))
}
