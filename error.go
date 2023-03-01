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

package whois

import "errors"

var (
	// ErrDomainEmpty is domain is an empty string
	ErrDomainEmpty = errors.New("whois: domain is empty")

	// ErrWhoisServerNotFound is no whois server found
	ErrWhoisServerNotFound = errors.New("whois: no whois server found for domain")
)
