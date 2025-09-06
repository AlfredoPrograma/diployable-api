FROM golang:1.25-trixie

RUN apt-get update && apt-get install -y \
    curl \
    zsh \
    && rm -rf /var/lib/apt/lists/*

# Install oh-my-zsh
RUN chsh -s $(which zsh) \
  && sh -c "$(curl -fsSL https://raw.githubusercontent.com/ohmyzsh/ohmyzsh/master/tools/install.sh)"

# Install Go related tooling
RUN go install github.com/air-verse/air@latest \
    && go install github.com/pressly/goose/v3/cmd/goose@latest \
    && go install github.com/sqlc-dev/sqlc/cmd/sqlc@latest

EXPOSE 8080