# `ptt`

`ptt` (**P**ublic **T**ransport in the **T**erminal) is a simple command line application that displays information
about various public transport services.

It is divided into commands and sub-commands. The currently supported top-level commands are:

| Command | Sub-command  | Purpose                                                                                                                                                                                                                                                                                                                                             | Arguments                                                                         |
|---------|--------------|-----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|-----------------------------------------------------------------------------------|
| `tfl`   | `status`     | View the status of one or more specified TfL lines.                                                                                                                                                                                                                                                                                                 | A comma-separated list of line IDs.                                               |
| `tfl`   | `modestatus` | View the status of all lines of one or more specified TfL transport modes.                                                                                                                                                                                                                                                                          | A comma-separated list mode IDs.                                                  |
| `tfl`   | `arrivals`   | View the next arrivals at the given stop. **NOTE:** Because this command displays arrivals rather than departures, if the stop you are checking is the stop from which the given service starts, nothing will be reported. This is a limitation of the TfL API: https://techforum.tfl.gov.uk/t/how-to-find-departures-from-terminal-stations/72/26. | The NaPTAN ID of the relevant stop.                                               |
| `tfl`   | `bikes`      | Search bike availability at the specified Santander Cycles station.                                                                                                                                                                                                                                                                                 | The ID of the relevant bike point.                                                |
| `tfl`   | `search`     | Search for information about TfL modes, lines and stations. This command is divided into sub-(sub-)commands and can be used to get the relevant ID values to use for the other commands.                                                                                                                                                            | Call with the `-h` flag for more information on sub-commands and their arguments. |
| `nre`   | `departures` | View next departures at the specified National Rail station.                                                                                                                                                                                                                                                                                        | The CRS code of the relevant station.                                             |
| `waqi`  |              | View the Air Quality Index for the given city.                                                                                                                                                                                                                                                                                                      | The name of a city.                                                               |

## Input ID codes

The `tfl` commands generally take line, mode, stop or bike station IDs in the form used by the TfL API.  Some
examples of what these ID look like:

