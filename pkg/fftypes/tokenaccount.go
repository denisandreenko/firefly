// Copyright © 2021 Kaleido, Inc.
//
// SPDX-License-Identifier: Apache-2.0
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

package fftypes

import (
	"crypto/sha256"
	"encoding/json"
)

type TokenAccountIdentifier struct {
	Namespace  string `json:"namespace"`
	PoolID     *UUID  `json:"poolId,omitempty"`
	TokenIndex string `json:"tokenIndex,omitempty"`
	Identity   string `json:"identity,omitempty"`
}

func (t *TokenAccountIdentifier) Hash() *Bytes32 {
	b, _ := json.Marshal(&t)
	var b32 Bytes32 = sha256.Sum256(b)
	return &b32
}

type TokenAccount struct {
	Identifier TokenAccountIdentifier `json:"identifier"`
	Balance    int64                  `json:"balance"`
	Hash       *Bytes32               `json:"hash,omitempty"`
}
