FROM golang:1.22-alpine AS build
ENV CGO_ENABLED="0"

# These are all the tools installed by the "Go: Install/Update Tools"
# command, plus gofumpt. gofumpt is a stricter gofmt that enforces additional
# formatting rules for better consistency and readabilty.
RUN --mount=type=cache,target=/root/.cache/go-build \
    --mount=type=cache,target=/go/pkg \
    <<-EOT
    apk --no-cache add curl
    go install github.com/cweill/gotests/gotests@latest
    go install github.com/fatih/gomodifytags@latest
    go install github.com/josharian/impl@latest
    go install github.com/haya14busa/goplay/cmd/goplay@latest
    go install github.com/go-delve/delve/cmd/dlv@latest
    go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest
    go install golang.org/x/tools/gopls@latest
    go install honnef.co/go/tools/cmd/staticcheck@latest
    go install mvdan.cc/gofumpt@latest
EOT


FROM ghcr.io/ryboe/alpinecodespace:latest
ENV PATH="/home/vscode/go/bin:/usr/local/go/bin:${PATH}"
COPY --from=build /usr/local/go/ /usr/local/go/
COPY --from=build /go/bin/ /home/vscode/go/bin
RUN sudo mkdir /home/vscode/go/pkg
RUN sudo chown -R vscode /home/vscode/go
