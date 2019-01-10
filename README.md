# pixela-client-go

THIS IS VERY ALPHA VERSION. API COULD CHANGE.

Unofficial [pixe.la](https://pixe.la) API client & CLI by golang.


## Synopsys

```
# create user (authentication data stores $HOME/.pixela.yaml by default)
$ pixela user create <username> <token>

# create graph (default timezone is 'UTC')
$ pixela graph create <graph id> <graph name> <unit> <type> <color> [timezone]

# record quantity
$ pixela pixel record <graph id> <date> <quantity>
```


## Feature implement matrix

pixela-client-go now implement following feature.


|Target/Methods |User    |Graph   |Pixel   |Webhook |
|---------------|--------|--------|--------|--------|
|create         |**done**|**done**|**done**|not yet |
|get(definition)|N/A     |not yet |N/A     |N/A     |
|get(data)      |N/A     |not yet |**done**|not yet |
|update         |not yet |not yet |not yet |N/A     |
|update(inc)    |N/A     |N/A     |not yet |N/A     |
|update(dec)    |N/A     |N/A     |not yet |N/A     |
|delete         |not yet |not yet |not yet |not yet |
|invoke         |N/A     |N/A     |N/A     |not yet |


## Author

[noissefnoc](noissefnoc@gmail.com)