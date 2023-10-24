## Pushing to a remote

1. create tag 
```bash
git tag -a v1.0.0 -m "version 1.0.0"
```
2. push tag
```bash
git push origin v1.0.0
```
3. push tag to remote
```bash
GOPROXY=proxy.golang.org go list -m github.com/oooiik/go-telegram-bot-api@v1.0.0
```