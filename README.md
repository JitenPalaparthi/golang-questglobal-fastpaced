# Go Channels

- close 
- range loop for a channel
- select 

- only sender can close a channel 
- upon closing a channel , it not only closes the channel but also it sends a value.It sends the default value of the type (according to the type inference of go)
- range loop on a channel , it receives the value until the channel is closed
- when multiple go routines are trying to fetch values from a channel (buffered), the order is not guaranteed 