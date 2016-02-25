package icat4json

import (
	"fmt"
	"testing"
)

func TestGetW(t *testing.T) {
	json, err := Get(ToolICATW)
	if err != nil {
		t.Errorf("Get() error = \"%v\", want nil.", err)
	} else if json == "" {
		t.Errorf("Get() = \"\", not want \"\".")
	}
}

func TestGetH(t *testing.T) {
	json, err := Get(ToolICATH)
	if err != nil {
		t.Errorf("Get() error = \"%v\", want nil.", err)
	} else if json == "" {
		t.Errorf("Get() = \"\", not want \"\".")
	}
}

func TestGetError(t *testing.T) {
	json, err := Get("aaa")
	if err == nil {
		t.Errorf("Get() error = nil, not want nil.")
		fmt.Println(json)
	} else {
		fmt.Println(err)
	}
}

func TestDecode(t *testing.T) {
	json := JSON(`{"itemdata":[{"item_title":"Microsoft \u88fd\u54c1\u306e\u8106\u5f31\u6027\u5bfe\u7b56\u306b\u3064\u3044\u3066(2016\u5e7402\u6708)","item_link":"http:\/\/www.ipa.go.jp\/security\/ciadr\/vul\/20160210-ms.html","item_date":"2016-02-10T12:00:00+09:00","item_identifier":[]},{"item_title":"Adobe Flash Player \u306e\u8106\u5f31\u6027\u5bfe\u7b56\u306b\u3064\u3044\u3066(APSB16-04)(CVE-2016-0985\u7b49)","item_link":"http:\/\/www.ipa.go.jp\/security\/ciadr\/vul\/20160210-adobeflashplayer.html","item_date":"2016-02-10T12:00:00+09:00","item_identifier":[]},{"item_title":"Oracle Java \u306e\u8106\u5f31\u6027\u5bfe\u7b56\u306b\u3064\u3044\u3066(CVE-2016-0494\u7b49)","item_link":"http:\/\/www.ipa.go.jp\/security\/ciadr\/vul\/20160120-jre.html","item_date":"2016-01-20T12:00:00+09:00","item_identifier":[]},{"item_title":"Microsoft \u88fd\u54c1\u306e\u8106\u5f31\u6027\u5bfe\u7b56\u306b\u3064\u3044\u3066(2016\u5e7401\u6708)","item_link":"http:\/\/www.ipa.go.jp\/security\/ciadr\/vul\/20160113-ms.html","item_date":"2016-01-13T12:00:00+09:00","item_identifier":[]},{"item_title":"Adobe Reader \u304a\u3088\u3073 Acrobat \u306e\u8106\u5f31\u6027\u5bfe\u7b56\u306b\u3064\u3044\u3066(APSB16-02)(CVE-2016-0932\u7b49)","item_link":"http:\/\/www.ipa.go.jp\/security\/ciadr\/vul\/20160113-adobereader.html","item_date":"2016-01-13T12:00:00+09:00","item_identifier":[]},{"item_title":"\u3010\u6ce8\u610f\u559a\u8d77\u3011\u30a4\u30f3\u30bf\u30fc\u30cd\u30c3\u30c8\u306b\u63a5\u7d9a\u3059\u308b\u8907\u5408\u6a5f\u7b49\u306e\u30aa\u30d5\u30a3\u30b9\u6a5f\u5668\u306e\u518d\u70b9\u691c\u3092\uff01","item_link":"http:\/\/www.ipa.go.jp\/security\/ciadr\/vul\/20160106-printer.html","item_date":"2016-01-06T16:40:00+09:00","item_identifier":[]},{"item_title":"\u300cDX\u30e9\u30a4\u30d6\u30e9\u30ea\u300d\u306b\u304a\u3051\u308b\u30d0\u30c3\u30d5\u30a1\u30aa\u30fc\u30d0\u30fc\u30d5\u30ed\u30fc\u306e\u8106\u5f31\u6027\u5bfe\u7b56\u306b\u3064\u3044\u3066(JVN#49476817)","item_link":"http:\/\/www.ipa.go.jp\/security\/ciadr\/vul\/20160105-jvn.html","item_date":"2016-01-05T14:00:00+09:00","item_identifier":[]}],"docTitle":"IPA\u30bb\u30ad\u30e5\u30ea\u30c6\u30a3\u30bb\u30f3\u30bf\u30fc:\u91cd\u8981\u306a\u30bb\u30ad\u30e5\u30ea\u30c6\u30a3\u60c5\u5831","docTitleFix":null,"docLink":"https:\/\/www.ipa.go.jp\/security\/vuln\/icat.html","docDate":"2016-02-10T11:58:22+09:00"}`)
	data, err := json.Decode()
	if err != nil {
		t.Errorf("Decode() error = \"%v\", want nil.", err)
	} else if data == nil {
		t.Errorf("Decode() = nil, not want nil.")
	} else {
		fmt.Println(data)
	}
}

func TestDecodeError(t *testing.T) {
	json := JSON(`{"messageid":"IRSS-0103","message":"\u30c7\u30fc\u30bf\u306e\u8aad\u8fbc\u307f\u306b\u5931\u6557\u3057\u307e\u3057\u305f\u3002\u3057\u3070\u3089\u304f\u3057\u3066\u304b\u3089\u3001\u30da\u30fc\u30b8\u306e\u518d\u8aad\u8fbc\u307f\u3092\u3057\u3066\u304f\u3060\u3055\u3044\u3002"}`)
	_, err := json.Decode()
	if err == nil {
		t.Errorf("Decode() error = nil, not want nil.")
	} else {
		fmt.Println(err)
	}
}

func TestDecodeNil(t *testing.T) {
	json := JSON("")
	_, err := json.Decode()
	if err == nil {
		t.Errorf("Decode() error = nil, not want nil.")
	} else {
		fmt.Println(err)
	}
}
