package app

type Event struct {
	Message  string
	Name     string
	Foo      int
	Bar      int
	SubEvent SubEvent
}

type SubEvent struct {
	BarBar string
	FooFoo SubSubEvent
}

type SubSubEvent struct {
	FooFooFoo int
}
