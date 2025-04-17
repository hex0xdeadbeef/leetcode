package main

import (
	"sync"
	"testing"
)

// BenchmarkPersonParallel тестирует структуру Person с cache-line padding
func BenchmarkPersonParallel(b *testing.B) {
	const n = 1000
	persons := make([]Person, n)
	for i := range persons {
		persons[i] = Person{person: person{name: "John", age: 30, isAdult: true}}
	}

	var mu sync.Mutex // Добавляем мьютекс для синхронизации доступа

	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		i := 0
		for pb.Next() {
			mu.Lock()
			p := persons[i%n] // Копируем элемент перед изменением
			p.BalanceAge()
			persons[i%n] = p
			mu.Unlock()
			i++
		}
	})
}

// BenchmarkPlainPersonParallel тестирует структуру PlainPerson без cache-line padding
func BenchmarkPlainPersonParallel(b *testing.B) {
	const n = 1000
	persons := make([]PlainPerson, n)
	for i := range persons {
		persons[i] = PlainPerson{person: person{name: "John", age: 30, isAdult: true}}
	}

	var mu sync.Mutex

	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		i := 0
		for pb.Next() {
			mu.Lock()
			p := persons[i%n]
			p.BalanceAge()
			persons[i%n] = p
			mu.Unlock()
			i++
		}
	})
}
