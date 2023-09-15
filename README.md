# Go Channels

- close 
- range loop for a channel
- select 

- only sender can close a channel 
- range loop on a channel , it receives the value until the channel is closed
- when multiple go routines are trying to fetch values from a channel (buffered), the order is not guaranteed 