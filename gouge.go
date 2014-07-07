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

func (self *page) element(name string, content string){
    self.content += "<" + name + ">" + content + "</" + name + ">"
}

