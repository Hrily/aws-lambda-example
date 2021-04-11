function FinancialsTable({ financials }) {
  console.log(financials);
  return (
    <table>
      <tr key="header">
        <th key="year">Year</th>
        <th key="eps">EPS (in Rs.)</th>
      </tr>
      {financials.map((data) => {
        return (
          <tr key={data.year}>
            <th key={data.year}>{data.year}</th>
            <td key={data.eps}>{data.eps}</td>
          </tr>
        );
      })}
    </table>
  );
}

export default FinancialsTable;
