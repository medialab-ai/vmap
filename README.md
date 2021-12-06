# vmap
## XML VMAP Wrapper

This repo contains a struct with XML tags used to marshal/unmarshal VMAP objects. It was partially generated via goxsd on golang 1.17.3 using the IAB's VMAP xsd file. It references vast 3 and we support 4.2 so I updated the referenced to use 4.2 but the import isn't actually used.

A few types are imported github.com/haxqer/vast since we use that for our VAST parsing. It also has some comments updated from official IAB documentation instead of placeholders to make golint be happy.

# Resources
- https://github.com/ivarg/goxsd
- https://github.com/InteractiveAdvertisingBureau/vmap/blob/master/xsd/vmap.xsd
- https://raw.githubusercontent.com/InteractiveAdvertisingBureau/vast/master/vast_4.2.xsd
- https://github.com/haxqer/vast
