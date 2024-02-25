package logger

// LoggerInterface is a custom logger interface

type LoggerInterface interface {
	Print(v ...interface{})
	Printf(format string, v ...interface{})
	Println(v ...interface{})
	Fatal(v ...interface{})
	Fatalf(format string, v ...interface{})
	Fatalln(v ...interface{})
}

// FakeLogger is a fake implementation of LoggerInterface
type FakeLogger struct {
	// You can add fields here to track logs or other information
}

func (f *FakeLogger) Print(v ...interface{}) {
	// Implement the Print method as needed
}

func (f *FakeLogger) Printf(format string, v ...interface{}) {
	// Implement the Printf method as needed
}

func (f *FakeLogger) Println(v ...interface{}) {
	// Implement the Println method as needed
}

func (f *FakeLogger) Fatal(v ...interface{}) {
	// Implement the Fatal method as needed
}

func (f *FakeLogger) Fatalf(format string, v ...interface{}) {
	// Implement the Fatalf method as needed
}

func (f *FakeLogger) Fatalln(v ...interface{}) {
	// Implement the Fatalln method as needed
}

// NewFakeLogger creates a new instance of FakeLogger
func NewFakeLogger() LoggerInterface {
	return &FakeLogger{}
}
