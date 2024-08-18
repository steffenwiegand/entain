## Review document for Entain BE Technical Test

The purpose of this document is to give some feedback and additional thoughts on the coding test.

### Overall Feedback
To the reviewer. I am coming from a lifelong .NET background and am basically new to golang. 
The biggest challenge for me was setup and configuration which chewed up quite a bit of time at the start. For context, I wasn't really able to get the solution up and running properly on my windows machine and ran into a fair few issues with some packages, specifically the protocol buffer (package proto). Once I went down the path of getting setup inside an oracle virtual box with Ubuntu everything got a lot better. Again more setup work involed to get it all setup and running. 
I was able to get through all 5 tasks incl. testing. From a code structure and overall golang best practises point of view there may be things I am not yet aware of that might be missing. Keen to receive feedback on this one.

### Future changes and considerations to the submitted solution

I would like to highlight some areas that in my view woud benefit with further time spent on them.

1. I would like to see the unit tests in both the sports- and races service with more test cases and expansion on the test data provided.

2. Task2 and SQL Injection. Personally, when we talk about production readiness it is obviously a risk having a consumer pass in a string to sort the result of a recesList or eventsList. I have put in some protection against this via a regular expression, but I would change this to an enum that is passed in with a strict choice of columns a consumer is allowed to sort by in whichever preferred order a user chooses.




