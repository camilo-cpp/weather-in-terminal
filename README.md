# Weather-in-Terminal

A command-line tool that displays current weather conditions for a specified location using real-time data from an external API. It includes temperature, weather conditions, sunrise and sunset times, and precipitation chances, with emojis for visual representation.

## Features

- **Real-time Weather Data**: Fetches current weather conditions from an external API.
- **User-friendly Output**: Displays temperature, weather conditions, sunrise and sunset times, and precipitation chances.
- **Emoji Representation**: Uses emojis to visually represent different weather conditions.

## Installation

1. Clone the repository:
   ```sh
   git clone https://github.com/camilo-cpp/weather-in-terminal.git
   ```
2. Navigate to the project directory:
   ```sh
   cd weather-in-terminal
   ```
3. Install dependencies:
   ```sh
   go get ./...
   ```

## Usage

Run the application with the following command:

```sh
go run cmd/api/main.go
```

## Global Installation on Windows

To install the application globally on Windows, follow these steps:

1. Build the executable.

```sh
go build -o weather.exe cmd/api/main.go
```

2. Move the executable to the System32 directory.

```sh
move weather.exe C:\Windows\System32\
```

3. After this, you can use the `weather` command in any terminal.

## Example Output

Here is an example of the output you can expect when running the application:

<div align="center">

![Weather Output 1](https://res.cloudinary.com/dpbb73v5v/image/upload/v1724468314/weather/pxxlogns5vrsda68isls.png)

</div>

Please note that the actual values may vary depending on the location and real-time weather data.

<div align="center">

![Weather Output 2](https://res.cloudinary.com/dpbb73v5v/image/upload/v1724468314/weather/zlhdrrdbkadbsshqzjlg.png)

</div>
