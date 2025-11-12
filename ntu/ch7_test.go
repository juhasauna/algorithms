package ntu

import (
	"algorithms/ut"
	"testing"
)

func Test_ch7(t *testing.T) {
	tests := []struct {
		name string
		f    func(*testing.T)
	}{
		{"FindEulerianCircuitPseudo", FindEulerianCircuitPseudoTest},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.f(t)
		})
	}
}

func FindEulerianCircuitPseudoTest(t *testing.T) {
	tests := []struct {
		graph *ut.EulerianGraph
	}{
		// {ut.GetEulerianGraphK3()},
		{ut.GetEulerianGraphK5()},
		// {ut.GetEulerianGraphK5_a()},
		// {ut.GetEulerianGraphBowtie()},

		// {ut.GetNotEulerianGraphK5_a()},
	}
	for _, tt := range tests {
		x := CH7{}
		gotPath, gotEdges := x.FindEulerianCircuitPseudo(tt.graph)

		gotPathLen := len(gotPath)
		gotEdgesLen := len(gotEdges)
		wantLen := tt.graph.EdgeCount()
		if gotPathLen != gotEdgesLen {
			t.Errorf("FAIL RESULT LENGTHS(%d/%d)\ngotPath: %v\ngotEdges: %v", gotPathLen, gotEdgesLen, gotPath, gotEdges)
		} else if gotPathLen != wantLen {
			t.Errorf("FAIL gotPath: %v\ngotEdges: %v, wantLen: %d, gotLen:%d", gotPath, gotEdges, wantLen, gotPathLen)
		} else if err := ut.IsValidEulerianCircuit(gotPath); err != nil {
			t.Errorf("FAIL gotPath: %v\ngotEdges: %v\n%v", gotPath, gotEdges, err)
		} else {
			t.Logf("SUCCESS gotPath: %v\ngotEdges: %v", gotPath, gotEdges)
		}
		hierGot := tt.graph.HierholzerFirst()
		t.Logf("%+v", hierGot)
	}
}
