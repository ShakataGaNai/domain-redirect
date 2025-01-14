# Why?
Because it's fun taking [snofox's nice clean golang](https://github.com/SnoFox/domain-redirect) and making it "better" by converting it to python.

# domain-redirect
A very simple python http server to redirect all subdomains it receives of a domain to a different domain

My reason for writing it: I have muscle memory that I refuse to get rid of for my existing internal LAN names.
However, I am migrating my Kubernetes cluster pods to be exposed via Tailscale rather than directly on the LAN.

So I'm just going to have this service do a temporary redirect for my old domain to my *.ts.net domain until
Tailscale yields and allows custom DNS records *or* I find a new solution after upgrading my Kubernetes cluster.

# Usage
Set the following environment variables:
- `LISTEN_ADDR` - optional; an `ip:port` string of what interface to listen on. By default, all interfaces on port 8080.
- `REDIRECT_DOMAIN` - required; we cut off the last two domain labels and replace with it. Ex, setting `example.com` will redirect `foo.bar.example.local` to `foo.bar.example.com`
