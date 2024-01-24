/*
Copyright 2023 The Ketches Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

	http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package image

import (
	"log"
	"strings"

	"github.com/containers/image/docker/reference"
)

// Parse parses the image name and returns the registry and name(name is not contains registry domain).
func Parse(image string) (registry, name string, err error) {
	named, err := reference.ParseDockerRef(image)
	if err != nil {
		log.Println("Parse image failed.", err.Error(), image)
		return
	}

	registry = reference.Domain(named)
	name = strings.TrimPrefix(reference.TagNameOnly(named).String(), registry+"/")
	return
}
