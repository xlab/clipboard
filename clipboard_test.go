package clipboard

import (
	"testing"

	"gopkg.in/qml.v1"
)

func init() {
	qml.SetupTesting()
}

func TestCopyAndPaste(t *testing.T) {
	engine := qml.NewEngine()
	clip := New(engine)
	//======================
	a := "Привет, мир!"

	err := clip.WriteAll(a)
	if err != nil {
		t.Fatal(err)
	}

	b, err := clip.ReadAll()
	if err != nil {
		t.Fatal(err)
	}

	if a != b {
		t.Errorf("expected '%s', got '%s'", a, b)
	}
}

func BenchmarkReadAll(b *testing.B) {
	engine := qml.NewEngine()
	clip := New(engine)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		clip.ReadAll()
	}
}

func BenchmarkWriteAll(b *testing.B) {
	engine := qml.NewEngine()
	clip := New(engine)
	b.ResetTimer()
	text := "Жарю сосиски"
	for i := 0; i < b.N; i++ {
		clip.WriteAll(text)
	}
}
