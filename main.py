from flask import Flask, redirect, request
import os
import logging

app = Flask(__name__)

# Get the REDIRECT_DOMAIN environment variable
redir_domain = os.getenv("REDIRECT_DOMAIN")
if not redir_domain:
    raise Exception("Environment variable REDIRECT_DOMAIN must be set!")

# Get the LISTEN_ADDR environment variable or set default
listen_addr = os.getenv("LISTEN_ADDR", "127.0.0.1:8080")

@app.route('/', defaults={'path': ''})
@app.route('/<path:path>')
def redirect_to(path):
    # Splitting the host to get subdomain
    subdomain = '.'.join(request.host.split('.')[:-2])
    
    # Constructing the redirect URL
    redir = f"https://{subdomain}.{redir_domain}/{path}"
    
    # Logging the redirect
    logging.info(f"Got a request: {request.host} looking for /{path}; redirecting to {redir}")

    # Performing the redirect
    return redirect(redir, code=302)

if __name__ == "__main__":
    logging.basicConfig(level=logging.INFO)
    logging.info(f"Listening at {listen_addr}, redirecting to {redir_domain}")
    # Start the Flask application
    host, port = listen_addr.split(":")
    app.run(host=host, port=int(port))
