# OAuth1 Changelog

## v0.3.0 (2015-09-13)

* Added `NoContext` which may be used in most cases.
* Allowed Transport Base http.RoundTripper to be set through a ctx.
* Changed `NewClient` to require a context.Context.
* Changed `Config.Client` to require a context.Context.

## v.0.2.0 (2015-08-30)

* Improved OAuth 1 spec compliance and test coverage.
* Added `func StaticTokenSource(*Token) TokenSource`
* Added `ParseAuthorizationCallback` function. Removed `Config.HandleAuthorizationCallback` method.
* Changed `Config` method signatures to allow an interface to be defined for the OAuth1 authorization flow. Gives users of this package (and downstream packages) the freedom to use other implementations if they wish.
* Removed `RequestToken` in favor of passing token and secret value strings.
* Removed `ReuseTokenSource` struct, it was effectively a static source. Replaced by `StaticTokenSource`.

## v0.1.0 (2015-04-26)

* Initial OAuth1 support for obtaining authorization and making authorized requests.