%%
%% This is file `zhawthesis.cls',
%% generated with the docstrip utility.
%%
%% The original source files were:
%%
%% zhawthesis.dtx  (with options: `class')
%% ----------------------------------------------------------------
%% zhawthesis --- A LaTeX class for writing a thesis at ZHAW.
%% E-mail: camerden@students.zhaw.ch
%% Released under the LaTeX Project Public License v1.3c or later
%% See http://www.latex-project.org/lppl.txt
%% ----------------------------------------------------------------
%% 
\NeedsTeXFormat{LaTeX2e}
\ProvidesClass{zhawthesis}[2019/02/28 v1.0 Initial version]
%% ========================================================================== %%
%%                                 Thesis class                               %%
%% ========================================================================== %%

\def\zhaw@babellangde{nswissgerman}
\def\zhaw@babellangen{english}

\newcommand{\zhaw@lang}{de}
\newcommand{\zhaw@babellang}{\zhaw@langde}
\newif\if@german\@germantrue

\newcommand{\zhaw@thesistype}{}
\newcommand{\zhaw@thesistypelong}{}

\DeclareOption{ba}{
  \renewcommand{\zhaw@thesistype}{BA}
  \renewcommand{\zhaw@thesistypelong}{%
    \if@german{Bachelorarbeit}\else{Bachelor thesis}\fi%
  }
}
\DeclareOption{pa}{
  \renewcommand{\zhaw@thesistype}{PA}
  \renewcommand{\zhaw@thesistypelong}{%
    \if@german{Projektarbeit}\else{Project work}\fi%
  }
}

\DeclareOption{german}{
  \@germantrue
  \renewcommand{\zhaw@lang}{de}
  \renewcommand{\zhaw@babellang}{\zhaw@babellangde}
  \PassOptionsToClass{\zhaw@babellang}{scrbook}
}
\DeclareOption{english}{
  \@germanfalse
  \renewcommand{\zhaw@lang}{en}
  \renewcommand{\zhaw@babellang}{\zhaw@babellangen}
  \PassOptionsToClass{\zhaw@babellang}{scrbook}
}

\DeclareOption*{\PassOptionsToClass{\CurrentOption}{scrbook}}
\ProcessOptions

\RequirePackage[T1]{fontenc}
\RequirePackage[utf8]{inputenc}

\LoadClass[listof=totoc,bibliography=totoc]{scrbook}

\RequirePackage{microtype}
\RequirePackage{graphicx}

\if@german
  \RequirePackage[\zhaw@babellangen,main=\zhaw@babellangde]{babel}
\else
  \RequirePackage[\zhaw@babellangde,main=\zhaw@babellangen]{babel}
\fi
\babeltags{de=\zhaw@babellangde,en=\zhaw@babellangen}
\RequirePackage{iflang}

\RequirePackage[autostyle=try,strict=true,german=swiss,english=british]{csquotes}

