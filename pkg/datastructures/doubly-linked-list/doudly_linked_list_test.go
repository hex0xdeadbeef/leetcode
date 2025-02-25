package doublelinkedlist

import "testing"

// TestInsertTail проверяет вставку элементов в хвост списка.
func TestInsertTail(t *testing.T) {
	list := New()

	// Вставка в пустой список.
	list.InsertTail(10)
	if list.head == nil || list.tail == nil {
		t.Fatalf("После InsertTail в пустом списке head или tail равны nil")
	}
	if list.head != list.tail {
		t.Errorf("При единственном элементе head и tail должны совпадать")
	}
	if list.head.Val != 10 {
		t.Errorf("Ожидалось значение 10, получено %d", list.head.Val)
	}
	if list.count != 1 {
		t.Errorf("Ожидалось cnt = 1, получено %d", list.count)
	}

	// Вставка второго элемента.
	list.InsertTail(20)
	if list.head == nil || list.tail == nil {
		t.Fatalf("После второй вставки head или tail равны nil")
	}
	if list.head.Val != 10 {
		t.Errorf("Ожидалось, что head останется равным 10, получено %d", list.head.Val)
	}
	if list.tail.Val != 20 {
		t.Errorf("Ожидалось, что tail станет равным 20, получено %d", list.tail.Val)
	}
	// При единственном элементе до вставки, InsertTail должен установить:
	// list.head.prev -> новый узел, а новый узел.next -> старый head.
	if list.head.prev == nil || list.head.prev.Val != 20 {
		t.Errorf("Ожидалось, что head.prev указывает на элемент со значением 20")
	}
	if list.tail.next == nil || list.tail.next.Val != 10 {
		t.Errorf("Ожидалось, что tail.next указывает на элемент со значением 10")
	}
	if list.count != 2 {
		t.Errorf("Ожидалось cnt = 2, получено %d", list.count)
	}

	// Вставка третьего элемента (случай default).
	list.InsertTail(30)
	if list.tail.Val != 30 {
		t.Errorf("Ожидалось, что tail станет равным 30, получено %d", list.tail.Val)
	}
	// Проверяем, что новый tail ссылается на предыдущий tail (20).
	if list.tail.next == nil || list.tail.next.Val != 20 {
		t.Errorf("Ожидалось, что tail.next указывает на элемент со значением 20")
	}
	if list.count != 3 {
		t.Errorf("Ожидалось cnt = 3, получено %d", list.count)
	}
}

// TestInsertHead проверяет вставку элементов в голову списка.
func TestInsertHead(t *testing.T) {
	list := New()

	// Вставка в пустой список.
	list.InsertHead(100)
	if list.head == nil {
		t.Fatalf("После InsertHead в пустом списке head равен nil")
	}
	// Согласно реализации, в пустом списке tail не устанавливается.
	if list.tail != nil {
		t.Errorf("Ожидалось, что tail останется nil, получено %v", list.tail)
	}
	if list.head.Val != 100 {
		t.Errorf("Ожидалось значение 100, получено %d", list.head.Val)
	}
	if list.count != 1 {
		t.Errorf("Ожидалось cnt = 1, получено %d", list.count)
	}

	// Вставка второго элемента.
	list.InsertHead(200)
	// В случае cnt == onlyHead происходит следующее:
	// list.tail = старый head, старый head.next = новый узел, новый узел.prev = старый head, list.head = новый узел.
	if list.head.Val != 200 {
		t.Errorf("Ожидалось, что head станет равным 200, получено %d", list.head.Val)
	}
	if list.tail == nil || list.tail.Val != 100 {
		t.Errorf("Ожидалось, что tail станет равным 100, получено %v", list.tail)
	}
	if list.tail.next == nil || list.tail.next.Val != 200 {
		t.Errorf("Ожидалось, что tail.next указывает на head с значением 200")
	}
	if list.head.prev == nil || list.head.prev.Val != 100 {
		t.Errorf("Ожидалось, что head.prev указывает на tail с значением 100")
	}
	if list.count != 2 {
		t.Errorf("Ожидалось cnt = 2, получено %d", list.count)
	}

	// Вставка третьего элемента (случай default).
	list.InsertHead(300)
	if list.head.Val != 300 {
		t.Errorf("Ожидалось, что head станет равным 300, получено %d", list.head.Val)
	}
	if list.head.prev == nil || list.head.prev.Val != 200 {
		t.Errorf("Ожидалось, что head.prev указывает на предыдущий head со значением 200")
	}
	if list.count != 3 {
		t.Errorf("Ожидалось cnt = 3, получено %d", list.count)
	}
}

// TestRemoveTail проверяет удаление элементов с хвоста списка.
func TestRemoveTail(t *testing.T) {
	list := New()

	// Удаление из пустого списка должно вернуть nil.
	if node := list.PopTail(); node != nil {
		t.Errorf("Ожидалось, что RemoveTail из пустого списка вернет nil")
	}

	// Вставляем один элемент и удаляем его.
	list.InsertTail(50)
	node := list.PopTail()
	if node == nil || node.Val != 50 {
		t.Errorf("Ожидалось, что RemoveTail вернет узел со значением 50")
	}
	if list.head != nil || list.tail != nil {
		t.Errorf("После удаления единственного элемента список должен стать пустым")
	}
	if list.count != 0 {
		t.Errorf("Ожидалось cnt = 0, получено %d", list.count)
	}

	// Вставляем несколько элементов.
	list.InsertTail(10)
	list.InsertTail(20)
	list.InsertTail(30)
	// Ожидаем: head = 10, tail = 30, связь: tail.next -> 20, затем 20.next -> 10.
	node = list.PopTail()
	if node == nil || node.Val != 30 {
		t.Errorf("Ожидалось, что RemoveTail вернет узел со значением 30, получено %v", node)
	}
	if list.tail == nil || list.tail.Val != 20 {
		t.Errorf("После удаления tail должен быть узел со значением 20")
	}
	if list.count != 2 {
		t.Errorf("Ожидалось cnt = 2, получено %d", list.count)
	}
}

// TestRemoveHead проверяет удаление элементов с головы списка.
func TestRemoveHead(t *testing.T) {
	list := New()

	// Удаление из пустого списка должно вернуть nil.
	if node := list.PopHead(); node != nil {
		t.Errorf("Ожидалось, что RemoveHead из пустого списка вернет nil")
	}

	// Вставляем один элемент и удаляем его.
	list.InsertHead(70)
	node := list.PopHead()
	if node == nil || node.Val != 70 {
		t.Errorf("Ожидалось, что RemoveHead вернет узел со значением 70")
	}
	if list.head != nil || list.tail != nil {
		t.Errorf("После удаления единственного элемента список должен стать пустым")
	}
	if list.count != 0 {
		t.Errorf("Ожидалось cnt = 0, получено %d", list.count)
	}

	// Вставляем несколько элементов.
	list.InsertHead(100)
	list.InsertHead(200)
	list.InsertHead(300)
	// Согласно логике InsertHead, последний вставленный (300) является head, а tail = 100.
	node = list.PopHead()
	if node == nil || node.Val != 300 {
		t.Errorf("Ожидалось, что RemoveHead вернет узел со значением 300, получено %v", node)
	}
	if list.head == nil || list.head.Val != 200 {
		t.Errorf("После удаления head должен быть узел со значением 200")
	}
	if list.count != 2 {
		t.Errorf("Ожидалось cnt = 2, получено %d", list.count)
	}
}
