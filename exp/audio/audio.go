// Copyright 2015 Hajime Hoshi
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

package audio

import (
	"io"

	"github.com/hajimehoshi/ebiten/internal/audio"
)

type ReadSeekCloser interface {
	io.ReadSeeker
	io.Closer
}

// Queue queues the given data to the given channel.
// The given data is queued to the end of the buffer.
// This may not be played immediately when data already exists in the buffer.
//
// src's format must be linear PCM (16bits, 2 channel stereo, little endian)
// without a header (e.g. RIFF header).
//
// TODO: Pass sample rate and num of channels.
func Queue(src ReadSeekCloser, sampleRate int) error {
	return audio.Queue(src, sampleRate)
}
