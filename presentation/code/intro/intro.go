package main

import "fmt"

// START OMIT
func main() {
	p := &presentation{
		slides: make([]slide, 0, 10),
		author: "Ramon",
	}
	p.slides = append(p.slides, newIntroSlide(p.author))
	fmt.Println(p)
}

type presentation struct {
	slides []slide
	author string
}

func (p *presentation) String() string {
	return fmt.Sprintf("Author: %v, Number of Slides: %v",
		p.author, len(p.slides))
}

func newIntroSlide(author string) slide {
	return slide{text: "Author: " + author}
}

// END OMIT

type slide struct {
	text string
}
