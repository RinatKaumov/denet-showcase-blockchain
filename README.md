# DenetShow - Простая реализация блокчейна на Go

DenetShow - это учебный проект, демонстрирующий базовую реализацию блокчейна с использованием алгоритма Proof-of-Work (PoW) и параллельной обработкой.

## Структура проекта

```
denetshow/
├── core/                 # Основные компоненты блокчейна
│   ├── block.go         # Реализация структуры блока
│   ├── blockchain.go    # Логика работы блокчейна
│   └── transaction.go   # Определение транзакций
├── crypto/              # Криптографические функции
├── main.go              # Точка входа и демонстрация
└── pow_test.go          # Тесты производительности PoW
```

## Основные компоненты

### Блок (Block)
- Заголовок (Header):
  - Height (uint32) - высота блока
  - Timestamp (int64) - временная метка
  - PrevHash ([]byte) - хеш предыдущего блока
  - DataHash ([]byte) - хеш транзакций
  - Nonce (uint64) - число для PoW
- Список транзакций
- Хеш блока

### Транзакция (Transaction)
- From (string) - отправитель
- To (string) - получатель
- Amount (int) - сумма

### Блокчейн (Blockchain)
- Массив блоков
- Функция создания нового блокчейна с генезис-блоком
- Метод добавления новых блоков

## Особенности реализации

1. **Proof-of-Work**:
   - Параллельная реализация майнинга
   - Настраиваемое количество ядер процессора
   - Тесты производительности для разных конфигураций

2. **Криптография**:
   - Использование SHA-256 для хеширования
   - Сериализация данных с помощью gob

3. **Производительность**:
   - Параллельная обработка PoW
   - Оптимизированная структура данных
   - Тесты производительности для разных конфигураций CPU

## Использование

```go
// Создание нового блокчейна
chain := core.NewBlockchain()

// Добавление блока с транзакциями
chain.AddBlock([]core.Transaction{
    core.NewTransaction("Alice", "Bob", 10),
    core.NewTransaction("Bob", "Charlie", 5),
})

// Вывод информации о блоках
for _, block := range chain.Blocks {
    fmt.Printf("Height: %d\n", block.Header.Height)
    fmt.Printf("Hash: %x\n", block.Hash)
    fmt.Printf("PrevHash: %x\n", block.Header.PrevHash)
    fmt.Printf("Nonce: %d\n", block.Header.Nonce)
}
```

## Тестирование

Для запуска тестов производительности PoW:
```bash
go test -v pow_test.go
```

Тесты показывают время нахождения nonce для разного количества ядер процессора.

## Требования

- Go 1.23.7 или выше
- Многоядерный процессор для оптимальной производительности PoW
