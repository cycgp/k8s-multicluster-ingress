// Copyright 2017 Google Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package instances

import (
	compute "google.golang.org/api/compute/v1"
)

func newInstance(igURL string) *compute.Instance {
	return &compute.Instance{
		Name: igURL,
		Zone: igURL,
		Tags: &compute.Tags{
			Items: []string{igURL},
		},
	}
}

// NewFakeInstanceGetter returns a new fake.
func NewFakeInstanceGetter() InstanceGetterInterface {
	return &FakeInstanceGetter{}
}

// FakeInstanceGetter to be used for tests.
type FakeInstanceGetter struct {
}

// Ensure this implements InstanceGetterInterface
var _ InstanceGetterInterface = &FakeInstanceGetter{}

// GetInstance retuns a new instance.
func (g *FakeInstanceGetter) GetInstance(igURL string) (*compute.Instance, error) {
	return newInstance(igURL), nil
}
