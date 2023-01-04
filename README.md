# OTN (One Time Note)

OTN (One Time Note) is a Golang microservice to be able to pipe command output to a url. The special feature of OTN over other tools that do simialr things is that OTN endpoints can only ever be viewed once and they expire after 5 minutes. This ensure that your information you saved does not live out there on a random endpoing for someone to find. 

I currently host my own instance of OTN which you are free to use at http://otn.maxfus.co.
## Usage
`<command> | curl -F "content=<-" otn.maxfus.co`
