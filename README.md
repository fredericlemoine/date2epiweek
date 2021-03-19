# date2epiweek

A simple command that uses [epiweek](https://github.com/jmeekhof/epiweek).

It takes a metadata file as input, a separator, the date column index, and prints the epi week for each line.

```bash
Usage of ./date2epiweek:
  -column int
    	Column with date (YYYY-MM-DD)
  -header
    	If input has a header
  -help
    	help
  -metadata string
    	Meta data file (default "stdin")
  -sep string
    	Column seprator, default tab (default "\t")
```
