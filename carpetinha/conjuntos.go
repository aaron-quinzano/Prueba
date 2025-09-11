package carpetinha

type IntSet struct {
	elements map[int]bool
}

func NewIntSet(elements ...int) *IntSet {
	s := &IntSet{elements: make(map[int]bool)}
	for _, v := range elements {
		s.Add(v)
	}
	return s
}

func (s *IntSet) Add(element int) {
	s.elements[element] = true
}

func (s *IntSet) Remove(element int) {
	delete(s.elements, element)
}

func (s *IntSet) Contains(element int) bool {
	return s.elements[element]
}

func (s *IntSet) Size() int {
	return len(s.elements)
}

func (s *IntSet) Values() []int {
	values := make([]int, 0, s.Size())
	for k := range s.elements {
		values = append(values, k)
	}
	return values
}

// Dado un conjunto A y un conjunto B, la unión de los conjuntos A y B será otro
// conjunto que estará formado por todos los elementos de A, con todos los
// elementos de B sin repetir ningún elemento. O(n)
func (s *IntSet) Union(other *IntSet) *IntSet {
	// Implementar
	union := NewIntSet(s.Values()...)
	for _, e := range other.Values() {
		union.Add(e)
	}
	return union
}

// Dado un conjunto A y un conjunto B, la intersección de los conjuntos A y B
// será otro conjunto que estará formado por los elementos de A y los elementos
// de B que sean comunes, los elementos no comunes entre A y B, serán excluidos. O(n)
func (s *IntSet) Intersection(other *IntSet) *IntSet {
	// Implementar
	inter := NewIntSet()
	for _, e := range s.Values() {
		if other.Contains(e) {
			inter.Add(e)
		}
	}
	return inter
}

// Dado un conjunto A y un conjunto B, la diferencia de los conjuntos A y B será
// otro conjunto que estará formado por los elementos de A que no estan
// presentes en B.
func (s *IntSet) Difference(other *IntSet) *IntSet {
	// Implementar
	diff := NewIntSet()
	for _, e := range s.Values() {
		if !other.Contains(e) {
			diff.Add(e)
		}
	}
	return diff
}

// Dado un conjunto A y un conjunto B, la diferencia simétrica de los conjuntos
// A y B será otro conjunto que estará formado por todos los elementos no
// comunes a los conjuntos A y B. O(n)
func (s *IntSet) SymmetricDifference(other *IntSet) *IntSet {
	// Implementar
	return s.Union(other).Difference(s.Intersection(other))
}

// Un conjunto A es igual a un conjunto B si ambos conjuntos tienen los mismos
// elementos.
func (s *IntSet) Equal(other *IntSet) bool {
	// Implementar
	if s.Size() != other.Size() {
		return false
	}
	for _, e := range s.Values() {
		if !other.Contains(e) {
			return false
		}
	}
	return true
}

// El conjunto `other` es subconjunto de `s` si todos los elementos de `other`
// están incluidos en `s`.
func (s *IntSet) Subset(other *IntSet) bool {
	// Implementar
	return other.Equal(s.Intersection(other))
}

// El conjunto `other` es superconjunto de `s` si `other` contiene todos los
// elementos de `s`.
func (s *IntSet) Superset(other *IntSet) bool {
	// Implementar
	return s.Equal(other.Intersection(s))
}
