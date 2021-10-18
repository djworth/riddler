const path = require('path')
module.exports = {
  // See <http://truffleframework.com/docs/advanced/configuration>
  // to customize your Truffle configuration!
  networks: {
    development: {
      host: "127.0.0.1",     // Localhost (default: none)
      port: 8545,            // Standard Ethereum port (default: none)
      network_id: "*"       // Any network (default: none)
    }
  },
  compilers: {
    solc: {
      version: "^0.8.0"
    }
  },
  contracts_build_directory: path.join(__dirname, "vapp/src/contracts"),
};
