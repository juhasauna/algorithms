package kolman

import (
	"fmt"
	"log"
	"reflect"
	"sort"
	"strconv"
)

const keySeparator = ";"

var none = struct{}{}

type relation struct {
	from string
	to   string
}
type set map[string]struct{}

func (x set) getFirst() string {
	if len(x) == 0 {
		log.Fatalf("getting items from empty set")
	}
	s := []string{}
	for key := range x {
		s = append(s, key)
	}
	sort.Slice(s, func(i, j int) bool {
		return s[i] < s[j]
	})
	return s[0]
}

func newSet(keys ...string) set {
	new := make(set)
	new.add(keys...)
	return new
}
func (x *set) add(keys ...string) {
	temp := *x
	for _, v := range keys {
		temp[v] = none
	}
}
func (x *set) eq(y set) bool {
	for key := range *x {
		if _, ok := y[key]; !ok {
			return false
		}
	}
	return true
}

func (x *set) cartesianProduct(y set) relations {
	result := relations{relationSet: make(map[string]relation)}
	for i := range *x {
		for j := range y {
			result.add(relation{i, j})
		}
	}
	return result
}

func (x *set) cartesianSquare() relations {
	return x.cartesianProduct(*x)
}

type relations struct {
	relationSet map[string]relation
}

func (x *relations) range_() set {
	rng := make(set)
	for _, v := range x.relationSet {
		rng.add(v.to)
	}
	return rng
}

func newRelationsFromMatrix(values [][]int) relations {
	result := relations{relationSet: make(map[string]relation)}
	for i, row := range values {
		for j, cell := range row {
			if cell == 1 {
				r := relation{from: strconv.Itoa(i), to: strconv.Itoa(j)}
				result.add(r)
			}
		}
	}
	return result
}
func newRelations(values []relation) relations {
	r := relations{relationSet: make(map[string]relation)}
	for _, v := range values {
		r.add(v)
	}
	return r
}
func (x relation) key() string {
	return x.from + keySeparator + x.to
}
func (x relation) keyReverse() string {
	return x.to + keySeparator + x.from
}
func (x *relations) add(new relation) {
	if v, ok := x.relationSet[new.key()]; ok {
		log.Fatalf("Tried to add an existing relation to set: %v", v)
	}
	x.relationSet[new.key()] = new
}

func (x *relations) isAntisymmetric() bool {
	for _, relation := range x.relationSet {
		if relation.from != relation.to {
			if _, ok := x.relationSet[relation.keyReverse()]; ok {
				return false
			}
		}
	}
	return true
}
func (x *relations) isAsymmetric() bool {
	for _, relation := range x.relationSet {
		if _, ok := x.relationSet[relation.keyReverse()]; ok {
			return false
		}
	}
	return true
}
func (x *relations) isIrreflexive() bool {
	for element := range x.domain() {
		key := relation{element, element}.key()
		if _, ok := x.relationSet[key]; ok {
			return false
		}
	}
	return true
}

func (x *relations) isReflexive() bool {
	domain := x.domain()
	if len(domain) == 0 {
		return false
	}
	for element := range domain {
		key := relation{element, element}.key()
		if _, ok := x.relationSet[key]; !ok {
			return false
		}
	}
	return true
}
func (x *relations) isSymmetric() bool {
	for _, relation := range x.relationSet {
		if _, ok := x.relationSet[relation.keyReverse()]; !ok {
			return false
		}
	}
	return true
}

func isReflexive(r []relation) bool {
	m := make(map[string]bool)
	for _, rel := range r {
		if v, ok := m[rel.from]; !ok {
			m[rel.from] = rel.from == rel.to
		} else {
			if !v {
				m[rel.from] = rel.from == rel.to
			}
		}
	}
	for _, v := range m {
		if !v {
			return false
		}
	}
	return true
}
func isReflexiveMatrix(m [][]int) bool {
	diagonalSize := func() int {
		if len(m) < len(m[0]) {
			return len(m)
		}
		return len(m[0])
	}
	for i := 0; i < diagonalSize(); i++ {
		if m[i][i] == 0 {
			return false
		}
	}
	return true
}
func isSymmetricMatrix(m [][]int) bool {
	if len(m) != len(m[0]) {
		return false
	}
	size := len(m)

	for i := 0; i < size; i++ {
		for j := 0; j < size; j++ {
			if m[i][j] != m[j][i] {
				return false
			}
		}
	}
	return true
}

