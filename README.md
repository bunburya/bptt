# `ltt`

`ltt` (**L**ondon **T**ransport in the **T**erminal) is a simple command line application that displays information
about various London public transport services.

## Line status

This command displays the current service status of each given line.  Where the TfL API reports a number of different
statuses, the most severe one is reported.

The general form is:
```
ltt status [--color] LINE_ID LINE_ID [...]
```

In the above example, `LINE_ID` is the id of the line in the form used by the TfL API. For example, the Northern line
is `northern`, the Hammersmith & City line is `hammersmith-city`. Bus route IDs are generally just the route name, with
any letters being lowercased.

For example:
```
ltt status bakerloo northern victoria 470 rb1 london-cable-car
```

will produce something like:

```
470                     Good Service
Bakerloo                Minor Delays
IFS Cloud Cable Car     Good Service
Northern                Good Service
RB1                     Good Service
Victoria                Good Service
```

The optional `--color` flag will use colour in the output.

## Arrivals

This command displays real time arrival predictions for the given stop. 

**NOTE:** Because this command displays arrivals rather than departures, if the stop you are checking is the stop from
which the given service starts, nothing will be reported. This is a limitation of the TfL API:
https://techforum.tfl.gov.uk/t/how-to-find-departures-from-terminal-stations/72/26.

Format:
```
ltt arrivals [--count N] [--lines LINE_ID,LINE_ID,[...]] STOP_ID
```

In the above example, `STOP_ID` is the [NaPTAN ID](https://beta-naptan.dft.gov.uk/) of the relevant stop. The optional
`--lines` argument (a comma-separated list of line IDs) will display only arrivals for the given lines. The `--count`
argument, if provided, will display at most the given number of arrivals.

For example:
```
ltt arrivals --count 5 490000254D
```

will produce something like:

```
59      Clapham Park, Atkins Road       2m54s 
172     Brockley Rise                   5m39s 
68      West Norwood                    10m22s
59      Clapham Park, Atkins Road       13m4s 
176     Penge                           16m58s
```