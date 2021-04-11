import { gql, useQuery } from "@apollo/client";
import { DEFAULT_SYMBOL } from "../constants";

const companiesQuery = gql`
  query {
    companies {
      symbol
      name
    }
  }
`;

const placeholderCompany = {
  name: "-- Select a company --",
  symbol: DEFAULT_SYMBOL,
};

function CompaniesSelector({ setSymbol }) {
  const { loading, err, data } = useQuery(companiesQuery);

  if (loading) {
    return <div>Loading companies...</div>;
  }

  if (err) {
    return (
      <div class="error">Error occured while getting companies: ${err}</div>
    );
  }

  const companies = [placeholderCompany, ...data.companies];

  const onInputChange = function (event) {
    setSymbol(event.target.value);
  };

  return (
    <form>
      <label htmlFor="companies">Choose a company: </label>
      <select onChange={onInputChange} name="companies" id="companies">
        {companies.map((company) => {
          return (
            <option key={company.symbol} value={company.symbol}>
              {company.name}
            </option>
          );
        })}
      </select>
    </form>
  );
}

export default CompaniesSelector;
