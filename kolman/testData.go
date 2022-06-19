package kolman

import "log"

var testRelationsData relationsTestData
var boolMatrixData boolMatrixTestData

func initTestData() {
	initBoolMatrixTestData()
	initRelationsTestData()
}

type boolMatrixTestData struct {
	test1   boolMatrix
	cp04_04 map[int]boolMatrix
	data    map[string]boolMatrix
}

func (x boolMatrixTestData) get(key string) boolMatrix {
	item, ok := x.data[key]
	if !ok {
		log.Fatalf("key: %s doesn't exists in boolMatrixTestData map", key)
	}
	return item
}

func initBoolMatrixTestData() {
	x := boolMatrixTestData{
		test1: boolMatrix{
			{1, 1, 1, 0},
			{0, 1, 1, 0},
			{1, 1, 1, 0},
			{1, 0, 1, 1},
		},
		data:    make(map[string]boolMatrix),
		cp04_04: make(map[int]boolMatrix),
	}
	x.data["cp04.8_examp02"] = boolMatrix{
		{0, 1, 0, 0},
		{1, 0, 1, 0},
		{0, 0, 0, 1},
		{0, 0, 0, 0},
	}
	x.data["cp04.8_examp02_answer"] = boolMatrix{
		{1, 1, 1, 1},
		{1, 1, 1, 1},
		{0, 0, 0, 1},
		{0, 0, 0, 0},
	}
	x.data["cp04.8_examp03"] = boolMatrix{
		{1, 1, 0, 0, 0},
		{1, 1, 0, 0, 0},
		{0, 0, 1, 1, 0},
		{0, 0, 1, 1, 1},
		{0, 0, 0, 1, 1},
	}
	x.data["cp04.8_examp03_answer"] = boolMatrix{
		{1, 1, 0, 0, 0},
		{1, 1, 0, 0, 0},
		{0, 0, 1, 1, 1},
		{0, 0, 1, 1, 1},
		{0, 0, 1, 1, 1},
	}
	x.data["diag3"] = boolMatrix{
		{1, 0, 0},
		{0, 1, 0},
		{0, 0, 1},
	}
	x.data["cp01.5_ex13_A"] = boolMatrix{
		{1, 1, 0},
		{0, 1, 0},
		{1, 1, 0},
		{0, 0, 1},
	}
	x.data["cp01.5_ex13_B"] = boolMatrix{
		{1, 0, 0, 0},
		{0, 1, 1, 0},
		{1, 0, 1, 1},
	}
	x.data["cp01.5_ex13_product"] = boolMatrix{
		{1, 1, 1, 0},
		{0, 1, 1, 0},
		{1, 1, 1, 0},
		{1, 0, 1, 1},
	}
	x.data["cp06.1_ex11"] = boolMatrix{
		{1, 1, 1, 1, 1},
		{0, 1, 0, 1, 1},
		{0, 0, 1, 0, 1},
		{0, 0, 0, 1, 1},
		{0, 0, 0, 0, 1},
	}
	x.data["lessThan4x4"] = boolMatrix{
		{0, 1, 1, 1, 1},
		{0, 0, 1, 1, 1},
		{0, 0, 0, 1, 1},
		{0, 0, 0, 0, 1},
		{0, 0, 0, 0, 0},
	}

	x.cp04_04[11] = boolMatrix{
		{0, 1, 0, 1},
		{1, 0, 1, 1},
		{0, 1, 0, 0},
		{1, 1, 0, 0},
	}
	x.cp04_04[12] = boolMatrix{
		{1, 1, 0, 0},
		{1, 1, 0, 0},
		{0, 0, 1, 0},
		{0, 0, 0, 1},
	}
	x.cp04_04[13] = boolMatrix{
		{1, 1, 1, 1},
		{1, 1, 1, 1},
		{0, 1, 1, 1},
		{0, 0, 1, 1},
	}
	x.cp04_04[3701] = boolMatrix{ // Ex.37 a) reflexive and symmetric, but not transitive.
		{1, 1, 0, 0},
		{1, 1, 1, 0},
		{0, 1, 1, 0},
		{0, 0, 0, 1},
	}
	x.cp04_04[3702] = boolMatrix{ // Ex.37 b) reflexive and transitive, but not symmetric.
		{1, 1, 1, 0},
		{0, 1, 0, 0},
		{0, 1, 1, 0},
		{0, 0, 0, 1},
	}
	x.cp04_04[3801] = boolMatrix{ // Ex.38 a) irreflexive and transitive, but not symmetric.
		{0, 1, 1, 0},
		{0, 0, 0, 0},
		{0, 1, 0, 0},
		{0, 0, 0, 0},
	}
	x.cp04_04[3802] = boolMatrix{ // Ex.38 b) antisymmetric and reflexive, but not transitive.
		{1, 0, 0, 0},
		{1, 1, 0, 0},
		{0, 1, 1, 0},
		{0, 0, 0, 1},
	}
	x.cp04_04[3901] = boolMatrix{ // Ex.39 a) transitive, reflexive, and symmetric.
		{1, 1, 0, 0},
		{1, 1, 0, 0},
		{0, 0, 1, 0},
		{0, 0, 0, 1},
	}
	x.cp04_04[3902] = boolMatrix{ // Ex.39 b) asymmetric and transitive.
		{0, 1, 0, 0},
		{0, 0, 0, 0},
		{0, 0, 0, 0},
		{0, 0, 0, 0},
	}
	boolMatrixData = x
}

