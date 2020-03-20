LATEXMK ?= latexmk
PDFLATEX ?= pdflatex
PLANTUML ?= plantuml
MAIN_TEX ?= ./thesis.tex

MKARGS = -interaction=nonstopmode -use-make -shell-escape

default: pdf

dvi: $(MAIN_TEX:.tex=.dvi)
pdf: $(MAIN_TEX:.tex=.pdf)

# Rules to compile zhawthesis document class
zhawthesis.cls: zhawthesis.dtx
	'$(PDFLATEX)' zhawthesis.dtx

# Rule to compile glossary
thesis.gls:
	makeglossaries thesis
skins.sty:



# Dependencies of thesis TeX
-include $(MAIN_TEX).d

# Thesis document creation rules (DVI, PDF)
$(MAIN_TEX:.tex=.dvi): $(MAIN_TEX)
	'$(LATEXMK)' -dvi -pdf- -ps- -deps-out='$<.d' $(MKARGS) '$<'

$(MAIN_TEX:.tex=.pdf): $(MAIN_TEX)
	'$(LATEXMK)' -pdf -dvi- -ps- -deps-out='$<.d' $(MKARGS) '$<'

# Continuous recompilation when the source files change
watch:
	'$(LATEXMK)' -pvc -pdf -dvi- -ps- $(MKARGS) '$(MAIN_TEX)'

# Open the compiled pdf
view: $(MAIN_TEX:.tex=.pdf)
	'$(shell command -v xdg-open open | head -n 1)' '$<' >/dev/null 2>&1 &

# Clean up temprary files
clean:
	'$(LATEXMK)' -C '$(MAIN_TEX)'
	$(RM) zhawthesis.aux zhawthesis.glo zhawthesis.hd zhawthesis.idx zhawthesis.log zhawthesis.out
	$(RM) *.synctex.gz *.bbl *.ilg
	$(RM) '$(MAIN_TEX).d'

.PHONY: clean dvi pdf view watch


# Resource compilation rules
%-%.pdf: %.mp
	cd '$(dir $(@))' && mpost '$(realpath $(@:.pdf=.mp))'

%.eps: %.plantuml
	$(PLANTUML) -teps -v '$<'

%.pdf: %.eps
	epstopdf -o='$@' '$<'

%.pdf: %.svg
	rsvg-convert -f pdf -o '$@' -d 300 -p 300 '$<'


logos/zhaw/%.pdf: logos/zhaw/%.ai
	inkscape --without-gui --export-file='$@' --export-area-drawing --export-margin=1 '$<'
