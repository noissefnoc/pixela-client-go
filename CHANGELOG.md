# Changelog

## [0.0.6] - 2019-04-21

### Added

* update according to [Pixela v1.10.0 release](https://github.com/a-know/Pixela/releases/tag/v1.10.0)
    * `graph stat` subcommand

## [0.0.5] - 2019-04-13

* update according to [Pixela v1.9.0 release](https://github.com/a-know/Pixela/releases/tag/v1.9.0)
    * `graph svg --mode` accepts new mode `line` that returns line chart. 

## [0.0.4] - 2019-04-06

### Added

* `graph detail` subcommand

### Changed

* adjust subcommand to`pi` command (official CLI)
    * `pixel create` to `pixel post`
    * `graph def` to `graph get`
    * `graph inc` to `graph increment`
    * `graph dec` to `graph decrement`


## [0.0.3] - 2019-03-06

### Added

* `graph create` accept `selfSufficient` flag.
* `graph update` accept `selfSufficient` flag.
* `graph pixels` subcommand
*  now install via `homebrew` 


## [0.0.2] - 2019-01-28

### Added

* User input validation according to official document validation rule.
* `graph create` subcommand now detect `timezone` as flag option.
* `user create` subcommand now detect `agreeTermsOfService` and `notMinor` as flag option.
* `pixel create` and `pixel update` accept `optionalData` flag.
* `graph svg` accept `date` and `mode` flag.


## [0.0.1] - 2019-01-14

### Added

* Basic API support (does not full support optional/advanced usage).
    * UNSUPPORTED: `optionalData` at `pixel create` and `pixel update`