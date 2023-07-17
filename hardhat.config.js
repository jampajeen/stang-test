/** @type import('hardhat/config').HardhatUserConfig */

require("@nomiclabs/hardhat-waffle");

module.exports = {
  solidity: "0.8.18",
  networks: {
    ganache: {
      url: "http://127.0.0.1:8545",
      accounts: [
        "0x4c98187970f46181fb35e19811f4a4b2b80de97fd95fb8ab9989a94eedb416aa",
      ],
    }
  }
};
