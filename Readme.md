# GfyWatch
> Upload Overwatch videos to GfyWatch to share with your friends!

This is a very small program that sits in the background and watches your Overwatch "videos" directory (on my machine, `Documents\Overwatch\videos\overwatch`), uploading new files as they arrive.

I wrote this to streamline my personal workflow around "get POTG, manually upload to Gfycat".

This tool is curently _very beta_ - I have it working well for me, but your feedback would be most valuable.

Read on for more information about features to come.

## Install

Either:

`go get -u github.com/bmhatfield/gfywatch`

and then

`go build`

or download from the releases page.

> TODO: This tool needs _way_ more tests.

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

You'll need (for now) API credentials from Gfycat: https://developers.gfycat.com/signup/#/apiform - I realize this is annoying but is currently required.

> TODO: Make it easier to sign in to GfyWatch

> TODO: Add `flag` support to GfyWatch and allow the directory and config file to be specified explicitly.

> Note: Gfywatch only watches for `.mp4` files at the moment - it could also support `.webm` trivially.

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