XGB - X Go Binding
==================

This is a fork of the [xgb package](https://code.google.com/p/x-go-binding), a
Go equivalent of the [XCB](http://xcb.freedesktop.org/), the X protocol 
C-language binding.

It's forked mainly to be used as the backend of my widget package 
[sparta](https://github.com/js-arias/sparta).

There is a [more powerful](https://github.com/BurntSushi/xgb) implementation, 
I used it initially, but for some reason it produce BadWindow errors in some
of my machines, so I switch back to the original xgb and made some small 
modifications.

Authorship and license
----------------------

Unless otherwise noted, the XGB source files are distributed
under the BSD-style license found in the LICENSE file.

Contributions should follow the 
[same procedure](http://golang.org/doc/contribute.html) as for the Go project.

