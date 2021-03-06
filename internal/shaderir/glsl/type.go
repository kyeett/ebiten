// Copyright 2020 The Ebiten Authors
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

package glsl

import (
	"fmt"

	"github.com/hajimehoshi/ebiten/internal/shaderir"
)

func typeString(t *shaderir.Type) (string, string) {
	switch t.Main {
	case shaderir.Array:
		t0, t1 := typeString(&t.Sub[0])
		return t0 + t1, fmt.Sprintf("[%d]", t.Length)
	case shaderir.Struct:
		panic("glsl: a struct is not implemented")
	default:
		return basicTypeString(t.Main), ""
	}
}

func basicTypeString(t shaderir.BasicType) string {
	switch t {
	case shaderir.None:
		return "?(none)"
	case shaderir.Bool:
		return "bool"
	case shaderir.Int:
		return "int"
	case shaderir.Float:
		return "float"
	case shaderir.Vec2:
		return "vec2"
	case shaderir.Vec3:
		return "vec3"
	case shaderir.Vec4:
		return "vec4"
	case shaderir.Mat2:
		return "mat2"
	case shaderir.Mat3:
		return "mat3"
	case shaderir.Mat4:
		return "mat4"
	case shaderir.Array:
		return "?(array)"
	case shaderir.Struct:
		return "?(struct)"
	default:
		return fmt.Sprintf("?(unknown type: %d)", t)
	}
}

func builtinFuncString(f shaderir.BuiltinFunc) string {
	switch f {
	case shaderir.Dfdx:
		return "dFdx"
	case shaderir.Dfdy:
		return "dFdy"
	default:
		return string(f)
	}
}
