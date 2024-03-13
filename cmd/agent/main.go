package main

import (
	"github.com/artforteam2018/yametrics/internal/agent/app"
)

func main() {

	app.Run()
}

// Хранение памяти в HEAP работает через алгоритм Heap allocator (TCMalloc)
// SYS - stack + heap memory
// MALLOC, FREES - count of heap objects
// ALLOC, HEAPALLOC - активная память кучи
// HEAPSYS - Сколько выделено памяти для HEAP :: HEAPIDLE + HEAPINUSE
// HEAPIDLE - Сколько памяти свободно
// HEAPINUSE - как HEAPALLOC, только еще включаются незаполненные блоки памяти
// HEAPRELEASED - Сколько памяти из зарезервированной возвращено ОС, т.к. не нужно здесь и сейчас
// HEAPOBJECTS - Malloc - Frees
// STACKINUSE - используемые стаком блоки памяти
// STACKSYS - чаще всего равно STACKINUSE (еще выделенная память из OS в CGO)
// MSPANINUSE, MSPANSYS - память, занимаемая структурами, хранящими HEAP память
// MCACHEINUSE, MCACHESYS - память, занимаемая структурами, хранящими КЕШ HEAP памяти
// BuckHashSys, GCSys, OtherSys - системное потребление памяти (GC, ...)
// LastGC, PauseEnd - Последний обход GC
// PauseTotalNs, PauseNs - Сколько секунд был StopTheWorld
// NumGC, NumForcedGC - Кол-во обходов GC
// NextGC - Пока непонятно, почитать гайд по GC
// GCCPUFraction - Доля процессорного времени, занятого GC

// ТО ЧТО передается в fmt.Println(ПОЧЕМУ?) попадает в HEAP.
// go build -gcflags "-m" покажет, что попадает в HEAP

// МАССИВ можно создать с КОНСТАНТНЫМ размером
// МАССИВ передается копией, потому что это фактическое значение, а не ссылка, значит передать напрямую его нельзя
// МАССИВ не может быть расширен
// МАССИВ не уходит из стека в HEAP при cap > 2^16

// СЛАЙС, в которую записывается СЛАЙС данных - записывает ссылку на изначальный массив/слайс
// СЛАЙС может быть расширен через Append
// СЛАЙС передается по значению
// При переполнении капасити - создается полностью новый СЛАЙС

// И СЛАЙС И МАССИВ заполнены нулевым значением типа по умолчанию
// СРЕЗ МАССИВА будет держать ссылку на изначальный массив данных

// ПРО ООП
// Инкапсуляция Деление на методы и свойства, их защита от влияния извне - Struct
// Наследование (has-a) композиция типов Person { Address };; (is-a) интерфейсы, объединяющие типы
// Полиморфизм реализуется через интерфейсы

// Наследование бывает двух типов - через реализацию и через интерфейс. Через реализацию - плохо, потому что ломает концепции
// Полиморфизм - через него и работает наследование, это использование виртуальных методов класса-предка (который выступает интерфейсом)
// Полиморфизм - любой мезанизм, при котором в зависимости от типа объекта, функция выполняет разное действие

// Замыкания - это функция, использующая более недоступную ниоткуда переменную извне (из функции, которая ее создала)

// https://www.youtube.com/watch?v=q4HoWwdZUHs
