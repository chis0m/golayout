# How Sessions Work with Token-Based Authentication

## Initial Authentication:

- Upon successful authentication (e.g., username and password check), the server issues a JWT access token and creates a session for the user. This session contains necessary user information and permissions.
- The session ID can be sent back to the client, typically within a secure, HTTPOnly cookie, reducing the risk of XSS attacks.

## Using Access Tokens with Sessions:

- The client sends the JWT access token in the HTTP header for accessing protected resources. The session ID is automatically sent with each request via cookies.
- The server validates the JWT as usual for authentication and then checks the session ID to verify that the session exists and is valid. This two-step process enhances security by ensuring the token is not only valid but also corresponds to a valid, active session.

## Session Expiry and Invalidation:

- Sessions can be configured with a timeout, automatically expiring after a period of inactivity, which requires the user to re-authenticate.
- If a user logs out or if suspicious activity is detected, the server can immediately invalidate the session, rendering the session ID (and by extension, the JWT) useless for further requests, even if the JWT itself has not expired.

## Refresh Tokens and Sessions:

- Instead of managing refresh tokens in a separate datastore, the session itself can serve as the basis for issuing new access tokens. When an access token expires, the presence of a valid session can allow the server to issue a new access token without requiring the user to log in again.
