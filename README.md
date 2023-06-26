# Logs
For Saving Logs


# To Build Applicate Docker Command
```
docker-compose build 
```

# To run Application Docker Command
```
docker-compose up 
```

# To run Application Script 
```
DOCKER_COMMAND=test docker-compose up
```

http://drive.google.com/uc?export=view&id=13ob5VLKnNIN0eEuietCT2XGnIL-Lz1Cj
# Application Architecture
![Application Architecture](http://drive.google.com/uc?export=view&id=13ob5VLKnNIN0eEuietCT2XGnIL-Lz1Cj) 


# Sample Cure
```
curl --location 'http://localhost:8000/logs' \
--header 'Content-Type: application/json' \
--data '{
    "id": 12347,
    "unix_ts": 1684129671,
    "user_id": 123456,
    "event_name": "login"
}'
```
