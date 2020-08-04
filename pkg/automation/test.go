/*
Copyright 2020 Google LLC

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    https://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package automation

import (
	"fmt"

	"google.golang.org/api/compute/v1"
)

func test (toTest string, value interface{}, valueMatcher string) (boolean, error) {
	if (value == nil) {
		if (valueMatcher == nil) {
			// return  error
		}
		
		r, err := regexp.Compile("^" + valueMatcher + "$")
	
		if err != nil {
			// return error
		}

		return r.matchString(toTest)
	}

	valueString, ok = operation.(string) 
	if (!ok) {
		// handle error (if nil, it's fine)
	}

	return valueString == toTest
}

func (s *googleService) TestMachineType(project string, zone string, instance string, value interface{}, valueMatcher *string) (boolean, error) {
	machineInstance := s.getInstance(project, zone, instance)
	machineType := machineInstance.MachineType

	return test (machineType, value, valueMatcher)
}

func (s *googleService) TestStatus(project string, zone string, instance string, value interface{}, valueMatcher string) {
	machineInstance := s.getInstance(project, zone, instance)
	status := machineInstance.Status

	return test (status, value, valueMatcher)

}