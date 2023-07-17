# stang-test

### <u>Steps to test</u>
1. Update rpcurl in file config.yml
2. Execute command docker-compose up
3. You can query transaction by using URL http://localhost:8083/api/addresses/{ETH_ADDRESS} 
eg. http://localhost:8083/api/addresses/0x2Bfd6Cbc525c1e4D32F02a769aeb080DA8C10efa

PS. If you want to test with local ganache you can use my config.yml configuration (you can also install ganache by running yarn install from this project)


### <u>Architecture design</u>
![alt text](https://raw.githubusercontent.com/jampajeen/stang-test/main/stang_test.drawio.png?token=GHSAT0AAAAAACETK5R2PIHJNMFTPGSQYYXSZFVIK4Q)