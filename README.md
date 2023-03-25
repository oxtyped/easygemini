# EasyGemini

EasyGemini is a super-simple gemini static site server that quickly gets you started in bringing your gemini site online.

### Opinionated

EasyGemini currently makes a lot of assumptions and defaults, such as automatically requesting for certificates from letsencrypt.

### Roadmap

- Caddyfile-style configuration mode
- Support for proxying requests to Cloud-based Storage Systems (S3, GCS etc)

#### Features
- 🔐 Handles automatically the cert generation (using lets-encrypt)
- 🌏 Only supports (1) site per process

#### Non-goal
- This server will not be able to support more than 1 site per process

#### Quickstart

1. Pull the entire source code
2. Download the [respective binary](https://github.com/oxtyped/easygemini/releases) and put it in the project repository root
3. Put your gemini files in the `sites/` directory. Replace `index.gmi` as necessary

```
$ ezgem start --site oxtyped.net -b 0.0.0.0 -p 1965
```
---

#### Currently Supports
- Local Filepath (default: `sites/*`)

#### Upcoming
- Amazon S3
- DigitalOcean Spaces
- Hetzner Storage
- GCP GCS
