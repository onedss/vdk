package format

import (
	"github.com/onedss/vdk/av/avutil"
	"github.com/onedss/vdk/format/aac"
	"github.com/onedss/vdk/format/flv"
	"github.com/onedss/vdk/format/mp4"
	"github.com/onedss/vdk/format/rtmp"
	"github.com/onedss/vdk/format/rtsp"
	"github.com/onedss/vdk/format/ts"
)

func RegisterAll() {
	avutil.DefaultHandlers.Add(mp4.Handler)
	avutil.DefaultHandlers.Add(ts.Handler)
	avutil.DefaultHandlers.Add(rtmp.Handler)
	avutil.DefaultHandlers.Add(rtsp.Handler)
	avutil.DefaultHandlers.Add(flv.Handler)
	avutil.DefaultHandlers.Add(aac.Handler)
}
