# Functional Go

<!--- `Build this document with docker run -it -v (pwd):/build/ pandoc/latex -s /build/README.md -o /build/README.pdf` -->

See `thesis.pdf` for the actual thesis document. That itself links to `thesis/thesis.pdf`.

Structure and Files:

- `thesis`: contains the latex code plus a pre-compiled version of the document.
  There are also further tools in that directory that are needed to build the
  document.
- `work`: contains subfolders for practical work that was done.
  - `common-list-functions`: See the thesis' appendix on what it is.
    Contains git submodules and a shell script to acquire the results
    that are described in the thesis.
  - `examples`: contains different examples of functional Go code. Some code snippets
    are described in the thesis, some have been developed out of curiosity.
  - `funcheck`: git submodule that contains the implementation of `funcheck' as described
    in the thesis.
  - `go`: git submodule with the Go compiler code.
- `proposals`: contains old proposals that where submitted.

See Appendix 1 in `thesis.pdf` for more information.
