// Package initiate This package allows us to initiate Time Sensitive components (Like registering the windows service) as early as possible in the startup process
package initiate


var (
    StopCh = make(chan bool)
)

