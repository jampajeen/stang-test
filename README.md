# stang-test

Steps to test
1. Execute command docker-compose up -d
2. Update rpcurl in file config.yml
3. Execute command go run .
4. You can query transaction by using URL http://localhost:8083/api/addresses/{ETH_ADDRESS} 
eg. http://localhost:8083/api/addresses/0x2Bfd6Cbc525c1e4D32F02a769aeb080DA8C10efa


PS. If you want to test with local ganache you can use my config.yml configuration (you can also install ganache by running yarn install from this project)