- Mode: `tube`, `bus`, `river-bus`, etc. Use `ptt tfl search mode` to see all modes supported by the TfL API (but note that some modes, such as `walking`, won't work with commands like `modestatus`).
- Line: `northern`, `waterloo-city`, `sl4`, `suffragette`, etc. Use `tfl search line <mode-id>` to see all line IDs for the given mode. 
- Stop: These are [NaPTAN IDs](https://www.data.gov.uk/dataset/ff93ffc1-6656-47d8-9155-85ea0b8f2251/naptan), eg, `490011218B`. There are a few ways to access these:
  - You can download a full list from the above link.
  - You can view the stop on [OpenStreetMap](https://www.openstreetmap.org) and look for the `naptan:AtcoCode` tag.
  - You can try `ptt tfl search stop <search string>`. However, note that often that command (like the API endpoint it calls) will only return the NaPTAN ID for a "hub", which is a collection of stops. Hubs themselves (as opposed to the stops) will not yield any arrivals data, and the stops won't always be returned in the data.
- Bike point: These are of the form `BikePoints_<n>`, where `<n>` is an integer, eg, `BikePoints_427`. Use `ptt search bike <search string>` to search for relevant bike point IDs.

The `nre departures` command takes the CRS code (which must be uppercase) of the relevant rail station, eg, `KGX` for King's Cross.
There isn't currently a way to search for CRS codes from within `ptt` but you can look them up at http://www.railwaycodes.org.uk/crs/crs0.shtm.

For convenience you can set aliases for stop IDs, bike point IDs and CRS codes in the configuration file.

## API keys

`ptt` is just a simple interface to a number of APIs. These APIs generally take API keys. For the TfL API, a key is
optional but recommended. Currently the API will serve requests without an API key, but there is no guarantee this will
continue or you won't get rate limited. The other APIs require API keys. You can get these at:

- Transport for London: https://api-portal.tfl.gov.uk/
- National Rail Enquiries: https://raildata.org.uk
- Air Quality Index: https://aqicn.org/data-platform/token/

The relevant API key can be passed using the `--api-key` flag or can be specified in the config file.

`ptt` does not perform any caching of data so it is up to the user to ensure it is used in compliance with the
applicable API terms.

## Configuration

The behaviour of `ptt` can be configured by command line flags or though a TOML-based configuration file. You can always
call `ptt` or its sub-commands with `-h` to view a list of supported flags.

You can point `ptt` to a specific config file location by passing the `--config` flag. Otherwise, it will use the
[ConfigDir](https://github.com/kirsle/configdir) library to determine the "normal" place to look for a config file (for
example, on Linux it will look for `$HOME/.config/ptt/config.toml`).

The repo contains a documented example `config.toml` file.

## Output

`ptt` will output the desired data in a table format. Spaces are used for padding and tabs to separate columns, so it
works best in a terminal environment using a monospace font. 

Passing the `--color` flag will use color in the output, where appropriate. For example, bad news (disrupted status,
late services, etc) may be displayed in yellow or red, whereas good news may be displayed in green. This uses ANSI
escape codes to format the output, so again, it works best in a terminal that supports such escape codes.

The `--col-size` flag allows you to fix the size of each column in the output. It should be a comma-separated list of
integers. Each integer is the width in terminal characters of the corresponding column. If the text to be displayed in a
column is longer than the desired width, it will be cut off, as indicated by a "…" character. A value of `0` for a
column hides that column completely.

Any negative value means there is no fixed width for that column (the width of the column will be the length of the
longest piece of text in that column, which is also the behaviour if the `--col-size` flag is not set). Similarly, if
fewer integers are passed via the `--col-size` than there are columns in the output, any extra columns in the output
will be treated as unfixed.

Technically, column widths can also be set in `config.toml` using the `column_size` setting, but this can only be set
globally (rather than per-command) so it is not really recommended to do it this way unless you will only ever be
calling the same command (as different commands will output different numbers of columns).

Remember that tabs are used to separate columns, so column `n` won't always start one character after the end of column
`n-1`.

Passing the `--header` flag will include a header row in the output with a short name for each column. Passing the
`--timestamp` flag will include a "Last updated" timestamp at the bottom of the output.

## Some random examples

```
$ ptt tfl modestatus tube
Bakerloo          	Good Service
Central           	Good Service
Circle            	Good Service
District          	Good Service
Hammersmith & City	Minor Delays
Jubilee           	Good Service
Metropolitan      	Good Service
Northern          	Good Service
Piccadilly        	Good Service
Victoria          	Good Service
Waterloo & City   	Good Service
```

```
$ ptt --header tfl status 470 rb1 london-cable-car
Line               	Status      
470                	Good Service
IFS Cloud Cable Car	Good Service
RB1                	Good Service
```

```
$ ptt tfl search bike southwark --header
Name                          	ID            	Latitude 	Longitude
Webber Street , Southwark     	BikePoints_80 	51.500693	-0.102091
Colombo Street, Southwark     	BikePoints_240	51.505459	-0.105692
Southwark Station 2, Southwark	BikePoints_421	51.504044	-0.104778
Blackfriars Road, Southwark   	BikePoints_792	51.505461	-0.104540
Southwark Street, Bankside    	BikePoints_803	51.505409	-0.098341
```

```
$ ptt --header tfl bikes BikePoints_240
Station                  	Bikes	E-Bikes	Empty docks
Colombo Street, Southwark	7    	1      	6          
```

```
$ go run . tfl arrivals 490008275H --col-size -1,9,-1
17      Archway         34s   
46      Paddingt…       4m36s 
17      Archway         14m54s
46      Paddingt…       19m44s
17      Archway         30m6s 
```

```
$ ptt waqi London --header
AQI	Description
9  	Good 
```