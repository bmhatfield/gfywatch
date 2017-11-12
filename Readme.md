# GfyWatch
> Upload Overwatch videos to GfyWatch to share with your friends!

![Tracer Being Annoying](http://giant.gfycat.com/SociableSlowGrunion.webm)

## Install

Either:

`go get -u github.com/bmhatfield/gfywatch`

and then

`go build`

or download from the releases page.

> TODO: No tests lol

## Configure

Currently GfyWatch is configured to watch whatever directory it's running in. You'll also need a file called `grant.json` in that directory with your Gfycat credentials. It looks like this:

```json
{
    "client_id":     "YourGfyCatAPIClient",
    "client_secret": "YourGfyCatAPISecret",
    "username":     "YourGfyCatWebsiteUsername",
    "password":     "YourGfyCatWebsitePassword",
    "grant_type":    "password"
}
```

> TODO: Add `flag` support to GfyWatch and allow the directory and config file to be specified explicitly.

> Note: Gfywatch only watches for `.mp4` files at the moment.

## Naming, Tagging, Etc

GfyWatch will automatically infer as much as it can from the filename you set in Overwatch. The intent is touchless operation - render a POTG in Overwatch, have it show up on GfyCat.

Here's some tips you can follow to get the most out of this:

* It will use the title you give it in Overwatch
* It will also analyze your title for tags based upon keywords
  * It knows about all of Overwatch's heroes and will add additional tags based upon them
  * If you want your video to be marked as a POTG, put POTG in the title.
  * It knows about double, triple, quadruple (etc) and will use those as descriptor words for the tags. (ie; "double kill")
  * It has a few other niceties about the way it tags videos.

> TODO: Additional Tag analysis
> TODO: Optionally cut trailing overwatch bumper.

## Notifications

GfyWatch will log to the `cmd` window that it runs in; however it is built with `0xAX/notificator`. This works great on my Mac, but I haven been unable to propery get it working with GrowlForWindows.

> TODO: Get notifications working on Windows

## Issues / Questions

Please use Github Issues for problems running GfyWatch. Because this is a small side project to streamline sharing my Overwatch plays, PRs are much easier to incorporate than to research and fix issues.