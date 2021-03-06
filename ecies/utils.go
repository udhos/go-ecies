/*
	utils.go

	Utility script for reading, writing files and hex encoding/decoding

	THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS
	"AS IS" AND ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT
	LIMITED TO, THE IMPLIED WARRANTIES OF MERCHANTABILITY AND FITNESS FOR
	A PARTICULAR PURPOSE ARE DISCLAIMED. IN NO EVENT SHALL THE COPYRIGHT
	OWNER OR CONTRIBUTORS BE LIABLE FOR ANY DIRECT, INDIRECT, INCIDENTAL,
	SPECIAL, EXEMPLARY, OR CONSEQUENTIAL DAMAGES (INCLUDING, BUT NOT
	LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR SERVICES; LOSS OF USE,
	DATA, OR PROFITS; OR BUSINESS INTERRUPTION) HOWEVER CAUSED AND ON ANY
	THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY, OR TORT
	(INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE
	OF THIS SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.

	utils.go Daniel Havir, 2018
*/

package ecies

import (
	"io"
	"strconv"
)

func getCryptoRandVec(rand io.Reader, len int) ([]byte, error) {
	out := make([]byte, len)
	_, err := io.ReadFull(rand, out)
	return out, err
}

func to32ByteArray(in []byte) *[32]byte {
	if len(in) != 32 {
		panic("Input array size does not match. Expected 32, but got " + strconv.Itoa(len(in)))
	}
	var out [32]byte
	for i := 0; i < 32; i++ {
		out[i] = in[i]
	}

	return &out
}

func to16ByteArray(in []byte) *[16]byte {
	if len(in) != 16 {
		panic("Input array size does not match. Expected 16, but got " + strconv.Itoa(len(in)))
	}
	var out [16]byte
	for i := 0; i < 16; i++ {
		out[i] = in[i]
	}

	return &out
}
