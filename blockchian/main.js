import crypto from "crypto-js";

class Block {
  constructor(index, timestamp, data, previousHash = "") {
    this.index = index;
    this.timestamp = timestamp;
    this.data = data;
    this.previousHash = previousHash;
    this.hash = this.calcuHash();
    this.nonce = 0;
  }

  calcuHash() {
    return crypto
      .SHA256(
        JSON.stringify(this.data) +
          this.index +
          this.timestamp +
          this.previousHash +
          this.nonce
      )
      .toString();
  }

  mineBlock(difficulty) {
    while (
      this.hash.substring(0, difficulty) !== Array(difficulty + 1).join("0")
    ) {
      this.nonce++;
      this.hash = this.calcuHash();
    }
    console.log("block mined :" + this.hash);
  }
}

class BlockChain {
  constructor() {
    this.chain = [this.createGenisisBlock()];
    this.difficulty = 6; 
  }

  createGenisisBlock() {
    return new Block(0, Date.now(), "GENISIS BLOCK", "-1");
  }

  getLatestBlock() {
    return this.chain[this.chain.length - 1];
  }

  addNewBlock(newBlock) {
    newBlock.previousHash = this.getLatestBlock().hash;
    // newBlock.hash = newBlock.calcuHash();
    newBlock.mineBlock(this.difficulty);
    this.chain.push(newBlock);
  }

  isChainValid() {
    for (let i = 1; i < this.chain.length; i++) {
      const currentBlock = this.chain[i];

      const prevBlock = this.chain[i - 1];
      if (currentBlock.hash !== currentBlock.calcuHash()) {
        return false;
      }

      if (prevBlock.hash !== currentBlock.previousHash) {
        return false;
      }
    }
    return true;
  }
}

const b = new BlockChain();
b.createGenisisBlock();

console.log("mining block 1...");
b.addNewBlock(new Block(1, "1/02/2000", 10));

///////

console.log("mining block 2...");
b.addNewBlock(new Block(1, "4/02/2000", 100));


console.log("Is blockchain valid?: ", b.isChainValid());
