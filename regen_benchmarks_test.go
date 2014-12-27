/*
Copyright 2014 Zachary Klippenstein

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

package regen

import (
	"github.com/zach-klippenstein/goregen/util"
	"testing"
)

const BigFancyRegexp = `
POST (/[-a-zA-Z0-9_.]{3,12}){3,6}
Content-Length: [0-9]{2,3}
X-Auth-Token: [a-zA-Z0-9+/]{64}

([A-Za-z0-9+/]{64}
){3,15}[A-Za-z0-9+/]{60}([A-Za-z0-9+/]{2}==|[A-Za-z0-9+/]{3}=)
`

var rng = util.NewRand(42)

// Benchmarks the code that creates generators.
// Doesn't actually run the generators.
func BenchmarkCreation(b *testing.B) {
	// Create everything here to save allocations in the loop.
	args := &GeneratorArgs{rng, 0, NewSerialExecutor()}

	for i := 0; i < b.N; i++ {
		NewGenerator(BigFancyRegexp, args)
	}
}

func BenchmarkSerialGeneration(b *testing.B) {
	args := &GeneratorArgs{
		Rng:      rng,
		Executor: NewSerialExecutor(),
	}
	generator, err := NewGenerator(BigFancyRegexp, args)
	if err != nil {
		panic(err)
	}
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		generator.Generate()
	}
}

func BenchmarkParallelGeneration(b *testing.B) {
	args := &GeneratorArgs{
		Rng:      rng,
		Executor: NewForkJoinExecutor(),
	}
	generator, err := NewGenerator(BigFancyRegexp, args)
	if err != nil {
		panic(err)
	}
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		generator.Generate()
	}
}