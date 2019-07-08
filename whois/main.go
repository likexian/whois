/*
 * Copyright 2014-2019 Li Kexian
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
 * Go module for domain and ip whois info query
 * https://www.likexian.com/
 */

package main

import (
	"fmt"
	"github.com/mailgun/whois-go"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println(fmt.Sprintf("usage:\n\t%s domain [server]", os.Args[0]))
		os.Exit(1)
	}

	var server string
	if len(os.Args) > 2 {
		server = os.Args[2]
	}

	result, err := whois.Whois(os.Args[1], server)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	fmt.Println(result)
	os.Exit(0)
}
