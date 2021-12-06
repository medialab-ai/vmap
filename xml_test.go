package vmap

import (
	"encoding/xml"
	"strings"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestParsingSimple(t *testing.T) {
	t.Parallel()
	is := assert.New(t)
	vmap := &VMAP{}
	is.Nil(xml.Unmarshal([]byte(vmapPostRollBumper), vmap))
	if !is.NotNil(vmap) {
		t.FailNow()
	}

	// this is 2 because the number of ad breaks is 2
	is.Equal(2, len(vmap.AdBreak))
	adBreak := vmap.AdBreak[0]
	is.Equal(adBreak.TimeOffset.position, OffsetPositionEnd)
	is.Equal("postroll", adBreak.BreakID)
	is.NotEmpty(adBreak.Extensions.Extension)
	source := adBreak.AdSource
	is.Equal("postroll-pre-bumper", source.ID)
	is.Equal(false, source.AllowMultipleAds)
	is.Equal(true, source.FollowRedirects)
	tag := source.AdTagURI
	if !is.NotNil(tag) {
		t.FailNow()
	}

	is.Equal("vast3", tag.TemplateType)
	is.True(strings.Contains(tag.Text, "doubleclick"))
}

func TestOffSetParsing(t *testing.T) {
	dur := Duration(time.Hour*1 + time.Minute*2 + time.Second*3)
	cases := []struct {
		label    string
		input    string
		expected *Offset
	}{
		{"position_start", "start", &Offset{position: OffsetPositionStart}},
		{"position_end", "end", &Offset{position: OffsetPositionEnd}},
		{"percent", "23%", &Offset{percent: 23}},
		{"time_duration", "01:02:03", &Offset{timeDur: dur}},
	}

	for _, c := range cases {
		t.Run(c.label, func(t2 *testing.T) {
			is := assert.New(t)
			off := &Offset{}
			off.UnmarshalText([]byte(c.input))
			is.EqualValues(c.expected, off)
		})
	}
}

func TestMarshal(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	vmap1 := &VMAP{}
	vmap2 := &VMAP{}
	is.Nil(xml.Unmarshal([]byte(vmapPostRollBumper), vmap1))
	is.NotEmpty(vmap1)
	bytes, err := xml.Marshal(vmap1)
	is.Nil(err)
	is.NotEmpty(bytes)
	is.Nil(xml.Unmarshal(bytes, vmap2))
	is.EqualValues(vmap1, vmap2)
}

const (
	// vmapPostRollBumper taken from https://developers.google.com/interactive-media-ads/docs/sdks/html5/client-side/tags
	vmapPostRollBumper = `<vmap:VMAP xmlns:vmap="http://www.iab.net/videosuite/vmap" version="1.0">
	<vmap:AdBreak timeOffset="end" breakType="linear" breakId="postroll">
	<vmap:AdSource id="postroll-pre-bumper" allowMultipleAds="false" followRedirects="true">
	<vmap:AdTagURI templateType="vast3">
	<![CDATA[ https://pubads.g.doubleclick.net/gampad/ads?slotname=/124319096/external/ad_rule_samples&sz=640x480&ciu_szs=300x250&cust_params=deployment%3Ddevsite%26sample_ar%3Dpostonlybumper&url=&unviewed_position_start=1&output=xml_vast3&impl=s&env=vp&gdfp_req=1&ad_rule=0&useragent=Mozilla/5.0+(Macintosh%3B+Intel+Mac+OS+X+10_15_7)+AppleWebKit/537.36+(KHTML,+like+Gecko)+Chrome/96.0.4664.55+Safari/537.36,gzip(gfe)&vad_type=linear&vpos=postroll&pod=1&bumper=before&min_ad_duration=0&max_ad_duration=10000&vrid=6136&sb=1&cmsid=496&video_doc_id=short_onecue&kfa=0&tfcd=0 ]]>
	</vmap:AdTagURI>
	</vmap:AdSource>
	<vmap:Extensions>
	<vmap:Extension type="bumper" suppress_bumper="true"/>
	</vmap:Extensions>
	</vmap:AdBreak>
	<vmap:AdBreak timeOffset="end" breakType="linear" breakId="postroll">
	<vmap:AdSource id="postroll-ad-1" allowMultipleAds="false" followRedirects="true">
	<vmap:AdTagURI templateType="vast3">
	<![CDATA[ https://pubads.g.doubleclick.net/gampad/ads?slotname=/124319096/external/ad_rule_samples&sz=640x480&ciu_szs=300x250&cust_params=deployment%3Ddevsite%26sample_ar%3Dpostonlybumper&url=&unviewed_position_start=1&output=xml_vast3&impl=s&env=vp&gdfp_req=1&ad_rule=0&useragent=Mozilla/5.0+(Macintosh%3B+Intel+Mac+OS+X+10_15_7)+AppleWebKit/537.36+(KHTML,+like+Gecko)+Chrome/96.0.4664.55+Safari/537.36,gzip(gfe)&vad_type=linear&vpos=postroll&pod=1&ppos=1&lip=true&min_ad_duration=0&max_ad_duration=30000&vrid=6136&cmsid=496&video_doc_id=short_onecue&kfa=0&tfcd=0 ]]>
	</vmap:AdTagURI>
	</vmap:AdSource>
	</vmap:AdBreak>
	</vmap:VMAP>`
)
