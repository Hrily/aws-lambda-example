import { useQuery } from "@apollo/client";
import gql from "graphql-tag";
import FinancialsTable from "./FinancialsTable";

const getFinancialsQuery = gql`
  query GetFinancials($symbol: String!) {
    companyFinancials(symbol: $symbol) {
      company {
        name
        symbol
      }
      financials {
        year
        eps
      }
    }
  }
`;

function Financials({ symbol }) {
  const { loading, err, data } = useQuery(getFinancialsQuery, {
    variables: { symbol: symbol },
  });

  if (loading) {
    return <div>Loading Financials...</div>;
  }

  if (err) {
    return (
      <div class="error">Error occured while getting companies: ${err}</div>
    );
  }
  console.log("data", data);

  return (
    <span>
      <div>Company Name: {data.companyFinancials.company.name}</div>
      <div>Company Symbol: {data.companyFinancials.company.symbol}</div>
      <FinancialsTable financials={data.companyFinancials.financials} />
    </span>
  );
}

export default Financials;
