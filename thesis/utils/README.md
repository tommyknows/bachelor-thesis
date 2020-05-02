# Delim

This is a small helper utility for building the latex document.

It is called by [thesis.tex](../thesis.tex) with the following `newcommand`:

```tex
\newcommand{\gofilerange}[4][]{%
    \immediate\write18{./utils/delim -file="#2" -start="#3" -end="#4"}
    \IfFileExists{code.lineno}
      {\CatchFileEdef{\linenumber}{./code.lineno}{\endlinechar=-1 }}
      {\def\linenumber{0}}
    \edef\flags{firstnumber=\linenumber,#1}
    \expandafter\gofile\expandafter[\flags]{./code.snippet}
}
```

Basically, `delim` receives the filename and a start- and end-comment.
It then writes the files content between the two comments into `code.snippet`,
and the starting line number into `code.lineno`.

These are then used by the `minted` package to display the code section
within the document.

Use it in the document like so:

```tex
\gofilerange{this/file.go}{start-section}{end-section}
```
