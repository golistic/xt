xt - Minimal Test Functions Wrapping Go's testing
=================================================

Copyright (c) 2022, Geert JM Vanderkelen

The Go xt package offers a minimal set of functions helping writing tests;
wrapping around Go testing package. This is not as sophisticated as others,
but it covers the basics, and that is usually enough.

Overview
--------

* `OK(t *testing.T, err error)`
* `KO(t *testing.T, err error)`
* `Assert(t *testing.T, condition bool, messages ...string)`
* `Eq(t *testing.T, expect, have any, messages ...string)`
* `MatchString(t *testing.T, pattern, s string, messages ...string)`
* `Panics(t *testing.T, f func())`
* `PanicsEq(t *testing.T, exp string, f func())`


License
-------

Distributed under the MIT license. See LICENSE.txt for more information.
