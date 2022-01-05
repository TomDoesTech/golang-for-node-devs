const { MongoClient } = require("mongodb");

const URI = "mongodb://localhost:27017";

const databaseStr = "products-api";

const collections = {
  products: "products",
};

const client = new MongoClient(`${URI}`);

async function connectToDB() {
  try {
    await client.connect();
    console.log("Database connected");
  } catch {
    await client.close();
    console.error("Could not connect to database");
    process.exit(1);
  }
}

const database = client.db(databaseStr);

module.exports = {
  database,
  collections,
  connectToDB,
};
