// Copyright Istio Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package model

import (
	"errors"

	"istio.io/istio/pkg/ledger"
)

// DisabledLedger is an empty mock of the ledger.Ledger interface
// which we will substitute when distribution tracking is disabled.
type DisabledLedger struct {
	ledger.Ledger
}

func (d *DisabledLedger) Put(key, value string) (string, error) {
	return "", nil
}

func (d *DisabledLedger) Delete(key string) error {
	return nil
}

func (d *DisabledLedger) Get(key string) (string, error) {
	return "", nil
}

func (d *DisabledLedger) RootHash() string {
	return ""
}

func (d *DisabledLedger) GetPreviousValue(previousHash, key string) (result string, err error) {
	return "", errors.New("distribution tracking is disabled")
}
