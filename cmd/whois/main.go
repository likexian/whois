/*
 * Copyright 2014-2023 Li Kexian
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 *
 * Go module for domain and ip whois information query
 * https://www.likexian.com/
 */

package main

import (
	"context"
	"flag"
	"fmt"
	"os"
	"time"

	"github.com/likexian/gokit/xjson"
	"github.com/likexian/gokit/xversion"
	"github.com/likexian/whois"
	whoisparser "github.com/likexian/whois-parser"
	"golang.org/x/net/proxy"
)

func main() {
	updateMessage := make(chan string)
	go checkUpdate(updateMessage, "v"+whois.Version())

	server := flag.String("h", "", "specify the whois server")
	outJSON := flag.Bool("j", false, "output format as json")
	version := flag.Bool("v", false, "show the whois version")
	flag.Parse()

	if *version {
		fmt.Println("whois version " + whois.Version())
		fmt.Println(whois.Author())
		os.Exit(0)
	}

	if len(flag.Args()) == 0 {
		fmt.Printf("Usage:\n\t%s [-j] [-h server] domain\n", os.Args[0])
		fmt.Printf(`
domain:
  domain or IPv4 or IPv6 or ASN for query

options:
  -h string specify the whois server
  -j        output format as json
  -v        show the whois version
`)
		os.Exit(1)
	}

	text, err := whois.NewClient().
		SetDialer(proxy.FromEnvironment()).Whois(flag.Args()[0], *server)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	if *outJSON {
		info, err := whoisparser.Parse(text)
		if err != nil {
			fmt.Println(err.Error())
			os.Exit(1)
		}
		data, err := xjson.PrettyDumps(info)
		if err != nil {
			fmt.Println(err.Error())
			os.Exit(1)
		}
		fmt.Println(data)
		os.Exit(0)
	}

	fmt.Println(text)

	message := <-updateMessage
	if message != "" {
		fmt.Println(message)
	}

	os.Exit(0)
}

// checkUpdate checks version update
func checkUpdate(updateMessage chan string, version string) {
	checkPoint := "https://release.likexian.com/whois/update"
	cacheFile := fmt.Sprintf("%s/whois.update.cache", os.TempDir())

	req := &xversion.CheckUpdateRequest{
		Product:       "whois",
		Current:       version,
		CacheFile:     cacheFile,
		CacheDuration: 24 * time.Hour,
		CheckPoint:    checkPoint,
	}

	ctx := context.Background()
	rsp, err := req.Run(ctx)
	if err == nil && rsp.Outdated {
		if version != rsp.Current {
			_ = os.Remove(cacheFile)
		} else {
			emergency := "NOTICE"
			if rsp.Emergency {
				emergency = "WARNING"
			}
			message := fmt.Sprintf("%% %s: Your version of whois is outdate, the latest is %s.\n",
				emergency, rsp.Latest)
			message += fmt.Sprintf("%% You can update it by downloading from %s", rsp.ProductURL)
			updateMessage <- message
		}
	}

	updateMessage <- ""
}
