
# go-suffix-map

`go-suffix-map` supports a suffix tree with pointers to aribitrary objects.  The idea is that lookup happens on shared suffix between search key and store key, with specialisations to be supported.

The default specialisation is that suffix matching is on either dot (`.`) separation boundaries, or full string match.

## Examples of default lookup

| Lookup Key | Storage key | Matches? | Comment |
| -----------| ----------- | -------- | ------- |
| subSubSubDomainBeta | domainA.subDomainB.subSubDomainC.subSubSubDomainBeta | true | matches on `.` boundary |
| omainBeta | domainA.subDomainB.subSubDomainC.subSubSubDomainBeta | false | does not match on `.` boundary |
| subSubDomainC.subSubSubDomainBeta | domainA.subDomainB.subSubDomainC.subSubSubDomainBeta | true | matches on `.` boundary |
| ainC.subSubSubDomainBeta | domainA.subDomainB.subSubDomainC.subSubSubDomainBeta | false | does not match on `.` boundary |
| subDomainB.subSubDomainC.subSubSubDomainBeta | domainA.subDomainB.subSubDomainC.subSubSubDomainBeta | true | matches on `.` boundary |
| ainA.subDomainB.subSubDomainC.subSubSubDomainBeta | domainA.subDomainB.subSubDomainC.subSubSubDomainBeta | false | does not match on `.` boundary |
| domainA.subDomainB.subSubDomainC.subSubSubDomainBeta | domainA.subDomainB.subSubDomainC.subSubSubDomainBeta | true | full string match |


## TODO

- Convert to `golang` generics.
- Possibly change core implementation to [Ukkonen's algorithm](https://en.wikipedia.org/wiki/Ukkonen%27s_algorithm).
- Add init config object to support specialisations.
