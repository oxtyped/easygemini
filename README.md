# EasyGemini

EasyGemini is a super-simple [gemini](https://gemini.circumlunar.space/) static site server that quickly gets you started in bringing your gemini site online.

### Opinionated

EasyGemini currently makes a lot of assumptions and defaults, such as automatically generating a self-signed certificate and loading it for you. At its core, EasyGemini is powered by [adano/go-gemini](https://git.sr.ht/~adnano/go-gemini).

### Roadmap

- Caddyfile-style configuration mode
- Support for proxying requests to Cloud-based Storage Systems (S3, GCS etc)
- CGI-BIN support (Ref: [RFC3875](https://www.rfc-editor.org/rfc/rfc3875))

#### Features
- 🔐 Handles the cert generation automatically
- 🌏 Only supports (1) site per process

#### Non-goal
- This server will NOT be able to support more than 1 site per process. In order to run multiple sites in a single host, it is advisable to use a loadbalancer with SNI support (eg. Haproxy) to proxy the requests to the right process.

#### Quickstart

1. Pull the entire source code
2. Download the [respective binary](https://github.com/oxtyped/easygemini/releases) and put it in the project repository root
3. Put your gemini files in the `sites/` directory. Replace `index.gmi` as necessary

```
$ easygemini serve
```

4. Grab a [gemini browser](https://geminiquickst.art/) and head on over to `gemini://localhost:1965`

---

#### Currently Supports
- Local Filepath (default: `sites/*`)

#### Upcoming
- Amazon S3
- DigitalOcean Spaces
- Hetzner Storage
- GCP GCS
