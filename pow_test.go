package main

import (
	"denetshow/core"
	"denetshow/crypto"
	"fmt"
	"runtime"
	"testing"
	"time"
)

func TestProofOfWork(t *testing.T) {
	// Заголовок таблицы
	fmt.Println("| Ядра | Время (сек) до нахождения nonce |")
	fmt.Println("|------|-------------------------------|")

	// Тестируем разные количества ядер
	for _, cpu := range []int{1, 2, 4, 8} {
		// Устанавливаем количество ядер
		runtime.GOMAXPROCS(cpu)

		// Создаем блок с транзакциями
		transactions := []core.Transaction{
			core.NewTransaction("Alice", "Bob", 10),
			core.NewTransaction("Bob", "Charlie", 5),
		}
		block := core.NewBlock(transactions, []byte("prevhash"), 1)

		// Передаем данные в ProofOfWork
		powInput := crypto.BlockHeaderData{
			PrevHash:  block.Header.PrevHash,
			DataHash:  block.Header.DataHash,
			Timestamp: block.Header.Timestamp,
			Height:    block.Header.Height,
		}

		pow := crypto.NewProofOfWork(powInput)

		// Замеряем время работы майнинга
		start := time.Now()
		_, _ = pow.RunParallel()
		duration := time.Since(start)

		// Выводим результаты
		fmt.Printf("| %d    | %.2f                          |\n", cpu, duration.Seconds())
	}
}
