package crypto

import (
	"bytes"
	"crypto/sha256"
	"encoding/binary"
	"fmt"
	"math"
	"math/big"
	"runtime"
	"sync"
)

const targetBits = 24

type BlockHeaderData struct {
	PrevHash  []byte
	DataHash  []byte
	Timestamp int64
	Height    uint32
}

type ProofOfWork struct {
	Header BlockHeaderData
	Target *big.Int
}

func NewProofOfWork(h BlockHeaderData) *ProofOfWork {
	target := big.NewInt(1)
	target.Lsh(target, uint(256-targetBits))
	return &ProofOfWork{Header: h, Target: target}
}

func (pow *ProofOfWork) prepareData(nonce uint64) []byte {
	data := bytes.Join([][]byte{
		pow.Header.PrevHash,
		pow.Header.DataHash,
		Uint64ToBytes(uint64(pow.Header.Timestamp)),
		Uint32ToBytes(pow.Header.Height),
		Uint64ToBytes(nonce),
	}, []byte{})
	return data
}

func (pow *ProofOfWork) RunParallel() (uint64, []byte) {
	numCPU := runtime.NumCPU()
	result := make(chan struct {
		nonce uint64
		hash  []byte
	}, 1)
	done := make(chan struct{})
	var wg sync.WaitGroup

	wg.Add(numCPU)

	for i := 0; i < numCPU; i++ {
		go func(start uint64) {
			defer wg.Done()
			for nonce := start; nonce < math.MaxUint64; nonce += uint64(numCPU) {
				select {
				case <-done:
					return
				default:
					data := pow.prepareData(nonce)
					hash := sha256.Sum256(data)

					var hashInt big.Int
					hashInt.SetBytes(hash[:])
					if hashInt.Cmp(pow.Target) == -1 {
						select {
						case result <- struct {
							nonce uint64
							hash  []byte
						}{nonce, hash[:]}:
							fmt.Printf("Goroutine %d нашла nonce: %d\n", start, nonce)
						case <-done:
							// другой уже отправил — выходим
						}
						return
					}
				}
			}
		}(uint64(i))
	}

	// Ждём, пока кто-то найдёт
	res := <-result
	close(done)

	// Ждём завершения всех горутин
	wg.Wait()
	return res.nonce, res.hash
}

func Uint64ToBytes(num uint64) []byte {
	buf := new(bytes.Buffer)
	_ = binary.Write(buf, binary.BigEndian, num)
	return buf.Bytes()
}

func Uint32ToBytes(num uint32) []byte {
	buf := new(bytes.Buffer)
	_ = binary.Write(buf, binary.BigEndian, num)
	return buf.Bytes()
}