type relationsTestData struct {
	test1         relations
	trans_sym_3   relations
	equivalence_4 relations
	cp04_04       map[int]relations
	cp04_05       map[int]relations
	integerFuncs  map[string]func(int, int) bool
}

func initRelationsTestData() {
	setA := newSet("1", "2", "3", "4")
	testRelationsData = relationsTestData{
		test1:         newRelations([]relation{{"a", "a"}, {"a", "b"}, {"c", "c"}, {"c", "d"}}),
		trans_sym_3:   newRelations([]relation{{"a", "a"}, {"a", "b"}, {"a", "c"}, {"b", "c"}, {"c", "c"}}),
		equivalence_4: newRelations([]relation{{"1", "1"}, {"1", "2"}, {"2", "1"}, {"2", "2"}, {"3", "4"}, {"4", "3"}, {"3", "3"}, {"4", "4"}}),
		cp04_04:       make(map[int]relations),
		cp04_05:       make(map[int]relations),
		integerFuncs:  make(map[string]func(int, int) bool),
	}
	testRelationsData.cp04_04[1] = newRelations([]relation{{"1", "1"}, {"1", "2"}, {"2", "1"}, {"2", "2"}, {"3", "3"}, {"3", "4"}, {"4", "3"}, {"4", "4"}})
	testRelationsData.cp04_04[2] = newRelations([]relation{{"1", "2"}, {"1", "3"}, {"1", "4"}, {"2", "3"}, {"2", "4"}, {"3", "4"}})
	testRelationsData.cp04_04[6] = setA.cartesianSquare()
	testRelationsData.cp04_04[8] = newRelations([]relation{{"1", "3"}, {"4", "2"}, {"2", "4"}, {"3", "1"}, {"2", "2"}})
	testRelationsData.cp04_04[108] = newRelations([]relation{{"4", "2"}, {"2", "4"}})
	testRelationsData.cp04_04[10] = newRelations([]relation{{"1", "2"}, {"1", "3"}, {"1", "4"}, {"5", "2"}, {"5", "3"}, {"5", "4"}})
	testRelationsData.cp04_04[11] = boolMatrixData.cp04_04[11].toRelations()
	testRelationsData.cp04_04[12] = boolMatrixData.cp04_04[12].toRelations()
	testRelationsData.cp04_04[13] = boolMatrixData.cp04_04[13].toRelations()
	testRelationsData.cp04_04[3701] = boolMatrixData.cp04_04[3701].toRelations()
	testRelationsData.cp04_04[3702] = boolMatrixData.cp04_04[3702].toRelations()
	testRelationsData.cp04_04[3801] = boolMatrixData.cp04_04[3801].toRelations()
	testRelationsData.cp04_04[3802] = boolMatrixData.cp04_04[3802].toRelations()
	testRelationsData.cp04_04[3901] = boolMatrixData.cp04_04[3901].toRelations()
	testRelationsData.cp04_04[3902] = boolMatrixData.cp04_04[3902].toRelations()
	testRelationsData.integerFuncs["cp04_04_13"] = func(from, to int) bool { return from <= (to + 1) }

	testRelationsData.cp04_05[5] = newRelations([]relation{{"a", "a"}, {"b", "a"}, {"b", "b"}, {"c", "c"}, {"d", "d"}, {"d", "c"}})
	testRelationsData.cp04_05[7] = newRelations([]relation{{"1", "1"}, {"1", "2"}, {"2", "1"}, {"2", "2"}, {"3", "1"}, {"3", "3"}, {"1", "3"}, {"4", "1"}, {"4", "4"}})

}
