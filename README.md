## Gomock
Mock server written in Go. Mock your API in seconds!
## Installation 
    There are 2 option for installation
    
### Compile from source
    git clone github.com/CommoDor64/gomock
    cd gomock
    go run build OR go run install
    
### Use pre-compiled binary
    git clone github.com/CommoDor64/gomock
    cd bin
    ./gomock <...options>
## Usage
The programme expects API payload models as JSON files  

    app/
        ...app files
        __mock__/
                 users.json
                 kids.json
                 bars.json
It will read the json and generate a RESTful API accordignly  

    API : /users
          /kids
          /bars
## Run
    ./gomock -port=<port_number_default_8001> -dir=<mock_dir_location>
