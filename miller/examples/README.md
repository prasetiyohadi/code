# Examples

This directory contains example files for miller from https://miller.readthedocs.io/en/latest/10min/

## Example commands

```bash
Example ```

```bash
mlr --csv cat example.csv
mlr --csv head -n 4 example.csv
mlr --csv tail -n 4 example.csv
mlr --icsv --ojson tail -n 2 example.csv
mlr --icsv --opprint sort -f shape example.csv
mlr --icsv --opprint sort -f shape -nr index example.csv
mlr --icsv --opprint cut f flag,shape example.csv
mlr --icsv --opprint cut -f flag,shape example.csv
mlr --icsv --opprint cut -o -f flag,shape example.csv
mlr --icsv --opprint cut -x -f flag,shape example.csv
mlr --icsv --opprint put '$[[3]] = "NEW"' example.csv
mlr --icsv --opprint put '$[[[3]]] = "NEW"' example.csv
mlr --icsv --opprint filter '$color == "red"' example.csv
mlr --icsv --opprint filter '$color == "red" && $flag == "true"' example.csv
mlr --icsv --opprint put '\n$ratio = $quantity / $rate;\n$color_shape = $color . "_" . $shape\n' example.csv
mlr --icsv --opprint --from example.csv put '\n$y = $index + 1;\n$z = $y**2 + $k;\n'
mlr --csv cat spaces.csv
mlr --c2p sort -nr 'Total MWh' spaces.csv
mlr --c2p put '${Total KWh} = ${Total MWh} * 1000' spaces.csv
mlr --csv cat data/a.csv data/b.csv
mlr --csv sort -nr index example.csv | mlr --icsv --opprint head -n 3
mlr --csv sort -nr index then head -n 3 example.csv
mlr --icsv --opprint --from example.csv \\nsort -nr index \\nthen head -n 3 \\nthen cut -f shape,quantity
mlr --icsv --opprint sort -nr index then head -n 3 example.csv
mlr --icsv --opprint sort -f shape -nr index then head -n 1 -g shape example.csv
mlr --icsv --opprint --from example.csv \\nstats1 -a count,min,mean,max -f quantity -g shape
mlr --icsv --opprint --from example.csv \\nstats1 -a count,min,mean,max -f quantity -g shape,color
mlr --icsv --oxtab --from example.csv \\nstats1 -a p0,p10,p25,p50,p75,p90,p99,p100 -f rate
mlr --icsv --opprint cat example.csv
```
