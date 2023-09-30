// pages/index.tsx

import { useQuery } from "react-query";
import React from "react";
// import useDebounce from "../utils/useDebounce";
// import searchPokemons from "../utils/searchPokemons";

export default function IndexPage() {
  const [searchValue, setSearchValue] = React.useState("");
//   const debounedSearchValue = useDebounce(searchValue, 300);

//   const { isLoading, isError, isSuccess, data } = useQuery(
//     // ["searchPokemons", debounedSearchValue],
//     // () => searchPokemons(debounedSearchValue),
//   );

  return (
    <div className="home">
      <h1>Search Your Pokemon</h1>
      <input
        type="text"
        onChange={({ target: { value } }) => setSearchValue(value)}
        value={searchValue}
      />
    </div>
  );
}