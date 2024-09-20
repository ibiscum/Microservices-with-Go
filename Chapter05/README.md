## Generate proto buffer code

    protoc -I=./api --go_out=. --go-grpc_out=. movie.proto

## Get details of gRPC services

### Metadata service

    grpcurl -plaintext localhost:8081 list
    grpcurl -plaintext localhost:8081 describe MetadataService
    grpcurl -plaintext localhost:8081 describe MetadataService.GetMetadata
    grpcurl -plaintext localhost:8081 describe .GetMetadataRequest

### Request metadata entry

    grpcurl -plaintext -d '{"movie_id": "1"}' localhost:8081 MetadataService/GetMetadata

### Rating service

    grpcurl -plaintext localhost:8082 list
    grpcurl -plaintext localhost:8082 describe RatingService
    grpcurl -plaintext localhost:8082 describe RatingService.GetAggregatedRating
    grpcurl -plaintext localhost:8082 describe .GetAggregatedRatingRequest

    grpcurl -plaintext localhost:8082 describe RatingService.PutRating
    grpcurl -plaintext localhost:8082 describe .PutRatingRequest

### Request rating entry

    grpcurl -plaintext -d '{"record_id": "1", "record_type": "movie"}' localhost:8082 RatingService/GetAggregatedRating

### Add rating entry

    grpcurl -plaintext -d '{"user_id": "u5", "record_id": "1", "record_type": "movie", "rating_value": 9}' localhost:8082 RatingService/PutRating

### Movie service

    grpcurl -plaintext localhost:8083 list
    grpcurl -plaintext localhost:8083 describe MovieService
    grpcurl -plaintext localhost:8083 describe MovieService.GetMovieDetails
    grpcurl -plaintext localhost:8083 describe .GetMovieDetailsRequest

### Request movie details

    grpcurl -plaintext -d '{"movie_id": "1"}' localhost:8083 MovieService/GetMovieDetails

