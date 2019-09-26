# hexreplace

A simple binary file hex replacer.

## usage

```
>hexreplace --help
Usage: hexreplace [OPTION]... SOURCE DEST
Replace hex in SOURCE, save as DEST.

OPTION:
  -n string
        new hex
  -o string
        old hex
```

A simple:
```
hexreplace -o 68007400740070003A002F002F006C006F00610064006D00610072006B0064006F0077006E002F0000000000 -n 68007400740070003A002F002F006C006F00630061006C0068006F00730074003A003500350035002F000000 Evernote.exe D:\Evernote.exe
```