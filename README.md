# mal-cli - A CLI for MyAnimeList

mal-cli is a CLI for MyAnimeList, written in GO. It allows you to search for anime and manga, view details about them, add them to your list, and edit your list.

## Installation

```bash
go install github.com/scandar/mal-cli@latest
```

## Usage

```bash
mal-cli [command]
```

## Authentication

Before using the tool you need to authenticate with MyAnimeList. You can do this by running the following command:

```bash
mal-cli auth
```

## Commands

```bash
anime-details Get anime details
anime-list    Authenticated user's anime list
auth          Authenticate with MyAnimeList
completion    Generate the autocompletion script for the specified shell
delete-anime  Delete an entry from the user's anime list
delete-manga  Delete an entry from the user's manga list
find-anime    Search for anime by name
find-manga    Search for manga by name
help          Help about any command
manga-details Get manga details
manga-list    Authenticated user's manga list
me            Get authenticated user info
update-anime  Update an entry in the user's anime list
update-manga  Update an entry in the user's manga list
```
