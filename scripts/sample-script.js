
const test = async () => {
    const [signer] = await hre.ethers.getSigners();
    const tx = await signer.sendTransaction({
    // from: signer.address,
    from: "0x2Bfd6Cbc525c1e4D32F02a769aeb080DA8C10efa",
    to: "0xCF07e038e632f2a3F784761c6457B7cCD8F01de6",
    value: hre.ethers.utils.parseEther("0.1"),
    });

    console.log("tx", tx);
}

test();
