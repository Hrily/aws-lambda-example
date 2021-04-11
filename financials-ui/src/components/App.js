import CompaniesSelector from "./CompaniesSelector";
import "./../styles/App.css";
import { DEFAULT_SYMBOL } from "../constants";
import { useState } from "react";
import Financials from "./Financials";

function App() {
  const [symbol, setSymbol] = useState(DEFAULT_SYMBOL);

  const renderFinancials = () => {
    if (symbol !== DEFAULT_SYMBOL) {
      return <Financials symbol={symbol} />;
    }
  };

  return (
    <div className="App">
      <header className="App-header">
        <CompaniesSelector setSymbol={setSymbol} />
        {renderFinancials()}
      </header>
    </div>
  );
}

export default App;
