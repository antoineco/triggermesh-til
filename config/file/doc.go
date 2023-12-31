/*
Copyright 2021 TriggerMesh Inc.

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

// Package file contains utilities for parsing and decoding Bridge Description Files.
//
// It is responsible for configuration parsing and static validation, but not
// for the runtime aspects of the TriggerMesh Integration Language, such as
// evaluating expressions, which is the responsibility of the "lang" package.
package file
