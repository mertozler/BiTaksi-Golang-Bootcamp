# BiTaksi Golang Bootcamp Project

A Driver Location API that uses location data stored in a MongoDB collection.
A Matching API Service that finds the nearest driver with the rider location using the Driver Location API.


## Installation

The project is currently connected to MongoDB Atlas and can be accessed via any ip address. (This will be removed when the project is made public.) If the cluster fails or you want to try it in your local, you can edit the .env file. 

DriveLocation-API/.env
```bash
MONGOURI=Write your mongo URI here
```

## How To Use?
First you need to run the DriverLocation API. In order to provide a match, the rider's location is taken from the matching api and sent to the driver location api. First, the driver location api checks whether there is data in the database, if there is no data, it takes the driver locations from the coordinates.csv file and transforms it into a GEOJson object and assigns it to MongoDB.
After the DriverLocation API is running, you can run the Matching API. In order to request from Matching, you need to have JWT authorization. (Bearer) If there is no JWT key, the desired result will not be obtained, since auth cannot be provided. Key required for this:

```
eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhdXRoZW50aWNhdGVkIjp0cnVlfQ.iG8Rux4vSqgoBvG2OMggjK9Q5QGyvATykfH8qKbJTAs
```

## Documentation
For documentation of Driverlocation-API go to http://localhost:7000/swagger/, for documentation of Matching-API go to http://localhost:8000/swagger/.


## Contributing
Pull requests are welcome. For major changes, please open an issue first to discuss what you would like to change.

Please make sure to update tests as appropriate.

## License
[MIT](https://choosealicense.com/licenses/mit/)
