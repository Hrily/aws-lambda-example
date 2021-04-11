import { ApolloServer } from "apollo-server";
import { getAllCompanies, getCompany, getFinancialsFor } from "./ddb.js";

// 1
const typeDefs = `
	type Query {
		companies: [Company!]
		companyFinancials(symbol: String!): CompanyFinancials!
	}

  type Company {
    symbol: String!
    name: String!
  }

  type Financial {
    year: Int!
    eps: Float!
  }

	type CompanyFinancials {
		company: Company!
    financials: [Financial!]!
	}
`;

// 2
const resolvers = {
  Query: {
    companies: async () => {
      let companies = await getAllCompanies();
      return companies;
    },
    companyFinancials: async (_, args) => {
      let company = await getCompany(args.symbol);
      let financials = await getFinancialsFor(args.symbol);
      return {
        company: company,
        financials: financials,
      };
    },
  },
  Company: {
    symbol: (parent) => parent.symbol,
    name: (parent) => parent.name,
  },
  Financial: {
    year: (parent) => parent.year,
    eps: (parent) => parent.eps,
  },
};

// 3
const server = new ApolloServer({
  typeDefs,
  resolvers,
});

server.listen().then(({ url }) => console.log(`Server is running on ${url}`));
