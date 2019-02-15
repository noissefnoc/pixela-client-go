# Changelog

## [Unreleased]

### Added

* User input validation according to official document validation rule.
* `graph create` subcommand now detect `timezone` as flag option.
* `user create` subcommand now detect `agreeTermsOfService` and `notMinor` as flag option.
* `pixel create` and `pixel update` accept `optionalData` flag.
* `graph svg` accept `date` and `mode` flag.
* `graph create` accept `selfSufficient` flag.
 `graph update` accept `selfSufficient` flag.


## [0.0.1] - 2019-01-14

### Added

* Basic API support (does not full support optional/advanced usage).
    * UNSUPPORTED: `optionalData` at `pixel create` and `pixel update`