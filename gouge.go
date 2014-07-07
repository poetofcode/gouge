package gouge

type page struct {
    content string
}

func template(content_maker func(*page)) string{
    self := &page{}
    
    content_maker(self)

    return self.content
}

func (self *page) div(content string){
    self.content += "<div>" + content + "</div>"    
}

func (self *page) element(name string, subcontent interface{}){
	switch subcontent.(type) {
		case string:
			self.content += "<" + name + ">" + subcontent.(string) + "</" + name + ">"
		case func(*page):
			self.content += "<" + name + ">"
			subcontent.(func(* page))(self)
			self.content += "</" + name + ">"
		default:
	}
}

func (self *page) write(content string){
	self.content += content
}
