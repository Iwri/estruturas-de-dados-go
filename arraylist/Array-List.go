package main

/*
Operações implementadas:
- Criar a lista;
- Adicionar(Add);
- Obter(Get);
- Definir(Set);
- Remover (Remove);
- Redimensionar (Resize auto);
- Inserção em posição específica.
*/

type ArrayList struct {
	data     []int // Array interno
	size     int   // Quantidade de elementos
	capacity int   // Tamanho do array interno
}

func NewArrayList() *ArrayList {
	cap := 2

	return &ArrayList{
		data:     make([]int, cap),
		size:     0,
		capacity: cap,
	}
}

func (list *ArrayList) resize() {
	newCapacity := list.capacity * 2
	newData := make([]int, newCapacity)

	for i := 0; i < list.size; i++ {
		newData[i] = list.data[i]
	}

	list.data = newData
	list.capacity = newCapacity
}

func (list *ArrayList) Add(value int) {
	if list.size == list.capacity {
		list.resize()
	}

	list.data[list.size] = value
	list.size++
}

func (list *ArrayList) Get(index int) int {
	if index < 0 || index >= list.size {
		panic("index out of bounds")
	}
	return list.data[index]
}

func (list *ArrayList) Set(index int, value int) {
	if index < 0 || index >= list.size {
		panic("index out of bounds")
	}
	list.data[index] = value
}

func (list *ArrayList) Remove(index int) int {
	if index < 0 || index >= list.size {
		panic("index out of bounds")
	}
	removed := list.data[index]

	for i := index; i < list.size-1; i++ {
		list.data[i] = list.data[i+1]
	}

	list.size--
	return removed
}

func (list *ArrayList) AddAt(index int, value int) {
	if index < 0 || index > list.size {
		panic("index out of bounds")
	}

	if list.size == list.capacity {
		list.resize()
	}

	// Desloca elementos para a direita
	for i := list.size; i > index; i-- { // Começa do final para evitar sobrescrever
		list.data[i] = list.data[i-1]
	}

	list.data[index] = value
	list.size++
}

// Teste
func main() {
	list := NewArrayList()

	list.Add(10)
	list.Add(20)
	list.Add(30)
	list.Add(40)

	list.AddAt(1, 99)

	for i := 0; i < list.size; i++ {
		println(list.Get(i))
	}
}
