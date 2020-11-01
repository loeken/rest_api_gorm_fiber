# Run
go run ./main.go


# create a release record
curl -H "Content-Type: application/json" localhost:3000/api/v1/release -X POST --data '{"title": "release #1", "artist": "rZumA"}'  

# get all releases
curl localhost:3000/api/v1/release

# get a release by id
curl localhost:3000/api/v1/release


# delete release
curl localhost:3000/api/v1/release/1 -X DELETE

# update release
curl -H "Content-Type: application/json" localhost:3000/api/v1/release/1 -X PUT  --data '{"title": "real release #1", "artist": "RzumA"}' 