func (x *relations) domain() set {
	// Not sure if this is always correct. Certainly if e.g. R is subset of A x A, then it's fine.
	m := make(set)
	for _, relation := range x.relationSet {
		m[relation.from] = none
		m[relation.to] = none
	}
	return m
}

func (x *relations) sortedDomain() []string {
	s := []string{}
	for key := range x.domain() {
		s = append(s, key)
	}
	sort.Slice(s, func(i, j int) bool {
		return s[i] < s[j]
	})
	return s
}

func (x *relations) toMatrix() (result boolMatrix) {
	domain := x.sortedDomain()
	for i, row := range domain {
		result = append(result, []int{})
		for j, col := range domain {
			result[i] = append(result[i], 0) // Init as 0 (false). I.e. no relation between row index and colindex.
			key := relation{row, col}.key()
			if _, ok := x.relationSet[key]; ok {
				result[i][j] = 1
			}
		}
	}
	return result
}

type boolMatrix [][]int

func (x *boolMatrix) rmRow(i int) boolMatrix {
	x_ := *x
	fmt.Println(x_)
	x_ = append(x_[:i], x_[i+1:]...)
	fmt.Println(x_)
	return x_
}
func (x *boolMatrix) rmCol(j int) {
	x_ := *x
	for i, row := range x_ {
		row = append(row[:j], row[j+1:]...)
		x_[i] = row
	}
	x = &x_
}
func (x *boolMatrix) swapRows(a, b int) {
	x_ := *x
	for i, v := range x_[a] {
		temp := x_[b][i]
		x_[b][i] = v
		x_[a][i] = temp
	}
}

func (x boolMatrix) rows() int {
	return len(x)
}
func (x boolMatrix) columns() int {
	return len(x[0])
}
func (x boolMatrix) union(b boolMatrix) (result boolMatrix) {
	for i, row := range x {
		result = append(result, []int{})
		for j, v := range row {
			if v+b[i][j] > 0 {
				result[i] = append(result[i], 1)
			}
		}
	}
	return result
}
func (x boolMatrix) intersection(b boolMatrix) (result boolMatrix) {
	for i, row := range x {
		result = append(result, []int{})
		for j, v := range row {
			result[i] = append(result[i], 0)
			if v+b[i][j] == 2 {
				result[i][j] = 1
			}
		}
	}
	return result
}
func (x boolMatrix) complement() (result boolMatrix) {
	for i, row := range x {
		result = append(result, []int{})
		for j, v := range row {
			result[i] = append(result[i], 0)
			if v == 0 {
				result[i][j] = 1
			}
		}
	}
	return result
}

func (x boolMatrix) transpose() (result boolMatrix) { // TODO: test & refactor
	rLen := x.rows()
	cLen := x.columns()
	r, c := 0, 0
	resRow := []int{}
	if rLen < cLen {
		for r < cLen {
			for c < rLen {
				cc := c % cLen
				resRow = append(resRow, x[cc][r])
				c++
			}
			c = 0
			result = append(result, resRow)
			resRow = []int{}
			r++
		}
	} else {
		for r < cLen {
			rr := r % rLen
			for c < rLen {
				resRow = append(resRow, x[c][rr])
				c++
			}
			c = 0
			result = append(result, resRow)
			resRow = []int{}
			r++
		}
	}
	return result
}

func initSquareBoolMatrix(size int) (result boolMatrix) {
	for i := 0; i < size; i++ {
		result = append(result, []int{})
		for j := 0; j < size; j++ {
			result[i] = append(result[i], 0)
		}
	}
	return result
}

