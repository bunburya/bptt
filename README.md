# `ltt`

`ltt` (**L**ondon **T**ransport in the **T**erminal) is a simple command line application that displays information
about various London public transport services.

It is divided into sub-commands. The currently supported commands are:

- `tfl status`: View the status of TfL lines.
- `tfl arrivals`: View next arrivals at TfL stop.
- `nre departures`: View next departures at a National Rail station.

## `tfl status`

This command displays the current service status of each given line.  Where the TfL API reports a number of different
statuses, the most severe one is reported.

The general form is:
```
ltt tfl status LINE_ID...
```

In the above example, `LINE_ID` is the ID of the line in the form used by the TfL API. For example, the Northern line
is `northern`, the Hammersmith & City line is `hammersmith-city`. Bus route IDs are generally just the route name, with
any letters being lowercased. Multiple line IDs can be provided.

For example:
```
$ ltt tfl status bakerloo northern victoria 470 rb1 london-cable-car
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

This command supports coloured output.

## `tfl arrivals`

This command displays real time arrival predictions for the given stop. 

**NOTE:** Because this command displays arrivals rather than departures, if the stop you are checking is the stop from
which the given service starts, nothing will be reported. This is a limitation of the TfL API:
https://techforum.tfl.gov.uk/t/how-to-find-departures-from-terminal-stations/72/26.

Format:
```
ltt tfl arrivals [--count N] [--lines LINE_ID,LINE_ID,[...]] STOP_ID
```

In the above example, `STOP_ID` is the [NaPTAN ID](https://beta-naptan.dft.gov.uk/) of the relevant stop. The optional
`--lines` argument (a comma-separated list of line IDs) will display only arrivals for the given lines. The `--count`
argument, if provided, will display at most the given number of arrivals. By default, the next five arrivals will be
displayed.

For example:
```
$ ltt tfl arrivals --count 5 490000254D
```

will produce something like:

```
59      Clapham Park, Atkins Road       2m54s 
172     Brockley Rise                   5m39s 
68      West Norwood                    10m22s
59      Clapham Park, Atkins Road       13m4s 
176     Penge                           16m58s
```

## `nre departures`

This command displays real time departure predictions for the given National Rail station.

Format:
```
ltt nre departures STATION_ID
```

In the above example, `STATION_ID` is the CRS code of the relevant station.

For example:
```
$ ltt nre departures --token [api_token_here] KGX
```

will produce something like:

```
Leeds           23:26   On time
Cambridge       23:27   On time
Cambridge       00:02   On time
Cambridge       00:32   On time
Royston         00:36   On time
Royston         01:07   On time
Finsbury Park   01:10   On time
```