[![CircleCI](https://circleci.com/gh/broderickhyman/albiondata-client/tree/master.svg?style=svg)](https://circleci.com/gh/broderickhyman/albiondata-client/tree/master) [![Go Report Card](https://goreportcard.com/badge/github.com/broderickhyman/albiondata-client)](https://goreportcard.com/report/github.com/broderickhyman/albiondata-client)

# Albion Data - Client
Distributed client for the [Albion Online Data](https://www.albion-online-data.com/)
project.

A quick note on the legality of this application and if it
violates the Terms and Conditions for Albion Online. Here is
the response from SBI when asked if we are allowed to do
monitor network packets relating to Albion Online:
> Our position is quite simple. As long as you just look and
analyze we are ok with it. The moment you modify or manipulate
something or somehow interfere with our services we will react
(e.g. perma-ban, take legal action, whatever).

~ MadDave - Technical Lead for Albion Online

Source: https://forum.albiononline.com/index.php/Thread/51604-Is-it-allowed-to-scan-your-internet-trafic-and-pick-up-logs/?postID=512670#post512670

This client monitors local network traffic, identifies UDP packets
that contain relevant data for Albion Online, and ships the information
off to a central NATS server that anyone can subscribe to.

[Client download stats](https://www.somsubhra.com/github-release-stats/?username=broderickhyman&repository=albiondata-client)

### Contributing
This process is run on a [DigitalOcean Droplet](https://www.digitalocean.com) in order to ensure almost perfect uptime and high performance for the users. If you find this project beneficial to you then please consider a donation, thanks!!

[Become a Patron on Patreon!](https://www.patreon.com/bePatron?u=10422119)

[![ko-fi](https://www.ko-fi.com/img/donate_sm.png)](https://ko-fi.com/E1E5K69V)

# Contributions
Many thanks to the original developers:
- [Regner](https://github.com/Regner)
- [pcdummy](https://github.com/pcdummy)
- [Ultraporing](https://github.com/Ultraporing)

# Downloads
Downloads can be found here: https://github.com/broderickhyman/albiondata-client/releases

# Related Projects
- [albiondata-deduper-dotNet](https://github.com/BroderickHyman/albiondata-deduper-dotNet)
- [albiondata-sql-dotNet](https://github.com/BroderickHyman/albiondata-sql-dotNet)
- [albiondata-api-dotNet](https://github.com/BroderickHyman/albiondata-api-dotNet)
- [AlbionData.Models](https://github.com/broderickhyman/albiondata-models-dotNet) [![NuGet](https://img.shields.io/nuget/v/AlbionData.Models.svg)](https://www.nuget.org/packages/AlbionData.Models/)
- [albion-data-website](https://github.com/broderickhyman/albion-data-website)

# Contact Us
The best way to get in touch with us is on the Albion Online Fansites Discord server in either the #proj-albiondata or the #developers channel. A permanent invite link can be found here: [https://discord.gg/TjWdq24](https://discord.gg/TjWdq24)

# Developer Setup
### Mac/Linux Setup
- Install [Dep](https://github.com/golang/dep)
  - Any OS: `go get -u github.com/golang/dep/cmd/dep`
  - Mac with Homebrew: `brew install dep`
- Install dependencies using `dep ensure`

### Windows Setup
[Windows Setup Guide](https://github.com/broderickhyman/albiondata-client/wiki/Building-in-Windows)

# License
This project, and all contributed code, are licensed under the MIT
License. A copy of the MIT License may be found in the repository.