func (x boolMatrix) product(right boolMatrix) boolMatrix {
	leftRowLen, leftColLen := len(x), len(x[0])
	rightColLen, rightRowLen := len(right[0]), len(right)
	if leftRowLen != rightColLen {
		log.Fatalf("unable to multiply given matrices, left rowLen (%d) should equal right colLen (%d)", leftRowLen, rightColLen)
	}
	if leftColLen != rightRowLen {
		log.Fatalf("unable to multiply given matrices, left colLen (%d) should equal right rowLen (%d)", leftColLen, rightRowLen)
	}
	resultSize := leftColLen
	if resultSize < leftRowLen {
		resultSize = leftRowLen
	}
	result := initSquareBoolMatrix(resultSize)

	for i, row := range x {
		resultRow := 0
		for j := 0; j < rightColLen; j++ {
			temp := 0
			for rightJ := 0; rightJ < rightRowLen; rightJ++ {
				if right[rightJ][j] == 1 && row[rightJ] == 1 {
					temp = 1
				}
			}
			result[i][resultRow] = temp
			resultRow++
		}
	}
	return result
}

func (x *relations) isTransitive() bool {
	// From the book: TRANS algorithm
	m := x.toMatrix()
	for i, row := range m {
		for j, v := range row {
			if v == 1 && i != j {
				for k := 0; k < len(m); k++ {
					if m[j][k] == 1 && m[i][k] != 1 {
						return false
					}
				}
			}
		}
	}
	return true
}
func (x *relations) isTransitive_bySquaring() bool {
	// This does not work
	m := x.toMatrix()
	squared := m.product(m)
	return reflect.DeepEqual(m, squared)
}

func (x boolMatrix) toRelations() (result relations) {
	result.relationSet = make(map[string]relation)
	for i, row := range x {
		for j, v := range row {
			if v == 1 {
				result.add(relation{strconv.Itoa(i), strconv.Itoa(j)})
			}
		}
	}
	return result
}
func (x boolMatrix) isEquivalenceRelation() bool {
	relations := x.toRelations()
	return relations.isReflexive() && relations.isSymmetric() && relations.isTransitive()
}

type properties struct {
	reflexive     int
	irreflexive   int
	symmetric     int
	asymmetric    int
	antisymmetric int
	transitive    int
}

func (x relations) getProperties() (p properties) {
	if x.isReflexive() {
		p.reflexive = 1
	}
	if x.isIrreflexive() {
		p.irreflexive = 1
	}
	if x.isSymmetric() {
		p.symmetric = 1
	}
	if x.isAsymmetric() {
		p.asymmetric = 1
	}
	if x.isAntisymmetric() {
		p.antisymmetric = 1
	}
	if x.isTransitive() {
		p.transitive = 1
	}
	return p
}

func integerRelations(from, to int, f func(int, int) bool) relations {
	rels := []relation{}
	for i := from; i < to; i++ {
		for j := from; j < to; j++ {
			if f(i, j) {
				rels = append(rels, relation{strconv.Itoa(i), strconv.Itoa(j)})
			}
		}
	}
	return newRelations(rels)
}

// poset AKA partially ordered set
func (x boolMatrix) isPoset() bool {
	// Appendix C. Chapter 6 coding ex. 1. Write a subroutine that determines if a relation R represented by its matrix is a partial order.
	result := x.toRelations().getProperties() == properties{reflexive: 1, antisymmetric: 1, transitive: 1}
	return result
}

// Warshall's algorithm
func (x boolMatrix) transitiveClosure(iter int) boolMatrix {
	var xk boolMatrix
	for _, row := range x {
		xk = append(xk, copySlice(row))
	}
	kRowValues := copySlice(x[iter])
	kColValues := x.getColValues(iter)
	for i, rowVal := range kRowValues {
		if rowVal == 0 {
			continue
		}
		for j, colVal := range kColValues {
			if colVal == 0 {
				continue
			}
			xk[j][i] = 1
		}
	}
	if iter == len(kRowValues)-1 || iter == len(kColValues)-1 {
		return xk
	}
	return xk.transitiveClosure(iter + 1)
}

func copySlice(src []int) []int {
	var tempRow []int
	for _, v := range src {
		tempRow = append(tempRow, v)
	}
	return tempRow
}

func (x boolMatrix) getColValues(i int) (result []int) {
	for _, row := range x {
		result = append(result, row[i])
	}
	return result
}
