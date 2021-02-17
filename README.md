# connector-accuweather

Connector used to provide weather products (weather today and forecast) from  Accuweather

## Building

```
docker build -t pot-accuweather .
```

## Running

NOTE THAT THE **ACCUWEATHER_TOKEN** ENVIRONMENT VARIABLE HAS TO BE SET FOR THE CONNECTOR TO WORK!

The public key to validate the response signature is provided from the following path: `/public.key`

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

# A file path to a private key to use when signing the responses from the connector.
# The public key will be generated from the private key so no need to provide it.
# If not provided a private key will be generated when the connector starts.
PRIVATE_KEY={/path/to/private.key}

# The key that will be used for the payload signature verification.
# Can be either a url or a file path. To use an url it must start with either http:// or https://.
# Will fallback for predefined list of keys if not provided.
POT_PUBLIC_KEY={platform of trust public key}

# The URI of the context that will be used for the data payload 
POT_RESPONSE_CONTEXT=https://standards.oftrust.net/v2/Context/DataProductOutput/Forecast/Weather/AccuWeather/

# The URI of the context that will be used for the parameters payload 
POT_PARAMETER_CONTEXT=https://standards.oftrust.net/v2/Context/DataProductParameters/Forecast/Weather/AccuWeather/
```
