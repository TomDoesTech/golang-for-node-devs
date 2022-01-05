const { database, collections } = require("../db");

async function createProduct(req, res) {
  const collection = database.collection(collections.products);

  const doc = {
    title: req.body.title,
    createdAt: new Date(),
    updatedAt: new Date(),
  };

  const result = await collection.insertOne(doc);

  return res.send(result);
}

async function getProducts(req, res) {
  const collection = database.collection(collections.products);

  const cur = collection.find({});

  let products = [];

  await cur.forEach((p) => {
    products.push(p);
  });

  return res.send(products);
}

module.exports = {
  createProduct,
  getProducts,
};
