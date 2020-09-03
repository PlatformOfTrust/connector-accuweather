# connector-accuweather

Connector used to provide weather products (weather today and forecast) from  Accuweather

## Building

```
docker build -t pot-accuweather .
```

## Running

NOTE THAT THE **ACCUWEATHER_TOKEN** AND **POT_SECRET** ENVIRONMENT VARIABLES HAS TO BE SET FOR THE CONNECTOR TO WORK!

### Basic
```
docker run -p 8080:8080 pot-accuweather
```

### Environment mounted from a .env file
```
docker run -v $(pwd)/.env:/connector/.env -p 8080:8080 pot-accuweather
```

### Custom port
```
docker run -e PORT=8888 -p 8888:8888 pot-accuweather
```

## Configuration

### Environment variables with their default values and descriptions

```bash
# The port that the server will start listening on
PORT=8080

# The access token that will be used to fetch data from the accuweather APIs
ACCUWEATHER_TOKEN={get the token from your accuweather account}

# The secret that will be used for the payload signing and validation
POT_SECRET={platform of trust secret for the connector}

# The URI of the public key that will be used for the signature payload of the response
POT_CREATOR=https://example.com/public-key

# The URI of the context that will be used for the data payload 
POT_RESPONSE_CONTEXT=https://standards.oftrust.net/v2/Context/DataProductOutput/Forecast/Weather/AccuWeather/

# The URI of the context that will be used for the parameters payload 
POT_PARAMETER_CONTEXT=https://standards.oftrust.net/v2/Context/DataProductParameters/Forecast/Weather/AccuWeather/
```
