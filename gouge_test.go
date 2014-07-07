package gouge

import(
	"testing"
	"reflect"
)

// Test Helpers
func expect(t *testing.T, a interface{}, b interface{}) {
	if a != b {
		t.Errorf("Expected %v (type %v) - Got %v (type %v)", b, reflect.TypeOf(b), a, reflect.TypeOf(a))
	}
}

func refute(t *testing.T, a interface{}, b interface{}) {
	if a == b {
		t.Errorf("Did not expect %v (type %v) - Got %v (type %v)", b, reflect.TypeOf(b), a, reflect.TypeOf(a))
	}
}

func TestWrite(t *testing.T){
	result := template(func(self *page){
		self.write("hello world!")
	})

	expect(t, result, "hello world!")
}

func TestElement(t *testing.T){
	result := template(func(self *page){
		self.element("html", "hello world!")
	})
	
	expect(t, result, "<html>hello world!</html>")
}

func TestElementWithContentMaker(t* testing.T) {
	result := template(func(self *page){
		self.element("html", func(self *page){
			self.write("hello world!")
		})
	})

	expect(t, result, "<html>hello world!</html>")
}
