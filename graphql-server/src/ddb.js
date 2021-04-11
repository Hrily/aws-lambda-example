// Load the AWS SDK for Node.js
import pkg from "aws-sdk";
const { config, DynamoDB } = pkg;

const companiesTable = "companies";
const financialsTable = "financials";

// Set the region
config.update({
  accessKeyId: "test",
  secretAccessKey: "test",
  region: "us-east-1",
});

// Create the DynamoDB service object
var ddb = new DynamoDB({
  endpoint: "http://localstack:4566",
  apiVersion: "2012-08-10",
});

function getAllCompanies() {
  return new Promise((resolve, reject) => {
    ddb.scan(
      {
        TableName: companiesTable,
      },
      function (err, data) {
        if (err) {
          console.log("Error", err);
          reject(err);
        } else {
          let companies = [];
          for (let item of data.Items) {
            companies.push({
              symbol: item.Symbol.S,
              name: item.Name.S,
            });
          }
          resolve(companies);
        }
      }
    );
  });
}

function getCompany(symbol) {
  return new Promise((resolve, reject) => {
    let params = {
      TableName: companiesTable,
      Key: {
        Symbol: { S: symbol },
      },
    };
    ddb.getItem(params, function (err, data) {
      if (err) {
        console.log("Error", err);
        reject(err);
      } else {
        let company = {
          symbol: data.Item.Symbol.S,
          name: data.Item.Name.S,
        };
        resolve(company);
      }
    });
  });
}

// getFinancialsFor a company symbol
// gets all the financial data for all years available
// TODO: add year level filter
function getFinancialsFor(symbol) {
  console.log(`getting financials for ${symbol}`);
  return new Promise((resolve, reject) => {
    let params = {
      ExpressionAttributeValues: {
        ":s": { S: symbol },
      },
      KeyConditionExpression: "Symbol = :s",
      TableName: financialsTable,
    };
    ddb.query(params, function (err, data) {
      if (err) {
        console.log("Error", err);
        reject(err);
      } else {
        let financials = [];
        for (let item of data.Items) {
          financials.push({
            year: item.Year.N,
            eps: item.EPS.N,
          });
        }
        resolve(financials);
      }
    });
  });
}

export { getAllCompanies, getCompany, getFinancialsFor };
