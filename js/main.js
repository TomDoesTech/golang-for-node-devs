const express = require("express");
const { connectToDB } = require("./db");
const { createProduct, getProducts } = require("./products");

const app = express();

app.use(express.json());

app.post("/api/products", createProduct);

app.get("/api/products", getProducts);

async function main() {
  app.listen(4000, () => {
    console.log("App listening at http://localhost:4000");
  });

  await connectToDB();

  console.log("Ready for work");
}

/*
 * In Go we won't need to execute main
 */
main();
