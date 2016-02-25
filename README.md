# icat for JSON with Golang

[![Build Status](https://travis-ci.org/spiegel-im-spiegel/icat4json.svg?branch=master)](https://travis-ci.org/spiegel-im-spiegel/icat4json)
[![GitHub license](https://img.shields.io/badge/license-CC0-blue.svg)](https://raw.githubusercontent.com/spiegel-im-spiegel/icat4json/master/LICENSE)

## Example

```go:sample.go
package main

import (
	"fmt"
	"log"

	"github.com/spiegel-im-spiegel/icat4json"
)

func main() {
	json, err := icat4json.Get(icat4json.ToolICATW)
	if err != nil {
		log.Fatal(err)
	}
	data, err := json.Decode()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Title: %v\n", data.Title)
	fmt.Printf("  URL: %v\n", data.Link)
	fmt.Printf(" Date: %v\n", data.Date)
	fmt.Print("Items:\n")
	for _, item := range data.Itemdata {
		fmt.Printf("\t%v: %v (%v)\n", item.Date, item.Title, item.Link)
	}
}
```

```
$ go run sample.go
Title: IPAセキュリティセンター:重要なセキュリティ情報
  URL: https://www.ipa.go.jp/security/vuln/icat.html
 Date: 2016-02-10 11:58:22 +0900 JST
Items:
	2016-02-10 12:00:00 +0900 JST: Microsoft 製品の脆弱性対策について(2016年02月) (http://www.ipa.go.jp/security/ciadr/vul/20160210-ms.html)
	2016-02-10 12:00:00 +0900 JST: Adobe Flash Player の脆弱性対策について(APSB16-04)(CVE-2016-0985等) (http://www.ipa.go.jp/security/ciadr/vul/20160210-adobeflashplayer.html)
	2016-01-20 12:00:00 +0900 JST: Oracle Java の脆弱性対策について(CVE-2016-0494等) (http://www.ipa.go.jp/security/ciadr/vul/20160120-jre.html)
	2016-01-13 12:00:00 +0900 JST: Microsoft 製品の脆弱性対策について(2016年01月) (http://www.ipa.go.jp/security/ciadr/vul/20160113-ms.html)
	2016-01-13 12:00:00 +0900 JST: Adobe Reader および Acrobat の脆弱性対策について(APSB16-02)(CVE-2016-0932等) (http://www.ipa.go.jp/security/ciadr/vul/20160113-adobereader.html)
	2016-01-06 16:40:00 +0900 JST: 【注意喚起】インターネットに接続する複合機等のオフィス機器の再点検を！ (http://www.ipa.go.jp/security/ciadr/vul/20160106-printer.html)
	2016-01-05 14:00:00 +0900 JST: 「DXライブラリ」におけるバッファオーバーフローの脆弱性対策について(JVN#49476817) (http://www.ipa.go.jp/security/ciadr/vul/20160105-jvn.html)
```

## Reference

- [サイバーセキュリティ注意喚起サービス「icat for JSON」：IPA 独立行政法人 情報処理推進機構](https://www.ipa.go.jp/security/vuln/icat.html)
- [icat for JSON について - Qiita](http://qiita.com/spiegel-im-spiegel/items/4acefe47d3dda688a03e)
- [icat4json 公開 — しっぽのさきっちょ | text.Baldanders.info](http://text.baldanders.info/remark/2016/02/icat4json/)

## License

These codes are licensed under CC0.

[![CC0](http://i.creativecommons.org/p/zero/1.0/88x31.png "CC0")](http://creativecommons.org/publicdomain/zero/1.0/deed)
