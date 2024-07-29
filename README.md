# Install Pipx
brew install pipx
pipx ensurepath


# Install pre-commit
pipx install pre-commit

# Run it without committing

pre-commit run --all-files --verbose

# Autoupdate

pre-commit autoupdate

# Clean unused hooks

pre-commit clean

# Install pre-commit script

pre-commit install

# Uninstall pre-commit script

pre-commit uninstall



# Install golangci-lint

brew install golangci-lint
brew upgrade golangci-lint

# Install gobrew
curl -sL https://raw.githubusercontent.com/kevincobain2000/gobrew/master/git.io.sh | sh


# Ignore linting issues

//nolint:forbidigo,staticcheck

or

//nolint:all



Long line

	fmt.Println("fff", 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 171, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 171, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 171, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17) //nolint:forbidigo

myMap := map[string]string{"first key": "first value", "second key": "second value", "third key": "third value", "fourth key": "fourth value", "fifth key": "fifth value"}