\RequirePackage[output-decimal-marker={.},group-separator={'}]{siunitx}
\sisetup{detect-all}

\RequirePackage[pdfpagelabels]{hyperref}
\hypersetup{%
  unicode,
  pdfauthor={\@author},
  pdftitle={\@title},
  pdfsubject={\zhaw@thesistypelong},
  pdfkeywords={Thesis;LaTeX},  % TODO: Complete
  pdfcreator={pdfLaTeX},
  pdfduplex={DuplexFlipLongEdge},
  pdflang={\zhaw@lang},
  bookmarksopen,
  bookmarksnumbered
}

%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%

\RequirePackage[automark]{scrlayer-scrpage}
\automark[chapter]{chapter}
\clearpairofpagestyles
\lehead{\@title}
\rohead{\@author}
\ofoot[\pagemark]{\pagemark}

%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%

\newcommand{\zhaw@institute}{}
\providecommand{\zhawinstitute}[1]{\renewcommand{\zhaw@institute}{#1}}

\newcommand{\zhaw@logocolour}{pantone2945}
\providecommand{\zhawlogocolour}[1]{\renewcommand{\zhaw@logocolour}{#1}}

\newcommand{\zhaw@major}{}
\providecommand{\major}[1]{\renewcommand{\zhaw@major}{#1}}

\newcommand{\zhaw@semester}{}
\providecommand{\zhawsemester}[1]{\renewcommand{\zhaw@semester}{#1}}

\newcommand{\zhaw@mainsupervisor}{}
\providecommand{\mainsupervisor}[1]{\renewcommand{\zhaw@mainsupervisor}{#1}}

\newcommand{\zhaw@subsupervisor}{}
\providecommand{\subsupervisor}[1]{\renewcommand{\zhaw@subsupervisor}{#1}}

\newcommand\zhaw@industrypartner{}
\providecommand{\industrypartner}[1]{\renewcommand{\zhaw@industrypartner}{#1}}

\newcommand{\zhaw@extsupervisor}{}
\providecommand{\externalsupervisor}[1]{\renewcommand{\zhaw@extsupervisor}{#1}}

\newcommand{\zhaw@titlepagerow}[2]{%
  \begin{minipage}[t]{0.3\textwidth}
    \hrule\vskip 5mm
    \textbf{#1}
  \end{minipage}
  % \hskip 0.03\textwidth
  \hfil
  \begin{minipage}[t]{0.67\textwidth}
    \hrule\vskip 5mm
    #2 \\
  \end{minipage}%
}

\newcommand{\zhaw@maketitlepage}{
  \thispagestyle{empty}
  \clearpage
  \begin{titlepage}
    \sffamily
    % Logo
    \begin{picture}(0,0)
      \put(-30,-50){
        \includegraphics[width=84.2mm]{logos/zhaw/\zhaw@lang-zhaw-\zhaw@institute-\zhaw@logocolour}
      }
    \end{picture}

    \vskip 10mm
    \hskip 26mm
    \begin{minipage}[b]{0.91\textwidth}
      \vskip 20mm
      {\huge
        % Projekt Name (max. 2 Zeilen)
        \textbf{\underline{{\zhaw@thesistypelong} in {\zhaw@major}}} \\
        \textbf{\underline{{\zhaw@semester}}} \\

        % Projekt Titel (max. 4 Zeilen)
        {
          \huge{\@title} \\[1.25ex]
          \large{\@subtitle}
        }
        \vspace{9mm}
      }

      \zhaw@titlepagerow{\IfLanguageName{nswissgerman}{Autoren}{Authors}}{%
        \@author \\
      }

      \if\zhaw@mainsupervisor\empty\else
      \zhaw@titlepagerow{\IfLanguageName{nswissgerman}{Hauptbetreuung}{Main supervisor}}{%
        \zhaw@mainsupervisor \\
      }
      \fi

      \if\zhaw@subsupervisor\empty\else
      \zhaw@titlepagerow{\IfLanguageName{nswissgerman}{Nebenbetreuung}{Sub supervisor}}{%
        \zhaw@subsupervisor \\
      }
      \fi

      \if\zhaw@industrypartner\empty\else
      \zhaw@titlepagerow{\IfLanguageName{nswissgerman}{Industriepartner}{Industrial partner}}{%
        \zhaw@industrypartner \\
      }
      \fi

      \if\zhaw@extsupervisor\empty\else
      \zhaw@titlepagerow{\IfLanguageName{nswissgerman}{Externe Betreuung}{External supervisor}}{%
        \zhaw@extsupervisor \\
      }
      \fi

      \zhaw@titlepagerow{\IfLanguageName{nswissgerman}{Datum}{Date}}{%
        \@date \\
      }
    \end{minipage}
    \vskip 5mm
  \end{titlepage}
}
\renewcommand{\maketitle}{\zhaw@maketitlepage}

\RequirePackage{pdfpages}
\providecommand{\makedeclarationoforiginality}{%
  \cleardoublepage%
  \if@german%
    \includepdf{includes/Erklaerung_Selbstaendigkeit_SoE_\zhaw@thesistype_de.pdf}
  \else%
    \includepdf{includes/Declaration_of_Originality_SoE_\zhaw@thesistype_en.pdf}
  \fi%
}

\AtBeginDocument{%
  \pagestyle{scrheadings}

  % Set 1st level itemize bullet symbol to --
  \def\labelitemi{--}
}

%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%

\providecommand\norm[1]{\left\lVert#1\right\rVert}

\def\signed #1{{\leavevmode\unskip\nobreak\hfil\penalty50\hskip2em
  \hbox{}\nobreak\hfil(#1)%
  \parfillskip=0pt \finalhyphendemerits=0 \endgraf}}

%% 
%% Copyright (C) 2019 by Dennis Camera <camerden@students.zhaw.ch>
%% 
%% This work may be distributed and/or modified under the
%% conditions of the LaTeX Project Public License (LPPL), either
%% version 1.3c of this license or (at your option) any later
%% version.  The latest version of this license is in the file:
%% 
%% http://www.latex-project.org/lppl.txt
%% 
%% This work is "maintained" (as per LPPL maintenance status).
%% 
%% This work consists of the file  zhawthesis.dtx
%% and the derived files           zhawthesis.ins,
%%                                 zhawthesis.pdf and
%%                                 zhawthesis.cls.
%% 
%%
%% End of file `zhawthesis.cls'.
